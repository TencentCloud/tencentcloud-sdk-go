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

package v20190722

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-22"

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


func NewDescribeCaptchaAppIdInfoRequest() (request *DescribeCaptchaAppIdInfoRequest) {
    request = &DescribeCaptchaAppIdInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaAppIdInfo")
    
    
    return
}

func NewDescribeCaptchaAppIdInfoResponse() (response *DescribeCaptchaAppIdInfoResponse) {
    response = &DescribeCaptchaAppIdInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaAppIdInfo
// 查询安全验证码应用APPId信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaAppIdInfo(request *DescribeCaptchaAppIdInfoRequest) (response *DescribeCaptchaAppIdInfoResponse, err error) {
    return c.DescribeCaptchaAppIdInfoWithContext(context.Background(), request)
}

// DescribeCaptchaAppIdInfo
// 查询安全验证码应用APPId信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaAppIdInfoWithContext(ctx context.Context, request *DescribeCaptchaAppIdInfoRequest) (response *DescribeCaptchaAppIdInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaAppIdInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaAppIdInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaAppIdInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaDataRequest() (request *DescribeCaptchaDataRequest) {
    request = &DescribeCaptchaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaData")
    
    
    return
}

func NewDescribeCaptchaDataResponse() (response *DescribeCaptchaDataResponse) {
    response = &DescribeCaptchaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaData
// 安全验证码分类查询数据接口，请求量type=0、通过量type=1、验证量type=2、拦截量type=3  分钟级查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaData(request *DescribeCaptchaDataRequest) (response *DescribeCaptchaDataResponse, err error) {
    return c.DescribeCaptchaDataWithContext(context.Background(), request)
}

// DescribeCaptchaData
// 安全验证码分类查询数据接口，请求量type=0、通过量type=1、验证量type=2、拦截量type=3  分钟级查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaDataWithContext(ctx context.Context, request *DescribeCaptchaDataRequest) (response *DescribeCaptchaDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaDataSumRequest() (request *DescribeCaptchaDataSumRequest) {
    request = &DescribeCaptchaDataSumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaDataSum")
    
    
    return
}

func NewDescribeCaptchaDataSumResponse() (response *DescribeCaptchaDataSumResponse) {
    response = &DescribeCaptchaDataSumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaDataSum
// 安全验证码查询请求数据概况，例如：按照时间段查询数据  昨日请求量、昨日恶意比例、昨日验证量、昨日通过量、昨日恶意拦截量……
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaDataSum(request *DescribeCaptchaDataSumRequest) (response *DescribeCaptchaDataSumResponse, err error) {
    return c.DescribeCaptchaDataSumWithContext(context.Background(), request)
}

// DescribeCaptchaDataSum
// 安全验证码查询请求数据概况，例如：按照时间段查询数据  昨日请求量、昨日恶意比例、昨日验证量、昨日通过量、昨日恶意拦截量……
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaDataSumWithContext(ctx context.Context, request *DescribeCaptchaDataSumRequest) (response *DescribeCaptchaDataSumResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaDataSumRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaDataSum require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaDataSumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniDataRequest() (request *DescribeCaptchaMiniDataRequest) {
    request = &DescribeCaptchaMiniDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniData")
    
    
    return
}

func NewDescribeCaptchaMiniDataResponse() (response *DescribeCaptchaMiniDataResponse) {
    response = &DescribeCaptchaMiniDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaMiniData
// 安全验证码小程序插件分类查询数据接口，请求量type=0、通过量type=1、验证量type=2、拦截量type=3 小时级查询（五小时左右延迟）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniData(request *DescribeCaptchaMiniDataRequest) (response *DescribeCaptchaMiniDataResponse, err error) {
    return c.DescribeCaptchaMiniDataWithContext(context.Background(), request)
}

// DescribeCaptchaMiniData
// 安全验证码小程序插件分类查询数据接口，请求量type=0、通过量type=1、验证量type=2、拦截量type=3 小时级查询（五小时左右延迟）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniDataWithContext(ctx context.Context, request *DescribeCaptchaMiniDataRequest) (response *DescribeCaptchaMiniDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaMiniData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaMiniDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniDataSumRequest() (request *DescribeCaptchaMiniDataSumRequest) {
    request = &DescribeCaptchaMiniDataSumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniDataSum")
    
    
    return
}

func NewDescribeCaptchaMiniDataSumResponse() (response *DescribeCaptchaMiniDataSumResponse) {
    response = &DescribeCaptchaMiniDataSumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaMiniDataSum
// 安全验证码小程序插件查询请求数据概况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniDataSum(request *DescribeCaptchaMiniDataSumRequest) (response *DescribeCaptchaMiniDataSumResponse, err error) {
    return c.DescribeCaptchaMiniDataSumWithContext(context.Background(), request)
}

// DescribeCaptchaMiniDataSum
// 安全验证码小程序插件查询请求数据概况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniDataSumWithContext(ctx context.Context, request *DescribeCaptchaMiniDataSumRequest) (response *DescribeCaptchaMiniDataSumResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniDataSumRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaMiniDataSum require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaMiniDataSumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniOperDataRequest() (request *DescribeCaptchaMiniOperDataRequest) {
    request = &DescribeCaptchaMiniOperDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniOperData")
    
    
    return
}

