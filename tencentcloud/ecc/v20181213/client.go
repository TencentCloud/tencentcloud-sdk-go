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

package v20181213

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-12-13"

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


func NewECCRequest() (request *ECCRequest) {
    request = &ECCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecc", APIVersion, "ECC")
    return
}

func NewECCResponse() (response *ECCResponse) {
    response = &ECCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 接口请求域名： ecc.tencentcloudapi.com 
// 纯文本英语作文批改
func (c *Client) ECC(request *ECCRequest) (response *ECCResponse, err error) {
    if request == nil {
        request = NewECCRequest()
    }
    response = NewECCResponse()
    err = c.Send(request, response)
    return
}

func NewEHOCRRequest() (request *EHOCRRequest) {
    request = &EHOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecc", APIVersion, "EHOCR")
    return
}

func NewEHOCRResponse() (response *EHOCRResponse) {
    response = &EHOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// https://ecc.tencentcloudapi.com/?Action=EHOCR
// 作文识别
func (c *Client) EHOCR(request *EHOCRRequest) (response *EHOCRResponse, err error) {
    if request == nil {
        request = NewEHOCRRequest()
    }
    response = NewEHOCRResponse()
    err = c.Send(request, response)
    return
}
