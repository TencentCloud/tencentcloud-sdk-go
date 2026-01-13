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

package v20250513

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-05-13"

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


func NewAddChunkRequest() (request *AddChunkRequest) {
    request = &AddChunkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "AddChunk")
    
    
    return
}

func NewAddChunkResponse() (response *AddChunkResponse) {
    response = &AddChunkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddChunk
// 文档切片新增
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddChunk(request *AddChunkRequest) (response *AddChunkResponse, err error) {
    return c.AddChunkWithContext(context.Background(), request)
}

// AddChunk
// 文档切片新增
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddChunkWithContext(ctx context.Context, request *AddChunkRequest) (response *AddChunkResponse, err error) {
    if request == nil {
        request = NewAddChunkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "AddChunk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddChunk require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddChunkResponse()
    err = c.Send(request, response)
    return
}

func NewAddSceneRequest() (request *AddSceneRequest) {
    request = &AddSceneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "AddScene")
    
    
    return
}

func NewAddSceneResponse() (response *AddSceneResponse) {
    response = &AddSceneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddScene
// 新增场景
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddScene(request *AddSceneRequest) (response *AddSceneResponse, err error) {
    return c.AddSceneWithContext(context.Background(), request)
}

// AddScene
// 新增场景
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddSceneWithContext(ctx context.Context, request *AddSceneRequest) (response *AddSceneResponse, err error) {
    if request == nil {
        request = NewAddSceneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "AddScene")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddScene require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddSceneResponse()
    err = c.Send(request, response)
    return
}

func NewChatAIRequest() (request *ChatAIRequest) {
    request = &ChatAIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "ChatAI")
    
    
    return
}

func NewChatAIResponse() (response *ChatAIResponse) {
    response = &ChatAIResponse{} 
    return

}

// ChatAI
// 提供DataAgent 产品服务API
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChatAI(request *ChatAIRequest) (response *ChatAIResponse, err error) {
    return c.ChatAIWithContext(context.Background(), request)
}

// ChatAI
// 提供DataAgent 产品服务API
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChatAIWithContext(ctx context.Context, request *ChatAIRequest) (response *ChatAIResponse, err error) {
    if request == nil {
        request = NewChatAIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "ChatAI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatAI require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatAIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataAgentSessionRequest() (request *CreateDataAgentSessionRequest) {
    request = &CreateDataAgentSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "CreateDataAgentSession")
    
    
    return
}

func NewCreateDataAgentSessionResponse() (response *CreateDataAgentSessionResponse) {
    response = &CreateDataAgentSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataAgentSession
// 生成DataAgent 会话ID
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDataAgentSession(request *CreateDataAgentSessionRequest) (response *CreateDataAgentSessionResponse, err error) {
    return c.CreateDataAgentSessionWithContext(context.Background(), request)
}

// CreateDataAgentSession
// 生成DataAgent 会话ID
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDataAgentSessionWithContext(ctx context.Context, request *CreateDataAgentSessionRequest) (response *CreateDataAgentSessionResponse, err error) {
    if request == nil {
        request = NewCreateDataAgentSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "CreateDataAgentSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataAgentSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataAgentSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteChunkRequest() (request *DeleteChunkRequest) {
    request = &DeleteChunkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "DeleteChunk")
    
    
    return
}

func NewDeleteChunkResponse() (response *DeleteChunkResponse) {
    response = &DeleteChunkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteChunk
// 文档切片删除
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteChunk(request *DeleteChunkRequest) (response *DeleteChunkResponse, err error) {
    return c.DeleteChunkWithContext(context.Background(), request)
}

// DeleteChunk
// 文档切片删除
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteChunkWithContext(ctx context.Context, request *DeleteChunkRequest) (response *DeleteChunkResponse, err error) {
    if request == nil {
        request = NewDeleteChunkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "DeleteChunk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteChunk require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteChunkResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDataAgentSessionRequest() (request *DeleteDataAgentSessionRequest) {
    request = &DeleteDataAgentSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "DeleteDataAgentSession")
    
    
    return
}

func NewDeleteDataAgentSessionResponse() (response *DeleteDataAgentSessionResponse) {
    response = &DeleteDataAgentSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDataAgentSession
// 删除会话
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDataAgentSession(request *DeleteDataAgentSessionRequest) (response *DeleteDataAgentSessionResponse, err error) {
    return c.DeleteDataAgentSessionWithContext(context.Background(), request)
}

// DeleteDataAgentSession
// 删除会话
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDataAgentSessionWithContext(ctx context.Context, request *DeleteDataAgentSessionRequest) (response *DeleteDataAgentSessionResponse, err error) {
    if request == nil {
        request = NewDeleteDataAgentSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "DeleteDataAgentSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataAgentSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataAgentSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSceneRequest() (request *DeleteSceneRequest) {
    request = &DeleteSceneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "DeleteScene")
    
    
    return
}

func NewDeleteSceneResponse() (response *DeleteSceneResponse) {
    response = &DeleteSceneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteScene
// 删除场景
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteScene(request *DeleteSceneRequest) (response *DeleteSceneResponse, err error) {
    return c.DeleteSceneWithContext(context.Background(), request)
}

// DeleteScene
// 删除场景
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSceneWithContext(ctx context.Context, request *DeleteSceneRequest) (response *DeleteSceneResponse, err error) {
    if request == nil {
        request = NewDeleteSceneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "DeleteScene")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScene require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSceneResponse()
    err = c.Send(request, response)
    return
}

func NewGetJobsByKnowledgeBaseIdRequest() (request *GetJobsByKnowledgeBaseIdRequest) {
    request = &GetJobsByKnowledgeBaseIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "GetJobsByKnowledgeBaseId")
    
    
    return
}

func NewGetJobsByKnowledgeBaseIdResponse() (response *GetJobsByKnowledgeBaseIdResponse) {
    response = &GetJobsByKnowledgeBaseIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetJobsByKnowledgeBaseId
// 根据知识库id查询jobs 列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetJobsByKnowledgeBaseId(request *GetJobsByKnowledgeBaseIdRequest) (response *GetJobsByKnowledgeBaseIdResponse, err error) {
    return c.GetJobsByKnowledgeBaseIdWithContext(context.Background(), request)
}

// GetJobsByKnowledgeBaseId
// 根据知识库id查询jobs 列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetJobsByKnowledgeBaseIdWithContext(ctx context.Context, request *GetJobsByKnowledgeBaseIdRequest) (response *GetJobsByKnowledgeBaseIdResponse, err error) {
    if request == nil {
        request = NewGetJobsByKnowledgeBaseIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "GetJobsByKnowledgeBaseId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetJobsByKnowledgeBaseId require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetJobsByKnowledgeBaseIdResponse()
    err = c.Send(request, response)
    return
}

func NewGetKnowledgeBaseFileListRequest() (request *GetKnowledgeBaseFileListRequest) {
    request = &GetKnowledgeBaseFileListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "GetKnowledgeBaseFileList")
    
    
    return
}

func NewGetKnowledgeBaseFileListResponse() (response *GetKnowledgeBaseFileListResponse) {
    response = &GetKnowledgeBaseFileListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetKnowledgeBaseFileList
// 获取知识库文件信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetKnowledgeBaseFileList(request *GetKnowledgeBaseFileListRequest) (response *GetKnowledgeBaseFileListResponse, err error) {
    return c.GetKnowledgeBaseFileListWithContext(context.Background(), request)
}

// GetKnowledgeBaseFileList
// 获取知识库文件信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetKnowledgeBaseFileListWithContext(ctx context.Context, request *GetKnowledgeBaseFileListRequest) (response *GetKnowledgeBaseFileListResponse, err error) {
    if request == nil {
        request = NewGetKnowledgeBaseFileListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "GetKnowledgeBaseFileList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetKnowledgeBaseFileList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetKnowledgeBaseFileListResponse()
    err = c.Send(request, response)
    return
}

func NewGetKnowledgeBaseListRequest() (request *GetKnowledgeBaseListRequest) {
    request = &GetKnowledgeBaseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "GetKnowledgeBaseList")
    
    
    return
}

func NewGetKnowledgeBaseListResponse() (response *GetKnowledgeBaseListResponse) {
    response = &GetKnowledgeBaseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetKnowledgeBaseList
// 获取知识库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetKnowledgeBaseList(request *GetKnowledgeBaseListRequest) (response *GetKnowledgeBaseListResponse, err error) {
    return c.GetKnowledgeBaseListWithContext(context.Background(), request)
}

// GetKnowledgeBaseList
// 获取知识库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetKnowledgeBaseListWithContext(ctx context.Context, request *GetKnowledgeBaseListRequest) (response *GetKnowledgeBaseListResponse, err error) {
    if request == nil {
        request = NewGetKnowledgeBaseListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "GetKnowledgeBaseList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetKnowledgeBaseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetKnowledgeBaseListResponse()
    err = c.Send(request, response)
    return
}

func NewGetSessionDetailsRequest() (request *GetSessionDetailsRequest) {
    request = &GetSessionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "GetSessionDetails")
    
    
    return
}

func NewGetSessionDetailsResponse() (response *GetSessionDetailsResponse) {
    response = &GetSessionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSessionDetails
// 获取用户会话记录详情列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSessionDetails(request *GetSessionDetailsRequest) (response *GetSessionDetailsResponse, err error) {
    return c.GetSessionDetailsWithContext(context.Background(), request)
}

// GetSessionDetails
// 获取用户会话记录详情列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSessionDetailsWithContext(ctx context.Context, request *GetSessionDetailsRequest) (response *GetSessionDetailsResponse, err error) {
    if request == nil {
        request = NewGetSessionDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "GetSessionDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSessionDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSessionDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewGetUploadJobDetailsRequest() (request *GetUploadJobDetailsRequest) {
    request = &GetUploadJobDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "GetUploadJobDetails")
    
    
    return
}

