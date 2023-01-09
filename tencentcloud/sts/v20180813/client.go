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

package v20180813

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-08-13"

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


func NewAssumeRoleRequest() (request *AssumeRoleRequest) {
    request = &AssumeRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sts", APIVersion, "AssumeRole")
    
    
    return
}

func NewAssumeRoleResponse() (response *AssumeRoleResponse) {
    response = &AssumeRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssumeRole
// 申请扮演角色
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INTERNALERROR_PBSERIALIZEERROR = "InternalError.PbSerializeError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERLIMIT = "InvalidParameter.OverLimit"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssumeRole(request *AssumeRoleRequest) (response *AssumeRoleResponse, err error) {
    return c.AssumeRoleWithContext(context.Background(), request)
}

// AssumeRole
// 申请扮演角色
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INTERNALERROR_PBSERIALIZEERROR = "InternalError.PbSerializeError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERLIMIT = "InvalidParameter.OverLimit"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssumeRoleWithContext(ctx context.Context, request *AssumeRoleRequest) (response *AssumeRoleResponse, err error) {
    if request == nil {
        request = NewAssumeRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssumeRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssumeRoleResponse()
    err = c.Send(request, response)
    return
}

func NewAssumeRoleWithSAMLRequest() (request *AssumeRoleWithSAMLRequest) {
    request = &AssumeRoleWithSAMLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.SetSkipSign(true)
    request.Init().WithApiInfo("sts", APIVersion, "AssumeRoleWithSAML")
    
    
    return
}

func NewAssumeRoleWithSAMLResponse() (response *AssumeRoleWithSAMLResponse) {
    response = &AssumeRoleWithSAMLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssumeRoleWithSAML
// 本接口（AssumeRoleWithSAML）用于根据 SAML 断言申请角色临时凭证。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INTERNALERROR_PBSERIALIZEERROR = "InternalError.PbSerializeError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRoleWithSAML(request *AssumeRoleWithSAMLRequest) (response *AssumeRoleWithSAMLResponse, err error) {
    return c.AssumeRoleWithSAMLWithContext(context.Background(), request)
}

// AssumeRoleWithSAML
// 本接口（AssumeRoleWithSAML）用于根据 SAML 断言申请角色临时凭证。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INTERNALERROR_PBSERIALIZEERROR = "InternalError.PbSerializeError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRoleWithSAMLWithContext(ctx context.Context, request *AssumeRoleWithSAMLRequest) (response *AssumeRoleWithSAMLResponse, err error) {
    if request == nil {
        request = NewAssumeRoleWithSAMLRequest()
    }
    
    request.SetContext(ctx)
    
    response = NewAssumeRoleWithSAMLResponse()
    err = c.Send(request, response)
    return
}

func NewAssumeRoleWithWebIdentityRequest() (request *AssumeRoleWithWebIdentityRequest) {
    request = &AssumeRoleWithWebIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.SetSkipSign(true)
    request.Init().WithApiInfo("sts", APIVersion, "AssumeRoleWithWebIdentity")
    
    
    return
}

func NewAssumeRoleWithWebIdentityResponse() (response *AssumeRoleWithWebIdentityResponse) {
    response = &AssumeRoleWithWebIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssumeRoleWithWebIdentity
// 申请OIDC角色临时密钥
//
// 可能返回的错误码:
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_OVERLIMIT = "InvalidParameter.OverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_WEBIDENTITYTOKENERROR = "InvalidParameter.WebIdentityTokenError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRoleWithWebIdentity(request *AssumeRoleWithWebIdentityRequest) (response *AssumeRoleWithWebIdentityResponse, err error) {
    return c.AssumeRoleWithWebIdentityWithContext(context.Background(), request)
}

// AssumeRoleWithWebIdentity
// 申请OIDC角色临时密钥
//
// 可能返回的错误码:
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_OVERLIMIT = "InvalidParameter.OverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_WEBIDENTITYTOKENERROR = "InvalidParameter.WebIdentityTokenError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRoleWithWebIdentityWithContext(ctx context.Context, request *AssumeRoleWithWebIdentityRequest) (response *AssumeRoleWithWebIdentityResponse, err error) {
    if request == nil {
        request = NewAssumeRoleWithWebIdentityRequest()
    }
    
    request.SetContext(ctx)
    
    response = NewAssumeRoleWithWebIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewGetCallerIdentityRequest() (request *GetCallerIdentityRequest) {
    request = &GetCallerIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sts", APIVersion, "GetCallerIdentity")
    
    
    return
}

func NewGetCallerIdentityResponse() (response *GetCallerIdentityResponse) {
    response = &GetCallerIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCallerIdentity
// 获取当前调用者的身份信息。
//
// 接口支持主账号，子账号长期密钥以及AssumeRole，GetFederationToken生成的临时凭据的身份获取。
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSKEYILLEGAL = "AuthFailure.AccessKeyIllegal"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INVALIDPARAMETER_ACCESSKEYNOTSUPPORT = "InvalidParameter.AccessKeyNotSupport"
func (c *Client) GetCallerIdentity(request *GetCallerIdentityRequest) (response *GetCallerIdentityResponse, err error) {
    return c.GetCallerIdentityWithContext(context.Background(), request)
}

// GetCallerIdentity
// 获取当前调用者的身份信息。
//
// 接口支持主账号，子账号长期密钥以及AssumeRole，GetFederationToken生成的临时凭据的身份获取。
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSKEYILLEGAL = "AuthFailure.AccessKeyIllegal"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INVALIDPARAMETER_ACCESSKEYNOTSUPPORT = "InvalidParameter.AccessKeyNotSupport"
func (c *Client) GetCallerIdentityWithContext(ctx context.Context, request *GetCallerIdentityRequest) (response *GetCallerIdentityResponse, err error) {
    if request == nil {
        request = NewGetCallerIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCallerIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCallerIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewGetFederationTokenRequest() (request *GetFederationTokenRequest) {
    request = &GetFederationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sts", APIVersion, "GetFederationToken")
    
    
    return
}

func NewGetFederationTokenResponse() (response *GetFederationTokenResponse) {
    response = &GetFederationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFederationToken
// 获取联合身份临时访问凭证
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetFederationToken(request *GetFederationTokenRequest) (response *GetFederationTokenResponse, err error) {
    return c.GetFederationTokenWithContext(context.Background(), request)
}

// GetFederationToken
// 获取联合身份临时访问凭证
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetFederationTokenWithContext(ctx context.Context, request *GetFederationTokenRequest) (response *GetFederationTokenResponse, err error) {
    if request == nil {
        request = NewGetFederationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFederationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFederationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewQueryApiKeyRequest() (request *QueryApiKeyRequest) {
    request = &QueryApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sts", APIVersion, "QueryApiKey")
    
    
    return
}

func NewQueryApiKeyResponse() (response *QueryApiKeyResponse) {
    response = &QueryApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryApiKey
// 拉取API密钥列表
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) QueryApiKey(request *QueryApiKeyRequest) (response *QueryApiKeyResponse, err error) {
    return c.QueryApiKeyWithContext(context.Background(), request)
}

// QueryApiKey
// 拉取API密钥列表
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) QueryApiKeyWithContext(ctx context.Context, request *QueryApiKeyRequest) (response *QueryApiKeyResponse, err error) {
    if request == nil {
        request = NewQueryApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryApiKeyResponse()
    err = c.Send(request, response)
    return
}
