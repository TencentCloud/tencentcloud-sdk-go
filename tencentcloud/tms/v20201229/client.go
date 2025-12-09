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

package v20201229

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-29"

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


func NewCreateFinancialLLMTaskRequest() (request *CreateFinancialLLMTaskRequest) {
    request = &CreateFinancialLLMTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tms", APIVersion, "CreateFinancialLLMTask")
    
    
    return
}

func NewCreateFinancialLLMTaskResponse() (response *CreateFinancialLLMTaskResponse) {
    response = &CreateFinancialLLMTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFinancialLLMTask
// 本接口适用于“金融大模型审校”服务。在对接前，请参考快速入门文档并配置好业务基础信息。
//
// - **快速入门**：[快速入门文档](https://cloud.tencent.com/document/product/1124/124604)
//
// 
//
// ### 接口功能说明：
//
// 由于大模型审校服务耗时较长，通常达到分钟级，因此采用异步模式，整体流程分为两步：
//
// 1. 创建金融大模型审校任务（详见本文档）。
//
// 2. 查询审校结果（详见 [查询结果文档](https://cloud.tencent.com/document/product/1124/124463)）。
//
// 
//
// ### 接口调用说明：
//
// - **请求域名**：tms.tencentcloudapi.com
//
// - **并发限制**：每个账号最多可同时进行3个审校任务。
//
// - **支持的文件格式**：纯文本、PDF、DOC、DOCX。
//
// 
//
// ### 文件限制说明：
//
// - **文档大小限制**：PDF/DOC/DOCX 格式文件不超过 200M（该大小为Base64编码后）。
//
// - **文档下载时长**：不超过 15 秒（建议将文档存储在腾讯云 URL，以确保更高的下载稳定性）。
//
// 可能返回的错误码:
//  INTERNALERROR_QUERYREQLIMITED = "InternalError.QueryReqLimited"
//  INVALIDPARAMETERVALUE_INVALIDBIZTYPE = "InvalidParameterValue.InvalidBizType"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDCONTENTTYPE = "InvalidParameterValue.InvalidContentType"
//  INVALIDPARAMETERVALUE_INVALIDFILENAME = "InvalidParameterValue.InvalidFileName"
//  INVALIDPARAMETERVALUE_INVALIDFILETYPE = "InvalidParameterValue.InvalidFileType"
//  INVALIDPARAMETERVALUE_INVALIDFILEURL = "InvalidParameterValue.InvalidFileUrl"
//  INVALIDPARAMETERVALUE_INVALIDREQUIREMENT = "InvalidParameterValue.InvalidRequirement"
//  INVALIDPARAMETERVALUE_INVALIDVERBOSE = "InvalidParameterValue.InvalidVerbose"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) CreateFinancialLLMTask(request *CreateFinancialLLMTaskRequest) (response *CreateFinancialLLMTaskResponse, err error) {
    return c.CreateFinancialLLMTaskWithContext(context.Background(), request)
}

// CreateFinancialLLMTask
// 本接口适用于“金融大模型审校”服务。在对接前，请参考快速入门文档并配置好业务基础信息。
//
// - **快速入门**：[快速入门文档](https://cloud.tencent.com/document/product/1124/124604)
//
// 
//
// ### 接口功能说明：
//
// 由于大模型审校服务耗时较长，通常达到分钟级，因此采用异步模式，整体流程分为两步：
//
// 1. 创建金融大模型审校任务（详见本文档）。
//
// 2. 查询审校结果（详见 [查询结果文档](https://cloud.tencent.com/document/product/1124/124463)）。
//
// 
//
// ### 接口调用说明：
//
// - **请求域名**：tms.tencentcloudapi.com
//
// - **并发限制**：每个账号最多可同时进行3个审校任务。
//
// - **支持的文件格式**：纯文本、PDF、DOC、DOCX。
//
// 
//
// ### 文件限制说明：
//
// - **文档大小限制**：PDF/DOC/DOCX 格式文件不超过 200M（该大小为Base64编码后）。
//
// - **文档下载时长**：不超过 15 秒（建议将文档存储在腾讯云 URL，以确保更高的下载稳定性）。
//
// 可能返回的错误码:
//  INTERNALERROR_QUERYREQLIMITED = "InternalError.QueryReqLimited"
//  INVALIDPARAMETERVALUE_INVALIDBIZTYPE = "InvalidParameterValue.InvalidBizType"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDCONTENTTYPE = "InvalidParameterValue.InvalidContentType"
//  INVALIDPARAMETERVALUE_INVALIDFILENAME = "InvalidParameterValue.InvalidFileName"
//  INVALIDPARAMETERVALUE_INVALIDFILETYPE = "InvalidParameterValue.InvalidFileType"
//  INVALIDPARAMETERVALUE_INVALIDFILEURL = "InvalidParameterValue.InvalidFileUrl"
//  INVALIDPARAMETERVALUE_INVALIDREQUIREMENT = "InvalidParameterValue.InvalidRequirement"
//  INVALIDPARAMETERVALUE_INVALIDVERBOSE = "InvalidParameterValue.InvalidVerbose"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) CreateFinancialLLMTaskWithContext(ctx context.Context, request *CreateFinancialLLMTaskRequest) (response *CreateFinancialLLMTaskResponse, err error) {
    if request == nil {
        request = NewCreateFinancialLLMTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tms", APIVersion, "CreateFinancialLLMTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFinancialLLMTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFinancialLLMTaskResponse()
    err = c.Send(request, response)
    return
}