func NewGetUploadJobDetailsResponse() (response *GetUploadJobDetailsResponse) {
    response = &GetUploadJobDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetUploadJobDetails
// 查询上传任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetUploadJobDetails(request *GetUploadJobDetailsRequest) (response *GetUploadJobDetailsResponse, err error) {
    return c.GetUploadJobDetailsWithContext(context.Background(), request)
}

// GetUploadJobDetails
// 查询上传任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetUploadJobDetailsWithContext(ctx context.Context, request *GetUploadJobDetailsRequest) (response *GetUploadJobDetailsResponse, err error) {
    if request == nil {
        request = NewGetUploadJobDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "GetUploadJobDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUploadJobDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUploadJobDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyChunkRequest() (request *ModifyChunkRequest) {
    request = &ModifyChunkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "ModifyChunk")
    
    
    return
}

func NewModifyChunkResponse() (response *ModifyChunkResponse) {
    response = &ModifyChunkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyChunk
// 编辑修改分片
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyChunk(request *ModifyChunkRequest) (response *ModifyChunkResponse, err error) {
    return c.ModifyChunkWithContext(context.Background(), request)
}

// ModifyChunk
// 编辑修改分片
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyChunkWithContext(ctx context.Context, request *ModifyChunkRequest) (response *ModifyChunkResponse, err error) {
    if request == nil {
        request = NewModifyChunkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "ModifyChunk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyChunk require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyChunkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKnowledgeBaseRequest() (request *ModifyKnowledgeBaseRequest) {
    request = &ModifyKnowledgeBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "ModifyKnowledgeBase")
    
    
    return
}

func NewModifyKnowledgeBaseResponse() (response *ModifyKnowledgeBaseResponse) {
    response = &ModifyKnowledgeBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyKnowledgeBase
// 操作知识库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyKnowledgeBase(request *ModifyKnowledgeBaseRequest) (response *ModifyKnowledgeBaseResponse, err error) {
    return c.ModifyKnowledgeBaseWithContext(context.Background(), request)
}

// ModifyKnowledgeBase
// 操作知识库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyKnowledgeBaseWithContext(ctx context.Context, request *ModifyKnowledgeBaseRequest) (response *ModifyKnowledgeBaseResponse, err error) {
    if request == nil {
        request = NewModifyKnowledgeBaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "ModifyKnowledgeBase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyKnowledgeBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyKnowledgeBaseResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserAuthorityRequest() (request *ModifyUserAuthorityRequest) {
    request = &ModifyUserAuthorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "ModifyUserAuthority")
    
    
    return
}