func NewDescribeCaptchaMiniOperDataResponse() (response *DescribeCaptchaMiniOperDataResponse) {
    response = &DescribeCaptchaMiniOperDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaMiniOperData
// 安全验证码小程序插件用户操作数据查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniOperData(request *DescribeCaptchaMiniOperDataRequest) (response *DescribeCaptchaMiniOperDataResponse, err error) {
    return c.DescribeCaptchaMiniOperDataWithContext(context.Background(), request)
}

// DescribeCaptchaMiniOperData
// 安全验证码小程序插件用户操作数据查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniOperDataWithContext(ctx context.Context, request *DescribeCaptchaMiniOperDataRequest) (response *DescribeCaptchaMiniOperDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniOperDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaMiniOperData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaMiniOperDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniResultRequest() (request *DescribeCaptchaMiniResultRequest) {
    request = &DescribeCaptchaMiniResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniResult")
    
    
    return
}

func NewDescribeCaptchaMiniResultResponse() (response *DescribeCaptchaMiniResultResponse) {
    response = &DescribeCaptchaMiniResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaMiniResult
// 核查验证码票据结果(小程序插件) 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniResult(request *DescribeCaptchaMiniResultRequest) (response *DescribeCaptchaMiniResultResponse, err error) {
    return c.DescribeCaptchaMiniResultWithContext(context.Background(), request)
}

// DescribeCaptchaMiniResult
// 核查验证码票据结果(小程序插件) 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniResultWithContext(ctx context.Context, request *DescribeCaptchaMiniResultRequest) (response *DescribeCaptchaMiniResultResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaMiniResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaMiniResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniRiskResultRequest() (request *DescribeCaptchaMiniRiskResultRequest) {
    request = &DescribeCaptchaMiniRiskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniRiskResult")
    
    
    return
}

func NewDescribeCaptchaMiniRiskResultResponse() (response *DescribeCaptchaMiniRiskResultResponse) {
    response = &DescribeCaptchaMiniRiskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaMiniRiskResult
// 核查验证码小程序插件票据接入风控结果(已停用)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniRiskResult(request *DescribeCaptchaMiniRiskResultRequest) (response *DescribeCaptchaMiniRiskResultResponse, err error) {
    return c.DescribeCaptchaMiniRiskResultWithContext(context.Background(), request)
}

// DescribeCaptchaMiniRiskResult
// 核查验证码小程序插件票据接入风控结果(已停用)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaMiniRiskResultWithContext(ctx context.Context, request *DescribeCaptchaMiniRiskResultRequest) (response *DescribeCaptchaMiniRiskResultResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniRiskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaMiniRiskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaMiniRiskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaOperDataRequest() (request *DescribeCaptchaOperDataRequest) {
    request = &DescribeCaptchaOperDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaOperData")
    
    
    return
}

