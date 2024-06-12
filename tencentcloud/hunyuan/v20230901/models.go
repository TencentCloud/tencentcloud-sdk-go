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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ChatCompletionsRequestParams struct {
	// 模型名称，可选值包括 hunyuan-lite、hunyuan-standard、hunyuan-standard-256K、hunyuan-pro。
	// 各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。
	// 
	// 注意：
	// 不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message.Role 可选值：system、user、assistant。
	// 其中，system 角色可选，如存在则必须位于列表的最开始。user 和 assistant 需交替出现（一问一答），以 user 提问开始和结束，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user assistant user ...]。
	// 3. Messages 中 Content 总长度不能超过模型输入长度上限（可参考 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 文档），超过则会截断最前面的内容，只保留尾部内容。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 流式输出审核开关。
	// 说明：
	// 1. 当使用流式输出（Stream 字段值为 true）时，该字段生效。
	// 2. 输出审核有流式和同步两种模式，**流式模式首包响应更快**。未传值时默认为流式模式（true）。
	// 3. 如果值为 true，将对输出内容进行分段审核，审核通过的内容流式输出返回。如果出现审核不过，响应中的 FinishReason 值为 sensitive。
	// 4. 如果值为 false，则不使用流式输出审核，需要审核完所有输出内容后再返回结果。
	// 
	// 注意：
	// 当选择流式输出审核时，可能会出现部分内容已输出，但中间某一段响应中的 FinishReason 值为 sensitive，此时说明安全审核未通过。如果业务场景有实时文字上屏的需求，需要自行撤回已上屏的内容，并建议自定义替换为一条提示语，如 “这个问题我不方便回答，不如我们换个话题试试”，以保障终端体验。
	StreamModeration *bool `json:"StreamModeration,omitnil,omitempty" name:"StreamModeration"`

	// 说明：
	// 1. 影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2. 默认 1.0，取值区间为 [0.0, 1.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2. 默认 1.0，取值区间为 [0.0, 2.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 功能增强（如搜索）开关。
	// 说明：
	// 1. hunyuan-lite 无功能增强（如搜索）能力，该参数对 hunyuan-lite 版本不生效。
	// 2. 未传值时默认打开开关。
	// 3. 关闭时将直接由主模型生成回复内容，可以降低响应时延（对于流式输出时的首字时延尤为明显）。但在少数场景里，回复效果可能会下降。
	// 4. 安全审核能力不属于功能增强范围，不受此字段影响。
	EnableEnhancement *bool `json:"EnableEnhancement,omitnil,omitempty" name:"EnableEnhancement"`
}

type ChatCompletionsRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称，可选值包括 hunyuan-lite、hunyuan-standard、hunyuan-standard-256K、hunyuan-pro。
	// 各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。
	// 
	// 注意：
	// 不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message.Role 可选值：system、user、assistant。
	// 其中，system 角色可选，如存在则必须位于列表的最开始。user 和 assistant 需交替出现（一问一答），以 user 提问开始和结束，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user assistant user ...]。
	// 3. Messages 中 Content 总长度不能超过模型输入长度上限（可参考 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 文档），超过则会截断最前面的内容，只保留尾部内容。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 流式输出审核开关。
	// 说明：
	// 1. 当使用流式输出（Stream 字段值为 true）时，该字段生效。
	// 2. 输出审核有流式和同步两种模式，**流式模式首包响应更快**。未传值时默认为流式模式（true）。
	// 3. 如果值为 true，将对输出内容进行分段审核，审核通过的内容流式输出返回。如果出现审核不过，响应中的 FinishReason 值为 sensitive。
	// 4. 如果值为 false，则不使用流式输出审核，需要审核完所有输出内容后再返回结果。
	// 
	// 注意：
	// 当选择流式输出审核时，可能会出现部分内容已输出，但中间某一段响应中的 FinishReason 值为 sensitive，此时说明安全审核未通过。如果业务场景有实时文字上屏的需求，需要自行撤回已上屏的内容，并建议自定义替换为一条提示语，如 “这个问题我不方便回答，不如我们换个话题试试”，以保障终端体验。
	StreamModeration *bool `json:"StreamModeration,omitnil,omitempty" name:"StreamModeration"`

	// 说明：
	// 1. 影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2. 默认 1.0，取值区间为 [0.0, 1.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2. 默认 1.0，取值区间为 [0.0, 2.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 功能增强（如搜索）开关。
	// 说明：
	// 1. hunyuan-lite 无功能增强（如搜索）能力，该参数对 hunyuan-lite 版本不生效。
	// 2. 未传值时默认打开开关。
	// 3. 关闭时将直接由主模型生成回复内容，可以降低响应时延（对于流式输出时的首字时延尤为明显）。但在少数场景里，回复效果可能会下降。
	// 4. 安全审核能力不属于功能增强范围，不受此字段影响。
	EnableEnhancement *bool `json:"EnableEnhancement,omitnil,omitempty" name:"EnableEnhancement"`
}

