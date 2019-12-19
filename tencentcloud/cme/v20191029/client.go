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

package v20191029

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-29"

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


func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "CreateProject")
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建云剪的编辑项目，支持创建视频剪辑及直播剪辑两大类项目。
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoginStatusRequest() (request *DeleteLoginStatusRequest) {
    request = &DeleteLoginStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DeleteLoginStatus")
    return
}

func NewDeleteLoginStatusResponse() (response *DeleteLoginStatusResponse) {
    response = &DeleteLoginStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户登录态，使用户登出云剪平台。
func (c *Client) DeleteLoginStatus(request *DeleteLoginStatusRequest) (response *DeleteLoginStatusResponse, err error) {
    if request == nil {
        request = NewDeleteLoginStatusRequest()
    }
    response = NewDeleteLoginStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DeleteProject")
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除云剪编辑项目。
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginStatusRequest() (request *DescribeLoginStatusRequest) {
    request = &DescribeLoginStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeLoginStatus")
    return
}

func NewDescribeLoginStatusResponse() (response *DescribeLoginStatusResponse) {
    response = &DescribeLoginStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指定用户的登录态。
func (c *Client) DescribeLoginStatus(request *DescribeLoginStatusRequest) (response *DescribeLoginStatusResponse, err error) {
    if request == nil {
        request = NewDescribeLoginStatusRequest()
    }
    response = NewDescribeLoginStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeProjects")
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 支持根据多种条件过滤出项目列表。
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeTaskDetail")
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取任务详情信息，包含下面几个部分：
// <li>任务基础信息：包括任务状态、错误信息、创建时间等；</li>
// <li>导出项目输出信息：包括输出的素材 Id 等。</li>
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeTasks")
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 支持各种条件筛选，返回对应的任务基础信息列表。
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewExportVideoEditProjectRequest() (request *ExportVideoEditProjectRequest) {
    request = &ExportVideoEditProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "ExportVideoEditProject")
    return
}

func NewExportVideoEditProjectResponse() (response *ExportVideoEditProjectResponse) {
    response = &ExportVideoEditProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导出视频编辑项目，支持指定输出的模板。
func (c *Client) ExportVideoEditProject(request *ExportVideoEditProjectRequest) (response *ExportVideoEditProjectResponse, err error) {
    if request == nil {
        request = NewExportVideoEditProjectRequest()
    }
    response = NewExportVideoEditProjectResponse()
    err = c.Send(request, response)
    return
}

func NewImportMediaToProjectRequest() (request *ImportMediaToProjectRequest) {
    request = &ImportMediaToProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "ImportMediaToProject")
    return
}

func NewImportMediaToProjectResponse() (response *ImportMediaToProjectResponse) {
    response = &ImportMediaToProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将云点播中的媒资添加到素材库中，提供给后续的视频编辑。
func (c *Client) ImportMediaToProject(request *ImportMediaToProjectRequest) (response *ImportMediaToProjectResponse, err error) {
    if request == nil {
        request = NewImportMediaToProjectRequest()
    }
    response = NewImportMediaToProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
    request = &ModifyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "ModifyProject")
    return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
    response = &ModifyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改云剪编辑项目的信息。
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}
