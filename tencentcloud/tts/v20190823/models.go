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

package v20190823

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateTtsTaskRequestParams struct {
	// 合成语音的源文本，按UTF-8编码统一计算，最多支持10万字符
	Text *string `json:"Text,omitempty" name:"Text"`

	// 模型类型，1-默认模型。
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// 音量大小，范围：[0，10]，分别对应11个等级的音量，默认为0，代表正常音量。没有静音选项。
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，2]，分别对应不同语速：<li>-2代表0.6倍</li><li>-1代表0.8倍</li><li>0代表1.0倍（默认）</li><li>1代表1.2倍</li><li>2代表1.5倍</li>如果需要更细化的语速，可以保留小数点后 2 位，例如0.5 1.1 1.8等。<br>参数值与实际语速转换，可参考[代码示例](https://sdk-1300466766.cos.ap-shanghai.myqcloud.com/sample/speed_sample.tar.gz)
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// 项目id，用户自定义，默认为0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色，请参见[购买指南](https://cloud.tencent.com/document/product/1073/34112)。完整的音色 ID 列表请参见[音色列表](https://cloud.tencent.com/document/product/1073/92668)。
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 主语言类型：<li>1-中文（默认）</li><li>2-英文</li><li>3-日文</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// 音频采样率：<li>16000：16k（默认）</li><li>8000：8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 返回音频格式，可取值：mp3（默认），wav，pcm
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 回调 URL，用户自行搭建的用于接收识别结果的服务URL。如果用户使用轮询方式获取识别结果，则无需提交该参数。[回调说明](https://cloud.tencent.com/document/product/1073/55746)
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 是否开启时间戳功能，默认为false。
	EnableSubtitle *bool `json:"EnableSubtitle,omitempty" name:"EnableSubtitle"`

	// 旁白与对白文本解析，分别合成相应风格（仅适用于旁对白音色10510000、100510000），默认 false
	VoiceoverDialogueSplit *bool `json:"VoiceoverDialogueSplit,omitempty" name:"VoiceoverDialogueSplit"`
}

type CreateTtsTaskRequest struct {
	*tchttp.BaseRequest
	
	// 合成语音的源文本，按UTF-8编码统一计算，最多支持10万字符
	Text *string `json:"Text,omitempty" name:"Text"`

	// 模型类型，1-默认模型。
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// 音量大小，范围：[0，10]，分别对应11个等级的音量，默认为0，代表正常音量。没有静音选项。
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，2]，分别对应不同语速：<li>-2代表0.6倍</li><li>-1代表0.8倍</li><li>0代表1.0倍（默认）</li><li>1代表1.2倍</li><li>2代表1.5倍</li>如果需要更细化的语速，可以保留小数点后 2 位，例如0.5 1.1 1.8等。<br>参数值与实际语速转换，可参考[代码示例](https://sdk-1300466766.cos.ap-shanghai.myqcloud.com/sample/speed_sample.tar.gz)
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// 项目id，用户自定义，默认为0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色，请参见[购买指南](https://cloud.tencent.com/document/product/1073/34112)。完整的音色 ID 列表请参见[音色列表](https://cloud.tencent.com/document/product/1073/92668)。
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 主语言类型：<li>1-中文（默认）</li><li>2-英文</li><li>3-日文</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// 音频采样率：<li>16000：16k（默认）</li><li>8000：8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 返回音频格式，可取值：mp3（默认），wav，pcm
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 回调 URL，用户自行搭建的用于接收识别结果的服务URL。如果用户使用轮询方式获取识别结果，则无需提交该参数。[回调说明](https://cloud.tencent.com/document/product/1073/55746)
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 是否开启时间戳功能，默认为false。
	EnableSubtitle *bool `json:"EnableSubtitle,omitempty" name:"EnableSubtitle"`

	// 旁白与对白文本解析，分别合成相应风格（仅适用于旁对白音色10510000、100510000），默认 false
	VoiceoverDialogueSplit *bool `json:"VoiceoverDialogueSplit,omitempty" name:"VoiceoverDialogueSplit"`
}

func (r *CreateTtsTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTtsTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "ModelType")
	delete(f, "Volume")
	delete(f, "Speed")
	delete(f, "ProjectId")
	delete(f, "VoiceType")
	delete(f, "PrimaryLanguage")
	delete(f, "SampleRate")
	delete(f, "Codec")
	delete(f, "CallbackUrl")
	delete(f, "EnableSubtitle")
	delete(f, "VoiceoverDialogueSplit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTtsTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTtsTaskRespData struct {
	// 任务ID，可通过此ID在轮询接口获取合成状态与结果。注意：TaskId数据类型为string
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CreateTtsTaskResponseParams struct {
	// 任务 id
	Data *CreateTtsTaskRespData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTtsTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTtsTaskResponseParams `json:"Response"`
}

func (r *CreateTtsTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTtsTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTtsTaskStatusRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTtsTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTtsTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTtsTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTtsTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTtsTaskStatusRespData struct {
	// 任务标识。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态码，0：任务等待，1：任务执行中，2：任务成功，3：任务失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 任务状态，waiting：任务等待，doing：任务执行中，success：任务成功，failed：任务失败。
	StatusStr *string `json:"StatusStr,omitempty" name:"StatusStr"`

	// 合成音频COS地址（链接有效期1天）。
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 时间戳信息，若未开启时间戳，则返回空数组。
	Subtitles []*Subtitle `json:"Subtitles,omitempty" name:"Subtitles"`

	// 失败原因说明。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

// Predefined struct for user
type DescribeTtsTaskStatusResponseParams struct {
	// 任务状态返回
	Data *DescribeTtsTaskStatusRespData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTtsTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTtsTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTtsTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTtsTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Subtitle struct {
	// ⽂本信息。
	Text *string `json:"Text,omitempty" name:"Text"`

	// ⽂本对应tts语⾳开始时间戳，单位ms。
	BeginTime *int64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// ⽂本对应tts语⾳结束时间戳，单位ms。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 该字在整句中的开始位置，从0开始。
	BeginIndex *int64 `json:"BeginIndex,omitempty" name:"BeginIndex"`

	// 该字在整句中的结束位置，从0开始。
	EndIndex *int64 `json:"EndIndex,omitempty" name:"EndIndex"`

	// 该字的音素
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phoneme *string `json:"Phoneme,omitempty" name:"Phoneme"`
}

// Predefined struct for user
type TextToVoiceRequestParams struct {
	// 合成语音的源文本，按UTF-8编码统一计算。
	// 中文最大支持150个汉字（全角标点符号算一个汉字）；英文最大支持500个字母（半角标点符号算一个字母）。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 一次请求对应一个SessionId，会原样返回，建议传入类似于uuid的字符串防止重复。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 音量大小，范围[0，10]，对应音量大小。默认为0，代表正常音量，值越大音量越高。
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，6]，分别对应不同语速：<li>-2代表0.6倍</li><li>-1代表0.8倍</li><li>0代表1.0倍（默认）</li><li>1代表1.2倍</li><li>2代表1.5倍</li><li>6代表2.5倍</li>如果需要更细化的语速，可以保留小数点后 2 位，例如0.5 1.1 1.8等。<br>参数值与实际语速转换，可参考[代码示例](https://sdk-1300466766.cos.ap-shanghai.myqcloud.com/sample/speed_sample.tar.gz)
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// 项目id，用户自定义，默认为0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 模型类型，1-默认模型。
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// 音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色，请参见[购买指南](https://cloud.tencent.com/document/product/1073/34112)。完整的音色 ID 列表请参见[音色列表](https://cloud.tencent.com/document/product/1073/92668)。
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 主语言类型：<li>1-中文（默认）</li><li>2-英文</li><li>3-日文</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// 音频采样率：
	// <li>24000：24k（部分音色支持，请参见[音色列表](https://cloud.tencent.com/document/product/1073/92668)）</li>
	// <li>16000：16k（默认）</li>
	// <li>8000：8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 返回音频格式，可取值：wav（默认），mp3，pcm
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 是否开启时间戳功能，默认为false。
	EnableSubtitle *bool `json:"EnableSubtitle,omitempty" name:"EnableSubtitle"`

	// 断句敏感阈值，默认值为：0，取值范围：[0,1,2]。该值越大越不容易断句，模型会更倾向于仅按照标点符号断句。此参数建议不要随意调整，可能会影响合成效果。
	SegmentRate *uint64 `json:"SegmentRate,omitempty" name:"SegmentRate"`

	// 控制合成音频的情感，仅支持多情感音色使用。取值: neutral(中性)、sad(悲伤)、happy(高兴)、angry(生气)、fear(恐惧)、news(新闻)、story(故事)、radio(广播)、poetry(诗歌)、call(客服)、撒娇(sajiao)、厌恶(disgusted)、震惊(amaze)、平静(peaceful)、兴奋(exciting)、傲娇(aojiao)、解说(jieshuo)
	EmotionCategory *string `json:"EmotionCategory,omitempty" name:"EmotionCategory"`

	// 控制合成音频情感程度，取值范围为[50,200],默认为100；只有EmotionCategory不为空时生效；
	EmotionIntensity *int64 `json:"EmotionIntensity,omitempty" name:"EmotionIntensity"`
}

type TextToVoiceRequest struct {
	*tchttp.BaseRequest
	
	// 合成语音的源文本，按UTF-8编码统一计算。
	// 中文最大支持150个汉字（全角标点符号算一个汉字）；英文最大支持500个字母（半角标点符号算一个字母）。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 一次请求对应一个SessionId，会原样返回，建议传入类似于uuid的字符串防止重复。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 音量大小，范围[0，10]，对应音量大小。默认为0，代表正常音量，值越大音量越高。
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，6]，分别对应不同语速：<li>-2代表0.6倍</li><li>-1代表0.8倍</li><li>0代表1.0倍（默认）</li><li>1代表1.2倍</li><li>2代表1.5倍</li><li>6代表2.5倍</li>如果需要更细化的语速，可以保留小数点后 2 位，例如0.5 1.1 1.8等。<br>参数值与实际语速转换，可参考[代码示例](https://sdk-1300466766.cos.ap-shanghai.myqcloud.com/sample/speed_sample.tar.gz)
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// 项目id，用户自定义，默认为0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 模型类型，1-默认模型。
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// 音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色，请参见[购买指南](https://cloud.tencent.com/document/product/1073/34112)。完整的音色 ID 列表请参见[音色列表](https://cloud.tencent.com/document/product/1073/92668)。
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 主语言类型：<li>1-中文（默认）</li><li>2-英文</li><li>3-日文</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// 音频采样率：
	// <li>24000：24k（部分音色支持，请参见[音色列表](https://cloud.tencent.com/document/product/1073/92668)）</li>
	// <li>16000：16k（默认）</li>
	// <li>8000：8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 返回音频格式，可取值：wav（默认），mp3，pcm
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 是否开启时间戳功能，默认为false。
	EnableSubtitle *bool `json:"EnableSubtitle,omitempty" name:"EnableSubtitle"`

	// 断句敏感阈值，默认值为：0，取值范围：[0,1,2]。该值越大越不容易断句，模型会更倾向于仅按照标点符号断句。此参数建议不要随意调整，可能会影响合成效果。
	SegmentRate *uint64 `json:"SegmentRate,omitempty" name:"SegmentRate"`

	// 控制合成音频的情感，仅支持多情感音色使用。取值: neutral(中性)、sad(悲伤)、happy(高兴)、angry(生气)、fear(恐惧)、news(新闻)、story(故事)、radio(广播)、poetry(诗歌)、call(客服)、撒娇(sajiao)、厌恶(disgusted)、震惊(amaze)、平静(peaceful)、兴奋(exciting)、傲娇(aojiao)、解说(jieshuo)
	EmotionCategory *string `json:"EmotionCategory,omitempty" name:"EmotionCategory"`

	// 控制合成音频情感程度，取值范围为[50,200],默认为100；只有EmotionCategory不为空时生效；
	EmotionIntensity *int64 `json:"EmotionIntensity,omitempty" name:"EmotionIntensity"`
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
	delete(f, "Volume")
	delete(f, "Speed")
	delete(f, "ProjectId")
	delete(f, "ModelType")
	delete(f, "VoiceType")
	delete(f, "PrimaryLanguage")
	delete(f, "SampleRate")
	delete(f, "Codec")
	delete(f, "EnableSubtitle")
	delete(f, "SegmentRate")
	delete(f, "EmotionCategory")
	delete(f, "EmotionIntensity")
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

	// 时间戳信息，若未开启时间戳，则返回空数组。
	Subtitles []*Subtitle `json:"Subtitles,omitempty" name:"Subtitles"`

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