func (r *ChatCompletionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Messages")
	delete(f, "Stream")
	delete(f, "StreamModeration")
	delete(f, "TopP")
	delete(f, "Temperature")
	delete(f, "EnableEnhancement")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatCompletionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatCompletionsResponseParams struct {
	// Unix 时间戳，单位为秒。
	Created *int64 `json:"Created,omitnil,omitempty" name:"Created"`

	// Token 统计信息。
	// 按照总 Token 数量计费。
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 免责声明。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 本轮对话的 ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 回复内容。
	Choices []*Choice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// 错误信息。
	// 如果流式返回中服务处理异常，返回该错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *ErrorMsg `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChatCompletionsResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ChatCompletionsResponseParams `json:"Response"`
}

func (r *ChatCompletionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Choice struct {
	// 结束标志位，可能为 stop 或 sensitive。
	// stop 表示输出正常结束，sensitive 只在开启流式输出审核时会出现，表示安全审核未通过。
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 增量返回值，流式调用时使用该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Delta *Delta `json:"Delta,omitnil,omitempty" name:"Delta"`

	// 返回值，非流式调用时使用该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *Message `json:"Message,omitnil,omitempty" name:"Message"`
}

type Delta struct {
	// 角色名称。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容详情。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type EmbeddingData struct {
	// Embedding 信息，目前为 1024 维浮点数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Embedding []*float64 `json:"Embedding,omitnil,omitempty" name:"Embedding"`

	// 下标，目前不支持批量，因此固定为 0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 目前固定为 "embedding"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`
}

type EmbeddingUsage struct {
	// 输入 Token 数。
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 总 Token 数。
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type ErrorMsg struct {
	// 错误提示信息。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 错误码。
	// 4000 服务内部异常。
	// 4001 请求模型超时。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`
}

// Predefined struct for user
type GetEmbeddingRequestParams struct {
	// 输入文本。总长度不超过 1024 个 Token，超过则会截断最后面的内容。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`
}

type GetEmbeddingRequest struct {
	*tchttp.BaseRequest
	
	// 输入文本。总长度不超过 1024 个 Token，超过则会截断最后面的内容。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`
}

func (r *GetEmbeddingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmbeddingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Input")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEmbeddingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmbeddingResponseParams struct {
	// 返回的 Embedding 信息。当前不支持批量，所以数组元素数目为 1。
	Data []*EmbeddingData `json:"Data,omitnil,omitempty" name:"Data"`

	// Token 使用计数，按照总 Token 数量收费。
	Usage *EmbeddingUsage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEmbeddingResponse struct {
	*tchttp.BaseResponse
	Response *GetEmbeddingResponseParams `json:"Response"`
}

func (r *GetEmbeddingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmbeddingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTokenCountRequestParams struct {
	// 输入文本
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`
}

type GetTokenCountRequest struct {
	*tchttp.BaseRequest
	
	// 输入文本
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`
}

func (r *GetTokenCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTokenCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTokenCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTokenCountResponseParams struct {
	// token计数
	TokenCount *int64 `json:"TokenCount,omitnil,omitempty" name:"TokenCount"`

	// 字符计数
	CharacterCount *int64 `json:"CharacterCount,omitnil,omitempty" name:"CharacterCount"`

	// 切分后的列表
	Tokens []*string `json:"Tokens,omitnil,omitempty" name:"Tokens"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTokenCountResponse struct {
	*tchttp.BaseResponse
	Response *GetTokenCountResponseParams `json:"Response"`
}

func (r *GetTokenCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTokenCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Message struct {
	// 角色，可选值包括 system、user、assistant。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 文本内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type QueryHunyuanImageJobRequestParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanImageJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanImageJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanImageJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanImageJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanImageJobResponseParams struct {
	// 当前任务状态码：
	// 1：等待中、2：运行中、4：处理失败、5：处理完成。
	JobStatusCode *string `json:"JobStatusCode,omitnil,omitempty" name:"JobStatusCode"`

	// 当前任务状态：排队中、处理中、处理失败或者处理完成。
	JobStatusMsg *string `json:"JobStatusMsg,omitnil,omitempty" name:"JobStatusMsg"`

	// 任务处理失败错误码。
	JobErrorCode *string `json:"JobErrorCode,omitnil,omitempty" name:"JobErrorCode"`

	// 任务处理失败错误信息。
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`

	// 生成图 URL 列表，有效期1小时，请及时保存。
	ResultImage []*string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 结果 detail 数组，Success 代表成功。
	ResultDetails []*string `json:"ResultDetails,omitnil,omitempty" name:"ResultDetails"`

	// 对应 SubmitTextToImageProJob 接口中 Revise 参数。开启扩写时，返回扩写后的 prompt 文本。 如果关闭扩写，将直接返回原始输入的 prompt。
	RevisedPrompt []*string `json:"RevisedPrompt,omitnil,omitempty" name:"RevisedPrompt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanImageJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanImageJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanImageJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanImageJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanImageJobRequestParams struct {
	// 文本描述。 算法将根据输入的文本智能生成与之相关的图像。 不能为空，推荐使用中文。最多可传100个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 绘画风格。
	// 请在 [混元生图风格列表](https://cloud.tencent.com/document/product/1729/105846) 中选择期望的风格，传入风格编号。
	// 不传默认不指定风格。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3），不传默认使用1024:1024。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// prompt 扩写开关。1为开启，0为关闭，不传默认开启。
	// 开启扩写后，将自动扩写原始输入的 prompt 并使用扩写后的 prompt 生成图片，返回生成图片结果时将一并返回扩写后的 prompt 文本。
	// 如果关闭扩写，将直接使用原始输入的 prompt 生成图片。
	// 建议开启，在多数场景下可提升生成图片效果、丰富生成图片细节。
	Revise *int64 `json:"Revise,omitnil,omitempty" name:"Revise"`
}

type SubmitHunyuanImageJobRequest struct {
	*tchttp.BaseRequest
	
	// 文本描述。 算法将根据输入的文本智能生成与之相关的图像。 不能为空，推荐使用中文。最多可传100个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 绘画风格。
	// 请在 [混元生图风格列表](https://cloud.tencent.com/document/product/1729/105846) 中选择期望的风格，传入风格编号。
	// 不传默认不指定风格。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3），不传默认使用1024:1024。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// prompt 扩写开关。1为开启，0为关闭，不传默认开启。
	// 开启扩写后，将自动扩写原始输入的 prompt 并使用扩写后的 prompt 生成图片，返回生成图片结果时将一并返回扩写后的 prompt 文本。
	// 如果关闭扩写，将直接使用原始输入的 prompt 生成图片。
	// 建议开启，在多数场景下可提升生成图片效果、丰富生成图片细节。
	Revise *int64 `json:"Revise,omitnil,omitempty" name:"Revise"`
}

func (r *SubmitHunyuanImageJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanImageJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "Style")
	delete(f, "Resolution")
	delete(f, "LogoAdd")
	delete(f, "Revise")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanImageJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanImageJobResponseParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanImageJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanImageJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanImageJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanImageJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Usage struct {
	// 输入 Token 数量。
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 输出 Token 数量。
	CompletionTokens *int64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// 总 Token 数量。
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}