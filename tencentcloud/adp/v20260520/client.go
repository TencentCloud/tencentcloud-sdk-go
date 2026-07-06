// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20260520

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2026-05-20"

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


func NewCopyAgentFromAppRequest() (request *CopyAgentFromAppRequest) {
    request = &CopyAgentFromAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CopyAgentFromApp")
    
    
    return
}

func NewCopyAgentFromAppResponse() (response *CopyAgentFromAppResponse) {
    response = &CopyAgentFromAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyAgentFromApp
// 创建Agent
func (c *Client) CopyAgentFromApp(request *CopyAgentFromAppRequest) (response *CopyAgentFromAppResponse, err error) {
    return c.CopyAgentFromAppWithContext(context.Background(), request)
}

// CopyAgentFromApp
// 创建Agent
func (c *Client) CopyAgentFromAppWithContext(ctx context.Context, request *CopyAgentFromAppRequest) (response *CopyAgentFromAppResponse, err error) {
    if request == nil {
        request = NewCopyAgentFromAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CopyAgentFromApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyAgentFromApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyAgentFromAppResponse()
    err = c.Send(request, response)
    return
}

func NewCopyAppRequest() (request *CopyAppRequest) {
    request = &CopyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CopyApp")
    
    
    return
}

func NewCopyAppResponse() (response *CopyAppResponse) {
    response = &CopyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyApp
// 复制应用
func (c *Client) CopyApp(request *CopyAppRequest) (response *CopyAppResponse, err error) {
    return c.CopyAppWithContext(context.Background(), request)
}

// CopyApp
// 复制应用
func (c *Client) CopyAppWithContext(ctx context.Context, request *CopyAppRequest) (response *CopyAppResponse, err error) {
    if request == nil {
        request = NewCopyAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CopyApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentRequest() (request *CreateAgentRequest) {
    request = &CreateAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateAgent")
    
    
    return
}

func NewCreateAgentResponse() (response *CreateAgentResponse) {
    response = &CreateAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgent
// 创建Agent
func (c *Client) CreateAgent(request *CreateAgentRequest) (response *CreateAgentResponse, err error) {
    return c.CreateAgentWithContext(context.Background(), request)
}

// CreateAgent
// 创建Agent
func (c *Client) CreateAgentWithContext(ctx context.Context, request *CreateAgentRequest) (response *CreateAgentResponse, err error) {
    if request == nil {
        request = NewCreateAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateApp")
    
    
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApp
// 创建应用
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    return c.CreateAppWithContext(context.Background(), request)
}

// CreateApp
// 创建应用
func (c *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConversationRequest() (request *CreateConversationRequest) {
    request = &CreateConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateConversation")
    
    
    return
}

func NewCreateConversationResponse() (response *CreateConversationResponse) {
    response = &CreateConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConversation
// 新建会话
func (c *Client) CreateConversation(request *CreateConversationRequest) (response *CreateConversationResponse, err error) {
    return c.CreateConversationWithContext(context.Background(), request)
}

// CreateConversation
// 新建会话
func (c *Client) CreateConversationWithContext(ctx context.Context, request *CreateConversationRequest) (response *CreateConversationResponse, err error) {
    if request == nil {
        request = NewCreateConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConversationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePluginRequest() (request *CreatePluginRequest) {
    request = &CreatePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreatePlugin")
    
    
    return
}

func NewCreatePluginResponse() (response *CreatePluginResponse) {
    response = &CreatePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePlugin
// 获取插件详情
func (c *Client) CreatePlugin(request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
    return c.CreatePluginWithContext(context.Background(), request)
}

// CreatePlugin
// 获取插件详情
func (c *Client) CreatePluginWithContext(ctx context.Context, request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
    if request == nil {
        request = NewCreatePluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreatePlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePluginResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReleaseRequest() (request *CreateReleaseRequest) {
    request = &CreateReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateRelease")
    
    
    return
}

func NewCreateReleaseResponse() (response *CreateReleaseResponse) {
    response = &CreateReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRelease
// 新增发布任务
func (c *Client) CreateRelease(request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    return c.CreateReleaseWithContext(context.Background(), request)
}

// CreateRelease
// 新增发布任务
func (c *Client) CreateReleaseWithContext(ctx context.Context, request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    if request == nil {
        request = NewCreateReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSkillRequest() (request *CreateSkillRequest) {
    request = &CreateSkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateSkill")
    
    
    return
}

func NewCreateSkillResponse() (response *CreateSkillResponse) {
    response = &CreateSkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSkill
// 创建skill
func (c *Client) CreateSkill(request *CreateSkillRequest) (response *CreateSkillResponse, err error) {
    return c.CreateSkillWithContext(context.Background(), request)
}

// CreateSkill
// 创建skill
func (c *Client) CreateSkillWithContext(ctx context.Context, request *CreateSkillRequest) (response *CreateSkillResponse, err error) {
    if request == nil {
        request = NewCreateSkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateSkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSkillResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSkillShareRequest() (request *CreateSkillShareRequest) {
    request = &CreateSkillShareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateSkillShare")
    
    
    return
}

func NewCreateSkillShareResponse() (response *CreateSkillShareResponse) {
    response = &CreateSkillShareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSkillShare
// 提交自定义Skill至企业级共享审批（两段式：提交→审批→回调创建共享任务）
func (c *Client) CreateSkillShare(request *CreateSkillShareRequest) (response *CreateSkillShareResponse, err error) {
    return c.CreateSkillShareWithContext(context.Background(), request)
}

// CreateSkillShare
// 提交自定义Skill至企业级共享审批（两段式：提交→审批→回调创建共享任务）
func (c *Client) CreateSkillShareWithContext(ctx context.Context, request *CreateSkillShareRequest) (response *CreateSkillShareResponse, err error) {
    if request == nil {
        request = NewCreateSkillShareRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateSkillShare")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSkillShare require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSkillShareResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSpaceRequest() (request *CreateSpaceRequest) {
    request = &CreateSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateSpace")
    
    
    return
}

func NewCreateSpaceResponse() (response *CreateSpaceResponse) {
    response = &CreateSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSpace
// 创建空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSpace(request *CreateSpaceRequest) (response *CreateSpaceResponse, err error) {
    return c.CreateSpaceWithContext(context.Background(), request)
}

// CreateSpace
// 创建空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSpaceWithContext(ctx context.Context, request *CreateSpaceRequest) (response *CreateSpaceResponse, err error) {
    if request == nil {
        request = NewCreateSpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateSpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSpaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVariableRequest() (request *CreateVariableRequest) {
    request = &CreateVariableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateVariable")
    
    
    return
}

func NewCreateVariableResponse() (response *CreateVariableResponse) {
    response = &CreateVariableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVariable
// 创建参数变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateVariable(request *CreateVariableRequest) (response *CreateVariableResponse, err error) {
    return c.CreateVariableWithContext(context.Background(), request)
}

// CreateVariable
// 创建参数变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateVariableWithContext(ctx context.Context, request *CreateVariableRequest) (response *CreateVariableResponse, err error) {
    if request == nil {
        request = NewCreateVariableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateVariable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVariable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVariableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWebSocketTokenRequest() (request *CreateWebSocketTokenRequest) {
    request = &CreateWebSocketTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateWebSocketToken")
    
    
    return
}

func NewCreateWebSocketTokenResponse() (response *CreateWebSocketTokenResponse) {
    response = &CreateWebSocketTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWebSocketToken
// 创建 WebSocket Token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWebSocketToken(request *CreateWebSocketTokenRequest) (response *CreateWebSocketTokenResponse, err error) {
    return c.CreateWebSocketTokenWithContext(context.Background(), request)
}

// CreateWebSocketToken
// 创建 WebSocket Token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWebSocketTokenWithContext(ctx context.Context, request *CreateWebSocketTokenRequest) (response *CreateWebSocketTokenResponse, err error) {
    if request == nil {
        request = NewCreateWebSocketTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateWebSocketToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWebSocketToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWebSocketTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceCredentialRequest() (request *CreateWorkspaceCredentialRequest) {
    request = &CreateWorkspaceCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateWorkspaceCredential")
    
    
    return
}

func NewCreateWorkspaceCredentialResponse() (response *CreateWorkspaceCredentialResponse) {
    response = &CreateWorkspaceCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkspaceCredential
// 创建工作空间凭证
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWorkspaceCredential(request *CreateWorkspaceCredentialRequest) (response *CreateWorkspaceCredentialResponse, err error) {
    return c.CreateWorkspaceCredentialWithContext(context.Background(), request)
}

// CreateWorkspaceCredential
// 创建工作空间凭证
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWorkspaceCredentialWithContext(ctx context.Context, request *CreateWorkspaceCredentialRequest) (response *CreateWorkspaceCredentialResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateWorkspaceCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaceCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentRequest() (request *DeleteAgentRequest) {
    request = &DeleteAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteAgent")
    
    
    return
}

func NewDeleteAgentResponse() (response *DeleteAgentResponse) {
    response = &DeleteAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAgent
// 删除Agent
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAgent(request *DeleteAgentRequest) (response *DeleteAgentResponse, err error) {
    return c.DeleteAgentWithContext(context.Background(), request)
}

// DeleteAgent
// 删除Agent
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAgentWithContext(ctx context.Context, request *DeleteAgentRequest) (response *DeleteAgentResponse, err error) {
    if request == nil {
        request = NewDeleteAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppRequest() (request *DeleteAppRequest) {
    request = &DeleteAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteApp")
    
    
    return
}

func NewDeleteAppResponse() (response *DeleteAppResponse) {
    response = &DeleteAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApp
// 删除应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteApp(request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
    return c.DeleteAppWithContext(context.Background(), request)
}

// DeleteApp
// 删除应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAppWithContext(ctx context.Context, request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
    if request == nil {
        request = NewDeleteAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConversationRequest() (request *DeleteConversationRequest) {
    request = &DeleteConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteConversation")
    
    
    return
}

func NewDeleteConversationResponse() (response *DeleteConversationResponse) {
    response = &DeleteConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConversation
// 删除会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteConversation(request *DeleteConversationRequest) (response *DeleteConversationResponse, err error) {
    return c.DeleteConversationWithContext(context.Background(), request)
}

// DeleteConversation
// 删除会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteConversationWithContext(ctx context.Context, request *DeleteConversationRequest) (response *DeleteConversationResponse, err error) {
    if request == nil {
        request = NewDeleteConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConversationResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePluginRequest() (request *DeletePluginRequest) {
    request = &DeletePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeletePlugin")
    
    
    return
}

func NewDeletePluginResponse() (response *DeletePluginResponse) {
    response = &DeletePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePlugin
// 修改插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeletePlugin(request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
    return c.DeletePluginWithContext(context.Background(), request)
}

// DeletePlugin
// 修改插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeletePluginWithContext(ctx context.Context, request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
    if request == nil {
        request = NewDeletePluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeletePlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePluginResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSkillRequest() (request *DeleteSkillRequest) {
    request = &DeleteSkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteSkill")
    
    
    return
}

func NewDeleteSkillResponse() (response *DeleteSkillResponse) {
    response = &DeleteSkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSkill
// 删除自定义 Skill  鉴权：创建者 ∨ (编辑权限 ∧ 删除权限） 拒绝场景：非 Custom 类型 / 已共享 / 安全检测中 / 上架审批中 / 下架审批中
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSkill(request *DeleteSkillRequest) (response *DeleteSkillResponse, err error) {
    return c.DeleteSkillWithContext(context.Background(), request)
}

// DeleteSkill
// 删除自定义 Skill  鉴权：创建者 ∨ (编辑权限 ∧ 删除权限） 拒绝场景：非 Custom 类型 / 已共享 / 安全检测中 / 上架审批中 / 下架审批中
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSkillWithContext(ctx context.Context, request *DeleteSkillRequest) (response *DeleteSkillResponse, err error) {
    if request == nil {
        request = NewDeleteSkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteSkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSkillResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSkillShareRequest() (request *DeleteSkillShareRequest) {
    request = &DeleteSkillShareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteSkillShare")
    
    
    return
}

func NewDeleteSkillShareResponse() (response *DeleteSkillShareResponse) {
    response = &DeleteSkillShareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSkillShare
// 提交共享 Skill 下架审批（v2，两段式：提交→审批→回调下架共享 Skill） 鉴权：删除权 拒绝场景：未共享 / 上架审批中 / 下架审批中
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSkillShare(request *DeleteSkillShareRequest) (response *DeleteSkillShareResponse, err error) {
    return c.DeleteSkillShareWithContext(context.Background(), request)
}

// DeleteSkillShare
// 提交共享 Skill 下架审批（v2，两段式：提交→审批→回调下架共享 Skill） 鉴权：删除权 拒绝场景：未共享 / 上架审批中 / 下架审批中
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSkillShareWithContext(ctx context.Context, request *DeleteSkillShareRequest) (response *DeleteSkillShareResponse, err error) {
    if request == nil {
        request = NewDeleteSkillShareRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteSkillShare")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSkillShare require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSkillShareResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSpaceRequest() (request *DeleteSpaceRequest) {
    request = &DeleteSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteSpace")
    
    
    return
}

func NewDeleteSpaceResponse() (response *DeleteSpaceResponse) {
    response = &DeleteSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSpace
// 删除空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSpace(request *DeleteSpaceRequest) (response *DeleteSpaceResponse, err error) {
    return c.DeleteSpaceWithContext(context.Background(), request)
}

// DeleteSpace
// 删除空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSpaceWithContext(ctx context.Context, request *DeleteSpaceRequest) (response *DeleteSpaceResponse, err error) {
    if request == nil {
        request = NewDeleteSpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteSpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSpaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVariableRequest() (request *DeleteVariableRequest) {
    request = &DeleteVariableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteVariable")
    
    
    return
}

func NewDeleteVariableResponse() (response *DeleteVariableResponse) {
    response = &DeleteVariableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVariable
// 删除参数变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVariable(request *DeleteVariableRequest) (response *DeleteVariableResponse, err error) {
    return c.DeleteVariableWithContext(context.Background(), request)
}

// DeleteVariable
// 删除参数变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVariableWithContext(ctx context.Context, request *DeleteVariableRequest) (response *DeleteVariableResponse, err error) {
    if request == nil {
        request = NewDeleteVariableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteVariable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVariable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVariableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentDetailRequest() (request *DescribeAgentDetailRequest) {
    request = &DescribeAgentDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAgentDetail")
    
    
    return
}

func NewDescribeAgentDetailResponse() (response *DescribeAgentDetailResponse) {
    response = &DescribeAgentDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentDetail
// 查询 Agent 详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAgentDetail(request *DescribeAgentDetailRequest) (response *DescribeAgentDetailResponse, err error) {
    return c.DescribeAgentDetailWithContext(context.Background(), request)
}

// DescribeAgentDetail
// 查询 Agent 详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAgentDetailWithContext(ctx context.Context, request *DescribeAgentDetailRequest) (response *DescribeAgentDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAgentDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAgentDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentReleasePreviewListRequest() (request *DescribeAgentReleasePreviewListRequest) {
    request = &DescribeAgentReleasePreviewListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAgentReleasePreviewList")
    
    
    return
}

func NewDescribeAgentReleasePreviewListResponse() (response *DescribeAgentReleasePreviewListResponse) {
    response = &DescribeAgentReleasePreviewListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentReleasePreviewList
// 获取应用下 Agent 的发布预览列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentReleasePreviewList(request *DescribeAgentReleasePreviewListRequest) (response *DescribeAgentReleasePreviewListResponse, err error) {
    return c.DescribeAgentReleasePreviewListWithContext(context.Background(), request)
}

// DescribeAgentReleasePreviewList
// 获取应用下 Agent 的发布预览列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentReleasePreviewListWithContext(ctx context.Context, request *DescribeAgentReleasePreviewListRequest) (response *DescribeAgentReleasePreviewListResponse, err error) {
    if request == nil {
        request = NewDescribeAgentReleasePreviewListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAgentReleasePreviewList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentReleasePreviewList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentReleasePreviewListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentSummaryListRequest() (request *DescribeAgentSummaryListRequest) {
    request = &DescribeAgentSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAgentSummaryList")
    
    
    return
}

func NewDescribeAgentSummaryListResponse() (response *DescribeAgentSummaryListResponse) {
    response = &DescribeAgentSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentSummaryList
// 查询 Agent 摘要信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentSummaryList(request *DescribeAgentSummaryListRequest) (response *DescribeAgentSummaryListResponse, err error) {
    return c.DescribeAgentSummaryListWithContext(context.Background(), request)
}

// DescribeAgentSummaryList
// 查询 Agent 摘要信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentSummaryListWithContext(ctx context.Context, request *DescribeAgentSummaryListRequest) (response *DescribeAgentSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeAgentSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAgentSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppRequest() (request *DescribeAppRequest) {
    request = &DescribeAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeApp")
    
    
    return
}

func NewDescribeAppResponse() (response *DescribeAppResponse) {
    response = &DescribeAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApp
// 获取应用信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApp(request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    return c.DescribeAppWithContext(context.Background(), request)
}

// DescribeApp
// 获取应用信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppWithContext(ctx context.Context, request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    if request == nil {
        request = NewDescribeAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppSummaryListRequest() (request *DescribeAppSummaryListRequest) {
    request = &DescribeAppSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppSummaryList")
    
    
    return
}

func NewDescribeAppSummaryListResponse() (response *DescribeAppSummaryListResponse) {
    response = &DescribeAppSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppSummaryList
// 获取应用摘要列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppSummaryList(request *DescribeAppSummaryListRequest) (response *DescribeAppSummaryListResponse, err error) {
    return c.DescribeAppSummaryListWithContext(context.Background(), request)
}

// DescribeAppSummaryList
// 获取应用摘要列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppSummaryListWithContext(ctx context.Context, request *DescribeAppSummaryListRequest) (response *DescribeAppSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeAppSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConversationRequest() (request *DescribeConversationRequest) {
    request = &DescribeConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConversation")
    
    
    return
}

func NewDescribeConversationResponse() (response *DescribeConversationResponse) {
    response = &DescribeConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConversation
// 查看会话信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversation(request *DescribeConversationRequest) (response *DescribeConversationResponse, err error) {
    return c.DescribeConversationWithContext(context.Background(), request)
}

// DescribeConversation
// 查看会话信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationWithContext(ctx context.Context, request *DescribeConversationRequest) (response *DescribeConversationResponse, err error) {
    if request == nil {
        request = NewDescribeConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConversationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConversationListRequest() (request *DescribeConversationListRequest) {
    request = &DescribeConversationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConversationList")
    
    
    return
}

func NewDescribeConversationListResponse() (response *DescribeConversationListResponse) {
    response = &DescribeConversationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConversationList
// 获取会话列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationList(request *DescribeConversationListRequest) (response *DescribeConversationListResponse, err error) {
    return c.DescribeConversationListWithContext(context.Background(), request)
}

// DescribeConversationList
// 获取会话列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationListWithContext(ctx context.Context, request *DescribeConversationListRequest) (response *DescribeConversationListResponse, err error) {
    if request == nil {
        request = NewDescribeConversationListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConversationList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConversationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConversationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConversationMessageListRequest() (request *DescribeConversationMessageListRequest) {
    request = &DescribeConversationMessageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConversationMessageList")
    
    
    return
}

func NewDescribeConversationMessageListResponse() (response *DescribeConversationMessageListResponse) {
    response = &DescribeConversationMessageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConversationMessageList
// 获取会话历史消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationMessageList(request *DescribeConversationMessageListRequest) (response *DescribeConversationMessageListResponse, err error) {
    return c.DescribeConversationMessageListWithContext(context.Background(), request)
}

// DescribeConversationMessageList
// 获取会话历史消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationMessageListWithContext(ctx context.Context, request *DescribeConversationMessageListRequest) (response *DescribeConversationMessageListResponse, err error) {
    if request == nil {
        request = NewDescribeConversationMessageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConversationMessageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConversationMessageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConversationMessageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLatestReleaseRequest() (request *DescribeLatestReleaseRequest) {
    request = &DescribeLatestReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeLatestRelease")
    
    
    return
}

func NewDescribeLatestReleaseResponse() (response *DescribeLatestReleaseResponse) {
    response = &DescribeLatestReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLatestRelease
// 拉取最新发布信息(包含发布时间、状态、渠道)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLatestRelease(request *DescribeLatestReleaseRequest) (response *DescribeLatestReleaseResponse, err error) {
    return c.DescribeLatestReleaseWithContext(context.Background(), request)
}

// DescribeLatestRelease
// 拉取最新发布信息(包含发布时间、状态、渠道)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLatestReleaseWithContext(ctx context.Context, request *DescribeLatestReleaseRequest) (response *DescribeLatestReleaseResponse, err error) {
    if request == nil {
        request = NewDescribeLatestReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeLatestRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLatestRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLatestReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelListRequest() (request *DescribeModelListRequest) {
    request = &DescribeModelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeModelList")
    
    
    return
}

func NewDescribeModelListResponse() (response *DescribeModelListResponse) {
    response = &DescribeModelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelList
// 查询模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeModelList(request *DescribeModelListRequest) (response *DescribeModelListResponse, err error) {
    return c.DescribeModelListWithContext(context.Background(), request)
}

// DescribeModelList
// 查询模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeModelListWithContext(ctx context.Context, request *DescribeModelListRequest) (response *DescribeModelListResponse, err error) {
    if request == nil {
        request = NewDescribeModelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeModelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginRequest() (request *DescribePluginRequest) {
    request = &DescribePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribePlugin")
    
    
    return
}

func NewDescribePluginResponse() (response *DescribePluginResponse) {
    response = &DescribePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePlugin
// 获取插件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePlugin(request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
    return c.DescribePluginWithContext(context.Background(), request)
}

// DescribePlugin
// 获取插件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePluginWithContext(ctx context.Context, request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
    if request == nil {
        request = NewDescribePluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribePlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginSummaryListRequest() (request *DescribePluginSummaryListRequest) {
    request = &DescribePluginSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribePluginSummaryList")
    
    
    return
}

func NewDescribePluginSummaryListResponse() (response *DescribePluginSummaryListResponse) {
    response = &DescribePluginSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePluginSummaryList
// 获取插件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePluginSummaryList(request *DescribePluginSummaryListRequest) (response *DescribePluginSummaryListResponse, err error) {
    return c.DescribePluginSummaryListWithContext(context.Background(), request)
}

// DescribePluginSummaryList
// 获取插件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePluginSummaryListWithContext(ctx context.Context, request *DescribePluginSummaryListRequest) (response *DescribePluginSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribePluginSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribePluginSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePluginSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseListRequest() (request *DescribeReleaseListRequest) {
    request = &DescribeReleaseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeReleaseList")
    
    
    return
}

func NewDescribeReleaseListResponse() (response *DescribeReleaseListResponse) {
    response = &DescribeReleaseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReleaseList
// 发布记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeReleaseList(request *DescribeReleaseListRequest) (response *DescribeReleaseListResponse, err error) {
    return c.DescribeReleaseListWithContext(context.Background(), request)
}

// DescribeReleaseList
// 发布记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeReleaseListWithContext(ctx context.Context, request *DescribeReleaseListRequest) (response *DescribeReleaseListResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeReleaseList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleaseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseSummaryRequest() (request *DescribeReleaseSummaryRequest) {
    request = &DescribeReleaseSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeReleaseSummary")
    
    
    return
}

func NewDescribeReleaseSummaryResponse() (response *DescribeReleaseSummaryResponse) {
    response = &DescribeReleaseSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReleaseSummary
// 查询发布任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeReleaseSummary(request *DescribeReleaseSummaryRequest) (response *DescribeReleaseSummaryResponse, err error) {
    return c.DescribeReleaseSummaryWithContext(context.Background(), request)
}

// DescribeReleaseSummary
// 查询发布任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeReleaseSummaryWithContext(ctx context.Context, request *DescribeReleaseSummaryRequest) (response *DescribeReleaseSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeReleaseSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleaseSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillCategoryListRequest() (request *DescribeSkillCategoryListRequest) {
    request = &DescribeSkillCategoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSkillCategoryList")
    
    
    return
}

func NewDescribeSkillCategoryListResponse() (response *DescribeSkillCategoryListResponse) {
    response = &DescribeSkillCategoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillCategoryList
// 查询 Skill 分类列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillCategoryList(request *DescribeSkillCategoryListRequest) (response *DescribeSkillCategoryListResponse, err error) {
    return c.DescribeSkillCategoryListWithContext(context.Background(), request)
}

// DescribeSkillCategoryList
// 查询 Skill 分类列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillCategoryListWithContext(ctx context.Context, request *DescribeSkillCategoryListRequest) (response *DescribeSkillCategoryListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillCategoryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSkillCategoryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillCategoryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillCategoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillDetailRequest() (request *DescribeSkillDetailRequest) {
    request = &DescribeSkillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSkillDetail")
    
    
    return
}

func NewDescribeSkillDetailResponse() (response *DescribeSkillDetailResponse) {
    response = &DescribeSkillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillDetail
// 查询skill详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillDetail(request *DescribeSkillDetailRequest) (response *DescribeSkillDetailResponse, err error) {
    return c.DescribeSkillDetailWithContext(context.Background(), request)
}

// DescribeSkillDetail
// 查询skill详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillDetailWithContext(ctx context.Context, request *DescribeSkillDetailRequest) (response *DescribeSkillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeSkillDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSkillDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillReferenceListRequest() (request *DescribeSkillReferenceListRequest) {
    request = &DescribeSkillReferenceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSkillReferenceList")
    
    
    return
}

func NewDescribeSkillReferenceListResponse() (response *DescribeSkillReferenceListResponse) {
    response = &DescribeSkillReferenceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillReferenceList
// 查询某个 Skill 被引用的详情列表（按 SkillRefType 分组：OpenClaw / cloud agent / 企业助手 agent） 鉴权：同 DescribeSkillDetail（能看该 Skill 即可查）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillReferenceList(request *DescribeSkillReferenceListRequest) (response *DescribeSkillReferenceListResponse, err error) {
    return c.DescribeSkillReferenceListWithContext(context.Background(), request)
}

// DescribeSkillReferenceList
// 查询某个 Skill 被引用的详情列表（按 SkillRefType 分组：OpenClaw / cloud agent / 企业助手 agent） 鉴权：同 DescribeSkillDetail（能看该 Skill 即可查）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillReferenceListWithContext(ctx context.Context, request *DescribeSkillReferenceListRequest) (response *DescribeSkillReferenceListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillReferenceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSkillReferenceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillReferenceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillReferenceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillSummaryListRequest() (request *DescribeSkillSummaryListRequest) {
    request = &DescribeSkillSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSkillSummaryList")
    
    
    return
}

func NewDescribeSkillSummaryListResponse() (response *DescribeSkillSummaryListResponse) {
    response = &DescribeSkillSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillSummaryList
// 查询 Skill 列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillSummaryList(request *DescribeSkillSummaryListRequest) (response *DescribeSkillSummaryListResponse, err error) {
    return c.DescribeSkillSummaryListWithContext(context.Background(), request)
}

// DescribeSkillSummaryList
// 查询 Skill 列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillSummaryListWithContext(ctx context.Context, request *DescribeSkillSummaryListRequest) (response *DescribeSkillSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSkillSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpaceListRequest() (request *DescribeSpaceListRequest) {
    request = &DescribeSpaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSpaceList")
    
    
    return
}

func NewDescribeSpaceListResponse() (response *DescribeSpaceListResponse) {
    response = &DescribeSpaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpaceList
// 获取空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSpaceList(request *DescribeSpaceListRequest) (response *DescribeSpaceListResponse, err error) {
    return c.DescribeSpaceListWithContext(context.Background(), request)
}

// DescribeSpaceList
// 获取空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSpaceListWithContext(ctx context.Context, request *DescribeSpaceListRequest) (response *DescribeSpaceListResponse, err error) {
    if request == nil {
        request = NewDescribeSpaceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSpaceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSystemVariableListRequest() (request *DescribeSystemVariableListRequest) {
    request = &DescribeSystemVariableListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSystemVariableList")
    
    
    return
}

func NewDescribeSystemVariableListResponse() (response *DescribeSystemVariableListResponse) {
    response = &DescribeSystemVariableListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSystemVariableList
// 获取系统变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSystemVariableList(request *DescribeSystemVariableListRequest) (response *DescribeSystemVariableListResponse, err error) {
    return c.DescribeSystemVariableListWithContext(context.Background(), request)
}

// DescribeSystemVariableList
// 获取系统变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSystemVariableListWithContext(ctx context.Context, request *DescribeSystemVariableListRequest) (response *DescribeSystemVariableListResponse, err error) {
    if request == nil {
        request = NewDescribeSystemVariableListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSystemVariableList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSystemVariableList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSystemVariableListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVariableRequest() (request *DescribeVariableRequest) {
    request = &DescribeVariableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeVariable")
    
    
    return
}

func NewDescribeVariableResponse() (response *DescribeVariableResponse) {
    response = &DescribeVariableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVariable
// 获取参数变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeVariable(request *DescribeVariableRequest) (response *DescribeVariableResponse, err error) {
    return c.DescribeVariableWithContext(context.Background(), request)
}

// DescribeVariable
// 获取参数变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeVariableWithContext(ctx context.Context, request *DescribeVariableRequest) (response *DescribeVariableResponse, err error) {
    if request == nil {
        request = NewDescribeVariableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeVariable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVariable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVariableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVariableListRequest() (request *DescribeVariableListRequest) {
    request = &DescribeVariableListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeVariableList")
    
    
    return
}

func NewDescribeVariableListResponse() (response *DescribeVariableListResponse) {
    response = &DescribeVariableListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVariableList
// 获取参数变量列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeVariableList(request *DescribeVariableListRequest) (response *DescribeVariableListResponse, err error) {
    return c.DescribeVariableListWithContext(context.Background(), request)
}

// DescribeVariableList
// 获取参数变量列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeVariableListWithContext(ctx context.Context, request *DescribeVariableListRequest) (response *DescribeVariableListResponse, err error) {
    if request == nil {
        request = NewDescribeVariableListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeVariableList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVariableList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVariableListResponse()
    err = c.Send(request, response)
    return
}

func NewFavoritePluginRequest() (request *FavoritePluginRequest) {
    request = &FavoritePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "FavoritePlugin")
    
    
    return
}

func NewFavoritePluginResponse() (response *FavoritePluginResponse) {
    response = &FavoritePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FavoritePlugin
// 收藏插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FavoritePlugin(request *FavoritePluginRequest) (response *FavoritePluginResponse, err error) {
    return c.FavoritePluginWithContext(context.Background(), request)
}

// FavoritePlugin
// 收藏插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FavoritePluginWithContext(ctx context.Context, request *FavoritePluginRequest) (response *FavoritePluginResponse, err error) {
    if request == nil {
        request = NewFavoritePluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "FavoritePlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FavoritePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewFavoritePluginResponse()
    err = c.Send(request, response)
    return
}

func NewFavoriteSkillRequest() (request *FavoriteSkillRequest) {
    request = &FavoriteSkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "FavoriteSkill")
    
    
    return
}

func NewFavoriteSkillResponse() (response *FavoriteSkillResponse) {
    response = &FavoriteSkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FavoriteSkill
// 收藏skill
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FavoriteSkill(request *FavoriteSkillRequest) (response *FavoriteSkillResponse, err error) {
    return c.FavoriteSkillWithContext(context.Background(), request)
}

// FavoriteSkill
// 收藏skill
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FavoriteSkillWithContext(ctx context.Context, request *FavoriteSkillRequest) (response *FavoriteSkillResponse, err error) {
    if request == nil {
        request = NewFavoriteSkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "FavoriteSkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FavoriteSkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewFavoriteSkillResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentRequest() (request *ModifyAgentRequest) {
    request = &ModifyAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyAgent")
    
    
    return
}

func NewModifyAgentResponse() (response *ModifyAgentResponse) {
    response = &ModifyAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgent
// 修改Agent配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyAgent(request *ModifyAgentRequest) (response *ModifyAgentResponse, err error) {
    return c.ModifyAgentWithContext(context.Background(), request)
}

// ModifyAgent
// 修改Agent配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyAgentWithContext(ctx context.Context, request *ModifyAgentRequest) (response *ModifyAgentResponse, err error) {
    if request == nil {
        request = NewModifyAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppRequest() (request *ModifyAppRequest) {
    request = &ModifyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyApp")
    
    
    return
}

func NewModifyAppResponse() (response *ModifyAppResponse) {
    response = &ModifyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApp
// 修改应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    return c.ModifyAppWithContext(context.Background(), request)
}

// ModifyApp
// 修改应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    if request == nil {
        request = NewModifyAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConversationRequest() (request *ModifyConversationRequest) {
    request = &ModifyConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyConversation")
    
    
    return
}

func NewModifyConversationResponse() (response *ModifyConversationResponse) {
    response = &ModifyConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConversation
// 修改会话信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyConversation(request *ModifyConversationRequest) (response *ModifyConversationResponse, err error) {
    return c.ModifyConversationWithContext(context.Background(), request)
}

// ModifyConversation
// 修改会话信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyConversationWithContext(ctx context.Context, request *ModifyConversationRequest) (response *ModifyConversationResponse, err error) {
    if request == nil {
        request = NewModifyConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConversationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPluginRequest() (request *ModifyPluginRequest) {
    request = &ModifyPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyPlugin")
    
    
    return
}

func NewModifyPluginResponse() (response *ModifyPluginResponse) {
    response = &ModifyPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPlugin
// 修改插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyPlugin(request *ModifyPluginRequest) (response *ModifyPluginResponse, err error) {
    return c.ModifyPluginWithContext(context.Background(), request)
}

// ModifyPlugin
// 修改插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyPluginWithContext(ctx context.Context, request *ModifyPluginRequest) (response *ModifyPluginResponse, err error) {
    if request == nil {
        request = NewModifyPluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyPlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPluginResponse()
    err = c.Send(request, response)
    return
}

func NewModifySkillRequest() (request *ModifySkillRequest) {
    request = &ModifySkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifySkill")
    
    
    return
}

func NewModifySkillResponse() (response *ModifySkillResponse) {
    response = &ModifySkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySkill
// Skill修改
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifySkill(request *ModifySkillRequest) (response *ModifySkillResponse, err error) {
    return c.ModifySkillWithContext(context.Background(), request)
}

// ModifySkill
// Skill修改
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifySkillWithContext(ctx context.Context, request *ModifySkillRequest) (response *ModifySkillResponse, err error) {
    if request == nil {
        request = NewModifySkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifySkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySkillResponse()
    err = c.Send(request, response)
    return
}

func NewModifySpaceRequest() (request *ModifySpaceRequest) {
    request = &ModifySpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifySpace")
    
    
    return
}

func NewModifySpaceResponse() (response *ModifySpaceResponse) {
    response = &ModifySpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySpace
// 编辑空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifySpace(request *ModifySpaceRequest) (response *ModifySpaceResponse, err error) {
    return c.ModifySpaceWithContext(context.Background(), request)
}

// ModifySpace
// 编辑空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifySpaceWithContext(ctx context.Context, request *ModifySpaceRequest) (response *ModifySpaceResponse, err error) {
    if request == nil {
        request = NewModifySpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifySpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySpaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVariableRequest() (request *ModifyVariableRequest) {
    request = &ModifyVariableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyVariable")
    
    
    return
}

func NewModifyVariableResponse() (response *ModifyVariableResponse) {
    response = &ModifyVariableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVariable
// 更新参数变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyVariable(request *ModifyVariableRequest) (response *ModifyVariableResponse, err error) {
    return c.ModifyVariableWithContext(context.Background(), request)
}

// ModifyVariable
// 更新参数变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyVariableWithContext(ctx context.Context, request *ModifyVariableRequest) (response *ModifyVariableResponse, err error) {
    if request == nil {
        request = NewModifyVariableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyVariable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVariable require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVariableResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseSkillRequest() (request *ReleaseSkillRequest) {
    request = &ReleaseSkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ReleaseSkill")
    
    
    return
}

func NewReleaseSkillResponse() (response *ReleaseSkillResponse) {
    response = &ReleaseSkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseSkill
// 上架skill
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ReleaseSkill(request *ReleaseSkillRequest) (response *ReleaseSkillResponse, err error) {
    return c.ReleaseSkillWithContext(context.Background(), request)
}

// ReleaseSkill
// 上架skill
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ReleaseSkillWithContext(ctx context.Context, request *ReleaseSkillRequest) (response *ReleaseSkillResponse, err error) {
    if request == nil {
        request = NewReleaseSkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ReleaseSkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseSkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseSkillResponse()
    err = c.Send(request, response)
    return
}

func NewResetConversationRequest() (request *ResetConversationRequest) {
    request = &ResetConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ResetConversation")
    
    
    return
}

func NewResetConversationResponse() (response *ResetConversationResponse) {
    response = &ResetConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetConversation
// 重置会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResetConversation(request *ResetConversationRequest) (response *ResetConversationResponse, err error) {
    return c.ResetConversationWithContext(context.Background(), request)
}

// ResetConversation
// 重置会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResetConversationWithContext(ctx context.Context, request *ResetConversationRequest) (response *ResetConversationResponse, err error) {
    if request == nil {
        request = NewResetConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ResetConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetConversationResponse()
    err = c.Send(request, response)
    return
}

func NewRetryReleaseRequest() (request *RetryReleaseRequest) {
    request = &RetryReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "RetryRelease")
    
    
    return
}

func NewRetryReleaseResponse() (response *RetryReleaseResponse) {
    response = &RetryReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryRelease
// 重试发布(发布暂停之后再次重新发布)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryRelease(request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    return c.RetryReleaseWithContext(context.Background(), request)
}

// RetryRelease
// 重试发布(发布暂停之后再次重新发布)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryReleaseWithContext(ctx context.Context, request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    if request == nil {
        request = NewRetryReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "RetryRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackReleaseRequest() (request *RollbackReleaseRequest) {
    request = &RollbackReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "RollbackRelease")
    
    
    return
}

func NewRollbackReleaseResponse() (response *RollbackReleaseResponse) {
    response = &RollbackReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackRelease
// 回滚发布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RollbackRelease(request *RollbackReleaseRequest) (response *RollbackReleaseResponse, err error) {
    return c.RollbackReleaseWithContext(context.Background(), request)
}

// RollbackRelease
// 回滚发布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RollbackReleaseWithContext(ctx context.Context, request *RollbackReleaseRequest) (response *RollbackReleaseResponse, err error) {
    if request == nil {
        request = NewRollbackReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "RollbackRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewUnfavoritePluginRequest() (request *UnfavoritePluginRequest) {
    request = &UnfavoritePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "UnfavoritePlugin")
    
    
    return
}

func NewUnfavoritePluginResponse() (response *UnfavoritePluginResponse) {
    response = &UnfavoritePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnfavoritePlugin
// 取消收藏插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnfavoritePlugin(request *UnfavoritePluginRequest) (response *UnfavoritePluginResponse, err error) {
    return c.UnfavoritePluginWithContext(context.Background(), request)
}

// UnfavoritePlugin
// 取消收藏插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnfavoritePluginWithContext(ctx context.Context, request *UnfavoritePluginRequest) (response *UnfavoritePluginResponse, err error) {
    if request == nil {
        request = NewUnfavoritePluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "UnfavoritePlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnfavoritePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnfavoritePluginResponse()
    err = c.Send(request, response)
    return
}

func NewUnfavoriteSkillRequest() (request *UnfavoriteSkillRequest) {
    request = &UnfavoriteSkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "UnfavoriteSkill")
    
    
    return
}

func NewUnfavoriteSkillResponse() (response *UnfavoriteSkillResponse) {
    response = &UnfavoriteSkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnfavoriteSkill
// 取消收藏skill
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnfavoriteSkill(request *UnfavoriteSkillRequest) (response *UnfavoriteSkillResponse, err error) {
    return c.UnfavoriteSkillWithContext(context.Background(), request)
}

// UnfavoriteSkill
// 取消收藏skill
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnfavoriteSkillWithContext(ctx context.Context, request *UnfavoriteSkillRequest) (response *UnfavoriteSkillResponse, err error) {
    if request == nil {
        request = NewUnfavoriteSkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "UnfavoriteSkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnfavoriteSkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnfavoriteSkillResponse()
    err = c.Send(request, response)
    return
}
