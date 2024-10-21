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

package v20210111

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-01-11"

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


func NewAddSmsSignRequest() (request *AddSmsSignRequest) {
    request = &AddSmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "AddSmsSign")
    
    
    return
}

func NewAddSmsSignResponse() (response *AddSmsSignResponse) {
    response = &AddSmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddSmsSign
// 本接口 (AddSmsSign) 用于添加短信签名。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>添加短信签名前，请先认真参阅 <a href="https://cloud.tencent.com/document/product/382/39022">腾讯云短信签名审核标准。</a></li><li>个人认证用户不支持使用 API 申请短信签名，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 申请短信签名。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_PARAMETERSOTHERERROR = "FailedOperation.ParametersOtherError"
//  FAILEDOPERATION_SIGNNUMBERLIMIT = "FailedOperation.SignNumberLimit"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDUSEDMETHOD = "InvalidParameterValue.InvalidUsedMethod"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  INVALIDPARAMETERVALUE_SIGNALREADYPASSEDCHECK = "InvalidParameterValue.SignAlreadyPassedCheck"
//  INVALIDPARAMETERVALUE_SIGNEXISTANDUNAPPROVED = "InvalidParameterValue.SignExistAndUnapproved"
//  INVALIDPARAMETERVALUE_SIGNNAMELENGTHTOOLONG = "InvalidParameterValue.SignNameLengthTooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddSmsSign(request *AddSmsSignRequest) (response *AddSmsSignResponse, err error) {
    return c.AddSmsSignWithContext(context.Background(), request)
}

// AddSmsSign
// 本接口 (AddSmsSign) 用于添加短信签名。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>添加短信签名前，请先认真参阅 <a href="https://cloud.tencent.com/document/product/382/39022">腾讯云短信签名审核标准。</a></li><li>个人认证用户不支持使用 API 申请短信签名，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 申请短信签名。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_PARAMETERSOTHERERROR = "FailedOperation.ParametersOtherError"
//  FAILEDOPERATION_SIGNNUMBERLIMIT = "FailedOperation.SignNumberLimit"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDUSEDMETHOD = "InvalidParameterValue.InvalidUsedMethod"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  INVALIDPARAMETERVALUE_SIGNALREADYPASSEDCHECK = "InvalidParameterValue.SignAlreadyPassedCheck"
//  INVALIDPARAMETERVALUE_SIGNEXISTANDUNAPPROVED = "InvalidParameterValue.SignExistAndUnapproved"
//  INVALIDPARAMETERVALUE_SIGNNAMELENGTHTOOLONG = "InvalidParameterValue.SignNameLengthTooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddSmsSignWithContext(ctx context.Context, request *AddSmsSignRequest) (response *AddSmsSignResponse, err error) {
    if request == nil {
        request = NewAddSmsSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSmsSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddSmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewAddSmsTemplateRequest() (request *AddSmsTemplateRequest) {
    request = &AddSmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "AddSmsTemplate")
    
    
    return
}

func NewAddSmsTemplateResponse() (response *AddSmsTemplateResponse) {
    response = &AddSmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddSmsTemplate
// 本接口 (AddSmsTemplate) 用于创建短信模板。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>申请短信模板前，请先认真参阅 <a href="https://cloud.tencent.com/document/product/382/39023">腾讯云短信正文模板审核标准。</a></li><li>个人认证用户不支持使用 API 申请短信正文模板，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 申请短信正文模板。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURE = "FailedOperation.MissingSignature"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"
//  INTERNALERROR_JSONPARSEFAIL = "InternalError.JsonParseFail"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_DIRTYWORDFOUND = "InvalidParameter.DirtyWordFound"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDSMSTYPE = "InvalidParameterValue.InvalidSmsType"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"
//  INVALIDPARAMETERVALUE_MARKETINGTEMPLATEWITHOUTUNSUBSCRIBE = "InvalidParameterValue.MarketingTemplateWithoutUnsubscribe"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTEMPLATEVARIABLE = "InvalidParameterValue.UnsupportedTemplateVariable"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddSmsTemplate(request *AddSmsTemplateRequest) (response *AddSmsTemplateResponse, err error) {
    return c.AddSmsTemplateWithContext(context.Background(), request)
}

// AddSmsTemplate
// 本接口 (AddSmsTemplate) 用于创建短信模板。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>申请短信模板前，请先认真参阅 <a href="https://cloud.tencent.com/document/product/382/39023">腾讯云短信正文模板审核标准。</a></li><li>个人认证用户不支持使用 API 申请短信正文模板，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 申请短信正文模板。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURE = "FailedOperation.MissingSignature"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"
//  INTERNALERROR_JSONPARSEFAIL = "InternalError.JsonParseFail"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_DIRTYWORDFOUND = "InvalidParameter.DirtyWordFound"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDSMSTYPE = "InvalidParameterValue.InvalidSmsType"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"
//  INVALIDPARAMETERVALUE_MARKETINGTEMPLATEWITHOUTUNSUBSCRIBE = "InvalidParameterValue.MarketingTemplateWithoutUnsubscribe"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTEMPLATEVARIABLE = "InvalidParameterValue.UnsupportedTemplateVariable"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddSmsTemplateWithContext(ctx context.Context, request *AddSmsTemplateRequest) (response *AddSmsTemplateResponse, err error) {
    if request == nil {
        request = NewAddSmsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSmsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddSmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCallbackStatusStatisticsRequest() (request *CallbackStatusStatisticsRequest) {
    request = &CallbackStatusStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "CallbackStatusStatistics")
    
    
    return
}

