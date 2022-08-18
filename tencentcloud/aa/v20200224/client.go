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

package v20200224

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-24"

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


func NewQueryActivityAntiRushRequest() (request *QueryActivityAntiRushRequest) {
    request = &QueryActivityAntiRushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aa", APIVersion, "QueryActivityAntiRush")
    
    
    return
}

func NewQueryActivityAntiRushResponse() (response *QueryActivityAntiRushResponse) {
    response = &QueryActivityAntiRushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryActivityAntiRush
// 腾讯云活动防刷（ActivityAntiRush，AA）是针对电商、O2O、P2P、游戏、支付等行业在促销活动中遇到“羊毛党”恶意刷取优惠福利的行为时，通过防刷引擎，精准识别出“薅羊毛”恶意行为的活动防刷服务，避免了企业被刷带来的巨大经济损失。
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INVALIDPARAMETER_CAPSIGERROR = "InvalidParameter.CapSigError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_URLERROR = "InvalidParameter.UrlError"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_BADBODY = "InvalidParameterValue.BadBody"
//  INVALIDPARAMETERVALUE_BODYTOOLARGE = "InvalidParameterValue.BodyTooLarge"
//  INVALIDPARAMETERVALUE_CAPMISMATCH = "InvalidParameterValue.CapMisMatch"
//  INVALIDPARAMETERVALUE_HTTPMETHODERROR = "InvalidParameterValue.HttpMethodError"
//  LIMITEXCEEDED_FREQCNT = "LimitExceeded.FreqCnt"
//  LIMITEXCEEDED_IPFREQCNT = "LimitExceeded.IpFreqCnt"
//  LIMITEXCEEDED_KEYFREQCNT = "LimitExceeded.KeyFreqCnt"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
func (c *Client) QueryActivityAntiRush(request *QueryActivityAntiRushRequest) (response *QueryActivityAntiRushResponse, err error) {
    return c.QueryActivityAntiRushWithContext(context.Background(), request)
}

// QueryActivityAntiRush
// 腾讯云活动防刷（ActivityAntiRush，AA）是针对电商、O2O、P2P、游戏、支付等行业在促销活动中遇到“羊毛党”恶意刷取优惠福利的行为时，通过防刷引擎，精准识别出“薅羊毛”恶意行为的活动防刷服务，避免了企业被刷带来的巨大经济损失。
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INVALIDPARAMETER_CAPSIGERROR = "InvalidParameter.CapSigError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_URLERROR = "InvalidParameter.UrlError"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_BADBODY = "InvalidParameterValue.BadBody"
//  INVALIDPARAMETERVALUE_BODYTOOLARGE = "InvalidParameterValue.BodyTooLarge"
//  INVALIDPARAMETERVALUE_CAPMISMATCH = "InvalidParameterValue.CapMisMatch"
//  INVALIDPARAMETERVALUE_HTTPMETHODERROR = "InvalidParameterValue.HttpMethodError"
//  LIMITEXCEEDED_FREQCNT = "LimitExceeded.FreqCnt"
//  LIMITEXCEEDED_IPFREQCNT = "LimitExceeded.IpFreqCnt"
//  LIMITEXCEEDED_KEYFREQCNT = "LimitExceeded.KeyFreqCnt"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
func (c *Client) QueryActivityAntiRushWithContext(ctx context.Context, request *QueryActivityAntiRushRequest) (response *QueryActivityAntiRushResponse, err error) {
    if request == nil {
        request = NewQueryActivityAntiRushRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryActivityAntiRush require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryActivityAntiRushResponse()
    err = c.Send(request, response)
    return
}
