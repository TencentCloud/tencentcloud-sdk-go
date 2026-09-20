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


func NewCheckLabelRequest() (request *CheckLabelRequest) {
    request = &CheckLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CheckLabel")
    
    
    return
}

func NewCheckLabelResponse() (response *CheckLabelResponse) {
    response = &CheckLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckLabel
// 校验标签
func (c *Client) CheckLabel(request *CheckLabelRequest) (response *CheckLabelResponse, err error) {
    return c.CheckLabelWithContext(context.Background(), request)
}

// CheckLabel
// 校验标签
func (c *Client) CheckLabelWithContext(ctx context.Context, request *CheckLabelRequest) (response *CheckLabelResponse, err error) {
    if request == nil {
        request = NewCheckLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CheckLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckLabelResponse()
    err = c.Send(request, response)
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
// 复制 Agent（目前仅支持claw模式））
func (c *Client) CopyAgentFromApp(request *CopyAgentFromAppRequest) (response *CopyAgentFromAppResponse, err error) {
    return c.CopyAgentFromAppWithContext(context.Background(), request)
}

// CopyAgentFromApp
// 复制 Agent（目前仅支持claw模式））
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
// 创建Agent（目前仅支持claw模式）
func (c *Client) CreateAgent(request *CreateAgentRequest) (response *CreateAgentResponse, err error) {
    return c.CreateAgentWithContext(context.Background(), request)
}

// CreateAgent
// 创建Agent（目前仅支持claw模式）
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

func NewCreateAppTriggerRequest() (request *CreateAppTriggerRequest) {
    request = &CreateAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateAppTrigger")
    
    
    return
}

func NewCreateAppTriggerResponse() (response *CreateAppTriggerResponse) {
    response = &CreateAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAppTrigger
// CreateAppTrigger
func (c *Client) CreateAppTrigger(request *CreateAppTriggerRequest) (response *CreateAppTriggerResponse, err error) {
    return c.CreateAppTriggerWithContext(context.Background(), request)
}

// CreateAppTrigger
// CreateAppTrigger
func (c *Client) CreateAppTriggerWithContext(ctx context.Context, request *CreateAppTriggerRequest) (response *CreateAppTriggerResponse, err error) {
    if request == nil {
        request = NewCreateAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCategoryRequest() (request *CreateCategoryRequest) {
    request = &CreateCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateCategory")
    
    
    return
}

func NewCreateCategoryResponse() (response *CreateCategoryResponse) {
    response = &CreateCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCategory
// 创建分类
func (c *Client) CreateCategory(request *CreateCategoryRequest) (response *CreateCategoryResponse, err error) {
    return c.CreateCategoryWithContext(context.Background(), request)
}

// CreateCategory
// 创建分类
func (c *Client) CreateCategoryWithContext(ctx context.Context, request *CreateCategoryRequest) (response *CreateCategoryResponse, err error) {
    if request == nil {
        request = NewCreateCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChannelRequest() (request *CreateChannelRequest) {
    request = &CreateChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateChannel")
    
    
    return
}

func NewCreateChannelResponse() (response *CreateChannelResponse) {
    response = &CreateChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateChannel
// 创建渠道（通过scene区分B端应用发布渠道与C端IM渠道）
func (c *Client) CreateChannel(request *CreateChannelRequest) (response *CreateChannelResponse, err error) {
    return c.CreateChannelWithContext(context.Background(), request)
}

// CreateChannel
// 创建渠道（通过scene区分B端应用发布渠道与C端IM渠道）
func (c *Client) CreateChannelWithContext(ctx context.Context, request *CreateChannelRequest) (response *CreateChannelResponse, err error) {
    if request == nil {
        request = NewCreateChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChannelResponse()
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

func NewCreateKBRequest() (request *CreateKBRequest) {
    request = &CreateKBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateKB")
    
    
    return
}

func NewCreateKBResponse() (response *CreateKBResponse) {
    response = &CreateKBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKB
// 创建知识库
func (c *Client) CreateKB(request *CreateKBRequest) (response *CreateKBResponse, err error) {
    return c.CreateKBWithContext(context.Background(), request)
}

// CreateKB
// 创建知识库
func (c *Client) CreateKBWithContext(ctx context.Context, request *CreateKBRequest) (response *CreateKBResponse, err error) {
    if request == nil {
        request = NewCreateKBRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateKB")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKB require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKBResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLabelRequest() (request *CreateLabelRequest) {
    request = &CreateLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateLabel")
    
    
    return
}

func NewCreateLabelResponse() (response *CreateLabelResponse) {
    response = &CreateLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLabel
// 创建标签
func (c *Client) CreateLabel(request *CreateLabelRequest) (response *CreateLabelResponse, err error) {
    return c.CreateLabelWithContext(context.Background(), request)
}

// CreateLabel
// 创建标签
func (c *Client) CreateLabelWithContext(ctx context.Context, request *CreateLabelRequest) (response *CreateLabelResponse, err error) {
    if request == nil {
        request = NewCreateLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLabelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMsgRecordCategoryRequest() (request *CreateMsgRecordCategoryRequest) {
    request = &CreateMsgRecordCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateMsgRecordCategory")
    
    
    return
}

func NewCreateMsgRecordCategoryResponse() (response *CreateMsgRecordCategoryResponse) {
    response = &CreateMsgRecordCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMsgRecordCategory
// 创建一条消息记录分类，支持指定分类名称与父分类（ParentId 为 0 时表示一级分类）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateMsgRecordCategory(request *CreateMsgRecordCategoryRequest) (response *CreateMsgRecordCategoryResponse, err error) {
    return c.CreateMsgRecordCategoryWithContext(context.Background(), request)
}

// CreateMsgRecordCategory
// 创建一条消息记录分类，支持指定分类名称与父分类（ParentId 为 0 时表示一级分类）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateMsgRecordCategoryWithContext(ctx context.Context, request *CreateMsgRecordCategoryRequest) (response *CreateMsgRecordCategoryResponse, err error) {
    if request == nil {
        request = NewCreateMsgRecordCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateMsgRecordCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMsgRecordCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMsgRecordCategoryResponse()
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
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreatePlugin(request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
    return c.CreatePluginWithContext(context.Background(), request)
}

// CreatePlugin
// 获取插件详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewCreateQAGenerationTaskRequest() (request *CreateQAGenerationTaskRequest) {
    request = &CreateQAGenerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateQAGenerationTask")
    
    
    return
}

func NewCreateQAGenerationTaskResponse() (response *CreateQAGenerationTaskResponse) {
    response = &CreateQAGenerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQAGenerationTask
// 创建 QA 生成任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateQAGenerationTask(request *CreateQAGenerationTaskRequest) (response *CreateQAGenerationTaskResponse, err error) {
    return c.CreateQAGenerationTaskWithContext(context.Background(), request)
}

// CreateQAGenerationTask
// 创建 QA 生成任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateQAGenerationTaskWithContext(ctx context.Context, request *CreateQAGenerationTaskRequest) (response *CreateQAGenerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateQAGenerationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateQAGenerationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQAGenerationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQAGenerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQAListRequest() (request *CreateQAListRequest) {
    request = &CreateQAListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateQAList")
    
    
    return
}

func NewCreateQAListResponse() (response *CreateQAListResponse) {
    response = &CreateQAListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQAList
// 批量创建 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateQAList(request *CreateQAListRequest) (response *CreateQAListResponse, err error) {
    return c.CreateQAListWithContext(context.Background(), request)
}

// CreateQAList
// 批量创建 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateQAListWithContext(ctx context.Context, request *CreateQAListRequest) (response *CreateQAListResponse, err error) {
    if request == nil {
        request = NewCreateQAListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateQAList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQAList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQAListResponse()
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
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRelease(request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    return c.CreateReleaseWithContext(context.Background(), request)
}

// CreateRelease
// 新增发布任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewCreateSimilarQuestionRequest() (request *CreateSimilarQuestionRequest) {
    request = &CreateSimilarQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateSimilarQuestion")
    
    
    return
}

func NewCreateSimilarQuestionResponse() (response *CreateSimilarQuestionResponse) {
    response = &CreateSimilarQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSimilarQuestion
// 创建相似问生成任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSimilarQuestion(request *CreateSimilarQuestionRequest) (response *CreateSimilarQuestionResponse, err error) {
    return c.CreateSimilarQuestionWithContext(context.Background(), request)
}

// CreateSimilarQuestion
// 创建相似问生成任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSimilarQuestionWithContext(ctx context.Context, request *CreateSimilarQuestionRequest) (response *CreateSimilarQuestionResponse, err error) {
    if request == nil {
        request = NewCreateSimilarQuestionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateSimilarQuestion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSimilarQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSimilarQuestionResponse()
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
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSkill(request *CreateSkillRequest) (response *CreateSkillResponse, err error) {
    return c.CreateSkillWithContext(context.Background(), request)
}

// CreateSkill
// 创建skill
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSkillShare(request *CreateSkillShareRequest) (response *CreateSkillShareResponse, err error) {
    return c.CreateSkillShareWithContext(context.Background(), request)
}

// CreateSkillShare
// 提交自定义Skill至企业级共享审批（两段式：提交→审批→回调创建共享任务）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewDeleteAppTriggerRequest() (request *DeleteAppTriggerRequest) {
    request = &DeleteAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteAppTrigger")
    
    
    return
}

func NewDeleteAppTriggerResponse() (response *DeleteAppTriggerResponse) {
    response = &DeleteAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAppTrigger
// DeleteAppTrigger
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAppTrigger(request *DeleteAppTriggerRequest) (response *DeleteAppTriggerResponse, err error) {
    return c.DeleteAppTriggerWithContext(context.Background(), request)
}

// DeleteAppTrigger
// DeleteAppTrigger
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAppTriggerWithContext(ctx context.Context, request *DeleteAppTriggerRequest) (response *DeleteAppTriggerResponse, err error) {
    if request == nil {
        request = NewDeleteAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCategoryRequest() (request *DeleteCategoryRequest) {
    request = &DeleteCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteCategory")
    
    
    return
}

func NewDeleteCategoryResponse() (response *DeleteCategoryResponse) {
    response = &DeleteCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCategory
// 删除分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCategory(request *DeleteCategoryRequest) (response *DeleteCategoryResponse, err error) {
    return c.DeleteCategoryWithContext(context.Background(), request)
}

// DeleteCategory
// 删除分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCategoryWithContext(ctx context.Context, request *DeleteCategoryRequest) (response *DeleteCategoryResponse, err error) {
    if request == nil {
        request = NewDeleteCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteChannelRequest() (request *DeleteChannelRequest) {
    request = &DeleteChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteChannel")
    
    
    return
}

func NewDeleteChannelResponse() (response *DeleteChannelResponse) {
    response = &DeleteChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteChannel
// 删除渠道（通过scene区分场景）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteChannel(request *DeleteChannelRequest) (response *DeleteChannelResponse, err error) {
    return c.DeleteChannelWithContext(context.Background(), request)
}

// DeleteChannel
// 删除渠道（通过scene区分场景）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteChannelWithContext(ctx context.Context, request *DeleteChannelRequest) (response *DeleteChannelResponse, err error) {
    if request == nil {
        request = NewDeleteChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteChannelResponse()
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

func NewDeleteDocListRequest() (request *DeleteDocListRequest) {
    request = &DeleteDocListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteDocList")
    
    
    return
}

func NewDeleteDocListResponse() (response *DeleteDocListResponse) {
    response = &DeleteDocListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDocList
// 批量删除文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDocList(request *DeleteDocListRequest) (response *DeleteDocListResponse, err error) {
    return c.DeleteDocListWithContext(context.Background(), request)
}

// DeleteDocList
// 批量删除文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDocListWithContext(ctx context.Context, request *DeleteDocListRequest) (response *DeleteDocListResponse, err error) {
    if request == nil {
        request = NewDeleteDocListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteDocList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDocList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKBRequest() (request *DeleteKBRequest) {
    request = &DeleteKBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteKB")
    
    
    return
}

func NewDeleteKBResponse() (response *DeleteKBResponse) {
    response = &DeleteKBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteKB
// 删除知识库
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteKB(request *DeleteKBRequest) (response *DeleteKBResponse, err error) {
    return c.DeleteKBWithContext(context.Background(), request)
}

// DeleteKB
// 删除知识库
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteKBWithContext(ctx context.Context, request *DeleteKBRequest) (response *DeleteKBResponse, err error) {
    if request == nil {
        request = NewDeleteKBRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteKB")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKB require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKBResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLabelListRequest() (request *DeleteLabelListRequest) {
    request = &DeleteLabelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteLabelList")
    
    
    return
}

func NewDeleteLabelListResponse() (response *DeleteLabelListResponse) {
    response = &DeleteLabelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLabelList
// 批量删除标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteLabelList(request *DeleteLabelListRequest) (response *DeleteLabelListResponse, err error) {
    return c.DeleteLabelListWithContext(context.Background(), request)
}

// DeleteLabelList
// 批量删除标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteLabelListWithContext(ctx context.Context, request *DeleteLabelListRequest) (response *DeleteLabelListResponse, err error) {
    if request == nil {
        request = NewDeleteLabelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteLabelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLabelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLabelListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMsgRecordCategoryRequest() (request *DeleteMsgRecordCategoryRequest) {
    request = &DeleteMsgRecordCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteMsgRecordCategory")
    
    
    return
}

func NewDeleteMsgRecordCategoryResponse() (response *DeleteMsgRecordCategoryResponse) {
    response = &DeleteMsgRecordCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMsgRecordCategory
// 删除指定的消息记录分类
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMsgRecordCategory(request *DeleteMsgRecordCategoryRequest) (response *DeleteMsgRecordCategoryResponse, err error) {
    return c.DeleteMsgRecordCategoryWithContext(context.Background(), request)
}

// DeleteMsgRecordCategory
// 删除指定的消息记录分类
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMsgRecordCategoryWithContext(ctx context.Context, request *DeleteMsgRecordCategoryRequest) (response *DeleteMsgRecordCategoryResponse, err error) {
    if request == nil {
        request = NewDeleteMsgRecordCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteMsgRecordCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMsgRecordCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMsgRecordCategoryResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePlugin(request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
    return c.DeletePluginWithContext(context.Background(), request)
}

// DeletePlugin
// 修改插件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewDeleteQAListRequest() (request *DeleteQAListRequest) {
    request = &DeleteQAListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteQAList")
    
    
    return
}

func NewDeleteQAListResponse() (response *DeleteQAListResponse) {
    response = &DeleteQAListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQAList
// 批量删除 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteQAList(request *DeleteQAListRequest) (response *DeleteQAListResponse, err error) {
    return c.DeleteQAListWithContext(context.Background(), request)
}

// DeleteQAList
// 批量删除 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteQAListWithContext(ctx context.Context, request *DeleteQAListRequest) (response *DeleteQAListResponse, err error) {
    if request == nil {
        request = NewDeleteQAListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteQAList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQAList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQAListResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSkill(request *DeleteSkillRequest) (response *DeleteSkillResponse, err error) {
    return c.DeleteSkillWithContext(context.Background(), request)
}

// DeleteSkill
// 删除自定义 Skill  鉴权：创建者 ∨ (编辑权限 ∧ 删除权限） 拒绝场景：非 Custom 类型 / 已共享 / 安全检测中 / 上架审批中 / 下架审批中
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSkillShare(request *DeleteSkillShareRequest) (response *DeleteSkillShareResponse, err error) {
    return c.DeleteSkillShareWithContext(context.Background(), request)
}

// DeleteSkillShare
// 提交共享 Skill 下架审批（v2，两段式：提交→审批→回调下架共享 Skill） 鉴权：删除权 拒绝场景：未共享 / 上架审批中 / 下架审批中
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewDescribeAccountListRequest() (request *DescribeAccountListRequest) {
    request = &DescribeAccountListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAccountList")
    
    
    return
}

func NewDescribeAccountListResponse() (response *DescribeAccountListResponse) {
    response = &DescribeAccountListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountList
// 查看企业下的员工列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAccountList(request *DescribeAccountListRequest) (response *DescribeAccountListResponse, err error) {
    return c.DescribeAccountListWithContext(context.Background(), request)
}

// DescribeAccountList
// 查看企业下的员工列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAccountListWithContext(ctx context.Context, request *DescribeAccountListRequest) (response *DescribeAccountListResponse, err error) {
    if request == nil {
        request = NewDescribeAccountListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAccountList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountListResponse()
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

func NewDescribeAppStatisticsOverviewRequest() (request *DescribeAppStatisticsOverviewRequest) {
    request = &DescribeAppStatisticsOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppStatisticsOverview")
    
    
    return
}

func NewDescribeAppStatisticsOverviewResponse() (response *DescribeAppStatisticsOverviewResponse) {
    response = &DescribeAppStatisticsOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppStatisticsOverview
// 查询应用视图下的调用统计总览，包含总调用次数、调用成功率、总tokens平均耗时及首tokens平均耗时；RAG 应用额外返回各回复方式的调用次数及占比，用于绘制饼图
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppStatisticsOverview(request *DescribeAppStatisticsOverviewRequest) (response *DescribeAppStatisticsOverviewResponse, err error) {
    return c.DescribeAppStatisticsOverviewWithContext(context.Background(), request)
}

// DescribeAppStatisticsOverview
// 查询应用视图下的调用统计总览，包含总调用次数、调用成功率、总tokens平均耗时及首tokens平均耗时；RAG 应用额外返回各回复方式的调用次数及占比，用于绘制饼图
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppStatisticsOverviewWithContext(ctx context.Context, request *DescribeAppStatisticsOverviewRequest) (response *DescribeAppStatisticsOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeAppStatisticsOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppStatisticsOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppStatisticsOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppStatisticsOverviewResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppSummaryList(request *DescribeAppSummaryListRequest) (response *DescribeAppSummaryListResponse, err error) {
    return c.DescribeAppSummaryListWithContext(context.Background(), request)
}

// DescribeAppSummaryList
// 获取应用摘要列表
//
// 可能返回的错误码:
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

func NewDescribeAppTriggerRequest() (request *DescribeAppTriggerRequest) {
    request = &DescribeAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppTrigger")
    
    
    return
}

func NewDescribeAppTriggerResponse() (response *DescribeAppTriggerResponse) {
    response = &DescribeAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppTrigger
// DescribeAppTrigger
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTrigger(request *DescribeAppTriggerRequest) (response *DescribeAppTriggerResponse, err error) {
    return c.DescribeAppTriggerWithContext(context.Background(), request)
}

// DescribeAppTrigger
// DescribeAppTrigger
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerWithContext(ctx context.Context, request *DescribeAppTriggerRequest) (response *DescribeAppTriggerResponse, err error) {
    if request == nil {
        request = NewDescribeAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppTriggerInstanceRequest() (request *DescribeAppTriggerInstanceRequest) {
    request = &DescribeAppTriggerInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppTriggerInstance")
    
    
    return
}

func NewDescribeAppTriggerInstanceResponse() (response *DescribeAppTriggerInstanceResponse) {
    response = &DescribeAppTriggerInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppTriggerInstance
// DescribeAppTriggerInstance
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerInstance(request *DescribeAppTriggerInstanceRequest) (response *DescribeAppTriggerInstanceResponse, err error) {
    return c.DescribeAppTriggerInstanceWithContext(context.Background(), request)
}

// DescribeAppTriggerInstance
// DescribeAppTriggerInstance
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerInstanceWithContext(ctx context.Context, request *DescribeAppTriggerInstanceRequest) (response *DescribeAppTriggerInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeAppTriggerInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppTriggerInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppTriggerInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppTriggerInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppTriggerRunLogListRequest() (request *DescribeAppTriggerRunLogListRequest) {
    request = &DescribeAppTriggerRunLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppTriggerRunLogList")
    
    
    return
}

func NewDescribeAppTriggerRunLogListResponse() (response *DescribeAppTriggerRunLogListResponse) {
    response = &DescribeAppTriggerRunLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppTriggerRunLogList
// DescribeAppTriggerRunLogList
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerRunLogList(request *DescribeAppTriggerRunLogListRequest) (response *DescribeAppTriggerRunLogListResponse, err error) {
    return c.DescribeAppTriggerRunLogListWithContext(context.Background(), request)
}

// DescribeAppTriggerRunLogList
// DescribeAppTriggerRunLogList
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerRunLogListWithContext(ctx context.Context, request *DescribeAppTriggerRunLogListRequest) (response *DescribeAppTriggerRunLogListResponse, err error) {
    if request == nil {
        request = NewDescribeAppTriggerRunLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppTriggerRunLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppTriggerRunLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppTriggerRunLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppTriggerSummaryListRequest() (request *DescribeAppTriggerSummaryListRequest) {
    request = &DescribeAppTriggerSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppTriggerSummaryList")
    
    
    return
}

func NewDescribeAppTriggerSummaryListResponse() (response *DescribeAppTriggerSummaryListResponse) {
    response = &DescribeAppTriggerSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppTriggerSummaryList
// DescribeAppTriggerSummaryList
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerSummaryList(request *DescribeAppTriggerSummaryListRequest) (response *DescribeAppTriggerSummaryListResponse, err error) {
    return c.DescribeAppTriggerSummaryListWithContext(context.Background(), request)
}

// DescribeAppTriggerSummaryList
// DescribeAppTriggerSummaryList
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerSummaryListWithContext(ctx context.Context, request *DescribeAppTriggerSummaryListRequest) (response *DescribeAppTriggerSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeAppTriggerSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppTriggerSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppTriggerSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppTriggerSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogListRequest() (request *DescribeAuditLogListRequest) {
    request = &DescribeAuditLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAuditLogList")
    
    
    return
}

func NewDescribeAuditLogListResponse() (response *DescribeAuditLogListResponse) {
    response = &DescribeAuditLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditLogList
// 查看操作日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAuditLogList(request *DescribeAuditLogListRequest) (response *DescribeAuditLogListResponse, err error) {
    return c.DescribeAuditLogListWithContext(context.Background(), request)
}

// DescribeAuditLogList
// 查看操作日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAuditLogListWithContext(ctx context.Context, request *DescribeAuditLogListRequest) (response *DescribeAuditLogListResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAuditLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogMetaRequest() (request *DescribeAuditLogMetaRequest) {
    request = &DescribeAuditLogMetaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAuditLogMeta")
    
    
    return
}

func NewDescribeAuditLogMetaResponse() (response *DescribeAuditLogMetaResponse) {
    response = &DescribeAuditLogMetaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditLogMeta
// 获取审计日志元信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAuditLogMeta(request *DescribeAuditLogMetaRequest) (response *DescribeAuditLogMetaResponse, err error) {
    return c.DescribeAuditLogMetaWithContext(context.Background(), request)
}

// DescribeAuditLogMeta
// 获取审计日志元信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAuditLogMetaWithContext(ctx context.Context, request *DescribeAuditLogMetaRequest) (response *DescribeAuditLogMetaResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogMetaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAuditLogMeta")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLogMeta require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogMetaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCategoryListRequest() (request *DescribeCategoryListRequest) {
    request = &DescribeCategoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeCategoryList")
    
    
    return
}

func NewDescribeCategoryListResponse() (response *DescribeCategoryListResponse) {
    response = &DescribeCategoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCategoryList
// 查询分类列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCategoryList(request *DescribeCategoryListRequest) (response *DescribeCategoryListResponse, err error) {
    return c.DescribeCategoryListWithContext(context.Background(), request)
}

// DescribeCategoryList
// 查询分类列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCategoryListWithContext(ctx context.Context, request *DescribeCategoryListRequest) (response *DescribeCategoryListResponse, err error) {
    if request == nil {
        request = NewDescribeCategoryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeCategoryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCategoryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCategoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelRequest() (request *DescribeChannelRequest) {
    request = &DescribeChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeChannel")
    
    
    return
}

func NewDescribeChannelResponse() (response *DescribeChannelResponse) {
    response = &DescribeChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChannel
// 获取渠道详情（scene区分场景）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeChannel(request *DescribeChannelRequest) (response *DescribeChannelResponse, err error) {
    return c.DescribeChannelWithContext(context.Background(), request)
}

// DescribeChannel
// 获取渠道详情（scene区分场景）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeChannelWithContext(ctx context.Context, request *DescribeChannelRequest) (response *DescribeChannelResponse, err error) {
    if request == nil {
        request = NewDescribeChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelListRequest() (request *DescribeChannelListRequest) {
    request = &DescribeChannelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeChannelList")
    
    
    return
}

func NewDescribeChannelListResponse() (response *DescribeChannelListResponse) {
    response = &DescribeChannelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChannelList
// 获取渠道列表（scene区分场景）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeChannelList(request *DescribeChannelListRequest) (response *DescribeChannelListResponse, err error) {
    return c.DescribeChannelListWithContext(context.Background(), request)
}

// DescribeChannelList
// 获取渠道列表（scene区分场景）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeChannelListWithContext(ctx context.Context, request *DescribeChannelListRequest) (response *DescribeChannelListResponse, err error) {
    if request == nil {
        request = NewDescribeChannelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeChannelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChannelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChannelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrencyLimitDetailListRequest() (request *DescribeConcurrencyLimitDetailListRequest) {
    request = &DescribeConcurrencyLimitDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConcurrencyLimitDetailList")
    
    
    return
}

func NewDescribeConcurrencyLimitDetailListResponse() (response *DescribeConcurrencyLimitDetailListResponse) {
    response = &DescribeConcurrencyLimitDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrencyLimitDetailList
// 查询并发超限明细，包含QPM/TPM超限与专属并发超限记录，返回超限发生时间、空间、应用、模型及请求内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConcurrencyLimitDetailList(request *DescribeConcurrencyLimitDetailListRequest) (response *DescribeConcurrencyLimitDetailListResponse, err error) {
    return c.DescribeConcurrencyLimitDetailListWithContext(context.Background(), request)
}

// DescribeConcurrencyLimitDetailList
// 查询并发超限明细，包含QPM/TPM超限与专属并发超限记录，返回超限发生时间、空间、应用、模型及请求内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConcurrencyLimitDetailListWithContext(ctx context.Context, request *DescribeConcurrencyLimitDetailListRequest) (response *DescribeConcurrencyLimitDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrencyLimitDetailListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConcurrencyLimitDetailList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrencyLimitDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrencyLimitDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConflictQARequest() (request *DescribeConflictQARequest) {
    request = &DescribeConflictQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConflictQA")
    
    
    return
}

func NewDescribeConflictQAResponse() (response *DescribeConflictQAResponse) {
    response = &DescribeConflictQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConflictQA
// 查询冲突问详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConflictQA(request *DescribeConflictQARequest) (response *DescribeConflictQAResponse, err error) {
    return c.DescribeConflictQAWithContext(context.Background(), request)
}

// DescribeConflictQA
// 查询冲突问详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConflictQAWithContext(ctx context.Context, request *DescribeConflictQARequest) (response *DescribeConflictQAResponse, err error) {
    if request == nil {
        request = NewDescribeConflictQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConflictQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConflictQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConflictQAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConflictQASummaryListRequest() (request *DescribeConflictQASummaryListRequest) {
    request = &DescribeConflictQASummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConflictQASummaryList")
    
    
    return
}

func NewDescribeConflictQASummaryListResponse() (response *DescribeConflictQASummaryListResponse) {
    response = &DescribeConflictQASummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConflictQASummaryList
// 查询冲突问列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConflictQASummaryList(request *DescribeConflictQASummaryListRequest) (response *DescribeConflictQASummaryListResponse, err error) {
    return c.DescribeConflictQASummaryListWithContext(context.Background(), request)
}

// DescribeConflictQASummaryList
// 查询冲突问列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConflictQASummaryListWithContext(ctx context.Context, request *DescribeConflictQASummaryListRequest) (response *DescribeConflictQASummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeConflictQASummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConflictQASummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConflictQASummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConflictQASummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumptionDetailListRequest() (request *DescribeConsumptionDetailListRequest) {
    request = &DescribeConsumptionDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConsumptionDetailList")
    
    
    return
}

func NewDescribeConsumptionDetailListResponse() (response *DescribeConsumptionDetailListResponse) {
    response = &DescribeConsumptionDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumptionDetailList
// 查询资源消耗明细，包含计费相关字段（消耗类型、消耗目标、消耗场景、套餐包及PU消耗等）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConsumptionDetailList(request *DescribeConsumptionDetailListRequest) (response *DescribeConsumptionDetailListResponse, err error) {
    return c.DescribeConsumptionDetailListWithContext(context.Background(), request)
}

// DescribeConsumptionDetailList
// 查询资源消耗明细，包含计费相关字段（消耗类型、消耗目标、消耗场景、套餐包及PU消耗等）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConsumptionDetailListWithContext(ctx context.Context, request *DescribeConsumptionDetailListRequest) (response *DescribeConsumptionDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeConsumptionDetailListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConsumptionDetailList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumptionDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumptionDetailListResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversation(request *DescribeConversationRequest) (response *DescribeConversationResponse, err error) {
    return c.DescribeConversationWithContext(context.Background(), request)
}

// DescribeConversation
// 查看会话信息
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationList(request *DescribeConversationListRequest) (response *DescribeConversationListResponse, err error) {
    return c.DescribeConversationListWithContext(context.Background(), request)
}

// DescribeConversationList
// 获取会话列表
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationMessageList(request *DescribeConversationMessageListRequest) (response *DescribeConversationMessageListResponse, err error) {
    return c.DescribeConversationMessageListWithContext(context.Background(), request)
}

// DescribeConversationMessageList
// 获取会话历史消息
//
// 可能返回的错误码:
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

func NewDescribeDocRequest() (request *DescribeDocRequest) {
    request = &DescribeDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeDoc")
    
    
    return
}

func NewDescribeDocResponse() (response *DescribeDocResponse) {
    response = &DescribeDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDoc
// 查询文档详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDoc(request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    return c.DescribeDocWithContext(context.Background(), request)
}

// DescribeDoc
// 查询文档详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDocWithContext(ctx context.Context, request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    if request == nil {
        request = NewDescribeDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocSummaryListRequest() (request *DescribeDocSummaryListRequest) {
    request = &DescribeDocSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeDocSummaryList")
    
    
    return
}

func NewDescribeDocSummaryListResponse() (response *DescribeDocSummaryListResponse) {
    response = &DescribeDocSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDocSummaryList
// 查询文档摘要列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDocSummaryList(request *DescribeDocSummaryListRequest) (response *DescribeDocSummaryListResponse, err error) {
    return c.DescribeDocSummaryListWithContext(context.Background(), request)
}

// DescribeDocSummaryList
// 查询文档摘要列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDocSummaryListWithContext(ctx context.Context, request *DescribeDocSummaryListRequest) (response *DescribeDocSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeDocSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeDocSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDocSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKBRequest() (request *DescribeKBRequest) {
    request = &DescribeKBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeKB")
    
    
    return
}

func NewDescribeKBResponse() (response *DescribeKBResponse) {
    response = &DescribeKBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKB
// 查询知识库详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKB(request *DescribeKBRequest) (response *DescribeKBResponse, err error) {
    return c.DescribeKBWithContext(context.Background(), request)
}

// DescribeKB
// 查询知识库详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKBWithContext(ctx context.Context, request *DescribeKBRequest) (response *DescribeKBResponse, err error) {
    if request == nil {
        request = NewDescribeKBRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeKB")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKB require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKBResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKBSummaryListRequest() (request *DescribeKBSummaryListRequest) {
    request = &DescribeKBSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeKBSummaryList")
    
    
    return
}

func NewDescribeKBSummaryListResponse() (response *DescribeKBSummaryListResponse) {
    response = &DescribeKBSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKBSummaryList
// 查询知识库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKBSummaryList(request *DescribeKBSummaryListRequest) (response *DescribeKBSummaryListResponse, err error) {
    return c.DescribeKBSummaryListWithContext(context.Background(), request)
}

// DescribeKBSummaryList
// 查询知识库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKBSummaryListWithContext(ctx context.Context, request *DescribeKBSummaryListRequest) (response *DescribeKBSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeKBSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeKBSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKBSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKBSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLabelRequest() (request *DescribeLabelRequest) {
    request = &DescribeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeLabel")
    
    
    return
}

func NewDescribeLabelResponse() (response *DescribeLabelResponse) {
    response = &DescribeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLabel
// 查询标签详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLabel(request *DescribeLabelRequest) (response *DescribeLabelResponse, err error) {
    return c.DescribeLabelWithContext(context.Background(), request)
}

// DescribeLabel
// 查询标签详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLabelWithContext(ctx context.Context, request *DescribeLabelRequest) (response *DescribeLabelResponse, err error) {
    if request == nil {
        request = NewDescribeLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLabelSummaryListRequest() (request *DescribeLabelSummaryListRequest) {
    request = &DescribeLabelSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeLabelSummaryList")
    
    
    return
}

func NewDescribeLabelSummaryListResponse() (response *DescribeLabelSummaryListResponse) {
    response = &DescribeLabelSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLabelSummaryList
// 查询标签列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLabelSummaryList(request *DescribeLabelSummaryListRequest) (response *DescribeLabelSummaryListResponse, err error) {
    return c.DescribeLabelSummaryListWithContext(context.Background(), request)
}

// DescribeLabelSummaryList
// 查询标签列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLabelSummaryListWithContext(ctx context.Context, request *DescribeLabelSummaryListRequest) (response *DescribeLabelSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeLabelSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeLabelSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLabelSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLabelSummaryListResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLatestRelease(request *DescribeLatestReleaseRequest) (response *DescribeLatestReleaseResponse, err error) {
    return c.DescribeLatestReleaseWithContext(context.Background(), request)
}

// DescribeLatestRelease
// 拉取最新发布信息(包含发布时间、状态、渠道)
//
// 可能返回的错误码:
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

func NewDescribeMetricOverviewListRequest() (request *DescribeMetricOverviewListRequest) {
    request = &DescribeMetricOverviewListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeMetricOverviewList")
    
    
    return
}

func NewDescribeMetricOverviewListResponse() (response *DescribeMetricOverviewListResponse) {
    response = &DescribeMetricOverviewListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricOverviewList
// 查询看板总览KPI卡片数据，通过resource_type区分资源看板与业务看板域，返回各域KPI指标列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricOverviewList(request *DescribeMetricOverviewListRequest) (response *DescribeMetricOverviewListResponse, err error) {
    return c.DescribeMetricOverviewListWithContext(context.Background(), request)
}

// DescribeMetricOverviewList
// 查询看板总览KPI卡片数据，通过resource_type区分资源看板与业务看板域，返回各域KPI指标列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricOverviewListWithContext(ctx context.Context, request *DescribeMetricOverviewListRequest) (response *DescribeMetricOverviewListResponse, err error) {
    if request == nil {
        request = NewDescribeMetricOverviewListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeMetricOverviewList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricOverviewList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricOverviewListResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeModelList(request *DescribeModelListRequest) (response *DescribeModelListResponse, err error) {
    return c.DescribeModelListWithContext(context.Background(), request)
}

// DescribeModelList
// 查询模型列表
//
// 可能返回的错误码:
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

func NewDescribeMsgRecordCategoryListRequest() (request *DescribeMsgRecordCategoryListRequest) {
    request = &DescribeMsgRecordCategoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeMsgRecordCategoryList")
    
    
    return
}

func NewDescribeMsgRecordCategoryListResponse() (response *DescribeMsgRecordCategoryListResponse) {
    response = &DescribeMsgRecordCategoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMsgRecordCategoryList
// 查询应用的消息记录分类树，返回分类及子分类、各分类下记录数量与操作权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMsgRecordCategoryList(request *DescribeMsgRecordCategoryListRequest) (response *DescribeMsgRecordCategoryListResponse, err error) {
    return c.DescribeMsgRecordCategoryListWithContext(context.Background(), request)
}

// DescribeMsgRecordCategoryList
// 查询应用的消息记录分类树，返回分类及子分类、各分类下记录数量与操作权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMsgRecordCategoryListWithContext(ctx context.Context, request *DescribeMsgRecordCategoryListRequest) (response *DescribeMsgRecordCategoryListResponse, err error) {
    if request == nil {
        request = NewDescribeMsgRecordCategoryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeMsgRecordCategoryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMsgRecordCategoryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMsgRecordCategoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsgRecordListRequest() (request *DescribeMsgRecordListRequest) {
    request = &DescribeMsgRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeMsgRecordList")
    
    
    return
}

func NewDescribeMsgRecordListResponse() (response *DescribeMsgRecordListResponse) {
    response = &DescribeMsgRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMsgRecordList
// 查询应用的对话消息记录列表，支持按渠道类型、反馈类型、意图、调用结果等条件过滤，并支持游标分页与按创建时间排序
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMsgRecordList(request *DescribeMsgRecordListRequest) (response *DescribeMsgRecordListResponse, err error) {
    return c.DescribeMsgRecordListWithContext(context.Background(), request)
}

// DescribeMsgRecordList
// 查询应用的对话消息记录列表，支持按渠道类型、反馈类型、意图、调用结果等条件过滤，并支持游标分页与按创建时间排序
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMsgRecordListWithContext(ctx context.Context, request *DescribeMsgRecordListRequest) (response *DescribeMsgRecordListResponse, err error) {
    if request == nil {
        request = NewDescribeMsgRecordListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeMsgRecordList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMsgRecordList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMsgRecordListResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePlugin(request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
    return c.DescribePluginWithContext(context.Background(), request)
}

// DescribePlugin
// 获取插件详情
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePluginSummaryList(request *DescribePluginSummaryListRequest) (response *DescribePluginSummaryListResponse, err error) {
    return c.DescribePluginSummaryListWithContext(context.Background(), request)
}

// DescribePluginSummaryList
// 获取插件列表
//
// 可能返回的错误码:
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

func NewDescribeQARequest() (request *DescribeQARequest) {
    request = &DescribeQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeQA")
    
    
    return
}

func NewDescribeQAResponse() (response *DescribeQAResponse) {
    response = &DescribeQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQA
// 查询 QA 详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeQA(request *DescribeQARequest) (response *DescribeQAResponse, err error) {
    return c.DescribeQAWithContext(context.Background(), request)
}

// DescribeQA
// 查询 QA 详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeQAWithContext(ctx context.Context, request *DescribeQARequest) (response *DescribeQAResponse, err error) {
    if request == nil {
        request = NewDescribeQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQASummaryListRequest() (request *DescribeQASummaryListRequest) {
    request = &DescribeQASummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeQASummaryList")
    
    
    return
}

func NewDescribeQASummaryListResponse() (response *DescribeQASummaryListResponse) {
    response = &DescribeQASummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQASummaryList
// 查询 QA 列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeQASummaryList(request *DescribeQASummaryListRequest) (response *DescribeQASummaryListResponse, err error) {
    return c.DescribeQASummaryListWithContext(context.Background(), request)
}

// DescribeQASummaryList
// 查询 QA 列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeQASummaryListWithContext(ctx context.Context, request *DescribeQASummaryListRequest) (response *DescribeQASummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeQASummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeQASummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQASummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQASummaryListResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeReleaseList(request *DescribeReleaseListRequest) (response *DescribeReleaseListResponse, err error) {
    return c.DescribeReleaseListWithContext(context.Background(), request)
}

// DescribeReleaseList
// 发布记录列表
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeReleaseSummary(request *DescribeReleaseSummaryRequest) (response *DescribeReleaseSummaryResponse, err error) {
    return c.DescribeReleaseSummaryWithContext(context.Background(), request)
}

// DescribeReleaseSummary
// 查询发布任务
//
// 可能返回的错误码:
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

func NewDescribeResourceSummaryRequest() (request *DescribeResourceSummaryRequest) {
    request = &DescribeResourceSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeResourceSummary")
    
    
    return
}

func NewDescribeResourceSummaryResponse() (response *DescribeResourceSummaryResponse) {
    response = &DescribeResourceSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceSummary
// 获取用户资源套餐和增值包用量信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeResourceSummary(request *DescribeResourceSummaryRequest) (response *DescribeResourceSummaryResponse, err error) {
    return c.DescribeResourceSummaryWithContext(context.Background(), request)
}

// DescribeResourceSummary
// 获取用户资源套餐和增值包用量信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeResourceSummaryWithContext(ctx context.Context, request *DescribeResourceSummaryRequest) (response *DescribeResourceSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeResourceSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeResourceSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceSummaryResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillCategoryList(request *DescribeSkillCategoryListRequest) (response *DescribeSkillCategoryListResponse, err error) {
    return c.DescribeSkillCategoryListWithContext(context.Background(), request)
}

// DescribeSkillCategoryList
// 查询 Skill 分类列表
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillDetail(request *DescribeSkillDetailRequest) (response *DescribeSkillDetailResponse, err error) {
    return c.DescribeSkillDetailWithContext(context.Background(), request)
}

// DescribeSkillDetail
// 查询skill详情
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillReferenceList(request *DescribeSkillReferenceListRequest) (response *DescribeSkillReferenceListResponse, err error) {
    return c.DescribeSkillReferenceListWithContext(context.Background(), request)
}

// DescribeSkillReferenceList
// 查询某个 Skill 被引用的详情列表（按 SkillRefType 分组：OpenClaw / cloud agent / 企业助手 agent） 鉴权：同 DescribeSkillDetail（能看该 Skill 即可查）
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSkillSummaryList(request *DescribeSkillSummaryListRequest) (response *DescribeSkillSummaryListResponse, err error) {
    return c.DescribeSkillSummaryListWithContext(context.Background(), request)
}

// DescribeSkillSummaryList
// 查询 Skill 列表
//
// 可能返回的错误码:
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

func NewDescribeUsageDetailListRequest() (request *DescribeUsageDetailListRequest) {
    request = &DescribeUsageDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeUsageDetailList")
    
    
    return
}

func NewDescribeUsageDetailListResponse() (response *DescribeUsageDetailListResponse) {
    response = &DescribeUsageDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsageDetailList
// 查询资源调用时序明细，支持模型和插件两类资源，按时间顺序返回每条调用记录的详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUsageDetailList(request *DescribeUsageDetailListRequest) (response *DescribeUsageDetailListResponse, err error) {
    return c.DescribeUsageDetailListWithContext(context.Background(), request)
}

// DescribeUsageDetailList
// 查询资源调用时序明细，支持模型和插件两类资源，按时间顺序返回每条调用记录的详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUsageDetailListWithContext(ctx context.Context, request *DescribeUsageDetailListRequest) (response *DescribeUsageDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeUsageDetailListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeUsageDetailList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsageDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageSummaryListRequest() (request *DescribeUsageSummaryListRequest) {
    request = &DescribeUsageSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeUsageSummaryList")
    
    
    return
}

func NewDescribeUsageSummaryListResponse() (response *DescribeUsageSummaryListResponse) {
    response = &DescribeUsageSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsageSummaryList
// 查询资源用量聚合明细，支持模型、插件、平台三类资源，按空间/应用维度聚合展示调用次数、Token消耗等指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUsageSummaryList(request *DescribeUsageSummaryListRequest) (response *DescribeUsageSummaryListResponse, err error) {
    return c.DescribeUsageSummaryListWithContext(context.Background(), request)
}

// DescribeUsageSummaryList
// 查询资源用量聚合明细，支持模型、插件、平台三类资源，按空间/应用维度聚合展示调用次数、Token消耗等指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUsageSummaryListWithContext(ctx context.Context, request *DescribeUsageSummaryListRequest) (response *DescribeUsageSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeUsageSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeUsageSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsageSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageSummaryListResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVariable(request *DescribeVariableRequest) (response *DescribeVariableResponse, err error) {
    return c.DescribeVariableWithContext(context.Background(), request)
}

// DescribeVariable
// 获取参数变量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVariableList(request *DescribeVariableListRequest) (response *DescribeVariableListResponse, err error) {
    return c.DescribeVariableListWithContext(context.Background(), request)
}

// DescribeVariableList
// 获取参数变量列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewExportQARequest() (request *ExportQARequest) {
    request = &ExportQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ExportQA")
    
    
    return
}

func NewExportQAResponse() (response *ExportQAResponse) {
    response = &ExportQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportQA
// 异步导出 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportQA(request *ExportQARequest) (response *ExportQAResponse, err error) {
    return c.ExportQAWithContext(context.Background(), request)
}

// ExportQA
// 异步导出 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportQAWithContext(ctx context.Context, request *ExportQARequest) (response *ExportQAResponse, err error) {
    if request == nil {
        request = NewExportQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ExportQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportQAResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) FavoritePlugin(request *FavoritePluginRequest) (response *FavoritePluginResponse, err error) {
    return c.FavoritePluginWithContext(context.Background(), request)
}

// FavoritePlugin
// 收藏插件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) FavoriteSkill(request *FavoriteSkillRequest) (response *FavoriteSkillResponse, err error) {
    return c.FavoriteSkillWithContext(context.Background(), request)
}

// FavoriteSkill
// 收藏skill
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewImportDocListRequest() (request *ImportDocListRequest) {
    request = &ImportDocListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ImportDocList")
    
    
    return
}

func NewImportDocListResponse() (response *ImportDocListResponse) {
    response = &ImportDocListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportDocList
// 批量导入文档
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ImportDocList(request *ImportDocListRequest) (response *ImportDocListResponse, err error) {
    return c.ImportDocListWithContext(context.Background(), request)
}

// ImportDocList
// 批量导入文档
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ImportDocListWithContext(ctx context.Context, request *ImportDocListRequest) (response *ImportDocListResponse, err error) {
    if request == nil {
        request = NewImportDocListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ImportDocList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportDocList require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportDocListResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAgent(request *ModifyAgentRequest) (response *ModifyAgentResponse, err error) {
    return c.ModifyAgentWithContext(context.Background(), request)
}

// ModifyAgent
// 修改Agent配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    return c.ModifyAppWithContext(context.Background(), request)
}

// ModifyApp
// 修改应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewModifyAppTriggerRequest() (request *ModifyAppTriggerRequest) {
    request = &ModifyAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyAppTrigger")
    
    
    return
}

func NewModifyAppTriggerResponse() (response *ModifyAppTriggerResponse) {
    response = &ModifyAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAppTrigger
// ModifyAppTrigger
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAppTrigger(request *ModifyAppTriggerRequest) (response *ModifyAppTriggerResponse, err error) {
    return c.ModifyAppTriggerWithContext(context.Background(), request)
}

// ModifyAppTrigger
// ModifyAppTrigger
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAppTriggerWithContext(ctx context.Context, request *ModifyAppTriggerRequest) (response *ModifyAppTriggerResponse, err error) {
    if request == nil {
        request = NewModifyAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCategoryRequest() (request *ModifyCategoryRequest) {
    request = &ModifyCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyCategory")
    
    
    return
}

func NewModifyCategoryResponse() (response *ModifyCategoryResponse) {
    response = &ModifyCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCategory
// 修改分类
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCategory(request *ModifyCategoryRequest) (response *ModifyCategoryResponse, err error) {
    return c.ModifyCategoryWithContext(context.Background(), request)
}

// ModifyCategory
// 修改分类
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCategoryWithContext(ctx context.Context, request *ModifyCategoryRequest) (response *ModifyCategoryResponse, err error) {
    if request == nil {
        request = NewModifyCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewModifyChannelRequest() (request *ModifyChannelRequest) {
    request = &ModifyChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyChannel")
    
    
    return
}

func NewModifyChannelResponse() (response *ModifyChannelResponse) {
    response = &ModifyChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyChannel
// 修改渠道（支持修改备注与企微机器人渠道回调机器人ID）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyChannel(request *ModifyChannelRequest) (response *ModifyChannelResponse, err error) {
    return c.ModifyChannelWithContext(context.Background(), request)
}

// ModifyChannel
// 修改渠道（支持修改备注与企微机器人渠道回调机器人ID）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyChannelWithContext(ctx context.Context, request *ModifyChannelRequest) (response *ModifyChannelResponse, err error) {
    if request == nil {
        request = NewModifyChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConflictQARequest() (request *ModifyConflictQARequest) {
    request = &ModifyConflictQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyConflictQA")
    
    
    return
}

func NewModifyConflictQAResponse() (response *ModifyConflictQAResponse) {
    response = &ModifyConflictQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConflictQA
// 修改冲突问
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyConflictQA(request *ModifyConflictQARequest) (response *ModifyConflictQAResponse, err error) {
    return c.ModifyConflictQAWithContext(context.Background(), request)
}

// ModifyConflictQA
// 修改冲突问
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyConflictQAWithContext(ctx context.Context, request *ModifyConflictQARequest) (response *ModifyConflictQAResponse, err error) {
    if request == nil {
        request = NewModifyConflictQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyConflictQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConflictQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConflictQAResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyConversation(request *ModifyConversationRequest) (response *ModifyConversationResponse, err error) {
    return c.ModifyConversationWithContext(context.Background(), request)
}

// ModifyConversation
// 修改会话信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewModifyDocRequest() (request *ModifyDocRequest) {
    request = &ModifyDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyDoc")
    
    
    return
}

func NewModifyDocResponse() (response *ModifyDocResponse) {
    response = &ModifyDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDoc
// 修改单个文档
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDoc(request *ModifyDocRequest) (response *ModifyDocResponse, err error) {
    return c.ModifyDocWithContext(context.Background(), request)
}

// ModifyDoc
// 修改单个文档
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDocWithContext(ctx context.Context, request *ModifyDocRequest) (response *ModifyDocResponse, err error) {
    if request == nil {
        request = NewModifyDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDocListRequest() (request *ModifyDocListRequest) {
    request = &ModifyDocListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyDocList")
    
    
    return
}

func NewModifyDocListResponse() (response *ModifyDocListResponse) {
    response = &ModifyDocListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDocList
// 批量修改文档
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDocList(request *ModifyDocListRequest) (response *ModifyDocListResponse, err error) {
    return c.ModifyDocListWithContext(context.Background(), request)
}

// ModifyDocList
// 批量修改文档
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDocListWithContext(ctx context.Context, request *ModifyDocListRequest) (response *ModifyDocListResponse, err error) {
    if request == nil {
        request = NewModifyDocListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyDocList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDocList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKBRequest() (request *ModifyKBRequest) {
    request = &ModifyKBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyKB")
    
    
    return
}

func NewModifyKBResponse() (response *ModifyKBResponse) {
    response = &ModifyKBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyKB
// 修改知识库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyKB(request *ModifyKBRequest) (response *ModifyKBResponse, err error) {
    return c.ModifyKBWithContext(context.Background(), request)
}

// ModifyKB
// 修改知识库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyKBWithContext(ctx context.Context, request *ModifyKBRequest) (response *ModifyKBResponse, err error) {
    if request == nil {
        request = NewModifyKBRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyKB")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyKB require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyKBResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLabelRequest() (request *ModifyLabelRequest) {
    request = &ModifyLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyLabel")
    
    
    return
}

func NewModifyLabelResponse() (response *ModifyLabelResponse) {
    response = &ModifyLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLabel
// 修改标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyLabel(request *ModifyLabelRequest) (response *ModifyLabelResponse, err error) {
    return c.ModifyLabelWithContext(context.Background(), request)
}

// ModifyLabel
// 修改标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyLabelWithContext(ctx context.Context, request *ModifyLabelRequest) (response *ModifyLabelResponse, err error) {
    if request == nil {
        request = NewModifyLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLabelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMsgRecordCategoryRequest() (request *ModifyMsgRecordCategoryRequest) {
    request = &ModifyMsgRecordCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyMsgRecordCategory")
    
    
    return
}

func NewModifyMsgRecordCategoryResponse() (response *ModifyMsgRecordCategoryResponse) {
    response = &ModifyMsgRecordCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMsgRecordCategory
// 修改指定消息记录分类的名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMsgRecordCategory(request *ModifyMsgRecordCategoryRequest) (response *ModifyMsgRecordCategoryResponse, err error) {
    return c.ModifyMsgRecordCategoryWithContext(context.Background(), request)
}

// ModifyMsgRecordCategory
// 修改指定消息记录分类的名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMsgRecordCategoryWithContext(ctx context.Context, request *ModifyMsgRecordCategoryRequest) (response *ModifyMsgRecordCategoryResponse, err error) {
    if request == nil {
        request = NewModifyMsgRecordCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyMsgRecordCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMsgRecordCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMsgRecordCategoryResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPlugin(request *ModifyPluginRequest) (response *ModifyPluginResponse, err error) {
    return c.ModifyPluginWithContext(context.Background(), request)
}

// ModifyPlugin
// 修改插件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewModifyQARequest() (request *ModifyQARequest) {
    request = &ModifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyQA")
    
    
    return
}

func NewModifyQAResponse() (response *ModifyQAResponse) {
    response = &ModifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQA
// 修改单个 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyQA(request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    return c.ModifyQAWithContext(context.Background(), request)
}

// ModifyQA
// 修改单个 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyQAWithContext(ctx context.Context, request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    if request == nil {
        request = NewModifyQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQAListRequest() (request *ModifyQAListRequest) {
    request = &ModifyQAListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyQAList")
    
    
    return
}

func NewModifyQAListResponse() (response *ModifyQAListResponse) {
    response = &ModifyQAListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQAList
// 批量修改 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyQAList(request *ModifyQAListRequest) (response *ModifyQAListResponse, err error) {
    return c.ModifyQAListWithContext(context.Background(), request)
}

// ModifyQAList
// 批量修改 QA
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyQAListWithContext(ctx context.Context, request *ModifyQAListRequest) (response *ModifyQAListResponse, err error) {
    if request == nil {
        request = NewModifyQAListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyQAList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQAList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAListResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySkill(request *ModifySkillRequest) (response *ModifySkillResponse, err error) {
    return c.ModifySkillWithContext(context.Background(), request)
}

// ModifySkill
// Skill修改
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewPauseAppTriggerRequest() (request *PauseAppTriggerRequest) {
    request = &PauseAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "PauseAppTrigger")
    
    
    return
}

func NewPauseAppTriggerResponse() (response *PauseAppTriggerResponse) {
    response = &PauseAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseAppTrigger
// PauseAppTrigger
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) PauseAppTrigger(request *PauseAppTriggerRequest) (response *PauseAppTriggerResponse, err error) {
    return c.PauseAppTriggerWithContext(context.Background(), request)
}

// PauseAppTrigger
// PauseAppTrigger
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) PauseAppTriggerWithContext(ctx context.Context, request *PauseAppTriggerRequest) (response *PauseAppTriggerResponse, err error) {
    if request == nil {
        request = NewPauseAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "PauseAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseAppTriggerResponse()
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
// 注意：当前Claw模式应用会话不支持重置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResetConversation(request *ResetConversationRequest) (response *ResetConversationResponse, err error) {
    return c.ResetConversationWithContext(context.Background(), request)
}

// ResetConversation
// 重置会话
//
// 注意：当前Claw模式应用会话不支持重置
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

func NewResumeAppTriggerRequest() (request *ResumeAppTriggerRequest) {
    request = &ResumeAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ResumeAppTrigger")
    
    
    return
}

func NewResumeAppTriggerResponse() (response *ResumeAppTriggerResponse) {
    response = &ResumeAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeAppTrigger
// ResumeAppTrigger
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResumeAppTrigger(request *ResumeAppTriggerRequest) (response *ResumeAppTriggerResponse, err error) {
    return c.ResumeAppTriggerWithContext(context.Background(), request)
}

// ResumeAppTrigger
// ResumeAppTrigger
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResumeAppTriggerWithContext(ctx context.Context, request *ResumeAppTriggerRequest) (response *ResumeAppTriggerResponse, err error) {
    if request == nil {
        request = NewResumeAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ResumeAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeAppTriggerResponse()
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

func NewRunAppTriggerNowRequest() (request *RunAppTriggerNowRequest) {
    request = &RunAppTriggerNowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "RunAppTriggerNow")
    
    
    return
}

func NewRunAppTriggerNowResponse() (response *RunAppTriggerNowResponse) {
    response = &RunAppTriggerNowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunAppTriggerNow
// RunAppTriggerNow
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunAppTriggerNow(request *RunAppTriggerNowRequest) (response *RunAppTriggerNowResponse, err error) {
    return c.RunAppTriggerNowWithContext(context.Background(), request)
}

// RunAppTriggerNow
// RunAppTriggerNow
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunAppTriggerNowWithContext(ctx context.Context, request *RunAppTriggerNowRequest) (response *RunAppTriggerNowResponse, err error) {
    if request == nil {
        request = NewRunAppTriggerNowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "RunAppTriggerNow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunAppTriggerNow require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunAppTriggerNowResponse()
    err = c.Send(request, response)
    return
}

func NewSearchKnowledgeRequest() (request *SearchKnowledgeRequest) {
    request = &SearchKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "SearchKnowledge")
    
    
    return
}

func NewSearchKnowledgeResponse() (response *SearchKnowledgeResponse) {
    response = &SearchKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchKnowledge
// 知识检索
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SearchKnowledge(request *SearchKnowledgeRequest) (response *SearchKnowledgeResponse, err error) {
    return c.SearchKnowledgeWithContext(context.Background(), request)
}

// SearchKnowledge
// 知识检索
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SearchKnowledgeWithContext(ctx context.Context, request *SearchKnowledgeRequest) (response *SearchKnowledgeResponse, err error) {
    if request == nil {
        request = NewSearchKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "SearchKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchKnowledgeResponse()
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
