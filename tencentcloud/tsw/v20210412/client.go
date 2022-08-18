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

package v20210412

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-04-12"

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


func NewDescribeComponentAlertObjectRequest() (request *DescribeComponentAlertObjectRequest) {
    request = &DescribeComponentAlertObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsw", APIVersion, "DescribeComponentAlertObject")
    
    
    return
}

func NewDescribeComponentAlertObjectResponse() (response *DescribeComponentAlertObjectResponse) {
    response = &DescribeComponentAlertObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComponentAlertObject
// 获取告警对象-组件告警
func (c *Client) DescribeComponentAlertObject(request *DescribeComponentAlertObjectRequest) (response *DescribeComponentAlertObjectResponse, err error) {
    return c.DescribeComponentAlertObjectWithContext(context.Background(), request)
}

// DescribeComponentAlertObject
// 获取告警对象-组件告警
func (c *Client) DescribeComponentAlertObjectWithContext(ctx context.Context, request *DescribeComponentAlertObjectRequest) (response *DescribeComponentAlertObjectResponse, err error) {
    if request == nil {
        request = NewDescribeComponentAlertObjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComponentAlertObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComponentAlertObjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceAlertObjectRequest() (request *DescribeServiceAlertObjectRequest) {
    request = &DescribeServiceAlertObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsw", APIVersion, "DescribeServiceAlertObject")
    
    
    return
}

func NewDescribeServiceAlertObjectResponse() (response *DescribeServiceAlertObjectResponse) {
    response = &DescribeServiceAlertObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceAlertObject
// 获取告警对象-服务告警表
func (c *Client) DescribeServiceAlertObject(request *DescribeServiceAlertObjectRequest) (response *DescribeServiceAlertObjectResponse, err error) {
    return c.DescribeServiceAlertObjectWithContext(context.Background(), request)
}

// DescribeServiceAlertObject
// 获取告警对象-服务告警表
func (c *Client) DescribeServiceAlertObjectWithContext(ctx context.Context, request *DescribeServiceAlertObjectRequest) (response *DescribeServiceAlertObjectResponse, err error) {
    if request == nil {
        request = NewDescribeServiceAlertObjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceAlertObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceAlertObjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenRequest() (request *DescribeTokenRequest) {
    request = &DescribeTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsw", APIVersion, "DescribeToken")
    
    
    return
}

func NewDescribeTokenResponse() (response *DescribeTokenResponse) {
    response = &DescribeTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeToken
// 查询token
func (c *Client) DescribeToken(request *DescribeTokenRequest) (response *DescribeTokenResponse, err error) {
    return c.DescribeTokenWithContext(context.Background(), request)
}

// DescribeToken
// 查询token
func (c *Client) DescribeTokenWithContext(ctx context.Context, request *DescribeTokenRequest) (response *DescribeTokenResponse, err error) {
    if request == nil {
        request = NewDescribeTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenResponse()
    err = c.Send(request, response)
    return
}
