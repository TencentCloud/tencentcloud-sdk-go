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

package v20190328

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-28"

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


func NewDescribeBRIRequest() (request *DescribeBRIRequest) {
    request = &DescribeBRIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bri", APIVersion, "DescribeBRI")
    return
}

func NewDescribeBRIResponse() (response *DescribeBRIResponse) {
    response = &DescribeBRIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 输入业务名 (bri_num, bri_dev, bri_ip, bri_apk, bri_url 五种之一)  及其 相应字段, 获取业务风险分数和标签。
// 
// 当业务名为bri_num时，必须填PhoneNumber字段.
// 
// 当业务名为bri_dev时, 必须填Imei字段.
// 
// 当业务名为bri_ip时，必须填Ip字段.
// 
// 当业务名为bri_apk时，必须填 (PackageName,CertMd5,FileSize) 三个字段 或者 FileMd5一个字段.
// 
// 当业务名为bri_url时，必须填Url字段.
func (c *Client) DescribeBRI(request *DescribeBRIRequest) (response *DescribeBRIResponse, err error) {
    if request == nil {
        request = NewDescribeBRIRequest()
    }
    response = NewDescribeBRIResponse()
    err = c.Send(request, response)
    return
}
