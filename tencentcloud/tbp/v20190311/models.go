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

package v20190311

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type PostAudioRequest struct {
	*tchttp.BaseRequest

	// 机器人标识
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 语音识别引擎类型,{8k_0、16k_0、16k_en}
	EngineModelType *string `json:"EngineModelType,omitempty" name:"EngineModelType"`

	// 输入音频文件编码格式。1：wav(pcm)；4：speex(sp)；6：silk
	AsrVoiceFormat *uint64 `json:"AsrVoiceFormat,omitempty" name:"AsrVoiceFormat"`

	// asr请求Id。长度为16位的字符串，同一句话VoiceId保持一致
	VoiceId *string `json:"VoiceId,omitempty" name:"VoiceId"`

	// 语音分片序列号。同一句话必须从0开始，依次递增
	Seq *uint64 `json:"Seq,omitempty" name:"Seq"`

	// 是否为尾包。传递最后一个语音分片时为true，其他为false
	IsEnd *bool `json:"IsEnd,omitempty" name:"IsEnd"`

	// 待识别音频字节流
	WaveData *string `json:"WaveData,omitempty" name:"WaveData"`

	// 子账户id，每个终端对应一个
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 机器人版本号。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`

	// 识别返回文本编码格式	0：UTF-8（默认）；1：GB2312；2：GBK；3：BIG5
	ResultTextFormat *uint64 `json:"ResultTextFormat,omitempty" name:"ResultTextFormat"`

	// 透传字段，传递给后台
	SessionAttributes *string `json:"SessionAttributes,omitempty" name:"SessionAttributes"`

	// 是否将机器人回答合成音频并返回url
	NeedTts *bool `json:"NeedTts,omitempty" name:"NeedTts"`

	// 音量大小，范围：[0，10]。默认值为0，代表正常音量
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，2]。0代表1.0倍
	Speed *int64 `json:"Speed,omitempty" name:"Speed"`

	// 音色,{0：女声,1:男声}
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 返回音频的采样率{8k,16k}。默认16k
	SampleRate *string `json:"SampleRate,omitempty" name:"SampleRate"`

	// 机器人环境{dev:测试;release:线上}。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotEnv *string `json:"BotEnv,omitempty" name:"BotEnv"`

	// TTS合成音频格式，{0：wav}。该字段在当前版本仅支持取值为0。
	TtsVoiceFormat *uint64 `json:"TtsVoiceFormat,omitempty" name:"TtsVoiceFormat"`
}

func (r *PostAudioRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PostAudioRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PostAudioResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前会话状态。取值:"start"/"continue"/"complete"
	// 注意：此字段可能返回 null，表示取不到有效值。
		DialogStatus *string `json:"DialogStatus,omitempty" name:"DialogStatus"`

		// 匹配到的机器人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		BotName *string `json:"BotName,omitempty" name:"BotName"`

		// 匹配到的意图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		IntentName *string `json:"IntentName,omitempty" name:"IntentName"`

		// 机器人应答文本
		ResponseText *string `json:"ResponseText,omitempty" name:"ResponseText"`

		// 语义解析的槽位结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		SlotInfoList []*SlotInfo `json:"SlotInfoList,omitempty" name:"SlotInfoList" list`

		// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
		SessionAttributes *string `json:"SessionAttributes,omitempty" name:"SessionAttributes"`

		// 用户说法。该说法是用户原生说法或ASR识别结果，未经过语义优化
	// 注意：此字段可能返回 null，表示取不到有效值。
		Question *string `json:"Question,omitempty" name:"Question"`

		// tts合成pcm音频存储链接。仅当请求参数NeedTts=true时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
		WaveUrl *string `json:"WaveUrl,omitempty" name:"WaveUrl"`

		// tts合成pcm音频
	// 注意：此字段可能返回 null，表示取不到有效值。
		WaveData *string `json:"WaveData,omitempty" name:"WaveData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PostAudioResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PostAudioResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PostTextRequest struct {
	*tchttp.BaseRequest

	// 机器人标识
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 请求的文本
	InputText *string `json:"InputText,omitempty" name:"InputText"`

	// 子账户id，每个终端对应一个
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 机器人版本号。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`

	// 透传字段，传递给后台
	SessionAttributes *string `json:"SessionAttributes,omitempty" name:"SessionAttributes"`

	// 是否将机器人回答合成音频并返回url
	NeedTts *bool `json:"NeedTts,omitempty" name:"NeedTts"`

	// 音量大小，范围：[0，10]。默认值为0，代表正常音量
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，2]。0代表1.0倍
	Speed *int64 `json:"Speed,omitempty" name:"Speed"`

	// 音色,{0：女声,1:男声}
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 返回音频的采样率{8k,16k}。默认16k
	SampleRate *string `json:"SampleRate,omitempty" name:"SampleRate"`

	// 机器人环境{dev:测试;release:线上}。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotEnv *string `json:"BotEnv,omitempty" name:"BotEnv"`

	// TTS合成音频格式，{0：wav}。该字段在当前版本仅支持取值为0。
	TtsVoiceFormat *uint64 `json:"TtsVoiceFormat,omitempty" name:"TtsVoiceFormat"`
}

func (r *PostTextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PostTextRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PostTextResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前会话状态。取值:"start"/"continue"/"complete"
	// 注意：此字段可能返回 null，表示取不到有效值。
		DialogStatus *string `json:"DialogStatus,omitempty" name:"DialogStatus"`

		// 匹配到的机器人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		BotName *string `json:"BotName,omitempty" name:"BotName"`

		// 匹配到的意图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		IntentName *string `json:"IntentName,omitempty" name:"IntentName"`

		// 机器人回答
		ResponseText *string `json:"ResponseText,omitempty" name:"ResponseText"`

		// 语义解析的槽位结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		SlotInfoList []*SlotInfo `json:"SlotInfoList,omitempty" name:"SlotInfoList" list`

		// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
		SessionAttributes *string `json:"SessionAttributes,omitempty" name:"SessionAttributes"`

		// 用户说法。该说法是用户原生说法或ASR识别结果，未经过语义优化
		Question *string `json:"Question,omitempty" name:"Question"`

		// tts合成pcm音频存储链接。仅当请求参数NeedTts=true时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
		WaveUrl *string `json:"WaveUrl,omitempty" name:"WaveUrl"`

		// tts合成的pcm音频。二进制数组经过base64编码(暂时不返回)
	// 注意：此字段可能返回 null，表示取不到有效值。
		WaveData *string `json:"WaveData,omitempty" name:"WaveData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PostTextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PostTextResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetRequest struct {
	*tchttp.BaseRequest

	// 机器人标识
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 子账户id，每个终端对应一个
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 机器人版本号。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotVersion *string `json:"BotVersion,omitempty" name:"BotVersion"`

	// 机器人环境{dev:测试;release:线上}。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotEnv *string `json:"BotEnv,omitempty" name:"BotEnv"`
}

