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

package v20210820

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-08-20"

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


func NewDescribeProjectRequest() (request *DescribeProjectRequest) {
    request = &DescribeProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeProject")
    
    
    return
}

func NewDescribeProjectResponse() (response *DescribeProjectResponse) {
    response = &DescribeProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProject
// 获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    return c.DescribeProjectWithContext(context.Background(), request)
}

// DescribeProject
// 获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeProjectWithContext(ctx context.Context, request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    if request == nil {
        request = NewDescribeProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRelatedInstancesRequest() (request *DescribeRelatedInstancesRequest) {
    request = &DescribeRelatedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRelatedInstances")
    
    
    return
}

func NewDescribeRelatedInstancesResponse() (response *DescribeRelatedInstancesResponse) {
    response = &DescribeRelatedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRelatedInstances
// 查询任务实例的关联实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRelatedInstances(request *DescribeRelatedInstancesRequest) (response *DescribeRelatedInstancesResponse, err error) {
    return c.DescribeRelatedInstancesWithContext(context.Background(), request)
}

// DescribeRelatedInstances
// 查询任务实例的关联实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRelatedInstancesWithContext(ctx context.Context, request *DescribeRelatedInstancesRequest) (response *DescribeRelatedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeRelatedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelatedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRelatedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInstancesRequest() (request *DescribeTaskInstancesRequest) {
    request = &DescribeTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskInstances")
    
    
    return
}

func NewDescribeTaskInstancesResponse() (response *DescribeTaskInstancesResponse) {
    response = &DescribeTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskInstances
// 查询任务实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskInstances(request *DescribeTaskInstancesRequest) (response *DescribeTaskInstancesResponse, err error) {
    return c.DescribeTaskInstancesWithContext(context.Background(), request)
}

// DescribeTaskInstances
// 查询任务实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskInstancesWithContext(ctx context.Context, request *DescribeTaskInstancesRequest) (response *DescribeTaskInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInstancesResponse()
    err = c.Send(request, response)
    return
}
