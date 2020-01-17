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

package v20190711

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-11"

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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// 添加短信签名
func (c *Client) AddSmsSign(request *AddSmsSignRequest) (response *AddSmsSignResponse, err error) {
    if request == nil {
        request = NewAddSmsSignRequest()
    }
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

// 添加短信模板
func (c *Client) AddSmsTemplate(request *AddSmsTemplateRequest) (response *AddSmsTemplateResponse, err error) {
    if request == nil {
        request = NewAddSmsTemplateRequest()
    }
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

// 统计用户回执的数据。
func (c *Client) CallbackStatusStatistics(request *CallbackStatusStatisticsRequest) (response *CallbackStatusStatisticsResponse, err error) {
    if request == nil {
        request = NewCallbackStatusStatisticsRequest()
    }
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

// 删除短信签名
func (c *Client) DeleteSmsSign(request *DeleteSmsSignRequest) (response *DeleteSmsSignResponse, err error) {
    if request == nil {
        request = NewDeleteSmsSignRequest()
    }
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

// 删除短信模板
func (c *Client) DeleteSmsTemplate(request *DeleteSmsTemplateRequest) (response *DeleteSmsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSmsTemplateRequest()
    }
    response = NewDeleteSmsTemplateResponse()
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

// 修改短信签名
func (c *Client) ModifySmsSign(request *ModifySmsSignRequest) (response *ModifySmsSignResponse, err error) {
    if request == nil {
        request = NewModifySmsSignRequest()
    }
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

// 修改短信模板
func (c *Client) ModifySmsTemplate(request *ModifySmsTemplateRequest) (response *ModifySmsTemplateResponse, err error) {
    if request == nil {
        request = NewModifySmsTemplateRequest()
    }
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

// 拉取短信回复状态。
func (c *Client) PullSmsReplyStatus(request *PullSmsReplyStatusRequest) (response *PullSmsReplyStatusResponse, err error) {
    if request == nil {
        request = NewPullSmsReplyStatusRequest()
    }
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

// 拉取单个号码短信回复状态。
func (c *Client) PullSmsReplyStatusByPhoneNumber(request *PullSmsReplyStatusByPhoneNumberRequest) (response *PullSmsReplyStatusByPhoneNumberResponse, err error) {
    if request == nil {
        request = NewPullSmsReplyStatusByPhoneNumberRequest()
    }
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

// 拉取短信下发状态。
func (c *Client) PullSmsSendStatus(request *PullSmsSendStatusRequest) (response *PullSmsSendStatusResponse, err error) {
    if request == nil {
        request = NewPullSmsSendStatusRequest()
    }
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

// 拉取单个号码短信下发状态。
func (c *Client) PullSmsSendStatusByPhoneNumber(request *PullSmsSendStatusByPhoneNumberRequest) (response *PullSmsSendStatusByPhoneNumberResponse, err error) {
    if request == nil {
        request = NewPullSmsSendStatusByPhoneNumberRequest()
    }
    response = NewPullSmsSendStatusByPhoneNumberResponse()
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

// 短信发送接口，用户给用户发短信验证码、通知类短信或营销短信。
// 
func (c *Client) SendSms(request *SendSmsRequest) (response *SendSmsResponse, err error) {
    if request == nil {
        request = NewSendSmsRequest()
    }
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

// 统计用户发送短信的数据。
func (c *Client) SendStatusStatistics(request *SendStatusStatisticsRequest) (response *SendStatusStatisticsResponse, err error) {
    if request == nil {
        request = NewSendStatusStatisticsRequest()
    }
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

// 用户套餐包信息统计。
func (c *Client) SmsPackagesStatistics(request *SmsPackagesStatisticsRequest) (response *SmsPackagesStatisticsResponse, err error) {
    if request == nil {
        request = NewSmsPackagesStatisticsRequest()
    }
    response = NewSmsPackagesStatisticsResponse()
    err = c.Send(request, response)
    return
}
