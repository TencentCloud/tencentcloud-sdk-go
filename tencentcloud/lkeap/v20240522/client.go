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

package v20240522

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-05-22"

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
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ChatCompletions")
    
    
    return
}

func NewChatCompletionsResponse() (response *ChatCompletionsResponse) {
    response = &ChatCompletionsResponse{} 
    return

}

// ChatCompletions
// ### 接口功能
//
// 
//
// 调用接口，发起一次对话请求。默认该接口下的单账号QPM上限为15000 ，TPM上限为1200000
//
// 如需使用OpenAI兼容接口， 请参考文档：[Deepseek OpenAI对话接口](https://cloud.tencent.com/document/product/1772/115969)
//
// 
//
// #### 在线体验
//
// 如您希望在网页内直接体验 DeepSeek 模型对话，推荐您前往[腾讯云智能体开发平台](https://cloud.tencent.com/product/lke)，使用[ DeepSeek 联网助手](https://lke.cloud.tencent.com/webim_exp/#/chat/wQrAwR)。
//
// 
//
// #### 已支持的模型
//
// - DeepSeek-V3-0324（model 参数值为**deepseek-v3-0324**）
//
//     - DeepSeek-V3-0324为671B参数MoE模型，在编程与技术能力、上下文理解与长文本处理等方面优势突出。
//
//     - 支持128K上下文长度，最大输出16k（不含思维链）。
//
//     - 注意：相比于DeepSeek-V3，DeepSeek-V3-0324仅更新了模型权重，未增加参数量。总模型大小为685B，其中包括671B的主模型权重和 14B 的多令牌预测（MTP）模块权重，后续均描述主模型参数量。
//
// - DeepSeek-V3（model 参数值为**deepseek-v3**）
//
//     - DeepSeek-V3为671B参数MoE模型，在百科知识、数学推理等多项任务上优势突出，评测成绩在主流榜单中位列开源模型榜首。
//
//     - 支持64K上下文长度，最大输出16k。
//
// - DeepSeek-R1（model 参数值为**deepseek-r1**）
//
//     - DeepSeek-R1为671B模型，使用强化学习训练，推理过程包含大量反思和验证，思维链长度可达数万字。 该系列模型在数学、代码以及各种复杂逻辑推理任务上推理效果优异，并为用户展现了完整的思考过程。
//
//     -  支持96K上下文长度，最大输入长度64k，最大输出16k（默认4k），最大思维链输出长度32k。
//
// - DeepSeek-Prover-V2（model 参数值为**deepseek-prover-v2**）
//
//     - DeepSeek-Prover-V2 为671B 参数 MoE 模型，在数学定理证明和复杂计算任务中表现出色。
//
//     -  支持64K上下文长度，最大输出16k。
//
// 
//
// 
//
// ### 计费说明
//
// 
//
// - 标准计费（2025年2月26日起生效），计费模式为后付费小时结，为保证您账户资源的正常使用，请提前[开通后付费](https://lke.cloud.tencent.com/lke#/app/system/charge/postpaid)并及时[充值](https://console.cloud.tencent.com/expense/recharge)。
//
// 
//
//     -  DeepSeek-R1 模型   | 输入：0.004元/千token | 输出（含思维链）：0.016元/千token
//
// 
//
//     - DeepSeek-V3 模型 | 输入：0.002元/千token | 输出：0.008元/千token
//
// 
//
//     - DeepSeek-V3-0324 模型 | 输入：0.002元/千token | 输出：0.008元/千token
//
// 
//
//     - DeepSeek-Prover-V2 模型 | 暂不计费
//
// 
//
// 
//
// ### Openai兼容协议接口
//
// 知识引擎原子能力大模型对话 API 兼容了 OpenAI 的接口规范，这意味着您可以直接使用 OpenAI 官方提供的 SDK 来调用大模型对话接口。您仅需要将 base_url 和 [api_key](https://cloud.tencent.com/document/product/1772/115970) 替换成相关配置，不需要对应用做额外修改，即可无缝将您的应用切换到相应的大模型。请参考文档：[Deepseek OpenAI对话接口](https://cloud.tencent.com/document/product/1772/115969)。
//
// > base_url：  https://api.lkeap.cloud.tencent.com/v1
//
// 
//
// > api_key的获取请参考[API KEY管理](https://cloud.tencent.com/document/product/1772/115970)
//
// 
//
// 
//
// ### 快速接入
//
// 1. 完成[实名认证](https://console.cloud.tencent.com/developer/auth)。
//
// 2. 主账户前往[控制台](https://console.cloud.tencent.com/lkeap)开通服务。若为子账户，需由主账号在[权限管理](https://console.cloud.tencent.com/cam)中为子账号授权，关联预设策略，策略名称：QcloudLKEAPFullAccess。
//
// 3. 通过API Explorer[在线调试](https://console.cloud.tencent.com/api/explorer?Product=lkeap&Version=2024-05-22&Action=ChatCompletions)。
//
// 4. 使用[官方 SDK ](https://cloud.tencent.com/document/product/1772/115963#SDK)调用本接口（支持Python/Java/PHP/Go/Node.js/.NET等语言）。
//
// 
//
// -----------
//
// 
//
// ### SDK调用示例
//
// 通过本地代码调用本接口（支持Python/Java/PHP/Go/Node.js/.NET等语言）：下面的代码以 Python 语言为例，展示如何访问腾讯云上的DeepSeek模型API的样例。
//
// （1）安装环境
//
// ```
//
// python3 -m pip install --upgrade tencentcloud-sdk-python-common
//
// python3 -m pip install --upgrade tencentcloud-sdk-python-lkeap
//
// ```
//
// 
//
// （2）示例代码
//
// 
//
// - 其中SecretKey和SecretID需要从腾讯云控制台获取
//
// 
//
// - 参数params中模型Model字段可以选择“deepseek-r1“和“deepseek-v3”
//
// 
//
// ```
//
// # -*- coding: utf-8 -*-
//
// import json
//
// 
//
// from tencentcloud.common.common_client import CommonClient
//
// from tencentcloud.common import credential
//
// from tencentcloud.common.exception.tencent_cloud_sdk_exception import TencentCloudSDKException
//
// from tencentcloud.common.profile.client_profile import ClientProfile
//
// from tencentcloud.common.profile.http_profile import HttpProfile
//
// 
//
// class NonStreamResponse(object):
//
//     def __init__(self):
//
//         self.response = ""
//
// 
//
//     def _deserialize(self, obj):
//
//         self.response = json.dumps(obj)
//
// 
//
// try:
//
//     # 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
//
//     # 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
//
//     # 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
//
//     cred = credential.Credential("", "")
//
// 
//
//     httpProfile = HttpProfile()
//
//     httpProfile.endpoint = "lkeap.tencentcloudapi.com"
//
//     httpProfile.reqTimeout = 40000  # 流式接口可能耗时较长
//
//     clientProfile = ClientProfile()
//
//     clientProfile.httpProfile = httpProfile
//
// 
//
//     params = "{\"Model\":\"deepseek-r1\",\"Messages\":[{\"Role\":\"user\",\"Content\":\"你好\"}],\"Stream\":true}";
//
//     common_client = CommonClient("lkeap", "2024-05-22", cred, "ap-guangzhou", profile=clientProfile)
//
//     resp = common_client._call_and_deserialize("ChatCompletions", json.loads(params), NonStreamResponse)
//
//     if isinstance(resp, NonStreamResponse):  # 非流式响应
//
//         print(resp.response)
//
//     else:  # 流式响应
//
//         for event in resp:
//
//             print(event)
//
// except TencentCloudSDKException as err:
//
//     print(err)
//
// 
//
// ```
//
// 
//
// **DeepSeek-R1使用建议**
//
// 
//
// 1. 将温度设置在 0.5-0.7 范围内（建议为0.6），以防止无休止的重复或不连贯的输出。
//
// 2. 避免添加system prompt，所有说明都应包含在user prompt中。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) ChatCompletions(request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
    return c.ChatCompletionsWithContext(context.Background(), request)
}

