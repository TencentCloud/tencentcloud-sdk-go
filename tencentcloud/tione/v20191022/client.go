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


func NewCreateCodeRepositoryRequest() (request *CreateCodeRepositoryRequest) {
    request = &CreateCodeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "CreateCodeRepository")
    return
}

func NewCreateCodeRepositoryResponse() (response *CreateCodeRepositoryResponse) {
    response = &CreateCodeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建存储库
func (c *Client) CreateCodeRepository(request *CreateCodeRepositoryRequest) (response *CreateCodeRepositoryResponse, err error) {
    if request == nil {
        request = NewCreateCodeRepositoryRequest()
    }
    response = NewCreateCodeRepositoryResponse()
    err = c.Send(request, response)
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

func NewCreateNotebookLifecycleScriptRequest() (request *CreateNotebookLifecycleScriptRequest) {
    request = &CreateNotebookLifecycleScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "CreateNotebookLifecycleScript")
    return
}

func NewCreateNotebookLifecycleScriptResponse() (response *CreateNotebookLifecycleScriptResponse) {
    response = &CreateNotebookLifecycleScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建Notebook生命周期脚本
func (c *Client) CreateNotebookLifecycleScript(request *CreateNotebookLifecycleScriptRequest) (response *CreateNotebookLifecycleScriptResponse, err error) {
    if request == nil {
        request = NewCreateNotebookLifecycleScriptRequest()
    }
    response = NewCreateNotebookLifecycleScriptResponse()
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

func NewDeleteCodeRepositoryRequest() (request *DeleteCodeRepositoryRequest) {
    request = &DeleteCodeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DeleteCodeRepository")
    return
}

func NewDeleteCodeRepositoryResponse() (response *DeleteCodeRepositoryResponse) {
    response = &DeleteCodeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除存储库
func (c *Client) DeleteCodeRepository(request *DeleteCodeRepositoryRequest) (response *DeleteCodeRepositoryResponse, err error) {
    if request == nil {
        request = NewDeleteCodeRepositoryRequest()
    }
    response = NewDeleteCodeRepositoryResponse()
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

func NewDeleteNotebookLifecycleScriptRequest() (request *DeleteNotebookLifecycleScriptRequest) {
    request = &DeleteNotebookLifecycleScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DeleteNotebookLifecycleScript")
    return
}

func NewDeleteNotebookLifecycleScriptResponse() (response *DeleteNotebookLifecycleScriptResponse) {
    response = &DeleteNotebookLifecycleScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除Notebook生命周期脚本
func (c *Client) DeleteNotebookLifecycleScript(request *DeleteNotebookLifecycleScriptRequest) (response *DeleteNotebookLifecycleScriptResponse, err error) {
    if request == nil {
        request = NewDeleteNotebookLifecycleScriptRequest()
    }
    response = NewDeleteNotebookLifecycleScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCodeRepositoriesRequest() (request *DescribeCodeRepositoriesRequest) {
    request = &DescribeCodeRepositoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeCodeRepositories")
    return
}

func NewDescribeCodeRepositoriesResponse() (response *DescribeCodeRepositoriesResponse) {
    response = &DescribeCodeRepositoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询存储库列表
func (c *Client) DescribeCodeRepositories(request *DescribeCodeRepositoriesRequest) (response *DescribeCodeRepositoriesResponse, err error) {
    if request == nil {
        request = NewDescribeCodeRepositoriesRequest()
    }
    response = NewDescribeCodeRepositoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCodeRepositoryRequest() (request *DescribeCodeRepositoryRequest) {
    request = &DescribeCodeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeCodeRepository")
    return
}

func NewDescribeCodeRepositoryResponse() (response *DescribeCodeRepositoryResponse) {
    response = &DescribeCodeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询存储库详情
func (c *Client) DescribeCodeRepository(request *DescribeCodeRepositoryRequest) (response *DescribeCodeRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeCodeRepositoryRequest()
    }
    response = NewDescribeCodeRepositoryResponse()
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

func NewDescribeNotebookLifecycleScriptRequest() (request *DescribeNotebookLifecycleScriptRequest) {
    request = &DescribeNotebookLifecycleScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookLifecycleScript")
    return
}

func NewDescribeNotebookLifecycleScriptResponse() (response *DescribeNotebookLifecycleScriptResponse) {
    response = &DescribeNotebookLifecycleScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看notebook生命周期脚本详情
func (c *Client) DescribeNotebookLifecycleScript(request *DescribeNotebookLifecycleScriptRequest) (response *DescribeNotebookLifecycleScriptResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookLifecycleScriptRequest()
    }
    response = NewDescribeNotebookLifecycleScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookLifecycleScriptsRequest() (request *DescribeNotebookLifecycleScriptsRequest) {
    request = &DescribeNotebookLifecycleScriptsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookLifecycleScripts")
    return
}

func NewDescribeNotebookLifecycleScriptsResponse() (response *DescribeNotebookLifecycleScriptsResponse) {
    response = &DescribeNotebookLifecycleScriptsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看notebook生命周期脚本列表
func (c *Client) DescribeNotebookLifecycleScripts(request *DescribeNotebookLifecycleScriptsRequest) (response *DescribeNotebookLifecycleScriptsResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookLifecycleScriptsRequest()
    }
    response = NewDescribeNotebookLifecycleScriptsResponse()
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

func NewUpdateCodeRepositoryRequest() (request *UpdateCodeRepositoryRequest) {
    request = &UpdateCodeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "UpdateCodeRepository")
    return
}

func NewUpdateCodeRepositoryResponse() (response *UpdateCodeRepositoryResponse) {
    response = &UpdateCodeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新存储库
func (c *Client) UpdateCodeRepository(request *UpdateCodeRepositoryRequest) (response *UpdateCodeRepositoryResponse, err error) {
    if request == nil {
        request = NewUpdateCodeRepositoryRequest()
    }
    response = NewUpdateCodeRepositoryResponse()
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

func NewUpdateNotebookLifecycleScriptRequest() (request *UpdateNotebookLifecycleScriptRequest) {
    request = &UpdateNotebookLifecycleScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "UpdateNotebookLifecycleScript")
    return
}

func NewUpdateNotebookLifecycleScriptResponse() (response *UpdateNotebookLifecycleScriptResponse) {
    response = &UpdateNotebookLifecycleScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新notebook生命周期脚本
func (c *Client) UpdateNotebookLifecycleScript(request *UpdateNotebookLifecycleScriptRequest) (response *UpdateNotebookLifecycleScriptResponse, err error) {
    if request == nil {
        request = NewUpdateNotebookLifecycleScriptRequest()
    }
    response = NewUpdateNotebookLifecycleScriptResponse()
    err = c.Send(request, response)
    return
}
