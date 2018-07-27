// Copyright 1999-2018 Tencent Ltd.
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

package v20180228

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-02-28"

type Client struct {
    common.Client
}

func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithSecretId(credential.SecretId, credential.SecretKey).
        WithProfile(clientProfile)
    return
}


func NewDescribePersonInfoRequest() (request *DescribePersonInfoRequest) {
    request = &DescribePersonInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("youmall", APIVersion, "DescribePersonInfo")
    return
}

func NewDescribePersonInfoResponse() (response *DescribePersonInfoResponse) {
    response = &DescribePersonInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 指定门店获取所有顾客详情列表，包含客户ID、图片、年龄、性别
func (c *Client) DescribePersonInfo(request *DescribePersonInfoRequest) (response *DescribePersonInfoResponse, err error) {
    if request == nil {
        request = NewDescribePersonInfoRequest()
    }
    response = NewDescribePersonInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonVisitInfoRequest() (request *DescribePersonVisitInfoRequest) {
    request = &DescribePersonVisitInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("youmall", APIVersion, "DescribePersonVisitInfo")
    return
}

func NewDescribePersonVisitInfoResponse() (response *DescribePersonVisitInfoResponse) {
    response = &DescribePersonVisitInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取门店指定时间范围内的所有用户到访信息记录，支持的时间范围：过去365天，含当天。
func (c *Client) DescribePersonVisitInfo(request *DescribePersonVisitInfoRequest) (response *DescribePersonVisitInfoResponse, err error) {
    if request == nil {
        request = NewDescribePersonVisitInfoRequest()
    }
    response = NewDescribePersonVisitInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShopHourTrafficInfoRequest() (request *DescribeShopHourTrafficInfoRequest) {
    request = &DescribeShopHourTrafficInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeShopHourTrafficInfo")
    return
}

func NewDescribeShopHourTrafficInfoResponse() (response *DescribeShopHourTrafficInfoResponse) {
    response = &DescribeShopHourTrafficInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按小时提供查询日期范围内门店的每天每小时累计客流人数数据，支持的时间范围：过去365天，含当天。
func (c *Client) DescribeShopHourTrafficInfo(request *DescribeShopHourTrafficInfoRequest) (response *DescribeShopHourTrafficInfoResponse, err error) {
    if request == nil {
        request = NewDescribeShopHourTrafficInfoRequest()
    }
    response = NewDescribeShopHourTrafficInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShopInfoRequest() (request *DescribeShopInfoRequest) {
    request = &DescribeShopInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeShopInfo")
    return
}

func NewDescribeShopInfoResponse() (response *DescribeShopInfoResponse) {
    response = &DescribeShopInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据客户身份标识获取客户下所有的门店信息列表
func (c *Client) DescribeShopInfo(request *DescribeShopInfoRequest) (response *DescribeShopInfoResponse, err error) {
    if request == nil {
        request = NewDescribeShopInfoRequest()
    }
    response = NewDescribeShopInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShopTrafficInfoRequest() (request *DescribeShopTrafficInfoRequest) {
    request = &DescribeShopTrafficInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeShopTrafficInfo")
    return
}

func NewDescribeShopTrafficInfoResponse() (response *DescribeShopTrafficInfoResponse) {
    response = &DescribeShopTrafficInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按天提供查询日期范围内门店的单日累计客流人数，支持的时间范围：过去365天，含当天。
func (c *Client) DescribeShopTrafficInfo(request *DescribeShopTrafficInfoRequest) (response *DescribeShopTrafficInfoResponse, err error) {
    if request == nil {
        request = NewDescribeShopTrafficInfoRequest()
    }
    response = NewDescribeShopTrafficInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneTrafficInfoRequest() (request *DescribeZoneTrafficInfoRequest) {
    request = &DescribeZoneTrafficInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeZoneTrafficInfo")
    return
}

func NewDescribeZoneTrafficInfoResponse() (response *DescribeZoneTrafficInfoResponse) {
    response = &DescribeZoneTrafficInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按天提供查询日期范围内，客户指定门店下的所有区域（优Mall部署时已配置区域）的累计客流人次和平均停留时间。支持的时间范围：过去365天，含当天。
func (c *Client) DescribeZoneTrafficInfo(request *DescribeZoneTrafficInfoRequest) (response *DescribeZoneTrafficInfoResponse, err error) {
    if request == nil {
        request = NewDescribeZoneTrafficInfoRequest()
    }
    response = NewDescribeZoneTrafficInfoResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterCallbackRequest() (request *RegisterCallbackRequest) {
    request = &RegisterCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("youmall", APIVersion, "RegisterCallback")
    return
}

func NewRegisterCallbackResponse() (response *RegisterCallbackResponse) {
    response = &RegisterCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 调用本接口在优Mall中注册自己集团的到店通知回调接口地址，接口协议为HTTP或HTTPS。注册后，若集团有特殊身份（例如老客）到店通知，优Mall后台将主动将到店信息push给该接口
func (c *Client) RegisterCallback(request *RegisterCallbackRequest) (response *RegisterCallbackResponse, err error) {
    if request == nil {
        request = NewRegisterCallbackRequest()
    }
    response = NewRegisterCallbackResponse()
    err = c.Send(request, response)
    return
}
