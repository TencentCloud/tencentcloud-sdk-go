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

package v20220331

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-03-31"

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


func NewCreateApiImportUserJobRequest() (request *CreateApiImportUserJobRequest) {
    request = &CreateApiImportUserJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "CreateApiImportUserJob")
    
    
    return
}

func NewCreateApiImportUserJobResponse() (response *CreateApiImportUserJobResponse) {
    response = &CreateApiImportUserJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApiImportUserJob
// 新建接口导入用户任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApiImportUserJob(request *CreateApiImportUserJobRequest) (response *CreateApiImportUserJobResponse, err error) {
    return c.CreateApiImportUserJobWithContext(context.Background(), request)
}

// CreateApiImportUserJob
// 新建接口导入用户任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApiImportUserJobWithContext(ctx context.Context, request *CreateApiImportUserJobRequest) (response *CreateApiImportUserJobResponse, err error) {
    if request == nil {
        request = NewCreateApiImportUserJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiImportUserJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiImportUserJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileExportUserJobRequest() (request *CreateFileExportUserJobRequest) {
    request = &CreateFileExportUserJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "CreateFileExportUserJob")
    
    
    return
}

func NewCreateFileExportUserJobResponse() (response *CreateFileExportUserJobResponse) {
    response = &CreateFileExportUserJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFileExportUserJob
// 新建文件导出用户任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFileExportUserJob(request *CreateFileExportUserJobRequest) (response *CreateFileExportUserJobResponse, err error) {
    return c.CreateFileExportUserJobWithContext(context.Background(), request)
}

// CreateFileExportUserJob
// 新建文件导出用户任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFileExportUserJobWithContext(ctx context.Context, request *CreateFileExportUserJobRequest) (response *CreateFileExportUserJobResponse, err error) {
    if request == nil {
        request = NewCreateFileExportUserJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileExportUserJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileExportUserJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUser
// 创建用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ATTRIBUTEFORMATERROR = "FailedOperation.AttributeFormatError"
//  FAILEDOPERATION_EMAILALREADYEXISTS = "FailedOperation.EmailAlreadyExists"
//  FAILEDOPERATION_EMAILISNULL = "FailedOperation.EmailIsNull"
//  FAILEDOPERATION_INDEXEDATTRIBUTETOOLONG = "FailedOperation.IndexedAttributeTooLong"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PASSWORDISNULL = "FailedOperation.PasswordIsNull"
//  FAILEDOPERATION_PHONENUMBERALREADYEXISTS = "FailedOperation.PhoneNumberAlreadyExists"
//  FAILEDOPERATION_PHONENUMBERISNULL = "FailedOperation.PhoneNumberIsNull"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERNAMEISNULL = "FailedOperation.UserNameIsNull"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 创建用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ATTRIBUTEFORMATERROR = "FailedOperation.AttributeFormatError"
//  FAILEDOPERATION_EMAILALREADYEXISTS = "FailedOperation.EmailAlreadyExists"
//  FAILEDOPERATION_EMAILISNULL = "FailedOperation.EmailIsNull"
//  FAILEDOPERATION_INDEXEDATTRIBUTETOOLONG = "FailedOperation.IndexedAttributeTooLong"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PASSWORDISNULL = "FailedOperation.PasswordIsNull"
//  FAILEDOPERATION_PHONENUMBERALREADYEXISTS = "FailedOperation.PhoneNumberAlreadyExists"
//  FAILEDOPERATION_PHONENUMBERISNULL = "FailedOperation.PhoneNumberIsNull"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERNAMEISNULL = "FailedOperation.UserNameIsNull"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
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
    
    request.Init().WithApiInfo("ciam", APIVersion, "CreateUserGroup")
    
    
    return
}

func NewCreateUserGroupResponse() (response *CreateUserGroupResponse) {
    response = &CreateUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserGroup
// 创建用户组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNAMEALREADYEXISTS = "FailedOperation.UserGroupNameAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) CreateUserGroup(request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    return c.CreateUserGroupWithContext(context.Background(), request)
}

