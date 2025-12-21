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

package v20250717

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-07-17"

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


func NewContinueAgentWorkRequest() (request *ContinueAgentWorkRequest) {
    request = &ContinueAgentWorkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "ContinueAgentWork")
    
    
    return
}

func NewContinueAgentWorkResponse() (response *ContinueAgentWorkResponse) {
    response = &ContinueAgentWorkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ContinueAgentWork
// 本接口（ContinueAgentWork）用于重启智能体实例的值守任务，通常在用户需要重启时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ContinueAgentWork(request *ContinueAgentWorkRequest) (response *ContinueAgentWorkResponse, err error) {
    return c.ContinueAgentWorkWithContext(context.Background(), request)
}

// ContinueAgentWork
// 本接口（ContinueAgentWork）用于重启智能体实例的值守任务，通常在用户需要重启时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ContinueAgentWorkWithContext(ctx context.Context, request *ContinueAgentWorkRequest) (response *ContinueAgentWorkResponse, err error) {
    if request == nil {
        request = NewContinueAgentWorkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "ContinueAgentWork")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ContinueAgentWork require credential")
    }

    request.SetContext(ctx)
    
    response = NewContinueAgentWorkResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentInstanceRequest() (request *CreateAgentInstanceRequest) {
    request = &CreateAgentInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "CreateAgentInstance")
    
    
    return
}

func NewCreateAgentInstanceResponse() (response *CreateAgentInstanceResponse) {
    response = &CreateAgentInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgentInstance
// 本接口（CreateAgentInstance）用于创建一个智能体实例，通常在用户购买一个智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAgentInstance(request *CreateAgentInstanceRequest) (response *CreateAgentInstanceResponse, err error) {
    return c.CreateAgentInstanceWithContext(context.Background(), request)
}

// CreateAgentInstance
// 本接口（CreateAgentInstance）用于创建一个智能体实例，通常在用户购买一个智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAgentInstanceWithContext(ctx context.Context, request *CreateAgentInstanceRequest) (response *CreateAgentInstanceResponse, err error) {
    if request == nil {
        request = NewCreateAgentInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "CreateAgentInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChatCompletionRequest() (request *CreateChatCompletionRequest) {
    request = &CreateChatCompletionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "CreateChatCompletion")
    
    
    return
}

func NewCreateChatCompletionResponse() (response *CreateChatCompletionResponse) {
    response = &CreateChatCompletionResponse{} 
    return

}

// CreateChatCompletion
// 用于创建一次会话的SSE接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateChatCompletion(request *CreateChatCompletionRequest) (response *CreateChatCompletionResponse, err error) {
    return c.CreateChatCompletionWithContext(context.Background(), request)
}

// CreateChatCompletion
// 用于创建一次会话的SSE接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateChatCompletionWithContext(ctx context.Context, request *CreateChatCompletionRequest) (response *CreateChatCompletionResponse, err error) {
    if request == nil {
        request = NewCreateChatCompletionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "CreateChatCompletion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChatCompletion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChatCompletionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentDutyTaskDetailRequest() (request *DescribeAgentDutyTaskDetailRequest) {
    request = &DescribeAgentDutyTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "DescribeAgentDutyTaskDetail")
    
    
    return
}

func NewDescribeAgentDutyTaskDetailResponse() (response *DescribeAgentDutyTaskDetailResponse) {
    response = &DescribeAgentDutyTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentDutyTaskDetail
// 查询智能体值守任务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAgentDutyTaskDetail(request *DescribeAgentDutyTaskDetailRequest) (response *DescribeAgentDutyTaskDetailResponse, err error) {
    return c.DescribeAgentDutyTaskDetailWithContext(context.Background(), request)
}

// DescribeAgentDutyTaskDetail
// 查询智能体值守任务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAgentDutyTaskDetailWithContext(ctx context.Context, request *DescribeAgentDutyTaskDetailRequest) (response *DescribeAgentDutyTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAgentDutyTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "DescribeAgentDutyTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentDutyTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentDutyTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentDutyTasksRequest() (request *DescribeAgentDutyTasksRequest) {
    request = &DescribeAgentDutyTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "DescribeAgentDutyTasks")
    
    
    return
}

func NewDescribeAgentDutyTasksResponse() (response *DescribeAgentDutyTasksResponse) {
    response = &DescribeAgentDutyTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentDutyTasks
// 查询智能体值守任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAgentDutyTasks(request *DescribeAgentDutyTasksRequest) (response *DescribeAgentDutyTasksResponse, err error) {
    return c.DescribeAgentDutyTasksWithContext(context.Background(), request)
}

