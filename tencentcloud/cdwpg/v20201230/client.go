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

package v20201230

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-30"

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


func NewCreateInstanceByApiRequest() (request *CreateInstanceByApiRequest) {
    request = &CreateInstanceByApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "CreateInstanceByApi")
    
    
    return
}

func NewCreateInstanceByApiResponse() (response *CreateInstanceByApiResponse) {
    response = &CreateInstanceByApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstanceByApi
// 创建集群
func (c *Client) CreateInstanceByApi(request *CreateInstanceByApiRequest) (response *CreateInstanceByApiResponse, err error) {
    return c.CreateInstanceByApiWithContext(context.Background(), request)
}

// CreateInstanceByApi
// 创建集群
func (c *Client) CreateInstanceByApiWithContext(ctx context.Context, request *CreateInstanceByApiRequest) (response *CreateInstanceByApiResponse, err error) {
    if request == nil {
        request = NewCreateInstanceByApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceByApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceByApiResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleInstancesRequest() (request *DescribeSimpleInstancesRequest) {
    request = &DescribeSimpleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeSimpleInstances")
    
    
    return
}

func NewDescribeSimpleInstancesResponse() (response *DescribeSimpleInstancesResponse) {
    response = &DescribeSimpleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSimpleInstances
// 获取集群实例列表
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DescribeSimpleInstances(request *DescribeSimpleInstancesRequest) (response *DescribeSimpleInstancesResponse, err error) {
    return c.DescribeSimpleInstancesWithContext(context.Background(), request)
}

// DescribeSimpleInstances
// 获取集群实例列表
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DescribeSimpleInstancesWithContext(ctx context.Context, request *DescribeSimpleInstancesRequest) (response *DescribeSimpleInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSimpleInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSimpleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyInstanceByApiRequest() (request *DestroyInstanceByApiRequest) {
    request = &DestroyInstanceByApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DestroyInstanceByApi")
    
    
    return
}

func NewDestroyInstanceByApiResponse() (response *DestroyInstanceByApiResponse) {
    response = &DestroyInstanceByApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyInstanceByApi
// 销毁集群
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DestroyInstanceByApi(request *DestroyInstanceByApiRequest) (response *DestroyInstanceByApiResponse, err error) {
    return c.DestroyInstanceByApiWithContext(context.Background(), request)
}

// DestroyInstanceByApi
// 销毁集群
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DestroyInstanceByApiWithContext(ctx context.Context, request *DestroyInstanceByApiRequest) (response *DestroyInstanceByApiResponse, err error) {
    if request == nil {
        request = NewDestroyInstanceByApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyInstanceByApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyInstanceByApiResponse()
    err = c.Send(request, response)
    return
}
