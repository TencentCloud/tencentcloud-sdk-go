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

package v20221115

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-11-15"

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


func NewCreateBPBrandRequest() (request *CreateBPBrandRequest) {
    request = &CreateBPBrandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPBrand")
    
    
    return
}

func NewCreateBPBrandResponse() (response *CreateBPBrandResponse) {
    response = &CreateBPBrandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPBrand
// 添加品牌
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPBrand(request *CreateBPBrandRequest) (response *CreateBPBrandResponse, err error) {
    return c.CreateBPBrandWithContext(context.Background(), request)
}

// CreateBPBrand
// 添加品牌
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPBrandWithContext(ctx context.Context, request *CreateBPBrandRequest) (response *CreateBPBrandResponse, err error) {
    if request == nil {
        request = NewCreateBPBrandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPBrand require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPBrandResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBPFakeAPPRequest() (request *CreateBPFakeAPPRequest) {
    request = &CreateBPFakeAPPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPFakeAPP")
    
    
    return
}

func NewCreateBPFakeAPPResponse() (response *CreateBPFakeAPPResponse) {
    response = &CreateBPFakeAPPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPFakeAPP
// 仿冒应用举报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPFakeAPP(request *CreateBPFakeAPPRequest) (response *CreateBPFakeAPPResponse, err error) {
    return c.CreateBPFakeAPPWithContext(context.Background(), request)
}

// CreateBPFakeAPP
// 仿冒应用举报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPFakeAPPWithContext(ctx context.Context, request *CreateBPFakeAPPRequest) (response *CreateBPFakeAPPResponse, err error) {
    if request == nil {
        request = NewCreateBPFakeAPPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPFakeAPP require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPFakeAPPResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBPFakeAPPListRequest() (request *CreateBPFakeAPPListRequest) {
    request = &CreateBPFakeAPPListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPFakeAPPList")
    
    
    return
}

func NewCreateBPFakeAPPListResponse() (response *CreateBPFakeAPPListResponse) {
    response = &CreateBPFakeAPPListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPFakeAPPList
// 批量仿冒应用举报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPFakeAPPList(request *CreateBPFakeAPPListRequest) (response *CreateBPFakeAPPListResponse, err error) {
    return c.CreateBPFakeAPPListWithContext(context.Background(), request)
}

// CreateBPFakeAPPList
// 批量仿冒应用举报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPFakeAPPListWithContext(ctx context.Context, request *CreateBPFakeAPPListRequest) (response *CreateBPFakeAPPListResponse, err error) {
    if request == nil {
        request = NewCreateBPFakeAPPListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPFakeAPPList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPFakeAPPListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBPFakeURLRequest() (request *CreateBPFakeURLRequest) {
    request = &CreateBPFakeURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPFakeURL")
    
    
    return
}

func NewCreateBPFakeURLResponse() (response *CreateBPFakeURLResponse) {
    response = &CreateBPFakeURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPFakeURL
// 仿冒网址举报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPFakeURL(request *CreateBPFakeURLRequest) (response *CreateBPFakeURLResponse, err error) {
    return c.CreateBPFakeURLWithContext(context.Background(), request)
}

// CreateBPFakeURL
// 仿冒网址举报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPFakeURLWithContext(ctx context.Context, request *CreateBPFakeURLRequest) (response *CreateBPFakeURLResponse, err error) {
    if request == nil {
        request = NewCreateBPFakeURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPFakeURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPFakeURLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBPFakeURLsRequest() (request *CreateBPFakeURLsRequest) {
    request = &CreateBPFakeURLsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPFakeURLs")
    
    
    return
}

func NewCreateBPFakeURLsResponse() (response *CreateBPFakeURLsResponse) {
    response = &CreateBPFakeURLsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPFakeURLs
// 批量仿冒网址举报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPFakeURLs(request *CreateBPFakeURLsRequest) (response *CreateBPFakeURLsResponse, err error) {
    return c.CreateBPFakeURLsWithContext(context.Background(), request)
}

// CreateBPFakeURLs
// 批量仿冒网址举报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPFakeURLsWithContext(ctx context.Context, request *CreateBPFakeURLsRequest) (response *CreateBPFakeURLsResponse, err error) {
    if request == nil {
        request = NewCreateBPFakeURLsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPFakeURLs require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPFakeURLsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBPWhiteListRequest() (request *CreateBPWhiteListRequest) {
    request = &CreateBPWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPWhiteList")
    
    
    return
}

func NewCreateBPWhiteListResponse() (response *CreateBPWhiteListResponse) {
    response = &CreateBPWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPWhiteList
// 添加白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPWhiteList(request *CreateBPWhiteListRequest) (response *CreateBPWhiteListResponse, err error) {
    return c.CreateBPWhiteListWithContext(context.Background(), request)
}

// CreateBPWhiteList
// 添加白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateBPWhiteListWithContext(ctx context.Context, request *CreateBPWhiteListRequest) (response *CreateBPWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateBPWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBPWhiteListRequest() (request *DeleteBPWhiteListRequest) {
    request = &DeleteBPWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DeleteBPWhiteList")
    
    
    return
}

func NewDeleteBPWhiteListResponse() (response *DeleteBPWhiteListResponse) {
    response = &DeleteBPWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBPWhiteList
// 删除白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DeleteBPWhiteList(request *DeleteBPWhiteListRequest) (response *DeleteBPWhiteListResponse, err error) {
    return c.DeleteBPWhiteListWithContext(context.Background(), request)
}

// DeleteBPWhiteList
// 删除白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DeleteBPWhiteListWithContext(ctx context.Context, request *DeleteBPWhiteListRequest) (response *DeleteBPWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteBPWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBPWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBPWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBPBrandsRequest() (request *DescribeBPBrandsRequest) {
    request = &DescribeBPBrandsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeBPBrands")
    
    
    return
}

func NewDescribeBPBrandsResponse() (response *DescribeBPBrandsResponse) {
    response = &DescribeBPBrandsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBPBrands
// 查询品牌列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBPBrands(request *DescribeBPBrandsRequest) (response *DescribeBPBrandsResponse, err error) {
    return c.DescribeBPBrandsWithContext(context.Background(), request)
}

// DescribeBPBrands
// 查询品牌列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBPBrandsWithContext(ctx context.Context, request *DescribeBPBrandsRequest) (response *DescribeBPBrandsResponse, err error) {
    if request == nil {
        request = NewDescribeBPBrandsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBPBrands require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBPBrandsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBPFakeAPPListRequest() (request *DescribeBPFakeAPPListRequest) {
    request = &DescribeBPFakeAPPListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeBPFakeAPPList")
    
    
    return
}

func NewDescribeBPFakeAPPListResponse() (response *DescribeBPFakeAPPListResponse) {
    response = &DescribeBPFakeAPPListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBPFakeAPPList
// 查询仿冒应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeBPFakeAPPList(request *DescribeBPFakeAPPListRequest) (response *DescribeBPFakeAPPListResponse, err error) {
    return c.DescribeBPFakeAPPListWithContext(context.Background(), request)
}

// DescribeBPFakeAPPList
// 查询仿冒应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeBPFakeAPPListWithContext(ctx context.Context, request *DescribeBPFakeAPPListRequest) (response *DescribeBPFakeAPPListResponse, err error) {
    if request == nil {
        request = NewDescribeBPFakeAPPListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBPFakeAPPList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBPFakeAPPListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBPFakeURLsRequest() (request *DescribeBPFakeURLsRequest) {
    request = &DescribeBPFakeURLsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeBPFakeURLs")
    
    
    return
}

func NewDescribeBPFakeURLsResponse() (response *DescribeBPFakeURLsResponse) {
    response = &DescribeBPFakeURLsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBPFakeURLs
// 查询仿冒网址列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeBPFakeURLs(request *DescribeBPFakeURLsRequest) (response *DescribeBPFakeURLsResponse, err error) {
    return c.DescribeBPFakeURLsWithContext(context.Background(), request)
}

// DescribeBPFakeURLs
// 查询仿冒网址列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeBPFakeURLsWithContext(ctx context.Context, request *DescribeBPFakeURLsRequest) (response *DescribeBPFakeURLsResponse, err error) {
    if request == nil {
        request = NewDescribeBPFakeURLsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBPFakeURLs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBPFakeURLsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBPWhiteListsRequest() (request *DescribeBPWhiteListsRequest) {
    request = &DescribeBPWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeBPWhiteLists")
    
    
    return
}

func NewDescribeBPWhiteListsResponse() (response *DescribeBPWhiteListsResponse) {
    response = &DescribeBPWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBPWhiteLists
// 查询白名单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeBPWhiteLists(request *DescribeBPWhiteListsRequest) (response *DescribeBPWhiteListsResponse, err error) {
    return c.DescribeBPWhiteListsWithContext(context.Background(), request)
}

// DescribeBPWhiteLists
// 查询白名单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"
//  FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"
//  FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"
//  FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"
//  FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"
//  FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"
//  FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"
//  INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"
//  INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeBPWhiteListsWithContext(ctx context.Context, request *DescribeBPWhiteListsRequest) (response *DescribeBPWhiteListsResponse, err error) {
    if request == nil {
        request = NewDescribeBPWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBPWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBPWhiteListsResponse()
    err = c.Send(request, response)
    return
}
