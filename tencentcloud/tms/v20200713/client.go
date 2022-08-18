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

package v20200713

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-07-13"

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


func NewAccountTipoffAccessRequest() (request *AccountTipoffAccessRequest) {
    request = &AccountTipoffAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tms", APIVersion, "AccountTipoffAccess")
    
    
    return
}

func NewAccountTipoffAccessResponse() (response *AccountTipoffAccessResponse) {
    response = &AccountTipoffAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AccountTipoffAccess
// 举报恶意账号
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRACTION = "InvalidParameter.ErrAction"
//  INVALIDPARAMETER_INVALIDPARAMETERCONTENT = "InvalidParameter.InvalidParameterContent"
//  INVALIDPARAMETERVALUE_INVALIDEVILCONTENT = "InvalidParameterValue.InvalidEvilContent"
//  INVALIDPARAMETERVALUE_INVALIDEVILTYPE = "InvalidParameterValue.InvalidEvilType"
//  INVALIDPARAMETERVALUE_INVALIDREPORTEDACCOUNT = "InvalidParameterValue.InvalidReportedAccount"
//  INVALIDPARAMETERVALUE_INVALIDREPORTEDACCOUNTLENGTH = "InvalidParameterValue.InvalidReportedAccountLength"
//  INVALIDPARAMETERVALUE_INVALIDREPORTEDACCOUNTTYPE = "InvalidParameterValue.InvalidReportedAccountType"
//  INVALIDPARAMETERVALUE_INVALIDSENDERACCOUNT = "InvalidParameterValue.InvalidSenderAccount"
//  INVALIDPARAMETERVALUE_INVALIDSENDERACCOUNTLENGTH = "InvalidParameterValue.InvalidSenderAccountLength"
//  INVALIDPARAMETERVALUE_INVALIDSENDERACCOUNTTYPE = "InvalidParameterValue.InvalidSenderAccountType"
//  INVALIDPARAMETERVALUE_INVALIDSENDERIP = "InvalidParameterValue.InvalidSenderIP"
func (c *Client) AccountTipoffAccess(request *AccountTipoffAccessRequest) (response *AccountTipoffAccessResponse, err error) {
    return c.AccountTipoffAccessWithContext(context.Background(), request)
}

// AccountTipoffAccess
// 举报恶意账号
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRACTION = "InvalidParameter.ErrAction"
//  INVALIDPARAMETER_INVALIDPARAMETERCONTENT = "InvalidParameter.InvalidParameterContent"
//  INVALIDPARAMETERVALUE_INVALIDEVILCONTENT = "InvalidParameterValue.InvalidEvilContent"
//  INVALIDPARAMETERVALUE_INVALIDEVILTYPE = "InvalidParameterValue.InvalidEvilType"
//  INVALIDPARAMETERVALUE_INVALIDREPORTEDACCOUNT = "InvalidParameterValue.InvalidReportedAccount"
//  INVALIDPARAMETERVALUE_INVALIDREPORTEDACCOUNTLENGTH = "InvalidParameterValue.InvalidReportedAccountLength"
//  INVALIDPARAMETERVALUE_INVALIDREPORTEDACCOUNTTYPE = "InvalidParameterValue.InvalidReportedAccountType"
//  INVALIDPARAMETERVALUE_INVALIDSENDERACCOUNT = "InvalidParameterValue.InvalidSenderAccount"
//  INVALIDPARAMETERVALUE_INVALIDSENDERACCOUNTLENGTH = "InvalidParameterValue.InvalidSenderAccountLength"
//  INVALIDPARAMETERVALUE_INVALIDSENDERACCOUNTTYPE = "InvalidParameterValue.InvalidSenderAccountType"
//  INVALIDPARAMETERVALUE_INVALIDSENDERIP = "InvalidParameterValue.InvalidSenderIP"
func (c *Client) AccountTipoffAccessWithContext(ctx context.Context, request *AccountTipoffAccessRequest) (response *AccountTipoffAccessResponse, err error) {
    if request == nil {
        request = NewAccountTipoffAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AccountTipoffAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewAccountTipoffAccessResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTextLibRequest() (request *DescribeTextLibRequest) {
    request = &DescribeTextLibRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tms", APIVersion, "DescribeTextLib")
    
    
    return
}

func NewDescribeTextLibResponse() (response *DescribeTextLibResponse) {
    response = &DescribeTextLibResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTextLib
// 控制台获取用户词库列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTextLib(request *DescribeTextLibRequest) (response *DescribeTextLibResponse, err error) {
    return c.DescribeTextLibWithContext(context.Background(), request)
}

// DescribeTextLib
// 控制台获取用户词库列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTextLibWithContext(ctx context.Context, request *DescribeTextLibRequest) (response *DescribeTextLibResponse, err error) {
    if request == nil {
        request = NewDescribeTextLibRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTextLib require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTextLibResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTextStatRequest() (request *DescribeTextStatRequest) {
    request = &DescribeTextStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tms", APIVersion, "DescribeTextStat")
    
    
    return
}

func NewDescribeTextStatResponse() (response *DescribeTextStatResponse) {
    response = &DescribeTextStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTextStat
// 控制台识别统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRTEXTTIMEOUT = "InternalError.ErrTextTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRTEXTCONTENTTYPE = "InvalidParameter.ErrTextContentType"
func (c *Client) DescribeTextStat(request *DescribeTextStatRequest) (response *DescribeTextStatResponse, err error) {
    return c.DescribeTextStatWithContext(context.Background(), request)
}

// DescribeTextStat
// 控制台识别统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRTEXTTIMEOUT = "InternalError.ErrTextTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRTEXTCONTENTTYPE = "InvalidParameter.ErrTextContentType"
func (c *Client) DescribeTextStatWithContext(ctx context.Context, request *DescribeTextStatRequest) (response *DescribeTextStatResponse, err error) {
    if request == nil {
        request = NewDescribeTextStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTextStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTextStatResponse()
    err = c.Send(request, response)
    return
}

func NewTextModerationRequest() (request *TextModerationRequest) {
    request = &TextModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tms", APIVersion, "TextModeration")
    
    
    return
}

func NewTextModerationResponse() (response *TextModerationResponse) {
    response = &TextModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextModeration
// 文本内容检测（Text Moderation）服务使用了深度学习技术，识别可能令人反感、不安全或不适宜的文本内容，同时支持用户配置词库黑白名单，打击自定义识别类型的图片。
//
// 可能返回的错误码:
//  INTERNALERROR_ERRTEXTTIMEOUT = "InternalError.ErrTextTimeOut"
//  INVALIDPARAMETER_ERRACTION = "InvalidParameter.ErrAction"
//  INVALIDPARAMETER_ERRTEXTCONTENTLEN = "InvalidParameter.ErrTextContentLen"
//  INVALIDPARAMETER_ERRTEXTCONTENTTYPE = "InvalidParameter.ErrTextContentType"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTTYPE = "InvalidParameterValue.ErrTextContentType"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) TextModeration(request *TextModerationRequest) (response *TextModerationResponse, err error) {
    return c.TextModerationWithContext(context.Background(), request)
}

// TextModeration
// 文本内容检测（Text Moderation）服务使用了深度学习技术，识别可能令人反感、不安全或不适宜的文本内容，同时支持用户配置词库黑白名单，打击自定义识别类型的图片。
//
// 可能返回的错误码:
//  INTERNALERROR_ERRTEXTTIMEOUT = "InternalError.ErrTextTimeOut"
//  INVALIDPARAMETER_ERRACTION = "InvalidParameter.ErrAction"
//  INVALIDPARAMETER_ERRTEXTCONTENTLEN = "InvalidParameter.ErrTextContentLen"
//  INVALIDPARAMETER_ERRTEXTCONTENTTYPE = "InvalidParameter.ErrTextContentType"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTTYPE = "InvalidParameterValue.ErrTextContentType"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) TextModerationWithContext(ctx context.Context, request *TextModerationRequest) (response *TextModerationResponse, err error) {
    if request == nil {
        request = NewTextModerationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextModerationResponse()
    err = c.Send(request, response)
    return
}
