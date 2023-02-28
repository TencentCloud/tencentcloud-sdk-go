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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// AddTeamMember
// 向一个团队中添加团队成员，并且指定成员的角色。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMID = "InvalidParameterValue.TeamId"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddTeamMember(request *AddTeamMemberRequest) (response *AddTeamMemberResponse, err error) {
    return c.AddTeamMemberWithContext(context.Background(), request)
}

// AddTeamMember
// 向一个团队中添加团队成员，并且指定成员的角色。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMID = "InvalidParameterValue.TeamId"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddTeamMemberWithContext(ctx context.Context, request *AddTeamMemberRequest) (response *AddTeamMemberResponse, err error) {
    if request == nil {
        request = NewAddTeamMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddTeamMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddTeamMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCopyProjectRequest() (request *CopyProjectRequest) {
    request = &CopyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "CopyProject")
    
    
    return
}

func NewCopyProjectResponse() (response *CopyProjectResponse) {
    response = &CopyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyProject
// 复制一个项目，包括项目素材及轨道数据。目前仅普通剪辑及模板制作项目可复制，其它类型的项目不支持复制。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CopyProject(request *CopyProjectRequest) (response *CopyProjectResponse, err error) {
    return c.CopyProjectWithContext(context.Background(), request)
}

// CopyProject
// 复制一个项目，包括项目素材及轨道数据。目前仅普通剪辑及模板制作项目可复制，其它类型的项目不支持复制。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CopyProjectWithContext(ctx context.Context, request *CopyProjectRequest) (response *CopyProjectResponse, err error) {
    if request == nil {
        request = NewCopyProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyProjectResponse()
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

// CreateClass
// 新增分类，用于管理素材。分类层数不能超过20。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE_CLASSEXIST = "InvalidParameterValue.ClassExist"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClass(request *CreateClassRequest) (response *CreateClassResponse, err error) {
    return c.CreateClassWithContext(context.Background(), request)
}

// CreateClass
// 新增分类，用于管理素材。分类层数不能超过20。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE_CLASSEXIST = "InvalidParameterValue.ClassExist"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClassWithContext(ctx context.Context, request *CreateClassRequest) (response *CreateClassResponse, err error) {
    if request == nil {
        request = NewCreateClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClass require credential")
    }

    request.SetContext(ctx)
    
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

// CreateLink
//  创建媒体链接或分类路径链接，将源资源信息链接到目标。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateLink(request *CreateLinkRequest) (response *CreateLinkResponse, err error) {
    return c.CreateLinkWithContext(context.Background(), request)
}

// CreateLink
//  创建媒体链接或分类路径链接，将源资源信息链接到目标。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateLinkWithContext(ctx context.Context, request *CreateLinkRequest) (response *CreateLinkResponse, err error) {
    if request == nil {
        request = NewCreateLinkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLink require credential")
    }

    request.SetContext(ctx)
    
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

// CreateProject
// 创建多媒体创作引擎项目，目前支持的项目类型有：
//
// <li>视频剪辑项目：用于普通视频剪辑；</li>
//
// <li>直播剪辑项目：用于直播流剪辑；</li>
//
// <li>导播台项目：用于云导播台；</li>
//
// <li>视频拆条：用于视频拆条；</li>
//
// <li>录制回放项目：用于直播录制回放；</li>
//
// <li>云转推项目：用于直播云转推。</li>
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATERECORDTASK = "FailedOperation.CreateRecordTask"
//  FAILEDOPERATION_RECORDNOTSUPPORT = "FailedOperation.RecordNotSupport"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ASPECTRATIO = "InvalidParameterValue.AspectRatio"
//  INVALIDPARAMETERVALUE_CATEGORY = "InvalidParameterValue.Category"
//  INVALIDPARAMETERVALUE_EXTERNALMEDIAINFONOTEXIST = "InvalidParameterValue.ExternalMediaInfoNotExist"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_SWITCHERPROJECTINPUT = "InvalidParameterValue.SwitcherProjectInput"
//  INVALIDPARAMETERVALUE_TRACKITEM = "InvalidParameterValue.TrackItem"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// 创建多媒体创作引擎项目，目前支持的项目类型有：
//
// <li>视频剪辑项目：用于普通视频剪辑；</li>
//
// <li>直播剪辑项目：用于直播流剪辑；</li>
//
// <li>导播台项目：用于云导播台；</li>
//
// <li>视频拆条：用于视频拆条；</li>
//
// <li>录制回放项目：用于直播录制回放；</li>
//
// <li>云转推项目：用于直播云转推。</li>
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATERECORDTASK = "FailedOperation.CreateRecordTask"
//  FAILEDOPERATION_RECORDNOTSUPPORT = "FailedOperation.RecordNotSupport"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ASPECTRATIO = "InvalidParameterValue.AspectRatio"
//  INVALIDPARAMETERVALUE_CATEGORY = "InvalidParameterValue.Category"
//  INVALIDPARAMETERVALUE_EXTERNALMEDIAINFONOTEXIST = "InvalidParameterValue.ExternalMediaInfoNotExist"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_SWITCHERPROJECTINPUT = "InvalidParameterValue.SwitcherProjectInput"
//  INVALIDPARAMETERVALUE_TRACKITEM = "InvalidParameterValue.TrackItem"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
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

// CreateTeam
// 创建一个团队。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OWNERREMARK = "InvalidParameterValue.OwnerRemark"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMID = "InvalidParameterValue.TeamId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTeam(request *CreateTeamRequest) (response *CreateTeamResponse, err error) {
    return c.CreateTeamWithContext(context.Background(), request)
}

// CreateTeam
// 创建一个团队。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OWNERREMARK = "InvalidParameterValue.OwnerRemark"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMID = "InvalidParameterValue.TeamId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTeamWithContext(ctx context.Context, request *CreateTeamRequest) (response *CreateTeamResponse, err error) {
    if request == nil {
        request = NewCreateTeamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTeam require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTeamResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoEncodingPresetRequest() (request *CreateVideoEncodingPresetRequest) {
    request = &CreateVideoEncodingPresetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "CreateVideoEncodingPreset")
    
    
    return
}

func NewCreateVideoEncodingPresetResponse() (response *CreateVideoEncodingPresetResponse) {
    response = &CreateVideoEncodingPresetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVideoEncodingPreset
// 指定导出的参数，创建一个视频编码配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) CreateVideoEncodingPreset(request *CreateVideoEncodingPresetRequest) (response *CreateVideoEncodingPresetResponse, err error) {
    return c.CreateVideoEncodingPresetWithContext(context.Background(), request)
}

// CreateVideoEncodingPreset
// 指定导出的参数，创建一个视频编码配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) CreateVideoEncodingPresetWithContext(ctx context.Context, request *CreateVideoEncodingPresetRequest) (response *CreateVideoEncodingPresetResponse, err error) {
    if request == nil {
        request = NewCreateVideoEncodingPresetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoEncodingPreset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoEncodingPresetResponse()
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

// DeleteClass
// 删除分类信息，删除时检验下述限制：
//
// <li>分类路径必须存在；</li>
//
// <li>分类下没有绑定素材。</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSEXIST = "InvalidParameterValue.ClassExist"
//  INVALIDPARAMETERVALUE_CLASSNOTEMPTY = "InvalidParameterValue.ClassNotEmpty"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteClass(request *DeleteClassRequest) (response *DeleteClassResponse, err error) {
    return c.DeleteClassWithContext(context.Background(), request)
}

