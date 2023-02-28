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

package v20191018

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-18"

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


func NewAddDeviceGroupMembersRequest() (request *AddDeviceGroupMembersRequest) {
    request = &AddDeviceGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "AddDeviceGroupMembers")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "AddUserGroupMembers")
    
    
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
func (c *Client) AddUserGroupMembersWithContext(ctx context.Context, request *AddUserGroupMembersRequest) (response *AddUserGroupMembersResponse, err error) {
    if request == nil {
        request = NewAddUserGroupMembersRequest()
    }
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "BindDeviceAccountPassword")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "BindDeviceAccountPrivateKey")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "BindDeviceResource")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDeviceResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindDeviceResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAclRequest() (request *CreateAclRequest) {
    request = &CreateAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "CreateAcl")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "CreateAssetSyncJob")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCmdTemplateRequest() (request *CreateCmdTemplateRequest) {
    request = &CreateCmdTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "CreateCmdTemplate")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "CreateDeviceAccount")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "CreateDeviceGroup")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeviceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "CreateUser")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserGroupRequest() (request *CreateUserGroupRequest) {
    request = &CreateUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "CreateUserGroup")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAclsRequest() (request *DeleteAclsRequest) {
    request = &DeleteAclsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeleteAcls")
    
    
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
func (c *Client) DeleteAclsWithContext(ctx context.Context, request *DeleteAclsRequest) (response *DeleteAclsResponse, err error) {
    if request == nil {
        request = NewDeleteAclsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAcls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAclsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCmdTemplatesRequest() (request *DeleteCmdTemplatesRequest) {
    request = &DeleteCmdTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeleteCmdTemplates")
    
    
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
func (c *Client) DeleteCmdTemplatesWithContext(ctx context.Context, request *DeleteCmdTemplatesRequest) (response *DeleteCmdTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteCmdTemplatesRequest()
    }
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeleteDeviceAccounts")
    
    
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
func (c *Client) DeleteDeviceAccountsWithContext(ctx context.Context, request *DeleteDeviceAccountsRequest) (response *DeleteDeviceAccountsResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceAccountsRequest()
    }
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeleteDeviceGroupMembers")
    
    
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
func (c *Client) DeleteDeviceGroupMembersWithContext(ctx context.Context, request *DeleteDeviceGroupMembersRequest) (response *DeleteDeviceGroupMembersResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceGroupMembersRequest()
    }
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeleteDeviceGroups")
    
    
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
func (c *Client) DeleteDeviceGroupsWithContext(ctx context.Context, request *DeleteDeviceGroupsRequest) (response *DeleteDeviceGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceGroupsRequest()
    }
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeleteDevices")
    
    
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
func (c *Client) DeleteDevicesWithContext(ctx context.Context, request *DeleteDevicesRequest) (response *DeleteDevicesResponse, err error) {
    if request == nil {
        request = NewDeleteDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserGroupMembersRequest() (request *DeleteUserGroupMembersRequest) {
    request = &DeleteUserGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeleteUserGroupMembers")
    
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteUserGroupMembers(request *DeleteUserGroupMembersRequest) (response *DeleteUserGroupMembersResponse, err error) {
    return c.DeleteUserGroupMembersWithContext(context.Background(), request)
}

// DeleteUserGroupMembers
// 删除用户组成员
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteUserGroupMembersWithContext(ctx context.Context, request *DeleteUserGroupMembersRequest) (response *DeleteUserGroupMembersResponse, err error) {
    if request == nil {
        request = NewDeleteUserGroupMembersRequest()
    }
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeleteUserGroups")
    
    
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
func (c *Client) DeleteUserGroupsWithContext(ctx context.Context, request *DeleteUserGroupsRequest) (response *DeleteUserGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteUserGroupsRequest()
    }
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeleteUsers")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DeployResource")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAclsRequest() (request *DescribeAclsRequest) {
    request = &DescribeAclsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeAcls")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAcls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAclsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetSyncStatusRequest() (request *DescribeAssetSyncStatusRequest) {
    request = &DescribeAssetSyncStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeAssetSyncStatus")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetSyncStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetSyncStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmdTemplatesRequest() (request *DescribeCmdTemplatesRequest) {
    request = &DescribeCmdTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeCmdTemplates")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmdTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmdTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDasbImageIdsRequest() (request *DescribeDasbImageIdsRequest) {
    request = &DescribeDasbImageIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbImageIds")
    
    
    return
}

func NewDescribeDasbImageIdsResponse() (response *DescribeDasbImageIdsResponse) {
    response = &DescribeDasbImageIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDasbImageIds
// 获取镜像列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDasbImageIds(request *DescribeDasbImageIdsRequest) (response *DescribeDasbImageIdsResponse, err error) {
    return c.DescribeDasbImageIdsWithContext(context.Background(), request)
}

// DescribeDasbImageIds
// 获取镜像列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDasbImageIdsWithContext(ctx context.Context, request *DescribeDasbImageIdsRequest) (response *DescribeDasbImageIdsResponse, err error) {
    if request == nil {
        request = NewDescribeDasbImageIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDasbImageIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDasbImageIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceAccountsRequest() (request *DescribeDeviceAccountsRequest) {
    request = &DescribeDeviceAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDeviceAccounts")
    
    
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
func (c *Client) DescribeDeviceAccountsWithContext(ctx context.Context, request *DescribeDeviceAccountsRequest) (response *DescribeDeviceAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceAccountsRequest()
    }
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDeviceGroupMembers")
    
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceGroupMembersWithContext(ctx context.Context, request *DescribeDeviceGroupMembersRequest) (response *DescribeDeviceGroupMembersResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceGroupMembersRequest()
    }
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDeviceGroups")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDevices")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
    request = &DescribeResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeResources")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserGroupMembersRequest() (request *DescribeUserGroupMembersRequest) {
    request = &DescribeUserGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeUserGroupMembers")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeUserGroups")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersRequest() (request *DescribeUsersRequest) {
    request = &DescribeUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeUsers")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersResponse()
    err = c.Send(request, response)
    return
}

func NewImportExternalDeviceRequest() (request *ImportExternalDeviceRequest) {
    request = &ImportExternalDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "ImportExternalDevice")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportExternalDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportExternalDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAclRequest() (request *ModifyAclRequest) {
    request = &ModifyAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "ModifyAcl")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAclResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceRequest() (request *ModifyDeviceRequest) {
    request = &ModifyDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "ModifyDevice")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "ModifyDeviceGroup")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "ModifyUser")
    
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewResetDeviceAccountPasswordRequest() (request *ResetDeviceAccountPasswordRequest) {
    request = &ResetDeviceAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dasb", APIVersion, "ResetDeviceAccountPassword")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "ResetDeviceAccountPrivateKey")
    
    
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
    
    request.Init().WithApiInfo("dasb", APIVersion, "ResetUser")
    
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResetUser(request *ResetUserRequest) (response *ResetUserResponse, err error) {
    return c.ResetUserWithContext(context.Background(), request)
}

// ResetUser
// 重置用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResetUserWithContext(ctx context.Context, request *ResetUserRequest) (response *ResetUserResponse, err error) {
    if request == nil {
        request = NewResetUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetUserResponse()
    err = c.Send(request, response)
    return
}
