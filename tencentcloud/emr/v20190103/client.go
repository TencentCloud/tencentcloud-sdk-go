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

package v20190103

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-01-03"

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


func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("emr", APIVersion, "CreateInstance")
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建EMR实例
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("emr", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询EMR实例
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateInstanceRequest() (request *InquiryPriceCreateInstanceRequest) {
    request = &InquiryPriceCreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceCreateInstance")
    return
}

func NewInquiryPriceCreateInstanceResponse() (response *InquiryPriceCreateInstanceResponse) {
    response = &InquiryPriceCreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建实例询价
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateInstanceRequest()
    }
    response = NewInquiryPriceCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceScaleOutInstanceRequest() (request *InquiryPriceScaleOutInstanceRequest) {
    request = &InquiryPriceScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceScaleOutInstance")
    return
}

func NewInquiryPriceScaleOutInstanceResponse() (response *InquiryPriceScaleOutInstanceResponse) {
    response = &InquiryPriceScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 扩容询价. 当扩容时候，请通过该接口查询价格。
func (c *Client) InquiryPriceScaleOutInstance(request *InquiryPriceScaleOutInstanceRequest) (response *InquiryPriceScaleOutInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceScaleOutInstanceRequest()
    }
    response = NewInquiryPriceScaleOutInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutInstanceRequest() (request *ScaleOutInstanceRequest) {
    request = &ScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("emr", APIVersion, "ScaleOutInstance")
    return
}

func NewScaleOutInstanceResponse() (response *ScaleOutInstanceResponse) {
    response = &ScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 实例扩容
func (c *Client) ScaleOutInstance(request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    if request == nil {
        request = NewScaleOutInstanceRequest()
    }
    response = NewScaleOutInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstanceRequest() (request *TerminateInstanceRequest) {
    request = &TerminateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("emr", APIVersion, "TerminateInstance")
    return
}

func NewTerminateInstanceResponse() (response *TerminateInstanceResponse) {
    response = &TerminateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 销毁EMR实例
func (c *Client) TerminateInstance(request *TerminateInstanceRequest) (response *TerminateInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateInstanceRequest()
    }
    response = NewTerminateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateTasksRequest() (request *TerminateTasksRequest) {
    request = &TerminateTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("emr", APIVersion, "TerminateTasks")
    return
}

func NewTerminateTasksResponse() (response *TerminateTasksResponse) {
    response = &TerminateTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 缩容Task节点
func (c *Client) TerminateTasks(request *TerminateTasksRequest) (response *TerminateTasksResponse, err error) {
    if request == nil {
        request = NewTerminateTasksRequest()
    }
    response = NewTerminateTasksResponse()
    err = c.Send(request, response)
    return
}
