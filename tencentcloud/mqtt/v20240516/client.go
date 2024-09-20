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

package v20240516

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-05-16"

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


func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeInstance")
    
    
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstance
// 查询实例信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

// DescribeInstance
// 查询实例信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
    request = &DescribeInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeInstanceList")
    
    
    return
}

func NewDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
    response = &DescribeInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceList
// 获取实例列表，Filters参数使用说明如下：
//
// 1. InstanceName, 名称模糊查询
//
// 2. InstanceId，实例ID查询
//
// 3. InstanceType, 实例类型查询，支持多选
//
// 3. InstanceStatus，实例状态查询，支持多选
//
// 
//
// 当使用TagFilters查询时，Filters参数失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    return c.DescribeInstanceListWithContext(context.Background(), request)
}

// DescribeInstanceList
// 获取实例列表，Filters参数使用说明如下：
//
// 1. InstanceName, 名称模糊查询
//
// 2. InstanceId，实例ID查询
//
// 3. InstanceType, 实例类型查询，支持多选
//
// 3. InstanceStatus，实例状态查询，支持多选
//
// 
//
// 当使用TagFilters查询时，Filters参数失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceListWithContext(ctx context.Context, request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceListResponse()
    err = c.Send(request, response)
    return
}