func NewGetFinancialLLMTaskResultRequest() (request *GetFinancialLLMTaskResultRequest) {
    request = &GetFinancialLLMTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tms", APIVersion, "GetFinancialLLMTaskResult")
    
    
    return
}

func NewGetFinancialLLMTaskResultResponse() (response *GetFinancialLLMTaskResultResponse) {
    response = &GetFinancialLLMTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFinancialLLMTaskResult
// 本接口适用于“金融大模型审校”服务的结果查询。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  INVALIDPARAMETERVALUE_TASKIDNOTFOUND = "InvalidParameterValue.TaskIdNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetFinancialLLMTaskResult(request *GetFinancialLLMTaskResultRequest) (response *GetFinancialLLMTaskResultResponse, err error) {
    return c.GetFinancialLLMTaskResultWithContext(context.Background(), request)
}

// GetFinancialLLMTaskResult
// 本接口适用于“金融大模型审校”服务的结果查询。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  INVALIDPARAMETERVALUE_TASKIDNOTFOUND = "InvalidParameterValue.TaskIdNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) GetFinancialLLMTaskResultWithContext(ctx context.Context, request *GetFinancialLLMTaskResultRequest) (response *GetFinancialLLMTaskResultResponse, err error) {
    if request == nil {
        request = NewGetFinancialLLMTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tms", APIVersion, "GetFinancialLLMTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFinancialLLMTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFinancialLLMTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewTextModerationRequest() (request *TextModerationRequest) {
    request = &TextModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tms", APIVersion, "TextModeration")
    
    
    return
}

func NewTextModerationResponse() (response *TextModerationResponse) {
    response = &TextModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextModeration
// 本接口提供“内容安全”和“AI生成识别”服务。在对接之前，请先参考快速入门，以配置业务基础信息。
//
// - **内容安全**：[快速入门](https://cloud.tencent.com/document/product/1124/37119)
//
// - **AI生成识别**：[快速入门](https://cloud.tencent.com/document/product/1124/118694)
//
// 
//
// ### 接口功能说明：
//
// - **内容安全**：对输入的文本，识别其中是否存在色情、违法等风险，返回处置建议、风险标签及对应的模型阈值。
//
// - **AI生成识别**：对输入的文本，判断其是否为AI工具生成，返回AI生成的概率分数。
//
// 
//
// ### 接口调用说明：
//
// - **请求域名**：tms.tencentcloudapi.com
//
// - **文本长度限制**：最长10,000个字符（以Unicode编码计量）。
//
// - **请求频率**：内容安全默认1000次/秒，AI生成识别默认50次/秒。
//
// - **支持语言**：中文、英文。
//
// 可能返回的错误码:
//  INTERNALERROR_ERRTEXTTIMEOUT = "InternalError.ErrTextTimeOut"
//  INTERNALERROR_QUERYREQLIMITED = "InternalError.QueryReqLimited"
//  INVALIDPARAMETER_ERRACTION = "InvalidParameter.ErrAction"
//  INVALIDPARAMETER_ERRTEXTCONTENTLEN = "InvalidParameter.ErrTextContentLen"
//  INVALIDPARAMETER_ERRTEXTCONTENTTYPE = "InvalidParameter.ErrTextContentType"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTLEN = "InvalidParameterValue.ErrTextContentLen"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTTYPE = "InvalidParameterValue.ErrTextContentType"
//  INVALIDPARAMETERVALUE_ERRTYPE = "InvalidParameterValue.ErrType"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) TextModeration(request *TextModerationRequest) (response *TextModerationResponse, err error) {
    return c.TextModerationWithContext(context.Background(), request)
}

// TextModeration
// 本接口提供“内容安全”和“AI生成识别”服务。在对接之前，请先参考快速入门，以配置业务基础信息。
//
// - **内容安全**：[快速入门](https://cloud.tencent.com/document/product/1124/37119)
//
// - **AI生成识别**：[快速入门](https://cloud.tencent.com/document/product/1124/118694)
//
// 
//
// ### 接口功能说明：
//
// - **内容安全**：对输入的文本，识别其中是否存在色情、违法等风险，返回处置建议、风险标签及对应的模型阈值。
//
// - **AI生成识别**：对输入的文本，判断其是否为AI工具生成，返回AI生成的概率分数。
//
// 
//
// ### 接口调用说明：
//
// - **请求域名**：tms.tencentcloudapi.com
//
// - **文本长度限制**：最长10,000个字符（以Unicode编码计量）。
//
// - **请求频率**：内容安全默认1000次/秒，AI生成识别默认50次/秒。
//
// - **支持语言**：中文、英文。
//
// 可能返回的错误码:
//  INTERNALERROR_ERRTEXTTIMEOUT = "InternalError.ErrTextTimeOut"
//  INTERNALERROR_QUERYREQLIMITED = "InternalError.QueryReqLimited"
//  INVALIDPARAMETER_ERRACTION = "InvalidParameter.ErrAction"
//  INVALIDPARAMETER_ERRTEXTCONTENTLEN = "InvalidParameter.ErrTextContentLen"
//  INVALIDPARAMETER_ERRTEXTCONTENTTYPE = "InvalidParameter.ErrTextContentType"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTLEN = "InvalidParameterValue.ErrTextContentLen"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTTYPE = "InvalidParameterValue.ErrTextContentType"
//  INVALIDPARAMETERVALUE_ERRTYPE = "InvalidParameterValue.ErrType"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) TextModerationWithContext(ctx context.Context, request *TextModerationRequest) (response *TextModerationResponse, err error) {
    if request == nil {
        request = NewTextModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tms", APIVersion, "TextModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextModerationResponse()
    err = c.Send(request, response)
    return
}
