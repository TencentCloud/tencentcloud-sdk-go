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

package v20230418

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-04-18"

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


func NewAccessDevicesRequest() (request *AccessDevicesRequest) {
    request = &AccessDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "AccessDevices")
    
    
    return
}

func NewAccessDevicesResponse() (response *AccessDevicesResponse) {
    response = &AccessDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AccessDevices
// 外部客户访问资产
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AccessDevices(request *AccessDevicesRequest) (response *AccessDevicesResponse, err error) {
    return c.AccessDevicesWithContext(context.Background(), request)
}

// AccessDevices
// 外部客户访问资产
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AccessDevicesWithContext(ctx context.Context, request *AccessDevicesRequest) (response *AccessDevicesResponse, err error) {
    if request == nil {
        request = NewAccessDevicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "AccessDevices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AccessDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewAccessDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewAddDeviceGroupMembersRequest() (request *AddDeviceGroupMembersRequest) {
    request = &AddDeviceGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "AddDeviceGroupMembers")
    
    
    return
}

func NewAddDeviceGroupMembersResponse() (response *AddDeviceGroupMembersResponse) {
    response = &AddDeviceGroupMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddDeviceGroupMembers
// 添加资产组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddDeviceGroupMembers(request *AddDeviceGroupMembersRequest) (response *AddDeviceGroupMembersResponse, err error) {
    return c.AddDeviceGroupMembersWithContext(context.Background(), request)
}

// AddDeviceGroupMembers
// 添加资产组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddDeviceGroupMembersWithContext(ctx context.Context, request *AddDeviceGroupMembersRequest) (response *AddDeviceGroupMembersResponse, err error) {
    if request == nil {
        request = NewAddDeviceGroupMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "AddDeviceGroupMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddDeviceGroupMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddDeviceGroupMembersResponse()
    err = c.Send(request, response)
    return
}

func NewAddUserGroupMembersRequest() (request *AddUserGroupMembersRequest) {
    request = &AddUserGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "AddUserGroupMembers")
    
    
    return
}

func NewAddUserGroupMembersResponse() (response *AddUserGroupMembersResponse) {
    response = &AddUserGroupMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUserGroupMembers
// 添加用户组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddUserGroupMembers(request *AddUserGroupMembersRequest) (response *AddUserGroupMembersResponse, err error) {
    return c.AddUserGroupMembersWithContext(context.Background(), request)
}

// AddUserGroupMembers
// 添加用户组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddUserGroupMembersWithContext(ctx context.Context, request *AddUserGroupMembersRequest) (response *AddUserGroupMembersResponse, err error) {
    if request == nil {
        request = NewAddUserGroupMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "AddUserGroupMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserGroupMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserGroupMembersResponse()
    err = c.Send(request, response)
    return
}

func NewBindDeviceAccountPasswordRequest() (request *BindDeviceAccountPasswordRequest) {
    request = &BindDeviceAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "BindDeviceAccountPassword")
    
    
    return
}

func NewBindDeviceAccountPasswordResponse() (response *BindDeviceAccountPasswordResponse) {
    response = &BindDeviceAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindDeviceAccountPassword
// 绑定主机账号密码
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BindDeviceAccountPassword(request *BindDeviceAccountPasswordRequest) (response *BindDeviceAccountPasswordResponse, err error) {
    return c.BindDeviceAccountPasswordWithContext(context.Background(), request)
}

// BindDeviceAccountPassword
// 绑定主机账号密码
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BindDeviceAccountPasswordWithContext(ctx context.Context, request *BindDeviceAccountPasswordRequest) (response *BindDeviceAccountPasswordResponse, err error) {
    if request == nil {
        request = NewBindDeviceAccountPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "BindDeviceAccountPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDeviceAccountPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindDeviceAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewBindDeviceAccountPrivateKeyRequest() (request *BindDeviceAccountPrivateKeyRequest) {
    request = &BindDeviceAccountPrivateKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "BindDeviceAccountPrivateKey")
    
    
    return
}

func NewBindDeviceAccountPrivateKeyResponse() (response *BindDeviceAccountPrivateKeyResponse) {
    response = &BindDeviceAccountPrivateKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindDeviceAccountPrivateKey
// 绑定主机账号私钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BindDeviceAccountPrivateKey(request *BindDeviceAccountPrivateKeyRequest) (response *BindDeviceAccountPrivateKeyResponse, err error) {
    return c.BindDeviceAccountPrivateKeyWithContext(context.Background(), request)
}

// BindDeviceAccountPrivateKey
// 绑定主机账号私钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BindDeviceAccountPrivateKeyWithContext(ctx context.Context, request *BindDeviceAccountPrivateKeyRequest) (response *BindDeviceAccountPrivateKeyResponse, err error) {
    if request == nil {
        request = NewBindDeviceAccountPrivateKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "BindDeviceAccountPrivateKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDeviceAccountPrivateKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindDeviceAccountPrivateKeyResponse()
    err = c.Send(request, response)
    return
}

func NewBindDeviceResourceRequest() (request *BindDeviceResourceRequest) {
    request = &BindDeviceResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "BindDeviceResource")
    
    
    return
}

func NewBindDeviceResourceResponse() (response *BindDeviceResourceResponse) {
    response = &BindDeviceResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindDeviceResource
// 修改资产绑定的堡垒机服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindDeviceResource(request *BindDeviceResourceRequest) (response *BindDeviceResourceResponse, err error) {
    return c.BindDeviceResourceWithContext(context.Background(), request)
}

// BindDeviceResource
// 修改资产绑定的堡垒机服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindDeviceResourceWithContext(ctx context.Context, request *BindDeviceResourceRequest) (response *BindDeviceResourceResponse, err error) {
    if request == nil {
        request = NewBindDeviceResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "BindDeviceResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDeviceResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindDeviceResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCheckLDAPConnectionRequest() (request *CheckLDAPConnectionRequest) {
    request = &CheckLDAPConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CheckLDAPConnection")
    
    
    return
}

func NewCheckLDAPConnectionResponse() (response *CheckLDAPConnectionResponse) {
    response = &CheckLDAPConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckLDAPConnection
// 测试LDAP连接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAILED = "FailedOperation.AuthFailed"
//  FAILEDOPERATION_CONNECTIONFAILED = "FailedOperation.ConnectionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckLDAPConnection(request *CheckLDAPConnectionRequest) (response *CheckLDAPConnectionResponse, err error) {
    return c.CheckLDAPConnectionWithContext(context.Background(), request)
}