func NewModifyUserAuthorityResponse() (response *ModifyUserAuthorityResponse) {
    response = &ModifyUserAuthorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserAuthority
// 修改对象权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserAuthority(request *ModifyUserAuthorityRequest) (response *ModifyUserAuthorityResponse, err error) {
    return c.ModifyUserAuthorityWithContext(context.Background(), request)
}

// ModifyUserAuthority
// 修改对象权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserAuthorityWithContext(ctx context.Context, request *ModifyUserAuthorityRequest) (response *ModifyUserAuthorityResponse, err error) {
    if request == nil {
        request = NewModifyUserAuthorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "ModifyUserAuthority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserAuthority require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserAuthorityResponse()
    err = c.Send(request, response)
    return
}

func NewQueryChunkListRequest() (request *QueryChunkListRequest) {
    request = &QueryChunkListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "QueryChunkList")
    
    
    return
}

func NewQueryChunkListResponse() (response *QueryChunkListResponse) {
    response = &QueryChunkListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryChunkList
// 文档切片查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryChunkList(request *QueryChunkListRequest) (response *QueryChunkListResponse, err error) {
    return c.QueryChunkListWithContext(context.Background(), request)
}

// QueryChunkList
// 文档切片查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryChunkListWithContext(ctx context.Context, request *QueryChunkListRequest) (response *QueryChunkListResponse, err error) {
    if request == nil {
        request = NewQueryChunkListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "QueryChunkList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryChunkList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryChunkListResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySceneListRequest() (request *QuerySceneListRequest) {
    request = &QuerySceneListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "QuerySceneList")
    
    
    return
}