// CreateUserGroup
// 创建用户组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNAMEALREADYEXISTS = "FailedOperation.UserGroupNameAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
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

func NewCreateUserStoreRequest() (request *CreateUserStoreRequest) {
    request = &CreateUserStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "CreateUserStore")
    
    
    return
}

func NewCreateUserStoreResponse() (response *CreateUserStoreResponse) {
    response = &CreateUserStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserStore
// 创建用户目录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREALREADYEXISTS = "FailedOperation.UserStoreAlreadyExists"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_USERSTORECREATESECRETFAIL = "InternalError.UserStoreCreateSecretFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) CreateUserStore(request *CreateUserStoreRequest) (response *CreateUserStoreResponse, err error) {
    return c.CreateUserStoreWithContext(context.Background(), request)
}

// CreateUserStore
// 创建用户目录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREALREADYEXISTS = "FailedOperation.UserStoreAlreadyExists"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_USERSTORECREATESECRETFAIL = "InternalError.UserStoreCreateSecretFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) CreateUserStoreWithContext(ctx context.Context, request *CreateUserStoreRequest) (response *CreateUserStoreResponse, err error) {
    if request == nil {
        request = NewCreateUserStoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserStoreResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserGroupsRequest() (request *DeleteUserGroupsRequest) {
    request = &DeleteUserGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "DeleteUserGroups")
    
    
    return
}

func NewDeleteUserGroupsResponse() (response *DeleteUserGroupsResponse) {
    response = &DeleteUserGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUserGroups
// 批量删除用户组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) DeleteUserGroups(request *DeleteUserGroupsRequest) (response *DeleteUserGroupsResponse, err error) {
    return c.DeleteUserGroupsWithContext(context.Background(), request)
}

// DeleteUserGroups
// 批量删除用户组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
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

func NewDeleteUserStoreRequest() (request *DeleteUserStoreRequest) {
    request = &DeleteUserStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "DeleteUserStore")
    
    
    return
}

func NewDeleteUserStoreResponse() (response *DeleteUserStoreResponse) {
    response = &DeleteUserStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUserStore
// 删除用户目录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  FAILEDOPERATION_USERSTORENOTEXIST = "FailedOperation.UserStoreNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_USERSTOREDELETEUSERFAIL = "InternalError.UserStoreDeleteUserFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) DeleteUserStore(request *DeleteUserStoreRequest) (response *DeleteUserStoreResponse, err error) {
    return c.DeleteUserStoreWithContext(context.Background(), request)
}

// DeleteUserStore
// 删除用户目录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  FAILEDOPERATION_USERSTORENOTEXIST = "FailedOperation.UserStoreNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_USERSTOREDELETEUSERFAIL = "InternalError.UserStoreDeleteUserFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) DeleteUserStoreWithContext(ctx context.Context, request *DeleteUserStoreRequest) (response *DeleteUserStoreResponse, err error) {
    if request == nil {
        request = NewDeleteUserStoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserStoreResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
    request = &DeleteUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "DeleteUsers")
    
    
    return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
    response = &DeleteUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUsers
// 批量删除用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    return c.DeleteUsersWithContext(context.Background(), request)
}

// DeleteUsers
// 批量删除用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeUserRequest() (request *DescribeUserRequest) {
    request = &DescribeUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "DescribeUser")
    
    
    return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
    response = &DescribeUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUser
// 多条件查询用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUERYUSERSPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersParameterMustInWhitelist"
//  FAILEDOPERATION_QUERYUSERSPARAMETERREPEAT = "FailedOperation.QueryUsersParameterRepeat"
//  FAILEDOPERATION_QUERYUSERSSORTPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersSortParameterMustInWhitelist"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    return c.DescribeUserWithContext(context.Background(), request)
}

// DescribeUser
// 多条件查询用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUERYUSERSPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersParameterMustInWhitelist"
//  FAILEDOPERATION_QUERYUSERSPARAMETERREPEAT = "FailedOperation.QueryUsersParameterRepeat"
//  FAILEDOPERATION_QUERYUSERSSORTPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersSortParameterMustInWhitelist"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserWithContext(ctx context.Context, request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    if request == nil {
        request = NewDescribeUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserByIdRequest() (request *DescribeUserByIdRequest) {
    request = &DescribeUserByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "DescribeUserById")
    
    
    return
}

