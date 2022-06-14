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
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
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
//  FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"
//  FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
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