// ChatCompletions
// ### 接口功能
//
// 
//
// 调用接口，发起一次对话请求。默认该接口下的单账号QPM上限为15000 ，TPM上限为1200000
//
// 如需使用OpenAI兼容接口， 请参考文档：[Deepseek OpenAI对话接口](https://cloud.tencent.com/document/product/1772/115969)
//
// 
//
// #### 在线体验
//
// 如您希望在网页内直接体验 DeepSeek 模型对话，推荐您前往[腾讯云智能体开发平台](https://cloud.tencent.com/product/lke)，使用[ DeepSeek 联网助手](https://lke.cloud.tencent.com/webim_exp/#/chat/wQrAwR)。
//
// 
//
// #### 已支持的模型
//
// - DeepSeek-V3-0324（model 参数值为**deepseek-v3-0324**）
//
//     - DeepSeek-V3-0324为671B参数MoE模型，在编程与技术能力、上下文理解与长文本处理等方面优势突出。
//
//     - 支持128K上下文长度，最大输出16k（不含思维链）。
//
//     - 注意：相比于DeepSeek-V3，DeepSeek-V3-0324仅更新了模型权重，未增加参数量。总模型大小为685B，其中包括671B的主模型权重和 14B 的多令牌预测（MTP）模块权重，后续均描述主模型参数量。
//
// - DeepSeek-V3（model 参数值为**deepseek-v3**）
//
//     - DeepSeek-V3为671B参数MoE模型，在百科知识、数学推理等多项任务上优势突出，评测成绩在主流榜单中位列开源模型榜首。
//
//     - 支持64K上下文长度，最大输出16k。
//
// - DeepSeek-R1（model 参数值为**deepseek-r1**）
//
//     - DeepSeek-R1为671B模型，使用强化学习训练，推理过程包含大量反思和验证，思维链长度可达数万字。 该系列模型在数学、代码以及各种复杂逻辑推理任务上推理效果优异，并为用户展现了完整的思考过程。
//
//     -  支持96K上下文长度，最大输入长度64k，最大输出16k（默认4k），最大思维链输出长度32k。
//
// - DeepSeek-Prover-V2（model 参数值为**deepseek-prover-v2**）
//
//     - DeepSeek-Prover-V2 为671B 参数 MoE 模型，在数学定理证明和复杂计算任务中表现出色。
//
//     -  支持64K上下文长度，最大输出16k。
//
// 
//
// 
//
// ### 计费说明
//
// 
//
// - 标准计费（2025年2月26日起生效），计费模式为后付费小时结，为保证您账户资源的正常使用，请提前[开通后付费](https://lke.cloud.tencent.com/lke#/app/system/charge/postpaid)并及时[充值](https://console.cloud.tencent.com/expense/recharge)。
//
// 
//
//     -  DeepSeek-R1 模型   | 输入：0.004元/千token | 输出（含思维链）：0.016元/千token
//
// 
//
//     - DeepSeek-V3 模型 | 输入：0.002元/千token | 输出：0.008元/千token
//
// 
//
//     - DeepSeek-V3-0324 模型 | 输入：0.002元/千token | 输出：0.008元/千token
//
// 
//
//     - DeepSeek-Prover-V2 模型 | 暂不计费
//
// 
//
// 
//
// ### Openai兼容协议接口
//
// 知识引擎原子能力大模型对话 API 兼容了 OpenAI 的接口规范，这意味着您可以直接使用 OpenAI 官方提供的 SDK 来调用大模型对话接口。您仅需要将 base_url 和 [api_key](https://cloud.tencent.com/document/product/1772/115970) 替换成相关配置，不需要对应用做额外修改，即可无缝将您的应用切换到相应的大模型。请参考文档：[Deepseek OpenAI对话接口](https://cloud.tencent.com/document/product/1772/115969)。
//
// > base_url：  https://api.lkeap.cloud.tencent.com/v1
//
// 
//
// > api_key的获取请参考[API KEY管理](https://cloud.tencent.com/document/product/1772/115970)
//
// 
//
// 
//
// ### 快速接入
//
// 1. 完成[实名认证](https://console.cloud.tencent.com/developer/auth)。
//
// 2. 主账户前往[控制台](https://console.cloud.tencent.com/lkeap)开通服务。若为子账户，需由主账号在[权限管理](https://console.cloud.tencent.com/cam)中为子账号授权，关联预设策略，策略名称：QcloudLKEAPFullAccess。
//
// 3. 通过API Explorer[在线调试](https://console.cloud.tencent.com/api/explorer?Product=lkeap&Version=2024-05-22&Action=ChatCompletions)。
//
// 4. 使用[官方 SDK ](https://cloud.tencent.com/document/product/1772/115963#SDK)调用本接口（支持Python/Java/PHP/Go/Node.js/.NET等语言）。
//
// 
//
// -----------
//
// 
//
// ### SDK调用示例
//
// 通过本地代码调用本接口（支持Python/Java/PHP/Go/Node.js/.NET等语言）：下面的代码以 Python 语言为例，展示如何访问腾讯云上的DeepSeek模型API的样例。
//
// （1）安装环境
//
// ```
//
// python3 -m pip install --upgrade tencentcloud-sdk-python-common
//
// python3 -m pip install --upgrade tencentcloud-sdk-python-lkeap
//
// ```
//
// 
//
// （2）示例代码
//
// 
//
// - 其中SecretKey和SecretID需要从腾讯云控制台获取
//
// 
//
// - 参数params中模型Model字段可以选择“deepseek-r1“和“deepseek-v3”
//
// 
//
// ```
//
// # -*- coding: utf-8 -*-
//
// import json
//
// 
//
// from tencentcloud.common.common_client import CommonClient
//
// from tencentcloud.common import credential
//
// from tencentcloud.common.exception.tencent_cloud_sdk_exception import TencentCloudSDKException
//
// from tencentcloud.common.profile.client_profile import ClientProfile
//
// from tencentcloud.common.profile.http_profile import HttpProfile
//
// 
//
// class NonStreamResponse(object):
//
//     def __init__(self):
//
//         self.response = ""
//
// 
//
//     def _deserialize(self, obj):
//
//         self.response = json.dumps(obj)
//
// 
//
// try:
//
//     # 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
//
//     # 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
//
//     # 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
//
//     cred = credential.Credential("", "")
//
// 
//
//     httpProfile = HttpProfile()
//
//     httpProfile.endpoint = "lkeap.tencentcloudapi.com"
//
//     httpProfile.reqTimeout = 40000  # 流式接口可能耗时较长
//
//     clientProfile = ClientProfile()
//
//     clientProfile.httpProfile = httpProfile
//
// 
//
//     params = "{\"Model\":\"deepseek-r1\",\"Messages\":[{\"Role\":\"user\",\"Content\":\"你好\"}],\"Stream\":true}";
//
//     common_client = CommonClient("lkeap", "2024-05-22", cred, "ap-guangzhou", profile=clientProfile)
//
//     resp = common_client._call_and_deserialize("ChatCompletions", json.loads(params), NonStreamResponse)
//
//     if isinstance(resp, NonStreamResponse):  # 非流式响应
//
//         print(resp.response)
//
//     else:  # 流式响应
//
//         for event in resp:
//
//             print(event)
//
// except TencentCloudSDKException as err:
//
//     print(err)
//
// 
//
// ```
//
// 
//
// **DeepSeek-R1使用建议**
//
// 
//
// 1. 将温度设置在 0.5-0.7 范围内（建议为0.6），以防止无休止的重复或不连贯的输出。
//
// 2. 避免添加system prompt，所有说明都应包含在user prompt中。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
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

