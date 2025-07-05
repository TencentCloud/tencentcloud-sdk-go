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
// 如需使用OpenAI兼容接口， 请参考文档：[OpenAI 兼容接口](https://cloud.tencent.com/document/product/1729/111007)
//
// 
//
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) ChatCompletions(request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
    return c.ChatCompletionsWithContext(context.Background(), request)
}

// ChatCompletions
// 如需使用OpenAI兼容接口， 请参考文档：[OpenAI 兼容接口](https://cloud.tencent.com/document/product/1729/111007)
//
// 
//
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
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

func NewChatTranslationsRequest() (request *ChatTranslationsRequest) {
    request = &ChatTranslationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "ChatTranslations")
    
    
    return
}

func NewChatTranslationsResponse() (response *ChatTranslationsResponse) {
    response = &ChatTranslationsResponse{} 
    return

}

// ChatTranslations
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
func (c *Client) ChatTranslations(request *ChatTranslationsRequest) (response *ChatTranslationsResponse, err error) {
    return c.ChatTranslationsWithContext(context.Background(), request)
}

// ChatTranslations
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
func (c *Client) ChatTranslationsWithContext(ctx context.Context, request *ChatTranslationsRequest) (response *ChatTranslationsResponse, err error) {
    if request == nil {
        request = NewChatTranslationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatTranslations require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatTranslationsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateThreadRequest() (request *CreateThreadRequest) {
    request = &CreateThreadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "CreateThread")
    
    
    return
}

func NewCreateThreadResponse() (response *CreateThreadResponse) {
    response = &CreateThreadResponse{} 
    return

}

// CreateThread
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) CreateThread(request *CreateThreadRequest) (response *CreateThreadResponse, err error) {
    return c.CreateThreadWithContext(context.Background(), request)
}

// CreateThread
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) CreateThreadWithContext(ctx context.Context, request *CreateThreadRequest) (response *CreateThreadResponse, err error) {
    if request == nil {
        request = NewCreateThreadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateThread require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateThreadResponse()
    err = c.Send(request, response)
    return
}

func NewFilesDeletionsRequest() (request *FilesDeletionsRequest) {
    request = &FilesDeletionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "FilesDeletions")
    
    
    return
}

func NewFilesDeletionsResponse() (response *FilesDeletionsResponse) {
    response = &FilesDeletionsResponse{} 
    return

}

// FilesDeletions
// 删除文件。
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) FilesDeletions(request *FilesDeletionsRequest) (response *FilesDeletionsResponse, err error) {
    return c.FilesDeletionsWithContext(context.Background(), request)
}

// FilesDeletions
// 删除文件。
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) FilesDeletionsWithContext(ctx context.Context, request *FilesDeletionsRequest) (response *FilesDeletionsResponse, err error) {
    if request == nil {
        request = NewFilesDeletionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FilesDeletions require credential")
    }

    request.SetContext(ctx)
    
    response = NewFilesDeletionsResponse()
    err = c.Send(request, response)
    return
}

func NewFilesListRequest() (request *FilesListRequest) {
    request = &FilesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "FilesList")
    
    
    return
}

func NewFilesListResponse() (response *FilesListResponse) {
    response = &FilesListResponse{} 
    return

}

// FilesList
// 文件列表。
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) FilesList(request *FilesListRequest) (response *FilesListResponse, err error) {
    return c.FilesListWithContext(context.Background(), request)
}

// FilesList
// 文件列表。
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) FilesListWithContext(ctx context.Context, request *FilesListRequest) (response *FilesListResponse, err error) {
    if request == nil {
        request = NewFilesListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FilesList require credential")
    }

    request.SetContext(ctx)
    
    response = NewFilesListResponse()
    err = c.Send(request, response)
    return
}

func NewFilesUploadsRequest() (request *FilesUploadsRequest) {
    request = &FilesUploadsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "FilesUploads")
    
    
    return
}

func NewFilesUploadsResponse() (response *FilesUploadsResponse) {
    response = &FilesUploadsResponse{} 
    return

}

