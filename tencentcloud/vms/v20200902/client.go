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

package v20200902

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-09-02"

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


func NewSendCodeVoiceRequest() (request *SendCodeVoiceRequest) {
    request = &SendCodeVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vms", APIVersion, "SendCodeVoice")
    
    
    return
}

func NewSendCodeVoiceResponse() (response *SendCodeVoiceResponse) {
    response = &SendCodeVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendCodeVoice
// 给用户发语音验证码（仅支持数字）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSUPSTREAMTIMEOUT = "FailedOperation.AccessUpstreamTimeout"
//  FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_INSUFFICIENTBALANCEINVOICEPACKAGE = "FailedOperation.InsufficientBalanceInVoicePackage"
//  FAILEDOPERATION_INVALIDJSONPARAMETERS = "FailedOperation.InvalidJsonParameters"
//  FAILEDOPERATION_INVALIDPARAMETERS = "FailedOperation.InvalidParameters"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PARAMETERSOTHERERROR = "FailedOperation.ParametersOtherError"
//  FAILEDOPERATION_PHONENUMBERUNAPPLIEDOREXPIRED = "FailedOperation.PhonenumberUnappliedOrExpired"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_ACCESSUPSTREAMTIMEOUT = "InternalError.AccessUpstreamTimeout"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_SSOSENDRECVFAIL = "InternalError.SsoSendRecvFail"
//  INTERNALERROR_UPSTREAMERROR = "InternalError.UpstreamError"
//  INVALIDPARAMETERVALUE_CALLEDNUMBERVERIFYFAIL = "InvalidParameterValue.CalledNumberVerifyFail"
//  INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_VOICESDKAPPIDVERIFYFAIL = "UnauthorizedOperation.VoiceSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendCodeVoice(request *SendCodeVoiceRequest) (response *SendCodeVoiceResponse, err error) {
    return c.SendCodeVoiceWithContext(context.Background(), request)
}

// SendCodeVoice
// 给用户发语音验证码（仅支持数字）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSUPSTREAMTIMEOUT = "FailedOperation.AccessUpstreamTimeout"
//  FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_INSUFFICIENTBALANCEINVOICEPACKAGE = "FailedOperation.InsufficientBalanceInVoicePackage"
//  FAILEDOPERATION_INVALIDJSONPARAMETERS = "FailedOperation.InvalidJsonParameters"
//  FAILEDOPERATION_INVALIDPARAMETERS = "FailedOperation.InvalidParameters"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PARAMETERSOTHERERROR = "FailedOperation.ParametersOtherError"
//  FAILEDOPERATION_PHONENUMBERUNAPPLIEDOREXPIRED = "FailedOperation.PhonenumberUnappliedOrExpired"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_ACCESSUPSTREAMTIMEOUT = "InternalError.AccessUpstreamTimeout"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_SSOSENDRECVFAIL = "InternalError.SsoSendRecvFail"
//  INTERNALERROR_UPSTREAMERROR = "InternalError.UpstreamError"
//  INVALIDPARAMETERVALUE_CALLEDNUMBERVERIFYFAIL = "InvalidParameterValue.CalledNumberVerifyFail"
//  INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_VOICESDKAPPIDVERIFYFAIL = "UnauthorizedOperation.VoiceSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendCodeVoiceWithContext(ctx context.Context, request *SendCodeVoiceRequest) (response *SendCodeVoiceResponse, err error) {
    if request == nil {
        request = NewSendCodeVoiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendCodeVoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendCodeVoiceResponse()
    err = c.Send(request, response)
    return
}

func NewSendTtsVoiceRequest() (request *SendTtsVoiceRequest) {
    request = &SendTtsVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vms", APIVersion, "SendTtsVoice")
    
    
    return
}

func NewSendTtsVoiceResponse() (response *SendTtsVoiceResponse) {
    response = &SendTtsVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendTtsVoice
// 给用户发送指定模板的语音通知。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSUPSTREAMTIMEOUT = "FailedOperation.AccessUpstreamTimeout"
//  FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_INSUFFICIENTBALANCEINVOICEPACKAGE = "FailedOperation.InsufficientBalanceInVoicePackage"
//  FAILEDOPERATION_INVALIDJSONPARAMETERS = "FailedOperation.InvalidJsonParameters"
//  FAILEDOPERATION_INVALIDPARAMETERS = "FailedOperation.InvalidParameters"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PARAMETERSOTHERERROR = "FailedOperation.ParametersOtherError"
//  FAILEDOPERATION_PHONENUMBERUNAPPLIEDOREXPIRED = "FailedOperation.PhonenumberUnappliedOrExpired"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_ACCESSUPSTREAMTIMEOUT = "InternalError.AccessUpstreamTimeout"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_SSOSENDRECVFAIL = "InternalError.SsoSendRecvFail"
//  INTERNALERROR_UPSTREAMERROR = "InternalError.UpstreamError"
//  INVALIDPARAMETERVALUE_CALLEDNUMBERVERIFYFAIL = "InvalidParameterValue.CalledNumberVerifyFail"
//  INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_VOICESDKAPPIDVERIFYFAIL = "UnauthorizedOperation.VoiceSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendTtsVoice(request *SendTtsVoiceRequest) (response *SendTtsVoiceResponse, err error) {
    return c.SendTtsVoiceWithContext(context.Background(), request)
}

// SendTtsVoice
// 给用户发送指定模板的语音通知。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSUPSTREAMTIMEOUT = "FailedOperation.AccessUpstreamTimeout"
//  FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_INSUFFICIENTBALANCEINVOICEPACKAGE = "FailedOperation.InsufficientBalanceInVoicePackage"
//  FAILEDOPERATION_INVALIDJSONPARAMETERS = "FailedOperation.InvalidJsonParameters"
//  FAILEDOPERATION_INVALIDPARAMETERS = "FailedOperation.InvalidParameters"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PARAMETERSOTHERERROR = "FailedOperation.ParametersOtherError"
//  FAILEDOPERATION_PHONENUMBERUNAPPLIEDOREXPIRED = "FailedOperation.PhonenumberUnappliedOrExpired"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_ACCESSUPSTREAMTIMEOUT = "InternalError.AccessUpstreamTimeout"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_SSOSENDRECVFAIL = "InternalError.SsoSendRecvFail"
//  INTERNALERROR_UPSTREAMERROR = "InternalError.UpstreamError"
//  INVALIDPARAMETERVALUE_CALLEDNUMBERVERIFYFAIL = "InvalidParameterValue.CalledNumberVerifyFail"
//  INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_VOICESDKAPPIDVERIFYFAIL = "UnauthorizedOperation.VoiceSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendTtsVoiceWithContext(ctx context.Context, request *SendTtsVoiceRequest) (response *SendTtsVoiceResponse, err error) {
    if request == nil {
        request = NewSendTtsVoiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendTtsVoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendTtsVoiceResponse()
    err = c.Send(request, response)
    return
}