func NewCallbackStatusStatisticsResponse() (response *CallbackStatusStatisticsResponse) {
    response = &CallbackStatusStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CallbackStatusStatistics
// 本接口 (CallbackStatusStatistics) 用于统计用户回执的数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_OFFSETVERIFYFAIL = "InvalidParameterValue.OffsetVerifyFail"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CallbackStatusStatistics(request *CallbackStatusStatisticsRequest) (response *CallbackStatusStatisticsResponse, err error) {
    return c.CallbackStatusStatisticsWithContext(context.Background(), request)
}

// CallbackStatusStatistics
// 本接口 (CallbackStatusStatistics) 用于统计用户回执的数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_OFFSETVERIFYFAIL = "InvalidParameterValue.OffsetVerifyFail"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CallbackStatusStatisticsWithContext(ctx context.Context, request *CallbackStatusStatisticsRequest) (response *CallbackStatusStatisticsResponse, err error) {
    if request == nil {
        request = NewCallbackStatusStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CallbackStatusStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewCallbackStatusStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmsSignRequest() (request *DeleteSmsSignRequest) {
    request = &DeleteSmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "DeleteSmsSign")
    
    
    return
}

func NewDeleteSmsSignResponse() (response *DeleteSmsSignResponse) {
    response = &DeleteSmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSmsSign
// 本接口 (DeleteSmsSign) 用于删除短信签名。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>个人认证用户不支持使用 API 删除短信签名，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 删除短信签名。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_SIGNIDNOTEXIST = "FailedOperation.SignIdNotExist"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSmsSign(request *DeleteSmsSignRequest) (response *DeleteSmsSignResponse, err error) {
    return c.DeleteSmsSignWithContext(context.Background(), request)
}

// DeleteSmsSign
// 本接口 (DeleteSmsSign) 用于删除短信签名。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>个人认证用户不支持使用 API 删除短信签名，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 删除短信签名。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_SIGNIDNOTEXIST = "FailedOperation.SignIdNotExist"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSmsSignWithContext(ctx context.Context, request *DeleteSmsSignRequest) (response *DeleteSmsSignResponse, err error) {
    if request == nil {
        request = NewDeleteSmsSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSmsSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmsTemplateRequest() (request *DeleteSmsTemplateRequest) {
    request = &DeleteSmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "DeleteSmsTemplate")
    
    
    return
}

func NewDeleteSmsTemplateResponse() (response *DeleteSmsTemplateResponse) {
    response = &DeleteSmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSmsTemplate
// 本接口 (DeleteSmsTemplate) 用于删除短信模板。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>个人认证用户不支持使用 API 删除短信正文模板，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 删除短信正文模板。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSmsTemplate(request *DeleteSmsTemplateRequest) (response *DeleteSmsTemplateResponse, err error) {
    return c.DeleteSmsTemplateWithContext(context.Background(), request)
}

// DeleteSmsTemplate
// 本接口 (DeleteSmsTemplate) 用于删除短信模板。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>个人认证用户不支持使用 API 删除短信正文模板，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 删除短信正文模板。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSmsTemplateWithContext(ctx context.Context, request *DeleteSmsTemplateRequest) (response *DeleteSmsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSmsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSmsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePhoneNumberInfoRequest() (request *DescribePhoneNumberInfoRequest) {
    request = &DescribePhoneNumberInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "DescribePhoneNumberInfo")
    
    
    return
}