// FilesUploads
// 上传用于不同用途的文件。
//
// 当前用途仅支持 hunyuan 等模型的文档理解。
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) FilesUploads(request *FilesUploadsRequest) (response *FilesUploadsResponse, err error) {
    return c.FilesUploadsWithContext(context.Background(), request)
}

// FilesUploads
// 上传用于不同用途的文件。
//
// 当前用途仅支持 hunyuan 等模型的文档理解。
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) FilesUploadsWithContext(ctx context.Context, request *FilesUploadsRequest) (response *FilesUploadsResponse, err error) {
    if request == nil {
        request = NewFilesUploadsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FilesUploads require credential")
    }

    request.SetContext(ctx)
    
    response = NewFilesUploadsResponse()
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

func NewGetThreadRequest() (request *GetThreadRequest) {
    request = &GetThreadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "GetThread")
    
    
    return
}

func NewGetThreadResponse() (response *GetThreadResponse) {
    response = &GetThreadResponse{} 
    return

}

// GetThread
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) GetThread(request *GetThreadRequest) (response *GetThreadResponse, err error) {
    return c.GetThreadWithContext(context.Background(), request)
}

// GetThread
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) GetThreadWithContext(ctx context.Context, request *GetThreadRequest) (response *GetThreadResponse, err error) {
    if request == nil {
        request = NewGetThreadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetThread require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetThreadResponse()
    err = c.Send(request, response)
    return
}

func NewGetThreadMessageRequest() (request *GetThreadMessageRequest) {
    request = &GetThreadMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "GetThreadMessage")
    
    
    return
}

func NewGetThreadMessageResponse() (response *GetThreadMessageResponse) {
    response = &GetThreadMessageResponse{} 
    return

}

// GetThreadMessage
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) GetThreadMessage(request *GetThreadMessageRequest) (response *GetThreadMessageResponse, err error) {
    return c.GetThreadMessageWithContext(context.Background(), request)
}

// GetThreadMessage
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) GetThreadMessageWithContext(ctx context.Context, request *GetThreadMessageRequest) (response *GetThreadMessageResponse, err error) {
    if request == nil {
        request = NewGetThreadMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetThreadMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetThreadMessageResponse()
    err = c.Send(request, response)
    return
}

func NewGetThreadMessageListRequest() (request *GetThreadMessageListRequest) {
    request = &GetThreadMessageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "GetThreadMessageList")
    
    
    return
}

func NewGetThreadMessageListResponse() (response *GetThreadMessageListResponse) {
    response = &GetThreadMessageListResponse{} 
    return

}

// GetThreadMessageList
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) GetThreadMessageList(request *GetThreadMessageListRequest) (response *GetThreadMessageListResponse, err error) {
    return c.GetThreadMessageListWithContext(context.Background(), request)
}

// GetThreadMessageList
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) GetThreadMessageListWithContext(ctx context.Context, request *GetThreadMessageListRequest) (response *GetThreadMessageListResponse, err error) {
    if request == nil {
        request = NewGetThreadMessageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetThreadMessageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetThreadMessageListResponse()
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

func NewGroupChatCompletionsRequest() (request *GroupChatCompletionsRequest) {
    request = &GroupChatCompletionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "GroupChatCompletions")
    
    
    return
}

func NewGroupChatCompletionsResponse() (response *GroupChatCompletionsResponse) {
    response = &GroupChatCompletionsResponse{} 
    return

}

// GroupChatCompletions
// 如需使用OpenAI兼容接口， 请参考文档：[OpenAI 兼容接口](https://cloud.tencent.com/document/product/1729/111007)
//
// 
//
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) GroupChatCompletions(request *GroupChatCompletionsRequest) (response *GroupChatCompletionsResponse, err error) {
    return c.GroupChatCompletionsWithContext(context.Background(), request)
}

