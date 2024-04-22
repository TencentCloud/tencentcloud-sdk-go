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
type ChatProRequestParams struct {
	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message 的 Role 当前可选值：system、user、assistant。其中，system 角色是可选的，如果存在，必须位于列表的最开始。此外，user 和 assistant 需交替出现（一问一答），以 user 提问开始和结束，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user assistant user ...]。
	// 3. Messages 中 Content 总长度不超过 16000 Token，超过则会截断最前面的内容，只保留尾部内容。建议不超过 4000 Token。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 说明：
	// 1. 影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2. 默认 1.0，取值区间为 [0.0, 1.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2. 默认 1.0，取值区间为 [0.0，2.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为流式调用。
	// 2. 流式调用时以 SSE 协议增量返回结果。
	// 3. 非流式调用时接口响应耗时较长，非必要不建议使用。
	// 4. 非流式调用时只返回一次最终结果，调用方式与普通 HTTP 请求无异。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 流式输出审核开关。
	// 说明：
	// 1. 当 Stream 字段值为 true 时，该字段有效。
	// 2. 未传值时默认不使用流式输出审核。
	// 3. 如果值为 true，将对输出内容进行分段审核，审核通过的内容流式输出返回。如果出现审核不过，响应中的 finish_reason 值为 sensitive。
	// 4. 如果值为 false，则需要审核完所有输出内容后再返回结果。
	// 
	// 注意：
	// 当选择流式输出审核时，可能会出现部分内容已输出，但中间某一段响应中的 finish_reason 值为 sensitive，此时说明安全审核未通过。如果业务场景有实时文字上屏的需求，需要自行撤回已上屏的内容，并建议自定义替换为一条提示语，如 “这个问题我不方便回答，不如我们换个话题试试”，以保障终端体验。
	StreamModeration *bool `json:"StreamModeration,omitnil,omitempty" name:"StreamModeration"`
}

type ChatProRequest struct {
	*tchttp.BaseRequest
	
	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message 的 Role 当前可选值：system、user、assistant。其中，system 角色是可选的，如果存在，必须位于列表的最开始。此外，user 和 assistant 需交替出现（一问一答），以 user 提问开始和结束，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user assistant user ...]。
	// 3. Messages 中 Content 总长度不超过 16000 Token，超过则会截断最前面的内容，只保留尾部内容。建议不超过 4000 Token。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 说明：
	// 1. 影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2. 默认 1.0，取值区间为 [0.0, 1.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2. 默认 1.0，取值区间为 [0.0，2.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为流式调用。
	// 2. 流式调用时以 SSE 协议增量返回结果。
	// 3. 非流式调用时接口响应耗时较长，非必要不建议使用。
	// 4. 非流式调用时只返回一次最终结果，调用方式与普通 HTTP 请求无异。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 流式输出审核开关。
	// 说明：
	// 1. 当 Stream 字段值为 true 时，该字段有效。
	// 2. 未传值时默认不使用流式输出审核。
	// 3. 如果值为 true，将对输出内容进行分段审核，审核通过的内容流式输出返回。如果出现审核不过，响应中的 finish_reason 值为 sensitive。
	// 4. 如果值为 false，则需要审核完所有输出内容后再返回结果。
	// 
	// 注意：
	// 当选择流式输出审核时，可能会出现部分内容已输出，但中间某一段响应中的 finish_reason 值为 sensitive，此时说明安全审核未通过。如果业务场景有实时文字上屏的需求，需要自行撤回已上屏的内容，并建议自定义替换为一条提示语，如 “这个问题我不方便回答，不如我们换个话题试试”，以保障终端体验。
	StreamModeration *bool `json:"StreamModeration,omitnil,omitempty" name:"StreamModeration"`
}