func NewDescribePhoneNumberInfoResponse() (response *DescribePhoneNumberInfoResponse) {
    response = &DescribePhoneNumberInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePhoneNumberInfo
// 本接口 (DescribePhoneNumberInfo) 用于提供电话号码的信息查询，包括国家（或地区）码、规范的 E.164 格式号码等。
//
// - 例如：查询号码 +86018501234444，可以得到国家码 86、规范的 E.164 号码 +8618501234444 等信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERPARSEFAIL = "FailedOperation.PhoneNumberParseFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  LIMITEXCEEDED_PHONENUMBERCOUNTLIMIT = "LimitExceeded.PhoneNumberCountLimit"
func (c *Client) DescribePhoneNumberInfo(request *DescribePhoneNumberInfoRequest) (response *DescribePhoneNumberInfoResponse, err error) {
    return c.DescribePhoneNumberInfoWithContext(context.Background(), request)
}

// DescribePhoneNumberInfo
// 本接口 (DescribePhoneNumberInfo) 用于提供电话号码的信息查询，包括国家（或地区）码、规范的 E.164 格式号码等。
//
// - 例如：查询号码 +86018501234444，可以得到国家码 86、规范的 E.164 号码 +8618501234444 等信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERPARSEFAIL = "FailedOperation.PhoneNumberParseFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  LIMITEXCEEDED_PHONENUMBERCOUNTLIMIT = "LimitExceeded.PhoneNumberCountLimit"
func (c *Client) DescribePhoneNumberInfoWithContext(ctx context.Context, request *DescribePhoneNumberInfoRequest) (response *DescribePhoneNumberInfoResponse, err error) {
    if request == nil {
        request = NewDescribePhoneNumberInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePhoneNumberInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePhoneNumberInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsSignListRequest() (request *DescribeSmsSignListRequest) {
    request = &DescribeSmsSignListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "DescribeSmsSignList")
    
    
    return
}

func NewDescribeSmsSignListResponse() (response *DescribeSmsSignListResponse) {
    response = &DescribeSmsSignListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSmsSignList
// 本接口 (DescribeSmsSignList) 用于查询短信签名状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>个人认证用户不支持使用 API 查询短信签名，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 查询短信签名。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_JSONPARSEFAIL = "InternalError.JsonParseFail"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSmsSignList(request *DescribeSmsSignListRequest) (response *DescribeSmsSignListResponse, err error) {
    return c.DescribeSmsSignListWithContext(context.Background(), request)
}

// DescribeSmsSignList
// 本接口 (DescribeSmsSignList) 用于查询短信签名状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>个人认证用户不支持使用 API 查询短信签名，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 查询短信签名。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_JSONPARSEFAIL = "InternalError.JsonParseFail"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSmsSignListWithContext(ctx context.Context, request *DescribeSmsSignListRequest) (response *DescribeSmsSignListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsSignListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmsSignList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmsSignListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsTemplateListRequest() (request *DescribeSmsTemplateListRequest) {
    request = &DescribeSmsTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "DescribeSmsTemplateList")
    
    
    return
}

func NewDescribeSmsTemplateListResponse() (response *DescribeSmsTemplateListResponse) {
    response = &DescribeSmsTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSmsTemplateList
// 本接口 (DescribeSmsTemplateList) 用于查询短信模板状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>个人认证用户不支持使用 API 查询短信正文模板，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 查询短信正文模板。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_PROHIBITSUBACCOUNTUSE = "FailedOperation.ProhibitSubAccountUse"
//  FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_TEMPLATEWITHDIRTYWORDS = "InvalidParameterValue.TemplateWithDirtyWords"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSmsTemplateList(request *DescribeSmsTemplateListRequest) (response *DescribeSmsTemplateListResponse, err error) {
    return c.DescribeSmsTemplateListWithContext(context.Background(), request)
}

// DescribeSmsTemplateList
// 本接口 (DescribeSmsTemplateList) 用于查询短信模板状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>个人认证用户不支持使用 API 查询短信正文模板，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 查询短信正文模板。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_PROHIBITSUBACCOUNTUSE = "FailedOperation.ProhibitSubAccountUse"
//  FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_TEMPLATEWITHDIRTYWORDS = "InvalidParameterValue.TemplateWithDirtyWords"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSmsTemplateListWithContext(ctx context.Context, request *DescribeSmsTemplateListRequest) (response *DescribeSmsTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsTemplateListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmsTemplateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmsTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmsSignRequest() (request *ModifySmsSignRequest) {
    request = &ModifySmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "ModifySmsSign")
    
    
    return
}

func NewModifySmsSignResponse() (response *ModifySmsSignResponse) {
    response = &ModifySmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySmsSign
// 本接口 (ModifySmsSign) 用于修改短信签名。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>修改短信签名前，请先认真参阅 <a href="https://cloud.tencent.com/document/product/382/39022">腾讯云短信签名审核标准。</a></li><li>个人认证用户不支持使用 API 修改短信签名，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 修改短信签名。</li><li>修改短信签名，仅当签名为<b>待审核</b>或<b>已拒绝</b>状态时，才能进行修改，<b>已审核通过</b>的签名不支持修改。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURELIST = "FailedOperation.MissingSignatureList"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_SIGNNUMBERLIMIT = "FailedOperation.SignNumberLimit"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDSIGNPURPOSE = "InvalidParameterValue.InvalidSignPurpose"
//  INVALIDPARAMETERVALUE_SIGNALREADYPASSEDCHECK = "InvalidParameterValue.SignAlreadyPassedCheck"
//  INVALIDPARAMETERVALUE_SIGNEXISTANDUNAPPROVED = "InvalidParameterValue.SignExistAndUnapproved"
//  INVALIDPARAMETERVALUE_SIGNNAMELENGTHTOOLONG = "InvalidParameterValue.SignNameLengthTooLong"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySmsSign(request *ModifySmsSignRequest) (response *ModifySmsSignResponse, err error) {
    return c.ModifySmsSignWithContext(context.Background(), request)
}