// DeleteClass
// 删除分类信息，删除时检验下述限制：
//
// <li>分类路径必须存在；</li>
//
// <li>分类下没有绑定素材。</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSEXIST = "InvalidParameterValue.ClassExist"
//  INVALIDPARAMETERVALUE_CLASSNOTEMPTY = "InvalidParameterValue.ClassNotEmpty"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteClassWithContext(ctx context.Context, request *DeleteClassRequest) (response *DeleteClassResponse, err error) {
    if request == nil {
        request = NewDeleteClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClass require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteLoginStatus
// 删除用户登录态，使用户登出多媒体创作引擎平台。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLoginStatus(request *DeleteLoginStatusRequest) (response *DeleteLoginStatusResponse, err error) {
    return c.DeleteLoginStatusWithContext(context.Background(), request)
}

// DeleteLoginStatus
// 删除用户登录态，使用户登出多媒体创作引擎平台。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLoginStatusWithContext(ctx context.Context, request *DeleteLoginStatusRequest) (response *DeleteLoginStatusResponse, err error) {
    if request == nil {
        request = NewDeleteLoginStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoginStatus require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteMaterial
// 根据媒体 Id 删除媒体。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMaterial(request *DeleteMaterialRequest) (response *DeleteMaterialResponse, err error) {
    return c.DeleteMaterialWithContext(context.Background(), request)
}

// DeleteMaterial
// 根据媒体 Id 删除媒体。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMaterialWithContext(ctx context.Context, request *DeleteMaterialRequest) (response *DeleteMaterialResponse, err error) {
    if request == nil {
        request = NewDeleteMaterialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMaterial require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteProject
// 删除项目。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SWITCHERONWORKING = "FailedOperation.SwitcherOnWorking"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    return c.DeleteProjectWithContext(context.Background(), request)
}

// DeleteProject
// 删除项目。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SWITCHERONWORKING = "FailedOperation.SwitcherOnWorking"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteTeam
// 删除一个团队。要删除团队，必须满足以下条件：
//
// <li>要删除的团队必须没有归属的素材；</li>
//
// <li>要删除的团队必须没有归属的分类。</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteTeam(request *DeleteTeamRequest) (response *DeleteTeamResponse, err error) {
    return c.DeleteTeamWithContext(context.Background(), request)
}

// DeleteTeam
// 删除一个团队。要删除团队，必须满足以下条件：
//
// <li>要删除的团队必须没有归属的素材；</li>
//
// <li>要删除的团队必须没有归属的分类。</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteTeamWithContext(ctx context.Context, request *DeleteTeamRequest) (response *DeleteTeamResponse, err error) {
    if request == nil {
        request = NewDeleteTeamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTeam require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteTeamMembers
// 将团队成员从团队中删除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MEMBERIDS = "InvalidParameterValue.MemberIds"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) DeleteTeamMembers(request *DeleteTeamMembersRequest) (response *DeleteTeamMembersResponse, err error) {
    return c.DeleteTeamMembersWithContext(context.Background(), request)
}

// DeleteTeamMembers
// 将团队成员从团队中删除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MEMBERIDS = "InvalidParameterValue.MemberIds"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) DeleteTeamMembersWithContext(ctx context.Context, request *DeleteTeamMembersRequest) (response *DeleteTeamMembersResponse, err error) {
    if request == nil {
        request = NewDeleteTeamMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTeamMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTeamMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVideoEncodingPresetRequest() (request *DeleteVideoEncodingPresetRequest) {
    request = &DeleteVideoEncodingPresetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "DeleteVideoEncodingPreset")
    
    
    return
}

func NewDeleteVideoEncodingPresetResponse() (response *DeleteVideoEncodingPresetResponse) {
    response = &DeleteVideoEncodingPresetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVideoEncodingPreset
// 删除指定 ID 的视频编码配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) DeleteVideoEncodingPreset(request *DeleteVideoEncodingPresetRequest) (response *DeleteVideoEncodingPresetResponse, err error) {
    return c.DeleteVideoEncodingPresetWithContext(context.Background(), request)
}

// DeleteVideoEncodingPreset
// 删除指定 ID 的视频编码配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) DeleteVideoEncodingPresetWithContext(ctx context.Context, request *DeleteVideoEncodingPresetRequest) (response *DeleteVideoEncodingPresetResponse, err error) {
    if request == nil {
        request = NewDeleteVideoEncodingPresetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVideoEncodingPreset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVideoEncodingPresetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// 获取平台中所有的已注册账号。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  RESOURCENOTFOUND_PLATFORMNOTFOUND = "ResourceNotFound.PlatformNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// 获取平台中所有的已注册账号。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  RESOURCENOTFOUND_PLATFORMNOTFOUND = "ResourceNotFound.PlatformNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
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

// DescribeClass
// 获取指定归属者下所有的分类信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClass(request *DescribeClassRequest) (response *DescribeClassResponse, err error) {
    return c.DescribeClassWithContext(context.Background(), request)
}

// DescribeClass
// 获取指定归属者下所有的分类信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassWithContext(ctx context.Context, request *DescribeClassRequest) (response *DescribeClassResponse, err error) {
    if request == nil {
        request = NewDescribeClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClass require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeJoinTeams
// 获取用户所加入的团队列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeJoinTeams(request *DescribeJoinTeamsRequest) (response *DescribeJoinTeamsResponse, err error) {
    return c.DescribeJoinTeamsWithContext(context.Background(), request)
}

// DescribeJoinTeams
// 获取用户所加入的团队列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeJoinTeamsWithContext(ctx context.Context, request *DescribeJoinTeamsRequest) (response *DescribeJoinTeamsResponse, err error) {
    if request == nil {
        request = NewDescribeJoinTeamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJoinTeams require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeLoginStatus
// 查询指定用户的登录态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoginStatus(request *DescribeLoginStatusRequest) (response *DescribeLoginStatusResponse, err error) {
    return c.DescribeLoginStatusWithContext(context.Background(), request)
}

// DescribeLoginStatus
// 查询指定用户的登录态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoginStatusWithContext(ctx context.Context, request *DescribeLoginStatusRequest) (response *DescribeLoginStatusResponse, err error) {
    if request == nil {
        request = NewDescribeLoginStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoginStatus require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeMaterials
// 根据媒体 Id 批量获取媒体详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMaterials(request *DescribeMaterialsRequest) (response *DescribeMaterialsResponse, err error) {
    return c.DescribeMaterialsWithContext(context.Background(), request)
}

// DescribeMaterials
// 根据媒体 Id 批量获取媒体详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMaterialsWithContext(ctx context.Context, request *DescribeMaterialsRequest) (response *DescribeMaterialsResponse, err error) {
    if request == nil {
        request = NewDescribeMaterialsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaterials require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaterialsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlatformsRequest() (request *DescribePlatformsRequest) {
    request = &DescribePlatformsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "DescribePlatforms")
    
    
    return
}

func NewDescribePlatformsResponse() (response *DescribePlatformsResponse) {
    response = &DescribePlatformsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePlatforms
// <li>支持获取所创建的所有平台列表信息；</li>
//
// <li>支持获取指定的平台列表信息。</li>
//
// 
//
// 关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
//
// 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePlatforms(request *DescribePlatformsRequest) (response *DescribePlatformsResponse, err error) {
    return c.DescribePlatformsWithContext(context.Background(), request)
}

