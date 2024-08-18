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

package v20230901

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-09-01"

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


func NewActivateServiceRequest() (request *ActivateServiceRequest) {
    request = &ActivateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "ActivateService")
    
    
    return
}

func NewActivateServiceResponse() (response *ActivateServiceResponse) {
    response = &ActivateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ActivateService
// 开通服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_PARTNERACCOUNTUNSUPPORT = "FailedOperation.PartnerAccountUnSupport"
//  FAILEDOPERATION_USERUNAUTHERROR = "FailedOperation.UserUnAuthError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
func (c *Client) ActivateService(request *ActivateServiceRequest) (response *ActivateServiceResponse, err error) {
    return c.ActivateServiceWithContext(context.Background(), request)
}

// ActivateService
// 开通服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_PARTNERACCOUNTUNSUPPORT = "FailedOperation.PartnerAccountUnSupport"
//  FAILEDOPERATION_USERUNAUTHERROR = "FailedOperation.UserUnAuthError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
func (c *Client) ActivateServiceWithContext(ctx context.Context, request *ActivateServiceRequest) (response *ActivateServiceResponse, err error) {
    if request == nil {
        request = NewActivateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewChatCompletionsRequest() (request *ChatCompletionsRequest) {
    request = &ChatCompletionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "ChatCompletions")
    
    
    return
}

func NewChatCompletionsResponse() (response *ChatCompletionsResponse) {
    response = &ChatCompletionsResponse{} 
    return

}

// ChatCompletions
// 腾讯混元大模型是由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力。本接口支持流式或非流式调用，当使用流式调用时为 SSE 协议。
//
// 
//
//  1. 本接口暂不支持返回图片内容。
//
//  2. 默认该接口下单账号限制并发数为  5 路，如您有提高并发限制的需求请 [购买](https://buy.cloud.tencent.com/hunyuan) 。
//
//  3. 请使用 SDK 调用本接口，每种开发语言的 SDK Git 仓库 examples/hunyuan/v20230901/ 目录下有提供示例供参考。SDK 链接在文档下方 “**开发者资源 - SDK**” 部分提供。
//
//  4. 我们推荐您使用 API Explorer，方便快速地在线调试接口和下载各语言的示例代码，[点击打开](https://console.cloud.tencent.com/api/explorer?Product=hunyuan&Version=2023-09-01&Action=ChatCompletions)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//  FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"
//  FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"
//  FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"
//  FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"
//  FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ChatCompletions(request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
    return c.ChatCompletionsWithContext(context.Background(), request)
}

// ChatCompletions
// 腾讯混元大模型是由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力。本接口支持流式或非流式调用，当使用流式调用时为 SSE 协议。
//
// 
//
//  1. 本接口暂不支持返回图片内容。
//
//  2. 默认该接口下单账号限制并发数为  5 路，如您有提高并发限制的需求请 [购买](https://buy.cloud.tencent.com/hunyuan) 。
//
//  3. 请使用 SDK 调用本接口，每种开发语言的 SDK Git 仓库 examples/hunyuan/v20230901/ 目录下有提供示例供参考。SDK 链接在文档下方 “**开发者资源 - SDK**” 部分提供。
//
//  4. 我们推荐您使用 API Explorer，方便快速地在线调试接口和下载各语言的示例代码，[点击打开](https://console.cloud.tencent.com/api/explorer?Product=hunyuan&Version=2023-09-01&Action=ChatCompletions)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//  FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"
//  FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"
//  FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"
//  FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"
//  FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ChatCompletionsWithContext(ctx context.Context, request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
    if request == nil {
        request = NewChatCompletionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatCompletions require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatCompletionsResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmbeddingRequest() (request *GetEmbeddingRequest) {
    request = &GetEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "GetEmbedding")
    
    
    return
}

func NewGetEmbeddingResponse() (response *GetEmbeddingResponse) {
    response = &GetEmbeddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetEmbedding
// 腾讯混元 Embedding 接口，可以将文本转化为高质量的向量数据。向量维度为1024维。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetEmbedding(request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
    return c.GetEmbeddingWithContext(context.Background(), request)
}

// GetEmbedding
// 腾讯混元 Embedding 接口，可以将文本转化为高质量的向量数据。向量维度为1024维。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetEmbeddingWithContext(ctx context.Context, request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
    if request == nil {
        request = NewGetEmbeddingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetEmbedding require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetEmbeddingResponse()
    err = c.Send(request, response)
    return
}

func NewGetTokenCountRequest() (request *GetTokenCountRequest) {
    request = &GetTokenCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "GetTokenCount")
    
    
    return
}

