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
    "context"
    "errors"
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


func NewDescribeRiskAssessmentRequest() (request *DescribeRiskAssessmentRequest) {
    request = &DescribeRiskAssessmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "DescribeRiskAssessment")
    
    
    return
}

func NewDescribeRiskAssessmentResponse() (response *DescribeRiskAssessmentResponse) {
    response = &DescribeRiskAssessmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskAssessment
// 此接口用于查询风险评估结果
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
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskAssessment(request *DescribeRiskAssessmentRequest) (response *DescribeRiskAssessmentResponse, err error) {
    return c.DescribeRiskAssessmentWithContext(context.Background(), request)
}

// DescribeRiskAssessment
// 此接口用于查询风险评估结果
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
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskAssessmentWithContext(ctx context.Context, request *DescribeRiskAssessmentRequest) (response *DescribeRiskAssessmentResponse, err error) {
    if request == nil {
        request = NewDescribeRiskAssessmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskAssessment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskAssessmentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskTrendsRequest() (request *DescribeRiskTrendsRequest) {
    request = &DescribeRiskTrendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "DescribeRiskTrends")
    
    
    return
}

func NewDescribeRiskTrendsResponse() (response *DescribeRiskTrendsResponse) {
    response = &DescribeRiskTrendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskTrends
// 以图表形式展示三种请求状态的趋势变化
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INVALIDPARAMETER_CAPSIGERROR = "InvalidParameter.CapSigError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_URLERROR = "InvalidParameter.UrlError"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_BADBODY = "InvalidParameterValue.BadBody"
//  INVALIDPARAMETERVALUE_BODYTOOLARGE = "InvalidParameterValue.BodyTooLarge"
//  INVALIDPARAMETERVALUE_CAPMISMATCH = "InvalidParameterValue.CapMisMatch"
//  INVALIDPARAMETERVALUE_HTTPMETHODERROR = "InvalidParameterValue.HttpMethodError"
//  INVALIDPARAMETERVALUE_INVALIDDATE = "InvalidParameterValue.InvalidDate"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDNUM = "InvalidParameterValue.InvalidNum"
//  INVALIDPARAMETERVALUE_INVALIDSRVID = "InvalidParameterValue.InvalidSrvId"
//  INVALIDPARAMETERVALUE_INVALIDSTRIDE = "InvalidParameterValue.InvalidStride"
//  LIMITEXCEEDED_FREQCNT = "LimitExceeded.FreqCnt"
//  LIMITEXCEEDED_IPFREQCNT = "LimitExceeded.IpFreqCnt"
//  LIMITEXCEEDED_KEYFREQCNT = "LimitExceeded.KeyFreqCnt"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
func (c *Client) DescribeRiskTrends(request *DescribeRiskTrendsRequest) (response *DescribeRiskTrendsResponse, err error) {
    return c.DescribeRiskTrendsWithContext(context.Background(), request)
}

// DescribeRiskTrends
// 以图表形式展示三种请求状态的趋势变化
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INVALIDPARAMETER_CAPSIGERROR = "InvalidParameter.CapSigError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_URLERROR = "InvalidParameter.UrlError"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_BADBODY = "InvalidParameterValue.BadBody"
//  INVALIDPARAMETERVALUE_BODYTOOLARGE = "InvalidParameterValue.BodyTooLarge"
//  INVALIDPARAMETERVALUE_CAPMISMATCH = "InvalidParameterValue.CapMisMatch"
//  INVALIDPARAMETERVALUE_HTTPMETHODERROR = "InvalidParameterValue.HttpMethodError"
//  INVALIDPARAMETERVALUE_INVALIDDATE = "InvalidParameterValue.InvalidDate"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDNUM = "InvalidParameterValue.InvalidNum"
//  INVALIDPARAMETERVALUE_INVALIDSRVID = "InvalidParameterValue.InvalidSrvId"
//  INVALIDPARAMETERVALUE_INVALIDSTRIDE = "InvalidParameterValue.InvalidStride"
//  LIMITEXCEEDED_FREQCNT = "LimitExceeded.FreqCnt"
//  LIMITEXCEEDED_IPFREQCNT = "LimitExceeded.IpFreqCnt"
//  LIMITEXCEEDED_KEYFREQCNT = "LimitExceeded.KeyFreqCnt"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
func (c *Client) DescribeRiskTrendsWithContext(ctx context.Context, request *DescribeRiskTrendsRequest) (response *DescribeRiskTrendsResponse, err error) {
    if request == nil {
        request = NewDescribeRiskTrendsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskTrends require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskTrendsResponse()
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
    return c.ManageMarketingRiskWithContext(context.Background(), request)
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
func (c *Client) ManageMarketingRiskWithContext(ctx context.Context, request *ManageMarketingRiskRequest) (response *ManageMarketingRiskResponse, err error) {
    if request == nil {
        request = NewManageMarketingRiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageMarketingRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewManageMarketingRiskResponse()
    err = c.Send(request, response)
    return
}