func NewDescribeUserByIdResponse() (response *DescribeUserByIdResponse) {
    response = &DescribeUserByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserById
// 根据ID查询用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserById(request *DescribeUserByIdRequest) (response *DescribeUserByIdResponse, err error) {
    return c.DescribeUserByIdWithContext(context.Background(), request)
}

// DescribeUserById
// 根据ID查询用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserByIdWithContext(ctx context.Context, request *DescribeUserByIdRequest) (response *DescribeUserByIdResponse, err error) {
    if request == nil {
        request = NewDescribeUserByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserByIdResponse()
    err = c.Send(request, response)
    return
}

func NewLinkAccountRequest() (request *LinkAccountRequest) {
    request = &LinkAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "LinkAccount")
    
    
    return
}

func NewLinkAccountResponse() (response *LinkAccountResponse) {
    response = &LinkAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LinkAccount
// 账号融合
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYBELINKED = "FailedOperation.AccountAlreadyBeLinked"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTATUSENUM = "FailedOperation.InvalidUserStatusEnum"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PRIMARYUSERNOTFOUND = "FailedOperation.PrimaryUserNotFound"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_SECONDARYUSERNOTFOUND = "FailedOperation.SecondaryUserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LinkAccount(request *LinkAccountRequest) (response *LinkAccountResponse, err error) {
    return c.LinkAccountWithContext(context.Background(), request)
}

// LinkAccount
// 账号融合
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYBELINKED = "FailedOperation.AccountAlreadyBeLinked"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTATUSENUM = "FailedOperation.InvalidUserStatusEnum"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PRIMARYUSERNOTFOUND = "FailedOperation.PrimaryUserNotFound"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_SECONDARYUSERNOTFOUND = "FailedOperation.SecondaryUserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LinkAccountWithContext(ctx context.Context, request *LinkAccountRequest) (response *LinkAccountResponse, err error) {
    if request == nil {
        request = NewLinkAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LinkAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewLinkAccountResponse()
    err = c.Send(request, response)
    return
}

func NewListJobsRequest() (request *ListJobsRequest) {
    request = &ListJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListJobs")
    
    
    return
}

func NewListJobsResponse() (response *ListJobsResponse) {
    response = &ListJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListJobs
// 查询任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListJobs(request *ListJobsRequest) (response *ListJobsResponse, err error) {
    return c.ListJobsWithContext(context.Background(), request)
}

// ListJobs
// 查询任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListJobsWithContext(ctx context.Context, request *ListJobsRequest) (response *ListJobsResponse, err error) {
    if request == nil {
        request = NewListJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListJobsResponse()
    err = c.Send(request, response)
    return
}

func NewListLogMessageByConditionRequest() (request *ListLogMessageByConditionRequest) {
    request = &ListLogMessageByConditionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListLogMessageByCondition")
    
    
    return
}

func NewListLogMessageByConditionResponse() (response *ListLogMessageByConditionResponse) {
    response = &ListLogMessageByConditionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListLogMessageByCondition
// 查询日志信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListLogMessageByCondition(request *ListLogMessageByConditionRequest) (response *ListLogMessageByConditionResponse, err error) {
    return c.ListLogMessageByConditionWithContext(context.Background(), request)
}

// ListLogMessageByCondition
// 查询日志信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListLogMessageByConditionWithContext(ctx context.Context, request *ListLogMessageByConditionRequest) (response *ListLogMessageByConditionResponse, err error) {
    if request == nil {
        request = NewListLogMessageByConditionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListLogMessageByCondition require credential")
    }

    request.SetContext(ctx)
    
    response = NewListLogMessageByConditionResponse()
    err = c.Send(request, response)
    return
}

func NewListUserRequest() (request *ListUserRequest) {
    request = &ListUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListUser")
    
    
    return
}

