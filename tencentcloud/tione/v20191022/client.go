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

package v20191022

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-22"

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


func NewCreateNotebookInstanceRequest() (request *CreateNotebookInstanceRequest) {
    request = &CreateNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "CreateNotebookInstance")
    return
}

func NewCreateNotebookInstanceResponse() (response *CreateNotebookInstanceResponse) {
    response = &CreateNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建Notebook实例
func (c *Client) CreateNotebookInstance(request *CreateNotebookInstanceRequest) (response *CreateNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewCreateNotebookInstanceRequest()
    }
    response = NewCreateNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePresignedNotebookInstanceUrlRequest() (request *CreatePresignedNotebookInstanceUrlRequest) {
    request = &CreatePresignedNotebookInstanceUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "CreatePresignedNotebookInstanceUrl")
    return
}

func NewCreatePresignedNotebookInstanceUrlResponse() (response *CreatePresignedNotebookInstanceUrlResponse) {
    response = &CreatePresignedNotebookInstanceUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建Notebook授权Url
func (c *Client) CreatePresignedNotebookInstanceUrl(request *CreatePresignedNotebookInstanceUrlRequest) (response *CreatePresignedNotebookInstanceUrlResponse, err error) {
    if request == nil {
        request = NewCreatePresignedNotebookInstanceUrlRequest()
    }
    response = NewCreatePresignedNotebookInstanceUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTrainingJobRequest() (request *CreateTrainingJobRequest) {
    request = &CreateTrainingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "CreateTrainingJob")
    return
}

func NewCreateTrainingJobResponse() (response *CreateTrainingJobResponse) {
    response = &CreateTrainingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建训练任务
func (c *Client) CreateTrainingJob(request *CreateTrainingJobRequest) (response *CreateTrainingJobResponse, err error) {
    if request == nil {
        request = NewCreateTrainingJobRequest()
    }
    response = NewCreateTrainingJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNotebookInstanceRequest() (request *DeleteNotebookInstanceRequest) {
    request = &DeleteNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DeleteNotebookInstance")
    return
}

func NewDeleteNotebookInstanceResponse() (response *DeleteNotebookInstanceResponse) {
    response = &DeleteNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除notebook实例
func (c *Client) DeleteNotebookInstance(request *DeleteNotebookInstanceRequest) (response *DeleteNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteNotebookInstanceRequest()
    }
    response = NewDeleteNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookInstanceRequest() (request *DescribeNotebookInstanceRequest) {
    request = &DescribeNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookInstance")
    return
}

func NewDescribeNotebookInstanceResponse() (response *DescribeNotebookInstanceResponse) {
    response = &DescribeNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Notebook实例详情
func (c *Client) DescribeNotebookInstance(request *DescribeNotebookInstanceRequest) (response *DescribeNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookInstanceRequest()
    }
    response = NewDescribeNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookInstancesRequest() (request *DescribeNotebookInstancesRequest) {
    request = &DescribeNotebookInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookInstances")
    return
}

func NewDescribeNotebookInstancesResponse() (response *DescribeNotebookInstancesResponse) {
    response = &DescribeNotebookInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Notebook实例列表
func (c *Client) DescribeNotebookInstances(request *DescribeNotebookInstancesRequest) (response *DescribeNotebookInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookInstancesRequest()
    }
    response = NewDescribeNotebookInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingJobRequest() (request *DescribeTrainingJobRequest) {
    request = &DescribeTrainingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingJob")
    return
}

func NewDescribeTrainingJobResponse() (response *DescribeTrainingJobResponse) {
    response = &DescribeTrainingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询训练任务
func (c *Client) DescribeTrainingJob(request *DescribeTrainingJobRequest) (response *DescribeTrainingJobResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingJobRequest()
    }
    response = NewDescribeTrainingJobResponse()
    err = c.Send(request, response)
    return
}

func NewStartNotebookInstanceRequest() (request *StartNotebookInstanceRequest) {
    request = &StartNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "StartNotebookInstance")
    return
}

func NewStartNotebookInstanceResponse() (response *StartNotebookInstanceResponse) {
    response = &StartNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启动Notebook实例
func (c *Client) StartNotebookInstance(request *StartNotebookInstanceRequest) (response *StartNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewStartNotebookInstanceRequest()
    }
    response = NewStartNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStopNotebookInstanceRequest() (request *StopNotebookInstanceRequest) {
    request = &StopNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "StopNotebookInstance")
    return
}

func NewStopNotebookInstanceResponse() (response *StopNotebookInstanceResponse) {
    response = &StopNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止Notebook实例
func (c *Client) StopNotebookInstance(request *StopNotebookInstanceRequest) (response *StopNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewStopNotebookInstanceRequest()
    }
    response = NewStopNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStopTrainingJobRequest() (request *StopTrainingJobRequest) {
    request = &StopTrainingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "StopTrainingJob")
    return
}

func NewStopTrainingJobResponse() (response *StopTrainingJobResponse) {
    response = &StopTrainingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止训练任务
func (c *Client) StopTrainingJob(request *StopTrainingJobRequest) (response *StopTrainingJobResponse, err error) {
    if request == nil {
        request = NewStopTrainingJobRequest()
    }
    response = NewStopTrainingJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNotebookInstanceRequest() (request *UpdateNotebookInstanceRequest) {
    request = &UpdateNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "UpdateNotebookInstance")
    return
}

func NewUpdateNotebookInstanceResponse() (response *UpdateNotebookInstanceResponse) {
    response = &UpdateNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新Notebook实例
func (c *Client) UpdateNotebookInstance(request *UpdateNotebookInstanceRequest) (response *UpdateNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateNotebookInstanceRequest()
    }
    response = NewUpdateNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}