func (r *ResetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前会话状态。取值:"start"/"continue"/"complete"
	// 注意：此字段可能返回 null，表示取不到有效值。
		DialogStatus *string `json:"DialogStatus,omitempty" name:"DialogStatus"`

		// 匹配到的机器人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		BotName *string `json:"BotName,omitempty" name:"BotName"`

		// 匹配到的意图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		IntentName *string `json:"IntentName,omitempty" name:"IntentName"`

		// 机器人回答
		ResponseText *string `json:"ResponseText,omitempty" name:"ResponseText"`

		// 语义解析的槽位结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		SlotInfoList []*SlotInfo `json:"SlotInfoList,omitempty" name:"SlotInfoList" list`

		// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
		SessionAttributes *string `json:"SessionAttributes,omitempty" name:"SessionAttributes"`

		// 用户说法。该说法是用户原生说法或ASR识别结果，未经过语义优化
	// 注意：此字段可能返回 null，表示取不到有效值。
		Question *string `json:"Question,omitempty" name:"Question"`

		// tts合成pcm音频存储链接。仅当请求参数NeedTts=true时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
		WaveUrl *string `json:"WaveUrl,omitempty" name:"WaveUrl"`

		// tts合成的pcm音频。二进制数组经过base64编码(暂时不返回)
	// 注意：此字段可能返回 null，表示取不到有效值。
		WaveData *string `json:"WaveData,omitempty" name:"WaveData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SlotInfo struct {

	// 槽位名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotName *string `json:"SlotName,omitempty" name:"SlotName"`

	// 槽位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotValue *string `json:"SlotValue,omitempty" name:"SlotValue"`
}