// DescribePlatforms
// <li>支持获取所创建的所有平台列表信息；</li>
//
// <li>支持获取指定的平台列表信息。</li>
//
// 
//
// 关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
//
// 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePlatformsWithContext(ctx context.Context, request *DescribePlatformsRequest) (response *DescribePlatformsResponse, err error) {
    if request == nil {
        request = NewDescribePlatformsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlatforms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePlatformsResponse()
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

// DescribeProjects
// 支持根据多种条件过滤出项目列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ASPECTRATIOSET = "InvalidParameterValue.AspectRatioSet"
//  INVALIDPARAMETERVALUE_CATEGORYSET = "InvalidParameterValue.CategorySet"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_SORTORDER = "InvalidParameterValue.SortOrder"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    return c.DescribeProjectsWithContext(context.Background(), request)
}

// DescribeProjects
// 支持根据多种条件过滤出项目列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ASPECTRATIOSET = "InvalidParameterValue.AspectRatioSet"
//  INVALIDPARAMETERVALUE_CATEGORYSET = "InvalidParameterValue.CategorySet"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_SORTORDER = "InvalidParameterValue.SortOrder"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjects require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeResourceAuthorization
// 查询资源被授权的情况。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeResourceAuthorization(request *DescribeResourceAuthorizationRequest) (response *DescribeResourceAuthorizationResponse, err error) {
    return c.DescribeResourceAuthorizationWithContext(context.Background(), request)
}

// DescribeResourceAuthorization
// 查询资源被授权的情况。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeResourceAuthorizationWithContext(ctx context.Context, request *DescribeResourceAuthorizationRequest) (response *DescribeResourceAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeResourceAuthorizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceAuthorization require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeSharedSpace
// 获取共享空间。当个人或团队A对个人或团队B授权某资源以后，个人或团队B的共享空间就会增加个人或团队A。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSharedSpace(request *DescribeSharedSpaceRequest) (response *DescribeSharedSpaceResponse, err error) {
    return c.DescribeSharedSpaceWithContext(context.Background(), request)
}

// DescribeSharedSpace
// 获取共享空间。当个人或团队A对个人或团队B授权某资源以后，个人或团队B的共享空间就会增加个人或团队A。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSharedSpaceWithContext(ctx context.Context, request *DescribeSharedSpaceRequest) (response *DescribeSharedSpaceResponse, err error) {
    if request == nil {
        request = NewDescribeSharedSpaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSharedSpace require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTaskDetail
// 获取任务详情信息，包含下面几个部分：
//
// <li>任务基础信息：包括任务状态、错误信息、创建时间等；</li>
//
// <li>导出项目输出信息：包括输出的素材 Id 等。</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    return c.DescribeTaskDetailWithContext(context.Background(), request)
}

// DescribeTaskDetail
// 获取任务详情信息，包含下面几个部分：
//
// <li>任务基础信息：包括任务状态、错误信息、创建时间等；</li>
//
// <li>导出项目输出信息：包括输出的素材 Id 等。</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskDetail require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTasks
// 获取任务列表，支持条件筛选，返回对应的任务基础信息列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// 获取任务列表，支持条件筛选，返回对应的任务基础信息列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTeamMembers
// 获取指定团队的成员信息。支持获取指定成员的信息，同时也支持分页拉取指定团队的所有成员信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTeamMembers(request *DescribeTeamMembersRequest) (response *DescribeTeamMembersResponse, err error) {
    return c.DescribeTeamMembersWithContext(context.Background(), request)
}

// DescribeTeamMembers
// 获取指定团队的成员信息。支持获取指定成员的信息，同时也支持分页拉取指定团队的所有成员信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTeamMembersWithContext(ctx context.Context, request *DescribeTeamMembersRequest) (response *DescribeTeamMembersResponse, err error) {
    if request == nil {
        request = NewDescribeTeamMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTeamMembers require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTeams
// 获取指定团队的信息，拉取团队信息列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTeams(request *DescribeTeamsRequest) (response *DescribeTeamsResponse, err error) {
    return c.DescribeTeamsWithContext(context.Background(), request)
}

// DescribeTeams
// 获取指定团队的信息，拉取团队信息列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTeamsWithContext(ctx context.Context, request *DescribeTeamsRequest) (response *DescribeTeamsResponse, err error) {
    if request == nil {
        request = NewDescribeTeamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTeams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTeamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoEncodingPresetsRequest() (request *DescribeVideoEncodingPresetsRequest) {
    request = &DescribeVideoEncodingPresetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "DescribeVideoEncodingPresets")
    
    
    return
}

func NewDescribeVideoEncodingPresetsResponse() (response *DescribeVideoEncodingPresetsResponse) {
    response = &DescribeVideoEncodingPresetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVideoEncodingPresets
// 查询视频编码配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) DescribeVideoEncodingPresets(request *DescribeVideoEncodingPresetsRequest) (response *DescribeVideoEncodingPresetsResponse, err error) {
    return c.DescribeVideoEncodingPresetsWithContext(context.Background(), request)
}

// DescribeVideoEncodingPresets
// 查询视频编码配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) DescribeVideoEncodingPresetsWithContext(ctx context.Context, request *DescribeVideoEncodingPresetsRequest) (response *DescribeVideoEncodingPresetsResponse, err error) {
    if request == nil {
        request = NewDescribeVideoEncodingPresetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoEncodingPresets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoEncodingPresetsResponse()
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

// ExportVideoByEditorTrackData
// 使用 [视频合成协议](https://cloud.tencent.com/document/product/1156/51225) 合成视频，支持导出视频到 CME 云媒资或者云点播媒资。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_EXPORTDESTINATION = "InvalidParameterValue.ExportDestination"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_THIRDYPARTYPUBLISHCHANNELID = "InvalidParameterValue.ThirdyPartyPublishChannelId"
//  INVALIDPARAMETERVALUE_TRACKDATA = "InvalidParameterValue.TrackData"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  LIMITEXCEEDED_BILLITEMSTORAGE = "LimitExceeded.BillItemStorage"
//  LIMITEXCEEDED_BILLITEMVIDEOEDITEXPORTDURATION = "LimitExceeded.BillItemVideoEditExportDuration"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExportVideoByEditorTrackData(request *ExportVideoByEditorTrackDataRequest) (response *ExportVideoByEditorTrackDataResponse, err error) {
    return c.ExportVideoByEditorTrackDataWithContext(context.Background(), request)
}

// ExportVideoByEditorTrackData
// 使用 [视频合成协议](https://cloud.tencent.com/document/product/1156/51225) 合成视频，支持导出视频到 CME 云媒资或者云点播媒资。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_EXPORTDESTINATION = "InvalidParameterValue.ExportDestination"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_THIRDYPARTYPUBLISHCHANNELID = "InvalidParameterValue.ThirdyPartyPublishChannelId"
//  INVALIDPARAMETERVALUE_TRACKDATA = "InvalidParameterValue.TrackData"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  LIMITEXCEEDED_BILLITEMSTORAGE = "LimitExceeded.BillItemStorage"
//  LIMITEXCEEDED_BILLITEMVIDEOEDITEXPORTDURATION = "LimitExceeded.BillItemVideoEditExportDuration"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExportVideoByEditorTrackDataWithContext(ctx context.Context, request *ExportVideoByEditorTrackDataRequest) (response *ExportVideoByEditorTrackDataResponse, err error) {
    if request == nil {
        request = NewExportVideoByEditorTrackDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVideoByEditorTrackData require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVideoByEditorTrackDataResponse()
    err = c.Send(request, response)
    return
}