// CheckLDAPConnection
// 测试LDAP连接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAILED = "FailedOperation.AuthFailed"
//  FAILEDOPERATION_CONNECTIONFAILED = "FailedOperation.ConnectionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckLDAPConnectionWithContext(ctx context.Context, request *CheckLDAPConnectionRequest) (response *CheckLDAPConnectionResponse, err error) {
    if request == nil {
        request = NewCheckLDAPConnectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CheckLDAPConnection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckLDAPConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckLDAPConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessWhiteListRuleRequest() (request *CreateAccessWhiteListRuleRequest) {
    request = &CreateAccessWhiteListRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateAccessWhiteListRule")
    
    
    return
}

func NewCreateAccessWhiteListRuleResponse() (response *CreateAccessWhiteListRuleResponse) {
    response = &CreateAccessWhiteListRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccessWhiteListRule
// 添加访问白名单规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAccessWhiteListRule(request *CreateAccessWhiteListRuleRequest) (response *CreateAccessWhiteListRuleResponse, err error) {
    return c.CreateAccessWhiteListRuleWithContext(context.Background(), request)
}

// CreateAccessWhiteListRule
// 添加访问白名单规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAccessWhiteListRuleWithContext(ctx context.Context, request *CreateAccessWhiteListRuleRequest) (response *CreateAccessWhiteListRuleResponse, err error) {
    if request == nil {
        request = NewCreateAccessWhiteListRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateAccessWhiteListRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccessWhiteListRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccessWhiteListRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAclRequest() (request *CreateAclRequest) {
    request = &CreateAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateAcl")
    
    
    return
}

func NewCreateAclResponse() (response *CreateAclResponse) {
    response = &CreateAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAcl
// 新建访问权限
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAcl(request *CreateAclRequest) (response *CreateAclResponse, err error) {
    return c.CreateAclWithContext(context.Background(), request)
}

// CreateAcl
// 新建访问权限
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAclWithContext(ctx context.Context, request *CreateAclRequest) (response *CreateAclResponse, err error) {
    if request == nil {
        request = NewCreateAclRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateAcl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAclResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetSyncJobRequest() (request *CreateAssetSyncJobRequest) {
    request = &CreateAssetSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateAssetSyncJob")
    
    
    return
}

func NewCreateAssetSyncJobResponse() (response *CreateAssetSyncJobResponse) {
    response = &CreateAssetSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAssetSyncJob
// 创建手工资产同步任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAILED = "FailedOperation.AuthFailed"
//  FAILEDOPERATION_CONNECTIONFAILED = "FailedOperation.ConnectionFailed"
//  FAILEDOPERATION_TOOFREQUENT = "FailedOperation.TooFrequent"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAssetSyncJob(request *CreateAssetSyncJobRequest) (response *CreateAssetSyncJobResponse, err error) {
    return c.CreateAssetSyncJobWithContext(context.Background(), request)
}

// CreateAssetSyncJob
// 创建手工资产同步任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAILED = "FailedOperation.AuthFailed"
//  FAILEDOPERATION_CONNECTIONFAILED = "FailedOperation.ConnectionFailed"
//  FAILEDOPERATION_TOOFREQUENT = "FailedOperation.TooFrequent"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAssetSyncJobWithContext(ctx context.Context, request *CreateAssetSyncJobRequest) (response *CreateAssetSyncJobResponse, err error) {
    if request == nil {
        request = NewCreateAssetSyncJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateAssetSyncJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChangePwdTaskRequest() (request *CreateChangePwdTaskRequest) {
    request = &CreateChangePwdTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateChangePwdTask")
    
    
    return
}

func NewCreateChangePwdTaskResponse() (response *CreateChangePwdTaskResponse) {
    response = &CreateChangePwdTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateChangePwdTask
// 创建修改密码任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAILED = "FailedOperation.AuthFailed"
//  FAILEDOPERATION_CONNECTIONFAILED = "FailedOperation.ConnectionFailed"
//  FAILEDOPERATION_TOOFREQUENT = "FailedOperation.TooFrequent"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateChangePwdTask(request *CreateChangePwdTaskRequest) (response *CreateChangePwdTaskResponse, err error) {
    return c.CreateChangePwdTaskWithContext(context.Background(), request)
}

// CreateChangePwdTask
// 创建修改密码任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAILED = "FailedOperation.AuthFailed"
//  FAILEDOPERATION_CONNECTIONFAILED = "FailedOperation.ConnectionFailed"
//  FAILEDOPERATION_TOOFREQUENT = "FailedOperation.TooFrequent"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateChangePwdTaskWithContext(ctx context.Context, request *CreateChangePwdTaskRequest) (response *CreateChangePwdTaskResponse, err error) {
    if request == nil {
        request = NewCreateChangePwdTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateChangePwdTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChangePwdTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChangePwdTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCmdTemplateRequest() (request *CreateCmdTemplateRequest) {
    request = &CreateCmdTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateCmdTemplate")
    
    
    return
}

func NewCreateCmdTemplateResponse() (response *CreateCmdTemplateResponse) {
    response = &CreateCmdTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCmdTemplate
// 新建高危命令模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCmdTemplate(request *CreateCmdTemplateRequest) (response *CreateCmdTemplateResponse, err error) {
    return c.CreateCmdTemplateWithContext(context.Background(), request)
}

// CreateCmdTemplate
// 新建高危命令模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCmdTemplateWithContext(ctx context.Context, request *CreateCmdTemplateRequest) (response *CreateCmdTemplateResponse, err error) {
    if request == nil {
        request = NewCreateCmdTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateCmdTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCmdTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCmdTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeviceAccountRequest() (request *CreateDeviceAccountRequest) {
    request = &CreateDeviceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateDeviceAccount")
    
    
    return
}

func NewCreateDeviceAccountResponse() (response *CreateDeviceAccountResponse) {
    response = &CreateDeviceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeviceAccount
// 新建主机账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeviceAccount(request *CreateDeviceAccountRequest) (response *CreateDeviceAccountResponse, err error) {
    return c.CreateDeviceAccountWithContext(context.Background(), request)
}

// CreateDeviceAccount
// 新建主机账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeviceAccountWithContext(ctx context.Context, request *CreateDeviceAccountRequest) (response *CreateDeviceAccountResponse, err error) {
    if request == nil {
        request = NewCreateDeviceAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateDeviceAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeviceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeviceGroupRequest() (request *CreateDeviceGroupRequest) {
    request = &CreateDeviceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateDeviceGroup")
    
    
    return
}

func NewCreateDeviceGroupResponse() (response *CreateDeviceGroupResponse) {
    response = &CreateDeviceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeviceGroup
// 新建资产组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeviceGroup(request *CreateDeviceGroupRequest) (response *CreateDeviceGroupResponse, err error) {
    return c.CreateDeviceGroupWithContext(context.Background(), request)
}

// CreateDeviceGroup
// 新建资产组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeviceGroupWithContext(ctx context.Context, request *CreateDeviceGroupRequest) (response *CreateDeviceGroupResponse, err error) {
    if request == nil {
        request = NewCreateDeviceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateDeviceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeviceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOperationTaskRequest() (request *CreateOperationTaskRequest) {
    request = &CreateOperationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateOperationTask")
    
    
    return
}

func NewCreateOperationTaskResponse() (response *CreateOperationTaskResponse) {
    response = &CreateOperationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOperationTask
// 创建运维任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_RESOURCENOTPRO = "FailedOperation.ResourceNotPro"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateOperationTask(request *CreateOperationTaskRequest) (response *CreateOperationTaskResponse, err error) {
    return c.CreateOperationTaskWithContext(context.Background(), request)
}

// CreateOperationTask
// 创建运维任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_RESOURCENOTPRO = "FailedOperation.ResourceNotPro"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateOperationTaskWithContext(ctx context.Context, request *CreateOperationTaskRequest) (response *CreateOperationTaskResponse, err error) {
    if request == nil {
        request = NewCreateOperationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateOperationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOperationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOperationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceRequest() (request *CreateResourceRequest) {
    request = &CreateResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateResource")
    
    
    return
}

func NewCreateResourceResponse() (response *CreateResourceResponse) {
    response = &CreateResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResource
// 创建堡垒机实例
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    return c.CreateResourceWithContext(context.Background(), request)
}

// CreateResource
// 创建堡垒机实例
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateResourceWithContext(ctx context.Context, request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSyncUserTaskRequest() (request *CreateSyncUserTaskRequest) {
    request = &CreateSyncUserTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateSyncUserTask")
    
    
    return
}

func NewCreateSyncUserTaskResponse() (response *CreateSyncUserTaskResponse) {
    response = &CreateSyncUserTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSyncUserTask
// 创建用户同步任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateSyncUserTask(request *CreateSyncUserTaskRequest) (response *CreateSyncUserTaskResponse, err error) {
    return c.CreateSyncUserTaskWithContext(context.Background(), request)
}

// CreateSyncUserTask
// 创建用户同步任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateSyncUserTaskWithContext(ctx context.Context, request *CreateSyncUserTaskRequest) (response *CreateSyncUserTaskResponse, err error) {
    if request == nil {
        request = NewCreateSyncUserTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateSyncUserTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSyncUserTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSyncUserTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// 新建用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 新建用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserDirectoryRequest() (request *CreateUserDirectoryRequest) {
    request = &CreateUserDirectoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateUserDirectory")
    
    
    return
}

func NewCreateUserDirectoryResponse() (response *CreateUserDirectoryResponse) {
    response = &CreateUserDirectoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserDirectory
// 创建用户目录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserDirectory(request *CreateUserDirectoryRequest) (response *CreateUserDirectoryResponse, err error) {
    return c.CreateUserDirectoryWithContext(context.Background(), request)
}

// CreateUserDirectory
// 创建用户目录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserDirectoryWithContext(ctx context.Context, request *CreateUserDirectoryRequest) (response *CreateUserDirectoryResponse, err error) {
    if request == nil {
        request = NewCreateUserDirectoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateUserDirectory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserDirectory require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserDirectoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserGroupRequest() (request *CreateUserGroupRequest) {
    request = &CreateUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "CreateUserGroup")
    
    
    return
}

func NewCreateUserGroupResponse() (response *CreateUserGroupResponse) {
    response = &CreateUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserGroup
// 新建用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserGroup(request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    return c.CreateUserGroupWithContext(context.Background(), request)
}

// CreateUserGroup
// 新建用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserGroupWithContext(ctx context.Context, request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    if request == nil {
        request = NewCreateUserGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "CreateUserGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessWhiteListRulesRequest() (request *DeleteAccessWhiteListRulesRequest) {
    request = &DeleteAccessWhiteListRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteAccessWhiteListRules")
    
    
    return
}

func NewDeleteAccessWhiteListRulesResponse() (response *DeleteAccessWhiteListRulesResponse) {
    response = &DeleteAccessWhiteListRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccessWhiteListRules
// 删除访问白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteAccessWhiteListRules(request *DeleteAccessWhiteListRulesRequest) (response *DeleteAccessWhiteListRulesResponse, err error) {
    return c.DeleteAccessWhiteListRulesWithContext(context.Background(), request)
}

// DeleteAccessWhiteListRules
// 删除访问白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteAccessWhiteListRulesWithContext(ctx context.Context, request *DeleteAccessWhiteListRulesRequest) (response *DeleteAccessWhiteListRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAccessWhiteListRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteAccessWhiteListRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccessWhiteListRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccessWhiteListRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAclsRequest() (request *DeleteAclsRequest) {
    request = &DeleteAclsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteAcls")
    
    
    return
}

func NewDeleteAclsResponse() (response *DeleteAclsResponse) {
    response = &DeleteAclsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAcls
// 删除访问权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAcls(request *DeleteAclsRequest) (response *DeleteAclsResponse, err error) {
    return c.DeleteAclsWithContext(context.Background(), request)
}

// DeleteAcls
// 删除访问权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAclsWithContext(ctx context.Context, request *DeleteAclsRequest) (response *DeleteAclsResponse, err error) {
    if request == nil {
        request = NewDeleteAclsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteAcls")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAcls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAclsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteChangePwdTaskRequest() (request *DeleteChangePwdTaskRequest) {
    request = &DeleteChangePwdTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteChangePwdTask")
    
    
    return
}

func NewDeleteChangePwdTaskResponse() (response *DeleteChangePwdTaskResponse) {
    response = &DeleteChangePwdTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteChangePwdTask
// 删除改密任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
func (c *Client) DeleteChangePwdTask(request *DeleteChangePwdTaskRequest) (response *DeleteChangePwdTaskResponse, err error) {
    return c.DeleteChangePwdTaskWithContext(context.Background(), request)
}

// DeleteChangePwdTask
// 删除改密任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
func (c *Client) DeleteChangePwdTaskWithContext(ctx context.Context, request *DeleteChangePwdTaskRequest) (response *DeleteChangePwdTaskResponse, err error) {
    if request == nil {
        request = NewDeleteChangePwdTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteChangePwdTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteChangePwdTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteChangePwdTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCmdTemplatesRequest() (request *DeleteCmdTemplatesRequest) {
    request = &DeleteCmdTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteCmdTemplates")
    
    
    return
}

func NewDeleteCmdTemplatesResponse() (response *DeleteCmdTemplatesResponse) {
    response = &DeleteCmdTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCmdTemplates
// 删除高危命令模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCmdTemplates(request *DeleteCmdTemplatesRequest) (response *DeleteCmdTemplatesResponse, err error) {
    return c.DeleteCmdTemplatesWithContext(context.Background(), request)
}

// DeleteCmdTemplates
// 删除高危命令模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCmdTemplatesWithContext(ctx context.Context, request *DeleteCmdTemplatesRequest) (response *DeleteCmdTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteCmdTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteCmdTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCmdTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCmdTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceAccountsRequest() (request *DeleteDeviceAccountsRequest) {
    request = &DeleteDeviceAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteDeviceAccounts")
    
    
    return
}

func NewDeleteDeviceAccountsResponse() (response *DeleteDeviceAccountsResponse) {
    response = &DeleteDeviceAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDeviceAccounts
// 删除主机账号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDeviceAccounts(request *DeleteDeviceAccountsRequest) (response *DeleteDeviceAccountsResponse, err error) {
    return c.DeleteDeviceAccountsWithContext(context.Background(), request)
}

// DeleteDeviceAccounts
// 删除主机账号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDeviceAccountsWithContext(ctx context.Context, request *DeleteDeviceAccountsRequest) (response *DeleteDeviceAccountsResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteDeviceAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDeviceAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceGroupMembersRequest() (request *DeleteDeviceGroupMembersRequest) {
    request = &DeleteDeviceGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteDeviceGroupMembers")
    
    
    return
}

func NewDeleteDeviceGroupMembersResponse() (response *DeleteDeviceGroupMembersResponse) {
    response = &DeleteDeviceGroupMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDeviceGroupMembers
// 删除资产组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDeviceGroupMembers(request *DeleteDeviceGroupMembersRequest) (response *DeleteDeviceGroupMembersResponse, err error) {
    return c.DeleteDeviceGroupMembersWithContext(context.Background(), request)
}

// DeleteDeviceGroupMembers
// 删除资产组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDeviceGroupMembersWithContext(ctx context.Context, request *DeleteDeviceGroupMembersRequest) (response *DeleteDeviceGroupMembersResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceGroupMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteDeviceGroupMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDeviceGroupMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceGroupMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceGroupsRequest() (request *DeleteDeviceGroupsRequest) {
    request = &DeleteDeviceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteDeviceGroups")
    
    
    return
}

func NewDeleteDeviceGroupsResponse() (response *DeleteDeviceGroupsResponse) {
    response = &DeleteDeviceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDeviceGroups
// 删除资产组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDeviceGroups(request *DeleteDeviceGroupsRequest) (response *DeleteDeviceGroupsResponse, err error) {
    return c.DeleteDeviceGroupsWithContext(context.Background(), request)
}

// DeleteDeviceGroups
// 删除资产组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDeviceGroupsWithContext(ctx context.Context, request *DeleteDeviceGroupsRequest) (response *DeleteDeviceGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteDeviceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDeviceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDevicesRequest() (request *DeleteDevicesRequest) {
    request = &DeleteDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteDevices")
    
    
    return
}

func NewDeleteDevicesResponse() (response *DeleteDevicesResponse) {
    response = &DeleteDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDevices
// 删除主机
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDevices(request *DeleteDevicesRequest) (response *DeleteDevicesResponse, err error) {
    return c.DeleteDevicesWithContext(context.Background(), request)
}

// DeleteDevices
// 删除主机
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDevicesWithContext(ctx context.Context, request *DeleteDevicesRequest) (response *DeleteDevicesResponse, err error) {
    if request == nil {
        request = NewDeleteDevicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteDevices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOperationTasksRequest() (request *DeleteOperationTasksRequest) {
    request = &DeleteOperationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteOperationTasks")
    
    
    return
}

func NewDeleteOperationTasksResponse() (response *DeleteOperationTasksResponse) {
    response = &DeleteOperationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOperationTasks
// 删除运维任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteOperationTasks(request *DeleteOperationTasksRequest) (response *DeleteOperationTasksResponse, err error) {
    return c.DeleteOperationTasksWithContext(context.Background(), request)
}

// DeleteOperationTasks
// 删除运维任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteOperationTasksWithContext(ctx context.Context, request *DeleteOperationTasksRequest) (response *DeleteOperationTasksResponse, err error) {
    if request == nil {
        request = NewDeleteOperationTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteOperationTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOperationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOperationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserDirectoryRequest() (request *DeleteUserDirectoryRequest) {
    request = &DeleteUserDirectoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteUserDirectory")
    
    
    return
}

func NewDeleteUserDirectoryResponse() (response *DeleteUserDirectoryResponse) {
    response = &DeleteUserDirectoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserDirectory
// 删除用户目录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUserDirectory(request *DeleteUserDirectoryRequest) (response *DeleteUserDirectoryResponse, err error) {
    return c.DeleteUserDirectoryWithContext(context.Background(), request)
}

// DeleteUserDirectory
// 删除用户目录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUserDirectoryWithContext(ctx context.Context, request *DeleteUserDirectoryRequest) (response *DeleteUserDirectoryResponse, err error) {
    if request == nil {
        request = NewDeleteUserDirectoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteUserDirectory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserDirectory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserDirectoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserGroupMembersRequest() (request *DeleteUserGroupMembersRequest) {
    request = &DeleteUserGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteUserGroupMembers")
    
    
    return
}

func NewDeleteUserGroupMembersResponse() (response *DeleteUserGroupMembersResponse) {
    response = &DeleteUserGroupMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserGroupMembers
// 删除用户组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUserGroupMembers(request *DeleteUserGroupMembersRequest) (response *DeleteUserGroupMembersResponse, err error) {
    return c.DeleteUserGroupMembersWithContext(context.Background(), request)
}

// DeleteUserGroupMembers
// 删除用户组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUserGroupMembersWithContext(ctx context.Context, request *DeleteUserGroupMembersRequest) (response *DeleteUserGroupMembersResponse, err error) {
    if request == nil {
        request = NewDeleteUserGroupMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteUserGroupMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserGroupMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserGroupMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserGroupsRequest() (request *DeleteUserGroupsRequest) {
    request = &DeleteUserGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteUserGroups")
    
    
    return
}

func NewDeleteUserGroupsResponse() (response *DeleteUserGroupsResponse) {
    response = &DeleteUserGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserGroups
// 删除用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUserGroups(request *DeleteUserGroupsRequest) (response *DeleteUserGroupsResponse, err error) {
    return c.DeleteUserGroupsWithContext(context.Background(), request)
}

// DeleteUserGroups
// 删除用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUserGroupsWithContext(ctx context.Context, request *DeleteUserGroupsRequest) (response *DeleteUserGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteUserGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteUserGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
    request = &DeleteUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeleteUsers")
    
    
    return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
    response = &DeleteUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUsers
// 删除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    return c.DeleteUsersWithContext(context.Background(), request)
}

// DeleteUsers
// 删除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteUsersWithContext(ctx context.Context, request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    if request == nil {
        request = NewDeleteUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeleteUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDeployResourceRequest() (request *DeployResourceRequest) {
    request = &DeployResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DeployResource")
    
    
    return
}

func NewDeployResourceResponse() (response *DeployResourceResponse) {
    response = &DeployResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployResource
// 开通服务，初始化资源，只针对新购资源
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCDEPLOYED = "FailedOperation.VPCDeployed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeployResource(request *DeployResourceRequest) (response *DeployResourceResponse, err error) {
    return c.DeployResourceWithContext(context.Background(), request)
}

// DeployResource
// 开通服务，初始化资源，只针对新购资源
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCDEPLOYED = "FailedOperation.VPCDeployed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeployResourceWithContext(ctx context.Context, request *DeployResourceRequest) (response *DeployResourceResponse, err error) {
    if request == nil {
        request = NewDeployResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DeployResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessWhiteListRulesRequest() (request *DescribeAccessWhiteListRulesRequest) {
    request = &DescribeAccessWhiteListRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeAccessWhiteListRules")
    
    
    return
}

func NewDescribeAccessWhiteListRulesResponse() (response *DescribeAccessWhiteListRulesResponse) {
    response = &DescribeAccessWhiteListRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessWhiteListRules
// 查询访问白名单规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessWhiteListRules(request *DescribeAccessWhiteListRulesRequest) (response *DescribeAccessWhiteListRulesResponse, err error) {
    return c.DescribeAccessWhiteListRulesWithContext(context.Background(), request)
}

// DescribeAccessWhiteListRules
// 查询访问白名单规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessWhiteListRulesWithContext(ctx context.Context, request *DescribeAccessWhiteListRulesRequest) (response *DescribeAccessWhiteListRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAccessWhiteListRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeAccessWhiteListRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessWhiteListRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessWhiteListRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountGroupsRequest() (request *DescribeAccountGroupsRequest) {
    request = &DescribeAccountGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeAccountGroups")
    
    
    return
}

func NewDescribeAccountGroupsResponse() (response *DescribeAccountGroupsResponse) {
    response = &DescribeAccountGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountGroups
// 获取账号组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccountGroups(request *DescribeAccountGroupsRequest) (response *DescribeAccountGroupsResponse, err error) {
    return c.DescribeAccountGroupsWithContext(context.Background(), request)
}

// DescribeAccountGroups
// 获取账号组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccountGroupsWithContext(ctx context.Context, request *DescribeAccountGroupsRequest) (response *DescribeAccountGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeAccountGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAclsRequest() (request *DescribeAclsRequest) {
    request = &DescribeAclsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeAcls")
    
    
    return
}

func NewDescribeAclsResponse() (response *DescribeAclsResponse) {
    response = &DescribeAclsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAcls
// 查询访问权限列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAcls(request *DescribeAclsRequest) (response *DescribeAclsResponse, err error) {
    return c.DescribeAclsWithContext(context.Background(), request)
}

// DescribeAcls
// 查询访问权限列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAclsWithContext(ctx context.Context, request *DescribeAclsRequest) (response *DescribeAclsResponse, err error) {
    if request == nil {
        request = NewDescribeAclsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeAcls")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAcls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAclsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetSyncFlagRequest() (request *DescribeAssetSyncFlagRequest) {
    request = &DescribeAssetSyncFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeAssetSyncFlag")
    
    
    return
}

func NewDescribeAssetSyncFlagResponse() (response *DescribeAssetSyncFlagResponse) {
    response = &DescribeAssetSyncFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetSyncFlag
// 查询资产自动同步开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetSyncFlag(request *DescribeAssetSyncFlagRequest) (response *DescribeAssetSyncFlagResponse, err error) {
    return c.DescribeAssetSyncFlagWithContext(context.Background(), request)
}

// DescribeAssetSyncFlag
// 查询资产自动同步开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetSyncFlagWithContext(ctx context.Context, request *DescribeAssetSyncFlagRequest) (response *DescribeAssetSyncFlagResponse, err error) {
    if request == nil {
        request = NewDescribeAssetSyncFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeAssetSyncFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetSyncFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetSyncFlagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetSyncStatusRequest() (request *DescribeAssetSyncStatusRequest) {
    request = &DescribeAssetSyncStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeAssetSyncStatus")
    
    
    return
}

func NewDescribeAssetSyncStatusResponse() (response *DescribeAssetSyncStatusResponse) {
    response = &DescribeAssetSyncStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetSyncStatus
// 查询资产同步状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAssetSyncStatus(request *DescribeAssetSyncStatusRequest) (response *DescribeAssetSyncStatusResponse, err error) {
    return c.DescribeAssetSyncStatusWithContext(context.Background(), request)
}

// DescribeAssetSyncStatus
// 查询资产同步状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAssetSyncStatusWithContext(ctx context.Context, request *DescribeAssetSyncStatusRequest) (response *DescribeAssetSyncStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAssetSyncStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeAssetSyncStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetSyncStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetSyncStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChangePwdTaskRequest() (request *DescribeChangePwdTaskRequest) {
    request = &DescribeChangePwdTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeChangePwdTask")
    
    
    return
}

func NewDescribeChangePwdTaskResponse() (response *DescribeChangePwdTaskResponse) {
    response = &DescribeChangePwdTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChangePwdTask
// 查询改密任务列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChangePwdTask(request *DescribeChangePwdTaskRequest) (response *DescribeChangePwdTaskResponse, err error) {
    return c.DescribeChangePwdTaskWithContext(context.Background(), request)
}

// DescribeChangePwdTask
// 查询改密任务列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChangePwdTaskWithContext(ctx context.Context, request *DescribeChangePwdTaskRequest) (response *DescribeChangePwdTaskResponse, err error) {
    if request == nil {
        request = NewDescribeChangePwdTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeChangePwdTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChangePwdTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChangePwdTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChangePwdTaskDetailRequest() (request *DescribeChangePwdTaskDetailRequest) {
    request = &DescribeChangePwdTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeChangePwdTaskDetail")
    
    
    return
}

func NewDescribeChangePwdTaskDetailResponse() (response *DescribeChangePwdTaskDetailResponse) {
    response = &DescribeChangePwdTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChangePwdTaskDetail
// 查询改密任务详情
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChangePwdTaskDetail(request *DescribeChangePwdTaskDetailRequest) (response *DescribeChangePwdTaskDetailResponse, err error) {
    return c.DescribeChangePwdTaskDetailWithContext(context.Background(), request)
}

// DescribeChangePwdTaskDetail
// 查询改密任务详情
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChangePwdTaskDetailWithContext(ctx context.Context, request *DescribeChangePwdTaskDetailRequest) (response *DescribeChangePwdTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeChangePwdTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeChangePwdTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChangePwdTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChangePwdTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmdTemplatesRequest() (request *DescribeCmdTemplatesRequest) {
    request = &DescribeCmdTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeCmdTemplates")
    
    
    return
}

func NewDescribeCmdTemplatesResponse() (response *DescribeCmdTemplatesResponse) {
    response = &DescribeCmdTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCmdTemplates
// 查询命令模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCmdTemplates(request *DescribeCmdTemplatesRequest) (response *DescribeCmdTemplatesResponse, err error) {
    return c.DescribeCmdTemplatesWithContext(context.Background(), request)
}

// DescribeCmdTemplates
// 查询命令模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCmdTemplatesWithContext(ctx context.Context, request *DescribeCmdTemplatesRequest) (response *DescribeCmdTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeCmdTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeCmdTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmdTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmdTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDepartmentsRequest() (request *DescribeDepartmentsRequest) {
    request = &DescribeDepartmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeDepartments")
    
    
    return
}

func NewDescribeDepartmentsResponse() (response *DescribeDepartmentsResponse) {
    response = &DescribeDepartmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDepartments
// 查询部门信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDepartments(request *DescribeDepartmentsRequest) (response *DescribeDepartmentsResponse, err error) {
    return c.DescribeDepartmentsWithContext(context.Background(), request)
}

// DescribeDepartments
// 查询部门信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDepartmentsWithContext(ctx context.Context, request *DescribeDepartmentsRequest) (response *DescribeDepartmentsResponse, err error) {
    if request == nil {
        request = NewDescribeDepartmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeDepartments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDepartments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDepartmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceAccountsRequest() (request *DescribeDeviceAccountsRequest) {
    request = &DescribeDeviceAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeDeviceAccounts")
    
    
    return
}

func NewDescribeDeviceAccountsResponse() (response *DescribeDeviceAccountsResponse) {
    response = &DescribeDeviceAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceAccounts
// 查询主机账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeviceAccounts(request *DescribeDeviceAccountsRequest) (response *DescribeDeviceAccountsResponse, err error) {
    return c.DescribeDeviceAccountsWithContext(context.Background(), request)
}

// DescribeDeviceAccounts
// 查询主机账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeviceAccountsWithContext(ctx context.Context, request *DescribeDeviceAccountsRequest) (response *DescribeDeviceAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeDeviceAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceGroupMembersRequest() (request *DescribeDeviceGroupMembersRequest) {
    request = &DescribeDeviceGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeDeviceGroupMembers")
    
    
    return
}

func NewDescribeDeviceGroupMembersResponse() (response *DescribeDeviceGroupMembersResponse) {
    response = &DescribeDeviceGroupMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceGroupMembers
// 查询资产组成员列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceGroupMembers(request *DescribeDeviceGroupMembersRequest) (response *DescribeDeviceGroupMembersResponse, err error) {
    return c.DescribeDeviceGroupMembersWithContext(context.Background(), request)
}

// DescribeDeviceGroupMembers
// 查询资产组成员列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceGroupMembersWithContext(ctx context.Context, request *DescribeDeviceGroupMembersRequest) (response *DescribeDeviceGroupMembersResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceGroupMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeDeviceGroupMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceGroupMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceGroupMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceGroupsRequest() (request *DescribeDeviceGroupsRequest) {
    request = &DescribeDeviceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeDeviceGroups")
    
    
    return
}

func NewDescribeDeviceGroupsResponse() (response *DescribeDeviceGroupsResponse) {
    response = &DescribeDeviceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceGroups
// 查询资产组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceGroups(request *DescribeDeviceGroupsRequest) (response *DescribeDeviceGroupsResponse, err error) {
    return c.DescribeDeviceGroupsWithContext(context.Background(), request)
}

// DescribeDeviceGroups
// 查询资产组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceGroupsWithContext(ctx context.Context, request *DescribeDeviceGroupsRequest) (response *DescribeDeviceGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeDeviceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
    request = &DescribeDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeDevices")
    
    
    return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
    response = &DescribeDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDevices
// 查询资产列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    return c.DescribeDevicesWithContext(context.Background(), request)
}

// DescribeDevices
// 查询资产列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDevicesWithContext(ctx context.Context, request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeDevices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeDomains")
    
    
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomains
// 查询网络域
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    return c.DescribeDomainsWithContext(context.Background(), request)
}

// DescribeDomains
// 查询网络域
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDomainsWithContext(ctx context.Context, request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLDAPUnitSetRequest() (request *DescribeLDAPUnitSetRequest) {
    request = &DescribeLDAPUnitSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeLDAPUnitSet")
    
    
    return
}

func NewDescribeLDAPUnitSetResponse() (response *DescribeLDAPUnitSetResponse) {
    response = &DescribeLDAPUnitSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLDAPUnitSet
// 获取LDAP ou 列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAILED = "FailedOperation.AuthFailed"
//  FAILEDOPERATION_CONNECTIONFAILED = "FailedOperation.ConnectionFailed"
//  FAILEDOPERATION_SEARCHFAILED = "FailedOperation.SearchFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLDAPUnitSet(request *DescribeLDAPUnitSetRequest) (response *DescribeLDAPUnitSetResponse, err error) {
    return c.DescribeLDAPUnitSetWithContext(context.Background(), request)
}

// DescribeLDAPUnitSet
// 获取LDAP ou 列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAILED = "FailedOperation.AuthFailed"
//  FAILEDOPERATION_CONNECTIONFAILED = "FailedOperation.ConnectionFailed"
//  FAILEDOPERATION_SEARCHFAILED = "FailedOperation.SearchFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLDAPUnitSetWithContext(ctx context.Context, request *DescribeLDAPUnitSetRequest) (response *DescribeLDAPUnitSetResponse, err error) {
    if request == nil {
        request = NewDescribeLDAPUnitSetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeLDAPUnitSet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLDAPUnitSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLDAPUnitSetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginEventRequest() (request *DescribeLoginEventRequest) {
    request = &DescribeLoginEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeLoginEvent")
    
    
    return
}

func NewDescribeLoginEventResponse() (response *DescribeLoginEventResponse) {
    response = &DescribeLoginEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoginEvent
// 查询登录日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLoginEvent(request *DescribeLoginEventRequest) (response *DescribeLoginEventResponse, err error) {
    return c.DescribeLoginEventWithContext(context.Background(), request)
}

// DescribeLoginEvent
// 查询登录日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLoginEventWithContext(ctx context.Context, request *DescribeLoginEventRequest) (response *DescribeLoginEventResponse, err error) {
    if request == nil {
        request = NewDescribeLoginEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeLoginEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoginEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoginEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOperationEventRequest() (request *DescribeOperationEventRequest) {
    request = &DescribeOperationEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeOperationEvent")
    
    
    return
}

func NewDescribeOperationEventResponse() (response *DescribeOperationEventResponse) {
    response = &DescribeOperationEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOperationEvent
// 查询操作日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeOperationEvent(request *DescribeOperationEventRequest) (response *DescribeOperationEventResponse, err error) {
    return c.DescribeOperationEventWithContext(context.Background(), request)
}

// DescribeOperationEvent
// 查询操作日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeOperationEventWithContext(ctx context.Context, request *DescribeOperationEventRequest) (response *DescribeOperationEventResponse, err error) {
    if request == nil {
        request = NewDescribeOperationEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeOperationEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOperationEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOperationEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOperationTaskRequest() (request *DescribeOperationTaskRequest) {
    request = &DescribeOperationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeOperationTask")
    
    
    return
}

func NewDescribeOperationTaskResponse() (response *DescribeOperationTaskResponse) {
    response = &DescribeOperationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOperationTask
// 获取运维任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeOperationTask(request *DescribeOperationTaskRequest) (response *DescribeOperationTaskResponse, err error) {
    return c.DescribeOperationTaskWithContext(context.Background(), request)
}

// DescribeOperationTask
// 获取运维任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeOperationTaskWithContext(ctx context.Context, request *DescribeOperationTaskRequest) (response *DescribeOperationTaskResponse, err error) {
    if request == nil {
        request = NewDescribeOperationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeOperationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOperationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOperationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
    request = &DescribeResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeResources")
    
    
    return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
    response = &DescribeResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResources
// 查询用户购买的堡垒机服务信息，包括资源ID、授权点数、VPC、过期时间等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
    return c.DescribeResourcesWithContext(context.Background(), request)
}

// DescribeResources
// 查询用户购买的堡垒机服务信息，包括资源ID、授权点数、VPC、过期时间等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeResourcesWithContext(ctx context.Context, request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecuritySettingRequest() (request *DescribeSecuritySettingRequest) {
    request = &DescribeSecuritySettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeSecuritySetting")
    
    
    return
}

func NewDescribeSecuritySettingResponse() (response *DescribeSecuritySettingResponse) {
    response = &DescribeSecuritySettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecuritySetting
// 查询安全配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecuritySetting(request *DescribeSecuritySettingRequest) (response *DescribeSecuritySettingResponse, err error) {
    return c.DescribeSecuritySettingWithContext(context.Background(), request)
}

// DescribeSecuritySetting
// 查询安全配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecuritySettingWithContext(ctx context.Context, request *DescribeSecuritySettingRequest) (response *DescribeSecuritySettingResponse, err error) {
    if request == nil {
        request = NewDescribeSecuritySettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeSecuritySetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecuritySetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecuritySettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceTypesRequest() (request *DescribeSourceTypesRequest) {
    request = &DescribeSourceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeSourceTypes")
    
    
    return
}

func NewDescribeSourceTypesResponse() (response *DescribeSourceTypesResponse) {
    response = &DescribeSourceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSourceTypes
// 获取认证源信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSourceTypes(request *DescribeSourceTypesRequest) (response *DescribeSourceTypesResponse, err error) {
    return c.DescribeSourceTypesWithContext(context.Background(), request)
}

// DescribeSourceTypes
// 获取认证源信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSourceTypesWithContext(ctx context.Context, request *DescribeSourceTypesRequest) (response *DescribeSourceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeSourceTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeSourceTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSourceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSourceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDirectoryRequest() (request *DescribeUserDirectoryRequest) {
    request = &DescribeUserDirectoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeUserDirectory")
    
    
    return
}

func NewDescribeUserDirectoryResponse() (response *DescribeUserDirectoryResponse) {
    response = &DescribeUserDirectoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserDirectory
// 获取用户目录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserDirectory(request *DescribeUserDirectoryRequest) (response *DescribeUserDirectoryResponse, err error) {
    return c.DescribeUserDirectoryWithContext(context.Background(), request)
}

// DescribeUserDirectory
// 获取用户目录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserDirectoryWithContext(ctx context.Context, request *DescribeUserDirectoryRequest) (response *DescribeUserDirectoryResponse, err error) {
    if request == nil {
        request = NewDescribeUserDirectoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeUserDirectory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDirectory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDirectoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserGroupMembersRequest() (request *DescribeUserGroupMembersRequest) {
    request = &DescribeUserGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeUserGroupMembers")
    
    
    return
}

func NewDescribeUserGroupMembersResponse() (response *DescribeUserGroupMembersResponse) {
    response = &DescribeUserGroupMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserGroupMembers
// 查询用户组成员列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserGroupMembers(request *DescribeUserGroupMembersRequest) (response *DescribeUserGroupMembersResponse, err error) {
    return c.DescribeUserGroupMembersWithContext(context.Background(), request)
}

// DescribeUserGroupMembers
// 查询用户组成员列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserGroupMembersWithContext(ctx context.Context, request *DescribeUserGroupMembersRequest) (response *DescribeUserGroupMembersResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeUserGroupMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserGroupMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserGroupMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserGroupsRequest() (request *DescribeUserGroupsRequest) {
    request = &DescribeUserGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeUserGroups")
    
    
    return
}

func NewDescribeUserGroupsResponse() (response *DescribeUserGroupsResponse) {
    response = &DescribeUserGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserGroups
// 查询用户组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserGroups(request *DescribeUserGroupsRequest) (response *DescribeUserGroupsResponse, err error) {
    return c.DescribeUserGroupsWithContext(context.Background(), request)
}

// DescribeUserGroups
// 查询用户组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserGroupsWithContext(ctx context.Context, request *DescribeUserGroupsRequest) (response *DescribeUserGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeUserGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserSyncStatusRequest() (request *DescribeUserSyncStatusRequest) {
    request = &DescribeUserSyncStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeUserSyncStatus")
    
    
    return
}

func NewDescribeUserSyncStatusResponse() (response *DescribeUserSyncStatusResponse) {
    response = &DescribeUserSyncStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserSyncStatus
// 获取用户同步状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserSyncStatus(request *DescribeUserSyncStatusRequest) (response *DescribeUserSyncStatusResponse, err error) {
    return c.DescribeUserSyncStatusWithContext(context.Background(), request)
}

// DescribeUserSyncStatus
// 获取用户同步状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserSyncStatusWithContext(ctx context.Context, request *DescribeUserSyncStatusRequest) (response *DescribeUserSyncStatusResponse, err error) {
    if request == nil {
        request = NewDescribeUserSyncStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeUserSyncStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserSyncStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserSyncStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersRequest() (request *DescribeUsersRequest) {
    request = &DescribeUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DescribeUsers")
    
    
    return
}

func NewDescribeUsersResponse() (response *DescribeUsersResponse) {
    response = &DescribeUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsers
// 查询用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUsers(request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    return c.DescribeUsersWithContext(context.Background(), request)
}

// DescribeUsers
// 查询用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUsersWithContext(ctx context.Context, request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    if request == nil {
        request = NewDescribeUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DescribeUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDisableExternalAccessRequest() (request *DisableExternalAccessRequest) {
    request = &DisableExternalAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DisableExternalAccess")
    
    
    return
}

func NewDisableExternalAccessResponse() (response *DisableExternalAccessResponse) {
    response = &DisableExternalAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableExternalAccess
// 关闭公网访问堡垒机
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DisableExternalAccess(request *DisableExternalAccessRequest) (response *DisableExternalAccessResponse, err error) {
    return c.DisableExternalAccessWithContext(context.Background(), request)
}

// DisableExternalAccess
// 关闭公网访问堡垒机
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DisableExternalAccessWithContext(ctx context.Context, request *DisableExternalAccessRequest) (response *DisableExternalAccessResponse, err error) {
    if request == nil {
        request = NewDisableExternalAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DisableExternalAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableExternalAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableExternalAccessResponse()
    err = c.Send(request, response)
    return
}

func NewDisableIntranetAccessRequest() (request *DisableIntranetAccessRequest) {
    request = &DisableIntranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "DisableIntranetAccess")
    
    
    return
}

func NewDisableIntranetAccessResponse() (response *DisableIntranetAccessResponse) {
    response = &DisableIntranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableIntranetAccess
// 关闭内网访问
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DisableIntranetAccess(request *DisableIntranetAccessRequest) (response *DisableIntranetAccessResponse, err error) {
    return c.DisableIntranetAccessWithContext(context.Background(), request)
}

// DisableIntranetAccess
// 关闭内网访问
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DisableIntranetAccessWithContext(ctx context.Context, request *DisableIntranetAccessRequest) (response *DisableIntranetAccessResponse, err error) {
    if request == nil {
        request = NewDisableIntranetAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "DisableIntranetAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableIntranetAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableIntranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewEnableExternalAccessRequest() (request *EnableExternalAccessRequest) {
    request = &EnableExternalAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "EnableExternalAccess")
    
    
    return
}

func NewEnableExternalAccessResponse() (response *EnableExternalAccessResponse) {
    response = &EnableExternalAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableExternalAccess
// 开启公网访问堡垒机
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableExternalAccess(request *EnableExternalAccessRequest) (response *EnableExternalAccessResponse, err error) {
    return c.EnableExternalAccessWithContext(context.Background(), request)
}

// EnableExternalAccess
// 开启公网访问堡垒机
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableExternalAccessWithContext(ctx context.Context, request *EnableExternalAccessRequest) (response *EnableExternalAccessResponse, err error) {
    if request == nil {
        request = NewEnableExternalAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "EnableExternalAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableExternalAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableExternalAccessResponse()
    err = c.Send(request, response)
    return
}

func NewEnableIntranetAccessRequest() (request *EnableIntranetAccessRequest) {
    request = &EnableIntranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "EnableIntranetAccess")
    
    
    return
}

func NewEnableIntranetAccessResponse() (response *EnableIntranetAccessResponse) {
    response = &EnableIntranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableIntranetAccess
// 开通内网访问
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
func (c *Client) EnableIntranetAccess(request *EnableIntranetAccessRequest) (response *EnableIntranetAccessResponse, err error) {
    return c.EnableIntranetAccessWithContext(context.Background(), request)
}

// EnableIntranetAccess
// 开通内网访问
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
func (c *Client) EnableIntranetAccessWithContext(ctx context.Context, request *EnableIntranetAccessRequest) (response *EnableIntranetAccessResponse, err error) {
    if request == nil {
        request = NewEnableIntranetAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "EnableIntranetAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableIntranetAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableIntranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewImportExternalDeviceRequest() (request *ImportExternalDeviceRequest) {
    request = &ImportExternalDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ImportExternalDevice")
    
    
    return
}

func NewImportExternalDeviceResponse() (response *ImportExternalDeviceResponse) {
    response = &ImportExternalDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportExternalDevice
// 导入外部资产信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImportExternalDevice(request *ImportExternalDeviceRequest) (response *ImportExternalDeviceResponse, err error) {
    return c.ImportExternalDeviceWithContext(context.Background(), request)
}

// ImportExternalDevice
// 导入外部资产信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImportExternalDeviceWithContext(ctx context.Context, request *ImportExternalDeviceRequest) (response *ImportExternalDeviceResponse, err error) {
    if request == nil {
        request = NewImportExternalDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ImportExternalDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportExternalDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportExternalDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessWhiteListAutoStatusRequest() (request *ModifyAccessWhiteListAutoStatusRequest) {
    request = &ModifyAccessWhiteListAutoStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyAccessWhiteListAutoStatus")
    
    
    return
}

func NewModifyAccessWhiteListAutoStatusResponse() (response *ModifyAccessWhiteListAutoStatusResponse) {
    response = &ModifyAccessWhiteListAutoStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccessWhiteListAutoStatus
// 修改访问白名单自动添加IP状态：开启或关闭自动添加IP
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAccessWhiteListAutoStatus(request *ModifyAccessWhiteListAutoStatusRequest) (response *ModifyAccessWhiteListAutoStatusResponse, err error) {
    return c.ModifyAccessWhiteListAutoStatusWithContext(context.Background(), request)
}

// ModifyAccessWhiteListAutoStatus
// 修改访问白名单自动添加IP状态：开启或关闭自动添加IP
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAccessWhiteListAutoStatusWithContext(ctx context.Context, request *ModifyAccessWhiteListAutoStatusRequest) (response *ModifyAccessWhiteListAutoStatusResponse, err error) {
    if request == nil {
        request = NewModifyAccessWhiteListAutoStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyAccessWhiteListAutoStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessWhiteListAutoStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessWhiteListAutoStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessWhiteListRuleRequest() (request *ModifyAccessWhiteListRuleRequest) {
    request = &ModifyAccessWhiteListRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyAccessWhiteListRule")
    
    
    return
}

func NewModifyAccessWhiteListRuleResponse() (response *ModifyAccessWhiteListRuleResponse) {
    response = &ModifyAccessWhiteListRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccessWhiteListRule
// 修改访问白名单规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAccessWhiteListRule(request *ModifyAccessWhiteListRuleRequest) (response *ModifyAccessWhiteListRuleResponse, err error) {
    return c.ModifyAccessWhiteListRuleWithContext(context.Background(), request)
}

// ModifyAccessWhiteListRule
// 修改访问白名单规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAccessWhiteListRuleWithContext(ctx context.Context, request *ModifyAccessWhiteListRuleRequest) (response *ModifyAccessWhiteListRuleResponse, err error) {
    if request == nil {
        request = NewModifyAccessWhiteListRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyAccessWhiteListRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessWhiteListRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessWhiteListRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessWhiteListStatusRequest() (request *ModifyAccessWhiteListStatusRequest) {
    request = &ModifyAccessWhiteListStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyAccessWhiteListStatus")
    
    
    return
}

func NewModifyAccessWhiteListStatusResponse() (response *ModifyAccessWhiteListStatusResponse) {
    response = &ModifyAccessWhiteListStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccessWhiteListStatus
// 修改访问白名单状态：开启或关闭放开全部来源IP。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAccessWhiteListStatus(request *ModifyAccessWhiteListStatusRequest) (response *ModifyAccessWhiteListStatusResponse, err error) {
    return c.ModifyAccessWhiteListStatusWithContext(context.Background(), request)
}

// ModifyAccessWhiteListStatus
// 修改访问白名单状态：开启或关闭放开全部来源IP。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAccessWhiteListStatusWithContext(ctx context.Context, request *ModifyAccessWhiteListStatusRequest) (response *ModifyAccessWhiteListStatusResponse, err error) {
    if request == nil {
        request = NewModifyAccessWhiteListStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyAccessWhiteListStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessWhiteListStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessWhiteListStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAclRequest() (request *ModifyAclRequest) {
    request = &ModifyAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyAcl")
    
    
    return
}

func NewModifyAclResponse() (response *ModifyAclResponse) {
    response = &ModifyAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAcl
// 修改访问权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAcl(request *ModifyAclRequest) (response *ModifyAclResponse, err error) {
    return c.ModifyAclWithContext(context.Background(), request)
}

// ModifyAcl
// 修改访问权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAclWithContext(ctx context.Context, request *ModifyAclRequest) (response *ModifyAclResponse, err error) {
    if request == nil {
        request = NewModifyAclRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyAcl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAclResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetSyncFlagRequest() (request *ModifyAssetSyncFlagRequest) {
    request = &ModifyAssetSyncFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyAssetSyncFlag")
    
    
    return
}

func NewModifyAssetSyncFlagResponse() (response *ModifyAssetSyncFlagResponse) {
    response = &ModifyAssetSyncFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAssetSyncFlag
// 修改资产自动同步开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAssetSyncFlag(request *ModifyAssetSyncFlagRequest) (response *ModifyAssetSyncFlagResponse, err error) {
    return c.ModifyAssetSyncFlagWithContext(context.Background(), request)
}

// ModifyAssetSyncFlag
// 修改资产自动同步开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAssetSyncFlagWithContext(ctx context.Context, request *ModifyAssetSyncFlagRequest) (response *ModifyAssetSyncFlagResponse, err error) {
    if request == nil {
        request = NewModifyAssetSyncFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyAssetSyncFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetSyncFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetSyncFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuthModeSettingRequest() (request *ModifyAuthModeSettingRequest) {
    request = &ModifyAuthModeSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyAuthModeSetting")
    
    
    return
}

func NewModifyAuthModeSettingResponse() (response *ModifyAuthModeSettingResponse) {
    response = &ModifyAuthModeSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuthModeSetting
// 修改认证方式配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAuthModeSetting(request *ModifyAuthModeSettingRequest) (response *ModifyAuthModeSettingResponse, err error) {
    return c.ModifyAuthModeSettingWithContext(context.Background(), request)
}

// ModifyAuthModeSetting
// 修改认证方式配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAuthModeSettingWithContext(ctx context.Context, request *ModifyAuthModeSettingRequest) (response *ModifyAuthModeSettingResponse, err error) {
    if request == nil {
        request = NewModifyAuthModeSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyAuthModeSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuthModeSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuthModeSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyChangePwdTaskRequest() (request *ModifyChangePwdTaskRequest) {
    request = &ModifyChangePwdTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyChangePwdTask")
    
    
    return
}

func NewModifyChangePwdTaskResponse() (response *ModifyChangePwdTaskResponse) {
    response = &ModifyChangePwdTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyChangePwdTask
// 更新修改密码任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyChangePwdTask(request *ModifyChangePwdTaskRequest) (response *ModifyChangePwdTaskResponse, err error) {
    return c.ModifyChangePwdTaskWithContext(context.Background(), request)
}

// ModifyChangePwdTask
// 更新修改密码任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyChangePwdTaskWithContext(ctx context.Context, request *ModifyChangePwdTaskRequest) (response *ModifyChangePwdTaskResponse, err error) {
    if request == nil {
        request = NewModifyChangePwdTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyChangePwdTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyChangePwdTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyChangePwdTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCmdTemplateRequest() (request *ModifyCmdTemplateRequest) {
    request = &ModifyCmdTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyCmdTemplate")
    
    
    return
}

func NewModifyCmdTemplateResponse() (response *ModifyCmdTemplateResponse) {
    response = &ModifyCmdTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCmdTemplate
// 修改高危命令模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCmdTemplate(request *ModifyCmdTemplateRequest) (response *ModifyCmdTemplateResponse, err error) {
    return c.ModifyCmdTemplateWithContext(context.Background(), request)
}

// ModifyCmdTemplate
// 修改高危命令模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCmdTemplateWithContext(ctx context.Context, request *ModifyCmdTemplateRequest) (response *ModifyCmdTemplateResponse, err error) {
    if request == nil {
        request = NewModifyCmdTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyCmdTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCmdTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCmdTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceRequest() (request *ModifyDeviceRequest) {
    request = &ModifyDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyDevice")
    
    
    return
}

func NewModifyDeviceResponse() (response *ModifyDeviceResponse) {
    response = &ModifyDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDevice
// 修改资产信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDevice(request *ModifyDeviceRequest) (response *ModifyDeviceResponse, err error) {
    return c.ModifyDeviceWithContext(context.Background(), request)
}

// ModifyDevice
// 修改资产信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDeviceWithContext(ctx context.Context, request *ModifyDeviceRequest) (response *ModifyDeviceResponse, err error) {
    if request == nil {
        request = NewModifyDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceGroupRequest() (request *ModifyDeviceGroupRequest) {
    request = &ModifyDeviceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyDeviceGroup")
    
    
    return
}

func NewModifyDeviceGroupResponse() (response *ModifyDeviceGroupResponse) {
    response = &ModifyDeviceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDeviceGroup
// 修改资产组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDeviceGroup(request *ModifyDeviceGroupRequest) (response *ModifyDeviceGroupResponse, err error) {
    return c.ModifyDeviceGroupWithContext(context.Background(), request)
}

// ModifyDeviceGroup
// 修改资产组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDeviceGroupWithContext(ctx context.Context, request *ModifyDeviceGroupRequest) (response *ModifyDeviceGroupResponse, err error) {
    if request == nil {
        request = NewModifyDeviceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyDeviceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLDAPSettingRequest() (request *ModifyLDAPSettingRequest) {
    request = &ModifyLDAPSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyLDAPSetting")
    
    
    return
}

func NewModifyLDAPSettingResponse() (response *ModifyLDAPSettingResponse) {
    response = &ModifyLDAPSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLDAPSetting
// 修改LDAP配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLDAPSetting(request *ModifyLDAPSettingRequest) (response *ModifyLDAPSettingResponse, err error) {
    return c.ModifyLDAPSettingWithContext(context.Background(), request)
}

// ModifyLDAPSetting
// 修改LDAP配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLDAPSettingWithContext(ctx context.Context, request *ModifyLDAPSettingRequest) (response *ModifyLDAPSettingResponse, err error) {
    if request == nil {
        request = NewModifyLDAPSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyLDAPSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLDAPSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLDAPSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOAuthSettingRequest() (request *ModifyOAuthSettingRequest) {
    request = &ModifyOAuthSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyOAuthSetting")
    
    
    return
}

func NewModifyOAuthSettingResponse() (response *ModifyOAuthSettingResponse) {
    response = &ModifyOAuthSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOAuthSetting
// 设置OAuth认证参数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyOAuthSetting(request *ModifyOAuthSettingRequest) (response *ModifyOAuthSettingResponse, err error) {
    return c.ModifyOAuthSettingWithContext(context.Background(), request)
}

// ModifyOAuthSetting
// 设置OAuth认证参数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyOAuthSettingWithContext(ctx context.Context, request *ModifyOAuthSettingRequest) (response *ModifyOAuthSettingResponse, err error) {
    if request == nil {
        request = NewModifyOAuthSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyOAuthSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOAuthSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOAuthSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOperationTaskRequest() (request *ModifyOperationTaskRequest) {
    request = &ModifyOperationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyOperationTask")
    
    
    return
}

func NewModifyOperationTaskResponse() (response *ModifyOperationTaskResponse) {
    response = &ModifyOperationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOperationTask
// 修改运维任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  FAILEDOPERATION_RESOURCENOTPRO = "FailedOperation.ResourceNotPro"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOperationTask(request *ModifyOperationTaskRequest) (response *ModifyOperationTaskResponse, err error) {
    return c.ModifyOperationTaskWithContext(context.Background(), request)
}

// ModifyOperationTask
// 修改运维任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  FAILEDOPERATION_RESOURCENOTPRO = "FailedOperation.ResourceNotPro"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOperationTaskWithContext(ctx context.Context, request *ModifyOperationTaskRequest) (response *ModifyOperationTaskResponse, err error) {
    if request == nil {
        request = NewModifyOperationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyOperationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOperationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOperationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReconnectionSettingRequest() (request *ModifyReconnectionSettingRequest) {
    request = &ModifyReconnectionSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyReconnectionSetting")
    
    
    return
}

func NewModifyReconnectionSettingResponse() (response *ModifyReconnectionSettingResponse) {
    response = &ModifyReconnectionSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyReconnectionSetting
// 修改运维资产连接重连次数
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyReconnectionSetting(request *ModifyReconnectionSettingRequest) (response *ModifyReconnectionSettingResponse, err error) {
    return c.ModifyReconnectionSettingWithContext(context.Background(), request)
}

// ModifyReconnectionSetting
// 修改运维资产连接重连次数
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyReconnectionSettingWithContext(ctx context.Context, request *ModifyReconnectionSettingRequest) (response *ModifyReconnectionSettingResponse, err error) {
    if request == nil {
        request = NewModifyReconnectionSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyReconnectionSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyReconnectionSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyReconnectionSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceRequest() (request *ModifyResourceRequest) {
    request = &ModifyResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyResource")
    
    
    return
}

func NewModifyResourceResponse() (response *ModifyResourceResponse) {
    response = &ModifyResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResource
// 资源变配
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCDEPLOYED = "FailedOperation.VPCDeployed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyResource(request *ModifyResourceRequest) (response *ModifyResourceResponse, err error) {
    return c.ModifyResourceWithContext(context.Background(), request)
}

// ModifyResource
// 资源变配
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCDEPLOYED = "FailedOperation.VPCDeployed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyResourceWithContext(ctx context.Context, request *ModifyResourceRequest) (response *ModifyResourceResponse, err error) {
    if request == nil {
        request = NewModifyResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUser
// 修改用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    return c.ModifyUserWithContext(context.Background(), request)
}

// ModifyUser
// 修改用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserDirectoryRequest() (request *ModifyUserDirectoryRequest) {
    request = &ModifyUserDirectoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyUserDirectory")
    
    
    return
}

func NewModifyUserDirectoryResponse() (response *ModifyUserDirectoryResponse) {
    response = &ModifyUserDirectoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserDirectory
// 修改用户目录信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserDirectory(request *ModifyUserDirectoryRequest) (response *ModifyUserDirectoryResponse, err error) {
    return c.ModifyUserDirectoryWithContext(context.Background(), request)
}

// ModifyUserDirectory
// 修改用户目录信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserDirectoryWithContext(ctx context.Context, request *ModifyUserDirectoryRequest) (response *ModifyUserDirectoryResponse, err error) {
    if request == nil {
        request = NewModifyUserDirectoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyUserDirectory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserDirectory require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserDirectoryResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserGroupRequest() (request *ModifyUserGroupRequest) {
    request = &ModifyUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ModifyUserGroup")
    
    
    return
}

func NewModifyUserGroupResponse() (response *ModifyUserGroupResponse) {
    response = &ModifyUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserGroup
// 修改用户组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserGroup(request *ModifyUserGroupRequest) (response *ModifyUserGroupResponse, err error) {
    return c.ModifyUserGroupWithContext(context.Background(), request)
}

// ModifyUserGroup
// 修改用户组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserGroupWithContext(ctx context.Context, request *ModifyUserGroupRequest) (response *ModifyUserGroupResponse, err error) {
    if request == nil {
        request = NewModifyUserGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ModifyUserGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewReplaySessionRequest() (request *ReplaySessionRequest) {
    request = &ReplaySessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ReplaySession")
    
    
    return
}

func NewReplaySessionResponse() (response *ReplaySessionResponse) {
    response = &ReplaySessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaySession
// 会话回放
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ReplaySession(request *ReplaySessionRequest) (response *ReplaySessionResponse, err error) {
    return c.ReplaySessionWithContext(context.Background(), request)
}

// ReplaySession
// 会话回放
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ReplaySessionWithContext(ctx context.Context, request *ReplaySessionRequest) (response *ReplaySessionResponse, err error) {
    if request == nil {
        request = NewReplaySessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ReplaySession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaySession require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaySessionResponse()
    err = c.Send(request, response)
    return
}

func NewResetDeviceAccountPasswordRequest() (request *ResetDeviceAccountPasswordRequest) {
    request = &ResetDeviceAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ResetDeviceAccountPassword")
    
    
    return
}

func NewResetDeviceAccountPasswordResponse() (response *ResetDeviceAccountPasswordResponse) {
    response = &ResetDeviceAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetDeviceAccountPassword
// 清除设备账号绑定密码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ResetDeviceAccountPassword(request *ResetDeviceAccountPasswordRequest) (response *ResetDeviceAccountPasswordResponse, err error) {
    return c.ResetDeviceAccountPasswordWithContext(context.Background(), request)
}

// ResetDeviceAccountPassword
// 清除设备账号绑定密码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ResetDeviceAccountPasswordWithContext(ctx context.Context, request *ResetDeviceAccountPasswordRequest) (response *ResetDeviceAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetDeviceAccountPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ResetDeviceAccountPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDeviceAccountPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetDeviceAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewResetDeviceAccountPrivateKeyRequest() (request *ResetDeviceAccountPrivateKeyRequest) {
    request = &ResetDeviceAccountPrivateKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ResetDeviceAccountPrivateKey")
    
    
    return
}

func NewResetDeviceAccountPrivateKeyResponse() (response *ResetDeviceAccountPrivateKeyResponse) {
    response = &ResetDeviceAccountPrivateKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetDeviceAccountPrivateKey
// 清除设备账号绑定的密钥
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ResetDeviceAccountPrivateKey(request *ResetDeviceAccountPrivateKeyRequest) (response *ResetDeviceAccountPrivateKeyResponse, err error) {
    return c.ResetDeviceAccountPrivateKeyWithContext(context.Background(), request)
}

// ResetDeviceAccountPrivateKey
// 清除设备账号绑定的密钥
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ResetDeviceAccountPrivateKeyWithContext(ctx context.Context, request *ResetDeviceAccountPrivateKeyRequest) (response *ResetDeviceAccountPrivateKeyResponse, err error) {
    if request == nil {
        request = NewResetDeviceAccountPrivateKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ResetDeviceAccountPrivateKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDeviceAccountPrivateKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetDeviceAccountPrivateKeyResponse()
    err = c.Send(request, response)
    return
}

func NewResetUserRequest() (request *ResetUserRequest) {
    request = &ResetUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "ResetUser")
    
    
    return
}

func NewResetUserResponse() (response *ResetUserResponse) {
    response = &ResetUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetUser
// 重置用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetUser(request *ResetUserRequest) (response *ResetUserResponse, err error) {
    return c.ResetUserWithContext(context.Background(), request)
}

// ResetUser
// 重置用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetUserWithContext(ctx context.Context, request *ResetUserRequest) (response *ResetUserResponse, err error) {
    if request == nil {
        request = NewResetUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "ResetUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetUserResponse()
    err = c.Send(request, response)
    return
}

func NewRunChangePwdTaskRequest() (request *RunChangePwdTaskRequest) {
    request = &RunChangePwdTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "RunChangePwdTask")
    
    
    return
}

func NewRunChangePwdTaskResponse() (response *RunChangePwdTaskResponse) {
    response = &RunChangePwdTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunChangePwdTask
// 执行改密任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TOOFREQUENT = "FailedOperation.TooFrequent"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RunChangePwdTask(request *RunChangePwdTaskRequest) (response *RunChangePwdTaskResponse, err error) {
    return c.RunChangePwdTaskWithContext(context.Background(), request)
}

// RunChangePwdTask
// 执行改密任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TOOFREQUENT = "FailedOperation.TooFrequent"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RunChangePwdTaskWithContext(ctx context.Context, request *RunChangePwdTaskRequest) (response *RunChangePwdTaskResponse, err error) {
    if request == nil {
        request = NewRunChangePwdTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "RunChangePwdTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunChangePwdTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunChangePwdTaskResponse()
    err = c.Send(request, response)
    return
}

func NewRunOperationTaskRequest() (request *RunOperationTaskRequest) {
    request = &RunOperationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "RunOperationTask")
    
    
    return
}

func NewRunOperationTaskResponse() (response *RunOperationTaskResponse) {
    response = &RunOperationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunOperationTask
// 执行运维任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RunOperationTask(request *RunOperationTaskRequest) (response *RunOperationTaskResponse, err error) {
    return c.RunOperationTaskWithContext(context.Background(), request)
}

// RunOperationTask
// 执行运维任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"
//  FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"
//  FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"
//  FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RunOperationTaskWithContext(ctx context.Context, request *RunOperationTaskRequest) (response *RunOperationTaskResponse, err error) {
    if request == nil {
        request = NewRunOperationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "RunOperationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunOperationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunOperationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSearchAuditLogRequest() (request *SearchAuditLogRequest) {
    request = &SearchAuditLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SearchAuditLog")
    
    
    return
}

func NewSearchAuditLogResponse() (response *SearchAuditLogResponse) {
    response = &SearchAuditLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchAuditLog
// 搜索审计日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchAuditLog(request *SearchAuditLogRequest) (response *SearchAuditLogResponse, err error) {
    return c.SearchAuditLogWithContext(context.Background(), request)
}

// SearchAuditLog
// 搜索审计日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchAuditLogWithContext(ctx context.Context, request *SearchAuditLogRequest) (response *SearchAuditLogResponse, err error) {
    if request == nil {
        request = NewSearchAuditLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SearchAuditLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchAuditLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchAuditLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchCommandRequest() (request *SearchCommandRequest) {
    request = &SearchCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SearchCommand")
    
    
    return
}

func NewSearchCommandResponse() (response *SearchCommandResponse) {
    response = &SearchCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchCommand
// 命令执行检索
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchCommand(request *SearchCommandRequest) (response *SearchCommandResponse, err error) {
    return c.SearchCommandWithContext(context.Background(), request)
}

// SearchCommand
// 命令执行检索
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchCommandWithContext(ctx context.Context, request *SearchCommandRequest) (response *SearchCommandResponse, err error) {
    if request == nil {
        request = NewSearchCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SearchCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchCommandResponse()
    err = c.Send(request, response)
    return
}

func NewSearchCommandBySidRequest() (request *SearchCommandBySidRequest) {
    request = &SearchCommandBySidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SearchCommandBySid")
    
    
    return
}

func NewSearchCommandBySidResponse() (response *SearchCommandBySidResponse) {
    response = &SearchCommandBySidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchCommandBySid
// 根据会话Id搜索Command
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchCommandBySid(request *SearchCommandBySidRequest) (response *SearchCommandBySidResponse, err error) {
    return c.SearchCommandBySidWithContext(context.Background(), request)
}

// SearchCommandBySid
// 根据会话Id搜索Command
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchCommandBySidWithContext(ctx context.Context, request *SearchCommandBySidRequest) (response *SearchCommandBySidResponse, err error) {
    if request == nil {
        request = NewSearchCommandBySidRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SearchCommandBySid")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchCommandBySid require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchCommandBySidResponse()
    err = c.Send(request, response)
    return
}

func NewSearchFileRequest() (request *SearchFileRequest) {
    request = &SearchFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SearchFile")
    
    
    return
}

func NewSearchFileResponse() (response *SearchFileResponse) {
    response = &SearchFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchFile
// 文件传输检索
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchFile(request *SearchFileRequest) (response *SearchFileResponse, err error) {
    return c.SearchFileWithContext(context.Background(), request)
}

// SearchFile
// 文件传输检索
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchFileWithContext(ctx context.Context, request *SearchFileRequest) (response *SearchFileResponse, err error) {
    if request == nil {
        request = NewSearchFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SearchFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchFileResponse()
    err = c.Send(request, response)
    return
}

func NewSearchFileBySidRequest() (request *SearchFileBySidRequest) {
    request = &SearchFileBySidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SearchFileBySid")
    
    
    return
}

func NewSearchFileBySidResponse() (response *SearchFileBySidResponse) {
    response = &SearchFileBySidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchFileBySid
// 搜索文件传输会话下文件操作列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchFileBySid(request *SearchFileBySidRequest) (response *SearchFileBySidResponse, err error) {
    return c.SearchFileBySidWithContext(context.Background(), request)
}

// SearchFileBySid
// 搜索文件传输会话下文件操作列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchFileBySidWithContext(ctx context.Context, request *SearchFileBySidRequest) (response *SearchFileBySidResponse, err error) {
    if request == nil {
        request = NewSearchFileBySidRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SearchFileBySid")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchFileBySid require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchFileBySidResponse()
    err = c.Send(request, response)
    return
}

func NewSearchSessionRequest() (request *SearchSessionRequest) {
    request = &SearchSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SearchSession")
    
    
    return
}

func NewSearchSessionResponse() (response *SearchSessionResponse) {
    response = &SearchSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchSession
// 搜索会话
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchSession(request *SearchSessionRequest) (response *SearchSessionResponse, err error) {
    return c.SearchSessionWithContext(context.Background(), request)
}

// SearchSession
// 搜索会话
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchSessionWithContext(ctx context.Context, request *SearchSessionRequest) (response *SearchSessionResponse, err error) {
    if request == nil {
        request = NewSearchSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SearchSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchSessionResponse()
    err = c.Send(request, response)
    return
}

func NewSearchSessionCommandRequest() (request *SearchSessionCommandRequest) {
    request = &SearchSessionCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SearchSessionCommand")
    
    
    return
}

func NewSearchSessionCommandResponse() (response *SearchSessionCommandResponse) {
    response = &SearchSessionCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchSessionCommand
// 命令检索
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchSessionCommand(request *SearchSessionCommandRequest) (response *SearchSessionCommandResponse, err error) {
    return c.SearchSessionCommandWithContext(context.Background(), request)
}

// SearchSessionCommand
// 命令检索
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchSessionCommandWithContext(ctx context.Context, request *SearchSessionCommandRequest) (response *SearchSessionCommandResponse, err error) {
    if request == nil {
        request = NewSearchSessionCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SearchSessionCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchSessionCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchSessionCommandResponse()
    err = c.Send(request, response)
    return
}

func NewSearchSubtaskResultByIdRequest() (request *SearchSubtaskResultByIdRequest) {
    request = &SearchSubtaskResultByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SearchSubtaskResultById")
    
    
    return
}

func NewSearchSubtaskResultByIdResponse() (response *SearchSubtaskResultByIdResponse) {
    response = &SearchSubtaskResultByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchSubtaskResultById
// 查询运维子任务执行结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchSubtaskResultById(request *SearchSubtaskResultByIdRequest) (response *SearchSubtaskResultByIdResponse, err error) {
    return c.SearchSubtaskResultByIdWithContext(context.Background(), request)
}

// SearchSubtaskResultById
// 查询运维子任务执行结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchSubtaskResultByIdWithContext(ctx context.Context, request *SearchSubtaskResultByIdRequest) (response *SearchSubtaskResultByIdResponse, err error) {
    if request == nil {
        request = NewSearchSubtaskResultByIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SearchSubtaskResultById")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchSubtaskResultById require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchSubtaskResultByIdResponse()
    err = c.Send(request, response)
    return
}

func NewSearchTaskResultRequest() (request *SearchTaskResultRequest) {
    request = &SearchTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SearchTaskResult")
    
    
    return
}

func NewSearchTaskResultResponse() (response *SearchTaskResultResponse) {
    response = &SearchTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchTaskResult
// 搜索运维任务执行结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchTaskResult(request *SearchTaskResultRequest) (response *SearchTaskResultResponse, err error) {
    return c.SearchTaskResultWithContext(context.Background(), request)
}

// SearchTaskResult
// 搜索运维任务执行结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchTaskResultWithContext(ctx context.Context, request *SearchTaskResultRequest) (response *SearchTaskResultResponse, err error) {
    if request == nil {
        request = NewSearchTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SearchTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewSetLDAPSyncFlagRequest() (request *SetLDAPSyncFlagRequest) {
    request = &SetLDAPSyncFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SetLDAPSyncFlag")
    
    
    return
}

func NewSetLDAPSyncFlagResponse() (response *SetLDAPSyncFlagResponse) {
    response = &SetLDAPSyncFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetLDAPSyncFlag
// 设置LDAP 立即同步标记
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetLDAPSyncFlag(request *SetLDAPSyncFlagRequest) (response *SetLDAPSyncFlagResponse, err error) {
    return c.SetLDAPSyncFlagWithContext(context.Background(), request)
}

// SetLDAPSyncFlag
// 设置LDAP 立即同步标记
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetLDAPSyncFlagWithContext(ctx context.Context, request *SetLDAPSyncFlagRequest) (response *SetLDAPSyncFlagResponse, err error) {
    if request == nil {
        request = NewSetLDAPSyncFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SetLDAPSyncFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetLDAPSyncFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetLDAPSyncFlagResponse()
    err = c.Send(request, response)
    return
}

func NewSyncDevicesToIOARequest() (request *SyncDevicesToIOARequest) {
    request = &SyncDevicesToIOARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SyncDevicesToIOA")
    
    
    return
}

func NewSyncDevicesToIOAResponse() (response *SyncDevicesToIOAResponse) {
    response = &SyncDevicesToIOAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncDevicesToIOA
// 同步资产到IOA
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncDevicesToIOA(request *SyncDevicesToIOARequest) (response *SyncDevicesToIOAResponse, err error) {
    return c.SyncDevicesToIOAWithContext(context.Background(), request)
}

// SyncDevicesToIOA
// 同步资产到IOA
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncDevicesToIOAWithContext(ctx context.Context, request *SyncDevicesToIOARequest) (response *SyncDevicesToIOAResponse, err error) {
    if request == nil {
        request = NewSyncDevicesToIOARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SyncDevicesToIOA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncDevicesToIOA require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncDevicesToIOAResponse()
    err = c.Send(request, response)
    return
}

func NewSyncUserToIOARequest() (request *SyncUserToIOARequest) {
    request = &SyncUserToIOARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "SyncUserToIOA")
    
    
    return
}

func NewSyncUserToIOAResponse() (response *SyncUserToIOAResponse) {
    response = &SyncUserToIOAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncUserToIOA
// 同步堡垒机本地用户到IOA
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) SyncUserToIOA(request *SyncUserToIOARequest) (response *SyncUserToIOAResponse, err error) {
    return c.SyncUserToIOAWithContext(context.Background(), request)
}

// SyncUserToIOA
// 同步堡垒机本地用户到IOA
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) SyncUserToIOAWithContext(ctx context.Context, request *SyncUserToIOARequest) (response *SyncUserToIOAResponse, err error) {
    if request == nil {
        request = NewSyncUserToIOARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "SyncUserToIOA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncUserToIOA require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncUserToIOAResponse()
    err = c.Send(request, response)
    return
}

func NewUnlockUserRequest() (request *UnlockUserRequest) {
    request = &UnlockUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bh", APIVersion, "UnlockUser")
    
    
    return
}

func NewUnlockUserResponse() (response *UnlockUserResponse) {
    response = &UnlockUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnlockUser
// 解锁用户
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UnlockUser(request *UnlockUserRequest) (response *UnlockUserResponse, err error) {
    return c.UnlockUserWithContext(context.Background(), request)
}

// UnlockUser
// 解锁用户
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UnlockUserWithContext(ctx context.Context, request *UnlockUserRequest) (response *UnlockUserResponse, err error) {
    if request == nil {
        request = NewUnlockUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bh", APIVersion, "UnlockUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnlockUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnlockUserResponse()
    err = c.Send(request, response)
    return
}