// ModifySmsSign
// 本接口 (ModifySmsSign) 用于修改短信签名。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>修改短信签名前，请先认真参阅 <a href="https://cloud.tencent.com/document/product/382/39022">腾讯云短信签名审核标准。</a></li><li>个人认证用户不支持使用 API 修改短信签名，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 修改短信签名。</li><li>修改短信签名，仅当签名为<b>待审核</b>或<b>已拒绝</b>状态时，才能进行修改，<b>已审核通过</b>的签名不支持修改。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURELIST = "FailedOperation.MissingSignatureList"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_SIGNNUMBERLIMIT = "FailedOperation.SignNumberLimit"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDSIGNPURPOSE = "InvalidParameterValue.InvalidSignPurpose"
//  INVALIDPARAMETERVALUE_SIGNALREADYPASSEDCHECK = "InvalidParameterValue.SignAlreadyPassedCheck"
//  INVALIDPARAMETERVALUE_SIGNEXISTANDUNAPPROVED = "InvalidParameterValue.SignExistAndUnapproved"
//  INVALIDPARAMETERVALUE_SIGNNAMELENGTHTOOLONG = "InvalidParameterValue.SignNameLengthTooLong"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySmsSignWithContext(ctx context.Context, request *ModifySmsSignRequest) (response *ModifySmsSignResponse, err error) {
    if request == nil {
        request = NewModifySmsSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySmsSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmsTemplateRequest() (request *ModifySmsTemplateRequest) {
    request = &ModifySmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "ModifySmsTemplate")
    
    
    return
}

func NewModifySmsTemplateResponse() (response *ModifySmsTemplateResponse) {
    response = &ModifySmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySmsTemplate
// 本接口 (ModifySmsTemplate) 用于修改短信模板。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>修改短信正文模板前，请先认真参阅 <a href="https://cloud.tencent.com/document/product/382/39023">腾讯云短信正文模板审核标准。</a></li><li>个人认证用户不支持使用 API 修改短信正文模板，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 修改短信正文模板。</li><li>修改短信模板，仅当正文模板为<b>待审核</b>或<b>已拒绝</b>状态时，才能进行修改，<b>已审核通过</b>的正文模板不支持修改。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"
//  FAILEDOPERATION_MISSINGTEMPLATELIST = "FailedOperation.MissingTemplateList"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_DIRTYWORDFOUND = "InvalidParameter.DirtyWordFound"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDSMSTYPE = "InvalidParameterValue.InvalidSmsType"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"
//  INVALIDPARAMETERVALUE_MARKETINGTEMPLATEWITHOUTUNSUBSCRIBE = "InvalidParameterValue.MarketingTemplateWithoutUnsubscribe"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTEMPLATEVARIABLE = "InvalidParameterValue.UnsupportedTemplateVariable"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySmsTemplate(request *ModifySmsTemplateRequest) (response *ModifySmsTemplateResponse, err error) {
    return c.ModifySmsTemplateWithContext(context.Background(), request)
}