func NewCreateAttributeLabelRequest() (request *CreateAttributeLabelRequest) {
    request = &CreateAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateAttributeLabel")
    
    
    return
}

func NewCreateAttributeLabelResponse() (response *CreateAttributeLabelResponse) {
    response = &CreateAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAttributeLabel
// 用于为问答对创建属性标签，以便对内容进行分类和管理。 使用场景：当需要为问答对添加分类标签和属性标记时使用，比如为问答对添加“售后”标签。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateAttributeLabel(request *CreateAttributeLabelRequest) (response *CreateAttributeLabelResponse, err error) {
    return c.CreateAttributeLabelWithContext(context.Background(), request)
}

// CreateAttributeLabel
// 用于为问答对创建属性标签，以便对内容进行分类和管理。 使用场景：当需要为问答对添加分类标签和属性标记时使用，比如为问答对添加“售后”标签。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateAttributeLabelWithContext(ctx context.Context, request *CreateAttributeLabelRequest) (response *CreateAttributeLabelResponse, err error) {
    if request == nil {
        request = NewCreateAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKnowledgeBaseRequest() (request *CreateKnowledgeBaseRequest) {
    request = &CreateKnowledgeBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateKnowledgeBase")
    
    
    return
}

func NewCreateKnowledgeBaseResponse() (response *CreateKnowledgeBaseResponse) {
    response = &CreateKnowledgeBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKnowledgeBase
// 用于在系统中创建一个新的知识库。知识库是一个用于存储和管理知识条目的集合，可以包括文档、问答对、属性标签等。创建知识库后，可以向其中添加各种知识条目，以便在后续的知识检索中使用。 使用场景：当需要在系统中建立一个新的知识库以存储和管理特定领域或项目的知识条目时使用。例如，一个用户可能需要创建一个知识库，以存储用户指南、常见问题解答和技术文档。体验用户最多可创建3个知识库，付费用户最多可创建100个知识库。
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateKnowledgeBase(request *CreateKnowledgeBaseRequest) (response *CreateKnowledgeBaseResponse, err error) {
    return c.CreateKnowledgeBaseWithContext(context.Background(), request)
}

// CreateKnowledgeBase
// 用于在系统中创建一个新的知识库。知识库是一个用于存储和管理知识条目的集合，可以包括文档、问答对、属性标签等。创建知识库后，可以向其中添加各种知识条目，以便在后续的知识检索中使用。 使用场景：当需要在系统中建立一个新的知识库以存储和管理特定领域或项目的知识条目时使用。例如，一个用户可能需要创建一个知识库，以存储用户指南、常见问题解答和技术文档。体验用户最多可创建3个知识库，付费用户最多可创建100个知识库。
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateKnowledgeBaseWithContext(ctx context.Context, request *CreateKnowledgeBaseRequest) (response *CreateKnowledgeBaseResponse, err error) {
    if request == nil {
        request = NewCreateKnowledgeBaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKnowledgeBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKnowledgeBaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQARequest() (request *CreateQARequest) {
    request = &CreateQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateQA")
    
    
    return
}

func NewCreateQAResponse() (response *CreateQAResponse) {
    response = &CreateQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQA
// 用于创建新的问答对。问答对可以在SearchKnowledge接口知识检索时提供匹配的答案。 使用场景：当需要添加新的知识点和对应的问答对时使用，比如为产品添加新的常见问题解答。
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateQA(request *CreateQARequest) (response *CreateQAResponse, err error) {
    return c.CreateQAWithContext(context.Background(), request)
}

// CreateQA
// 用于创建新的问答对。问答对可以在SearchKnowledge接口知识检索时提供匹配的答案。 使用场景：当需要添加新的知识点和对应的问答对时使用，比如为产品添加新的常见问题解答。
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateQAWithContext(ctx context.Context, request *CreateQARequest) (response *CreateQAResponse, err error) {
    if request == nil {
        request = NewCreateQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQAResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReconstructDocumentFlowRequest() (request *CreateReconstructDocumentFlowRequest) {
    request = &CreateReconstructDocumentFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateReconstructDocumentFlow")
    
    
    return
}

func NewCreateReconstructDocumentFlowResponse() (response *CreateReconstructDocumentFlowResponse) {
    response = &CreateReconstructDocumentFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReconstructDocumentFlow
// 本接口为异步接口的发起请求接口，用于发起文档解析任务。
//
// 文档解析支持将图片或PDF、DOCX、PPTX、EXCEL等文件转换成Markdown格式文件，可解析包括表格、公式、图片、标题、段落、页眉、页脚等内容元素，并将内容智能转换成阅读顺序。具体支持文件类型请查看下方输入参数列表。
//
// 
//
// 体验期间单账号限制qps仅为1，若有正式接入需要请与产研团队沟通开放。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWFILETYPEERROR = "FailedOperation.UnKnowFileTypeError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateReconstructDocumentFlow(request *CreateReconstructDocumentFlowRequest) (response *CreateReconstructDocumentFlowResponse, err error) {
    return c.CreateReconstructDocumentFlowWithContext(context.Background(), request)
}

// CreateReconstructDocumentFlow
// 本接口为异步接口的发起请求接口，用于发起文档解析任务。
//
// 文档解析支持将图片或PDF、DOCX、PPTX、EXCEL等文件转换成Markdown格式文件，可解析包括表格、公式、图片、标题、段落、页眉、页脚等内容元素，并将内容智能转换成阅读顺序。具体支持文件类型请查看下方输入参数列表。
//
// 
//
// 体验期间单账号限制qps仅为1，若有正式接入需要请与产研团队沟通开放。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWFILETYPEERROR = "FailedOperation.UnKnowFileTypeError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateReconstructDocumentFlowWithContext(ctx context.Context, request *CreateReconstructDocumentFlowRequest) (response *CreateReconstructDocumentFlowResponse, err error) {
    if request == nil {
        request = NewCreateReconstructDocumentFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReconstructDocumentFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReconstructDocumentFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSplitDocumentFlowRequest() (request *CreateSplitDocumentFlowRequest) {
    request = &CreateSplitDocumentFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateSplitDocumentFlow")
    
    
    return
}

func NewCreateSplitDocumentFlowResponse() (response *CreateSplitDocumentFlowResponse) {
    response = &CreateSplitDocumentFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSplitDocumentFlow
// 用于创建一个文档拆分任务，支持多种文件类型，具备mllm能力，能够解析并深入理解图表中的信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateSplitDocumentFlow(request *CreateSplitDocumentFlowRequest) (response *CreateSplitDocumentFlowResponse, err error) {
    return c.CreateSplitDocumentFlowWithContext(context.Background(), request)
}

// CreateSplitDocumentFlow
// 用于创建一个文档拆分任务，支持多种文件类型，具备mllm能力，能够解析并深入理解图表中的信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateSplitDocumentFlowWithContext(ctx context.Context, request *CreateSplitDocumentFlowRequest) (response *CreateSplitDocumentFlowResponse, err error) {
    if request == nil {
        request = NewCreateSplitDocumentFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSplitDocumentFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSplitDocumentFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttributeLabelsRequest() (request *DeleteAttributeLabelsRequest) {
    request = &DeleteAttributeLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DeleteAttributeLabels")
    
    
    return
}

func NewDeleteAttributeLabelsResponse() (response *DeleteAttributeLabelsResponse) {
    response = &DeleteAttributeLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAttributeLabels
// 用于删除不再需要的属性标签。 使用场景：当某个标签不再使用时，可以将其删除以保持标签系统的整洁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DeleteAttributeLabels(request *DeleteAttributeLabelsRequest) (response *DeleteAttributeLabelsResponse, err error) {
    return c.DeleteAttributeLabelsWithContext(context.Background(), request)
}

// DeleteAttributeLabels
// 用于删除不再需要的属性标签。 使用场景：当某个标签不再使用时，可以将其删除以保持标签系统的整洁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DeleteAttributeLabelsWithContext(ctx context.Context, request *DeleteAttributeLabelsRequest) (response *DeleteAttributeLabelsResponse, err error) {
    if request == nil {
        request = NewDeleteAttributeLabelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAttributeLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAttributeLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDocsRequest() (request *DeleteDocsRequest) {
    request = &DeleteDocsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DeleteDocs")
    
    
    return
}

func NewDeleteDocsResponse() (response *DeleteDocsResponse) {
    response = &DeleteDocsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDocs
// 用于删除已有的文档。 使用场景：当某个文档不再需要时，可以将其删除以保持文档库的整洁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DeleteDocs(request *DeleteDocsRequest) (response *DeleteDocsResponse, err error) {
    return c.DeleteDocsWithContext(context.Background(), request)
}

// DeleteDocs
// 用于删除已有的文档。 使用场景：当某个文档不再需要时，可以将其删除以保持文档库的整洁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DeleteDocsWithContext(ctx context.Context, request *DeleteDocsRequest) (response *DeleteDocsResponse, err error) {
    if request == nil {
        request = NewDeleteDocsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDocs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKnowledgeBaseRequest() (request *DeleteKnowledgeBaseRequest) {
    request = &DeleteKnowledgeBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DeleteKnowledgeBase")
    
    
    return
}

func NewDeleteKnowledgeBaseResponse() (response *DeleteKnowledgeBaseResponse) {
    response = &DeleteKnowledgeBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteKnowledgeBase
// 用于从系统中删除一个现有的知识库。删除知识库将移除该知识库及其所有关联的知识条目（如文档、问答对、属性标签等）。该操作是不可逆的，请在执行前确认是否需要删除。**使用场景**：当某个知识库不再需要时，可以使用此接口将其从系统中删除。例如，一个项目结束后，其相关的知识库可能不再需要存储，可以使用该接口进行删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DeleteKnowledgeBase(request *DeleteKnowledgeBaseRequest) (response *DeleteKnowledgeBaseResponse, err error) {
    return c.DeleteKnowledgeBaseWithContext(context.Background(), request)
}

// DeleteKnowledgeBase
// 用于从系统中删除一个现有的知识库。删除知识库将移除该知识库及其所有关联的知识条目（如文档、问答对、属性标签等）。该操作是不可逆的，请在执行前确认是否需要删除。**使用场景**：当某个知识库不再需要时，可以使用此接口将其从系统中删除。例如，一个项目结束后，其相关的知识库可能不再需要存储，可以使用该接口进行删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DeleteKnowledgeBaseWithContext(ctx context.Context, request *DeleteKnowledgeBaseRequest) (response *DeleteKnowledgeBaseResponse, err error) {
    if request == nil {
        request = NewDeleteKnowledgeBaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKnowledgeBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKnowledgeBaseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQAsRequest() (request *DeleteQAsRequest) {
    request = &DeleteQAsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DeleteQAs")
    
    
    return
}

func NewDeleteQAsResponse() (response *DeleteQAsResponse) {
    response = &DeleteQAsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQAs
// 用于删除已有的问答对。 使用场景：当某个问答对不再适用或需要移除时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DeleteQAs(request *DeleteQAsRequest) (response *DeleteQAsResponse, err error) {
    return c.DeleteQAsWithContext(context.Background(), request)
}

// DeleteQAs
// 用于删除已有的问答对。 使用场景：当某个问答对不再适用或需要移除时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DeleteQAsWithContext(ctx context.Context, request *DeleteQAsRequest) (response *DeleteQAsResponse, err error) {
    if request == nil {
        request = NewDeleteQAsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQAs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQAsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocRequest() (request *DescribeDocRequest) {
    request = &DescribeDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DescribeDoc")
    
    
    return
}

func NewDescribeDocResponse() (response *DescribeDocResponse) {
    response = &DescribeDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDoc
// 用于查询特定文档的详细信息。 使用场景：当需要查看某个文档的具体内容和属性时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DescribeDoc(request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    return c.DescribeDocWithContext(context.Background(), request)
}

// DescribeDoc
// 用于查询特定文档的详细信息。 使用场景：当需要查看某个文档的具体内容和属性时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) DescribeDocWithContext(ctx context.Context, request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    if request == nil {
        request = NewDescribeDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocResponse()
    err = c.Send(request, response)
    return
}

func NewGetCharacterUsageRequest() (request *GetCharacterUsageRequest) {
    request = &GetCharacterUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetCharacterUsage")
    
    
    return
}

func NewGetCharacterUsageResponse() (response *GetCharacterUsageResponse) {
    response = &GetCharacterUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCharacterUsage
// 获取字符使用量统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) GetCharacterUsage(request *GetCharacterUsageRequest) (response *GetCharacterUsageResponse, err error) {
    return c.GetCharacterUsageWithContext(context.Background(), request)
}

// GetCharacterUsage
// 获取字符使用量统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) GetCharacterUsageWithContext(ctx context.Context, request *GetCharacterUsageRequest) (response *GetCharacterUsageResponse, err error) {
    if request == nil {
        request = NewGetCharacterUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCharacterUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCharacterUsageResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmbeddingRequest() (request *GetEmbeddingRequest) {
    request = &GetEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetEmbedding")
    
    
    return
}

func NewGetEmbeddingResponse() (response *GetEmbeddingResponse) {
    response = &GetEmbeddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetEmbedding
// 本接口（GetEmbedding）调用文本表示模型，将文本转化为用数值表示的向量形式，可用于文本检索、信息推荐、知识挖掘等场景。
//
// 本接口（GetEmbedding）有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetEmbedding(request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
    return c.GetEmbeddingWithContext(context.Background(), request)
}

// GetEmbedding
// 本接口（GetEmbedding）调用文本表示模型，将文本转化为用数值表示的向量形式，可用于文本检索、信息推荐、知识挖掘等场景。
//
// 本接口（GetEmbedding）有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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

func NewGetReconstructDocumentResultRequest() (request *GetReconstructDocumentResultRequest) {
    request = &GetReconstructDocumentResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetReconstructDocumentResult")
    
    
    return
}

func NewGetReconstructDocumentResultResponse() (response *GetReconstructDocumentResultResponse) {
    response = &GetReconstructDocumentResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetReconstructDocumentResult
// 本接口为异步接口的查询结果接口，用于获取文档解析处理结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetReconstructDocumentResult(request *GetReconstructDocumentResultRequest) (response *GetReconstructDocumentResultResponse, err error) {
    return c.GetReconstructDocumentResultWithContext(context.Background(), request)
}

// GetReconstructDocumentResult
// 本接口为异步接口的查询结果接口，用于获取文档解析处理结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetReconstructDocumentResultWithContext(ctx context.Context, request *GetReconstructDocumentResultRequest) (response *GetReconstructDocumentResultResponse, err error) {
    if request == nil {
        request = NewGetReconstructDocumentResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetReconstructDocumentResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetReconstructDocumentResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetSplitDocumentResultRequest() (request *GetSplitDocumentResultRequest) {
    request = &GetSplitDocumentResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetSplitDocumentResult")
    
    
    return
}

func NewGetSplitDocumentResultResponse() (response *GetSplitDocumentResultResponse) {
    response = &GetSplitDocumentResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSplitDocumentResult
// 查询文档拆分任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetSplitDocumentResult(request *GetSplitDocumentResultRequest) (response *GetSplitDocumentResultResponse, err error) {
    return c.GetSplitDocumentResultWithContext(context.Background(), request)
}

// GetSplitDocumentResult
// 查询文档拆分任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetSplitDocumentResultWithContext(ctx context.Context, request *GetSplitDocumentResultRequest) (response *GetSplitDocumentResultResponse, err error) {
    if request == nil {
        request = NewGetSplitDocumentResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSplitDocumentResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSplitDocumentResultResponse()
    err = c.Send(request, response)
    return
}

func NewImportQAsRequest() (request *ImportQAsRequest) {
    request = &ImportQAsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ImportQAs")
    
    
    return
}

func NewImportQAsResponse() (response *ImportQAsResponse) {
    response = &ImportQAsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportQAs
// 用于批量导入问答对，最多支持50,000条数据导入。通过此接口，可以将多个问答对一次性导入系统中，以便在后续的知识检索中使用。导入的问答对可以来自外部系统、文件或其他数据源。使用场景：当需要一次性导入大量的问答对时使用。例如，一个公司可能需要将其产品的常见问题解答从一个文档或外部系统批量导入到知识库中，以便用户可以通过知识检索系统进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ImportQAs(request *ImportQAsRequest) (response *ImportQAsResponse, err error) {
    return c.ImportQAsWithContext(context.Background(), request)
}

// ImportQAs
// 用于批量导入问答对，最多支持50,000条数据导入。通过此接口，可以将多个问答对一次性导入系统中，以便在后续的知识检索中使用。导入的问答对可以来自外部系统、文件或其他数据源。使用场景：当需要一次性导入大量的问答对时使用。例如，一个公司可能需要将其产品的常见问题解答从一个文档或外部系统批量导入到知识库中，以便用户可以通过知识检索系统进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ImportQAsWithContext(ctx context.Context, request *ImportQAsRequest) (response *ImportQAsResponse, err error) {
    if request == nil {
        request = NewImportQAsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportQAs require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportQAsResponse()
    err = c.Send(request, response)
    return
}

func NewListAttributeLabelsRequest() (request *ListAttributeLabelsRequest) {
    request = &ListAttributeLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ListAttributeLabels")
    
    
    return
}

func NewListAttributeLabelsResponse() (response *ListAttributeLabelsResponse) {
    response = &ListAttributeLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAttributeLabels
// 用于获取所有属性标签的列表。 使用场景：用于查看当前系统中所有已有的属性标签，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListAttributeLabels(request *ListAttributeLabelsRequest) (response *ListAttributeLabelsResponse, err error) {
    return c.ListAttributeLabelsWithContext(context.Background(), request)
}

// ListAttributeLabels
// 用于获取所有属性标签的列表。 使用场景：用于查看当前系统中所有已有的属性标签，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListAttributeLabelsWithContext(ctx context.Context, request *ListAttributeLabelsRequest) (response *ListAttributeLabelsResponse, err error) {
    if request == nil {
        request = NewListAttributeLabelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAttributeLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAttributeLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewListDocsRequest() (request *ListDocsRequest) {
    request = &ListDocsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ListDocs")
    
    
    return
}

func NewListDocsResponse() (response *ListDocsResponse) {
    response = &ListDocsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDocs
// 用于获取所有文档的列表。 使用场景：用于查看当前系统中所有已有的文档，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListDocs(request *ListDocsRequest) (response *ListDocsResponse, err error) {
    return c.ListDocsWithContext(context.Background(), request)
}

// ListDocs
// 用于获取所有文档的列表。 使用场景：用于查看当前系统中所有已有的文档，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListDocsWithContext(ctx context.Context, request *ListDocsRequest) (response *ListDocsResponse, err error) {
    if request == nil {
        request = NewListDocsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDocs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDocsResponse()
    err = c.Send(request, response)
    return
}

func NewListQAsRequest() (request *ListQAsRequest) {
    request = &ListQAsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ListQAs")
    
    
    return
}

func NewListQAsResponse() (response *ListQAsResponse) {
    response = &ListQAsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQAs
// 用于获取所有问答对的列表。 使用场景：用于查看当前系统中所有已有的问答对，方便进行管理和维护。
//
// 可能返回的错误码:
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListQAs(request *ListQAsRequest) (response *ListQAsResponse, err error) {
    return c.ListQAsWithContext(context.Background(), request)
}

// ListQAs
// 用于获取所有问答对的列表。 使用场景：用于查看当前系统中所有已有的问答对，方便进行管理和维护。
//
// 可能返回的错误码:
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListQAsWithContext(ctx context.Context, request *ListQAsRequest) (response *ListQAsResponse, err error) {
    if request == nil {
        request = NewListQAsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQAs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQAsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAttributeLabelRequest() (request *ModifyAttributeLabelRequest) {
    request = &ModifyAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ModifyAttributeLabel")
    
    
    return
}

func NewModifyAttributeLabelResponse() (response *ModifyAttributeLabelResponse) {
    response = &ModifyAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAttributeLabel
// 用于修改已有的属性标签。 使用场景：当需要更改属性标签的名称或描述时使用，比如将“售后”标签改为“售前”。
//
// 可能返回的错误码:
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyAttributeLabel(request *ModifyAttributeLabelRequest) (response *ModifyAttributeLabelResponse, err error) {
    return c.ModifyAttributeLabelWithContext(context.Background(), request)
}

// ModifyAttributeLabel
// 用于修改已有的属性标签。 使用场景：当需要更改属性标签的名称或描述时使用，比如将“售后”标签改为“售前”。
//
// 可能返回的错误码:
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyAttributeLabelWithContext(ctx context.Context, request *ModifyAttributeLabelRequest) (response *ModifyAttributeLabelResponse, err error) {
    if request == nil {
        request = NewModifyAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQARequest() (request *ModifyQARequest) {
    request = &ModifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ModifyQA")
    
    
    return
}

func NewModifyQAResponse() (response *ModifyQAResponse) {
    response = &ModifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQA
// 用于修改已有的问答对。 使用场景：当需要更新问答对的内容或答案时使用。
//
// 可能返回的错误码:
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyQA(request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    return c.ModifyQAWithContext(context.Background(), request)
}

// ModifyQA
// 用于修改已有的问答对。 使用场景：当需要更新问答对的内容或答案时使用。
//
// 可能返回的错误码:
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyQAWithContext(ctx context.Context, request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    if request == nil {
        request = NewModifyQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRewriteRequest() (request *QueryRewriteRequest) {
    request = &QueryRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "QueryRewrite")
    
    
    return
}

func NewQueryRewriteResponse() (response *QueryRewriteResponse) {
    response = &QueryRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryRewrite
// 多轮改写（QueryRewrite）主要用于多轮对话中，进行指代消解和省略补全。使用本接口，无需输入prompt描述，根据对话历史即可生成更精确的用户查询。在应用场景上，本接口可应用于智能问答、对话式搜索等多种场景。
//
// 开通[产品体验](https://lke.cloud.tencent.com/lke/#/trialProduct)后可获得50wtoken体验额度。本接口（QueryRewrite）有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryRewrite(request *QueryRewriteRequest) (response *QueryRewriteResponse, err error) {
    return c.QueryRewriteWithContext(context.Background(), request)
}

// QueryRewrite
// 多轮改写（QueryRewrite）主要用于多轮对话中，进行指代消解和省略补全。使用本接口，无需输入prompt描述，根据对话历史即可生成更精确的用户查询。在应用场景上，本接口可应用于智能问答、对话式搜索等多种场景。
//
// 开通[产品体验](https://lke.cloud.tencent.com/lke/#/trialProduct)后可获得50wtoken体验额度。本接口（QueryRewrite）有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryRewriteWithContext(ctx context.Context, request *QueryRewriteRequest) (response *QueryRewriteResponse, err error) {
    if request == nil {
        request = NewQueryRewriteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryRewrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewReconstructDocumentSSERequest() (request *ReconstructDocumentSSERequest) {
    request = &ReconstructDocumentSSERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ReconstructDocumentSSE")
    
    
    return
}

func NewReconstructDocumentSSEResponse() (response *ReconstructDocumentSSEResponse) {
    response = &ReconstructDocumentSSEResponse{} 
    return

}

// ReconstructDocumentSSE
// 准实时文档解析接口，使用HTTP SSE 协议通信。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReconstructDocumentSSE(request *ReconstructDocumentSSERequest) (response *ReconstructDocumentSSEResponse, err error) {
    return c.ReconstructDocumentSSEWithContext(context.Background(), request)
}

// ReconstructDocumentSSE
// 准实时文档解析接口，使用HTTP SSE 协议通信。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReconstructDocumentSSEWithContext(ctx context.Context, request *ReconstructDocumentSSERequest) (response *ReconstructDocumentSSEResponse, err error) {
    if request == nil {
        request = NewReconstructDocumentSSERequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReconstructDocumentSSE require credential")
    }

    request.SetContext(ctx)
    
    response = NewReconstructDocumentSSEResponse()
    err = c.Send(request, response)
    return
}

func NewRetrieveKnowledgeRequest() (request *RetrieveKnowledgeRequest) {
    request = &RetrieveKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "RetrieveKnowledge")
    
    
    return
}

func NewRetrieveKnowledgeResponse() (response *RetrieveKnowledgeResponse) {
    response = &RetrieveKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetrieveKnowledge
// 用于检索知识库中的文档和问答对内容。 使用场景：适用于查询长期存储在知识库中的文档和问答对，比如产品手册、用户指南等内容的检索。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RetrieveKnowledge(request *RetrieveKnowledgeRequest) (response *RetrieveKnowledgeResponse, err error) {
    return c.RetrieveKnowledgeWithContext(context.Background(), request)
}

// RetrieveKnowledge
// 用于检索知识库中的文档和问答对内容。 使用场景：适用于查询长期存储在知识库中的文档和问答对，比如产品手册、用户指南等内容的检索。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RetrieveKnowledgeWithContext(ctx context.Context, request *RetrieveKnowledgeRequest) (response *RetrieveKnowledgeResponse, err error) {
    if request == nil {
        request = NewRetrieveKnowledgeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetrieveKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetrieveKnowledgeResponse()
    err = c.Send(request, response)
    return
}

func NewRetrieveKnowledgeRealtimeRequest() (request *RetrieveKnowledgeRealtimeRequest) {
    request = &RetrieveKnowledgeRealtimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "RetrieveKnowledgeRealtime")
    
    
    return
}

func NewRetrieveKnowledgeRealtimeResponse() (response *RetrieveKnowledgeRealtimeResponse) {
    response = &RetrieveKnowledgeRealtimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetrieveKnowledgeRealtime
// 用于实时检索在UploadDocRealtime接口上传的实时文档内容。 使用场景：适用于在会话中对文档进行问答的场景
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RetrieveKnowledgeRealtime(request *RetrieveKnowledgeRealtimeRequest) (response *RetrieveKnowledgeRealtimeResponse, err error) {
    return c.RetrieveKnowledgeRealtimeWithContext(context.Background(), request)
}

// RetrieveKnowledgeRealtime
// 用于实时检索在UploadDocRealtime接口上传的实时文档内容。 使用场景：适用于在会话中对文档进行问答的场景
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RetrieveKnowledgeRealtimeWithContext(ctx context.Context, request *RetrieveKnowledgeRealtimeRequest) (response *RetrieveKnowledgeRealtimeResponse, err error) {
    if request == nil {
        request = NewRetrieveKnowledgeRealtimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetrieveKnowledgeRealtime require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetrieveKnowledgeRealtimeResponse()
    err = c.Send(request, response)
    return
}

func NewRunRerankRequest() (request *RunRerankRequest) {
    request = &RunRerankRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "RunRerank")
    
    
    return
}

func NewRunRerankResponse() (response *RunRerankResponse) {
    response = &RunRerankResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunRerank
// 基于知识引擎精调模型技术的rerank模型，支持对多路召回的结果进行重排序，根据query与切片内容的相关性，按分数由高到低对切片进行排序，并输出对应的打分结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunRerank(request *RunRerankRequest) (response *RunRerankResponse, err error) {
    return c.RunRerankWithContext(context.Background(), request)
}

// RunRerank
// 基于知识引擎精调模型技术的rerank模型，支持对多路召回的结果进行重排序，根据query与切片内容的相关性，按分数由高到低对切片进行排序，并输出对应的打分结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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

func NewUploadDocRequest() (request *UploadDocRequest) {
    request = &UploadDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "UploadDoc")
    
    
    return
}

func NewUploadDocResponse() (response *UploadDocResponse) {
    response = &UploadDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadDoc
// 用于上传文档内容。上传的文档将存储在知识库中，可以通过RetrieveKnowledge[知识库内容检索接口](https://cloud.tencent.com/document/product/1772/115349)进行检索。 
//
// 使用场景：适用于需要长期存储和检索的文档内容，如产品手册、用户指南等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UploadDoc(request *UploadDocRequest) (response *UploadDocResponse, err error) {
    return c.UploadDocWithContext(context.Background(), request)
}

// UploadDoc
// 用于上传文档内容。上传的文档将存储在知识库中，可以通过RetrieveKnowledge[知识库内容检索接口](https://cloud.tencent.com/document/product/1772/115349)进行检索。 
//
// 使用场景：适用于需要长期存储和检索的文档内容，如产品手册、用户指南等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UploadDocWithContext(ctx context.Context, request *UploadDocRequest) (response *UploadDocResponse, err error) {
    if request == nil {
        request = NewUploadDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadDocResponse()
    err = c.Send(request, response)
    return
}
