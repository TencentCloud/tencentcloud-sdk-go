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

package v20180522

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ChatRequestParams struct {
	// 聊天输入文本
	Text *string `json:"Text,omitempty" name:"Text"`

	// 腾讯云项目 ID，可填 0，总长度不超过 1024 字节。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// json格式，比如 {"id":"test","gender":"male"}。记录当前与机器人交互的用户id，非必须但强烈建议传入，否则多轮聊天功能会受影响
	User *string `json:"User,omitempty" name:"User"`
}

type ChatRequest struct {
	*tchttp.BaseRequest
	
	// 聊天输入文本
	Text *string `json:"Text,omitempty" name:"Text"`

	// 腾讯云项目 ID，可填 0，总长度不超过 1024 字节。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// json格式，比如 {"id":"test","gender":"male"}。记录当前与机器人交互的用户id，非必须但强烈建议传入，否则多轮聊天功能会受影响
	User *string `json:"User,omitempty" name:"User"`
}

func (r *ChatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "ProjectId")
	delete(f, "User")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatResponseParams struct {
	// 聊天输出文本
	Answer *string `json:"Answer,omitempty" name:"Answer"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChatResponse struct {
	*tchttp.BaseResponse
	Response *ChatResponseParams `json:"Response"`
}

func (r *ChatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SentenceRecognitionRequestParams struct {
	// 腾讯云项目 ID，可填 0，总长度不超过 1024 字节。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子服务类型。2，一句话识别。
	SubServiceType *uint64 `json:"SubServiceType,omitempty" name:"SubServiceType"`

	// 引擎类型。8k：电话 8k 通用模型；16k：16k 通用模型。只支持单声道音频识别。
	EngSerViceType *string `json:"EngSerViceType,omitempty" name:"EngSerViceType"`

	// 语音数据来源。0：语音 URL；1：语音数据（post body）。
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 识别音频的音频格式（支持mp3,wav）。
	VoiceFormat *string `json:"VoiceFormat,omitempty" name:"VoiceFormat"`

	// 用户端对此任务的唯一标识，用户自助生成，用于用户查找识别结果。
	UsrAudioKey *string `json:"UsrAudioKey,omitempty" name:"UsrAudioKey"`

	// 语音 URL，公网可下载。当 SourceType 值为 0 时须填写该字段，为 1 时不填；URL 的长度大于 0，小于 2048，需进行urlencode编码。音频时间长度要小于60s。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 语音数据，当SourceType 值为1时必须填写，为0可不写。要base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。音频数据要小于600kB。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据长度，当 SourceType 值为1时必须填写，为0可不写（此数据长度为数据未进行base64编码时的数据长度）。
	DataLen *int64 `json:"DataLen,omitempty" name:"DataLen"`
}

type SentenceRecognitionRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云项目 ID，可填 0，总长度不超过 1024 字节。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子服务类型。2，一句话识别。
	SubServiceType *uint64 `json:"SubServiceType,omitempty" name:"SubServiceType"`

	// 引擎类型。8k：电话 8k 通用模型；16k：16k 通用模型。只支持单声道音频识别。
	EngSerViceType *string `json:"EngSerViceType,omitempty" name:"EngSerViceType"`

	// 语音数据来源。0：语音 URL；1：语音数据（post body）。
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 识别音频的音频格式（支持mp3,wav）。
	VoiceFormat *string `json:"VoiceFormat,omitempty" name:"VoiceFormat"`

	// 用户端对此任务的唯一标识，用户自助生成，用于用户查找识别结果。
	UsrAudioKey *string `json:"UsrAudioKey,omitempty" name:"UsrAudioKey"`

	// 语音 URL，公网可下载。当 SourceType 值为 0 时须填写该字段，为 1 时不填；URL 的长度大于 0，小于 2048，需进行urlencode编码。音频时间长度要小于60s。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 语音数据，当SourceType 值为1时必须填写，为0可不写。要base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。音频数据要小于600kB。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据长度，当 SourceType 值为1时必须填写，为0可不写（此数据长度为数据未进行base64编码时的数据长度）。
	DataLen *int64 `json:"DataLen,omitempty" name:"DataLen"`
}

func (r *SentenceRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentenceRecognitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SubServiceType")
	delete(f, "EngSerViceType")
	delete(f, "SourceType")
	delete(f, "VoiceFormat")
	delete(f, "UsrAudioKey")
	delete(f, "Url")
	delete(f, "Data")
	delete(f, "DataLen")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SentenceRecognitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SentenceRecognitionResponseParams struct {
	// 识别结果。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SentenceRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *SentenceRecognitionResponseParams `json:"Response"`
}

func (r *SentenceRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentenceRecognitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SimultaneousInterpretingRequestParams struct {
	// 腾讯云项目 ID，可填 0，总长度不超过 1024 字节。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子服务类型。0：离线语音识别。1：实时流式识别，2，一句话识别。3：同传。
	SubServiceType *uint64 `json:"SubServiceType,omitempty" name:"SubServiceType"`

	// 识别引擎类型。8k_zh： 8k 中文会场模型；16k_zh：16k 中文会场模型，8k_en： 8k 英文会场模型；16k_en：16k 英文会场模型。当前仅支持16K。
	RecEngineModelType *string `json:"RecEngineModelType,omitempty" name:"RecEngineModelType"`

	// 语音数据，要base64编码。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据长度。
	DataLen *uint64 `json:"DataLen,omitempty" name:"DataLen"`

	// 声音id，标识一句话。
	VoiceId *string `json:"VoiceId,omitempty" name:"VoiceId"`

	// 是否是一句话的结束。
	IsEnd *uint64 `json:"IsEnd,omitempty" name:"IsEnd"`

	// 声音编码的格式1:pcm，4:speex，6:silk，默认为1。
	VoiceFormat *uint64 `json:"VoiceFormat,omitempty" name:"VoiceFormat"`

	// 是否需要翻译结果，1表示需要翻译，0是不需要。
	OpenTranslate *uint64 `json:"OpenTranslate,omitempty" name:"OpenTranslate"`

	// 如果需要翻译，表示源语言类型，可取值：zh，en。
	SourceLanguage *string `json:"SourceLanguage,omitempty" name:"SourceLanguage"`

	// 如果需要翻译，表示目标语言类型，可取值：zh，en。
	TargetLanguage *string `json:"TargetLanguage,omitempty" name:"TargetLanguage"`

	// 表明当前语音分片的索引，从0开始
	Seq *uint64 `json:"Seq,omitempty" name:"Seq"`
}

type SimultaneousInterpretingRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云项目 ID，可填 0，总长度不超过 1024 字节。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子服务类型。0：离线语音识别。1：实时流式识别，2，一句话识别。3：同传。
	SubServiceType *uint64 `json:"SubServiceType,omitempty" name:"SubServiceType"`

	// 识别引擎类型。8k_zh： 8k 中文会场模型；16k_zh：16k 中文会场模型，8k_en： 8k 英文会场模型；16k_en：16k 英文会场模型。当前仅支持16K。
	RecEngineModelType *string `json:"RecEngineModelType,omitempty" name:"RecEngineModelType"`

	// 语音数据，要base64编码。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据长度。
	DataLen *uint64 `json:"DataLen,omitempty" name:"DataLen"`

	// 声音id，标识一句话。
	VoiceId *string `json:"VoiceId,omitempty" name:"VoiceId"`

	// 是否是一句话的结束。
	IsEnd *uint64 `json:"IsEnd,omitempty" name:"IsEnd"`

	// 声音编码的格式1:pcm，4:speex，6:silk，默认为1。
	VoiceFormat *uint64 `json:"VoiceFormat,omitempty" name:"VoiceFormat"`

	// 是否需要翻译结果，1表示需要翻译，0是不需要。
	OpenTranslate *uint64 `json:"OpenTranslate,omitempty" name:"OpenTranslate"`

	// 如果需要翻译，表示源语言类型，可取值：zh，en。
	SourceLanguage *string `json:"SourceLanguage,omitempty" name:"SourceLanguage"`

	// 如果需要翻译，表示目标语言类型，可取值：zh，en。
	TargetLanguage *string `json:"TargetLanguage,omitempty" name:"TargetLanguage"`

	// 表明当前语音分片的索引，从0开始
	Seq *uint64 `json:"Seq,omitempty" name:"Seq"`
}

func (r *SimultaneousInterpretingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SimultaneousInterpretingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SubServiceType")
	delete(f, "RecEngineModelType")
	delete(f, "Data")
	delete(f, "DataLen")
	delete(f, "VoiceId")
	delete(f, "IsEnd")
	delete(f, "VoiceFormat")
	delete(f, "OpenTranslate")
	delete(f, "SourceLanguage")
	delete(f, "TargetLanguage")
	delete(f, "Seq")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SimultaneousInterpretingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SimultaneousInterpretingResponseParams struct {
	// 语音识别的结果
	AsrText *string `json:"AsrText,omitempty" name:"AsrText"`

	// 机器翻译的结果
	NmtText *string `json:"NmtText,omitempty" name:"NmtText"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SimultaneousInterpretingResponse struct {
	*tchttp.BaseResponse
	Response *SimultaneousInterpretingResponseParams `json:"Response"`
}

func (r *SimultaneousInterpretingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SimultaneousInterpretingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToVoiceRequestParams struct {
	// 合成语音的源文本，按UTF-8编码统一计算。
	// 中文最大支持100个汉字（全角标点符号算一个汉字）；英文最大支持400个字母（半角标点符号算一个字母）。包含空格等字符时需要url encode再传输。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 一次请求对应一个SessionId，会原样返回，建议传入类似于uuid的字符串防止重复。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 模型类型，1-默认模型。
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// 音量大小，范围：[0，10]，分别对应11个等级的音量，默认为0，代表正常音量。没有静音选项。
	// 输入除以上整数之外的其他参数不生效，按默认值处理。
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，2]，分别对应不同语速：<li>-2代表0.6倍</li><li>-1代表0.8倍</li><li>0代表1.0倍（默认）</li><li>1代表1.2倍</li><li>2代表1.5倍</li>输入除以上整数之外的其他参数不生效，按默认值处理。
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// 项目id，用户自定义，默认为0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 音色<li>0-亲和女声(默认)</li><li>1-亲和男声</li><li>2-成熟男声</li><li>3-活力男声</li><li>4-温暖女声</li><li>5-情感女声</li><li>6-情感男声</li>
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 主语言类型：<li>1-中文（默认）</li><li>2-英文</li>
	PrimaryLanguage *uint64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// 音频采样率：<li>16000：16k（默认）</li><li>8000：8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 返回音频格式，可取值：wav（默认），mp3
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type TextToVoiceRequest struct {
	*tchttp.BaseRequest
	
	// 合成语音的源文本，按UTF-8编码统一计算。
	// 中文最大支持100个汉字（全角标点符号算一个汉字）；英文最大支持400个字母（半角标点符号算一个字母）。包含空格等字符时需要url encode再传输。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 一次请求对应一个SessionId，会原样返回，建议传入类似于uuid的字符串防止重复。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 模型类型，1-默认模型。
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// 音量大小，范围：[0，10]，分别对应11个等级的音量，默认为0，代表正常音量。没有静音选项。
	// 输入除以上整数之外的其他参数不生效，按默认值处理。
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，2]，分别对应不同语速：<li>-2代表0.6倍</li><li>-1代表0.8倍</li><li>0代表1.0倍（默认）</li><li>1代表1.2倍</li><li>2代表1.5倍</li>输入除以上整数之外的其他参数不生效，按默认值处理。
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// 项目id，用户自定义，默认为0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 音色<li>0-亲和女声(默认)</li><li>1-亲和男声</li><li>2-成熟男声</li><li>3-活力男声</li><li>4-温暖女声</li><li>5-情感女声</li><li>6-情感男声</li>
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 主语言类型：<li>1-中文（默认）</li><li>2-英文</li>
	PrimaryLanguage *uint64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// 音频采样率：<li>16000：16k（默认）</li><li>8000：8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 返回音频格式，可取值：wav（默认），mp3
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

func (r *TextToVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToVoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "SessionId")
	delete(f, "ModelType")
	delete(f, "Volume")
	delete(f, "Speed")
	delete(f, "ProjectId")
	delete(f, "VoiceType")
	delete(f, "PrimaryLanguage")
	delete(f, "SampleRate")
	delete(f, "Codec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToVoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToVoiceResponseParams struct {
	// base64编码的wav/mp3音频数据
	Audio *string `json:"Audio,omitempty" name:"Audio"`

	// 一次请求对应一个SessionId
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextToVoiceResponse struct {
	*tchttp.BaseResponse
	Response *TextToVoiceResponseParams `json:"Response"`
}

func (r *TextToVoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToVoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}