// GroupChatCompletions
// 如需使用OpenAI兼容接口， 请参考文档：[OpenAI 兼容接口](https://cloud.tencent.com/document/product/1729/111007)
//
// 
//
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) GroupChatCompletionsWithContext(ctx context.Context, request *GroupChatCompletionsRequest) (response *GroupChatCompletionsResponse, err error) {
    if request == nil {
        request = NewGroupChatCompletionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GroupChatCompletions require credential")
    }

    request.SetContext(ctx)
    
    response = NewGroupChatCompletionsResponse()
    err = c.Send(request, response)
    return
}

func NewImageQuestionRequest() (request *ImageQuestionRequest) {
    request = &ImageQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "ImageQuestion")
    
    
    return
}

func NewImageQuestionResponse() (response *ImageQuestionResponse) {
    response = &ImageQuestionResponse{} 
    return

}

// ImageQuestion
// 如需使用OpenAI兼容接口， 请参考文档：[OpenAI 兼容接口](https://cloud.tencent.com/document/product/1729/111007)
//
// 
//
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) ImageQuestion(request *ImageQuestionRequest) (response *ImageQuestionResponse, err error) {
    return c.ImageQuestionWithContext(context.Background(), request)
}

// ImageQuestion
// 如需使用OpenAI兼容接口， 请参考文档：[OpenAI 兼容接口](https://cloud.tencent.com/document/product/1729/111007)
//
// 
//
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) ImageQuestionWithContext(ctx context.Context, request *ImageQuestionRequest) (response *ImageQuestionResponse, err error) {
    if request == nil {
        request = NewImageQuestionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewQueryHunyuanImageChatJobRequest() (request *QueryHunyuanImageChatJobRequest) {
    request = &QueryHunyuanImageChatJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "QueryHunyuanImageChatJob")
    
    
    return
}

func NewQueryHunyuanImageChatJobResponse() (response *QueryHunyuanImageChatJobResponse) {
    response = &QueryHunyuanImageChatJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuanImageChatJob
// 混元生图（多轮对话）接口基于混元大模型，将根据输入的文本描述生成图像，支持通过多轮对话的方式不断调整图像内容。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本和前置对话 ID 等，提交一个混元生图多轮对话异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得在上一轮对话基础上继续生成的图像结果。
//
// 混元生图（多轮对话）默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) QueryHunyuanImageChatJob(request *QueryHunyuanImageChatJobRequest) (response *QueryHunyuanImageChatJobResponse, err error) {
    return c.QueryHunyuanImageChatJobWithContext(context.Background(), request)
}

// QueryHunyuanImageChatJob
// 混元生图（多轮对话）接口基于混元大模型，将根据输入的文本描述生成图像，支持通过多轮对话的方式不断调整图像内容。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本和前置对话 ID 等，提交一个混元生图多轮对话异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得在上一轮对话基础上继续生成的图像结果。
//
// 混元生图（多轮对话）默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) QueryHunyuanImageChatJobWithContext(ctx context.Context, request *QueryHunyuanImageChatJobRequest) (response *QueryHunyuanImageChatJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuanImageChatJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuanImageChatJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuanImageChatJobResponse()
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
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
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
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
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

func NewQueryHunyuanTo3DJobRequest() (request *QueryHunyuanTo3DJobRequest) {
    request = &QueryHunyuanTo3DJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "QueryHunyuanTo3DJob")
    
    
    return
}

func NewQueryHunyuanTo3DJobResponse() (response *QueryHunyuanTo3DJobResponse) {
    response = &QueryHunyuanTo3DJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuanTo3DJob
// 查询混元生3D任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) QueryHunyuanTo3DJob(request *QueryHunyuanTo3DJobRequest) (response *QueryHunyuanTo3DJobResponse, err error) {
    return c.QueryHunyuanTo3DJobWithContext(context.Background(), request)
}

// QueryHunyuanTo3DJob
// 查询混元生3D任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) QueryHunyuanTo3DJobWithContext(ctx context.Context, request *QueryHunyuanTo3DJobRequest) (response *QueryHunyuanTo3DJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuanTo3DJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuanTo3DJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuanTo3DJobResponse()
    err = c.Send(request, response)
    return
}