func NewExportVideoByTemplateRequest() (request *ExportVideoByTemplateRequest) {
    request = &ExportVideoByTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "ExportVideoByTemplate")
    
    
    return
}

func NewExportVideoByTemplateResponse() (response *ExportVideoByTemplateResponse) {
    response = &ExportVideoByTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVideoByTemplate
// 使用视频剪辑模板直接导出视频。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_EXPORTDESTINATION = "InvalidParameterValue.ExportDestination"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_MEDIAREPLACEMENTINFO = "InvalidParameterValue.MediaReplacementInfo"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_REPLACEMENTTYPE = "InvalidParameterValue.ReplacementType"
//  INVALIDPARAMETERVALUE_TRACKDATA = "InvalidParameterValue.TrackData"
//  INVALIDPARAMETERVALUE_VIDEOEDITTEMPLATEIDNOTEXIST = "InvalidParameterValue.VideoEditTemplateIdNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_BILLITEMSTORAGE = "LimitExceeded.BillItemStorage"
//  LIMITEXCEEDED_BILLITEMVIDEOEDITEXPORTDURATION = "LimitExceeded.BillItemVideoEditExportDuration"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) ExportVideoByTemplate(request *ExportVideoByTemplateRequest) (response *ExportVideoByTemplateResponse, err error) {
    return c.ExportVideoByTemplateWithContext(context.Background(), request)
}

// ExportVideoByTemplate
// 使用视频剪辑模板直接导出视频。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_EXPORTDESTINATION = "InvalidParameterValue.ExportDestination"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_MEDIAREPLACEMENTINFO = "InvalidParameterValue.MediaReplacementInfo"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_REPLACEMENTTYPE = "InvalidParameterValue.ReplacementType"
//  INVALIDPARAMETERVALUE_TRACKDATA = "InvalidParameterValue.TrackData"
//  INVALIDPARAMETERVALUE_VIDEOEDITTEMPLATEIDNOTEXIST = "InvalidParameterValue.VideoEditTemplateIdNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_BILLITEMSTORAGE = "LimitExceeded.BillItemStorage"
//  LIMITEXCEEDED_BILLITEMVIDEOEDITEXPORTDURATION = "LimitExceeded.BillItemVideoEditExportDuration"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) ExportVideoByTemplateWithContext(ctx context.Context, request *ExportVideoByTemplateRequest) (response *ExportVideoByTemplateResponse, err error) {
    if request == nil {
        request = NewExportVideoByTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVideoByTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVideoByTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewExportVideoByVideoSegmentationDataRequest() (request *ExportVideoByVideoSegmentationDataRequest) {
    request = &ExportVideoByVideoSegmentationDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "ExportVideoByVideoSegmentationData")
    
    
    return
}

func NewExportVideoByVideoSegmentationDataResponse() (response *ExportVideoByVideoSegmentationDataResponse) {
    response = &ExportVideoByVideoSegmentationDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVideoByVideoSegmentationData
// 使用视频智能拆条数据导出视频，将指定的视频拆条片段导出为一个视频(内测中，请勿使用)。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_EXPORTDESTINATION = "InvalidParameterValue.ExportDestination"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  LIMITEXCEEDED_BILLITEMSTORAGE = "LimitExceeded.BillItemStorage"
//  LIMITEXCEEDED_BILLITEMVIDEOEDITEXPORTDURATION = "LimitExceeded.BillItemVideoEditExportDuration"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) ExportVideoByVideoSegmentationData(request *ExportVideoByVideoSegmentationDataRequest) (response *ExportVideoByVideoSegmentationDataResponse, err error) {
    return c.ExportVideoByVideoSegmentationDataWithContext(context.Background(), request)
}

// ExportVideoByVideoSegmentationData
// 使用视频智能拆条数据导出视频，将指定的视频拆条片段导出为一个视频(内测中，请勿使用)。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_EXPORTDESTINATION = "InvalidParameterValue.ExportDestination"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  LIMITEXCEEDED_BILLITEMSTORAGE = "LimitExceeded.BillItemStorage"
//  LIMITEXCEEDED_BILLITEMVIDEOEDITEXPORTDURATION = "LimitExceeded.BillItemVideoEditExportDuration"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) ExportVideoByVideoSegmentationDataWithContext(ctx context.Context, request *ExportVideoByVideoSegmentationDataRequest) (response *ExportVideoByVideoSegmentationDataResponse, err error) {
    if request == nil {
        request = NewExportVideoByVideoSegmentationDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVideoByVideoSegmentationData require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVideoByVideoSegmentationDataResponse()
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

// ExportVideoEditProject
// 导出视频编辑项目，支持指定输出的模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_EXPORTDESTINATION = "InvalidParameterValue.ExportDestination"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  INVALIDPARAMETERVALUE_THIRDYPARTYPUBLISHCHANNELID = "InvalidParameterValue.ThirdyPartyPublishChannelId"
//  INVALIDPARAMETERVALUE_TRACKDATA = "InvalidParameterValue.TrackData"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  LIMITEXCEEDED_BILLITEMSTORAGE = "LimitExceeded.BillItemStorage"
//  LIMITEXCEEDED_BILLITEMVIDEOEDITEXPORTDURATION = "LimitExceeded.BillItemVideoEditExportDuration"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExportVideoEditProject(request *ExportVideoEditProjectRequest) (response *ExportVideoEditProjectResponse, err error) {
    return c.ExportVideoEditProjectWithContext(context.Background(), request)
}

// ExportVideoEditProject
// 导出视频编辑项目，支持指定输出的模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_EXPORTDESTINATION = "InvalidParameterValue.ExportDestination"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  INVALIDPARAMETERVALUE_THIRDYPARTYPUBLISHCHANNELID = "InvalidParameterValue.ThirdyPartyPublishChannelId"
//  INVALIDPARAMETERVALUE_TRACKDATA = "InvalidParameterValue.TrackData"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  LIMITEXCEEDED_BILLITEMSTORAGE = "LimitExceeded.BillItemStorage"
//  LIMITEXCEEDED_BILLITEMVIDEOEDITEXPORTDURATION = "LimitExceeded.BillItemVideoEditExportDuration"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExportVideoEditProjectWithContext(ctx context.Context, request *ExportVideoEditProjectRequest) (response *ExportVideoEditProjectResponse, err error) {
    if request == nil {
        request = NewExportVideoEditProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVideoEditProject require credential")
    }

    request.SetContext(ctx)
    
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

// FlattenListMedia
// 平铺分类路径下及其子分类下的所有媒体基础信息，返回当前分类及子分类中的所有媒体的基础信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) FlattenListMedia(request *FlattenListMediaRequest) (response *FlattenListMediaResponse, err error) {
    return c.FlattenListMediaWithContext(context.Background(), request)
}

// FlattenListMedia
// 平铺分类路径下及其子分类下的所有媒体基础信息，返回当前分类及子分类中的所有媒体的基础信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) FlattenListMediaWithContext(ctx context.Context, request *FlattenListMediaRequest) (response *FlattenListMediaResponse, err error) {
    if request == nil {
        request = NewFlattenListMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FlattenListMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewFlattenListMediaResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateVideoSegmentationSchemeByAiRequest() (request *GenerateVideoSegmentationSchemeByAiRequest) {
    request = &GenerateVideoSegmentationSchemeByAiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "GenerateVideoSegmentationSchemeByAi")
    
    
    return
}

