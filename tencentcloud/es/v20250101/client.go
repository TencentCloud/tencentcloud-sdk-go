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

package v20250101

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-01-01"

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


func NewChatCompletionsRequest() (request *ChatCompletionsRequest) {
    request = &ChatCompletionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ChatCompletions")
    
    
    return
}

func NewChatCompletionsResponse() (response *ChatCompletionsResponse) {
    response = &ChatCompletionsResponse{} 
    return

}

// ChatCompletions
// 本服务支持一系列高性能的大语言模型，包括DeepSeek以及腾讯自主研发的混元大模型，结合混合搜索等先进搜索技术，快速高效实现RAG，有效解决幻觉和知识更新问题。
//
// 本接口有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ChatCompletions(request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
    return c.ChatCompletionsWithContext(context.Background(), request)
}

// ChatCompletions
// 本服务支持一系列高性能的大语言模型，包括DeepSeek以及腾讯自主研发的混元大模型，结合混合搜索等先进搜索技术，快速高效实现RAG，有效解决幻觉和知识更新问题。
//
// 本接口有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ChatCompletionsWithContext(ctx context.Context, request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
    if request == nil {
        request = NewChatCompletionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatCompletions require credential")
    }

    request.SetContext(ctx)
    request.SetRootDomain("ai." + tchttp.RootDomain)
    
    response = NewChatCompletionsResponse()
    err = c.Send(request, response)
    return
}

func NewChunkDocumentRequest() (request *ChunkDocumentRequest) {
    request = &ChunkDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ChunkDocument")
    
    
    return
}

func NewChunkDocumentResponse() (response *ChunkDocumentResponse) {
    response = &ChunkDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChunkDocument
// 文本切片是将长文本分割为短片段的技术，用于适配模型输入、提升处理效率或信息检索，平衡片段长度与语义连贯性，适用于NLP、数据分析等场景。
//
// 本接口为分隔符规则切片接口，有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ChunkDocument(request *ChunkDocumentRequest) (response *ChunkDocumentResponse, err error) {
    return c.ChunkDocumentWithContext(context.Background(), request)
}

