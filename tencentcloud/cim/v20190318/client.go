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

package v20190318

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-18"

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


func NewDescribeSdkAppidRequest() (request *DescribeSdkAppidRequest) {
    request = &DescribeSdkAppidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cim", APIVersion, "DescribeSdkAppid")
    return
}

func NewDescribeSdkAppidResponse() (response *DescribeSdkAppidResponse) {
    response = &DescribeSdkAppidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取云通信IM中腾讯云账号对应的sdkappid
func (c *Client) DescribeSdkAppid(request *DescribeSdkAppidRequest) (response *DescribeSdkAppidResponse, err error) {
    if request == nil {
        request = NewDescribeSdkAppidRequest()
    }
    response = NewDescribeSdkAppidResponse()
    err = c.Send(request, response)
    return
}