func NewGenerateVideoSegmentationSchemeByAiResponse() (response *GenerateVideoSegmentationSchemeByAiResponse) {
    response = &GenerateVideoSegmentationSchemeByAiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateVideoSegmentationSchemeByAi
// <li>发起视频智能拆条任务，支持智能生成和平精英集锦、王者荣耀集锦、足球集锦、篮球集锦 、人物集锦、新闻拆条等任务。</li>
//
// <li>和平精英集锦和王者荣耀集锦根据击杀场景进行拆条，足球集锦和篮球集锦根据进球场景进行拆条，人物集锦根据人物人脸特征进行拆条，新闻拆条根据导播进行拆条。</li>
//
// <li>【本接口内测中，暂不建议使用】</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) GenerateVideoSegmentationSchemeByAi(request *GenerateVideoSegmentationSchemeByAiRequest) (response *GenerateVideoSegmentationSchemeByAiResponse, err error) {
    return c.GenerateVideoSegmentationSchemeByAiWithContext(context.Background(), request)
}

// GenerateVideoSegmentationSchemeByAi
// <li>发起视频智能拆条任务，支持智能生成和平精英集锦、王者荣耀集锦、足球集锦、篮球集锦 、人物集锦、新闻拆条等任务。</li>
//
// <li>和平精英集锦和王者荣耀集锦根据击杀场景进行拆条，足球集锦和篮球集锦根据进球场景进行拆条，人物集锦根据人物人脸特征进行拆条，新闻拆条根据导播进行拆条。</li>
//
// <li>【本接口内测中，暂不建议使用】</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) GenerateVideoSegmentationSchemeByAiWithContext(ctx context.Context, request *GenerateVideoSegmentationSchemeByAiRequest) (response *GenerateVideoSegmentationSchemeByAiResponse, err error) {
    if request == nil {
        request = NewGenerateVideoSegmentationSchemeByAiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateVideoSegmentationSchemeByAi require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateVideoSegmentationSchemeByAiResponse()
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

// GrantResourceAuthorization
// 资源归属者对个人或团队授予目标资源的相应权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GrantResourceAuthorization(request *GrantResourceAuthorizationRequest) (response *GrantResourceAuthorizationResponse, err error) {
    return c.GrantResourceAuthorizationWithContext(context.Background(), request)
}

// GrantResourceAuthorization
// 资源归属者对个人或团队授予目标资源的相应权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GrantResourceAuthorizationWithContext(ctx context.Context, request *GrantResourceAuthorizationRequest) (response *GrantResourceAuthorizationResponse, err error) {
    if request == nil {
        request = NewGrantResourceAuthorizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GrantResourceAuthorization require credential")
    }

    request.SetContext(ctx)
    
    response = NewGrantResourceAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewHandleStreamConnectProjectRequest() (request *HandleStreamConnectProjectRequest) {
    request = &HandleStreamConnectProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "HandleStreamConnectProject")
    
    
    return
}

func NewHandleStreamConnectProjectResponse() (response *HandleStreamConnectProjectResponse) {
    response = &HandleStreamConnectProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// HandleStreamConnectProject
// 对云转推项目进行操作。
//
// ### 操作类型<a id="Operation"></a>
//
// - `AddInput`（添加输入源），包括：
//
// 	- 添加直播拉流输入源，参见 [示例1](#.E7.A4.BA.E4.BE.8B1-.E6.B7.BB.E5.8A.A0.E7.9B.B4.E6.92.AD.E6.8B.89.E6.B5.81.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// 	- 添加直播推流输入源，参见 [示例2](#.E7.A4.BA.E4.BE.8B2-.E6.B7.BB.E5.8A.A0.E7.9B.B4.E6.92.AD.E6.8E.A8.E6.B5.81.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// 	- 添加点播拉流输入源，参见 [示例3](#.E7.A4.BA.E4.BE.8B3-.E6.B7.BB.E5.8A.A0.E7.82.B9.E6.92.AD.E6.8B.89.E6.B5.81.E8.BE.93.E5.85.A5.E6.BA.90.E4.B8.94.E5.BE.AA.E7.8E.AF.E6.92.AD.E6.94.BE)、[示例4](#.E7.A4.BA.E4.BE.8B4-.E6.B7.BB.E5.8A.A0.E7.82.B9.E6.92.AD.E6.8B.89.E6.B5.81.E8.BE.93.E5.85.A5.E6.BA.90.E4.B8.94.E5.8D.95.E6.AC.A1.E6.92.AD.E6.94.BE)；
//
// - `DeleteInput`（删除输入源），参见 [示例5](#.E7.A4.BA.E4.BE.8B5-.E5.88.A0.E9.99.A4.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// - `ModifyInput`（修改输入源），参见 [示例6](#.E7.A4.BA.E4.BE.8B6-.E4.BF.AE.E6.94.B9.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// - `AddOutput`（ 添加输出源），参见 [示例7](#.E7.A4.BA.E4.BE.8B7-.E6.B7.BB.E5.8A.A0.E8.BE.93.E5.87.BA.E6.BA.90)；
//
// - `DeleteOutput`（删除输出源），参见 [示例8](#.E7.A4.BA.E4.BE.8B8-.E5.88.A0.E9.99.A4.E8.BE.93.E5.87.BA.E6.BA.90)；
//
// - `ModifyOutput`（修改输出源），参见 [示例9](#.E7.A4.BA.E4.BE.8B9-.E4.BF.AE.E6.94.B9.E8.BE.93.E5.87.BA.E6.BA.90)；
//
// - `Start`（开启转推），参见 [示例10](#.E7.A4.BA.E4.BE.8B10-.E5.BC.80.E5.90.AF.E4.BA.91.E8.BD.AC.E6.8E.A8)；
//
// - `Stop`（停止转推），参见 [示例11](#.E7.A4.BA.E4.BE.8B11-.E5.81.9C.E6.AD.A2.E4.BA.91.E8.BD.AC.E6.8E.A8)；
//
// - `SwitchInput`（切换输入源），参见 [示例12](#.E7.A4.BA.E4.BE.8B12-.E5.88.87.E6.8D.A2.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// - `ModifyCurrentStopTime`（修改当前计划结束时间），参见 [示例13](#.E7.A4.BA.E4.BE.8B13-.E4.BF.AE.E6.94.B9.E8.BD.AC.E6.8E.A8.E7.BB.93.E6.9D.9F.E6.97.B6.E9.97.B4);
//
// - `DescribeInputPlayInfo`（查询播放进度），参见 [示例14](#.E7.A4.BA.E4.BE.8B14-.E6.9F.A5.E8.AF.A2.E7.82.B9.E6.92.AD.E8.BE.93.E5.85.A5.E6.BA.90.E6.92.AD.E6.94.BE.E8.BF.9B.E5.BA.A6)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STREAMCONNECT = "FailedOperation.StreamConnect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CATEGORY = "InvalidParameterValue.Category"
//  INVALIDPARAMETERVALUE_EXTINFOINVALID = "InvalidParameterValue.ExtInfoInvalid"
//  INVALIDPARAMETERVALUE_INPUT = "InvalidParameterValue.Input"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  INVALIDPARAMETERVALUE_STREAMCONNECT = "InvalidParameterValue.StreamConnect"
//  INVALIDPARAMETERVALUE_STREAMCONNECTINPUTINVALID = "InvalidParameterValue.StreamConnectInputInvalid"
//  INVALIDPARAMETERVALUE_STREAMCONNECTOUTPUTINVALID = "InvalidParameterValue.StreamConnectOutputInvalid"
//  INVALIDPARAMETERVALUE_STREAMINPUT = "InvalidParameterValue.StreamInput"
//  LIMITEXCEEDED_BILLITEMLIVEDISPATCHDURATION = "LimitExceeded.BillItemLiveDispatchDuration"
//  LIMITEXCEEDED_BILLITEMLIVEDISPATCHMAXCOUNT = "LimitExceeded.BillItemLiveDispatchMaxCount"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) HandleStreamConnectProject(request *HandleStreamConnectProjectRequest) (response *HandleStreamConnectProjectResponse, err error) {
    return c.HandleStreamConnectProjectWithContext(context.Background(), request)
}

