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

package v20220519

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-05-19"

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


func NewCreateDataRepositoryTaskRequest() (request *CreateDataRepositoryTaskRequest) {
    request = &CreateDataRepositoryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "CreateDataRepositoryTask")
    
    
    return
}

func NewCreateDataRepositoryTaskResponse() (response *CreateDataRepositoryTaskResponse) {
    response = &CreateDataRepositoryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDataRepositoryTask
// 创建数据流通任务,包括从将文件系统的数据上传到存储桶下, 以及从存储桶下载到文件系统里。
func (c *Client) CreateDataRepositoryTask(request *CreateDataRepositoryTaskRequest) (response *CreateDataRepositoryTaskResponse, err error) {
    return c.CreateDataRepositoryTaskWithContext(context.Background(), request)
}

// CreateDataRepositoryTask
// 创建数据流通任务,包括从将文件系统的数据上传到存储桶下, 以及从存储桶下载到文件系统里。
func (c *Client) CreateDataRepositoryTaskWithContext(ctx context.Context, request *CreateDataRepositoryTaskRequest) (response *CreateDataRepositoryTaskResponse, err error) {
    if request == nil {
        request = NewCreateDataRepositoryTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataRepositoryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataRepositoryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataRepositoryTaskStatusRequest() (request *DescribeDataRepositoryTaskStatusRequest) {
    request = &DescribeDataRepositoryTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeDataRepositoryTaskStatus")
    
    
    return
}

func NewDescribeDataRepositoryTaskStatusResponse() (response *DescribeDataRepositoryTaskStatusResponse) {
    response = &DescribeDataRepositoryTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataRepositoryTaskStatus
// 获取数据流通任务实时状态，用作客户端控制
func (c *Client) DescribeDataRepositoryTaskStatus(request *DescribeDataRepositoryTaskStatusRequest) (response *DescribeDataRepositoryTaskStatusResponse, err error) {
    return c.DescribeDataRepositoryTaskStatusWithContext(context.Background(), request)
}

// DescribeDataRepositoryTaskStatus
// 获取数据流通任务实时状态，用作客户端控制
func (c *Client) DescribeDataRepositoryTaskStatusWithContext(ctx context.Context, request *DescribeDataRepositoryTaskStatusRequest) (response *DescribeDataRepositoryTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDataRepositoryTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataRepositoryTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataRepositoryTaskStatusResponse()
    err = c.Send(request, response)
    return
}
