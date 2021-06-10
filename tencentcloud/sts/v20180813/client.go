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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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
func (c *Client) AssumeRole(request *AssumeRoleRequest) (response *AssumeRoleResponse, err error) {
    if request == nil {
        request = NewAssumeRoleRequest()
    }
    response = NewAssumeRoleResponse()
    err = c.Send(request, response)
    return
}

func NewAssumeRoleWithSAMLRequest() (request *AssumeRoleWithSAMLRequest) {
    request = &AssumeRoleWithSAMLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
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
    if request == nil {
        request = NewAssumeRoleWithSAMLRequest()
    }
    response = NewAssumeRoleWithSAMLResponse()
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
    if request == nil {
        request = NewGetFederationTokenRequest()
    }
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
    if request == nil {
        request = NewQueryApiKeyRequest()
    }
    response = NewQueryApiKeyResponse()
    err = c.Send(request, response)
    return
}
