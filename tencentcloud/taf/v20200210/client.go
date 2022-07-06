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

package v20200210

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-10"

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


func NewDetectFraudKOLRequest() (request *DetectFraudKOLRequest) {
    request = &DetectFraudKOLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("taf", APIVersion, "DetectFraudKOL")
    
    
    return
}

func NewDetectFraudKOLResponse() (response *DetectFraudKOLResponse) {
    response = &DetectFraudKOLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectFraudKOL
// 流量反欺诈-KOL欺诈识别
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) DetectFraudKOL(request *DetectFraudKOLRequest) (response *DetectFraudKOLResponse, err error) {
    return c.DetectFraudKOLWithContext(context.Background(), request)
}

// DetectFraudKOL
// 流量反欺诈-KOL欺诈识别
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) DetectFraudKOLWithContext(ctx context.Context, request *DetectFraudKOLRequest) (response *DetectFraudKOLResponse, err error) {
    if request == nil {
        request = NewDetectFraudKOLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectFraudKOL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectFraudKOLResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeCustomizedAudienceRequest() (request *RecognizeCustomizedAudienceRequest) {
    request = &RecognizeCustomizedAudienceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("taf", APIVersion, "RecognizeCustomizedAudience")
    
    
    return
}

func NewRecognizeCustomizedAudienceResponse() (response *RecognizeCustomizedAudienceResponse) {
    response = &RecognizeCustomizedAudienceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeCustomizedAudience
// 流量反欺诈-流量验准定制版
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) RecognizeCustomizedAudience(request *RecognizeCustomizedAudienceRequest) (response *RecognizeCustomizedAudienceResponse, err error) {
    return c.RecognizeCustomizedAudienceWithContext(context.Background(), request)
}

// RecognizeCustomizedAudience
// 流量反欺诈-流量验准定制版
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) RecognizeCustomizedAudienceWithContext(ctx context.Context, request *RecognizeCustomizedAudienceRequest) (response *RecognizeCustomizedAudienceResponse, err error) {
    if request == nil {
        request = NewRecognizeCustomizedAudienceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeCustomizedAudience require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeCustomizedAudienceResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeEffectiveFlowRequest() (request *RecognizeEffectiveFlowRequest) {
    request = &RecognizeEffectiveFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("taf", APIVersion, "RecognizeEffectiveFlow")
    
    
    return
}

func NewRecognizeEffectiveFlowResponse() (response *RecognizeEffectiveFlowResponse) {
    response = &RecognizeEffectiveFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeEffectiveFlow
// 该服务已不再对外提供能力
//
// 
//
// 筛选敏感易骚扰人群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FREQCNT = "LimitExceeded.FreqCnt"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RecognizeEffectiveFlow(request *RecognizeEffectiveFlowRequest) (response *RecognizeEffectiveFlowResponse, err error) {
    return c.RecognizeEffectiveFlowWithContext(context.Background(), request)
}

// RecognizeEffectiveFlow
// 该服务已不再对外提供能力
//
// 
//
// 筛选敏感易骚扰人群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FREQCNT = "LimitExceeded.FreqCnt"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RecognizeEffectiveFlowWithContext(ctx context.Context, request *RecognizeEffectiveFlowRequest) (response *RecognizeEffectiveFlowResponse, err error) {
    if request == nil {
        request = NewRecognizeEffectiveFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeEffectiveFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeEffectiveFlowResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizePreciseTargetAudienceRequest() (request *RecognizePreciseTargetAudienceRequest) {
    request = &RecognizePreciseTargetAudienceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("taf", APIVersion, "RecognizePreciseTargetAudience")
    
    
    return
}

func NewRecognizePreciseTargetAudienceResponse() (response *RecognizePreciseTargetAudienceResponse) {
    response = &RecognizePreciseTargetAudienceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizePreciseTargetAudience
// 流量反欺诈-流量验准高级版
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) RecognizePreciseTargetAudience(request *RecognizePreciseTargetAudienceRequest) (response *RecognizePreciseTargetAudienceResponse, err error) {
    return c.RecognizePreciseTargetAudienceWithContext(context.Background(), request)
}

// RecognizePreciseTargetAudience
// 流量反欺诈-流量验准高级版
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) RecognizePreciseTargetAudienceWithContext(ctx context.Context, request *RecognizePreciseTargetAudienceRequest) (response *RecognizePreciseTargetAudienceResponse, err error) {
    if request == nil {
        request = NewRecognizePreciseTargetAudienceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizePreciseTargetAudience require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizePreciseTargetAudienceResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeTargetAudienceRequest() (request *RecognizeTargetAudienceRequest) {
    request = &RecognizeTargetAudienceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("taf", APIVersion, "RecognizeTargetAudience")
    
    
    return
}

func NewRecognizeTargetAudienceResponse() (response *RecognizeTargetAudienceResponse) {
    response = &RecognizeTargetAudienceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeTargetAudience
// 流量反欺诈-流量验准
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) RecognizeTargetAudience(request *RecognizeTargetAudienceRequest) (response *RecognizeTargetAudienceResponse, err error) {
    return c.RecognizeTargetAudienceWithContext(context.Background(), request)
}

// RecognizeTargetAudience
// 流量反欺诈-流量验准
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) RecognizeTargetAudienceWithContext(ctx context.Context, request *RecognizeTargetAudienceRequest) (response *RecognizeTargetAudienceResponse, err error) {
    if request == nil {
        request = NewRecognizeTargetAudienceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeTargetAudience require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeTargetAudienceResponse()
    err = c.Send(request, response)
    return
}

func NewSendTrafficSecuritySmsMessageRequest() (request *SendTrafficSecuritySmsMessageRequest) {
    request = &SendTrafficSecuritySmsMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("taf", APIVersion, "SendTrafficSecuritySmsMessage")
    
    
    return
}

func NewSendTrafficSecuritySmsMessageResponse() (response *SendTrafficSecuritySmsMessageResponse) {
    response = &SendTrafficSecuritySmsMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendTrafficSecuritySmsMessage
// 流量安选产品，短信发送接口
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) SendTrafficSecuritySmsMessage(request *SendTrafficSecuritySmsMessageRequest) (response *SendTrafficSecuritySmsMessageResponse, err error) {
    return c.SendTrafficSecuritySmsMessageWithContext(context.Background(), request)
}

// SendTrafficSecuritySmsMessage
// 流量安选产品，短信发送接口
//
// 可能返回的错误码:
//  AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"
//  AUTHFAILURE_EXPIRED = "AuthFailure.Expired"
//  INTERNALERROR_ADDTASKFAIL = "InternalError.AddTaskFail"
//  INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"
//  INTERNALERROR_CONNECTDBTIMEOUT = "InternalError.ConnectDBTimeout"
//  INTERNALERROR_DETECTFAIL = "InternalError.DetectFail"
//  INTERNALERROR_DOWNLOADFAIL = "InternalError.DownloadFail"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_QUERYTASKFAIL = "InternalError.QueryTaskFail"
//  INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"
//  INTERNALERROR_SIGNATUREFAIL = "InternalError.SignatureFail"
//  INTERNALERROR_TASKIDNOTFOUND = "InternalError.TaskIdNotFound"
//  INTERNALERROR_UPDATETASKFAIL = "InternalError.UpdateTaskFail"
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
func (c *Client) SendTrafficSecuritySmsMessageWithContext(ctx context.Context, request *SendTrafficSecuritySmsMessageRequest) (response *SendTrafficSecuritySmsMessageResponse, err error) {
    if request == nil {
        request = NewSendTrafficSecuritySmsMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendTrafficSecuritySmsMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendTrafficSecuritySmsMessageResponse()
    err = c.Send(request, response)
    return
}
