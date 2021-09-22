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

package v20201103

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-11-03"

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


func NewDescribeRiskModelRequest() (request *DescribeRiskModelRequest) {
    request = &DescribeRiskModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rce", APIVersion, "DescribeRiskModel")
    return
}

func NewDescribeRiskModelResponse() (response *DescribeRiskModelResponse) {
    response = &DescribeRiskModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskModel
// 依托人工智能技术和腾讯丰富的风控实战经验，根据用户提供的数据和业务场景，给客户提供定制化模型服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CAPSIGERROR = "InvalidParameter.CapSigError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_URLERROR = "InvalidParameter.UrlError"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADBODY = "InvalidParameterValue.BadBody"
//  INVALIDPARAMETERVALUE_BODYTOOLARGE = "InvalidParameterValue.BodyTooLarge"
//  INVALIDPARAMETERVALUE_CAPMISMATCH = "InvalidParameterValue.CapMisMatch"
//  INVALIDPARAMETERVALUE_HTTPMETHODERROR = "InvalidParameterValue.HttpMethodError"
//  INVALIDPARAMETERVALUE_INVALIDDATE = "InvalidParameterValue.InvalidDate"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDNUM = "InvalidParameterValue.InvalidNum"
//  INVALIDPARAMETERVALUE_INVALIDSRVID = "InvalidParameterValue.InvalidSrvId"
//  INVALIDPARAMETERVALUE_INVALIDSTRIDE = "InvalidParameterValue.InvalidStride"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FREQCNT = "LimitExceeded.FreqCnt"
//  LIMITEXCEEDED_IPFREQCNT = "LimitExceeded.IpFreqCnt"
//  LIMITEXCEEDED_KEYFREQCNT = "LimitExceeded.KeyFreqCnt"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskModel(request *DescribeRiskModelRequest) (response *DescribeRiskModelResponse, err error) {
    if request == nil {
        request = NewDescribeRiskModelRequest()
    }
    response = NewDescribeRiskModelResponse()
    err = c.Send(request, response)
    return
}

func NewManageMarketingRiskRequest() (request *ManageMarketingRiskRequest) {
    request = &ManageMarketingRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rce", APIVersion, "ManageMarketingRisk")
    return
}

func NewManageMarketingRiskResponse() (response *ManageMarketingRiskResponse) {
    response = &ManageMarketingRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManageMarketingRisk
// 全栈式风控引擎（RiskControlEngine，RCE）是基于人工智能技术和腾讯20年风控实战沉淀，依托腾讯海量业务构建的风控引擎，以轻量级的 SaaS 服务方式接入，帮助您快速解决注册、登录、营销活动等关键场景遇到的欺诈问题，实时防御黑灰产作恶。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CAPSIGERROR = "InvalidParameter.CapSigError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_URLERROR = "InvalidParameter.UrlError"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADBODY = "InvalidParameterValue.BadBody"
//  INVALIDPARAMETERVALUE_BODYTOOLARGE = "InvalidParameterValue.BodyTooLarge"
//  INVALIDPARAMETERVALUE_CAPMISMATCH = "InvalidParameterValue.CapMisMatch"
//  INVALIDPARAMETERVALUE_HTTPMETHODERROR = "InvalidParameterValue.HttpMethodError"
//  INVALIDPARAMETERVALUE_INVALIDDATE = "InvalidParameterValue.InvalidDate"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDNUM = "InvalidParameterValue.InvalidNum"
//  INVALIDPARAMETERVALUE_INVALIDSRVID = "InvalidParameterValue.InvalidSrvId"
//  INVALIDPARAMETERVALUE_INVALIDSTRIDE = "InvalidParameterValue.InvalidStride"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FREQCNT = "LimitExceeded.FreqCnt"
//  LIMITEXCEEDED_IPFREQCNT = "LimitExceeded.IpFreqCnt"
//  LIMITEXCEEDED_KEYFREQCNT = "LimitExceeded.KeyFreqCnt"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageMarketingRisk(request *ManageMarketingRiskRequest) (response *ManageMarketingRiskResponse, err error) {
    if request == nil {
        request = NewManageMarketingRiskRequest()
    }
    response = NewManageMarketingRiskResponse()
    err = c.Send(request, response)
    return
}
