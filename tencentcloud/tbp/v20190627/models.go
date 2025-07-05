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

package v20190627

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Group struct {
	// 消息类型参考互联网MIME类型标准，当前仅支持"text/plain"。
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 返回内容以链接形式提供。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 普通文本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ResponseMessage struct {
	// 消息组列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupList []*Group `json:"GroupList,omitnil,omitempty" name:"GroupList"`
}

type SlotInfo struct {
	// 槽位名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotName *string `json:"SlotName,omitnil,omitempty" name:"SlotName"`

	// 槽位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotValue *string `json:"SlotValue,omitnil,omitempty" name:"SlotValue"`
}

// Predefined struct for user
type TextProcessRequestParams struct {
	// 机器人标识，用于定义抽象机器人。
	BotId *string `json:"BotId,omitnil,omitempty" name:"BotId"`

	// 机器人版本，取值"dev"或"release"，{调试版本：dev；线上版本：release}。
	BotEnv *string `json:"BotEnv,omitnil,omitempty" name:"BotEnv"`

	// 终端标识，每个终端(或线程)对应一个，区分并发多用户。
	TerminalId *string `json:"TerminalId,omitnil,omitempty" name:"TerminalId"`

	// 请求的文本。
	InputText *string `json:"InputText,omitnil,omitempty" name:"InputText"`

	// 透传字段，透传给用户自定义的WebService服务。
	SessionAttributes *string `json:"SessionAttributes,omitnil,omitempty" name:"SessionAttributes"`

	// 平台类型，{小程序：MiniProgram；小微：XiaoWei；公众号：OfficialAccount；企业微信: WXWork}。
	PlatformType *string `json:"PlatformType,omitnil,omitempty" name:"PlatformType"`

	// 当PlatformType为微信公众号或企业微信时，传递对应微信公众号或企业微信的唯一标识
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type TextProcessRequest struct {
	*tchttp.BaseRequest
	
	// 机器人标识，用于定义抽象机器人。
	BotId *string `json:"BotId,omitnil,omitempty" name:"BotId"`

	// 机器人版本，取值"dev"或"release"，{调试版本：dev；线上版本：release}。
	BotEnv *string `json:"BotEnv,omitnil,omitempty" name:"BotEnv"`

	// 终端标识，每个终端(或线程)对应一个，区分并发多用户。
	TerminalId *string `json:"TerminalId,omitnil,omitempty" name:"TerminalId"`

	// 请求的文本。
	InputText *string `json:"InputText,omitnil,omitempty" name:"InputText"`

	// 透传字段，透传给用户自定义的WebService服务。
	SessionAttributes *string `json:"SessionAttributes,omitnil,omitempty" name:"SessionAttributes"`

	// 平台类型，{小程序：MiniProgram；小微：XiaoWei；公众号：OfficialAccount；企业微信: WXWork}。
	PlatformType *string `json:"PlatformType,omitnil,omitempty" name:"PlatformType"`

	// 当PlatformType为微信公众号或企业微信时，传递对应微信公众号或企业微信的唯一标识
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *TextProcessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextProcessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotId")
	delete(f, "BotEnv")
	delete(f, "TerminalId")
	delete(f, "InputText")
	delete(f, "SessionAttributes")
	delete(f, "PlatformType")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextProcessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextProcessResponseParams struct {
	// 当前会话状态{会话开始: START; 会话中: COUTINUE; 会话结束: COMPLETE}。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DialogStatus *string `json:"DialogStatus,omitnil,omitempty" name:"DialogStatus"`

	// 匹配到的机器人名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotName *string `json:"BotName,omitnil,omitempty" name:"BotName"`

	// 匹配到的意图名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentName *string `json:"IntentName,omitnil,omitempty" name:"IntentName"`

	// 槽位信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotInfoList []*SlotInfo `json:"SlotInfoList,omitnil,omitempty" name:"SlotInfoList"`

	// 原始的用户说法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputText *string `json:"InputText,omitnil,omitempty" name:"InputText"`

	// 机器人应答。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseMessage *ResponseMessage `json:"ResponseMessage,omitnil,omitempty" name:"ResponseMessage"`

	// 透传字段，由用户自定义的WebService服务返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionAttributes *string `json:"SessionAttributes,omitnil,omitempty" name:"SessionAttributes"`

	// 结果类型 {中间逻辑出错:0; 任务型机器人:1; 问答型机器人:2; 闲聊型机器人:3; 未匹配上，返回预设兜底话术:5; 未匹配上，返回相似问题列表:6}。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// 机器人对话的应答文本。	
	ResponseText *string `json:"ResponseText,omitnil,omitempty" name:"ResponseText"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextProcessResponse struct {
	*tchttp.BaseResponse
	Response *TextProcessResponseParams `json:"Response"`
}

func (r *TextProcessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextResetRequestParams struct {
	// 机器人标识，用于定义抽象机器人。
	BotId *string `json:"BotId,omitnil,omitempty" name:"BotId"`

	// 机器人版本，取值"dev"或"release"，{调试版本：dev；线上版本：release}。
	BotEnv *string `json:"BotEnv,omitnil,omitempty" name:"BotEnv"`

	// 终端标识，每个终端(或线程)对应一个，区分并发多用户。
	TerminalId *string `json:"TerminalId,omitnil,omitempty" name:"TerminalId"`

	// 平台类型，{小程序：MiniProgram；小微：XiaoWei；公众号：OfficialAccount；企业微信: WXWork}。
	PlatformType *string `json:"PlatformType,omitnil,omitempty" name:"PlatformType"`

	// 当PlatformType为微信公众号或企业微信时，传递对应微信公众号或企业微信的唯一标识
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type TextResetRequest struct {
	*tchttp.BaseRequest
	
	// 机器人标识，用于定义抽象机器人。
	BotId *string `json:"BotId,omitnil,omitempty" name:"BotId"`

	// 机器人版本，取值"dev"或"release"，{调试版本：dev；线上版本：release}。
	BotEnv *string `json:"BotEnv,omitnil,omitempty" name:"BotEnv"`

	// 终端标识，每个终端(或线程)对应一个，区分并发多用户。
	TerminalId *string `json:"TerminalId,omitnil,omitempty" name:"TerminalId"`

	// 平台类型，{小程序：MiniProgram；小微：XiaoWei；公众号：OfficialAccount；企业微信: WXWork}。
	PlatformType *string `json:"PlatformType,omitnil,omitempty" name:"PlatformType"`

	// 当PlatformType为微信公众号或企业微信时，传递对应微信公众号或企业微信的唯一标识
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *TextResetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextResetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotId")
	delete(f, "BotEnv")
	delete(f, "TerminalId")
	delete(f, "PlatformType")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextResetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextResetResponseParams struct {
	// 当前会话状态{会话开始: START; 会话中: COUTINUE; 会话结束: COMPLETE}。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DialogStatus *string `json:"DialogStatus,omitnil,omitempty" name:"DialogStatus"`

	// 匹配到的机器人名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotName *string `json:"BotName,omitnil,omitempty" name:"BotName"`

	// 匹配到的意图名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentName *string `json:"IntentName,omitnil,omitempty" name:"IntentName"`

	// 槽位信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotInfoList []*SlotInfo `json:"SlotInfoList,omitnil,omitempty" name:"SlotInfoList"`

	// 原始的用户说法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputText *string `json:"InputText,omitnil,omitempty" name:"InputText"`

	// 机器人应答。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseMessage *ResponseMessage `json:"ResponseMessage,omitnil,omitempty" name:"ResponseMessage"`

	// 透传字段，由用户自定义的WebService服务返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionAttributes *string `json:"SessionAttributes,omitnil,omitempty" name:"SessionAttributes"`

	// 结果类型 {中间逻辑出错:0; 任务型机器人:1; 问答型机器人:2; 闲聊型机器人:3; 未匹配上，返回预设兜底话术:5; 未匹配上，返回相似问题列表:6}。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// 机器人对话的应答文本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseText *string `json:"ResponseText,omitnil,omitempty" name:"ResponseText"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextResetResponse struct {
	*tchttp.BaseResponse
	Response *TextResetResponseParams `json:"Response"`
}

func (r *TextResetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextResetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}