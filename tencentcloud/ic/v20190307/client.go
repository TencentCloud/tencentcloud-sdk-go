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

package v20190307

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-07"

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


func NewDescribeAppRequest() (request *DescribeAppRequest) {
    request = &DescribeAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "DescribeApp")
    return
}

func NewDescribeAppResponse() (response *DescribeAppResponse) {
    response = &DescribeAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApp
// 根据应用id查询物联卡应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeApp(request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    if request == nil {
        request = NewDescribeAppRequest()
    }
    response = NewDescribeAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCardRequest() (request *DescribeCardRequest) {
    request = &DescribeCardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "DescribeCard")
    return
}

func NewDescribeCardResponse() (response *DescribeCardResponse) {
    response = &DescribeCardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCard
// 查询卡片详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCard(request *DescribeCardRequest) (response *DescribeCardResponse, err error) {
    if request == nil {
        request = NewDescribeCardRequest()
    }
    response = NewDescribeCardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCardsRequest() (request *DescribeCardsRequest) {
    request = &DescribeCardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "DescribeCards")
    return
}

func NewDescribeCardsResponse() (response *DescribeCardsResponse) {
    response = &DescribeCardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCards
// 查询卡片列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCards(request *DescribeCardsRequest) (response *DescribeCardsResponse, err error) {
    if request == nil {
        request = NewDescribeCardsRequest()
    }
    response = NewDescribeCardsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserCardRemarkRequest() (request *ModifyUserCardRemarkRequest) {
    request = &ModifyUserCardRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "ModifyUserCardRemark")
    return
}

func NewModifyUserCardRemarkResponse() (response *ModifyUserCardRemarkResponse) {
    response = &ModifyUserCardRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUserCardRemark
// 编辑卡片备注
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyUserCardRemark(request *ModifyUserCardRemarkRequest) (response *ModifyUserCardRemarkResponse, err error) {
    if request == nil {
        request = NewModifyUserCardRemarkRequest()
    }
    response = NewModifyUserCardRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewRenewCardsRequest() (request *RenewCardsRequest) {
    request = &RenewCardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "RenewCards")
    return
}

func NewRenewCardsResponse() (response *RenewCardsResponse) {
    response = &RenewCardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewCards
// 批量为卡片续费，此接口建议调用至少间隔10s,如果出现返回deal lock failed相关的错误，请过10s再重试。
//
// 续费的必要条件：
//
// 1、单次续费的卡片不可以超过 100张。
//
// 2、接口只支持在控制台购买的卡片进行续费
//
// 3、销户和未激活的卡片不支持续费。
//
// 4、每张物联网卡，续费总周期不能超过24个月
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPNOTFOUND = "ResourceNotFound.AppNotFound"
func (c *Client) RenewCards(request *RenewCardsRequest) (response *RenewCardsResponse, err error) {
    if request == nil {
        request = NewRenewCardsRequest()
    }
    response = NewRenewCardsResponse()
    err = c.Send(request, response)
    return
}

func NewSendMultiSmsRequest() (request *SendMultiSmsRequest) {
    request = &SendMultiSmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "SendMultiSms")
    return
}

func NewSendMultiSmsResponse() (response *SendMultiSmsResponse) {
    response = &SendMultiSmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendMultiSms
// 群发短信
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendMultiSms(request *SendMultiSmsRequest) (response *SendMultiSmsResponse, err error) {
    if request == nil {
        request = NewSendMultiSmsRequest()
    }
    response = NewSendMultiSmsResponse()
    err = c.Send(request, response)
    return
}

func NewSendSmsRequest() (request *SendSmsRequest) {
    request = &SendSmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "SendSms")
    return
}

func NewSendSmsResponse() (response *SendSmsResponse) {
    response = &SendSmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendSms
// 发送短信息接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendSms(request *SendSmsRequest) (response *SendSmsResponse, err error) {
    if request == nil {
        request = NewSendSmsRequest()
    }
    response = NewSendSmsResponse()
    err = c.Send(request, response)
    return
}