func NewGetTokenCountResponse() (response *GetTokenCountResponse) {
    response = &GetTokenCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTokenCount
// 该接口用于计算文本对应Token数、字符数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetTokenCount(request *GetTokenCountRequest) (response *GetTokenCountResponse, err error) {
    return c.GetTokenCountWithContext(context.Background(), request)
}

// GetTokenCount
// 该接口用于计算文本对应Token数、字符数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetTokenCountWithContext(ctx context.Context, request *GetTokenCountRequest) (response *GetTokenCountResponse, err error) {
    if request == nil {
        request = NewGetTokenCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTokenCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTokenCountResponse()
    err = c.Send(request, response)
    return
}

func NewQueryHunyuanImageJobRequest() (request *QueryHunyuanImageJobRequest) {
    request = &QueryHunyuanImageJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "QueryHunyuanImageJob")
    
    
    return
}

func NewQueryHunyuanImageJobResponse() (response *QueryHunyuanImageJobResponse) {
    response = &QueryHunyuanImageJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuanImageJob
// 混元生图接口基于混元大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个混元生图异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。混元生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryHunyuanImageJob(request *QueryHunyuanImageJobRequest) (response *QueryHunyuanImageJobResponse, err error) {
    return c.QueryHunyuanImageJobWithContext(context.Background(), request)
}

// QueryHunyuanImageJob
// 混元生图接口基于混元大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个混元生图异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。混元生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryHunyuanImageJobWithContext(ctx context.Context, request *QueryHunyuanImageJobRequest) (response *QueryHunyuanImageJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuanImageJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuanImageJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuanImageJobResponse()
    err = c.Send(request, response)
    return
}

func NewSetPayModeRequest() (request *SetPayModeRequest) {
    request = &SetPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "SetPayMode")
    
    
    return
}

func NewSetPayModeResponse() (response *SetPayModeResponse) {
    response = &SetPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetPayMode
// 设置付费模式
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_SETPAYMODEEXCEED = "FailedOperation.SetPayModeExceed"
func (c *Client) SetPayMode(request *SetPayModeRequest) (response *SetPayModeResponse, err error) {
    return c.SetPayModeWithContext(context.Background(), request)
}

// SetPayMode
// 设置付费模式
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_SETPAYMODEEXCEED = "FailedOperation.SetPayModeExceed"
func (c *Client) SetPayModeWithContext(ctx context.Context, request *SetPayModeRequest) (response *SetPayModeResponse, err error) {
    if request == nil {
        request = NewSetPayModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHunyuanImageJobRequest() (request *SubmitHunyuanImageJobRequest) {
    request = &SubmitHunyuanImageJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "SubmitHunyuanImageJob")
    
    
    return
}

func NewSubmitHunyuanImageJobResponse() (response *SubmitHunyuanImageJobResponse) {
    response = &SubmitHunyuanImageJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanImageJob
// 混元生图接口基于混元大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个混元生图异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。混元生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) SubmitHunyuanImageJob(request *SubmitHunyuanImageJobRequest) (response *SubmitHunyuanImageJobResponse, err error) {
    return c.SubmitHunyuanImageJobWithContext(context.Background(), request)
}

// SubmitHunyuanImageJob
// 混元生图接口基于混元大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个混元生图异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。混元生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) SubmitHunyuanImageJobWithContext(ctx context.Context, request *SubmitHunyuanImageJobRequest) (response *SubmitHunyuanImageJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanImageJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanImageJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanImageJobResponse()
    err = c.Send(request, response)
    return
}

func NewTextToImageLiteRequest() (request *TextToImageLiteRequest) {
    request = &TextToImageLiteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "TextToImageLite")
    
    
    return
}

func NewTextToImageLiteResponse() (response *TextToImageLiteResponse) {
    response = &TextToImageLiteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextToImageLite
// 文生图轻量版接口根据输入的文本描述，智能生成与之相关的结果图。
//
// 文生图轻量版默认提供3个并发任务数，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) TextToImageLite(request *TextToImageLiteRequest) (response *TextToImageLiteResponse, err error) {
    return c.TextToImageLiteWithContext(context.Background(), request)
}

// TextToImageLite
// 文生图轻量版接口根据输入的文本描述，智能生成与之相关的结果图。
//
// 文生图轻量版默认提供3个并发任务数，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) TextToImageLiteWithContext(ctx context.Context, request *TextToImageLiteRequest) (response *TextToImageLiteResponse, err error) {
    if request == nil {
        request = NewTextToImageLiteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToImageLite require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToImageLiteResponse()
    err = c.Send(request, response)
    return
}
