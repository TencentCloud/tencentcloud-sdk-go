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

package v20200304

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-04"

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


func NewCreateEditingTaskRequest() (request *CreateEditingTaskRequest) {
    request = &CreateEditingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "CreateEditingTask")
    return
}

func NewCreateEditingTaskResponse() (response *CreateEditingTaskResponse) {
    response = &CreateEditingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建智能编辑任务，可以同时选择视频标签识别、分类识别、智能拆条、智能集锦、智能封面和片头片尾识别中的一项或者多项能力。
func (c *Client) CreateEditingTask(request *CreateEditingTaskRequest) (response *CreateEditingTaskResponse, err error) {
    if request == nil {
        request = NewCreateEditingTaskRequest()
    }
    response = NewCreateEditingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEditingTaskResultRequest() (request *DescribeEditingTaskResultRequest) {
    request = &DescribeEditingTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "DescribeEditingTaskResult")
    return
}

func NewDescribeEditingTaskResultResponse() (response *DescribeEditingTaskResultResponse) {
    response = &DescribeEditingTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取智能编辑任务结果。
func (c *Client) DescribeEditingTaskResult(request *DescribeEditingTaskResultRequest) (response *DescribeEditingTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeEditingTaskResultRequest()
    }
    response = NewDescribeEditingTaskResultResponse()
    err = c.Send(request, response)
    return
}
