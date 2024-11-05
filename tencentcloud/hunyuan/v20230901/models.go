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
type ActivateServiceRequestParams struct {
	// 开通之后，是否关闭后付费；默认为0，不关闭；1为关闭
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type ActivateServiceRequest struct {
	*tchttp.BaseRequest
	
	// 开通之后，是否关闭后付费；默认为0，不关闭；1为关闭
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

func (r *ActivateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActivateServiceResponse struct {
	*tchttp.BaseResponse
	Response *ActivateServiceResponseParams `json:"Response"`
}

func (r *ActivateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatCompletionsRequestParams struct {
	// 模型名称，可选值包括 hunyuan-lite、hunyuan-standard、hunyuan-standard-256K、hunyuan-pro、 hunyuan-code、 hunyuan-role、 hunyuan-functioncall、 hunyuan-vision、 hunyuan-turbo、 hunyuan-turbo-latest。
	// 各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。
	// 
	// 注意：
	// 不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message.Role 可选值：system、user、assistant、 tool（functioncall场景）。
	// 其中，system 角色可选，如存在则必须位于列表的最开始。user（tool） 和 assistant 需交替出现（一问一答），以 user 提问开始，user（tool）提问结束，其中tool可以连续出现多次，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user（tool tool ...） assistant user（tool tool ...） ...]。
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
	// 1. 影响输出文本的多样性。模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。
	// 2. 取值区间为 [0.0, 1.0]。取值越大，生成文本的多样性越强。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 影响模型输出多样性，模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。
	// 2. 取值区间为 [0.0, 2.0]。较高的数值会使输出更加多样化和不可预测，而较低的数值会使其更加集中和确定。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 功能增强（如搜索）开关。
	// 说明：
	// 1. hunyuan-lite 无功能增强（如搜索）能力，该参数对 hunyuan-lite 版本不生效。
	// 2. 未传值时默认打开开关。
	// 3. 关闭时将直接由主模型生成回复内容，可以降低响应时延（对于流式输出时的首字时延尤为明显）。但在少数场景里，回复效果可能会下降。
	// 4. 安全审核能力不属于功能增强范围，不受此字段影响。
	EnableEnhancement *bool `json:"EnableEnhancement,omitnil,omitempty" name:"EnableEnhancement"`

	// 可调用的工具列表，仅对 hunyuan-pro、hunyuan-turbo、hunyuan-functioncall 模型生效。
	Tools []*Tool `json:"Tools,omitnil,omitempty" name:"Tools"`

	// 工具使用选项，可选值包括 none、auto、custom。
	// 说明：
	// 1. 仅对 hunyuan-pro、hunyuan-turbo、hunyuan-functioncall 模型生效。
	// 2. none：不调用工具；auto：模型自行选择生成回复或调用工具；custom：强制模型调用指定的工具。
	// 3. 未设置时，默认值为auto
	ToolChoice *string `json:"ToolChoice,omitnil,omitempty" name:"ToolChoice"`

	// 强制模型调用指定的工具，当参数ToolChoice为custom时，此参数为必填
	CustomTool *Tool `json:"CustomTool,omitnil,omitempty" name:"CustomTool"`

	// 默认是false，在值为true且命中搜索时，接口会返回SearchInfo
	SearchInfo *bool `json:"SearchInfo,omitnil,omitempty" name:"SearchInfo"`

	// 搜索引文角标开关。
	// 说明：
	// 1. 配合EnableEnhancement和SearchInfo参数使用。打开后，回答中命中搜索的结果会在片段后增加角标标志，对应SearchInfo列表中的链接。
	// 2. false：开关关闭，true：开关打开。
	// 3. 未传值时默认开关关闭（false）。
	Citation *bool `json:"Citation,omitnil,omitempty" name:"Citation"`

	// 是否开启极速版搜索，默认false，不开启；在开启且命中搜索时，会启用极速版搜索，流式输出首字返回更快。
	EnableSpeedSearch *bool `json:"EnableSpeedSearch,omitnil,omitempty" name:"EnableSpeedSearch"`

	// 多媒体开关。
	// 详细介绍请阅读 [多媒体介绍](https://cloud.tencent.com/document/product/1729/111178) 中的说明。
	// 说明：
	// 1. 该参数目前仅对白名单内用户生效，如您想体验该功能请 [联系我们](https://cloud.tencent.com/act/event/Online_service)。
	// 2. 该参数仅在功能增强（如搜索）开关开启（EnableEnhancement=true）并且极速版搜索开关关闭（EnableSpeedSearch=false）时生效。
	// 3. hunyuan-lite 无多媒体能力，该参数对 hunyuan-lite 版本不生效。
	// 4. 未传值时默认关闭。
	// 5. 开启并搜索到对应的多媒体信息时，会输出对应的多媒体地址，可以定制个性化的图文消息。
	EnableMultimedia *bool `json:"EnableMultimedia,omitnil,omitempty" name:"EnableMultimedia"`

	// 是否开启搜索深度模式，默认是false，在值为true且命中搜索时，会请求深度搜索。
	EnableDeepSearch *bool `json:"EnableDeepSearch,omitnil,omitempty" name:"EnableDeepSearch"`

	// 说明： 1. 确保模型的输出是可复现的。 2. 取值区间为非0正整数，最大值10000。 3. 非必要不建议使用，不合理的取值会影响效果。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`
}

type ChatCompletionsRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称，可选值包括 hunyuan-lite、hunyuan-standard、hunyuan-standard-256K、hunyuan-pro、 hunyuan-code、 hunyuan-role、 hunyuan-functioncall、 hunyuan-vision、 hunyuan-turbo、 hunyuan-turbo-latest。
	// 各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。
	// 
	// 注意：
	// 不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message.Role 可选值：system、user、assistant、 tool（functioncall场景）。
	// 其中，system 角色可选，如存在则必须位于列表的最开始。user（tool） 和 assistant 需交替出现（一问一答），以 user 提问开始，user（tool）提问结束，其中tool可以连续出现多次，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user（tool tool ...） assistant user（tool tool ...） ...]。
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
	// 1. 影响输出文本的多样性。模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。
	// 2. 取值区间为 [0.0, 1.0]。取值越大，生成文本的多样性越强。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 影响模型输出多样性，模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。
	// 2. 取值区间为 [0.0, 2.0]。较高的数值会使输出更加多样化和不可预测，而较低的数值会使其更加集中和确定。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 功能增强（如搜索）开关。
	// 说明：
	// 1. hunyuan-lite 无功能增强（如搜索）能力，该参数对 hunyuan-lite 版本不生效。
	// 2. 未传值时默认打开开关。
	// 3. 关闭时将直接由主模型生成回复内容，可以降低响应时延（对于流式输出时的首字时延尤为明显）。但在少数场景里，回复效果可能会下降。
	// 4. 安全审核能力不属于功能增强范围，不受此字段影响。
	EnableEnhancement *bool `json:"EnableEnhancement,omitnil,omitempty" name:"EnableEnhancement"`

	// 可调用的工具列表，仅对 hunyuan-pro、hunyuan-turbo、hunyuan-functioncall 模型生效。
	Tools []*Tool `json:"Tools,omitnil,omitempty" name:"Tools"`

	// 工具使用选项，可选值包括 none、auto、custom。
	// 说明：
	// 1. 仅对 hunyuan-pro、hunyuan-turbo、hunyuan-functioncall 模型生效。
	// 2. none：不调用工具；auto：模型自行选择生成回复或调用工具；custom：强制模型调用指定的工具。
	// 3. 未设置时，默认值为auto
	ToolChoice *string `json:"ToolChoice,omitnil,omitempty" name:"ToolChoice"`

	// 强制模型调用指定的工具，当参数ToolChoice为custom时，此参数为必填
	CustomTool *Tool `json:"CustomTool,omitnil,omitempty" name:"CustomTool"`

	// 默认是false，在值为true且命中搜索时，接口会返回SearchInfo
	SearchInfo *bool `json:"SearchInfo,omitnil,omitempty" name:"SearchInfo"`

	// 搜索引文角标开关。
	// 说明：
	// 1. 配合EnableEnhancement和SearchInfo参数使用。打开后，回答中命中搜索的结果会在片段后增加角标标志，对应SearchInfo列表中的链接。
	// 2. false：开关关闭，true：开关打开。
	// 3. 未传值时默认开关关闭（false）。
	Citation *bool `json:"Citation,omitnil,omitempty" name:"Citation"`

	// 是否开启极速版搜索，默认false，不开启；在开启且命中搜索时，会启用极速版搜索，流式输出首字返回更快。
	EnableSpeedSearch *bool `json:"EnableSpeedSearch,omitnil,omitempty" name:"EnableSpeedSearch"`

	// 多媒体开关。
	// 详细介绍请阅读 [多媒体介绍](https://cloud.tencent.com/document/product/1729/111178) 中的说明。
	// 说明：
	// 1. 该参数目前仅对白名单内用户生效，如您想体验该功能请 [联系我们](https://cloud.tencent.com/act/event/Online_service)。
	// 2. 该参数仅在功能增强（如搜索）开关开启（EnableEnhancement=true）并且极速版搜索开关关闭（EnableSpeedSearch=false）时生效。
	// 3. hunyuan-lite 无多媒体能力，该参数对 hunyuan-lite 版本不生效。
	// 4. 未传值时默认关闭。
	// 5. 开启并搜索到对应的多媒体信息时，会输出对应的多媒体地址，可以定制个性化的图文消息。
	EnableMultimedia *bool `json:"EnableMultimedia,omitnil,omitempty" name:"EnableMultimedia"`

	// 是否开启搜索深度模式，默认是false，在值为true且命中搜索时，会请求深度搜索。
	EnableDeepSearch *bool `json:"EnableDeepSearch,omitnil,omitempty" name:"EnableDeepSearch"`

	// 说明： 1. 确保模型的输出是可复现的。 2. 取值区间为非0正整数，最大值10000。 3. 非必要不建议使用，不合理的取值会影响效果。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`
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
	delete(f, "Tools")
	delete(f, "ToolChoice")
	delete(f, "CustomTool")
	delete(f, "SearchInfo")
	delete(f, "Citation")
	delete(f, "EnableSpeedSearch")
	delete(f, "EnableMultimedia")
	delete(f, "EnableDeepSearch")
	delete(f, "Seed")
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

	// 本次请求的 RequestId。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 回复内容。
	Choices []*Choice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// 错误信息。
	// 如果流式返回中服务处理异常，返回该错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *ErrorMsg `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 多轮会话风险审核，值为1时，表明存在信息安全风险，建议终止客户多轮会话。
	ModerationLevel *string `json:"ModerationLevel,omitnil,omitempty" name:"ModerationLevel"`

	// 搜索结果信息
	SearchInfo *SearchInfo `json:"SearchInfo,omitnil,omitempty" name:"SearchInfo"`

	// 多媒体信息。
	// 说明：
	// 1. 可以用多媒体信息替换回复内容里的占位符，得到完整的消息。
	// 2. 可能会出现回复内容里存在占位符，但是因为审核等原因没有返回多媒体信息。
	Replaces []*Replace `json:"Replaces,omitnil,omitempty" name:"Replaces"`

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
	// 结束标志位，可能为 stop、 sensitive或者tool_calls。
	// stop 表示输出正常结束。
	// sensitive 只在开启流式输出审核时会出现，表示安全审核未通过。
	// tool_calls 标识函数调用。
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 增量返回值，流式调用时使用该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Delta *Delta `json:"Delta,omitnil,omitempty" name:"Delta"`

	// 返回值，非流式调用时使用该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *Message `json:"Message,omitnil,omitempty" name:"Message"`
}

type Content struct {
	// 内容类型
	// 注意：
	// 当前只支持传入单张图片，传入多张图片时，以第一个图片为准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当 Type 为 text 时使用，表示具体的文本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 图片的url，当 Type 为 image_url 时使用，表示具体的图片内容
	// 如"https://example.com/1.png" 或 图片的base64（注意 "data:image/jpeg;base64," 为必要部分）："data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAA......"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *ImageUrl `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type Delta struct {
	// 角色名称。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容详情。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 模型生成的工具调用，仅 hunyuan-functioncall 模型支持
	// 说明：
	// 对于每一次的输出值应该以Id为标识对Type、Name、Arguments字段进行合并。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolCalls []*ToolCall `json:"ToolCalls,omitnil,omitempty" name:"ToolCalls"`
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

type History struct {
	// 对话的 ID，用于唯一标识一轮对话
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 原始输入的 Prompt 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 扩写后的 Prompt 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	RevisedPrompt *string `json:"RevisedPrompt,omitnil,omitempty" name:"RevisedPrompt"`

	// 生成图的随机种子
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`
}

type ImageUrl struct {
	// 图片的 Url（以 http:// 或 https:// 开头）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type LogoParam struct {
	// 水印url
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// 水印base64，url和base64二选一传入
	LogoImage *string `json:"LogoImage,omitnil,omitempty" name:"LogoImage"`

	// 水印图片位于融合结果图中的坐标，将按照坐标对标识图片进行位置和大小的拉伸匹配
	LogoRect *LogoRect `json:"LogoRect,omitnil,omitempty" name:"LogoRect"`
}

type LogoRect struct {
	// 左上角X坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 左上角Y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 方框宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 方框高度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type Message struct {
	// 角色，可选值包括 system、user、assistant、 tool。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 文本内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 多种类型内容（目前支持图片和文本），仅 hunyuan-vision 模型支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contents []*Content `json:"Contents,omitnil,omitempty" name:"Contents"`

	// 当role为tool时传入，标识具体的函数调用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolCallId *string `json:"ToolCallId,omitnil,omitempty" name:"ToolCallId"`

	// 模型生成的工具调用，仅 hunyuan-pro 或者 hunyuan-functioncall 模型支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolCalls []*ToolCall `json:"ToolCalls,omitnil,omitempty" name:"ToolCalls"`
}

type Multimedia struct {
	// 多媒体类型，可选值包括 image、music、album、playlist。
	// 说明：
	// 1. image：图片；music：单曲，类型为单曲时，会返回详细歌手和歌曲信息；album：专辑；playlist：歌单。
	// 2. 当 type 为 music、album、playlist 时，需要配合 [QQ音乐SDK](https://developer.y.qq.com/) 使用。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 多媒体地址。
	// 说明：
	// 1. type 为 image 时，地址为图片的预览地址；其他类型时，地址为封面图地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 多媒体详情地址。
	// 说明：
	// 1. 仅 type 为 image 时，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 名称。
	// 说明：
	// 1. type 为 image 时，该字段为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 歌手名称。
	// 说明：
	// 1. 仅 type 为 music 时，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Singer *string `json:"Singer,omitnil,omitempty" name:"Singer"`

	// 歌曲详情。
	// 说明：
	// 1. 仅 type 为 music 时，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ext *SongExt `json:"Ext,omitnil,omitempty" name:"Ext"`
}

// Predefined struct for user
type QueryHunyuanImageChatJobRequestParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanImageChatJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanImageChatJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanImageChatJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanImageChatJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanImageChatJobResponseParams struct {
	// 当前任务状态码：
	// 1：等待中、2：运行中、4：处理失败、5：处理完成。
	JobStatusCode *string `json:"JobStatusCode,omitnil,omitempty" name:"JobStatusCode"`

	// 当前任务状态：排队中、处理中、处理失败或者处理完成。
	JobStatusMsg *string `json:"JobStatusMsg,omitnil,omitempty" name:"JobStatusMsg"`

	// 任务处理失败错误码。
	JobErrorCode *string `json:"JobErrorCode,omitnil,omitempty" name:"JobErrorCode"`

	// 任务处理失败错误信息。
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`

	// 本轮对话的 ChatId，ChatId 用于唯一标识一轮对话。
	// 一个对话组中，最多支持进行100轮对话。
	// 每轮对话数据有效期为7天，到期后 ChatId 失效，有效期内的历史对话数据可通过 History 查询，如有长期使用需求请及时保存输入输出数据。
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 生成图 URL 列表，有效期7天，请及时保存。
	ResultImage []*string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 结果 detail 数组，Success 代表成功。
	ResultDetails []*string `json:"ResultDetails,omitnil,omitempty" name:"ResultDetails"`

	// 本轮对话前置的历史对话数据（不含生成图）。
	History []*History `json:"History,omitnil,omitempty" name:"History"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanImageChatJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanImageChatJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanImageChatJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanImageChatJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type Replace struct {
	// 占位符序号
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 多媒体详情
	Multimedia []*Multimedia `json:"Multimedia,omitnil,omitempty" name:"Multimedia"`
}

type SearchInfo struct {
	// 搜索引文信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchResults []*SearchResult `json:"SearchResults,omitnil,omitempty" name:"SearchResults"`
}

type SearchResult struct {
	// 搜索引文序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 搜索引文标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 搜索引文链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

// Predefined struct for user
type SetPayModeRequestParams struct {
	// 设置后付费状态，0：后付费；1：预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type SetPayModeRequest struct {
	*tchttp.BaseRequest
	
	// 设置后付费状态，0：后付费；1：预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

func (r *SetPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPayModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetPayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPayModeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetPayModeResponse struct {
	*tchttp.BaseResponse
	Response *SetPayModeResponseParams `json:"Response"`
}

func (r *SetPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SongExt struct {
	// 歌曲id
	SongId *int64 `json:"SongId,omitnil,omitempty" name:"SongId"`

	// 歌曲mid
	SongMid *string `json:"SongMid,omitnil,omitempty" name:"SongMid"`

	// 歌曲是否为vip。1：vip歌曲； 0：普通歌曲。
	Vip *int64 `json:"Vip,omitnil,omitempty" name:"Vip"`
}

// Predefined struct for user
type SubmitHunyuanImageChatJobRequestParams struct {
	// 本轮对话的文本描述。
	// 提交一个任务请求对应发起一轮生图对话，每轮对话中可输入一条 Prompt，生成一张图像，支持通过多轮输入 Prompt 来不断调整图像内容。
	// 推荐使用中文，最多可传1024个 utf-8 字符。
	// 输入示例：
	// <li> 第一轮对话：一颗红色的苹果 </li>
	// <li> 第二轮对话：将苹果改为绿色 </li>
	// <li> 第三轮对话：苹果放在桌子上 </li>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 上传上一轮对话的 ChatId，本轮对话将在指定的上一轮对话结果基础上继续生成图像。
	// 如果不传代表新建一个对话组，重新开启一轮新的对话。
	// 一个对话组中，最多支持进行100轮对话。
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitHunyuanImageChatJobRequest struct {
	*tchttp.BaseRequest
	
	// 本轮对话的文本描述。
	// 提交一个任务请求对应发起一轮生图对话，每轮对话中可输入一条 Prompt，生成一张图像，支持通过多轮输入 Prompt 来不断调整图像内容。
	// 推荐使用中文，最多可传1024个 utf-8 字符。
	// 输入示例：
	// <li> 第一轮对话：一颗红色的苹果 </li>
	// <li> 第二轮对话：将苹果改为绿色 </li>
	// <li> 第三轮对话：苹果放在桌子上 </li>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 上传上一轮对话的 ChatId，本轮对话将在指定的上一轮对话结果基础上继续生成图像。
	// 如果不传代表新建一个对话组，重新开启一轮新的对话。
	// 一个对话组中，最多支持进行100轮对话。
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitHunyuanImageChatJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanImageChatJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "ChatId")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanImageChatJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanImageChatJobResponseParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanImageChatJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanImageChatJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanImageChatJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanImageChatJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanImageJobRequestParams struct {
	// 文本描述。 
	// 算法将根据输入的文本智能生成与之相关的图像。 
	// 不能为空，推荐使用中文。最多可传1024个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 绘画风格。
	// 请在 [混元生图风格列表](https://cloud.tencent.com/document/product/1729/105846) 中选择期望的风格，传入风格编号。
	// 不传默认不指定风格。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3），不传默认使用1024:1024。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 图片生成数量。
	// 支持1 ~ 4张，默认生成1张。
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 随机种子，默认随机。
	// 不传：随机种子生成。
	// 正数：固定种子生成。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`

	// prompt 扩写开关。1为开启，0为关闭，不传默认开启。
	// 开启扩写后，将自动扩写原始输入的 prompt 并使用扩写后的 prompt 生成图片，返回生成图片结果时将一并返回扩写后的 prompt 文本。
	// 如果关闭扩写，将直接使用原始输入的 prompt 生成图片。
	// 建议开启，在多数场景下可提升生成图片效果、丰富生成图片细节。
	Revise *int64 `json:"Revise,omitnil,omitempty" name:"Revise"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitHunyuanImageJobRequest struct {
	*tchttp.BaseRequest
	
	// 文本描述。 
	// 算法将根据输入的文本智能生成与之相关的图像。 
	// 不能为空，推荐使用中文。最多可传1024个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 绘画风格。
	// 请在 [混元生图风格列表](https://cloud.tencent.com/document/product/1729/105846) 中选择期望的风格，传入风格编号。
	// 不传默认不指定风格。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3），不传默认使用1024:1024。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 图片生成数量。
	// 支持1 ~ 4张，默认生成1张。
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 随机种子，默认随机。
	// 不传：随机种子生成。
	// 正数：固定种子生成。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`

	// prompt 扩写开关。1为开启，0为关闭，不传默认开启。
	// 开启扩写后，将自动扩写原始输入的 prompt 并使用扩写后的 prompt 生成图片，返回生成图片结果时将一并返回扩写后的 prompt 文本。
	// 如果关闭扩写，将直接使用原始输入的 prompt 生成图片。
	// 建议开启，在多数场景下可提升生成图片效果、丰富生成图片细节。
	Revise *int64 `json:"Revise,omitnil,omitempty" name:"Revise"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
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
	delete(f, "Num")
	delete(f, "Seed")
	delete(f, "Revise")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
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

// Predefined struct for user
type TextToImageLiteRequestParams struct {
	// 文本描述。
	// 算法将根据输入的文本智能生成与之相关的图像。建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。
	// 不能为空，推荐使用中文。最多可传256个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传256个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [文生图轻量版风格列表](https://cloud.tencent.com/document/product/1729/108992) 中选择期望的风格，传入风格编号。不传默认使用201（日系动漫风格）。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3）、1080:1920（9:16）、1920:1080（16:9），不传默认使用768:768。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按0处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 返回图像方式（base64 或 url) ，二选一，默认为 base64。url 有效期为1小时。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

type TextToImageLiteRequest struct {
	*tchttp.BaseRequest
	
	// 文本描述。
	// 算法将根据输入的文本智能生成与之相关的图像。建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。
	// 不能为空，推荐使用中文。最多可传256个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传256个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [文生图轻量版风格列表](https://cloud.tencent.com/document/product/1729/108992) 中选择期望的风格，传入风格编号。不传默认使用201（日系动漫风格）。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3）、1080:1920（9:16）、1920:1080（16:9），不传默认使用768:768。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按0处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 返回图像方式（base64 或 url) ，二选一，默认为 base64。url 有效期为1小时。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

func (r *TextToImageLiteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToImageLiteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "Style")
	delete(f, "Resolution")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToImageLiteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToImageLiteResponseParams struct {
	// 根据入参 RspImgType 填入不同，返回不同的内容。
	// 如果传入 base64 则返回生成图 Base64 编码。
	// 如果传入 url 则返回的生成图 URL , 有效期1小时，请及时保存。
	ResultImage *string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextToImageLiteResponse struct {
	*tchttp.BaseResponse
	Response *TextToImageLiteResponseParams `json:"Response"`
}

func (r *TextToImageLiteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToImageLiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tool struct {
	// 工具类型，当前只支持function
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 具体要调用的function
	Function *ToolFunction `json:"Function,omitnil,omitempty" name:"Function"`
}

type ToolCall struct {
	// 工具调用id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 工具调用类型，当前只支持function
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 具体的function调用
	Function *ToolCallFunction `json:"Function,omitnil,omitempty" name:"Function"`
}

type ToolCallFunction struct {
	// function名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// function参数，一般为json字符串
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`
}

type ToolFunction struct {
	// function名称，只能包含a-z，A-Z，0-9，\_或-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// function参数，一般为json字符串
	Parameters *string `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// function的简单描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type Usage struct {
	// 输入 Token 数量。
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 输出 Token 数量。
	CompletionTokens *int64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// 总 Token 数量。
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}