func NewDescribeCaptchaOperDataResponse() (response *DescribeCaptchaOperDataResponse) {
    response = &DescribeCaptchaOperDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaOperData
// 安全验证码用户操作数据查询，验证码加载耗时type = 1 、拦截情况type = 2、 一周通过平均尝试次数 type = 3、尝试次数分布 type = 4
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaOperData(request *DescribeCaptchaOperDataRequest) (response *DescribeCaptchaOperDataResponse, err error) {
    return c.DescribeCaptchaOperDataWithContext(context.Background(), request)
}

// DescribeCaptchaOperData
// 安全验证码用户操作数据查询，验证码加载耗时type = 1 、拦截情况type = 2、 一周通过平均尝试次数 type = 3、尝试次数分布 type = 4
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaOperDataWithContext(ctx context.Context, request *DescribeCaptchaOperDataRequest) (response *DescribeCaptchaOperDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaOperDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaOperData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaOperDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaResultRequest() (request *DescribeCaptchaResultRequest) {
    request = &DescribeCaptchaResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaResult")
    
    
    return
}

func NewDescribeCaptchaResultResponse() (response *DescribeCaptchaResultResponse) {
    response = &DescribeCaptchaResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaResult
// 核查验证码票据结果(Web及APP)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaResult(request *DescribeCaptchaResultRequest) (response *DescribeCaptchaResultResponse, err error) {
    return c.DescribeCaptchaResultWithContext(context.Background(), request)
}

// DescribeCaptchaResult
// 核查验证码票据结果(Web及APP)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaResultWithContext(ctx context.Context, request *DescribeCaptchaResultRequest) (response *DescribeCaptchaResultResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaTicketDataRequest() (request *DescribeCaptchaTicketDataRequest) {
    request = &DescribeCaptchaTicketDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaTicketData")
    
    
    return
}

func NewDescribeCaptchaTicketDataResponse() (response *DescribeCaptchaTicketDataResponse) {
    response = &DescribeCaptchaTicketDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaTicketData
// 安全验证码用户操作票据数据查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaTicketData(request *DescribeCaptchaTicketDataRequest) (response *DescribeCaptchaTicketDataResponse, err error) {
    return c.DescribeCaptchaTicketDataWithContext(context.Background(), request)
}

// DescribeCaptchaTicketData
// 安全验证码用户操作票据数据查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaTicketDataWithContext(ctx context.Context, request *DescribeCaptchaTicketDataRequest) (response *DescribeCaptchaTicketDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaTicketDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaTicketData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaTicketDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaUserAllAppIdRequest() (request *DescribeCaptchaUserAllAppIdRequest) {
    request = &DescribeCaptchaUserAllAppIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaUserAllAppId")
    
    
    return
}

func NewDescribeCaptchaUserAllAppIdResponse() (response *DescribeCaptchaUserAllAppIdResponse) {
    response = &DescribeCaptchaUserAllAppIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCaptchaUserAllAppId
// 安全验证码获取用户注册所有APPId和应用名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaUserAllAppId(request *DescribeCaptchaUserAllAppIdRequest) (response *DescribeCaptchaUserAllAppIdResponse, err error) {
    return c.DescribeCaptchaUserAllAppIdWithContext(context.Background(), request)
}

// DescribeCaptchaUserAllAppId
// 安全验证码获取用户注册所有APPId和应用名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaUserAllAppIdWithContext(ctx context.Context, request *DescribeCaptchaUserAllAppIdRequest) (response *DescribeCaptchaUserAllAppIdResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaUserAllAppIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaUserAllAppId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaUserAllAppIdResponse()
    err = c.Send(request, response)
    return
}

func NewGetRequestStatisticsRequest() (request *GetRequestStatisticsRequest) {
    request = &GetRequestStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "GetRequestStatistics")
    
    
    return
}

func NewGetRequestStatisticsResponse() (response *GetRequestStatisticsResponse) {
    response = &GetRequestStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRequestStatistics
// 查询单个CaptchaAppID验证的统计数据，包括：请求量、验证量、验证通过量、验证拦截量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetRequestStatistics(request *GetRequestStatisticsRequest) (response *GetRequestStatisticsResponse, err error) {
    return c.GetRequestStatisticsWithContext(context.Background(), request)
}

// GetRequestStatistics
// 查询单个CaptchaAppID验证的统计数据，包括：请求量、验证量、验证通过量、验证拦截量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetRequestStatisticsWithContext(ctx context.Context, request *GetRequestStatisticsRequest) (response *GetRequestStatisticsResponse, err error) {
    if request == nil {
        request = NewGetRequestStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRequestStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRequestStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewGetTicketStatisticsRequest() (request *GetTicketStatisticsRequest) {
    request = &GetTicketStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "GetTicketStatistics")
    
    
    return
}