func NewListUserResponse() (response *ListUserResponse) {
    response = &ListUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUser
// 查询用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListUser(request *ListUserRequest) (response *ListUserResponse, err error) {
    return c.ListUserWithContext(context.Background(), request)
}

// ListUser
// 查询用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListUserWithContext(ctx context.Context, request *ListUserRequest) (response *ListUserResponse, err error) {
    if request == nil {
        request = NewListUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserResponse()
    err = c.Send(request, response)
    return
}

func NewListUserByPropertyRequest() (request *ListUserByPropertyRequest) {
    request = &ListUserByPropertyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListUserByProperty")
    
    
    return
}

func NewListUserByPropertyResponse() (response *ListUserByPropertyResponse) {
    response = &ListUserByPropertyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUserByProperty
// 根据属性查询用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListUserByProperty(request *ListUserByPropertyRequest) (response *ListUserByPropertyResponse, err error) {
    return c.ListUserByPropertyWithContext(context.Background(), request)
}

// ListUserByProperty
// 根据属性查询用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListUserByPropertyWithContext(ctx context.Context, request *ListUserByPropertyRequest) (response *ListUserByPropertyResponse, err error) {
    if request == nil {
        request = NewListUserByPropertyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserByProperty require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserByPropertyResponse()
    err = c.Send(request, response)
    return
}

func NewListUserGroupsRequest() (request *ListUserGroupsRequest) {
    request = &ListUserGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListUserGroups")
    
    
    return
}

func NewListUserGroupsResponse() (response *ListUserGroupsResponse) {
    response = &ListUserGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUserGroups
// 查询用户组列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) ListUserGroups(request *ListUserGroupsRequest) (response *ListUserGroupsResponse, err error) {
    return c.ListUserGroupsWithContext(context.Background(), request)
}

// ListUserGroups
// 查询用户组列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) ListUserGroupsWithContext(ctx context.Context, request *ListUserGroupsRequest) (response *ListUserGroupsResponse, err error) {
    if request == nil {
        request = NewListUserGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListUserStoreRequest() (request *ListUserStoreRequest) {
    request = &ListUserStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListUserStore")
    
    
    return
}

func NewListUserStoreResponse() (response *ListUserStoreResponse) {
    response = &ListUserStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUserStore
// 查询用户目录列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) ListUserStore(request *ListUserStoreRequest) (response *ListUserStoreResponse, err error) {
    return c.ListUserStoreWithContext(context.Background(), request)
}

// ListUserStore
// 查询用户目录列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) ListUserStoreWithContext(ctx context.Context, request *ListUserStoreRequest) (response *ListUserStoreResponse, err error) {
    if request == nil {
        request = NewListUserStoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserStoreResponse()
    err = c.Send(request, response)
    return
}

func NewResetPasswordRequest() (request *ResetPasswordRequest) {
    request = &ResetPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ResetPassword")
    
    
    return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
    response = &ResetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetPassword
// 重置用户密码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetPassword(request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    return c.ResetPasswordWithContext(context.Background(), request)
}

// ResetPassword
// 重置用户密码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewSetPasswordRequest() (request *SetPasswordRequest) {
    request = &SetPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "SetPassword")
    
    
    return
}

func NewSetPasswordResponse() (response *SetPasswordResponse) {
    response = &SetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetPassword
// 设置用户密码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PASSWORDISNULL = "FailedOperation.PasswordIsNull"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetPassword(request *SetPasswordRequest) (response *SetPasswordResponse, err error) {
    return c.SetPasswordWithContext(context.Background(), request)
}

// SetPassword
// 设置用户密码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PASSWORDISNULL = "FailedOperation.PasswordIsNull"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetPasswordWithContext(ctx context.Context, request *SetPasswordRequest) (response *SetPasswordResponse, err error) {
    if request == nil {
        request = NewSetPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserRequest() (request *UpdateUserRequest) {
    request = &UpdateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "UpdateUser")
    
    
    return
}