// HandleStreamConnectProject
// 对云转推项目进行操作。
//
// ### 操作类型<a id="Operation"></a>
//
// - `AddInput`（添加输入源），包括：
//
// 	- 添加直播拉流输入源，参见 [示例1](#.E7.A4.BA.E4.BE.8B1-.E6.B7.BB.E5.8A.A0.E7.9B.B4.E6.92.AD.E6.8B.89.E6.B5.81.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// 	- 添加直播推流输入源，参见 [示例2](#.E7.A4.BA.E4.BE.8B2-.E6.B7.BB.E5.8A.A0.E7.9B.B4.E6.92.AD.E6.8E.A8.E6.B5.81.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// 	- 添加点播拉流输入源，参见 [示例3](#.E7.A4.BA.E4.BE.8B3-.E6.B7.BB.E5.8A.A0.E7.82.B9.E6.92.AD.E6.8B.89.E6.B5.81.E8.BE.93.E5.85.A5.E6.BA.90.E4.B8.94.E5.BE.AA.E7.8E.AF.E6.92.AD.E6.94.BE)、[示例4](#.E7.A4.BA.E4.BE.8B4-.E6.B7.BB.E5.8A.A0.E7.82.B9.E6.92.AD.E6.8B.89.E6.B5.81.E8.BE.93.E5.85.A5.E6.BA.90.E4.B8.94.E5.8D.95.E6.AC.A1.E6.92.AD.E6.94.BE)；
//
// - `DeleteInput`（删除输入源），参见 [示例5](#.E7.A4.BA.E4.BE.8B5-.E5.88.A0.E9.99.A4.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// - `ModifyInput`（修改输入源），参见 [示例6](#.E7.A4.BA.E4.BE.8B6-.E4.BF.AE.E6.94.B9.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// - `AddOutput`（ 添加输出源），参见 [示例7](#.E7.A4.BA.E4.BE.8B7-.E6.B7.BB.E5.8A.A0.E8.BE.93.E5.87.BA.E6.BA.90)；
//
// - `DeleteOutput`（删除输出源），参见 [示例8](#.E7.A4.BA.E4.BE.8B8-.E5.88.A0.E9.99.A4.E8.BE.93.E5.87.BA.E6.BA.90)；
//
// - `ModifyOutput`（修改输出源），参见 [示例9](#.E7.A4.BA.E4.BE.8B9-.E4.BF.AE.E6.94.B9.E8.BE.93.E5.87.BA.E6.BA.90)；
//
// - `Start`（开启转推），参见 [示例10](#.E7.A4.BA.E4.BE.8B10-.E5.BC.80.E5.90.AF.E4.BA.91.E8.BD.AC.E6.8E.A8)；
//
// - `Stop`（停止转推），参见 [示例11](#.E7.A4.BA.E4.BE.8B11-.E5.81.9C.E6.AD.A2.E4.BA.91.E8.BD.AC.E6.8E.A8)；
//
// - `SwitchInput`（切换输入源），参见 [示例12](#.E7.A4.BA.E4.BE.8B12-.E5.88.87.E6.8D.A2.E8.BE.93.E5.85.A5.E6.BA.90)；
//
// - `ModifyCurrentStopTime`（修改当前计划结束时间），参见 [示例13](#.E7.A4.BA.E4.BE.8B13-.E4.BF.AE.E6.94.B9.E8.BD.AC.E6.8E.A8.E7.BB.93.E6.9D.9F.E6.97.B6.E9.97.B4);
//
// - `DescribeInputPlayInfo`（查询播放进度），参见 [示例14](#.E7.A4.BA.E4.BE.8B14-.E6.9F.A5.E8.AF.A2.E7.82.B9.E6.92.AD.E8.BE.93.E5.85.A5.E6.BA.90.E6.92.AD.E6.94.BE.E8.BF.9B.E5.BA.A6)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STREAMCONNECT = "FailedOperation.StreamConnect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CATEGORY = "InvalidParameterValue.Category"
//  INVALIDPARAMETERVALUE_EXTINFOINVALID = "InvalidParameterValue.ExtInfoInvalid"
//  INVALIDPARAMETERVALUE_INPUT = "InvalidParameterValue.Input"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  INVALIDPARAMETERVALUE_STREAMCONNECT = "InvalidParameterValue.StreamConnect"
//  INVALIDPARAMETERVALUE_STREAMCONNECTINPUTINVALID = "InvalidParameterValue.StreamConnectInputInvalid"
//  INVALIDPARAMETERVALUE_STREAMCONNECTOUTPUTINVALID = "InvalidParameterValue.StreamConnectOutputInvalid"
//  INVALIDPARAMETERVALUE_STREAMINPUT = "InvalidParameterValue.StreamInput"
//  LIMITEXCEEDED_BILLITEMLIVEDISPATCHDURATION = "LimitExceeded.BillItemLiveDispatchDuration"
//  LIMITEXCEEDED_BILLITEMLIVEDISPATCHMAXCOUNT = "LimitExceeded.BillItemLiveDispatchMaxCount"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) HandleStreamConnectProjectWithContext(ctx context.Context, request *HandleStreamConnectProjectRequest) (response *HandleStreamConnectProjectResponse, err error) {
    if request == nil {
        request = NewHandleStreamConnectProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HandleStreamConnectProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewHandleStreamConnectProjectResponse()
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

// ImportMaterial
// 将云点播媒资文件导入到多媒体创作引擎媒体资源库。支持导入媒体归属团队或者个人。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_EXTERNALMEDIAINFONOTEXIST = "InvalidParameterValue.ExternalMediaInfoNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PREPROCESSDEFINITION = "InvalidParameterValue.PreProcessDefinition"
//  INVALIDPARAMETERVALUE_VODFILEID = "InvalidParameterValue.VodFileId"
//  INVALIDPARAMETERVALUE_VODFILENOTEXIST = "InvalidParameterValue.VodFileNotExist"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ImportMaterial(request *ImportMaterialRequest) (response *ImportMaterialResponse, err error) {
    return c.ImportMaterialWithContext(context.Background(), request)
}