func (r *ChatProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Messages")
	delete(f, "TopP")
	delete(f, "Temperature")
	delete(f, "Stream")
	delete(f, "StreamModeration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatProResponseParams struct {
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

type ChatProResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ChatProResponseParams `json:"Response"`
}

func (r *ChatProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatStdRequestParams struct {
	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message 的 Role 当前可选值：user、assistant。其中，user 和 assistant 需交替出现（一问一答），以 user 提问开始和结束，且 Content 不能为空。Role 的顺序示例：[user assistant user assistant user ...]。
	// 3. Messages 中 Content 总长度不超过 16000 Token，超过则会截断最前面的内容，只保留尾部内容。建议不超过 4000 Token。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 说明：
	// 1. 影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2. 默认 1.0，取值区间为 [0.0, 1.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2. 默认 1.0，取值区间为 [0.0，2.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为流式调用。
	// 2. 流式调用时以 SSE 协议增量返回结果。
	// 3. 非流式调用时接口响应耗时较长，非必要不建议使用。
	// 4. 非流式调用时只返回一次最终结果，调用方式与普通 HTTP 请求无异。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 流式输出审核开关。
	// 说明：
	// 1. 当 Stream 字段值为 true 时，该字段有效。
	// 2. 未传值时默认不使用流式输出审核。
	// 3. 如果值为 true，将对输出内容进行分段审核，审核通过的内容流式输出返回。如果出现审核不过，响应中的 finish_reason 值为 sensitive。
	// 4. 如果值为 false，则需要审核完所有输出内容后再返回结果。
	// 
	// 注意：
	// 当选择流式输出审核时，可能会出现部分内容已输出，但中间某一段响应中的 finish_reason 值为 sensitive，此时说明安全审核未通过。如果业务场景有实时文字上屏的需求，需要自行撤回已上屏的内容，并建议自定义替换为一条提示语，如 “这个问题我不方便回答，不如我们换个话题试试”，以保障终端体验。
	StreamModeration *bool `json:"StreamModeration,omitnil,omitempty" name:"StreamModeration"`
}

type ChatStdRequest struct {
	*tchttp.BaseRequest
	
	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message 的 Role 当前可选值：user、assistant。其中，user 和 assistant 需交替出现（一问一答），以 user 提问开始和结束，且 Content 不能为空。Role 的顺序示例：[user assistant user assistant user ...]。
	// 3. Messages 中 Content 总长度不超过 16000 Token，超过则会截断最前面的内容，只保留尾部内容。建议不超过 4000 Token。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 说明：
	// 1. 影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2. 默认 1.0，取值区间为 [0.0, 1.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2. 默认 1.0，取值区间为 [0.0，2.0]。
	// 3. 非必要不建议使用，不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为流式调用。
	// 2. 流式调用时以 SSE 协议增量返回结果。
	// 3. 非流式调用时接口响应耗时较长，非必要不建议使用。
	// 4. 非流式调用时只返回一次最终结果，调用方式与普通 HTTP 请求无异。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 流式输出审核开关。
	// 说明：
	// 1. 当 Stream 字段值为 true 时，该字段有效。
	// 2. 未传值时默认不使用流式输出审核。
	// 3. 如果值为 true，将对输出内容进行分段审核，审核通过的内容流式输出返回。如果出现审核不过，响应中的 finish_reason 值为 sensitive。
	// 4. 如果值为 false，则需要审核完所有输出内容后再返回结果。
	// 
	// 注意：
	// 当选择流式输出审核时，可能会出现部分内容已输出，但中间某一段响应中的 finish_reason 值为 sensitive，此时说明安全审核未通过。如果业务场景有实时文字上屏的需求，需要自行撤回已上屏的内容，并建议自定义替换为一条提示语，如 “这个问题我不方便回答，不如我们换个话题试试”，以保障终端体验。
	StreamModeration *bool `json:"StreamModeration,omitnil,omitempty" name:"StreamModeration"`
}

func (r *ChatStdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatStdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Messages")
	delete(f, "TopP")
	delete(f, "Temperature")
	delete(f, "Stream")
	delete(f, "StreamModeration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatStdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatStdResponseParams struct {
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

type ChatStdResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ChatStdResponseParams `json:"Response"`
}

func (r *ChatStdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatStdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Choice struct {
	// 结束标志位，为 stop 则表示尾包。
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
	// embedding 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Embedding []*float64 `json:"Embedding,omitnil,omitempty" name:"Embedding"`

	// 下标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// embedding
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
	// 输入文本。总长度不超过1024 个token, 超过则会截断最后面的内容。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`
}

type GetEmbeddingRequest struct {
	*tchttp.BaseRequest
	
	// 输入文本。总长度不超过1024 个token, 超过则会截断最后面的内容。
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
	// 返回的 embedding 信息。
	Data []*EmbeddingData `json:"Data,omitnil,omitempty" name:"Data"`

	// token 使用计数，按照总token数量收费。
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
	// 角色
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 消息内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type Usage struct {
	// 输入 Token 数量。
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 输出 Token 数量。
	CompletionTokens *int64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// 总 Token 数量。
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}