// DescribeAgentDutyTasks
// 查询智能体值守任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAgentDutyTasksWithContext(ctx context.Context, request *DescribeAgentDutyTasksRequest) (response *DescribeAgentDutyTasksResponse, err error) {
    if request == nil {
        request = NewDescribeAgentDutyTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "DescribeAgentDutyTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentDutyTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentDutyTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentInstanceRequest() (request *DescribeAgentInstanceRequest) {
    request = &DescribeAgentInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "DescribeAgentInstance")
    
    
    return
}

func NewDescribeAgentInstanceResponse() (response *DescribeAgentInstanceResponse) {
    response = &DescribeAgentInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentInstance
// 本接口（DescribeAgentInstance）用于查询智能体实例详情，通常在用户查询所购买的所有智能体实例详情时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgentInstance(request *DescribeAgentInstanceRequest) (response *DescribeAgentInstanceResponse, err error) {
    return c.DescribeAgentInstanceWithContext(context.Background(), request)
}

// DescribeAgentInstance
// 本接口（DescribeAgentInstance）用于查询智能体实例详情，通常在用户查询所购买的所有智能体实例详情时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgentInstanceWithContext(ctx context.Context, request *DescribeAgentInstanceRequest) (response *DescribeAgentInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeAgentInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "DescribeAgentInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentInstancesRequest() (request *DescribeAgentInstancesRequest) {
    request = &DescribeAgentInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "DescribeAgentInstances")
    
    
    return
}

func NewDescribeAgentInstancesResponse() (response *DescribeAgentInstancesResponse) {
    response = &DescribeAgentInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentInstances
// 本接口（DescribeAgentInstances）用于查询智能体实例列表，通常在用户查询所购买的所有智能体列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgentInstances(request *DescribeAgentInstancesRequest) (response *DescribeAgentInstancesResponse, err error) {
    return c.DescribeAgentInstancesWithContext(context.Background(), request)
}

// DescribeAgentInstances
// 本接口（DescribeAgentInstances）用于查询智能体实例列表，通常在用户查询所购买的所有智能体列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgentInstancesWithContext(ctx context.Context, request *DescribeAgentInstancesRequest) (response *DescribeAgentInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeAgentInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "DescribeAgentInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentsRequest() (request *DescribeAgentsRequest) {
    request = &DescribeAgentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "DescribeAgents")
    
    
    return
}

func NewDescribeAgentsResponse() (response *DescribeAgentsResponse) {
    response = &DescribeAgentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgents
// 本接口（DescribeAgents）用于查询智能体列表，通常在用户查询所购买的所有智能体列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgents(request *DescribeAgentsRequest) (response *DescribeAgentsResponse, err error) {
    return c.DescribeAgentsWithContext(context.Background(), request)
}

// DescribeAgents
// 本接口（DescribeAgents）用于查询智能体列表，通常在用户查询所购买的所有智能体列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgentsWithContext(ctx context.Context, request *DescribeAgentsRequest) (response *DescribeAgentsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "DescribeAgents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChatDetailRequest() (request *DescribeChatDetailRequest) {
    request = &DescribeChatDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "DescribeChatDetail")
    
    
    return
}

func NewDescribeChatDetailResponse() (response *DescribeChatDetailResponse) {
    response = &DescribeChatDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChatDetail
// 本接口（DescribeChatDetail）用于查询对话详情，通常在用户查询会话的历史记录时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeChatDetail(request *DescribeChatDetailRequest) (response *DescribeChatDetailResponse, err error) {
    return c.DescribeChatDetailWithContext(context.Background(), request)
}

// DescribeChatDetail
// 本接口（DescribeChatDetail）用于查询对话详情，通常在用户查询会话的历史记录时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeChatDetailWithContext(ctx context.Context, request *DescribeChatDetailRequest) (response *DescribeChatDetailResponse, err error) {
    if request == nil {
        request = NewDescribeChatDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "DescribeChatDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChatDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChatDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChatsRequest() (request *DescribeChatsRequest) {
    request = &DescribeChatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "DescribeChats")
    
    
    return
}

func NewDescribeChatsResponse() (response *DescribeChatsResponse) {
    response = &DescribeChatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChats
// 本接口（DescribeChats）用于查询对话列表，通常在用户查询会话列表时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeChats(request *DescribeChatsRequest) (response *DescribeChatsResponse, err error) {
    return c.DescribeChatsWithContext(context.Background(), request)
}

