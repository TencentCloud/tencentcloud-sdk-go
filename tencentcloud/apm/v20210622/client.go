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

package v20210622

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-06-22"

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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateApmInstanceRequest() (request *CreateApmInstanceRequest) {
    request = &CreateApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apm", APIVersion, "CreateApmInstance")
    
    
    return
}

func NewCreateApmInstanceResponse() (response *CreateApmInstanceResponse) {
    response = &CreateApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApmInstance
// 业务购买APM实例，调用该接口创建
func (c *Client) CreateApmInstance(request *CreateApmInstanceRequest) (response *CreateApmInstanceResponse, err error) {
    if request == nil {
        request = NewCreateApmInstanceRequest()
    }
    
    response = NewCreateApmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmAgentRequest() (request *DescribeApmAgentRequest) {
    request = &DescribeApmAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmAgent")
    
    
    return
}

func NewDescribeApmAgentResponse() (response *DescribeApmAgentResponse) {
    response = &DescribeApmAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApmAgent
// 获取Apm Agent信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApmAgent(request *DescribeApmAgentRequest) (response *DescribeApmAgentResponse, err error) {
    if request == nil {
        request = NewDescribeApmAgentRequest()
    }
    
    response = NewDescribeApmAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmInstancesRequest() (request *DescribeApmInstancesRequest) {
    request = &DescribeApmInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmInstances")
    
    
    return
}

func NewDescribeApmInstancesResponse() (response *DescribeApmInstancesResponse) {
    response = &DescribeApmInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApmInstances
// APM实例列表拉取
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApmInstances(request *DescribeApmInstancesRequest) (response *DescribeApmInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeApmInstancesRequest()
    }
    
    response = NewDescribeApmInstancesResponse()
    err = c.Send(request, response)
    return
}
