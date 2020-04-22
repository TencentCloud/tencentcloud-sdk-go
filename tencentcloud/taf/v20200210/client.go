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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewDetectAccountActivityRequest() (request *DetectAccountActivityRequest) {
    request = &DetectAccountActivityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("taf", APIVersion, "DetectAccountActivity")
    return
}

func NewDetectAccountActivityResponse() (response *DetectAccountActivityResponse) {
    response = &DetectAccountActivityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectAccountActivity
func (c *Client) DetectAccountActivity(request *DetectAccountActivityRequest) (response *DetectAccountActivityResponse, err error) {
    if request == nil {
        request = NewDetectAccountActivityRequest()
    }
    response = NewDetectAccountActivityResponse()
    err = c.Send(request, response)
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
func (c *Client) DetectFraudKOL(request *DetectFraudKOLRequest) (response *DetectFraudKOLResponse, err error) {
    if request == nil {
        request = NewDetectFraudKOLRequest()
    }
    response = NewDetectFraudKOLResponse()
    err = c.Send(request, response)
    return
}

func NewEnhanceTaDegreeRequest() (request *EnhanceTaDegreeRequest) {
    request = &EnhanceTaDegreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("taf", APIVersion, "EnhanceTaDegree")
    return
}

func NewEnhanceTaDegreeResponse() (response *EnhanceTaDegreeResponse) {
    response = &EnhanceTaDegreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnhanceTaDegree
func (c *Client) EnhanceTaDegree(request *EnhanceTaDegreeRequest) (response *EnhanceTaDegreeResponse, err error) {
    if request == nil {
        request = NewEnhanceTaDegreeRequest()
    }
    response = NewEnhanceTaDegreeResponse()
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

// 流量反欺诈-流量验准定制版
func (c *Client) RecognizeCustomizedAudience(request *RecognizeCustomizedAudienceRequest) (response *RecognizeCustomizedAudienceResponse, err error) {
    if request == nil {
        request = NewRecognizeCustomizedAudienceRequest()
    }
    response = NewRecognizeCustomizedAudienceResponse()
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
func (c *Client) RecognizePreciseTargetAudience(request *RecognizePreciseTargetAudienceRequest) (response *RecognizePreciseTargetAudienceResponse, err error) {
    if request == nil {
        request = NewRecognizePreciseTargetAudienceRequest()
    }
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

// 流量反欺诈-流量验准
func (c *Client) RecognizeTargetAudience(request *RecognizeTargetAudienceRequest) (response *RecognizeTargetAudienceResponse, err error) {
    if request == nil {
        request = NewRecognizeTargetAudienceRequest()
    }
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
func (c *Client) SendTrafficSecuritySmsMessage(request *SendTrafficSecuritySmsMessageRequest) (response *SendTrafficSecuritySmsMessageResponse, err error) {
    if request == nil {
        request = NewSendTrafficSecuritySmsMessageRequest()
    }
    response = NewSendTrafficSecuritySmsMessageResponse()
    err = c.Send(request, response)
    return
}