// DescribeChats
// 本接口（DescribeChats）用于查询对话列表，通常在用户查询会话列表时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeChatsWithContext(ctx context.Context, request *DescribeChatsRequest) (response *DescribeChatsResponse, err error) {
    if request == nil {
        request = NewDescribeChatsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "DescribeChats")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReportUrlRequest() (request *DescribeReportUrlRequest) {
    request = &DescribeReportUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "DescribeReportUrl")
    
    
    return
}

func NewDescribeReportUrlResponse() (response *DescribeReportUrlResponse) {
    response = &DescribeReportUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReportUrl
// 智能体报告地址生成并下载
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeReportUrl(request *DescribeReportUrlRequest) (response *DescribeReportUrlResponse, err error) {
    return c.DescribeReportUrlWithContext(context.Background(), request)
}

// DescribeReportUrl
// 智能体报告地址生成并下载
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeReportUrlWithContext(ctx context.Context, request *DescribeReportUrlRequest) (response *DescribeReportUrlResponse, err error) {
    if request == nil {
        request = NewDescribeReportUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "DescribeReportUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReportUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReportUrlResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateAgentInstanceRequest() (request *IsolateAgentInstanceRequest) {
    request = &IsolateAgentInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "IsolateAgentInstance")
    
    
    return
}

func NewIsolateAgentInstanceResponse() (response *IsolateAgentInstanceResponse) {
    response = &IsolateAgentInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateAgentInstance
// 本接口（IsolateAgentInstance）用于隔离智能体实例，通常在用户需要隔离智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) IsolateAgentInstance(request *IsolateAgentInstanceRequest) (response *IsolateAgentInstanceResponse, err error) {
    return c.IsolateAgentInstanceWithContext(context.Background(), request)
}

// IsolateAgentInstance
// 本接口（IsolateAgentInstance）用于隔离智能体实例，通常在用户需要隔离智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) IsolateAgentInstanceWithContext(ctx context.Context, request *IsolateAgentInstanceRequest) (response *IsolateAgentInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateAgentInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "IsolateAgentInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateAgentInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateAgentInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentInstanceParametersRequest() (request *ModifyAgentInstanceParametersRequest) {
    request = &ModifyAgentInstanceParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "ModifyAgentInstanceParameters")
    
    
    return
}

func NewModifyAgentInstanceParametersResponse() (response *ModifyAgentInstanceParametersResponse) {
    response = &ModifyAgentInstanceParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgentInstanceParameters
// 本接口（ModifyAgentInstanceParameters）用于修改智能体实例的参数列表，通常在用户需要配置智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAgentInstanceParameters(request *ModifyAgentInstanceParametersRequest) (response *ModifyAgentInstanceParametersResponse, err error) {
    return c.ModifyAgentInstanceParametersWithContext(context.Background(), request)
}

// ModifyAgentInstanceParameters
// 本接口（ModifyAgentInstanceParameters）用于修改智能体实例的参数列表，通常在用户需要配置智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAgentInstanceParametersWithContext(ctx context.Context, request *ModifyAgentInstanceParametersRequest) (response *ModifyAgentInstanceParametersResponse, err error) {
    if request == nil {
        request = NewModifyAgentInstanceParametersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "ModifyAgentInstanceParameters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgentInstanceParameters require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentInstanceParametersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyChatTitleRequest() (request *ModifyChatTitleRequest) {
    request = &ModifyChatTitleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "ModifyChatTitle")
    
    
    return
}

func NewModifyChatTitleResponse() (response *ModifyChatTitleResponse) {
    response = &ModifyChatTitleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyChatTitle
// 本接口（ModifyChatTitle）用于修改会话标题，通常在用户修改会话标题时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyChatTitle(request *ModifyChatTitleRequest) (response *ModifyChatTitleResponse, err error) {
    return c.ModifyChatTitleWithContext(context.Background(), request)
}

// ModifyChatTitle
// 本接口（ModifyChatTitle）用于修改会话标题，通常在用户修改会话标题时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyChatTitleWithContext(ctx context.Context, request *ModifyChatTitleRequest) (response *ModifyChatTitleResponse, err error) {
    if request == nil {
        request = NewModifyChatTitleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "ModifyChatTitle")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyChatTitle require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyChatTitleResponse()
    err = c.Send(request, response)
    return
}

func NewPauseAgentWorkRequest() (request *PauseAgentWorkRequest) {
    request = &PauseAgentWorkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "PauseAgentWork")
    
    
    return
}

