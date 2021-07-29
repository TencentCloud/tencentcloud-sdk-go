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

package v20200226

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-26"

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


func NewQueryAntiFraudRequest() (request *QueryAntiFraudRequest) {
    request = &QueryAntiFraudRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("af", APIVersion, "QueryAntiFraud")
    return
}

func NewQueryAntiFraudResponse() (response *QueryAntiFraudResponse) {
    response = &QueryAntiFraudResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAntiFraud
// 天御反欺诈服务，主要应用于银行、证券、保险、消费金融等金融行业客户，通过腾讯的大数据风控能力，
//
// 可以准确识别恶意用户信息，解决客户在支付、活动、理财，风控等业务环节遇到的欺诈威胁，降低企业
//
// 的损失。
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
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
func (c *Client) QueryAntiFraud(request *QueryAntiFraudRequest) (response *QueryAntiFraudResponse, err error) {
    if request == nil {
        request = NewQueryAntiFraudRequest()
    }
    response = NewQueryAntiFraudResponse()
    err = c.Send(request, response)
    return
}
