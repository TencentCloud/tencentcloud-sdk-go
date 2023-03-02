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

package v20210524

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-05-24"

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


func NewCreateCustomizeTemplatesRequest() (request *CreateCustomizeTemplatesRequest) {
    request = &CreateCustomizeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "CreateCustomizeTemplates")
    
    
    return
}

func NewCreateCustomizeTemplatesResponse() (response *CreateCustomizeTemplatesResponse) {
    response = &CreateCustomizeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomizeTemplates
// 添加自定义模板
func (c *Client) CreateCustomizeTemplates(request *CreateCustomizeTemplatesRequest) (response *CreateCustomizeTemplatesResponse, err error) {
    return c.CreateCustomizeTemplatesWithContext(context.Background(), request)
}

// CreateCustomizeTemplates
// 添加自定义模板
func (c *Client) CreateCustomizeTemplatesWithContext(ctx context.Context, request *CreateCustomizeTemplatesRequest) (response *CreateCustomizeTemplatesResponse, err error) {
    if request == nil {
        request = NewCreateCustomizeTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomizeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomizeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceByAgentRequest() (request *CreateWorkspaceByAgentRequest) {
    request = &CreateWorkspaceByAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "CreateWorkspaceByAgent")
    
    
    return
}

func NewCreateWorkspaceByAgentResponse() (response *CreateWorkspaceByAgentResponse) {
    response = &CreateWorkspaceByAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWorkspaceByAgent
// 云服务器方式创建工作空间
func (c *Client) CreateWorkspaceByAgent(request *CreateWorkspaceByAgentRequest) (response *CreateWorkspaceByAgentResponse, err error) {
    return c.CreateWorkspaceByAgentWithContext(context.Background(), request)
}

// CreateWorkspaceByAgent
// 云服务器方式创建工作空间
func (c *Client) CreateWorkspaceByAgentWithContext(ctx context.Context, request *CreateWorkspaceByAgentRequest) (response *CreateWorkspaceByAgentResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceByAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaceByAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceByAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceByTemplateRequest() (request *CreateWorkspaceByTemplateRequest) {
    request = &CreateWorkspaceByTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "CreateWorkspaceByTemplate")
    
    
    return
}

func NewCreateWorkspaceByTemplateResponse() (response *CreateWorkspaceByTemplateResponse) {
    response = &CreateWorkspaceByTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWorkspaceByTemplate
// 快速开始, 基于模板创建工作空间
func (c *Client) CreateWorkspaceByTemplate(request *CreateWorkspaceByTemplateRequest) (response *CreateWorkspaceByTemplateResponse, err error) {
    return c.CreateWorkspaceByTemplateWithContext(context.Background(), request)
}

// CreateWorkspaceByTemplate
// 快速开始, 基于模板创建工作空间
func (c *Client) CreateWorkspaceByTemplateWithContext(ctx context.Context, request *CreateWorkspaceByTemplateRequest) (response *CreateWorkspaceByTemplateResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceByTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaceByTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceByTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceByVersionControlRequest() (request *CreateWorkspaceByVersionControlRequest) {
    request = &CreateWorkspaceByVersionControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "CreateWorkspaceByVersionControl")
    
    
    return
}

func NewCreateWorkspaceByVersionControlResponse() (response *CreateWorkspaceByVersionControlResponse) {
    response = &CreateWorkspaceByVersionControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWorkspaceByVersionControl
// 根据模板创建工作空间
func (c *Client) CreateWorkspaceByVersionControl(request *CreateWorkspaceByVersionControlRequest) (response *CreateWorkspaceByVersionControlResponse, err error) {
    return c.CreateWorkspaceByVersionControlWithContext(context.Background(), request)
}

// CreateWorkspaceByVersionControl
// 根据模板创建工作空间
func (c *Client) CreateWorkspaceByVersionControlWithContext(ctx context.Context, request *CreateWorkspaceByVersionControlRequest) (response *CreateWorkspaceByVersionControlResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceByVersionControlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaceByVersionControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceByVersionControlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceTemporaryTokenRequest() (request *CreateWorkspaceTemporaryTokenRequest) {
    request = &CreateWorkspaceTemporaryTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "CreateWorkspaceTemporaryToken")
    
    
    return
}

func NewCreateWorkspaceTemporaryTokenResponse() (response *CreateWorkspaceTemporaryTokenResponse) {
    response = &CreateWorkspaceTemporaryTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWorkspaceTemporaryToken
// 为工作空间创建临时访问凭证，重复调用会创建新的 Token，旧的 Token 将会自动失效
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateWorkspaceTemporaryToken(request *CreateWorkspaceTemporaryTokenRequest) (response *CreateWorkspaceTemporaryTokenResponse, err error) {
    return c.CreateWorkspaceTemporaryTokenWithContext(context.Background(), request)
}

// CreateWorkspaceTemporaryToken
// 为工作空间创建临时访问凭证，重复调用会创建新的 Token，旧的 Token 将会自动失效
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateWorkspaceTemporaryTokenWithContext(ctx context.Context, request *CreateWorkspaceTemporaryTokenRequest) (response *CreateWorkspaceTemporaryTokenResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceTemporaryTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaceTemporaryToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceTemporaryTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomizeTemplatesByIdRequest() (request *DeleteCustomizeTemplatesByIdRequest) {
    request = &DeleteCustomizeTemplatesByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DeleteCustomizeTemplatesById")
    
    
    return
}

func NewDeleteCustomizeTemplatesByIdResponse() (response *DeleteCustomizeTemplatesByIdResponse) {
    response = &DeleteCustomizeTemplatesByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCustomizeTemplatesById
// 删除自定义模板
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCustomizeTemplatesById(request *DeleteCustomizeTemplatesByIdRequest) (response *DeleteCustomizeTemplatesByIdResponse, err error) {
    return c.DeleteCustomizeTemplatesByIdWithContext(context.Background(), request)
}

// DeleteCustomizeTemplatesById
// 删除自定义模板
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCustomizeTemplatesByIdWithContext(ctx context.Context, request *DeleteCustomizeTemplatesByIdRequest) (response *DeleteCustomizeTemplatesByIdResponse, err error) {
    if request == nil {
        request = NewDeleteCustomizeTemplatesByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomizeTemplatesById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomizeTemplatesByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomizeTemplatesRequest() (request *DescribeCustomizeTemplatesRequest) {
    request = &DescribeCustomizeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeCustomizeTemplates")
    
    
    return
}

func NewDescribeCustomizeTemplatesResponse() (response *DescribeCustomizeTemplatesResponse) {
    response = &DescribeCustomizeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomizeTemplates
// 获取所有模板列表
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCustomizeTemplates(request *DescribeCustomizeTemplatesRequest) (response *DescribeCustomizeTemplatesResponse, err error) {
    return c.DescribeCustomizeTemplatesWithContext(context.Background(), request)
}

// DescribeCustomizeTemplates
// 获取所有模板列表
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCustomizeTemplatesWithContext(ctx context.Context, request *DescribeCustomizeTemplatesRequest) (response *DescribeCustomizeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeCustomizeTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomizeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomizeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomizeTemplatesByIdRequest() (request *DescribeCustomizeTemplatesByIdRequest) {
    request = &DescribeCustomizeTemplatesByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeCustomizeTemplatesById")
    
    
    return
}

func NewDescribeCustomizeTemplatesByIdResponse() (response *DescribeCustomizeTemplatesByIdResponse) {
    response = &DescribeCustomizeTemplatesByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomizeTemplatesById
// 获取特定模板信息
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCustomizeTemplatesById(request *DescribeCustomizeTemplatesByIdRequest) (response *DescribeCustomizeTemplatesByIdResponse, err error) {
    return c.DescribeCustomizeTemplatesByIdWithContext(context.Background(), request)
}

// DescribeCustomizeTemplatesById
// 获取特定模板信息
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCustomizeTemplatesByIdWithContext(ctx context.Context, request *DescribeCustomizeTemplatesByIdRequest) (response *DescribeCustomizeTemplatesByIdResponse, err error) {
    if request == nil {
        request = NewDescribeCustomizeTemplatesByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomizeTemplatesById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomizeTemplatesByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomizeTemplatesPresetsRequest() (request *DescribeCustomizeTemplatesPresetsRequest) {
    request = &DescribeCustomizeTemplatesPresetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeCustomizeTemplatesPresets")
    
    
    return
}

func NewDescribeCustomizeTemplatesPresetsResponse() (response *DescribeCustomizeTemplatesPresetsResponse) {
    response = &DescribeCustomizeTemplatesPresetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomizeTemplatesPresets
// 获取创建模板的预置参数
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCustomizeTemplatesPresets(request *DescribeCustomizeTemplatesPresetsRequest) (response *DescribeCustomizeTemplatesPresetsResponse, err error) {
    return c.DescribeCustomizeTemplatesPresetsWithContext(context.Background(), request)
}

// DescribeCustomizeTemplatesPresets
// 获取创建模板的预置参数
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCustomizeTemplatesPresetsWithContext(ctx context.Context, request *DescribeCustomizeTemplatesPresetsRequest) (response *DescribeCustomizeTemplatesPresetsResponse, err error) {
    if request == nil {
        request = NewDescribeCustomizeTemplatesPresetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomizeTemplatesPresets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomizeTemplatesPresetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspaceEnvListRequest() (request *DescribeWorkspaceEnvListRequest) {
    request = &DescribeWorkspaceEnvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeWorkspaceEnvList")
    
    
    return
}

func NewDescribeWorkspaceEnvListResponse() (response *DescribeWorkspaceEnvListResponse) {
    response = &DescribeWorkspaceEnvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWorkspaceEnvList
// 环境列表接口返回信息
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeWorkspaceEnvList(request *DescribeWorkspaceEnvListRequest) (response *DescribeWorkspaceEnvListResponse, err error) {
    return c.DescribeWorkspaceEnvListWithContext(context.Background(), request)
}

// DescribeWorkspaceEnvList
// 环境列表接口返回信息
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeWorkspaceEnvListWithContext(ctx context.Context, request *DescribeWorkspaceEnvListRequest) (response *DescribeWorkspaceEnvListResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspaceEnvListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaceEnvList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspaceEnvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspaceNameExistRequest() (request *DescribeWorkspaceNameExistRequest) {
    request = &DescribeWorkspaceNameExistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeWorkspaceNameExist")
    
    
    return
}

func NewDescribeWorkspaceNameExistResponse() (response *DescribeWorkspaceNameExistResponse) {
    response = &DescribeWorkspaceNameExistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWorkspaceNameExist
// 检查工作空间是否存在
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeWorkspaceNameExist(request *DescribeWorkspaceNameExistRequest) (response *DescribeWorkspaceNameExistResponse, err error) {
    return c.DescribeWorkspaceNameExistWithContext(context.Background(), request)
}

// DescribeWorkspaceNameExist
// 检查工作空间是否存在
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeWorkspaceNameExistWithContext(ctx context.Context, request *DescribeWorkspaceNameExistRequest) (response *DescribeWorkspaceNameExistResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspaceNameExistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaceNameExist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspaceNameExistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspaceStatusRequest() (request *DescribeWorkspaceStatusRequest) {
    request = &DescribeWorkspaceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeWorkspaceStatus")
    
    
    return
}

func NewDescribeWorkspaceStatusResponse() (response *DescribeWorkspaceStatusResponse) {
    response = &DescribeWorkspaceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWorkspaceStatus
// 获取工作空间元信息
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeWorkspaceStatus(request *DescribeWorkspaceStatusRequest) (response *DescribeWorkspaceStatusResponse, err error) {
    return c.DescribeWorkspaceStatusWithContext(context.Background(), request)
}

// DescribeWorkspaceStatus
// 获取工作空间元信息
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeWorkspaceStatusWithContext(ctx context.Context, request *DescribeWorkspaceStatusRequest) (response *DescribeWorkspaceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspaceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspaceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspaceStatusListRequest() (request *DescribeWorkspaceStatusListRequest) {
    request = &DescribeWorkspaceStatusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeWorkspaceStatusList")
    
    
    return
}

func NewDescribeWorkspaceStatusListResponse() (response *DescribeWorkspaceStatusListResponse) {
    response = &DescribeWorkspaceStatusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWorkspaceStatusList
// 获取用户工作空间列表
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeWorkspaceStatusList(request *DescribeWorkspaceStatusListRequest) (response *DescribeWorkspaceStatusListResponse, err error) {
    return c.DescribeWorkspaceStatusListWithContext(context.Background(), request)
}

// DescribeWorkspaceStatusList
// 获取用户工作空间列表
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeWorkspaceStatusListWithContext(ctx context.Context, request *DescribeWorkspaceStatusListRequest) (response *DescribeWorkspaceStatusListResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspaceStatusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaceStatusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspaceStatusListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomizeTemplateVersionControlRequest() (request *ModifyCustomizeTemplateVersionControlRequest) {
    request = &ModifyCustomizeTemplateVersionControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "ModifyCustomizeTemplateVersionControl")
    
    
    return
}

func NewModifyCustomizeTemplateVersionControlResponse() (response *ModifyCustomizeTemplateVersionControlResponse) {
    response = &ModifyCustomizeTemplateVersionControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCustomizeTemplateVersionControl
// 修改模板默认代码仓库
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCustomizeTemplateVersionControl(request *ModifyCustomizeTemplateVersionControlRequest) (response *ModifyCustomizeTemplateVersionControlResponse, err error) {
    return c.ModifyCustomizeTemplateVersionControlWithContext(context.Background(), request)
}

// ModifyCustomizeTemplateVersionControl
// 修改模板默认代码仓库
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCustomizeTemplateVersionControlWithContext(ctx context.Context, request *ModifyCustomizeTemplateVersionControlRequest) (response *ModifyCustomizeTemplateVersionControlResponse, err error) {
    if request == nil {
        request = NewModifyCustomizeTemplateVersionControlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomizeTemplateVersionControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomizeTemplateVersionControlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomizeTemplatesFullByIdRequest() (request *ModifyCustomizeTemplatesFullByIdRequest) {
    request = &ModifyCustomizeTemplatesFullByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "ModifyCustomizeTemplatesFullById")
    
    
    return
}

func NewModifyCustomizeTemplatesFullByIdResponse() (response *ModifyCustomizeTemplatesFullByIdResponse) {
    response = &ModifyCustomizeTemplatesFullByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCustomizeTemplatesFullById
// 全量修改自定义模板，不忽略空
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCustomizeTemplatesFullById(request *ModifyCustomizeTemplatesFullByIdRequest) (response *ModifyCustomizeTemplatesFullByIdResponse, err error) {
    return c.ModifyCustomizeTemplatesFullByIdWithContext(context.Background(), request)
}

// ModifyCustomizeTemplatesFullById
// 全量修改自定义模板，不忽略空
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCustomizeTemplatesFullByIdWithContext(ctx context.Context, request *ModifyCustomizeTemplatesFullByIdRequest) (response *ModifyCustomizeTemplatesFullByIdResponse, err error) {
    if request == nil {
        request = NewModifyCustomizeTemplatesFullByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomizeTemplatesFullById require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomizeTemplatesFullByIdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomizeTemplatesPartByIdRequest() (request *ModifyCustomizeTemplatesPartByIdRequest) {
    request = &ModifyCustomizeTemplatesPartByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "ModifyCustomizeTemplatesPartById")
    
    
    return
}

func NewModifyCustomizeTemplatesPartByIdResponse() (response *ModifyCustomizeTemplatesPartByIdResponse) {
    response = &ModifyCustomizeTemplatesPartByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCustomizeTemplatesPartById
// 全量修改自定义模板，忽略空
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCustomizeTemplatesPartById(request *ModifyCustomizeTemplatesPartByIdRequest) (response *ModifyCustomizeTemplatesPartByIdResponse, err error) {
    return c.ModifyCustomizeTemplatesPartByIdWithContext(context.Background(), request)
}

// ModifyCustomizeTemplatesPartById
// 全量修改自定义模板，忽略空
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCustomizeTemplatesPartByIdWithContext(ctx context.Context, request *ModifyCustomizeTemplatesPartByIdRequest) (response *ModifyCustomizeTemplatesPartByIdResponse, err error) {
    if request == nil {
        request = NewModifyCustomizeTemplatesPartByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomizeTemplatesPartById require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomizeTemplatesPartByIdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkspaceAttributesRequest() (request *ModifyWorkspaceAttributesRequest) {
    request = &ModifyWorkspaceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "ModifyWorkspaceAttributes")
    
    
    return
}

func NewModifyWorkspaceAttributesResponse() (response *ModifyWorkspaceAttributesResponse) {
    response = &ModifyWorkspaceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWorkspaceAttributes
// 修改工作空间的名称和描述
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyWorkspaceAttributes(request *ModifyWorkspaceAttributesRequest) (response *ModifyWorkspaceAttributesResponse, err error) {
    return c.ModifyWorkspaceAttributesWithContext(context.Background(), request)
}

// ModifyWorkspaceAttributes
// 修改工作空间的名称和描述
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyWorkspaceAttributesWithContext(ctx context.Context, request *ModifyWorkspaceAttributesRequest) (response *ModifyWorkspaceAttributesResponse, err error) {
    if request == nil {
        request = NewModifyWorkspaceAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkspaceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkspaceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverWorkspaceRequest() (request *RecoverWorkspaceRequest) {
    request = &RecoverWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "RecoverWorkspace")
    
    
    return
}

func NewRecoverWorkspaceResponse() (response *RecoverWorkspaceResponse) {
    response = &RecoverWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecoverWorkspace
// 恢复工作空间
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RecoverWorkspace(request *RecoverWorkspaceRequest) (response *RecoverWorkspaceResponse, err error) {
    return c.RecoverWorkspaceWithContext(context.Background(), request)
}

// RecoverWorkspace
// 恢复工作空间
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RecoverWorkspaceWithContext(ctx context.Context, request *RecoverWorkspaceRequest) (response *RecoverWorkspaceResponse, err error) {
    if request == nil {
        request = NewRecoverWorkspaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveWorkspaceRequest() (request *RemoveWorkspaceRequest) {
    request = &RemoveWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "RemoveWorkspace")
    
    
    return
}

func NewRemoveWorkspaceResponse() (response *RemoveWorkspaceResponse) {
    response = &RemoveWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveWorkspace
// 删除工作空间
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RemoveWorkspace(request *RemoveWorkspaceRequest) (response *RemoveWorkspaceResponse, err error) {
    return c.RemoveWorkspaceWithContext(context.Background(), request)
}

// RemoveWorkspace
// 删除工作空间
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RemoveWorkspaceWithContext(ctx context.Context, request *RemoveWorkspaceRequest) (response *RemoveWorkspaceResponse, err error) {
    if request == nil {
        request = NewRemoveWorkspaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewRunWorkspaceRequest() (request *RunWorkspaceRequest) {
    request = &RunWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "RunWorkspace")
    
    
    return
}

func NewRunWorkspaceResponse() (response *RunWorkspaceResponse) {
    response = &RunWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunWorkspace
// 运行空间
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RunWorkspace(request *RunWorkspaceRequest) (response *RunWorkspaceResponse, err error) {
    return c.RunWorkspaceWithContext(context.Background(), request)
}

// RunWorkspace
// 运行空间
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RunWorkspaceWithContext(ctx context.Context, request *RunWorkspaceRequest) (response *RunWorkspaceResponse, err error) {
    if request == nil {
        request = NewRunWorkspaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewStopWorkspaceRequest() (request *StopWorkspaceRequest) {
    request = &StopWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "StopWorkspace")
    
    
    return
}

func NewStopWorkspaceResponse() (response *StopWorkspaceResponse) {
    response = &StopWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopWorkspace
// 停止运行空间
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) StopWorkspace(request *StopWorkspaceRequest) (response *StopWorkspaceResponse, err error) {
    return c.StopWorkspaceWithContext(context.Background(), request)
}

// StopWorkspace
// 停止运行空间
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) StopWorkspaceWithContext(ctx context.Context, request *StopWorkspaceRequest) (response *StopWorkspaceResponse, err error) {
    if request == nil {
        request = NewStopWorkspaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopWorkspaceResponse()
    err = c.Send(request, response)
    return
}