func NewPauseAgentWorkResponse() (response *PauseAgentWorkResponse) {
    response = &PauseAgentWorkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseAgentWork
// 本接口（PauseAgentWork）用于暂停智能体实例的值守任务，通常在用户需要暂停时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PauseAgentWork(request *PauseAgentWorkRequest) (response *PauseAgentWorkResponse, err error) {
    return c.PauseAgentWorkWithContext(context.Background(), request)
}

// PauseAgentWork
// 本接口（PauseAgentWork）用于暂停智能体实例的值守任务，通常在用户需要暂停时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PauseAgentWorkWithContext(ctx context.Context, request *PauseAgentWorkRequest) (response *PauseAgentWorkResponse, err error) {
    if request == nil {
        request = NewPauseAgentWorkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "PauseAgentWork")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseAgentWork require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseAgentWorkResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverAgentInstanceRequest() (request *RecoverAgentInstanceRequest) {
    request = &RecoverAgentInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "RecoverAgentInstance")
    
    
    return
}

func NewRecoverAgentInstanceResponse() (response *RecoverAgentInstanceResponse) {
    response = &RecoverAgentInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecoverAgentInstance
// 本接口（RecoverAgentInstance）用于解隔离智能体实例，通常在用户需要解隔离智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RecoverAgentInstance(request *RecoverAgentInstanceRequest) (response *RecoverAgentInstanceResponse, err error) {
    return c.RecoverAgentInstanceWithContext(context.Background(), request)
}

// RecoverAgentInstance
// 本接口（RecoverAgentInstance）用于解隔离智能体实例，通常在用户需要解隔离智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RecoverAgentInstanceWithContext(ctx context.Context, request *RecoverAgentInstanceRequest) (response *RecoverAgentInstanceResponse, err error) {
    if request == nil {
        request = NewRecoverAgentInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "RecoverAgentInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverAgentInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverAgentInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveChatRequest() (request *RemoveChatRequest) {
    request = &RemoveChatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "RemoveChat")
    
    
    return
}

func NewRemoveChatResponse() (response *RemoveChatResponse) {
    response = &RemoveChatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveChat
// 本接口（RemoveChat）用于删除会话，通常在用户删除会话时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RemoveChat(request *RemoveChatRequest) (response *RemoveChatResponse, err error) {
    return c.RemoveChatWithContext(context.Background(), request)
}

// RemoveChat
// 本接口（RemoveChat）用于删除会话，通常在用户删除会话时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RemoveChatWithContext(ctx context.Context, request *RemoveChatRequest) (response *RemoveChatResponse, err error) {
    if request == nil {
        request = NewRemoveChatRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "RemoveChat")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveChat require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveChatResponse()
    err = c.Send(request, response)
    return
}

func NewStartAgentTaskRequest() (request *StartAgentTaskRequest) {
    request = &StartAgentTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "StartAgentTask")
    
    
    return
}

func NewStartAgentTaskResponse() (response *StartAgentTaskResponse) {
    response = &StartAgentTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartAgentTask
// 该接口用于启动一个智能体的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartAgentTask(request *StartAgentTaskRequest) (response *StartAgentTaskResponse, err error) {
    return c.StartAgentTaskWithContext(context.Background(), request)
}

// StartAgentTask
// 该接口用于启动一个智能体的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartAgentTaskWithContext(ctx context.Context, request *StartAgentTaskRequest) (response *StartAgentTaskResponse, err error) {
    if request == nil {
        request = NewStartAgentTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "StartAgentTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAgentTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAgentTaskResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateAgentInstanceRequest() (request *TerminateAgentInstanceRequest) {
    request = &TerminateAgentInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdai", APIVersion, "TerminateAgentInstance")
    
    
    return
}

func NewTerminateAgentInstanceResponse() (response *TerminateAgentInstanceResponse) {
    response = &TerminateAgentInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateAgentInstance
// 本接口（TerminateAgentInstance）用于下线智能体实例，通常在用户需要下线智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) TerminateAgentInstance(request *TerminateAgentInstanceRequest) (response *TerminateAgentInstanceResponse, err error) {
    return c.TerminateAgentInstanceWithContext(context.Background(), request)
}

// TerminateAgentInstance
// 本接口（TerminateAgentInstance）用于下线智能体实例，通常在用户需要下线智能体实例时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) TerminateAgentInstanceWithContext(ctx context.Context, request *TerminateAgentInstanceRequest) (response *TerminateAgentInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateAgentInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdai", APIVersion, "TerminateAgentInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateAgentInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateAgentInstanceResponse()
    err = c.Send(request, response)
    return
}
