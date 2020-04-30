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


func NewAddTeamMemberRequest() (request *AddTeamMemberRequest) {
    request = &AddTeamMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "AddTeamMember")
    return
}

func NewAddTeamMemberResponse() (response *AddTeamMemberResponse) {
    response = &AddTeamMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 向一个团队中团队成员，并且指定成员的角色。
func (c *Client) AddTeamMember(request *AddTeamMemberRequest) (response *AddTeamMemberResponse, err error) {
    if request == nil {
        request = NewAddTeamMemberRequest()
    }
    response = NewAddTeamMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClassRequest() (request *CreateClassRequest) {
    request = &CreateClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "CreateClass")
    return
}

func NewCreateClassResponse() (response *CreateClassResponse) {
    response = &CreateClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增分类，用于管理素材。
// <li>分类层数不能超过10；</li>
// <li>子分类数不能超过10。</li>
func (c *Client) CreateClass(request *CreateClassRequest) (response *CreateClassResponse, err error) {
    if request == nil {
        request = NewCreateClassRequest()
    }
    response = NewCreateClassResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLinkRequest() (request *CreateLinkRequest) {
    request = &CreateLinkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "CreateLink")
    return
}

func NewCreateLinkResponse() (response *CreateLinkResponse) {
    response = &CreateLinkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  创建素材链接或分类路径链接，将源资源信息链接到目标。
func (c *Client) CreateLink(request *CreateLinkRequest) (response *CreateLinkResponse, err error) {
    if request == nil {
        request = NewCreateLinkRequest()
    }
    response = NewCreateLinkResponse()
    err = c.Send(request, response)
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

func NewCreateTeamRequest() (request *CreateTeamRequest) {
    request = &CreateTeamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "CreateTeam")
    return
}

func NewCreateTeamResponse() (response *CreateTeamResponse) {
    response = &CreateTeamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建一个团队。
func (c *Client) CreateTeam(request *CreateTeamRequest) (response *CreateTeamResponse, err error) {
    if request == nil {
        request = NewCreateTeamRequest()
    }
    response = NewCreateTeamResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClassRequest() (request *DeleteClassRequest) {
    request = &DeleteClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DeleteClass")
    return
}

func NewDeleteClassResponse() (response *DeleteClassResponse) {
    response = &DeleteClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除分类信息，删除时检验下述限制：
// <li>分类路径必须存在；</li>
// <li>分类下没有绑定素材。</li>
func (c *Client) DeleteClass(request *DeleteClassRequest) (response *DeleteClassResponse, err error) {
    if request == nil {
        request = NewDeleteClassRequest()
    }
    response = NewDeleteClassResponse()
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

func NewDeleteMaterialRequest() (request *DeleteMaterialRequest) {
    request = &DeleteMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DeleteMaterial")
    return
}

func NewDeleteMaterialResponse() (response *DeleteMaterialResponse) {
    response = &DeleteMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据素材 Id 删除素材。
func (c *Client) DeleteMaterial(request *DeleteMaterialRequest) (response *DeleteMaterialResponse, err error) {
    if request == nil {
        request = NewDeleteMaterialRequest()
    }
    response = NewDeleteMaterialResponse()
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

func NewDeleteTeamRequest() (request *DeleteTeamRequest) {
    request = &DeleteTeamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DeleteTeam")
    return
}

func NewDeleteTeamResponse() (response *DeleteTeamResponse) {
    response = &DeleteTeamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除一个团队。
// <li>要删除的团队必须没有归属的素材；</li>
// <li>要删除的团队必须没有归属的分类。</li>
func (c *Client) DeleteTeam(request *DeleteTeamRequest) (response *DeleteTeamResponse, err error) {
    if request == nil {
        request = NewDeleteTeamRequest()
    }
    response = NewDeleteTeamResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTeamMembersRequest() (request *DeleteTeamMembersRequest) {
    request = &DeleteTeamMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DeleteTeamMembers")
    return
}

func NewDeleteTeamMembersResponse() (response *DeleteTeamMembersResponse) {
    response = &DeleteTeamMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将团队成员从团队中删除，默认只有 Owner 及管理员才有此权限。
func (c *Client) DeleteTeamMembers(request *DeleteTeamMembersRequest) (response *DeleteTeamMembersResponse, err error) {
    if request == nil {
        request = NewDeleteTeamMembersRequest()
    }
    response = NewDeleteTeamMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassRequest() (request *DescribeClassRequest) {
    request = &DescribeClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeClass")
    return
}

func NewDescribeClassResponse() (response *DescribeClassResponse) {
    response = &DescribeClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定归属者下所有的分类信息。
func (c *Client) DescribeClass(request *DescribeClassRequest) (response *DescribeClassResponse, err error) {
    if request == nil {
        request = NewDescribeClassRequest()
    }
    response = NewDescribeClassResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJoinTeamsRequest() (request *DescribeJoinTeamsRequest) {
    request = &DescribeJoinTeamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeJoinTeams")
    return
}

func NewDescribeJoinTeamsResponse() (response *DescribeJoinTeamsResponse) {
    response = &DescribeJoinTeamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定的团队成员所加入的团队列表。
func (c *Client) DescribeJoinTeams(request *DescribeJoinTeamsRequest) (response *DescribeJoinTeamsResponse, err error) {
    if request == nil {
        request = NewDescribeJoinTeamsRequest()
    }
    response = NewDescribeJoinTeamsResponse()
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

func NewDescribeMaterialsRequest() (request *DescribeMaterialsRequest) {
    request = &DescribeMaterialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeMaterials")
    return
}

func NewDescribeMaterialsResponse() (response *DescribeMaterialsResponse) {
    response = &DescribeMaterialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据素材 Id 批量获取素材详情。
func (c *Client) DescribeMaterials(request *DescribeMaterialsRequest) (response *DescribeMaterialsResponse, err error) {
    if request == nil {
        request = NewDescribeMaterialsRequest()
    }
    response = NewDescribeMaterialsResponse()
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

func NewDescribeResourceAuthorizationRequest() (request *DescribeResourceAuthorizationRequest) {
    request = &DescribeResourceAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeResourceAuthorization")
    return
}

func NewDescribeResourceAuthorizationResponse() (response *DescribeResourceAuthorizationResponse) {
    response = &DescribeResourceAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指定资源的授权列表。
func (c *Client) DescribeResourceAuthorization(request *DescribeResourceAuthorizationRequest) (response *DescribeResourceAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeResourceAuthorizationRequest()
    }
    response = NewDescribeResourceAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSharedSpaceRequest() (request *DescribeSharedSpaceRequest) {
    request = &DescribeSharedSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeSharedSpace")
    return
}

func NewDescribeSharedSpaceResponse() (response *DescribeSharedSpaceResponse) {
    response = &DescribeSharedSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取共享空间。当实体A对实体B授权某资源以后，实体B的共享空间就会增加实体A。
func (c *Client) DescribeSharedSpace(request *DescribeSharedSpaceRequest) (response *DescribeSharedSpaceResponse, err error) {
    if request == nil {
        request = NewDescribeSharedSpaceRequest()
    }
    response = NewDescribeSharedSpaceResponse()
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

// 获取任务列表，支持条件筛选，返回对应的任务基础信息列表。
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTeamMembersRequest() (request *DescribeTeamMembersRequest) {
    request = &DescribeTeamMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeTeamMembers")
    return
}

func NewDescribeTeamMembersResponse() (response *DescribeTeamMembersResponse) {
    response = &DescribeTeamMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定成员 ID 的信息，同时支持拉取所有团队成员信息。
func (c *Client) DescribeTeamMembers(request *DescribeTeamMembersRequest) (response *DescribeTeamMembersResponse, err error) {
    if request == nil {
        request = NewDescribeTeamMembersRequest()
    }
    response = NewDescribeTeamMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTeamsRequest() (request *DescribeTeamsRequest) {
    request = &DescribeTeamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "DescribeTeams")
    return
}

func NewDescribeTeamsResponse() (response *DescribeTeamsResponse) {
    response = &DescribeTeamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定团队的信息。
func (c *Client) DescribeTeams(request *DescribeTeamsRequest) (response *DescribeTeamsResponse, err error) {
    if request == nil {
        request = NewDescribeTeamsRequest()
    }
    response = NewDescribeTeamsResponse()
    err = c.Send(request, response)
    return
}

func NewExportVideoByEditorTrackDataRequest() (request *ExportVideoByEditorTrackDataRequest) {
    request = &ExportVideoByEditorTrackDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "ExportVideoByEditorTrackData")
    return
}

func NewExportVideoByEditorTrackDataResponse() (response *ExportVideoByEditorTrackDataResponse) {
    response = &ExportVideoByEditorTrackDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用在线编辑轨道数据直接导出视频。
func (c *Client) ExportVideoByEditorTrackData(request *ExportVideoByEditorTrackDataRequest) (response *ExportVideoByEditorTrackDataResponse, err error) {
    if request == nil {
        request = NewExportVideoByEditorTrackDataRequest()
    }
    response = NewExportVideoByEditorTrackDataResponse()
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

func NewFlattenListMediaRequest() (request *FlattenListMediaRequest) {
    request = &FlattenListMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "FlattenListMedia")
    return
}

func NewFlattenListMediaResponse() (response *FlattenListMediaResponse) {
    response = &FlattenListMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 平铺分类路径下及其子分类下的所有素材。
func (c *Client) FlattenListMedia(request *FlattenListMediaRequest) (response *FlattenListMediaResponse, err error) {
    if request == nil {
        request = NewFlattenListMediaRequest()
    }
    response = NewFlattenListMediaResponse()
    err = c.Send(request, response)
    return
}

func NewGrantResourceAuthorizationRequest() (request *GrantResourceAuthorizationRequest) {
    request = &GrantResourceAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "GrantResourceAuthorization")
    return
}

func NewGrantResourceAuthorizationResponse() (response *GrantResourceAuthorizationResponse) {
    response = &GrantResourceAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 资源所属实体对目标实体授予目标资源的相应权限。
func (c *Client) GrantResourceAuthorization(request *GrantResourceAuthorizationRequest) (response *GrantResourceAuthorizationResponse, err error) {
    if request == nil {
        request = NewGrantResourceAuthorizationRequest()
    }
    response = NewGrantResourceAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewImportMaterialRequest() (request *ImportMaterialRequest) {
    request = &ImportMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "ImportMaterial")
    return
}

func NewImportMaterialResponse() (response *ImportMaterialResponse) {
    response = &ImportMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将云点播媒资文件导入到云剪素材库。
func (c *Client) ImportMaterial(request *ImportMaterialRequest) (response *ImportMaterialResponse, err error) {
    if request == nil {
        request = NewImportMaterialRequest()
    }
    response = NewImportMaterialResponse()
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

// 将云点播中的媒资添加到素材库中，供后续视频编辑使用。
func (c *Client) ImportMediaToProject(request *ImportMediaToProjectRequest) (response *ImportMediaToProjectResponse, err error) {
    if request == nil {
        request = NewImportMediaToProjectRequest()
    }
    response = NewImportMediaToProjectResponse()
    err = c.Send(request, response)
    return
}

func NewListMediaRequest() (request *ListMediaRequest) {
    request = &ListMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "ListMedia")
    return
}

func NewListMediaResponse() (response *ListMediaResponse) {
    response = &ListMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  浏览当前分类路径下的资源，包括素材和子分类。
func (c *Client) ListMedia(request *ListMediaRequest) (response *ListMediaResponse, err error) {
    if request == nil {
        request = NewListMediaRequest()
    }
    response = NewListMediaResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaterialRequest() (request *ModifyMaterialRequest) {
    request = &ModifyMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "ModifyMaterial")
    return
}

func NewModifyMaterialResponse() (response *ModifyMaterialResponse) {
    response = &ModifyMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改素材信息，支持修改素材名称、分类路径、标签等信息。
func (c *Client) ModifyMaterial(request *ModifyMaterialRequest) (response *ModifyMaterialResponse, err error) {
    if request == nil {
        request = NewModifyMaterialRequest()
    }
    response = NewModifyMaterialResponse()
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

func NewModifyTeamRequest() (request *ModifyTeamRequest) {
    request = &ModifyTeamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "ModifyTeam")
    return
}

func NewModifyTeamResponse() (response *ModifyTeamResponse) {
    response = &ModifyTeamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改团队信息，目前支持修改的操作有：
// <li>修改团队名称。</li>
func (c *Client) ModifyTeam(request *ModifyTeamRequest) (response *ModifyTeamResponse, err error) {
    if request == nil {
        request = NewModifyTeamRequest()
    }
    response = NewModifyTeamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTeamMemberRequest() (request *ModifyTeamMemberRequest) {
    request = &ModifyTeamMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "ModifyTeamMember")
    return
}

func NewModifyTeamMemberResponse() (response *ModifyTeamMemberResponse) {
    response = &ModifyTeamMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改团队成员信息，包括成员备注、角色等。
func (c *Client) ModifyTeamMember(request *ModifyTeamMemberRequest) (response *ModifyTeamMemberResponse, err error) {
    if request == nil {
        request = NewModifyTeamMemberRequest()
    }
    response = NewModifyTeamMemberResponse()
    err = c.Send(request, response)
    return
}

func NewMoveClassRequest() (request *MoveClassRequest) {
    request = &MoveClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "MoveClass")
    return
}

func NewMoveClassResponse() (response *MoveClassResponse) {
    response = &MoveClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 移动某一个分类到另外一个分类下，也可用于分类重命名。
// <li>如果 SourceClassPath = /素材/视频/NBA，DestinationClassPath = /素材/视频/篮球，当 DestinationClassPath 不存在时候，操作结果为重命名 ClassPath，如果 DestinationClassPath 存在时候，操作结果为产生新目录 /素材/视频/篮球/NBA。</li>
func (c *Client) MoveClass(request *MoveClassRequest) (response *MoveClassResponse, err error) {
    if request == nil {
        request = NewMoveClassRequest()
    }
    response = NewMoveClassResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeResourceAuthorizationRequest() (request *RevokeResourceAuthorizationRequest) {
    request = &RevokeResourceAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "RevokeResourceAuthorization")
    return
}

func NewRevokeResourceAuthorizationResponse() (response *RevokeResourceAuthorizationResponse) {
    response = &RevokeResourceAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  资源所属实体对目标实体回收目标资源的相应权限，若原本没有相应权限则不产生变更。
func (c *Client) RevokeResourceAuthorization(request *RevokeResourceAuthorizationRequest) (response *RevokeResourceAuthorizationResponse, err error) {
    if request == nil {
        request = NewRevokeResourceAuthorizationRequest()
    }
    response = NewRevokeResourceAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewSearchMaterialRequest() (request *SearchMaterialRequest) {
    request = &SearchMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cme", APIVersion, "SearchMaterial")
    return
}

func NewSearchMaterialResponse() (response *SearchMaterialResponse) {
    response = &SearchMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据检索条件搜索素材，返回素材的基本信息。
func (c *Client) SearchMaterial(request *SearchMaterialRequest) (response *SearchMaterialResponse, err error) {
    if request == nil {
        request = NewSearchMaterialRequest()
    }
    response = NewSearchMaterialResponse()
    err = c.Send(request, response)
    return
}
