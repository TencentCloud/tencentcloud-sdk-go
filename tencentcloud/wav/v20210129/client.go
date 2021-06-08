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

package v20210129

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-01-29"

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


func NewQueryActivityLiveCodeListRequest() (request *QueryActivityLiveCodeListRequest) {
    request = &QueryActivityLiveCodeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("wav", APIVersion, "QueryActivityLiveCodeList")
    return
}

func NewQueryActivityLiveCodeListResponse() (response *QueryActivityLiveCodeListResponse) {
    response = &QueryActivityLiveCodeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据游标拉取活动活码列表信息
func (c *Client) QueryActivityLiveCodeList(request *QueryActivityLiveCodeListRequest) (response *QueryActivityLiveCodeListResponse, err error) {
    if request == nil {
        request = NewQueryActivityLiveCodeListRequest()
    }
    response = NewQueryActivityLiveCodeListResponse()
    err = c.Send(request, response)
    return
}