func NewUpdateUserResponse() (response *UpdateUserResponse) {
    response = &UpdateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateUser
// 更新用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ATTRIBUTEFORMATERROR = "FailedOperation.AttributeFormatError"
//  FAILEDOPERATION_EMAILALREADYEXISTS = "FailedOperation.EmailAlreadyExists"
//  FAILEDOPERATION_INDEXEDATTRIBUTETOOLONG = "FailedOperation.IndexedAttributeTooLong"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PHONENUMBERALREADYEXISTS = "FailedOperation.PhoneNumberAlreadyExists"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    return c.UpdateUserWithContext(context.Background(), request)
}

// UpdateUser
// 更新用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ATTRIBUTEFORMATERROR = "FailedOperation.AttributeFormatError"
//  FAILEDOPERATION_EMAILALREADYEXISTS = "FailedOperation.EmailAlreadyExists"
//  FAILEDOPERATION_INDEXEDATTRIBUTETOOLONG = "FailedOperation.IndexedAttributeTooLong"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PHONENUMBERALREADYEXISTS = "FailedOperation.PhoneNumberAlreadyExists"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    if request == nil {
        request = NewUpdateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserGroupRequest() (request *UpdateUserGroupRequest) {
    request = &UpdateUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "UpdateUserGroup")
    
    
    return
}

func NewUpdateUserGroupResponse() (response *UpdateUserGroupResponse) {
    response = &UpdateUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateUserGroup
// 更新用户组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNAMEALREADYEXISTS = "FailedOperation.UserGroupNameAlreadyExists"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (response *UpdateUserGroupResponse, err error) {
    return c.UpdateUserGroupWithContext(context.Background(), request)
}

// UpdateUserGroup
// 更新用户组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNAMEALREADYEXISTS = "FailedOperation.UserGroupNameAlreadyExists"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) UpdateUserGroupWithContext(ctx context.Context, request *UpdateUserGroupRequest) (response *UpdateUserGroupResponse, err error) {
    if request == nil {
        request = NewUpdateUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserStatusRequest() (request *UpdateUserStatusRequest) {
    request = &UpdateUserStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "UpdateUserStatus")
    
    
    return
}

func NewUpdateUserStatusResponse() (response *UpdateUserStatusResponse) {
    response = &UpdateUserStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateUserStatus
// 更新用户状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTATUSENUM = "FailedOperation.InvalidUserStatusEnum"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  FAILEDOPERATION_USERSTATUSREQUIRED = "FailedOperation.UserStatusRequired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateUserStatus(request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
    return c.UpdateUserStatusWithContext(context.Background(), request)
}

// UpdateUserStatus
// 更新用户状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTATUSENUM = "FailedOperation.InvalidUserStatusEnum"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  FAILEDOPERATION_USERSTATUSREQUIRED = "FailedOperation.UserStatusRequired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateUserStatusWithContext(ctx context.Context, request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
    if request == nil {
        request = NewUpdateUserStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserStoreRequest() (request *UpdateUserStoreRequest) {
    request = &UpdateUserStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "UpdateUserStore")
    
    
    return
}

func NewUpdateUserStoreResponse() (response *UpdateUserStoreResponse) {
    response = &UpdateUserStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateUserStore
// 更新用户目录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREALREADYEXISTS = "FailedOperation.UserStoreAlreadyExists"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  FAILEDOPERATION_USERSTORENOTEXIST = "FailedOperation.UserStoreNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) UpdateUserStore(request *UpdateUserStoreRequest) (response *UpdateUserStoreResponse, err error) {
    return c.UpdateUserStoreWithContext(context.Background(), request)
}

// UpdateUserStore
// 更新用户目录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"
//  FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"
//  FAILEDOPERATION_USERSTOREALREADYEXISTS = "FailedOperation.UserStoreAlreadyExists"
//  FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"
//  FAILEDOPERATION_USERSTORENOTEXIST = "FailedOperation.UserStoreNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"
//  UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"
//  UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"
//  UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"
//  UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
func (c *Client) UpdateUserStoreWithContext(ctx context.Context, request *UpdateUserStoreRequest) (response *UpdateUserStoreResponse, err error) {
    if request == nil {
        request = NewUpdateUserStoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserStoreResponse()
    err = c.Send(request, response)
    return
}