// ImportMaterial
// 将云点播媒资文件导入到多媒体创作引擎媒体资源库。支持导入媒体归属团队或者个人。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_EXTERNALMEDIAINFONOTEXIST = "InvalidParameterValue.ExternalMediaInfoNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PREPROCESSDEFINITION = "InvalidParameterValue.PreProcessDefinition"
//  INVALIDPARAMETERVALUE_VODFILEID = "InvalidParameterValue.VodFileId"
//  INVALIDPARAMETERVALUE_VODFILENOTEXIST = "InvalidParameterValue.VodFileNotExist"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ImportMaterialWithContext(ctx context.Context, request *ImportMaterialRequest) (response *ImportMaterialResponse, err error) {
    if request == nil {
        request = NewImportMaterialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportMaterial require credential")
    }

    request.SetContext(ctx)
    
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

// ImportMediaToProject
// 将云点播中的媒资或者用户自有媒资文件添加到项目中与项目关联，供后续视频编辑使用。目前仅视频编辑项目和智能视频拆条项目有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EXTERNALMEDIAINFONOTEXIST = "InvalidParameterValue.ExternalMediaInfoNotExist"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PREPROCESSDEFINITION = "InvalidParameterValue.PreProcessDefinition"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  INVALIDPARAMETERVALUE_VODFILEID = "InvalidParameterValue.VodFileId"
//  INVALIDPARAMETERVALUE_VODFILENOTEXIST = "InvalidParameterValue.VodFileNotExist"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ImportMediaToProject(request *ImportMediaToProjectRequest) (response *ImportMediaToProjectResponse, err error) {
    return c.ImportMediaToProjectWithContext(context.Background(), request)
}

// ImportMediaToProject
// 将云点播中的媒资或者用户自有媒资文件添加到项目中与项目关联，供后续视频编辑使用。目前仅视频编辑项目和智能视频拆条项目有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR_CREATETASK = "InternalError.CreateTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EXTERNALMEDIAINFONOTEXIST = "InvalidParameterValue.ExternalMediaInfoNotExist"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PREPROCESSDEFINITION = "InvalidParameterValue.PreProcessDefinition"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  INVALIDPARAMETERVALUE_VODFILEID = "InvalidParameterValue.VodFileId"
//  INVALIDPARAMETERVALUE_VODFILENOTEXIST = "InvalidParameterValue.VodFileNotExist"
//  INVALIDPARAMETERVALUE_VODSUBAPPID = "InvalidParameterValue.VodSubAppid"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ImportMediaToProjectWithContext(ctx context.Context, request *ImportMediaToProjectRequest) (response *ImportMediaToProjectResponse, err error) {
    if request == nil {
        request = NewImportMediaToProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportMediaToProject require credential")
    }

    request.SetContext(ctx)
    
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

// ListMedia
//  浏览当前分类路径下的资源，包括媒体文件和子分类，返回媒资基础信息和分类信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListMedia(request *ListMediaRequest) (response *ListMediaResponse, err error) {
    return c.ListMediaWithContext(context.Background(), request)
}

// ListMedia
//  浏览当前分类路径下的资源，包括媒体文件和子分类，返回媒资基础信息和分类信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListMediaWithContext(ctx context.Context, request *ListMediaRequest) (response *ListMediaResponse, err error) {
    if request == nil {
        request = NewListMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListMedia require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyMaterial
// 修改媒体信息，支持修改媒体名称、分类路径、标签等信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyMaterial(request *ModifyMaterialRequest) (response *ModifyMaterialResponse, err error) {
    return c.ModifyMaterialWithContext(context.Background(), request)
}

// ModifyMaterial
// 修改媒体信息，支持修改媒体名称、分类路径、标签等信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyMaterialWithContext(ctx context.Context, request *ModifyMaterialRequest) (response *ModifyMaterialResponse, err error) {
    if request == nil {
        request = NewModifyMaterialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaterial require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyProject
// 修改项目信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ASPECTRATIO = "InvalidParameterValue.AspectRatio"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    return c.ModifyProjectWithContext(context.Background(), request)
}

// ModifyProject
// 修改项目信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ASPECTRATIO = "InvalidParameterValue.AspectRatio"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_NAMELENLIMT = "InvalidParameterValue.NameLenLimt"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_PROJECTID = "InvalidParameterValue.ProjectId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyProjectWithContext(ctx context.Context, request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProject require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyTeam
// 修改团队信息，目前支持修改的操作有：
//
// <li>修改团队名称。</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTeam(request *ModifyTeamRequest) (response *ModifyTeamResponse, err error) {
    return c.ModifyTeamWithContext(context.Background(), request)
}

// ModifyTeam
// 修改团队信息，目前支持修改的操作有：
//
// <li>修改团队名称。</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTeamWithContext(ctx context.Context, request *ModifyTeamRequest) (response *ModifyTeamResponse, err error) {
    if request == nil {
        request = NewModifyTeamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTeam require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyTeamMember
// 修改团队成员信息，包括成员备注、角色等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MEMBERNOTEXIST = "InvalidParameterValue.MemberNotExist"
//  INVALIDPARAMETERVALUE_NOTTEAMMEMBERERROR = "InvalidParameterValue.NotTeamMemberError"
//  INVALIDPARAMETERVALUE_OPERATOR = "InvalidParameterValue.Operator"
//  INVALIDPARAMETERVALUE_ROLE = "InvalidParameterValue.Role"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTeamMember(request *ModifyTeamMemberRequest) (response *ModifyTeamMemberResponse, err error) {
    return c.ModifyTeamMemberWithContext(context.Background(), request)
}

// ModifyTeamMember
// 修改团队成员信息，包括成员备注、角色等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MEMBERNOTEXIST = "InvalidParameterValue.MemberNotExist"
//  INVALIDPARAMETERVALUE_NOTTEAMMEMBERERROR = "InvalidParameterValue.NotTeamMemberError"
//  INVALIDPARAMETERVALUE_OPERATOR = "InvalidParameterValue.Operator"
//  INVALIDPARAMETERVALUE_ROLE = "InvalidParameterValue.Role"
//  INVALIDPARAMETERVALUE_TEAMNOTEXIST = "InvalidParameterValue.TeamNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTeamMemberWithContext(ctx context.Context, request *ModifyTeamMemberRequest) (response *ModifyTeamMemberResponse, err error) {
    if request == nil {
        request = NewModifyTeamMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTeamMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTeamMemberResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVideoEncodingPresetRequest() (request *ModifyVideoEncodingPresetRequest) {
    request = &ModifyVideoEncodingPresetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "ModifyVideoEncodingPreset")
    
    
    return
}

func NewModifyVideoEncodingPresetResponse() (response *ModifyVideoEncodingPresetResponse) {
    response = &ModifyVideoEncodingPresetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVideoEncodingPreset
// 修改视频编码配置信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) ModifyVideoEncodingPreset(request *ModifyVideoEncodingPresetRequest) (response *ModifyVideoEncodingPresetResponse, err error) {
    return c.ModifyVideoEncodingPresetWithContext(context.Background(), request)
}