func NewRunThreadRequest() (request *RunThreadRequest) {
    request = &RunThreadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "RunThread")
    
    
    return
}

func NewRunThreadResponse() (response *RunThreadResponse) {
    response = &RunThreadResponse{} 
    return

}

// RunThread
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) RunThread(request *RunThreadRequest) (response *RunThreadResponse, err error) {
    return c.RunThreadWithContext(context.Background(), request)
}

// RunThread
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
func (c *Client) RunThreadWithContext(ctx context.Context, request *RunThreadRequest) (response *RunThreadResponse, err error) {
    if request == nil {
        request = NewRunThreadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunThread require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunThreadResponse()
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

func NewSubmitHunyuanImageChatJobRequest() (request *SubmitHunyuanImageChatJobRequest) {
    request = &SubmitHunyuanImageChatJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "SubmitHunyuanImageChatJob")
    
    
    return
}

func NewSubmitHunyuanImageChatJobResponse() (response *SubmitHunyuanImageChatJobResponse) {
    response = &SubmitHunyuanImageChatJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanImageChatJob
// 混元生图（多轮对话）接口基于混元大模型，将根据输入的文本描述生成图像，支持通过多轮对话的方式不断调整图像内容。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本和前置对话 ID 等，提交一个混元生图多轮对话异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得在上一轮对话基础上继续生成的图像结果。
//
// 混元生图（多轮对话）默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  RESOURCEINSUFFICIENT_CHARGERESOURCEEXHAUST = "ResourceInsufficient.ChargeResourceExhaust"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) SubmitHunyuanImageChatJob(request *SubmitHunyuanImageChatJobRequest) (response *SubmitHunyuanImageChatJobResponse, err error) {
    return c.SubmitHunyuanImageChatJobWithContext(context.Background(), request)
}

// SubmitHunyuanImageChatJob
// 混元生图（多轮对话）接口基于混元大模型，将根据输入的文本描述生成图像，支持通过多轮对话的方式不断调整图像内容。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本和前置对话 ID 等，提交一个混元生图多轮对话异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得在上一轮对话基础上继续生成的图像结果。
//
// 混元生图（多轮对话）默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  RESOURCEINSUFFICIENT_CHARGERESOURCEEXHAUST = "ResourceInsufficient.ChargeResourceExhaust"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) SubmitHunyuanImageChatJobWithContext(ctx context.Context, request *SubmitHunyuanImageChatJobRequest) (response *SubmitHunyuanImageChatJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanImageChatJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanImageChatJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanImageChatJobResponse()
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

func NewSubmitHunyuanTo3DJobRequest() (request *SubmitHunyuanTo3DJobRequest) {
    request = &SubmitHunyuanTo3DJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "SubmitHunyuanTo3DJob")
    
    
    return
}

func NewSubmitHunyuanTo3DJobResponse() (response *SubmitHunyuanTo3DJobResponse) {
    response = &SubmitHunyuanTo3DJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanTo3DJob
// 提交混元生3D任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESIZE = "InvalidParameterValue.InvalidImageSize"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_CHARGERESOURCEEXHAUST = "ResourceInsufficient.ChargeResourceExhaust"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SubmitHunyuanTo3DJob(request *SubmitHunyuanTo3DJobRequest) (response *SubmitHunyuanTo3DJobResponse, err error) {
    return c.SubmitHunyuanTo3DJobWithContext(context.Background(), request)
}

// SubmitHunyuanTo3DJob
// 提交混元生3D任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESIZE = "InvalidParameterValue.InvalidImageSize"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_CHARGERESOURCEEXHAUST = "ResourceInsufficient.ChargeResourceExhaust"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SubmitHunyuanTo3DJobWithContext(ctx context.Context, request *SubmitHunyuanTo3DJobRequest) (response *SubmitHunyuanTo3DJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanTo3DJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanTo3DJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanTo3DJobResponse()
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
