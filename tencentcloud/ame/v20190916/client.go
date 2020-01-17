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

package v20190916

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-16"

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


func NewDescribeItemsRequest() (request *DescribeItemsRequest) {
    request = &DescribeItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ame", APIVersion, "DescribeItems")
    return
}

func NewDescribeItemsResponse() (response *DescribeItemsResponse) {
    response = &DescribeItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 分类内容下歌曲列表获取，根据CategoryID或CategoryCode
func (c *Client) DescribeItems(request *DescribeItemsRequest) (response *DescribeItemsResponse, err error) {
    if request == nil {
        request = NewDescribeItemsRequest()
    }
    response = NewDescribeItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLyricRequest() (request *DescribeLyricRequest) {
    request = &DescribeLyricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ame", APIVersion, "DescribeLyric")
    return
}

func NewDescribeLyricResponse() (response *DescribeLyricResponse) {
    response = &DescribeLyricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据接口的模式及歌曲ID来取得歌词信息。
func (c *Client) DescribeLyric(request *DescribeLyricRequest) (response *DescribeLyricResponse, err error) {
    if request == nil {
        request = NewDescribeLyricRequest()
    }
    response = NewDescribeLyricResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMusicRequest() (request *DescribeMusicRequest) {
    request = &DescribeMusicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ame", APIVersion, "DescribeMusic")
    return
}

func NewDescribeMusicResponse() (response *DescribeMusicResponse) {
    response = &DescribeMusicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据接口的模式及歌曲ID来取得对应权限的歌曲播放地址等信息。
func (c *Client) DescribeMusic(request *DescribeMusicRequest) (response *DescribeMusicResponse, err error) {
    if request == nil {
        request = NewDescribeMusicRequest()
    }
    response = NewDescribeMusicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStationsRequest() (request *DescribeStationsRequest) {
    request = &DescribeStationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ame", APIVersion, "DescribeStations")
    return
}

func NewDescribeStationsResponse() (response *DescribeStationsResponse) {
    response = &DescribeStationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取素材库列表时使用
func (c *Client) DescribeStations(request *DescribeStationsRequest) (response *DescribeStationsResponse, err error) {
    if request == nil {
        request = NewDescribeStationsRequest()
    }
    response = NewDescribeStationsResponse()
    err = c.Send(request, response)
    return
}

func NewReportDataRequest() (request *ReportDataRequest) {
    request = &ReportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ame", APIVersion, "ReportData")
    return
}

func NewReportDataResponse() (response *ReportDataResponse) {
    response = &ReportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 客户上报用户数据功能，为了更好的为用户提供优质服务
func (c *Client) ReportData(request *ReportDataRequest) (response *ReportDataResponse, err error) {
    if request == nil {
        request = NewReportDataRequest()
    }
    response = NewReportDataResponse()
    err = c.Send(request, response)
    return
}