// ModifyVideoEncodingPreset
// 修改视频编码配置信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) ModifyVideoEncodingPresetWithContext(ctx context.Context, request *ModifyVideoEncodingPresetRequest) (response *ModifyVideoEncodingPresetResponse, err error) {
    if request == nil {
        request = NewModifyVideoEncodingPresetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVideoEncodingPreset require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVideoEncodingPresetResponse()
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

// MoveClass
// 移动某一个分类到另外一个分类下，也可用于分类重命名。
//
// 如果 SourceClassPath = /素材/视频/NBA，DestinationClassPath = /素材/视频/篮球
//
// <li>当 DestinationClassPath 不存在时候，操作结果为重命名 ClassPath；</li>
//
// <li>当 DestinationClassPath 存在时候，操作结果为产生新目录 /素材/视频/篮球/NBA</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSEXIST = "InvalidParameterValue.ClassExist"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_DSTCLASSPATHNOTEXIST = "InvalidParameterValue.DstClassPathNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) MoveClass(request *MoveClassRequest) (response *MoveClassResponse, err error) {
    return c.MoveClassWithContext(context.Background(), request)
}

// MoveClass
// 移动某一个分类到另外一个分类下，也可用于分类重命名。
//
// 如果 SourceClassPath = /素材/视频/NBA，DestinationClassPath = /素材/视频/篮球
//
// <li>当 DestinationClassPath 不存在时候，操作结果为重命名 ClassPath；</li>
//
// <li>当 DestinationClassPath 存在时候，操作结果为产生新目录 /素材/视频/篮球/NBA</li>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSEXIST = "InvalidParameterValue.ClassExist"
//  INVALIDPARAMETERVALUE_CLASSNOTEXIST = "InvalidParameterValue.ClassNotExist"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_DSTCLASSPATHNOTEXIST = "InvalidParameterValue.DstClassPathNotExist"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) MoveClassWithContext(ctx context.Context, request *MoveClassRequest) (response *MoveClassResponse, err error) {
    if request == nil {
        request = NewMoveClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MoveClass require credential")
    }

    request.SetContext(ctx)
    
    response = NewMoveClassResponse()
    err = c.Send(request, response)
    return
}

func NewMoveResourceRequest() (request *MoveResourceRequest) {
    request = &MoveResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "MoveResource")
    
    
    return
}

func NewMoveResourceResponse() (response *MoveResourceResponse) {
    response = &MoveResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MoveResource
// 移动资源，支持跨个人或团队移动媒体以及分类。如果填写了Operator，则需要校验用户对媒体和分类资源的访问以及写权限。
//
// <li>当原始资源为媒体时，该接口效果为将该媒体移动到目标分类下面；</li>
//
// <li>当原始资源为分类时，该接口效果为将原始分类移动到目标分类或者是重命名。</li>
//
//  如果 SourceResource.Resource.Id = /素材/视频/NBA，DestinationResource.Resource.Id= /素材/视频/篮球 
//
// <li>当 DestinationResource.Resource.Id 不存在时候且原始资源与目标资源归属相同，操作结果为重命名原始分类；</li>
//
// <li>当 DestinationResource.Resource.Id 存在时候，操作结果为产生新目录 /素材/视频/篮球/NBA</li>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
func (c *Client) MoveResource(request *MoveResourceRequest) (response *MoveResourceResponse, err error) {
    return c.MoveResourceWithContext(context.Background(), request)
}

// MoveResource
// 移动资源，支持跨个人或团队移动媒体以及分类。如果填写了Operator，则需要校验用户对媒体和分类资源的访问以及写权限。
//
// <li>当原始资源为媒体时，该接口效果为将该媒体移动到目标分类下面；</li>
//
// <li>当原始资源为分类时，该接口效果为将原始分类移动到目标分类或者是重命名。</li>
//
//  如果 SourceResource.Resource.Id = /素材/视频/NBA，DestinationResource.Resource.Id= /素材/视频/篮球 
//
// <li>当 DestinationResource.Resource.Id 不存在时候且原始资源与目标资源归属相同，操作结果为重命名原始分类；</li>
//
// <li>当 DestinationResource.Resource.Id 存在时候，操作结果为产生新目录 /素材/视频/篮球/NBA</li>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_MATERIALID = "InvalidParameterValue.MaterialId"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
func (c *Client) MoveResourceWithContext(ctx context.Context, request *MoveResourceRequest) (response *MoveResourceResponse, err error) {
    if request == nil {
        request = NewMoveResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MoveResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewMoveResourceResponse()
    err = c.Send(request, response)
    return
}

func NewParseEventRequest() (request *ParseEventRequest) {
    request = &ParseEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cme", APIVersion, "ParseEvent")
    
    
    return
}

func NewParseEventResponse() (response *ParseEventResponse) {
    response = &ParseEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ParseEvent
// 该接口接受制作云回调给客户的事件内容，将其转化为对应的 EventContent 结构，请不要实际调用该接口，只需要将接收到的事件内容直接使用 JSON 解析到 EventContent  即可使用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) ParseEvent(request *ParseEventRequest) (response *ParseEventResponse, err error) {
    return c.ParseEventWithContext(context.Background(), request)
}

// ParseEvent
// 该接口接受制作云回调给客户的事件内容，将其转化为对应的 EventContent 结构，请不要实际调用该接口，只需要将接收到的事件内容直接使用 JSON 解析到 EventContent  即可使用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PLATFORM = "InvalidParameter.Platform"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
func (c *Client) ParseEventWithContext(ctx context.Context, request *ParseEventRequest) (response *ParseEventResponse, err error) {
    if request == nil {
        request = NewParseEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseEventResponse()
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

// RevokeResourceAuthorization
//  资源所属实体对目标实体撤销目标资源的相应权限，若原本没有相应权限则不产生变更。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RevokeResourceAuthorization(request *RevokeResourceAuthorizationRequest) (response *RevokeResourceAuthorizationResponse, err error) {
    return c.RevokeResourceAuthorizationWithContext(context.Background(), request)
}

// RevokeResourceAuthorization
//  资源所属实体对目标实体撤销目标资源的相应权限，若原本没有相应权限则不产生变更。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RevokeResourceAuthorizationWithContext(ctx context.Context, request *RevokeResourceAuthorizationRequest) (response *RevokeResourceAuthorizationResponse, err error) {
    if request == nil {
        request = NewRevokeResourceAuthorizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokeResourceAuthorization require credential")
    }

    request.SetContext(ctx)
    
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

// SearchMaterial
// 根据检索条件搜索媒体，返回媒体的基本信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_DATANOTFOUNDINDB = "InvalidParameterValue.DataNotFoundInDB"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchMaterial(request *SearchMaterialRequest) (response *SearchMaterialResponse, err error) {
    return c.SearchMaterialWithContext(context.Background(), request)
}

// SearchMaterial
// 根据检索条件搜索媒体，返回媒体的基本信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSPATH = "InvalidParameterValue.ClassPath"
//  INVALIDPARAMETERVALUE_DATANOTFOUNDINDB = "InvalidParameterValue.DataNotFoundInDB"
//  INVALIDPARAMETERVALUE_OWNERID = "InvalidParameterValue.OwnerId"
//  INVALIDPARAMETERVALUE_OWNERTYPE = "InvalidParameterValue.OwnerType"
//  INVALIDPARAMETERVALUE_PLATFORM = "InvalidParameterValue.Platform"
//  OPERATIONDENIED_PERMISSIONDENY = "OperationDenied.PermissionDeny"
//  RESOURCENOTFOUND_PLATFORM = "ResourceNotFound.Platform"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchMaterialWithContext(ctx context.Context, request *SearchMaterialRequest) (response *SearchMaterialResponse, err error) {
    if request == nil {
        request = NewSearchMaterialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchMaterial require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchMaterialResponse()
    err = c.Send(request, response)
    return
}
