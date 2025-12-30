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

package v20231130

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-11-30"

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


func NewCallbackWorkflowToolNodeRequest() (request *CallbackWorkflowToolNodeRequest) {
    request = &CallbackWorkflowToolNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CallbackWorkflowToolNode")
    
    
    return
}

func NewCallbackWorkflowToolNodeResponse() (response *CallbackWorkflowToolNodeResponse) {
    response = &CallbackWorkflowToolNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CallbackWorkflowToolNode
// 工作流工具节点异步回调
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CallbackWorkflowToolNode(request *CallbackWorkflowToolNodeRequest) (response *CallbackWorkflowToolNodeResponse, err error) {
    return c.CallbackWorkflowToolNodeWithContext(context.Background(), request)
}

// CallbackWorkflowToolNode
// 工作流工具节点异步回调
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CallbackWorkflowToolNodeWithContext(ctx context.Context, request *CallbackWorkflowToolNodeRequest) (response *CallbackWorkflowToolNodeResponse, err error) {
    if request == nil {
        request = NewCallbackWorkflowToolNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CallbackWorkflowToolNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CallbackWorkflowToolNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCallbackWorkflowToolNodeResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAttributeLabelExistRequest() (request *CheckAttributeLabelExistRequest) {
    request = &CheckAttributeLabelExistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CheckAttributeLabelExist")
    
    
    return
}

func NewCheckAttributeLabelExistResponse() (response *CheckAttributeLabelExistResponse) {
    response = &CheckAttributeLabelExistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckAttributeLabelExist
// 检查属性下的标签名是否存在
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckAttributeLabelExist(request *CheckAttributeLabelExistRequest) (response *CheckAttributeLabelExistResponse, err error) {
    return c.CheckAttributeLabelExistWithContext(context.Background(), request)
}

// CheckAttributeLabelExist
// 检查属性下的标签名是否存在
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckAttributeLabelExistWithContext(ctx context.Context, request *CheckAttributeLabelExistRequest) (response *CheckAttributeLabelExistResponse, err error) {
    if request == nil {
        request = NewCheckAttributeLabelExistRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CheckAttributeLabelExist")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAttributeLabelExist require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAttributeLabelExistResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAttributeLabelReferRequest() (request *CheckAttributeLabelReferRequest) {
    request = &CheckAttributeLabelReferRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CheckAttributeLabelRefer")
    
    
    return
}

func NewCheckAttributeLabelReferResponse() (response *CheckAttributeLabelReferResponse) {
    response = &CheckAttributeLabelReferResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckAttributeLabelRefer
// 检查属性标签引用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckAttributeLabelRefer(request *CheckAttributeLabelReferRequest) (response *CheckAttributeLabelReferResponse, err error) {
    return c.CheckAttributeLabelReferWithContext(context.Background(), request)
}

// CheckAttributeLabelRefer
// 检查属性标签引用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckAttributeLabelReferWithContext(ctx context.Context, request *CheckAttributeLabelReferRequest) (response *CheckAttributeLabelReferResponse, err error) {
    if request == nil {
        request = NewCheckAttributeLabelReferRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CheckAttributeLabelRefer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAttributeLabelRefer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAttributeLabelReferResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateApp")
    
    
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApp
// 创建知识引擎应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    return c.CreateAppWithContext(context.Background(), request)
}

// CreateApp
// 创建知识引擎应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAttributeLabelRequest() (request *CreateAttributeLabelRequest) {
    request = &CreateAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateAttributeLabel")
    
    
    return
}

func NewCreateAttributeLabelResponse() (response *CreateAttributeLabelResponse) {
    response = &CreateAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAttributeLabel
// 创建标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAttributeLabel(request *CreateAttributeLabelRequest) (response *CreateAttributeLabelResponse, err error) {
    return c.CreateAttributeLabelWithContext(context.Background(), request)
}

// CreateAttributeLabel
// 创建标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAttributeLabelWithContext(ctx context.Context, request *CreateAttributeLabelRequest) (response *CreateAttributeLabelResponse, err error) {
    if request == nil {
        request = NewCreateAttributeLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateAttributeLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDocCateRequest() (request *CreateDocCateRequest) {
    request = &CreateDocCateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateDocCate")
    
    
    return
}

func NewCreateDocCateResponse() (response *CreateDocCateResponse) {
    response = &CreateDocCateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDocCate
// 创建Doc分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDocCate(request *CreateDocCateRequest) (response *CreateDocCateResponse, err error) {
    return c.CreateDocCateWithContext(context.Background(), request)
}

// CreateDocCate
// 创建Doc分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDocCateWithContext(ctx context.Context, request *CreateDocCateRequest) (response *CreateDocCateResponse, err error) {
    if request == nil {
        request = NewCreateDocCateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateDocCate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDocCate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDocCateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQARequest() (request *CreateQARequest) {
    request = &CreateQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateQA")
    
    
    return
}

func NewCreateQAResponse() (response *CreateQAResponse) {
    response = &CreateQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQA
// 录入问答
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateQA(request *CreateQARequest) (response *CreateQAResponse, err error) {
    return c.CreateQAWithContext(context.Background(), request)
}

// CreateQA
// 录入问答
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateQAWithContext(ctx context.Context, request *CreateQARequest) (response *CreateQAResponse, err error) {
    if request == nil {
        request = NewCreateQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQAResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQACateRequest() (request *CreateQACateRequest) {
    request = &CreateQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateQACate")
    
    
    return
}

func NewCreateQACateResponse() (response *CreateQACateResponse) {
    response = &CreateQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQACate
// 创建QA分类
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateQACate(request *CreateQACateRequest) (response *CreateQACateResponse, err error) {
    return c.CreateQACateWithContext(context.Background(), request)
}

// CreateQACate
// 创建QA分类
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateQACateWithContext(ctx context.Context, request *CreateQACateRequest) (response *CreateQACateResponse, err error) {
    if request == nil {
        request = NewCreateQACateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateQACate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQACateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRejectedQuestionRequest() (request *CreateRejectedQuestionRequest) {
    request = &CreateRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateRejectedQuestion")
    
    
    return
}

func NewCreateRejectedQuestionResponse() (response *CreateRejectedQuestionResponse) {
    response = &CreateRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRejectedQuestion
// 创建拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateRejectedQuestion(request *CreateRejectedQuestionRequest) (response *CreateRejectedQuestionResponse, err error) {
    return c.CreateRejectedQuestionWithContext(context.Background(), request)
}

// CreateRejectedQuestion
// 创建拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateRejectedQuestionWithContext(ctx context.Context, request *CreateRejectedQuestionRequest) (response *CreateRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewCreateRejectedQuestionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateRejectedQuestion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReleaseRequest() (request *CreateReleaseRequest) {
    request = &CreateReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateRelease")
    
    
    return
}

func NewCreateReleaseResponse() (response *CreateReleaseResponse) {
    response = &CreateReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRelease
// 创建发布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateRelease(request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    return c.CreateReleaseWithContext(context.Background(), request)
}

// CreateRelease
// 创建发布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateReleaseWithContext(ctx context.Context, request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    if request == nil {
        request = NewCreateReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSharedKnowledgeRequest() (request *CreateSharedKnowledgeRequest) {
    request = &CreateSharedKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateSharedKnowledge")
    
    
    return
}

func NewCreateSharedKnowledgeResponse() (response *CreateSharedKnowledgeResponse) {
    response = &CreateSharedKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSharedKnowledge
// 创建共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSharedKnowledge(request *CreateSharedKnowledgeRequest) (response *CreateSharedKnowledgeResponse, err error) {
    return c.CreateSharedKnowledgeWithContext(context.Background(), request)
}

// CreateSharedKnowledge
// 创建共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSharedKnowledgeWithContext(ctx context.Context, request *CreateSharedKnowledgeRequest) (response *CreateSharedKnowledgeResponse, err error) {
    if request == nil {
        request = NewCreateSharedKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateSharedKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSharedKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSharedKnowledgeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVarRequest() (request *CreateVarRequest) {
    request = &CreateVarRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateVar")
    
    
    return
}

func NewCreateVarResponse() (response *CreateVarResponse) {
    response = &CreateVarResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVar
// 创建变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateVar(request *CreateVarRequest) (response *CreateVarResponse, err error) {
    return c.CreateVarWithContext(context.Background(), request)
}

// CreateVar
// 创建变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateVarWithContext(ctx context.Context, request *CreateVarRequest) (response *CreateVarResponse, err error) {
    if request == nil {
        request = NewCreateVarRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateVar")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVar require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVarResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRunRequest() (request *CreateWorkflowRunRequest) {
    request = &CreateWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateWorkflowRun")
    
    
    return
}

func NewCreateWorkflowRunResponse() (response *CreateWorkflowRunResponse) {
    response = &CreateWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflowRun
// 本接口用来创建工作流的异步运行实例，创建成功后工作流会在后台异步运行，接口返回工作流运行实例ID（WorkflowRunId）等信息。后面可通过调用DescribeWorkflowRun接口查工作流运行的详情。
//
// 注意：工作流的异步运行是基于应用的，需要先把对应的应用配置成“单工作流模式”，并且打开“异步调用”的开关，才能创建成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflowRun(request *CreateWorkflowRunRequest) (response *CreateWorkflowRunResponse, err error) {
    return c.CreateWorkflowRunWithContext(context.Background(), request)
}

// CreateWorkflowRun
// 本接口用来创建工作流的异步运行实例，创建成功后工作流会在后台异步运行，接口返回工作流运行实例ID（WorkflowRunId）等信息。后面可通过调用DescribeWorkflowRun接口查工作流运行的详情。
//
// 注意：工作流的异步运行是基于应用的，需要先把对应的应用配置成“单工作流模式”，并且打开“异步调用”的开关，才能创建成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflowRunWithContext(ctx context.Context, request *CreateWorkflowRunRequest) (response *CreateWorkflowRunResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "CreateWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentRequest() (request *DeleteAgentRequest) {
    request = &DeleteAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteAgent")
    
    
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
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteAgent")
    
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
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteApp")
    
    
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
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttributeLabelRequest() (request *DeleteAttributeLabelRequest) {
    request = &DeleteAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteAttributeLabel")
    
    
    return
}

func NewDeleteAttributeLabelResponse() (response *DeleteAttributeLabelResponse) {
    response = &DeleteAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAttributeLabel
// 删除属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAttributeLabel(request *DeleteAttributeLabelRequest) (response *DeleteAttributeLabelResponse, err error) {
    return c.DeleteAttributeLabelWithContext(context.Background(), request)
}

// DeleteAttributeLabel
// 删除属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAttributeLabelWithContext(ctx context.Context, request *DeleteAttributeLabelRequest) (response *DeleteAttributeLabelResponse, err error) {
    if request == nil {
        request = NewDeleteAttributeLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteAttributeLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDocRequest() (request *DeleteDocRequest) {
    request = &DeleteDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteDoc")
    
    
    return
}

func NewDeleteDocResponse() (response *DeleteDocResponse) {
    response = &DeleteDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDoc
// 删除文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDoc(request *DeleteDocRequest) (response *DeleteDocResponse, err error) {
    return c.DeleteDocWithContext(context.Background(), request)
}

// DeleteDoc
// 删除文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDocWithContext(ctx context.Context, request *DeleteDocRequest) (response *DeleteDocResponse, err error) {
    if request == nil {
        request = NewDeleteDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDocCateRequest() (request *DeleteDocCateRequest) {
    request = &DeleteDocCateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteDocCate")
    
    
    return
}

func NewDeleteDocCateResponse() (response *DeleteDocCateResponse) {
    response = &DeleteDocCateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDocCate
// Doc分类删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDocCate(request *DeleteDocCateRequest) (response *DeleteDocCateResponse, err error) {
    return c.DeleteDocCateWithContext(context.Background(), request)
}

// DeleteDocCate
// Doc分类删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDocCateWithContext(ctx context.Context, request *DeleteDocCateRequest) (response *DeleteDocCateResponse, err error) {
    if request == nil {
        request = NewDeleteDocCateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteDocCate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDocCate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocCateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQARequest() (request *DeleteQARequest) {
    request = &DeleteQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteQA")
    
    
    return
}

func NewDeleteQAResponse() (response *DeleteQAResponse) {
    response = &DeleteQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQA
// 删除问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQA(request *DeleteQARequest) (response *DeleteQAResponse, err error) {
    return c.DeleteQAWithContext(context.Background(), request)
}

// DeleteQA
// 删除问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQAWithContext(ctx context.Context, request *DeleteQARequest) (response *DeleteQAResponse, err error) {
    if request == nil {
        request = NewDeleteQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQAResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQACateRequest() (request *DeleteQACateRequest) {
    request = &DeleteQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteQACate")
    
    
    return
}

func NewDeleteQACateResponse() (response *DeleteQACateResponse) {
    response = &DeleteQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQACate
// 分类删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQACate(request *DeleteQACateRequest) (response *DeleteQACateResponse, err error) {
    return c.DeleteQACateWithContext(context.Background(), request)
}

// DeleteQACate
// 分类删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQACateWithContext(ctx context.Context, request *DeleteQACateRequest) (response *DeleteQACateResponse, err error) {
    if request == nil {
        request = NewDeleteQACateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteQACate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQACateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRejectedQuestionRequest() (request *DeleteRejectedQuestionRequest) {
    request = &DeleteRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteRejectedQuestion")
    
    
    return
}

func NewDeleteRejectedQuestionResponse() (response *DeleteRejectedQuestionResponse) {
    response = &DeleteRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRejectedQuestion
// 删除拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRejectedQuestion(request *DeleteRejectedQuestionRequest) (response *DeleteRejectedQuestionResponse, err error) {
    return c.DeleteRejectedQuestionWithContext(context.Background(), request)
}

// DeleteRejectedQuestion
// 删除拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRejectedQuestionWithContext(ctx context.Context, request *DeleteRejectedQuestionRequest) (response *DeleteRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewDeleteRejectedQuestionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteRejectedQuestion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSharedKnowledgeRequest() (request *DeleteSharedKnowledgeRequest) {
    request = &DeleteSharedKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteSharedKnowledge")
    
    
    return
}

func NewDeleteSharedKnowledgeResponse() (response *DeleteSharedKnowledgeResponse) {
    response = &DeleteSharedKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSharedKnowledge
// 删除共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSharedKnowledge(request *DeleteSharedKnowledgeRequest) (response *DeleteSharedKnowledgeResponse, err error) {
    return c.DeleteSharedKnowledgeWithContext(context.Background(), request)
}

// DeleteSharedKnowledge
// 删除共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSharedKnowledgeWithContext(ctx context.Context, request *DeleteSharedKnowledgeRequest) (response *DeleteSharedKnowledgeResponse, err error) {
    if request == nil {
        request = NewDeleteSharedKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteSharedKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSharedKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSharedKnowledgeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVarRequest() (request *DeleteVarRequest) {
    request = &DeleteVarRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteVar")
    
    
    return
}

func NewDeleteVarResponse() (response *DeleteVarResponse) {
    response = &DeleteVarResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVar
// 删除变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVar(request *DeleteVarRequest) (response *DeleteVarResponse, err error) {
    return c.DeleteVarWithContext(context.Background(), request)
}

// DeleteVar
// 删除变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVarWithContext(ctx context.Context, request *DeleteVarRequest) (response *DeleteVarResponse, err error) {
    if request == nil {
        request = NewDeleteVarRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DeleteVar")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVar require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVarResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppRequest() (request *DescribeAppRequest) {
    request = &DescribeAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeApp")
    
    
    return
}

func NewDescribeAppResponse() (response *DescribeAppResponse) {
    response = &DescribeAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApp
// 获取企业下应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeApp(request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    return c.DescribeAppWithContext(context.Background(), request)
}

// DescribeApp
// 获取企业下应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAppWithContext(ctx context.Context, request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    if request == nil {
        request = NewDescribeAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppAgentListRequest() (request *DescribeAppAgentListRequest) {
    request = &DescribeAppAgentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeAppAgentList")
    
    
    return
}

func NewDescribeAppAgentListResponse() (response *DescribeAppAgentListResponse) {
    response = &DescribeAppAgentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppAgentList
// 查询指定应用下的Agent列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAppAgentList(request *DescribeAppAgentListRequest) (response *DescribeAppAgentListResponse, err error) {
    return c.DescribeAppAgentListWithContext(context.Background(), request)
}

// DescribeAppAgentList
// 查询指定应用下的Agent列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAppAgentListWithContext(ctx context.Context, request *DescribeAppAgentListRequest) (response *DescribeAppAgentListResponse, err error) {
    if request == nil {
        request = NewDescribeAppAgentListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeAppAgentList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppAgentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppAgentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttributeLabelRequest() (request *DescribeAttributeLabelRequest) {
    request = &DescribeAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeAttributeLabel")
    
    
    return
}

func NewDescribeAttributeLabelResponse() (response *DescribeAttributeLabelResponse) {
    response = &DescribeAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAttributeLabel
// 查询属性标签详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAttributeLabel(request *DescribeAttributeLabelRequest) (response *DescribeAttributeLabelResponse, err error) {
    return c.DescribeAttributeLabelWithContext(context.Background(), request)
}

// DescribeAttributeLabel
// 查询属性标签详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAttributeLabelWithContext(ctx context.Context, request *DescribeAttributeLabelRequest) (response *DescribeAttributeLabelResponse, err error) {
    if request == nil {
        request = NewDescribeAttributeLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeAttributeLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallStatsGraphRequest() (request *DescribeCallStatsGraphRequest) {
    request = &DescribeCallStatsGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeCallStatsGraph")
    
    
    return
}

func NewDescribeCallStatsGraphResponse() (response *DescribeCallStatsGraphResponse) {
    response = &DescribeCallStatsGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCallStatsGraph
// 接口调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCallStatsGraph(request *DescribeCallStatsGraphRequest) (response *DescribeCallStatsGraphResponse, err error) {
    return c.DescribeCallStatsGraphWithContext(context.Background(), request)
}

// DescribeCallStatsGraph
// 接口调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCallStatsGraphWithContext(ctx context.Context, request *DescribeCallStatsGraphRequest) (response *DescribeCallStatsGraphResponse, err error) {
    if request == nil {
        request = NewDescribeCallStatsGraphRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeCallStatsGraph")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallStatsGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallStatsGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrencyUsageRequest() (request *DescribeConcurrencyUsageRequest) {
    request = &DescribeConcurrencyUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeConcurrencyUsage")
    
    
    return
}

func NewDescribeConcurrencyUsageResponse() (response *DescribeConcurrencyUsageResponse) {
    response = &DescribeConcurrencyUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrencyUsage
// 并发调用响应
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsage(request *DescribeConcurrencyUsageRequest) (response *DescribeConcurrencyUsageResponse, err error) {
    return c.DescribeConcurrencyUsageWithContext(context.Background(), request)
}

// DescribeConcurrencyUsage
// 并发调用响应
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsageWithContext(ctx context.Context, request *DescribeConcurrencyUsageRequest) (response *DescribeConcurrencyUsageResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrencyUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeConcurrencyUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrencyUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrencyUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrencyUsageGraphRequest() (request *DescribeConcurrencyUsageGraphRequest) {
    request = &DescribeConcurrencyUsageGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeConcurrencyUsageGraph")
    
    
    return
}

func NewDescribeConcurrencyUsageGraphResponse() (response *DescribeConcurrencyUsageGraphResponse) {
    response = &DescribeConcurrencyUsageGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrencyUsageGraph
// 并发调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsageGraph(request *DescribeConcurrencyUsageGraphRequest) (response *DescribeConcurrencyUsageGraphResponse, err error) {
    return c.DescribeConcurrencyUsageGraphWithContext(context.Background(), request)
}

// DescribeConcurrencyUsageGraph
// 并发调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsageGraphWithContext(ctx context.Context, request *DescribeConcurrencyUsageGraphRequest) (response *DescribeConcurrencyUsageGraphResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrencyUsageGraphRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeConcurrencyUsageGraph")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrencyUsageGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrencyUsageGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocRequest() (request *DescribeDocRequest) {
    request = &DescribeDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeDoc")
    
    
    return
}

func NewDescribeDocResponse() (response *DescribeDocResponse) {
    response = &DescribeDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDoc
// 文档详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDoc(request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    return c.DescribeDocWithContext(context.Background(), request)
}

// DescribeDoc
// 文档详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDocWithContext(ctx context.Context, request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    if request == nil {
        request = NewDescribeDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKnowledgeUsageRequest() (request *DescribeKnowledgeUsageRequest) {
    request = &DescribeKnowledgeUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeKnowledgeUsage")
    
    
    return
}

func NewDescribeKnowledgeUsageResponse() (response *DescribeKnowledgeUsageResponse) {
    response = &DescribeKnowledgeUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeUsage
// 查询知识库用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsage(request *DescribeKnowledgeUsageRequest) (response *DescribeKnowledgeUsageResponse, err error) {
    return c.DescribeKnowledgeUsageWithContext(context.Background(), request)
}

// DescribeKnowledgeUsage
// 查询知识库用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsageWithContext(ctx context.Context, request *DescribeKnowledgeUsageRequest) (response *DescribeKnowledgeUsageResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeKnowledgeUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKnowledgeUsagePieGraphRequest() (request *DescribeKnowledgeUsagePieGraphRequest) {
    request = &DescribeKnowledgeUsagePieGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeKnowledgeUsagePieGraph")
    
    
    return
}

func NewDescribeKnowledgeUsagePieGraphResponse() (response *DescribeKnowledgeUsagePieGraphResponse) {
    response = &DescribeKnowledgeUsagePieGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeUsagePieGraph
// 查询企业知识库容量饼图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsagePieGraph(request *DescribeKnowledgeUsagePieGraphRequest) (response *DescribeKnowledgeUsagePieGraphResponse, err error) {
    return c.DescribeKnowledgeUsagePieGraphWithContext(context.Background(), request)
}

// DescribeKnowledgeUsagePieGraph
// 查询企业知识库容量饼图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsagePieGraphWithContext(ctx context.Context, request *DescribeKnowledgeUsagePieGraphRequest) (response *DescribeKnowledgeUsagePieGraphResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeUsagePieGraphRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeKnowledgeUsagePieGraph")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeUsagePieGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeUsagePieGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeRunRequest() (request *DescribeNodeRunRequest) {
    request = &DescribeNodeRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeNodeRun")
    
    
    return
}

func NewDescribeNodeRunResponse() (response *DescribeNodeRunResponse) {
    response = &DescribeNodeRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNodeRun
// 通过DescribeWorkflowRun接口获取了工作流异步运行的整体内容，其中包含了基本的节点信息，再通过本接口可查看节点的运行详情（包括输入、输出、日志等）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeNodeRun(request *DescribeNodeRunRequest) (response *DescribeNodeRunResponse, err error) {
    return c.DescribeNodeRunWithContext(context.Background(), request)
}

// DescribeNodeRun
// 通过DescribeWorkflowRun接口获取了工作流异步运行的整体内容，其中包含了基本的节点信息，再通过本接口可查看节点的运行详情（包括输入、输出、日志等）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeNodeRunWithContext(ctx context.Context, request *DescribeNodeRunRequest) (response *DescribeNodeRunResponse, err error) {
    if request == nil {
        request = NewDescribeNodeRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeNodeRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodeRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeRunResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQARequest() (request *DescribeQARequest) {
    request = &DescribeQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeQA")
    
    
    return
}

func NewDescribeQAResponse() (response *DescribeQAResponse) {
    response = &DescribeQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQA
// 问答详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQA(request *DescribeQARequest) (response *DescribeQAResponse, err error) {
    return c.DescribeQAWithContext(context.Background(), request)
}

// DescribeQA
// 问答详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQAWithContext(ctx context.Context, request *DescribeQARequest) (response *DescribeQAResponse, err error) {
    if request == nil {
        request = NewDescribeQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReferRequest() (request *DescribeReferRequest) {
    request = &DescribeReferRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeRefer")
    
    
    return
}

func NewDescribeReferResponse() (response *DescribeReferResponse) {
    response = &DescribeReferResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRefer
// 获取来源详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRefer(request *DescribeReferRequest) (response *DescribeReferResponse, err error) {
    return c.DescribeReferWithContext(context.Background(), request)
}

// DescribeRefer
// 获取来源详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeReferWithContext(ctx context.Context, request *DescribeReferRequest) (response *DescribeReferResponse, err error) {
    if request == nil {
        request = NewDescribeReferRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeRefer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRefer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReferResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseRequest() (request *DescribeReleaseRequest) {
    request = &DescribeReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeRelease")
    
    
    return
}

func NewDescribeReleaseResponse() (response *DescribeReleaseResponse) {
    response = &DescribeReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRelease
// 发布详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRelease(request *DescribeReleaseRequest) (response *DescribeReleaseResponse, err error) {
    return c.DescribeReleaseWithContext(context.Background(), request)
}

// DescribeRelease
// 发布详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeReleaseWithContext(ctx context.Context, request *DescribeReleaseRequest) (response *DescribeReleaseResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseInfoRequest() (request *DescribeReleaseInfoRequest) {
    request = &DescribeReleaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeReleaseInfo")
    
    
    return
}

func NewDescribeReleaseInfoResponse() (response *DescribeReleaseInfoResponse) {
    response = &DescribeReleaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReleaseInfo
// 拉取发布按钮状态、最后发布时间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeReleaseInfo(request *DescribeReleaseInfoRequest) (response *DescribeReleaseInfoResponse, err error) {
    return c.DescribeReleaseInfoWithContext(context.Background(), request)
}

// DescribeReleaseInfo
// 拉取发布按钮状态、最后发布时间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeReleaseInfoWithContext(ctx context.Context, request *DescribeReleaseInfoRequest) (response *DescribeReleaseInfoResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeReleaseInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleaseInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRobotBizIDByAppKeyRequest() (request *DescribeRobotBizIDByAppKeyRequest) {
    request = &DescribeRobotBizIDByAppKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeRobotBizIDByAppKey")
    
    
    return
}

func NewDescribeRobotBizIDByAppKeyResponse() (response *DescribeRobotBizIDByAppKeyResponse) {
    response = &DescribeRobotBizIDByAppKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRobotBizIDByAppKey
// 通过appKey获取应用业务ID
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRobotBizIDByAppKey(request *DescribeRobotBizIDByAppKeyRequest) (response *DescribeRobotBizIDByAppKeyResponse, err error) {
    return c.DescribeRobotBizIDByAppKeyWithContext(context.Background(), request)
}

// DescribeRobotBizIDByAppKey
// 通过appKey获取应用业务ID
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRobotBizIDByAppKeyWithContext(ctx context.Context, request *DescribeRobotBizIDByAppKeyRequest) (response *DescribeRobotBizIDByAppKeyResponse, err error) {
    if request == nil {
        request = NewDescribeRobotBizIDByAppKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeRobotBizIDByAppKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRobotBizIDByAppKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRobotBizIDByAppKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchStatsGraphRequest() (request *DescribeSearchStatsGraphRequest) {
    request = &DescribeSearchStatsGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeSearchStatsGraph")
    
    
    return
}

func NewDescribeSearchStatsGraphResponse() (response *DescribeSearchStatsGraphResponse) {
    response = &DescribeSearchStatsGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSearchStatsGraph
// 查询搜索服务调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSearchStatsGraph(request *DescribeSearchStatsGraphRequest) (response *DescribeSearchStatsGraphResponse, err error) {
    return c.DescribeSearchStatsGraphWithContext(context.Background(), request)
}

// DescribeSearchStatsGraph
// 查询搜索服务调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSearchStatsGraphWithContext(ctx context.Context, request *DescribeSearchStatsGraphRequest) (response *DescribeSearchStatsGraphResponse, err error) {
    if request == nil {
        request = NewDescribeSearchStatsGraphRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeSearchStatsGraph")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchStatsGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchStatsGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSegmentsRequest() (request *DescribeSegmentsRequest) {
    request = &DescribeSegmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeSegments")
    
    
    return
}

func NewDescribeSegmentsResponse() (response *DescribeSegmentsResponse) {
    response = &DescribeSegmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSegments
// 获取片段详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSegments(request *DescribeSegmentsRequest) (response *DescribeSegmentsResponse, err error) {
    return c.DescribeSegmentsWithContext(context.Background(), request)
}

// DescribeSegments
// 获取片段详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSegmentsWithContext(ctx context.Context, request *DescribeSegmentsRequest) (response *DescribeSegmentsResponse, err error) {
    if request == nil {
        request = NewDescribeSegmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeSegments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSegments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSegmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSharedKnowledgeRequest() (request *DescribeSharedKnowledgeRequest) {
    request = &DescribeSharedKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeSharedKnowledge")
    
    
    return
}

func NewDescribeSharedKnowledgeResponse() (response *DescribeSharedKnowledgeResponse) {
    response = &DescribeSharedKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSharedKnowledge
// 查询共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSharedKnowledge(request *DescribeSharedKnowledgeRequest) (response *DescribeSharedKnowledgeResponse, err error) {
    return c.DescribeSharedKnowledgeWithContext(context.Background(), request)
}

// DescribeSharedKnowledge
// 查询共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSharedKnowledgeWithContext(ctx context.Context, request *DescribeSharedKnowledgeRequest) (response *DescribeSharedKnowledgeResponse, err error) {
    if request == nil {
        request = NewDescribeSharedKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeSharedKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSharedKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSharedKnowledgeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageCredentialRequest() (request *DescribeStorageCredentialRequest) {
    request = &DescribeStorageCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeStorageCredential")
    
    
    return
}

func NewDescribeStorageCredentialResponse() (response *DescribeStorageCredentialResponse) {
    response = &DescribeStorageCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStorageCredential
// 获取文件上传临时密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStorageCredential(request *DescribeStorageCredentialRequest) (response *DescribeStorageCredentialResponse, err error) {
    return c.DescribeStorageCredentialWithContext(context.Background(), request)
}

// DescribeStorageCredential
// 获取文件上传临时密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStorageCredentialWithContext(ctx context.Context, request *DescribeStorageCredentialRequest) (response *DescribeStorageCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeStorageCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeStorageCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStorageCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStorageCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenUsageRequest() (request *DescribeTokenUsageRequest) {
    request = &DescribeTokenUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeTokenUsage")
    
    
    return
}

func NewDescribeTokenUsageResponse() (response *DescribeTokenUsageResponse) {
    response = &DescribeTokenUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenUsage
// 接口调用token详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsage(request *DescribeTokenUsageRequest) (response *DescribeTokenUsageResponse, err error) {
    return c.DescribeTokenUsageWithContext(context.Background(), request)
}

// DescribeTokenUsage
// 接口调用token详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsageWithContext(ctx context.Context, request *DescribeTokenUsageRequest) (response *DescribeTokenUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTokenUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeTokenUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenUsageGraphRequest() (request *DescribeTokenUsageGraphRequest) {
    request = &DescribeTokenUsageGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeTokenUsageGraph")
    
    
    return
}

func NewDescribeTokenUsageGraphResponse() (response *DescribeTokenUsageGraphResponse) {
    response = &DescribeTokenUsageGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenUsageGraph
// 接口调用token折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsageGraph(request *DescribeTokenUsageGraphRequest) (response *DescribeTokenUsageGraphResponse, err error) {
    return c.DescribeTokenUsageGraphWithContext(context.Background(), request)
}

// DescribeTokenUsageGraph
// 接口调用token折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsageGraphWithContext(ctx context.Context, request *DescribeTokenUsageGraphRequest) (response *DescribeTokenUsageGraphResponse, err error) {
    if request == nil {
        request = NewDescribeTokenUsageGraphRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeTokenUsageGraph")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenUsageGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenUsageGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnsatisfiedReplyContextRequest() (request *DescribeUnsatisfiedReplyContextRequest) {
    request = &DescribeUnsatisfiedReplyContextRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeUnsatisfiedReplyContext")
    
    
    return
}

func NewDescribeUnsatisfiedReplyContextResponse() (response *DescribeUnsatisfiedReplyContextResponse) {
    response = &DescribeUnsatisfiedReplyContextResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnsatisfiedReplyContext
// 获取不满意回复上下文
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUnsatisfiedReplyContext(request *DescribeUnsatisfiedReplyContextRequest) (response *DescribeUnsatisfiedReplyContextResponse, err error) {
    return c.DescribeUnsatisfiedReplyContextWithContext(context.Background(), request)
}

// DescribeUnsatisfiedReplyContext
// 获取不满意回复上下文
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUnsatisfiedReplyContextWithContext(ctx context.Context, request *DescribeUnsatisfiedReplyContextRequest) (response *DescribeUnsatisfiedReplyContextResponse, err error) {
    if request == nil {
        request = NewDescribeUnsatisfiedReplyContextRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeUnsatisfiedReplyContext")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnsatisfiedReplyContext require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnsatisfiedReplyContextResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowRunRequest() (request *DescribeWorkflowRunRequest) {
    request = &DescribeWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeWorkflowRun")
    
    
    return
}

func NewDescribeWorkflowRunResponse() (response *DescribeWorkflowRunResponse) {
    response = &DescribeWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkflowRun
// 创建了工作流的异步运行实例后，通过本接口可以查询整体的运行详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWorkflowRun(request *DescribeWorkflowRunRequest) (response *DescribeWorkflowRunResponse, err error) {
    return c.DescribeWorkflowRunWithContext(context.Background(), request)
}

// DescribeWorkflowRun
// 创建了工作流的异步运行实例后，通过本接口可以查询整体的运行详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWorkflowRunWithContext(ctx context.Context, request *DescribeWorkflowRunRequest) (response *DescribeWorkflowRunResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "DescribeWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewExportAttributeLabelRequest() (request *ExportAttributeLabelRequest) {
    request = &ExportAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ExportAttributeLabel")
    
    
    return
}

func NewExportAttributeLabelResponse() (response *ExportAttributeLabelResponse) {
    response = &ExportAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportAttributeLabel
// 导出标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAttributeLabel(request *ExportAttributeLabelRequest) (response *ExportAttributeLabelResponse, err error) {
    return c.ExportAttributeLabelWithContext(context.Background(), request)
}

// ExportAttributeLabel
// 导出标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAttributeLabelWithContext(ctx context.Context, request *ExportAttributeLabelRequest) (response *ExportAttributeLabelResponse, err error) {
    if request == nil {
        request = NewExportAttributeLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ExportAttributeLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewExportQAListRequest() (request *ExportQAListRequest) {
    request = &ExportQAListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ExportQAList")
    
    
    return
}

func NewExportQAListResponse() (response *ExportQAListResponse) {
    response = &ExportQAListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportQAList
// 导出QA列表
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportQAList(request *ExportQAListRequest) (response *ExportQAListResponse, err error) {
    return c.ExportQAListWithContext(context.Background(), request)
}

// ExportQAList
// 导出QA列表
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportQAListWithContext(ctx context.Context, request *ExportQAListRequest) (response *ExportQAListResponse, err error) {
    if request == nil {
        request = NewExportQAListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ExportQAList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportQAList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportQAListResponse()
    err = c.Send(request, response)
    return
}

func NewExportUnsatisfiedReplyRequest() (request *ExportUnsatisfiedReplyRequest) {
    request = &ExportUnsatisfiedReplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ExportUnsatisfiedReply")
    
    
    return
}

func NewExportUnsatisfiedReplyResponse() (response *ExportUnsatisfiedReplyResponse) {
    response = &ExportUnsatisfiedReplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportUnsatisfiedReply
// 导出不满意回复
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportUnsatisfiedReply(request *ExportUnsatisfiedReplyRequest) (response *ExportUnsatisfiedReplyResponse, err error) {
    return c.ExportUnsatisfiedReplyWithContext(context.Background(), request)
}

// ExportUnsatisfiedReply
// 导出不满意回复
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportUnsatisfiedReplyWithContext(ctx context.Context, request *ExportUnsatisfiedReplyRequest) (response *ExportUnsatisfiedReplyResponse, err error) {
    if request == nil {
        request = NewExportUnsatisfiedReplyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ExportUnsatisfiedReply")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportUnsatisfiedReply require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportUnsatisfiedReplyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateQARequest() (request *GenerateQARequest) {
    request = &GenerateQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GenerateQA")
    
    
    return
}

func NewGenerateQAResponse() (response *GenerateQAResponse) {
    response = &GenerateQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateQA
// 文档生成问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GenerateQA(request *GenerateQARequest) (response *GenerateQAResponse, err error) {
    return c.GenerateQAWithContext(context.Background(), request)
}

// GenerateQA
// 文档生成问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GenerateQAWithContext(ctx context.Context, request *GenerateQARequest) (response *GenerateQAResponse, err error) {
    if request == nil {
        request = NewGenerateQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GenerateQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateQAResponse()
    err = c.Send(request, response)
    return
}

func NewGetAnswerTypeDataCountRequest() (request *GetAnswerTypeDataCountRequest) {
    request = &GetAnswerTypeDataCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetAnswerTypeDataCount")
    
    
    return
}

func NewGetAnswerTypeDataCountResponse() (response *GetAnswerTypeDataCountResponse) {
    response = &GetAnswerTypeDataCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAnswerTypeDataCount
// 回答类型数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetAnswerTypeDataCount(request *GetAnswerTypeDataCountRequest) (response *GetAnswerTypeDataCountResponse, err error) {
    return c.GetAnswerTypeDataCountWithContext(context.Background(), request)
}

// GetAnswerTypeDataCount
// 回答类型数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetAnswerTypeDataCountWithContext(ctx context.Context, request *GetAnswerTypeDataCountRequest) (response *GetAnswerTypeDataCountResponse, err error) {
    if request == nil {
        request = NewGetAnswerTypeDataCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GetAnswerTypeDataCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAnswerTypeDataCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAnswerTypeDataCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetAppKnowledgeCountRequest() (request *GetAppKnowledgeCountRequest) {
    request = &GetAppKnowledgeCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetAppKnowledgeCount")
    
    
    return
}

func NewGetAppKnowledgeCountResponse() (response *GetAppKnowledgeCountResponse) {
    response = &GetAppKnowledgeCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAppKnowledgeCount
// 获取模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAppKnowledgeCount(request *GetAppKnowledgeCountRequest) (response *GetAppKnowledgeCountResponse, err error) {
    return c.GetAppKnowledgeCountWithContext(context.Background(), request)
}

// GetAppKnowledgeCount
// 获取模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAppKnowledgeCountWithContext(ctx context.Context, request *GetAppKnowledgeCountRequest) (response *GetAppKnowledgeCountResponse, err error) {
    if request == nil {
        request = NewGetAppKnowledgeCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GetAppKnowledgeCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAppKnowledgeCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAppKnowledgeCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetAppSecretRequest() (request *GetAppSecretRequest) {
    request = &GetAppSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetAppSecret")
    
    
    return
}

func NewGetAppSecretResponse() (response *GetAppSecretResponse) {
    response = &GetAppSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAppSecret
// 获取应用密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAppSecret(request *GetAppSecretRequest) (response *GetAppSecretResponse, err error) {
    return c.GetAppSecretWithContext(context.Background(), request)
}

// GetAppSecret
// 获取应用密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAppSecretWithContext(ctx context.Context, request *GetAppSecretRequest) (response *GetAppSecretResponse, err error) {
    if request == nil {
        request = NewGetAppSecretRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GetAppSecret")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAppSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAppSecretResponse()
    err = c.Send(request, response)
    return
}

func NewGetDocPreviewRequest() (request *GetDocPreviewRequest) {
    request = &GetDocPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetDocPreview")
    
    
    return
}

func NewGetDocPreviewResponse() (response *GetDocPreviewResponse) {
    response = &GetDocPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDocPreview
// 获取文档预览信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetDocPreview(request *GetDocPreviewRequest) (response *GetDocPreviewResponse, err error) {
    return c.GetDocPreviewWithContext(context.Background(), request)
}

// GetDocPreview
// 获取文档预览信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetDocPreviewWithContext(ctx context.Context, request *GetDocPreviewRequest) (response *GetDocPreviewResponse, err error) {
    if request == nil {
        request = NewGetDocPreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GetDocPreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDocPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDocPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewGetLikeDataCountRequest() (request *GetLikeDataCountRequest) {
    request = &GetLikeDataCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetLikeDataCount")
    
    
    return
}

func NewGetLikeDataCountResponse() (response *GetLikeDataCountResponse) {
    response = &GetLikeDataCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetLikeDataCount
// 点赞点踩数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetLikeDataCount(request *GetLikeDataCountRequest) (response *GetLikeDataCountResponse, err error) {
    return c.GetLikeDataCountWithContext(context.Background(), request)
}

// GetLikeDataCount
// 点赞点踩数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetLikeDataCountWithContext(ctx context.Context, request *GetLikeDataCountRequest) (response *GetLikeDataCountResponse, err error) {
    if request == nil {
        request = NewGetLikeDataCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GetLikeDataCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLikeDataCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLikeDataCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetMsgRecordRequest() (request *GetMsgRecordRequest) {
    request = &GetMsgRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetMsgRecord")
    
    
    return
}

func NewGetMsgRecordResponse() (response *GetMsgRecordResponse) {
    response = &GetMsgRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMsgRecord
// 获取聊天历史
//
// 根据会话session id获取聊天历史（仅保留180天内的历史对话数据）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetMsgRecord(request *GetMsgRecordRequest) (response *GetMsgRecordResponse, err error) {
    return c.GetMsgRecordWithContext(context.Background(), request)
}

// GetMsgRecord
// 获取聊天历史
//
// 根据会话session id获取聊天历史（仅保留180天内的历史对话数据）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetMsgRecordWithContext(ctx context.Context, request *GetMsgRecordRequest) (response *GetMsgRecordResponse, err error) {
    if request == nil {
        request = NewGetMsgRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GetMsgRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMsgRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMsgRecordResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskStatusRequest() (request *GetTaskStatusRequest) {
    request = &GetTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetTaskStatus")
    
    
    return
}

func NewGetTaskStatusResponse() (response *GetTaskStatusResponse) {
    response = &GetTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskStatus
// 获取任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskStatus(request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    return c.GetTaskStatusWithContext(context.Background(), request)
}

// GetTaskStatus
// 获取任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskStatusWithContext(ctx context.Context, request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    if request == nil {
        request = NewGetTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GetTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetVarListRequest() (request *GetVarListRequest) {
    request = &GetVarListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetVarList")
    
    
    return
}

func NewGetVarListResponse() (response *GetVarListResponse) {
    response = &GetVarListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetVarList
// 查询自定义变量列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetVarList(request *GetVarListRequest) (response *GetVarListResponse, err error) {
    return c.GetVarListWithContext(context.Background(), request)
}

// GetVarList
// 查询自定义变量列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetVarListWithContext(ctx context.Context, request *GetVarListRequest) (response *GetVarListResponse, err error) {
    if request == nil {
        request = NewGetVarListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GetVarList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetVarList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetVarListResponse()
    err = c.Send(request, response)
    return
}

func NewGetWsTokenRequest() (request *GetWsTokenRequest) {
    request = &GetWsTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetWsToken")
    
    
    return
}

func NewGetWsTokenResponse() (response *GetWsTokenResponse) {
    response = &GetWsTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWsToken
// 获取ws token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetWsToken(request *GetWsTokenRequest) (response *GetWsTokenResponse, err error) {
    return c.GetWsTokenWithContext(context.Background(), request)
}

// GetWsToken
// 获取ws token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetWsTokenWithContext(ctx context.Context, request *GetWsTokenRequest) (response *GetWsTokenResponse, err error) {
    if request == nil {
        request = NewGetWsTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GetWsToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWsToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWsTokenResponse()
    err = c.Send(request, response)
    return
}

func NewGroupDocRequest() (request *GroupDocRequest) {
    request = &GroupDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GroupDoc")
    
    
    return
}

func NewGroupDocResponse() (response *GroupDocResponse) {
    response = &GroupDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GroupDoc
// Doc分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GroupDoc(request *GroupDocRequest) (response *GroupDocResponse, err error) {
    return c.GroupDocWithContext(context.Background(), request)
}

// GroupDoc
// Doc分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GroupDocWithContext(ctx context.Context, request *GroupDocRequest) (response *GroupDocResponse, err error) {
    if request == nil {
        request = NewGroupDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GroupDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GroupDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewGroupDocResponse()
    err = c.Send(request, response)
    return
}

func NewGroupQARequest() (request *GroupQARequest) {
    request = &GroupQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GroupQA")
    
    
    return
}

func NewGroupQAResponse() (response *GroupQAResponse) {
    response = &GroupQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GroupQA
// 用户将多个问答批量的分类到知识库的具体分类
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GroupQA(request *GroupQARequest) (response *GroupQAResponse, err error) {
    return c.GroupQAWithContext(context.Background(), request)
}

// GroupQA
// 用户将多个问答批量的分类到知识库的具体分类
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GroupQAWithContext(ctx context.Context, request *GroupQARequest) (response *GroupQAResponse, err error) {
    if request == nil {
        request = NewGroupQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "GroupQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GroupQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewGroupQAResponse()
    err = c.Send(request, response)
    return
}

func NewIgnoreUnsatisfiedReplyRequest() (request *IgnoreUnsatisfiedReplyRequest) {
    request = &IgnoreUnsatisfiedReplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "IgnoreUnsatisfiedReply")
    
    
    return
}

func NewIgnoreUnsatisfiedReplyResponse() (response *IgnoreUnsatisfiedReplyResponse) {
    response = &IgnoreUnsatisfiedReplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IgnoreUnsatisfiedReply
// 忽略不满意回复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IgnoreUnsatisfiedReply(request *IgnoreUnsatisfiedReplyRequest) (response *IgnoreUnsatisfiedReplyResponse, err error) {
    return c.IgnoreUnsatisfiedReplyWithContext(context.Background(), request)
}

// IgnoreUnsatisfiedReply
// 忽略不满意回复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IgnoreUnsatisfiedReplyWithContext(ctx context.Context, request *IgnoreUnsatisfiedReplyRequest) (response *IgnoreUnsatisfiedReplyResponse, err error) {
    if request == nil {
        request = NewIgnoreUnsatisfiedReplyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "IgnoreUnsatisfiedReply")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IgnoreUnsatisfiedReply require credential")
    }

    request.SetContext(ctx)
    
    response = NewIgnoreUnsatisfiedReplyResponse()
    err = c.Send(request, response)
    return
}

func NewIsTransferIntentRequest() (request *IsTransferIntentRequest) {
    request = &IsTransferIntentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "IsTransferIntent")
    
    
    return
}

func NewIsTransferIntentResponse() (response *IsTransferIntentResponse) {
    response = &IsTransferIntentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsTransferIntent
// 是否意图转人工
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IsTransferIntent(request *IsTransferIntentRequest) (response *IsTransferIntentResponse, err error) {
    return c.IsTransferIntentWithContext(context.Background(), request)
}

// IsTransferIntent
// 是否意图转人工
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IsTransferIntentWithContext(ctx context.Context, request *IsTransferIntentRequest) (response *IsTransferIntentResponse, err error) {
    if request == nil {
        request = NewIsTransferIntentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "IsTransferIntent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsTransferIntent require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsTransferIntentResponse()
    err = c.Send(request, response)
    return
}

func NewListAppRequest() (request *ListAppRequest) {
    request = &ListAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListApp")
    
    
    return
}

func NewListAppResponse() (response *ListAppResponse) {
    response = &ListAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListApp
// 获取企业下应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListApp(request *ListAppRequest) (response *ListAppResponse, err error) {
    return c.ListAppWithContext(context.Background(), request)
}

// ListApp
// 获取企业下应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAppWithContext(ctx context.Context, request *ListAppRequest) (response *ListAppResponse, err error) {
    if request == nil {
        request = NewListAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAppResponse()
    err = c.Send(request, response)
    return
}

func NewListAppKnowledgeDetailRequest() (request *ListAppKnowledgeDetailRequest) {
    request = &ListAppKnowledgeDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListAppKnowledgeDetail")
    
    
    return
}

func NewListAppKnowledgeDetailResponse() (response *ListAppKnowledgeDetailResponse) {
    response = &ListAppKnowledgeDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAppKnowledgeDetail
// 列表查询知识库容量详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAppKnowledgeDetail(request *ListAppKnowledgeDetailRequest) (response *ListAppKnowledgeDetailResponse, err error) {
    return c.ListAppKnowledgeDetailWithContext(context.Background(), request)
}

// ListAppKnowledgeDetail
// 列表查询知识库容量详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAppKnowledgeDetailWithContext(ctx context.Context, request *ListAppKnowledgeDetailRequest) (response *ListAppKnowledgeDetailResponse, err error) {
    if request == nil {
        request = NewListAppKnowledgeDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListAppKnowledgeDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAppKnowledgeDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAppKnowledgeDetailResponse()
    err = c.Send(request, response)
    return
}

func NewListAttributeLabelRequest() (request *ListAttributeLabelRequest) {
    request = &ListAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListAttributeLabel")
    
    
    return
}

func NewListAttributeLabelResponse() (response *ListAttributeLabelResponse) {
    response = &ListAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAttributeLabel
// 查询属性标签列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAttributeLabel(request *ListAttributeLabelRequest) (response *ListAttributeLabelResponse, err error) {
    return c.ListAttributeLabelWithContext(context.Background(), request)
}

// ListAttributeLabel
// 查询属性标签列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAttributeLabelWithContext(ctx context.Context, request *ListAttributeLabelRequest) (response *ListAttributeLabelResponse, err error) {
    if request == nil {
        request = NewListAttributeLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListAttributeLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewListChannelRequest() (request *ListChannelRequest) {
    request = &ListChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListChannel")
    
    
    return
}

func NewListChannelResponse() (response *ListChannelResponse) {
    response = &ListChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListChannel
// 获取发布渠道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListChannel(request *ListChannelRequest) (response *ListChannelResponse, err error) {
    return c.ListChannelWithContext(context.Background(), request)
}

// ListChannel
// 获取发布渠道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListChannelWithContext(ctx context.Context, request *ListChannelRequest) (response *ListChannelResponse, err error) {
    if request == nil {
        request = NewListChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewListChannelResponse()
    err = c.Send(request, response)
    return
}

func NewListDocRequest() (request *ListDocRequest) {
    request = &ListDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListDoc")
    
    
    return
}

func NewListDocResponse() (response *ListDocResponse) {
    response = &ListDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDoc
// 文档列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDoc(request *ListDocRequest) (response *ListDocResponse, err error) {
    return c.ListDocWithContext(context.Background(), request)
}

// ListDoc
// 文档列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDocWithContext(ctx context.Context, request *ListDocRequest) (response *ListDocResponse, err error) {
    if request == nil {
        request = NewListDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDocResponse()
    err = c.Send(request, response)
    return
}

func NewListDocCateRequest() (request *ListDocCateRequest) {
    request = &ListDocCateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListDocCate")
    
    
    return
}

func NewListDocCateResponse() (response *ListDocCateResponse) {
    response = &ListDocCateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDocCate
// 获取Doc分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDocCate(request *ListDocCateRequest) (response *ListDocCateResponse, err error) {
    return c.ListDocCateWithContext(context.Background(), request)
}

// ListDocCate
// 获取Doc分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDocCateWithContext(ctx context.Context, request *ListDocCateRequest) (response *ListDocCateResponse, err error) {
    if request == nil {
        request = NewListDocCateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListDocCate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDocCate require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDocCateResponse()
    err = c.Send(request, response)
    return
}

func NewListModelRequest() (request *ListModelRequest) {
    request = &ListModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListModel")
    
    
    return
}

func NewListModelResponse() (response *ListModelResponse) {
    response = &ListModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListModel
// 获取模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListModel(request *ListModelRequest) (response *ListModelResponse, err error) {
    return c.ListModelWithContext(context.Background(), request)
}

// ListModel
// 获取模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListModelWithContext(ctx context.Context, request *ListModelRequest) (response *ListModelResponse, err error) {
    if request == nil {
        request = NewListModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewListModelResponse()
    err = c.Send(request, response)
    return
}

func NewListQARequest() (request *ListQARequest) {
    request = &ListQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListQA")
    
    
    return
}

func NewListQAResponse() (response *ListQAResponse) {
    response = &ListQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQA
// 问答列表
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQA(request *ListQARequest) (response *ListQAResponse, err error) {
    return c.ListQAWithContext(context.Background(), request)
}

// ListQA
// 问答列表
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQAWithContext(ctx context.Context, request *ListQARequest) (response *ListQAResponse, err error) {
    if request == nil {
        request = NewListQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQAResponse()
    err = c.Send(request, response)
    return
}

func NewListQACateRequest() (request *ListQACateRequest) {
    request = &ListQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListQACate")
    
    
    return
}

func NewListQACateResponse() (response *ListQACateResponse) {
    response = &ListQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQACate
// 获取QA分类
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQACate(request *ListQACateRequest) (response *ListQACateResponse, err error) {
    return c.ListQACateWithContext(context.Background(), request)
}

// ListQACate
// 获取QA分类
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQACateWithContext(ctx context.Context, request *ListQACateRequest) (response *ListQACateResponse, err error) {
    if request == nil {
        request = NewListQACateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListQACate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQACateResponse()
    err = c.Send(request, response)
    return
}

func NewListReferShareKnowledgeRequest() (request *ListReferShareKnowledgeRequest) {
    request = &ListReferShareKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListReferShareKnowledge")
    
    
    return
}

func NewListReferShareKnowledgeResponse() (response *ListReferShareKnowledgeResponse) {
    response = &ListReferShareKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListReferShareKnowledge
// 查看应用引用了哪些共享知识库，可以看到共享知识库的基础信息，包括名称，id等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReferShareKnowledge(request *ListReferShareKnowledgeRequest) (response *ListReferShareKnowledgeResponse, err error) {
    return c.ListReferShareKnowledgeWithContext(context.Background(), request)
}

// ListReferShareKnowledge
// 查看应用引用了哪些共享知识库，可以看到共享知识库的基础信息，包括名称，id等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReferShareKnowledgeWithContext(ctx context.Context, request *ListReferShareKnowledgeRequest) (response *ListReferShareKnowledgeResponse, err error) {
    if request == nil {
        request = NewListReferShareKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListReferShareKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListReferShareKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReferShareKnowledgeResponse()
    err = c.Send(request, response)
    return
}

func NewListRejectedQuestionRequest() (request *ListRejectedQuestionRequest) {
    request = &ListRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListRejectedQuestion")
    
    
    return
}

func NewListRejectedQuestionResponse() (response *ListRejectedQuestionResponse) {
    response = &ListRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRejectedQuestion
// 获取拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestion(request *ListRejectedQuestionRequest) (response *ListRejectedQuestionResponse, err error) {
    return c.ListRejectedQuestionWithContext(context.Background(), request)
}

// ListRejectedQuestion
// 获取拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestionWithContext(ctx context.Context, request *ListRejectedQuestionRequest) (response *ListRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewListRejectedQuestionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListRejectedQuestion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewListRejectedQuestionPreviewRequest() (request *ListRejectedQuestionPreviewRequest) {
    request = &ListRejectedQuestionPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListRejectedQuestionPreview")
    
    
    return
}

func NewListRejectedQuestionPreviewResponse() (response *ListRejectedQuestionPreviewResponse) {
    response = &ListRejectedQuestionPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRejectedQuestionPreview
// 发布拒答问题预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestionPreview(request *ListRejectedQuestionPreviewRequest) (response *ListRejectedQuestionPreviewResponse, err error) {
    return c.ListRejectedQuestionPreviewWithContext(context.Background(), request)
}

// ListRejectedQuestionPreview
// 发布拒答问题预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestionPreviewWithContext(ctx context.Context, request *ListRejectedQuestionPreviewRequest) (response *ListRejectedQuestionPreviewResponse, err error) {
    if request == nil {
        request = NewListRejectedQuestionPreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListRejectedQuestionPreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRejectedQuestionPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRejectedQuestionPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewListReleaseRequest() (request *ListReleaseRequest) {
    request = &ListReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListRelease")
    
    
    return
}

func NewListReleaseResponse() (response *ListReleaseResponse) {
    response = &ListReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRelease
// 发布列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRelease(request *ListReleaseRequest) (response *ListReleaseResponse, err error) {
    return c.ListReleaseWithContext(context.Background(), request)
}

// ListRelease
// 发布列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseWithContext(ctx context.Context, request *ListReleaseRequest) (response *ListReleaseResponse, err error) {
    if request == nil {
        request = NewListReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewListReleaseConfigPreviewRequest() (request *ListReleaseConfigPreviewRequest) {
    request = &ListReleaseConfigPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListReleaseConfigPreview")
    
    
    return
}

func NewListReleaseConfigPreviewResponse() (response *ListReleaseConfigPreviewResponse) {
    response = &ListReleaseConfigPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListReleaseConfigPreview
// 发布配置项预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseConfigPreview(request *ListReleaseConfigPreviewRequest) (response *ListReleaseConfigPreviewResponse, err error) {
    return c.ListReleaseConfigPreviewWithContext(context.Background(), request)
}

// ListReleaseConfigPreview
// 发布配置项预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseConfigPreviewWithContext(ctx context.Context, request *ListReleaseConfigPreviewRequest) (response *ListReleaseConfigPreviewResponse, err error) {
    if request == nil {
        request = NewListReleaseConfigPreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListReleaseConfigPreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListReleaseConfigPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReleaseConfigPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewListReleaseDocPreviewRequest() (request *ListReleaseDocPreviewRequest) {
    request = &ListReleaseDocPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListReleaseDocPreview")
    
    
    return
}

func NewListReleaseDocPreviewResponse() (response *ListReleaseDocPreviewResponse) {
    response = &ListReleaseDocPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListReleaseDocPreview
// 发布文档预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseDocPreview(request *ListReleaseDocPreviewRequest) (response *ListReleaseDocPreviewResponse, err error) {
    return c.ListReleaseDocPreviewWithContext(context.Background(), request)
}

// ListReleaseDocPreview
// 发布文档预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseDocPreviewWithContext(ctx context.Context, request *ListReleaseDocPreviewRequest) (response *ListReleaseDocPreviewResponse, err error) {
    if request == nil {
        request = NewListReleaseDocPreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListReleaseDocPreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListReleaseDocPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReleaseDocPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewListReleaseQAPreviewRequest() (request *ListReleaseQAPreviewRequest) {
    request = &ListReleaseQAPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListReleaseQAPreview")
    
    
    return
}

func NewListReleaseQAPreviewResponse() (response *ListReleaseQAPreviewResponse) {
    response = &ListReleaseQAPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListReleaseQAPreview
// 文档列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseQAPreview(request *ListReleaseQAPreviewRequest) (response *ListReleaseQAPreviewResponse, err error) {
    return c.ListReleaseQAPreviewWithContext(context.Background(), request)
}

// ListReleaseQAPreview
// 文档列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseQAPreviewWithContext(ctx context.Context, request *ListReleaseQAPreviewRequest) (response *ListReleaseQAPreviewResponse, err error) {
    if request == nil {
        request = NewListReleaseQAPreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListReleaseQAPreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListReleaseQAPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReleaseQAPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewListSelectDocRequest() (request *ListSelectDocRequest) {
    request = &ListSelectDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListSelectDoc")
    
    
    return
}

func NewListSelectDocResponse() (response *ListSelectDocResponse) {
    response = &ListSelectDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSelectDoc
// 获取文档下拉列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSelectDoc(request *ListSelectDocRequest) (response *ListSelectDocResponse, err error) {
    return c.ListSelectDocWithContext(context.Background(), request)
}

// ListSelectDoc
// 获取文档下拉列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSelectDocWithContext(ctx context.Context, request *ListSelectDocRequest) (response *ListSelectDocResponse, err error) {
    if request == nil {
        request = NewListSelectDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListSelectDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSelectDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSelectDocResponse()
    err = c.Send(request, response)
    return
}

func NewListSharedKnowledgeRequest() (request *ListSharedKnowledgeRequest) {
    request = &ListSharedKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListSharedKnowledge")
    
    
    return
}

func NewListSharedKnowledgeResponse() (response *ListSharedKnowledgeResponse) {
    response = &ListSharedKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSharedKnowledge
// 列举共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSharedKnowledge(request *ListSharedKnowledgeRequest) (response *ListSharedKnowledgeResponse, err error) {
    return c.ListSharedKnowledgeWithContext(context.Background(), request)
}

// ListSharedKnowledge
// 列举共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSharedKnowledgeWithContext(ctx context.Context, request *ListSharedKnowledgeRequest) (response *ListSharedKnowledgeResponse, err error) {
    if request == nil {
        request = NewListSharedKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListSharedKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSharedKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSharedKnowledgeResponse()
    err = c.Send(request, response)
    return
}

func NewListUnsatisfiedReplyRequest() (request *ListUnsatisfiedReplyRequest) {
    request = &ListUnsatisfiedReplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListUnsatisfiedReply")
    
    
    return
}

func NewListUnsatisfiedReplyResponse() (response *ListUnsatisfiedReplyResponse) {
    response = &ListUnsatisfiedReplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUnsatisfiedReply
// 查询不满意回复列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListUnsatisfiedReply(request *ListUnsatisfiedReplyRequest) (response *ListUnsatisfiedReplyResponse, err error) {
    return c.ListUnsatisfiedReplyWithContext(context.Background(), request)
}

// ListUnsatisfiedReply
// 查询不满意回复列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListUnsatisfiedReplyWithContext(ctx context.Context, request *ListUnsatisfiedReplyRequest) (response *ListUnsatisfiedReplyResponse, err error) {
    if request == nil {
        request = NewListUnsatisfiedReplyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListUnsatisfiedReply")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUnsatisfiedReply require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUnsatisfiedReplyResponse()
    err = c.Send(request, response)
    return
}

func NewListUsageCallDetailRequest() (request *ListUsageCallDetailRequest) {
    request = &ListUsageCallDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListUsageCallDetail")
    
    
    return
}

func NewListUsageCallDetailResponse() (response *ListUsageCallDetailResponse) {
    response = &ListUsageCallDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUsageCallDetail
// 列表查询单次调用明细
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListUsageCallDetail(request *ListUsageCallDetailRequest) (response *ListUsageCallDetailResponse, err error) {
    return c.ListUsageCallDetailWithContext(context.Background(), request)
}

// ListUsageCallDetail
// 列表查询单次调用明细
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListUsageCallDetailWithContext(ctx context.Context, request *ListUsageCallDetailRequest) (response *ListUsageCallDetailResponse, err error) {
    if request == nil {
        request = NewListUsageCallDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListUsageCallDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsageCallDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsageCallDetailResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowRunsRequest() (request *ListWorkflowRunsRequest) {
    request = &ListWorkflowRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListWorkflowRuns")
    
    
    return
}

func NewListWorkflowRunsResponse() (response *ListWorkflowRunsResponse) {
    response = &ListWorkflowRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflowRuns
// 此接口可查询已创建的所有工作流异步运行实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflowRuns(request *ListWorkflowRunsRequest) (response *ListWorkflowRunsResponse, err error) {
    return c.ListWorkflowRunsWithContext(context.Background(), request)
}

// ListWorkflowRuns
// 此接口可查询已创建的所有工作流异步运行实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflowRunsWithContext(ctx context.Context, request *ListWorkflowRunsRequest) (response *ListWorkflowRunsResponse, err error) {
    if request == nil {
        request = NewListWorkflowRunsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ListWorkflowRuns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflowRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowRunsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppRequest() (request *ModifyAppRequest) {
    request = &ModifyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyApp")
    
    
    return
}

func NewModifyAppResponse() (response *ModifyAppResponse) {
    response = &ModifyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApp
// 修改应用请求结构体
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    return c.ModifyAppWithContext(context.Background(), request)
}

// ModifyApp
// 修改应用请求结构体
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    if request == nil {
        request = NewModifyAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ModifyApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAttributeLabelRequest() (request *ModifyAttributeLabelRequest) {
    request = &ModifyAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyAttributeLabel")
    
    
    return
}

func NewModifyAttributeLabelResponse() (response *ModifyAttributeLabelResponse) {
    response = &ModifyAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAttributeLabel
// 编辑属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAttributeLabel(request *ModifyAttributeLabelRequest) (response *ModifyAttributeLabelResponse, err error) {
    return c.ModifyAttributeLabelWithContext(context.Background(), request)
}

// ModifyAttributeLabel
// 编辑属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAttributeLabelWithContext(ctx context.Context, request *ModifyAttributeLabelRequest) (response *ModifyAttributeLabelResponse, err error) {
    if request == nil {
        request = NewModifyAttributeLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ModifyAttributeLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDocRequest() (request *ModifyDocRequest) {
    request = &ModifyDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyDoc")
    
    
    return
}

func NewModifyDocResponse() (response *ModifyDocResponse) {
    response = &ModifyDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDoc
// 修改文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyDoc(request *ModifyDocRequest) (response *ModifyDocResponse, err error) {
    return c.ModifyDocWithContext(context.Background(), request)
}

// ModifyDoc
// 修改文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyDocWithContext(ctx context.Context, request *ModifyDocRequest) (response *ModifyDocResponse, err error) {
    if request == nil {
        request = NewModifyDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ModifyDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDocAttrRangeRequest() (request *ModifyDocAttrRangeRequest) {
    request = &ModifyDocAttrRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyDocAttrRange")
    
    
    return
}

func NewModifyDocAttrRangeResponse() (response *ModifyDocAttrRangeResponse) {
    response = &ModifyDocAttrRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDocAttrRange
// 批量修改文档适用范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocAttrRange(request *ModifyDocAttrRangeRequest) (response *ModifyDocAttrRangeResponse, err error) {
    return c.ModifyDocAttrRangeWithContext(context.Background(), request)
}

// ModifyDocAttrRange
// 批量修改文档适用范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocAttrRangeWithContext(ctx context.Context, request *ModifyDocAttrRangeRequest) (response *ModifyDocAttrRangeResponse, err error) {
    if request == nil {
        request = NewModifyDocAttrRangeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ModifyDocAttrRange")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDocAttrRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocAttrRangeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDocCateRequest() (request *ModifyDocCateRequest) {
    request = &ModifyDocCateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyDocCate")
    
    
    return
}

func NewModifyDocCateResponse() (response *ModifyDocCateResponse) {
    response = &ModifyDocCateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDocCate
// 修改Doc分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocCate(request *ModifyDocCateRequest) (response *ModifyDocCateResponse, err error) {
    return c.ModifyDocCateWithContext(context.Background(), request)
}

// ModifyDocCate
// 修改Doc分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocCateWithContext(ctx context.Context, request *ModifyDocCateRequest) (response *ModifyDocCateResponse, err error) {
    if request == nil {
        request = NewModifyDocCateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ModifyDocCate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDocCate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocCateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQARequest() (request *ModifyQARequest) {
    request = &ModifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyQA")
    
    
    return
}

func NewModifyQAResponse() (response *ModifyQAResponse) {
    response = &ModifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQA
// 更新问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQA(request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    return c.ModifyQAWithContext(context.Background(), request)
}

// ModifyQA
// 更新问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQAWithContext(ctx context.Context, request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    if request == nil {
        request = NewModifyQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ModifyQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQAAttrRangeRequest() (request *ModifyQAAttrRangeRequest) {
    request = &ModifyQAAttrRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyQAAttrRange")
    
    
    return
}

func NewModifyQAAttrRangeResponse() (response *ModifyQAAttrRangeResponse) {
    response = &ModifyQAAttrRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQAAttrRange
// 批量修改问答适用范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQAAttrRange(request *ModifyQAAttrRangeRequest) (response *ModifyQAAttrRangeResponse, err error) {
    return c.ModifyQAAttrRangeWithContext(context.Background(), request)
}

// ModifyQAAttrRange
// 批量修改问答适用范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQAAttrRangeWithContext(ctx context.Context, request *ModifyQAAttrRangeRequest) (response *ModifyQAAttrRangeResponse, err error) {
    if request == nil {
        request = NewModifyQAAttrRangeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ModifyQAAttrRange")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQAAttrRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAAttrRangeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQACateRequest() (request *ModifyQACateRequest) {
    request = &ModifyQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyQACate")
    
    
    return
}

func NewModifyQACateResponse() (response *ModifyQACateResponse) {
    response = &ModifyQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQACate
// 更新QA分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQACate(request *ModifyQACateRequest) (response *ModifyQACateResponse, err error) {
    return c.ModifyQACateWithContext(context.Background(), request)
}

// ModifyQACate
// 更新QA分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQACateWithContext(ctx context.Context, request *ModifyQACateRequest) (response *ModifyQACateResponse, err error) {
    if request == nil {
        request = NewModifyQACateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ModifyQACate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQACateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRejectedQuestionRequest() (request *ModifyRejectedQuestionRequest) {
    request = &ModifyRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyRejectedQuestion")
    
    
    return
}

func NewModifyRejectedQuestionResponse() (response *ModifyRejectedQuestionResponse) {
    response = &ModifyRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRejectedQuestion
// 修改拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyRejectedQuestion(request *ModifyRejectedQuestionRequest) (response *ModifyRejectedQuestionResponse, err error) {
    return c.ModifyRejectedQuestionWithContext(context.Background(), request)
}

// ModifyRejectedQuestion
// 修改拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyRejectedQuestionWithContext(ctx context.Context, request *ModifyRejectedQuestionRequest) (response *ModifyRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewModifyRejectedQuestionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ModifyRejectedQuestion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewRateMsgRecordRequest() (request *RateMsgRecordRequest) {
    request = &RateMsgRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RateMsgRecord")
    
    
    return
}

func NewRateMsgRecordResponse() (response *RateMsgRecordResponse) {
    response = &RateMsgRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RateMsgRecord
// 点赞点踩消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RateMsgRecord(request *RateMsgRecordRequest) (response *RateMsgRecordResponse, err error) {
    return c.RateMsgRecordWithContext(context.Background(), request)
}

// RateMsgRecord
// 点赞点踩消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RateMsgRecordWithContext(ctx context.Context, request *RateMsgRecordRequest) (response *RateMsgRecordResponse, err error) {
    if request == nil {
        request = NewRateMsgRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "RateMsgRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RateMsgRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewRateMsgRecordResponse()
    err = c.Send(request, response)
    return
}

func NewReferShareKnowledgeRequest() (request *ReferShareKnowledgeRequest) {
    request = &ReferShareKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ReferShareKnowledge")
    
    
    return
}

func NewReferShareKnowledgeResponse() (response *ReferShareKnowledgeResponse) {
    response = &ReferShareKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReferShareKnowledge
// 应用引用共享知识库，可以引用一个或多个，每次都是全量覆盖
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ReferShareKnowledge(request *ReferShareKnowledgeRequest) (response *ReferShareKnowledgeResponse, err error) {
    return c.ReferShareKnowledgeWithContext(context.Background(), request)
}

// ReferShareKnowledge
// 应用引用共享知识库，可以引用一个或多个，每次都是全量覆盖
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ReferShareKnowledgeWithContext(ctx context.Context, request *ReferShareKnowledgeRequest) (response *ReferShareKnowledgeResponse, err error) {
    if request == nil {
        request = NewReferShareKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "ReferShareKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReferShareKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewReferShareKnowledgeResponse()
    err = c.Send(request, response)
    return
}

func NewRenameDocRequest() (request *RenameDocRequest) {
    request = &RenameDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RenameDoc")
    
    
    return
}

func NewRenameDocResponse() (response *RenameDocResponse) {
    response = &RenameDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenameDoc
// 文档重命名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RenameDoc(request *RenameDocRequest) (response *RenameDocResponse, err error) {
    return c.RenameDocWithContext(context.Background(), request)
}

// RenameDoc
// 文档重命名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RenameDocWithContext(ctx context.Context, request *RenameDocRequest) (response *RenameDocResponse, err error) {
    if request == nil {
        request = NewRenameDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "RenameDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenameDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenameDocResponse()
    err = c.Send(request, response)
    return
}

func NewRetryDocAuditRequest() (request *RetryDocAuditRequest) {
    request = &RetryDocAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RetryDocAudit")
    
    
    return
}

func NewRetryDocAuditResponse() (response *RetryDocAuditResponse) {
    response = &RetryDocAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryDocAudit
// 文档解析重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocAudit(request *RetryDocAuditRequest) (response *RetryDocAuditResponse, err error) {
    return c.RetryDocAuditWithContext(context.Background(), request)
}

// RetryDocAudit
// 文档解析重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocAuditWithContext(ctx context.Context, request *RetryDocAuditRequest) (response *RetryDocAuditResponse, err error) {
    if request == nil {
        request = NewRetryDocAuditRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "RetryDocAudit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryDocAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryDocAuditResponse()
    err = c.Send(request, response)
    return
}

func NewRetryDocParseRequest() (request *RetryDocParseRequest) {
    request = &RetryDocParseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RetryDocParse")
    
    
    return
}

func NewRetryDocParseResponse() (response *RetryDocParseResponse) {
    response = &RetryDocParseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryDocParse
// 文档解析重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocParse(request *RetryDocParseRequest) (response *RetryDocParseResponse, err error) {
    return c.RetryDocParseWithContext(context.Background(), request)
}

// RetryDocParse
// 文档解析重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocParseWithContext(ctx context.Context, request *RetryDocParseRequest) (response *RetryDocParseResponse, err error) {
    if request == nil {
        request = NewRetryDocParseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "RetryDocParse")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryDocParse require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryDocParseResponse()
    err = c.Send(request, response)
    return
}

func NewRetryReleaseRequest() (request *RetryReleaseRequest) {
    request = &RetryReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RetryRelease")
    
    
    return
}

func NewRetryReleaseResponse() (response *RetryReleaseResponse) {
    response = &RetryReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryRelease
// 发布暂停后重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryRelease(request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    return c.RetryReleaseWithContext(context.Background(), request)
}

// RetryRelease
// 发布暂停后重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryReleaseWithContext(ctx context.Context, request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    if request == nil {
        request = NewRetryReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "RetryRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewSaveDocRequest() (request *SaveDocRequest) {
    request = &SaveDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "SaveDoc")
    
    
    return
}

func NewSaveDocResponse() (response *SaveDocResponse) {
    response = &SaveDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SaveDoc
// 知识库文档问答保存。
//
// 将文件存储到应用的知识库内需要三步：
//
// 1.获取临时密钥，参考[接口文档](https://cloud.tencent.com/document/product/1759/105050)。获取临时密钥不同参数组合权限不一样，可参考 [智能体开发平台操作 cos 指南](https://cloud.tencent.com/document/product/1759/116238)
//
// 2.调用腾讯云提供的 cos 存储接口，将文件存储到智能体开发平台 cos 中：具体可参考[ COS SDK 概览](https://cloud.tencent.com/document/product/436/6474), 注意使用的是临时密钥的方式操作 COS 
//
// 3.调用本接口，将文件的基础信息存储到智能体开发平台中。
//
// 以上步骤可参考[文档](https://cloud.tencent.com/document/product/1759/108903)，文档最后有[代码demo](https://cloud.tencent.com/document/product/1759/108903#demo)，可作为参考。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveDoc(request *SaveDocRequest) (response *SaveDocResponse, err error) {
    return c.SaveDocWithContext(context.Background(), request)
}

// SaveDoc
// 知识库文档问答保存。
//
// 将文件存储到应用的知识库内需要三步：
//
// 1.获取临时密钥，参考[接口文档](https://cloud.tencent.com/document/product/1759/105050)。获取临时密钥不同参数组合权限不一样，可参考 [智能体开发平台操作 cos 指南](https://cloud.tencent.com/document/product/1759/116238)
//
// 2.调用腾讯云提供的 cos 存储接口，将文件存储到智能体开发平台 cos 中：具体可参考[ COS SDK 概览](https://cloud.tencent.com/document/product/436/6474), 注意使用的是临时密钥的方式操作 COS 
//
// 3.调用本接口，将文件的基础信息存储到智能体开发平台中。
//
// 以上步骤可参考[文档](https://cloud.tencent.com/document/product/1759/108903)，文档最后有[代码demo](https://cloud.tencent.com/document/product/1759/108903#demo)，可作为参考。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveDocWithContext(ctx context.Context, request *SaveDocRequest) (response *SaveDocResponse, err error) {
    if request == nil {
        request = NewSaveDocRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "SaveDoc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaveDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaveDocResponse()
    err = c.Send(request, response)
    return
}

func NewStopDocParseRequest() (request *StopDocParseRequest) {
    request = &StopDocParseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "StopDocParse")
    
    
    return
}

func NewStopDocParseResponse() (response *StopDocParseResponse) {
    response = &StopDocParseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopDocParse
// 终止文档解析
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopDocParse(request *StopDocParseRequest) (response *StopDocParseResponse, err error) {
    return c.StopDocParseWithContext(context.Background(), request)
}

// StopDocParse
// 终止文档解析
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopDocParseWithContext(ctx context.Context, request *StopDocParseRequest) (response *StopDocParseResponse, err error) {
    if request == nil {
        request = NewStopDocParseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "StopDocParse")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopDocParse require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopDocParseResponse()
    err = c.Send(request, response)
    return
}

func NewStopWorkflowRunRequest() (request *StopWorkflowRunRequest) {
    request = &StopWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "StopWorkflowRun")
    
    
    return
}

func NewStopWorkflowRunResponse() (response *StopWorkflowRunResponse) {
    response = &StopWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopWorkflowRun
// 此接口用来停止正在进行的工作流异步运行实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) StopWorkflowRun(request *StopWorkflowRunRequest) (response *StopWorkflowRunResponse, err error) {
    return c.StopWorkflowRunWithContext(context.Background(), request)
}

// StopWorkflowRun
// 此接口用来停止正在进行的工作流异步运行实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) StopWorkflowRunWithContext(ctx context.Context, request *StopWorkflowRunRequest) (response *StopWorkflowRunResponse, err error) {
    if request == nil {
        request = NewStopWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "StopWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSharedKnowledgeRequest() (request *UpdateSharedKnowledgeRequest) {
    request = &UpdateSharedKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "UpdateSharedKnowledge")
    
    
    return
}

func NewUpdateSharedKnowledgeResponse() (response *UpdateSharedKnowledgeResponse) {
    response = &UpdateSharedKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSharedKnowledge
// 更新共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateSharedKnowledge(request *UpdateSharedKnowledgeRequest) (response *UpdateSharedKnowledgeResponse, err error) {
    return c.UpdateSharedKnowledgeWithContext(context.Background(), request)
}

// UpdateSharedKnowledge
// 更新共享知识库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateSharedKnowledgeWithContext(ctx context.Context, request *UpdateSharedKnowledgeRequest) (response *UpdateSharedKnowledgeResponse, err error) {
    if request == nil {
        request = NewUpdateSharedKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "UpdateSharedKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSharedKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSharedKnowledgeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateVarRequest() (request *UpdateVarRequest) {
    request = &UpdateVarRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "UpdateVar")
    
    
    return
}

func NewUpdateVarResponse() (response *UpdateVarResponse) {
    response = &UpdateVarResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateVar
// 更新变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateVar(request *UpdateVarRequest) (response *UpdateVarResponse, err error) {
    return c.UpdateVarWithContext(context.Background(), request)
}

// UpdateVar
// 更新变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateVarWithContext(ctx context.Context, request *UpdateVarRequest) (response *UpdateVarResponse, err error) {
    if request == nil {
        request = NewUpdateVarRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "UpdateVar")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateVar require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateVarResponse()
    err = c.Send(request, response)
    return
}

func NewUploadAttributeLabelRequest() (request *UploadAttributeLabelRequest) {
    request = &UploadAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "UploadAttributeLabel")
    
    
    return
}

func NewUploadAttributeLabelResponse() (response *UploadAttributeLabelResponse) {
    response = &UploadAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadAttributeLabel
// 上传导入属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UploadAttributeLabel(request *UploadAttributeLabelRequest) (response *UploadAttributeLabelResponse, err error) {
    return c.UploadAttributeLabelWithContext(context.Background(), request)
}

// UploadAttributeLabel
// 上传导入属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UploadAttributeLabelWithContext(ctx context.Context, request *UploadAttributeLabelRequest) (response *UploadAttributeLabelResponse, err error) {
    if request == nil {
        request = NewUploadAttributeLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "UploadAttributeLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyQARequest() (request *VerifyQARequest) {
    request = &VerifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "VerifyQA")
    
    
    return
}

func NewVerifyQAResponse() (response *VerifyQAResponse) {
    response = &VerifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyQA
// 校验问答
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) VerifyQA(request *VerifyQARequest) (response *VerifyQAResponse, err error) {
    return c.VerifyQAWithContext(context.Background(), request)
}

// VerifyQA
// 校验问答
//
// 知识库相关背景知识介绍
//
// “知识库检索范围”文档：https://cloud.tencent.com/document/product/1759/112704
//
// “标签”文档：https://cloud.tencent.com/document/product/1759/112956
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) VerifyQAWithContext(ctx context.Context, request *VerifyQARequest) (response *VerifyQAResponse, err error) {
    if request == nil {
        request = NewVerifyQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lke", APIVersion, "VerifyQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyQAResponse()
    err = c.Send(request, response)
    return
}
