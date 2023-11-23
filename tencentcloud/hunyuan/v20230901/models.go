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
	// 1.长度最多为40, 按对话时间从旧到新在数组中排列。
	// 2.Message的Role当前可选值：user、assistant，其中，user和assistant需要交替出现(一问一答)，最后一个为user提问, 且Content不能为空。
	// 3.Messages中Content总长度不超过16000token，超过则会截断最前面的内容，只保留尾部内容。建议不超过4000token。
	Messages []*Message `json:"Messages,omitnil" name:"Messages"`

	// 说明：
	// 1.影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2.默认1.0，取值区间为[0.0, 1.0]。
	// 3.非必要不建议使用, 不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil" name:"TopP"`

	// 说明：
	// 1.较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2.默认1.0，取值区间为[0.0，2.0]。
	// 3.非必要不建议使用,不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil" name:"Temperature"`
}

type ChatProRequest struct {
	*tchttp.BaseRequest
	
	// 聊天上下文信息。
	// 说明：
	// 1.长度最多为40, 按对话时间从旧到新在数组中排列。
	// 2.Message的Role当前可选值：user、assistant，其中，user和assistant需要交替出现(一问一答)，最后一个为user提问, 且Content不能为空。
	// 3.Messages中Content总长度不超过16000token，超过则会截断最前面的内容，只保留尾部内容。建议不超过4000token。
	Messages []*Message `json:"Messages,omitnil" name:"Messages"`

	// 说明：
	// 1.影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2.默认1.0，取值区间为[0.0, 1.0]。
	// 3.非必要不建议使用, 不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil" name:"TopP"`

	// 说明：
	// 1.较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2.默认1.0，取值区间为[0.0，2.0]。
	// 3.非必要不建议使用,不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil" name:"Temperature"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

    type ChatProResponse struct {
    	tchttp.BaseSSEResponse
    }

// Predefined struct for user
type ChatStdRequestParams struct {
	// 聊天上下文信息。
	// 说明：
	// 1.长度最多为40, 按对话时间从旧到新在数组中排列。
	// 2.Message的Role当前可选值：user、assistant，其中，user和assistant需要交替出现(一问一答)，最后一个为user提问, 且Content不能为空。
	// 3.Messages中Content总长度不超过16000token，超过则会截断最前面的内容，只保留尾部内容。建议不超过4000token。
	Messages []*Message `json:"Messages,omitnil" name:"Messages"`

	// 说明：
	// 1.影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2.默认1.0，取值区间为[0.0, 1.0]。
	// 3.非必要不建议使用, 不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil" name:"TopP"`

	// 说明：
	// 1.较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2.默认1.0，取值区间为[0.0，2.0]。
	// 3.非必要不建议使用,不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil" name:"Temperature"`
}

type ChatStdRequest struct {
	*tchttp.BaseRequest
	
	// 聊天上下文信息。
	// 说明：
	// 1.长度最多为40, 按对话时间从旧到新在数组中排列。
	// 2.Message的Role当前可选值：user、assistant，其中，user和assistant需要交替出现(一问一答)，最后一个为user提问, 且Content不能为空。
	// 3.Messages中Content总长度不超过16000token，超过则会截断最前面的内容，只保留尾部内容。建议不超过4000token。
	Messages []*Message `json:"Messages,omitnil" name:"Messages"`

	// 说明：
	// 1.影响输出文本的多样性，取值越大，生成文本的多样性越强。
	// 2.默认1.0，取值区间为[0.0, 1.0]。
	// 3.非必要不建议使用, 不合理的取值会影响效果。
	TopP *float64 `json:"TopP,omitnil" name:"TopP"`

	// 说明：
	// 1.较高的数值会使输出更加随机，而较低的数值会使其更加集中和确定。
	// 2.默认1.0，取值区间为[0.0，2.0]。
	// 3.非必要不建议使用,不合理的取值会影响效果。
	Temperature *float64 `json:"Temperature,omitnil" name:"Temperature"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatStdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

    type ChatStdResponse struct {
    	tchttp.BaseSSEResponse
    }

type Choice struct {
	// 流式结束标志位，为 stop 则表示尾包。
	FinishReason *string `json:"FinishReason,omitnil" name:"FinishReason"`

	// 返回值。
	Delta *Delta `json:"Delta,omitnil" name:"Delta"`
}

type Delta struct {
	// 角色名称。
	Role *string `json:"Role,omitnil" name:"Role"`

	// 内容详情。
	Content *string `json:"Content,omitnil" name:"Content"`
}

type ErrorMsg struct {
	// 错误提示信息。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 错误码。
	// 4000 服务内部异常。
	// 4001 请求模型超时。
	Code *int64 `json:"Code,omitnil" name:"Code"`
}

// Predefined struct for user
type GetTokenCountRequestParams struct {
	// 输入文本
	Prompt *string `json:"Prompt,omitnil" name:"Prompt"`
}

type GetTokenCountRequest struct {
	*tchttp.BaseRequest
	
	// 输入文本
	Prompt *string `json:"Prompt,omitnil" name:"Prompt"`
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
	TokenCount *int64 `json:"TokenCount,omitnil" name:"TokenCount"`

	// 字符计数
	CharacterCount *int64 `json:"CharacterCount,omitnil" name:"CharacterCount"`

	// 切分后的列表
	Tokens []*string `json:"Tokens,omitnil" name:"Tokens"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Role *string `json:"Role,omitnil" name:"Role"`

	// 消息的内容
	Content *string `json:"Content,omitnil" name:"Content"`
}

type Usage struct {
	// 输入 token 数量。
	PromptTokens *int64 `json:"PromptTokens,omitnil" name:"PromptTokens"`

	// 输出 token 数量。
	CompletionTokens *int64 `json:"CompletionTokens,omitnil" name:"CompletionTokens"`

	// 总 token 数量。
	TotalTokens *int64 `json:"TotalTokens,omitnil" name:"TotalTokens"`
}