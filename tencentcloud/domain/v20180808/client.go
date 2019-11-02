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

package v20180808

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-08-08"

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


func NewCheckDomainRequest() (request *CheckDomainRequest) {
    request = &CheckDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "CheckDomain")
    return
}

func NewCheckDomainResponse() (response *CheckDomainResponse) {
    response = &CheckDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查域名是否可以注册
func (c *Client) CheckDomain(request *CheckDomainRequest) (response *CheckDomainResponse, err error) {
    if request == nil {
        request = NewCheckDomainRequest()
    }
    response = NewCheckDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainPriceListRequest() (request *DescribeDomainPriceListRequest) {
    request = &DescribeDomainPriceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainPriceList")
    return
}

func NewDescribeDomainPriceListResponse() (response *DescribeDomainPriceListResponse) {
    response = &DescribeDomainPriceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按照域名后缀获取对应的价格列表
func (c *Client) DescribeDomainPriceList(request *DescribeDomainPriceListRequest) (response *DescribeDomainPriceListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainPriceListRequest()
    }
    response = NewDescribeDomainPriceListResponse()
    err = c.Send(request, response)
    return
}
