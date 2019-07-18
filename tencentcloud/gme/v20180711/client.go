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

package v20180711

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-11"

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


func NewDescribeFilterResultRequest() (request *DescribeFilterResultRequest) {
    request = &DescribeFilterResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeFilterResult")
    return
}

func NewDescribeFilterResultResponse() (response *DescribeFilterResultResponse) {
    response = &DescribeFilterResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据应用ID和文件ID查询识别结果
func (c *Client) DescribeFilterResult(request *DescribeFilterResultRequest) (response *DescribeFilterResultResponse, err error) {
    if request == nil {
        request = NewDescribeFilterResultRequest()
    }
    response = NewDescribeFilterResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFilterResultListRequest() (request *DescribeFilterResultListRequest) {
    request = &DescribeFilterResultListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeFilterResultList")
    return
}

func NewDescribeFilterResultListResponse() (response *DescribeFilterResultListResponse) {
    response = &DescribeFilterResultListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据日期查询识别结果列表
func (c *Client) DescribeFilterResultList(request *DescribeFilterResultListRequest) (response *DescribeFilterResultListResponse, err error) {
    if request == nil {
        request = NewDescribeFilterResultListRequest()
    }
    response = NewDescribeFilterResultListResponse()
    err = c.Send(request, response)
    return
}

func NewVoiceFilterRequest() (request *VoiceFilterRequest) {
    request = &VoiceFilterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "VoiceFilter")
    return
}

func NewVoiceFilterResponse() (response *VoiceFilterResponse) {
    response = &VoiceFilterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于识别涉黄、涉政等违规音频，成功会回调配置在应用的回调地址。回调示例如下：
// {"BizId":0,"FileId":"test_file_id","FileName":"test_file_name","FileUrl":"test_file_url","OpenId":"test_open_id","TimeStamp":"0000-00-00 00:00:00","Data":[{"Type":1,"Word":"xx"}]}
// Type表示过滤类型，1：政治，2：色情，3：谩骂
func (c *Client) VoiceFilter(request *VoiceFilterRequest) (response *VoiceFilterResponse, err error) {
    if request == nil {
        request = NewVoiceFilterRequest()
    }
    response = NewVoiceFilterResponse()
    err = c.Send(request, response)
    return
}