func NewQuerySceneListResponse() (response *QuerySceneListResponse) {
    response = &QuerySceneListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QuerySceneList
// 查询场景列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QuerySceneList(request *QuerySceneListRequest) (response *QuerySceneListResponse, err error) {
    return c.QuerySceneListWithContext(context.Background(), request)
}

// QuerySceneList
// 查询场景列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QuerySceneListWithContext(ctx context.Context, request *QuerySceneListRequest) (response *QuerySceneListResponse, err error) {
    if request == nil {
        request = NewQuerySceneListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "QuerySceneList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QuerySceneList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQuerySceneListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryUserAuthorityRequest() (request *QueryUserAuthorityRequest) {
    request = &QueryUserAuthorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "QueryUserAuthority")
    
    
    return
}

func NewQueryUserAuthorityResponse() (response *QueryUserAuthorityResponse) {
    response = &QueryUserAuthorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryUserAuthority
// 查询对象权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryUserAuthority(request *QueryUserAuthorityRequest) (response *QueryUserAuthorityResponse, err error) {
    return c.QueryUserAuthorityWithContext(context.Background(), request)
}

// QueryUserAuthority
// 查询对象权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryUserAuthorityWithContext(ctx context.Context, request *QueryUserAuthorityRequest) (response *QueryUserAuthorityResponse, err error) {
    if request == nil {
        request = NewQueryUserAuthorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "QueryUserAuthority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryUserAuthority require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryUserAuthorityResponse()
    err = c.Send(request, response)
    return
}

func NewStopChatAIRequest() (request *StopChatAIRequest) {
    request = &StopChatAIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "StopChatAI")
    
    
    return
}

func NewStopChatAIResponse() (response *StopChatAIResponse) {
    response = &StopChatAIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopChatAI
// 中断DataAgent的回答输出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopChatAI(request *StopChatAIRequest) (response *StopChatAIResponse, err error) {
    return c.StopChatAIWithContext(context.Background(), request)
}

// StopChatAI
// 中断DataAgent的回答输出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopChatAIWithContext(ctx context.Context, request *StopChatAIRequest) (response *StopChatAIResponse, err error) {
    if request == nil {
        request = NewStopChatAIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "StopChatAI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopChatAI require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopChatAIResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSceneRequest() (request *UpdateSceneRequest) {
    request = &UpdateSceneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "UpdateScene")
    
    
    return
}

func NewUpdateSceneResponse() (response *UpdateSceneResponse) {
    response = &UpdateSceneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateScene
// 更新场景
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateScene(request *UpdateSceneRequest) (response *UpdateSceneResponse, err error) {
    return c.UpdateSceneWithContext(context.Background(), request)
}

// UpdateScene
// 更新场景
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateSceneWithContext(ctx context.Context, request *UpdateSceneRequest) (response *UpdateSceneResponse, err error) {
    if request == nil {
        request = NewUpdateSceneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "UpdateScene")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateScene require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSceneResponse()
    err = c.Send(request, response)
    return
}

func NewUploadAndCommitFileRequest() (request *UploadAndCommitFileRequest) {
    request = &UploadAndCommitFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataagent", APIVersion, "UploadAndCommitFile")
    
    
    return
}

func NewUploadAndCommitFileResponse() (response *UploadAndCommitFileResponse) {
    response = &UploadAndCommitFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadAndCommitFile
// 上传提交文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadAndCommitFile(request *UploadAndCommitFileRequest) (response *UploadAndCommitFileResponse, err error) {
    return c.UploadAndCommitFileWithContext(context.Background(), request)
}

// UploadAndCommitFile
// 上传提交文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadAndCommitFileWithContext(ctx context.Context, request *UploadAndCommitFileRequest) (response *UploadAndCommitFileResponse, err error) {
    if request == nil {
        request = NewUploadAndCommitFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dataagent", APIVersion, "UploadAndCommitFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadAndCommitFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadAndCommitFileResponse()
    err = c.Send(request, response)
    return
}
