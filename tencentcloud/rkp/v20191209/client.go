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

package v20191209

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-12-09"

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


func NewGetOpenIdRequest() (request *GetOpenIdRequest) {
    request = &GetOpenIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rkp", APIVersion, "GetOpenId")
    return
}

func NewGetOpenIdResponse() (response *GetOpenIdResponse) {
    response = &GetOpenIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据DevicceToken查询OpenID。
func (c *Client) GetOpenId(request *GetOpenIdRequest) (response *GetOpenIdResponse, err error) {
    if request == nil {
        request = NewGetOpenIdRequest()
    }
    response = NewGetOpenIdResponse()
    err = c.Send(request, response)
    return
}

func NewGetTokenRequest() (request *GetTokenRequest) {
    request = &GetTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rkp", APIVersion, "GetToken")
    return
}

func NewGetTokenResponse() (response *GetTokenResponse) {
    response = &GetTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取token接口。
func (c *Client) GetToken(request *GetTokenRequest) (response *GetTokenResponse, err error) {
    if request == nil {
        request = NewGetTokenRequest()
    }
    response = NewGetTokenResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDevAndRiskRequest() (request *QueryDevAndRiskRequest) {
    request = &QueryDevAndRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rkp", APIVersion, "QueryDevAndRisk")
    return
}

func NewQueryDevAndRiskResponse() (response *QueryDevAndRiskResponse) {
    response = &QueryDevAndRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 腾讯天御设备风险查询接口，输入由客户应用自主采集的设备信息， 通过腾讯大数据风控能力，可以准确根据输入设备信息，还原设备库中的设备ID，并且识别设备的风险，解决客户业务过程中的设备风险，降低企业损失。
func (c *Client) QueryDevAndRisk(request *QueryDevAndRiskRequest) (response *QueryDevAndRiskResponse, err error) {
    if request == nil {
        request = NewQueryDevAndRiskRequest()
    }
    response = NewQueryDevAndRiskResponse()
    err = c.Send(request, response)
    return
}