// ModifySmsTemplate
// 本接口 (ModifySmsTemplate) 用于修改短信模板。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>修改短信正文模板前，请先认真参阅 <a href="https://cloud.tencent.com/document/product/382/39023">腾讯云短信正文模板审核标准。</a></li><li>个人认证用户不支持使用 API 修改短信正文模板，请参阅了解 <a href="https://cloud.tencent.com/document/product/378/3629">实名认证基本介绍</a>，如果为个人认证请登录 <a href="https://console.cloud.tencent.com/smsv2">控制台</a> 修改短信正文模板。</li><li>修改短信模板，仅当正文模板为<b>待审核</b>或<b>已拒绝</b>状态时，才能进行修改，<b>已审核通过</b>的正文模板不支持修改。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"
//  FAILEDOPERATION_MISSINGTEMPLATELIST = "FailedOperation.MissingTemplateList"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_DIRTYWORDFOUND = "InvalidParameter.DirtyWordFound"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDSMSTYPE = "InvalidParameterValue.InvalidSmsType"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"
//  INVALIDPARAMETERVALUE_MARKETINGTEMPLATEWITHOUTUNSUBSCRIBE = "InvalidParameterValue.MarketingTemplateWithoutUnsubscribe"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTEMPLATEVARIABLE = "InvalidParameterValue.UnsupportedTemplateVariable"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySmsTemplateWithContext(ctx context.Context, request *ModifySmsTemplateRequest) (response *ModifySmsTemplateResponse, err error) {
    if request == nil {
        request = NewModifySmsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySmsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsReplyStatusRequest() (request *PullSmsReplyStatusRequest) {
    request = &PullSmsReplyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsReplyStatus")
    
    
    return
}

func NewPullSmsReplyStatusResponse() (response *PullSmsReplyStatusResponse) {
    response = &PullSmsReplyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullSmsReplyStatus
// 本接口 (PullSmsReplyStatus) 用于拉取短信回复状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>此接口需要联系  <a href="https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81">腾讯云短信小助手</a> 开通。</li><li>上行回复也支持 <a href="https://cloud.tencent.com/document/product/382/42907">配置回复回调</a> 的方式获取。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PullSmsReplyStatus(request *PullSmsReplyStatusRequest) (response *PullSmsReplyStatusResponse, err error) {
    return c.PullSmsReplyStatusWithContext(context.Background(), request)
}

// PullSmsReplyStatus
// 本接口 (PullSmsReplyStatus) 用于拉取短信回复状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>此接口需要联系  <a href="https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81">腾讯云短信小助手</a> 开通。</li><li>上行回复也支持 <a href="https://cloud.tencent.com/document/product/382/42907">配置回复回调</a> 的方式获取。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PullSmsReplyStatusWithContext(ctx context.Context, request *PullSmsReplyStatusRequest) (response *PullSmsReplyStatusResponse, err error) {
    if request == nil {
        request = NewPullSmsReplyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullSmsReplyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullSmsReplyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsReplyStatusByPhoneNumberRequest() (request *PullSmsReplyStatusByPhoneNumberRequest) {
    request = &PullSmsReplyStatusByPhoneNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsReplyStatusByPhoneNumber")
    
    
    return
}

func NewPullSmsReplyStatusByPhoneNumberResponse() (response *PullSmsReplyStatusByPhoneNumberResponse) {
    response = &PullSmsReplyStatusByPhoneNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullSmsReplyStatusByPhoneNumber
// 本接口 (PullSmsReplyStatusByPhoneNumber) 用于拉取单个号码短信回复状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>上行回复也支持 <a href="https://cloud.tencent.com/document/product/382/42907">配置回复回调</a> 的方式获取。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PullSmsReplyStatusByPhoneNumber(request *PullSmsReplyStatusByPhoneNumberRequest) (response *PullSmsReplyStatusByPhoneNumberResponse, err error) {
    return c.PullSmsReplyStatusByPhoneNumberWithContext(context.Background(), request)
}

// PullSmsReplyStatusByPhoneNumber
// 本接口 (PullSmsReplyStatusByPhoneNumber) 用于拉取单个号码短信回复状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>上行回复也支持 <a href="https://cloud.tencent.com/document/product/382/42907">配置回复回调</a> 的方式获取。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PullSmsReplyStatusByPhoneNumberWithContext(ctx context.Context, request *PullSmsReplyStatusByPhoneNumberRequest) (response *PullSmsReplyStatusByPhoneNumberResponse, err error) {
    if request == nil {
        request = NewPullSmsReplyStatusByPhoneNumberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullSmsReplyStatusByPhoneNumber require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullSmsReplyStatusByPhoneNumberResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsSendStatusRequest() (request *PullSmsSendStatusRequest) {
    request = &PullSmsSendStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsSendStatus")
    
    
    return
}

func NewPullSmsSendStatusResponse() (response *PullSmsSendStatusResponse) {
    response = &PullSmsSendStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullSmsSendStatus
// 本接口 (PullSmsSendStatus) 用于拉取短信下发状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>此接口需要联系  <a href="https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81">腾讯云短信小助手</a> 开通。</li><li>下发状态也支持 <a href="https://cloud.tencent.com/document/product/382/37809#.E7.9F.AD.E4.BF.A1.E7.8A.B6.E6.80.81.E5.9B.9E.E8.B0.83.E9.85.8D.E7.BD.AE">配置回调</a> 的方式获取。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PullSmsSendStatus(request *PullSmsSendStatusRequest) (response *PullSmsSendStatusResponse, err error) {
    return c.PullSmsSendStatusWithContext(context.Background(), request)
}

// PullSmsSendStatus
// 本接口 (PullSmsSendStatus) 用于拉取短信下发状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>此接口需要联系  <a href="https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81">腾讯云短信小助手</a> 开通。</li><li>下发状态也支持 <a href="https://cloud.tencent.com/document/product/382/37809#.E7.9F.AD.E4.BF.A1.E7.8A.B6.E6.80.81.E5.9B.9E.E8.B0.83.E9.85.8D.E7.BD.AE">配置回调</a> 的方式获取。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PullSmsSendStatusWithContext(ctx context.Context, request *PullSmsSendStatusRequest) (response *PullSmsSendStatusResponse, err error) {
    if request == nil {
        request = NewPullSmsSendStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullSmsSendStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullSmsSendStatusResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsSendStatusByPhoneNumberRequest() (request *PullSmsSendStatusByPhoneNumberRequest) {
    request = &PullSmsSendStatusByPhoneNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsSendStatusByPhoneNumber")
    
    
    return
}

func NewPullSmsSendStatusByPhoneNumberResponse() (response *PullSmsSendStatusByPhoneNumberResponse) {
    response = &PullSmsSendStatusByPhoneNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullSmsSendStatusByPhoneNumber
// 本接口 (PullSmsSendStatusByPhoneNumber) 用于拉取单个号码短信下发状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>下发状态也支持 <a href="https://cloud.tencent.com/document/product/382/37809#.E7.9F.AD.E4.BF.A1.E7.8A.B6.E6.80.81.E5.9B.9E.E8.B0.83.E9.85.8D.E7.BD.AE">配置回调</a> 的方式获取。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_JSONPARSEFAIL = "InternalError.JsonParseFail"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PullSmsSendStatusByPhoneNumber(request *PullSmsSendStatusByPhoneNumberRequest) (response *PullSmsSendStatusByPhoneNumberResponse, err error) {
    return c.PullSmsSendStatusByPhoneNumberWithContext(context.Background(), request)
}

// PullSmsSendStatusByPhoneNumber
// 本接口 (PullSmsSendStatusByPhoneNumber) 用于拉取单个号码短信下发状态。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>下发状态也支持 <a href="https://cloud.tencent.com/document/product/382/37809#.E7.9F.AD.E4.BF.A1.E7.8A.B6.E6.80.81.E5.9B.9E.E8.B0.83.E9.85.8D.E7.BD.AE">配置回调</a> 的方式获取。</li></ul></blockquote>
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_JSONPARSEFAIL = "InternalError.JsonParseFail"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PullSmsSendStatusByPhoneNumberWithContext(ctx context.Context, request *PullSmsSendStatusByPhoneNumberRequest) (response *PullSmsSendStatusByPhoneNumberResponse, err error) {
    if request == nil {
        request = NewPullSmsSendStatusByPhoneNumberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullSmsSendStatusByPhoneNumber require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullSmsSendStatusByPhoneNumberResponse()
    err = c.Send(request, response)
    return
}

func NewReportConversionRequest() (request *ReportConversionRequest) {
    request = &ReportConversionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "ReportConversion")
    
    
    return
}

func NewReportConversionResponse() (response *ReportConversionResponse) {
    response = &ReportConversionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReportConversion
// 本接口 (ReportConversion) 用于短信转化率上报。将已接收到短信的流水号上报到腾讯云短信服务。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>该接口当前以白名单方式对外开放，如有需要请联系  <a href="https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81">腾讯云短信小助手</a> 开通。</li></ul></blockquote>
//
// 可能返回的错误码:
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
func (c *Client) ReportConversion(request *ReportConversionRequest) (response *ReportConversionResponse, err error) {
    return c.ReportConversionWithContext(context.Background(), request)
}

// ReportConversion
// 本接口 (ReportConversion) 用于短信转化率上报。将已接收到短信的流水号上报到腾讯云短信服务。
//
// <blockquote class="d-mod-explain"><div class="d-mod-title d-explain-title" style="line-height: normal;"><i class="d-icon-explain"></i>说明：</div><p></p><ul><li>该接口当前以白名单方式对外开放，如有需要请联系  <a href="https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81">腾讯云短信小助手</a> 开通。</li></ul></blockquote>
//
// 可能返回的错误码:
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
func (c *Client) ReportConversionWithContext(ctx context.Context, request *ReportConversionRequest) (response *ReportConversionResponse, err error) {
    if request == nil {
        request = NewReportConversionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportConversion require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportConversionResponse()
    err = c.Send(request, response)
    return
}

func NewSendSmsRequest() (request *SendSmsRequest) {
    request = &SendSmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "SendSms")
    
    
    return
}

func NewSendSmsResponse() (response *SendSmsResponse) {
    response = &SendSmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendSms
// 本接口 (SendSms) 用于发送验证码、通知类短信和营销短信。支持国内短信与国际/港澳台短信。
//
// - 当前接口属于 2021-01-11 版本，如果您仍在使用 [2019-07-11 版本](https://cloud.tencent.com/document/product/382/38778)，建议您使用当前最新版本的接口，版本差异可参考[版本描述](https://cloud.tencent.com/document/product/382/63195#.E7.89.88.E6.9C.AC.E6.8F.8F.E8.BF.B0)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_INSUFFICIENTBALANCEINSMSPACKAGE = "FailedOperation.InsufficientBalanceInSmsPackage"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MARKETINGSENDTIMECONSTRAINT = "FailedOperation.MarketingSendTimeConstraint"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  FAILEDOPERATION_TEMPLATEPARAMSETNOTMATCHAPPROVEDTEMPLATE = "FailedOperation.TemplateParamSetNotMatchApprovedTemplate"
//  FAILEDOPERATION_TEMPLATEUNAPPROVEDORNOTEXIST = "FailedOperation.TemplateUnapprovedOrNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_PROHIBITEDUSEURLINTEMPLATEPARAMETER = "InvalidParameterValue.ProhibitedUseUrlInTemplateParameter"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERFORMATERROR = "InvalidParameterValue.TemplateParameterFormatError"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERLENGTHLIMIT = "InvalidParameterValue.TemplateParameterLengthLimit"
//  LIMITEXCEEDED_APPCOUNTRYORREGIONDAILYLIMIT = "LimitExceeded.AppCountryOrRegionDailyLimit"
//  LIMITEXCEEDED_APPCOUNTRYORREGIONINBLACKLIST = "LimitExceeded.AppCountryOrRegionInBlacklist"
//  LIMITEXCEEDED_APPDAILYLIMIT = "LimitExceeded.AppDailyLimit"
//  LIMITEXCEEDED_APPGLOBALDAILYLIMIT = "LimitExceeded.AppGlobalDailyLimit"
//  LIMITEXCEEDED_APPMAINLANDCHINADAILYLIMIT = "LimitExceeded.AppMainlandChinaDailyLimit"
//  LIMITEXCEEDED_DAILYLIMIT = "LimitExceeded.DailyLimit"
//  LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"
//  LIMITEXCEEDED_PHONENUMBERCOUNTLIMIT = "LimitExceeded.PhoneNumberCountLimit"
//  LIMITEXCEEDED_PHONENUMBERDAILYLIMIT = "LimitExceeded.PhoneNumberDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERONEHOURLIMIT = "LimitExceeded.PhoneNumberOneHourLimit"
//  LIMITEXCEEDED_PHONENUMBERSAMECONTENTDAILYLIMIT = "LimitExceeded.PhoneNumberSameContentDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERTHIRTYSECONDLIMIT = "LimitExceeded.PhoneNumberThirtySecondLimit"
//  MISSINGPARAMETER_EMPTYPHONENUMBERSET = "MissingParameter.EmptyPhoneNumberSet"
//  UNAUTHORIZEDOPERATION_INDIVIDUALUSERMARKETINGSMSPERMISSIONDENY = "UnauthorizedOperation.IndividualUserMarketingSmsPermissionDeny"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CHINESEMAINLANDTEMPLATETOGLOBALPHONE = "UnsupportedOperation.ChineseMainlandTemplateToGlobalPhone"
//  UNSUPPORTEDOPERATION_CONTAINDOMESTICANDINTERNATIONALPHONENUMBER = "UnsupportedOperation.ContainDomesticAndInternationalPhoneNumber"
//  UNSUPPORTEDOPERATION_GLOBALTEMPLATETOCHINESEMAINLANDPHONE = "UnsupportedOperation.GlobalTemplateToChineseMainlandPhone"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"
func (c *Client) SendSms(request *SendSmsRequest) (response *SendSmsResponse, err error) {
    return c.SendSmsWithContext(context.Background(), request)
}

// SendSms
// 本接口 (SendSms) 用于发送验证码、通知类短信和营销短信。支持国内短信与国际/港澳台短信。
//
// - 当前接口属于 2021-01-11 版本，如果您仍在使用 [2019-07-11 版本](https://cloud.tencent.com/document/product/382/38778)，建议您使用当前最新版本的接口，版本差异可参考[版本描述](https://cloud.tencent.com/document/product/382/63195#.E7.89.88.E6.9C.AC.E6.8F.8F.E8.BF.B0)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_INSUFFICIENTBALANCEINSMSPACKAGE = "FailedOperation.InsufficientBalanceInSmsPackage"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MARKETINGSENDTIMECONSTRAINT = "FailedOperation.MarketingSendTimeConstraint"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  FAILEDOPERATION_TEMPLATEPARAMSETNOTMATCHAPPROVEDTEMPLATE = "FailedOperation.TemplateParamSetNotMatchApprovedTemplate"
//  FAILEDOPERATION_TEMPLATEUNAPPROVEDORNOTEXIST = "FailedOperation.TemplateUnapprovedOrNotExist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_PROHIBITEDUSEURLINTEMPLATEPARAMETER = "InvalidParameterValue.ProhibitedUseUrlInTemplateParameter"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERFORMATERROR = "InvalidParameterValue.TemplateParameterFormatError"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERLENGTHLIMIT = "InvalidParameterValue.TemplateParameterLengthLimit"
//  LIMITEXCEEDED_APPCOUNTRYORREGIONDAILYLIMIT = "LimitExceeded.AppCountryOrRegionDailyLimit"
//  LIMITEXCEEDED_APPCOUNTRYORREGIONINBLACKLIST = "LimitExceeded.AppCountryOrRegionInBlacklist"
//  LIMITEXCEEDED_APPDAILYLIMIT = "LimitExceeded.AppDailyLimit"
//  LIMITEXCEEDED_APPGLOBALDAILYLIMIT = "LimitExceeded.AppGlobalDailyLimit"
//  LIMITEXCEEDED_APPMAINLANDCHINADAILYLIMIT = "LimitExceeded.AppMainlandChinaDailyLimit"
//  LIMITEXCEEDED_DAILYLIMIT = "LimitExceeded.DailyLimit"
//  LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"
//  LIMITEXCEEDED_PHONENUMBERCOUNTLIMIT = "LimitExceeded.PhoneNumberCountLimit"
//  LIMITEXCEEDED_PHONENUMBERDAILYLIMIT = "LimitExceeded.PhoneNumberDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERONEHOURLIMIT = "LimitExceeded.PhoneNumberOneHourLimit"
//  LIMITEXCEEDED_PHONENUMBERSAMECONTENTDAILYLIMIT = "LimitExceeded.PhoneNumberSameContentDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERTHIRTYSECONDLIMIT = "LimitExceeded.PhoneNumberThirtySecondLimit"
//  MISSINGPARAMETER_EMPTYPHONENUMBERSET = "MissingParameter.EmptyPhoneNumberSet"
//  UNAUTHORIZEDOPERATION_INDIVIDUALUSERMARKETINGSMSPERMISSIONDENY = "UnauthorizedOperation.IndividualUserMarketingSmsPermissionDeny"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CHINESEMAINLANDTEMPLATETOGLOBALPHONE = "UnsupportedOperation.ChineseMainlandTemplateToGlobalPhone"
//  UNSUPPORTEDOPERATION_CONTAINDOMESTICANDINTERNATIONALPHONENUMBER = "UnsupportedOperation.ContainDomesticAndInternationalPhoneNumber"
//  UNSUPPORTEDOPERATION_GLOBALTEMPLATETOCHINESEMAINLANDPHONE = "UnsupportedOperation.GlobalTemplateToChineseMainlandPhone"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"
func (c *Client) SendSmsWithContext(ctx context.Context, request *SendSmsRequest) (response *SendSmsResponse, err error) {
    if request == nil {
        request = NewSendSmsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendSms require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendSmsResponse()
    err = c.Send(request, response)
    return
}

func NewSendStatusStatisticsRequest() (request *SendStatusStatisticsRequest) {
    request = &SendStatusStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "SendStatusStatistics")
    
    
    return
}

func NewSendStatusStatisticsResponse() (response *SendStatusStatisticsResponse) {
    response = &SendStatusStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendStatusStatistics
// 本接口 (SendStatusStatistics) 用于统计用户发送短信的数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_PARSEBACKENDRESPONSEFAIL = "InternalError.ParseBackendResponseFail"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_OFFSETVERIFYFAIL = "InvalidParameterValue.OffsetVerifyFail"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendStatusStatistics(request *SendStatusStatisticsRequest) (response *SendStatusStatisticsResponse, err error) {
    return c.SendStatusStatisticsWithContext(context.Background(), request)
}

// SendStatusStatistics
// 本接口 (SendStatusStatistics) 用于统计用户发送短信的数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_PARSEBACKENDRESPONSEFAIL = "InternalError.ParseBackendResponseFail"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_OFFSETVERIFYFAIL = "InvalidParameterValue.OffsetVerifyFail"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendStatusStatisticsWithContext(ctx context.Context, request *SendStatusStatisticsRequest) (response *SendStatusStatisticsResponse, err error) {
    if request == nil {
        request = NewSendStatusStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendStatusStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendStatusStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewSmsPackagesStatisticsRequest() (request *SmsPackagesStatisticsRequest) {
    request = &SmsPackagesStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "SmsPackagesStatistics")
    
    
    return
}

func NewSmsPackagesStatisticsResponse() (response *SmsPackagesStatisticsResponse) {
    response = &SmsPackagesStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SmsPackagesStatistics
// 本接口 (SmsPackagesStatistics) 用于统计用户套餐包数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_PARSEBACKENDRESPONSEFAIL = "InternalError.ParseBackendResponseFail"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_OFFSETVERIFYFAIL = "InvalidParameterValue.OffsetVerifyFail"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SmsPackagesStatistics(request *SmsPackagesStatisticsRequest) (response *SmsPackagesStatisticsResponse, err error) {
    return c.SmsPackagesStatisticsWithContext(context.Background(), request)
}

// SmsPackagesStatistics
// 本接口 (SmsPackagesStatistics) 用于统计用户套餐包数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_PARSEBACKENDRESPONSEFAIL = "InternalError.ParseBackendResponseFail"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_OFFSETVERIFYFAIL = "InvalidParameterValue.OffsetVerifyFail"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SmsPackagesStatisticsWithContext(ctx context.Context, request *SmsPackagesStatisticsRequest) (response *SmsPackagesStatisticsResponse, err error) {
    if request == nil {
        request = NewSmsPackagesStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SmsPackagesStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewSmsPackagesStatisticsResponse()
    err = c.Send(request, response)
    return
}