// ChunkDocument
// 文本切片是将长文本分割为短片段的技术，用于适配模型输入、提升处理效率或信息检索，平衡片段长度与语义连贯性，适用于NLP、数据分析等场景。
//
// 本接口为分隔符规则切片接口，有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ChunkDocumentWithContext(ctx context.Context, request *ChunkDocumentRequest) (response *ChunkDocumentResponse, err error) {
    if request == nil {
        request = NewChunkDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChunkDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewChunkDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewChunkDocumentAsyncRequest() (request *ChunkDocumentAsyncRequest) {
    request = &ChunkDocumentAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ChunkDocumentAsync")
    
    
    return
}

func NewChunkDocumentAsyncResponse() (response *ChunkDocumentAsyncResponse) {
    response = &ChunkDocumentAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChunkDocumentAsync
// 文本切片是将长文本分割为短片段的技术，用于适配模型输入、提升处理效率或信息检索，平衡片段长度与语义连贯性，适用于NLP、数据分析等场景。
//
// 本接口为异步接口，有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ChunkDocumentAsync(request *ChunkDocumentAsyncRequest) (response *ChunkDocumentAsyncResponse, err error) {
    return c.ChunkDocumentAsyncWithContext(context.Background(), request)
}

// ChunkDocumentAsync
// 文本切片是将长文本分割为短片段的技术，用于适配模型输入、提升处理效率或信息检索，平衡片段长度与语义连贯性，适用于NLP、数据分析等场景。
//
// 本接口为异步接口，有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ChunkDocumentAsyncWithContext(ctx context.Context, request *ChunkDocumentAsyncRequest) (response *ChunkDocumentAsyncResponse, err error) {
    if request == nil {
        request = NewChunkDocumentAsyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChunkDocumentAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewChunkDocumentAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewGetDocumentChunkResultRequest() (request *GetDocumentChunkResultRequest) {
    request = &GetDocumentChunkResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetDocumentChunkResult")
    
    
    return
}

func NewGetDocumentChunkResultResponse() (response *GetDocumentChunkResultResponse) {
    response = &GetDocumentChunkResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDocumentChunkResult
// 获取文档切片结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetDocumentChunkResult(request *GetDocumentChunkResultRequest) (response *GetDocumentChunkResultResponse, err error) {
    return c.GetDocumentChunkResultWithContext(context.Background(), request)
}

// GetDocumentChunkResult
// 获取文档切片结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetDocumentChunkResultWithContext(ctx context.Context, request *GetDocumentChunkResultRequest) (response *GetDocumentChunkResultResponse, err error) {
    if request == nil {
        request = NewGetDocumentChunkResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDocumentChunkResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDocumentChunkResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetDocumentParseResultRequest() (request *GetDocumentParseResultRequest) {
    request = &GetDocumentParseResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetDocumentParseResult")
    
    
    return
}

func NewGetDocumentParseResultResponse() (response *GetDocumentParseResultResponse) {
    response = &GetDocumentParseResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDocumentParseResult
// 本接口用于获取文档解析异步处理结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetDocumentParseResult(request *GetDocumentParseResultRequest) (response *GetDocumentParseResultResponse, err error) {
    return c.GetDocumentParseResultWithContext(context.Background(), request)
}

// GetDocumentParseResult
// 本接口用于获取文档解析异步处理结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetDocumentParseResultWithContext(ctx context.Context, request *GetDocumentParseResultRequest) (response *GetDocumentParseResultResponse, err error) {
    if request == nil {
        request = NewGetDocumentParseResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDocumentParseResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDocumentParseResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetTextEmbeddingRequest() (request *GetTextEmbeddingRequest) {
    request = &GetTextEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetTextEmbedding")
    
    
    return
}

func NewGetTextEmbeddingResponse() (response *GetTextEmbeddingResponse) {
    response = &GetTextEmbeddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTextEmbedding
// Embedding是一种将高维数据映射到低维空间的技术，通常用于将非结构化数据，如文本、图像或音频转化为向量表示，使其更容易输入机器模型进行处理，并且向量之间的距离可以反映对象之间的相似性。
//
// 本接口有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTextEmbedding(request *GetTextEmbeddingRequest) (response *GetTextEmbeddingResponse, err error) {
    return c.GetTextEmbeddingWithContext(context.Background(), request)
}

// GetTextEmbedding
// Embedding是一种将高维数据映射到低维空间的技术，通常用于将非结构化数据，如文本、图像或音频转化为向量表示，使其更容易输入机器模型进行处理，并且向量之间的距离可以反映对象之间的相似性。
//
// 本接口有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTextEmbeddingWithContext(ctx context.Context, request *GetTextEmbeddingRequest) (response *GetTextEmbeddingResponse, err error) {
    if request == nil {
        request = NewGetTextEmbeddingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTextEmbedding require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTextEmbeddingResponse()
    err = c.Send(request, response)
    return
}

func NewParseDocumentRequest() (request *ParseDocumentRequest) {
    request = &ParseDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ParseDocument")
    
    
    return
}

func NewParseDocumentResponse() (response *ParseDocumentResponse) {
    response = &ParseDocumentResponse{} 
    return

}

// ParseDocument
// 本服务可将各类格式文档精准转换为标准格式，满足企业知识库建设、技术文档迁移、内容平台结构化存储等需求。
//
// 本接口有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ParseDocument(request *ParseDocumentRequest) (response *ParseDocumentResponse, err error) {
    return c.ParseDocumentWithContext(context.Background(), request)
}

// ParseDocument
// 本服务可将各类格式文档精准转换为标准格式，满足企业知识库建设、技术文档迁移、内容平台结构化存储等需求。
//
// 本接口有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ParseDocumentWithContext(ctx context.Context, request *ParseDocumentRequest) (response *ParseDocumentResponse, err error) {
    if request == nil {
        request = NewParseDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseDocument require credential")
    }

    request.SetContext(ctx)
    request.SetRootDomain("ai." + tchttp.RootDomain)
    
    response = NewParseDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewParseDocumentAsyncRequest() (request *ParseDocumentAsyncRequest) {
    request = &ParseDocumentAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ParseDocumentAsync")
    
    
    return
}

func NewParseDocumentAsyncResponse() (response *ParseDocumentAsyncResponse) {
    response = &ParseDocumentAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ParseDocumentAsync
// 本服务可将各类格式文档精准转换为标准格式，满足企业知识库建设、技术文档迁移、内容平台结构化存储等需求。
//
// 本接口为异步接口，有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ParseDocumentAsync(request *ParseDocumentAsyncRequest) (response *ParseDocumentAsyncResponse, err error) {
    return c.ParseDocumentAsyncWithContext(context.Background(), request)
}

// ParseDocumentAsync
// 本服务可将各类格式文档精准转换为标准格式，满足企业知识库建设、技术文档迁移、内容平台结构化存储等需求。
//
// 本接口为异步接口，有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ParseDocumentAsyncWithContext(ctx context.Context, request *ParseDocumentAsyncRequest) (response *ParseDocumentAsyncResponse, err error) {
    if request == nil {
        request = NewParseDocumentAsyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseDocumentAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseDocumentAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewRunRerankRequest() (request *RunRerankRequest) {
    request = &RunRerankRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RunRerank")
    
    
    return
}

func NewRunRerankResponse() (response *RunRerankResponse) {
    response = &RunRerankResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunRerank
// 重排是指在 RAG 过程中，通过评估文档与查询之间的相关性，将最相关的文档放在前面，确保语言模型在生成回答时优先考虑排名靠前的上下文，提高生成结果的准确性和可信度，也可以通过这种方式进行过滤，减少大模型成本。
//
// 本接口有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RunRerank(request *RunRerankRequest) (response *RunRerankResponse, err error) {
    return c.RunRerankWithContext(context.Background(), request)
}

// RunRerank
// 重排是指在 RAG 过程中，通过评估文档与查询之间的相关性，将最相关的文档放在前面，确保语言模型在生成回答时优先考虑排名靠前的上下文，提高生成结果的准确性和可信度，也可以通过这种方式进行过滤，减少大模型成本。
//
// 本接口有单账号调用上限控制，如您有提高并发限制的需求请[联系我们](https://cloud.tencent.com/act/event/Online_service)  。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RunRerankWithContext(ctx context.Context, request *RunRerankRequest) (response *RunRerankResponse, err error) {
    if request == nil {
        request = NewRunRerankRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunRerank require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunRerankResponse()
    err = c.Send(request, response)
    return
}