func NewGetTicketStatisticsResponse() (response *GetTicketStatisticsResponse) {
    response = &GetTicketStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTicketStatistics
// 查询单个CaptchaAppID票据校验数据，包括：票据校验量、票据校验通过量、票据校验拦截量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetTicketStatistics(request *GetTicketStatisticsRequest) (response *GetTicketStatisticsResponse, err error) {
    return c.GetTicketStatisticsWithContext(context.Background(), request)
}

// GetTicketStatistics
// 查询单个CaptchaAppID票据校验数据，包括：票据校验量、票据校验通过量、票据校验拦截量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetTicketStatisticsWithContext(ctx context.Context, request *GetTicketStatisticsRequest) (response *GetTicketStatisticsResponse, err error) {
    if request == nil {
        request = NewGetTicketStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTicketStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTicketStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewGetTotalRequestStatisticsRequest() (request *GetTotalRequestStatisticsRequest) {
    request = &GetTotalRequestStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "GetTotalRequestStatistics")
    
    
    return
}

func NewGetTotalRequestStatisticsResponse() (response *GetTotalRequestStatisticsResponse) {
    response = &GetTotalRequestStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTotalRequestStatistics
// 查询全部验证的统计数据，包括：总请求量、总验证量、总验证通过量、总验证拦截量等数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetTotalRequestStatistics(request *GetTotalRequestStatisticsRequest) (response *GetTotalRequestStatisticsResponse, err error) {
    return c.GetTotalRequestStatisticsWithContext(context.Background(), request)
}

// GetTotalRequestStatistics
// 查询全部验证的统计数据，包括：总请求量、总验证量、总验证通过量、总验证拦截量等数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetTotalRequestStatisticsWithContext(ctx context.Context, request *GetTotalRequestStatisticsRequest) (response *GetTotalRequestStatisticsResponse, err error) {
    if request == nil {
        request = NewGetTotalRequestStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTotalRequestStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTotalRequestStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewGetTotalTicketStatisticsRequest() (request *GetTotalTicketStatisticsRequest) {
    request = &GetTotalTicketStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "GetTotalTicketStatistics")
    
    
    return
}

func NewGetTotalTicketStatisticsResponse() (response *GetTotalTicketStatisticsResponse) {
    response = &GetTotalTicketStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTotalTicketStatistics
// 查询全部票据校验的统计数据，包括：总票据校验量、总票据校验通过量、总票据校验拦截量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetTotalTicketStatistics(request *GetTotalTicketStatisticsRequest) (response *GetTotalTicketStatisticsResponse, err error) {
    return c.GetTotalTicketStatisticsWithContext(context.Background(), request)
}

// GetTotalTicketStatistics
// 查询全部票据校验的统计数据，包括：总票据校验量、总票据校验通过量、总票据校验拦截量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetTotalTicketStatisticsWithContext(ctx context.Context, request *GetTotalTicketStatisticsRequest) (response *GetTotalTicketStatisticsResponse, err error) {
    if request == nil {
        request = NewGetTotalTicketStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTotalTicketStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTotalTicketStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCaptchaAppIdInfoRequest() (request *UpdateCaptchaAppIdInfoRequest) {
    request = &UpdateCaptchaAppIdInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "UpdateCaptchaAppIdInfo")
    
    
    return
}

func NewUpdateCaptchaAppIdInfoResponse() (response *UpdateCaptchaAppIdInfoResponse) {
    response = &UpdateCaptchaAppIdInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCaptchaAppIdInfo
// 更新验证码应用APPId信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) UpdateCaptchaAppIdInfo(request *UpdateCaptchaAppIdInfoRequest) (response *UpdateCaptchaAppIdInfoResponse, err error) {
    return c.UpdateCaptchaAppIdInfoWithContext(context.Background(), request)
}

// UpdateCaptchaAppIdInfo
// 更新验证码应用APPId信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) UpdateCaptchaAppIdInfoWithContext(ctx context.Context, request *UpdateCaptchaAppIdInfoRequest) (response *UpdateCaptchaAppIdInfoResponse, err error) {
    if request == nil {
        request = NewUpdateCaptchaAppIdInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCaptchaAppIdInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCaptchaAppIdInfoResponse()
    err = c.Send(request, response)
    return
}
