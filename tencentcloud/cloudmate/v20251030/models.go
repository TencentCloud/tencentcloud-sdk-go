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

package v20251030

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ChatContent struct {
	// 角色，可选值：
	//  - user
	//  - model_thinking
	//  - model_output
	//  - knowledge
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容分片
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parts []*ChatEventContentPart `json:"Parts,omitnil,omitempty" name:"Parts"`
}

type ChatEventContentPart struct {
	// 文本内容（流式或完整）
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 函数调用信息
	FunctionCall *string `json:"FunctionCall,omitnil,omitempty" name:"FunctionCall"`

	// 函数返回结果
	FunctionResponse *string `json:"FunctionResponse,omitnil,omitempty" name:"FunctionResponse"`
}

// Predefined struct for user
type CloudMateAgentRequestParams struct {
	// 空间 ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 用户提问内容
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 场景 ID
	ScenarioId *string `json:"ScenarioId,omitnil,omitempty" name:"ScenarioId"`

	// 会话 ID
	// 
	// - 填写上一次接口调用返回的会话 ID（SessionId），可在原有会话基础上继续对话
	// - 未填写会话 ID 时，创建新会话
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 是否使用流式响应，默认为 false，不启用流式响应
	Streaming *bool `json:"Streaming,omitnil,omitempty" name:"Streaming"`
}

type CloudMateAgentRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 用户提问内容
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 场景 ID
	ScenarioId *string `json:"ScenarioId,omitnil,omitempty" name:"ScenarioId"`

	// 会话 ID
	// 
	// - 填写上一次接口调用返回的会话 ID（SessionId），可在原有会话基础上继续对话
	// - 未填写会话 ID 时，创建新会话
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 是否使用流式响应，默认为 false，不启用流式响应
	Streaming *bool `json:"Streaming,omitnil,omitempty" name:"Streaming"`
}

func (r *CloudMateAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloudMateAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Message")
	delete(f, "ScenarioId")
	delete(f, "SessionId")
	delete(f, "Streaming")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloudMateAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloudMateAgentResponseParams struct {
	// 会话ID，用于后续继续对话
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Unix 时间戳
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 诊断内容
	Content *ChatContent `json:"Content,omitnil,omitempty" name:"Content"`

	// 是否为部分内容（流式场景）
	Partial *bool `json:"Partial,omitnil,omitempty" name:"Partial"`

	// 本轮对话是否完成
	TurnComplete *bool `json:"TurnComplete,omitnil,omitempty" name:"TurnComplete"`

	// 错误代码，无错误时无该字段
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 错误详细信息，无错误时无该字段
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloudMateAgentResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *CloudMateAgentResponseParams `json:"Response"`
}

func (r *CloudMateAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloudMateAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}