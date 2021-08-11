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

type CreateTtsTaskRequest struct {
	*tchttp.BaseRequest

	// 合成语音的源文本，按UTF-8编码统一计算，最多支持10万字符
	Text *string `json:"Text,omitempty" name:"Text"`

	// 模型类型，1-默认模型。
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// 音量大小，范围：[0，10]，分别对应11个等级的音量，默认为0，代表正常音量。没有静音选项。
	// 输入除以上整数之外的其他参数不生效，按默认值处理。
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，2]，分别对应不同语速：<li>-2代表0.6倍</li><li>-1代表0.8倍</li><li>0代表1.0倍（默认）</li><li>1代表1.2倍</li><li>2代表1.5倍</li>如果需要更细化的语速，可以保留小数点后一位，例如0.5 1.1 1.8等。<br>
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// 项目id，用户自定义，默认为0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 标准音色<li>1001-智瑜，情感女声(默认)</li><li>1002-智聆，通用女声</li><li>1003-智美，客服女声</li><li>1006-智云，通用男声</li><li>1050-WeJack，英文男声</li><li>1051-WeRose，英文女声</li>精品音色<br>精品音色拟真度更高，价格不同于普通音色，查看[购买指南](https://cloud.tencent.com/product/tts/pricing)<br><li>100510000-智逍遥，旁对白阅读风格</li><li>101001-智瑜，情感女声（精品）</li><li>101002-智聆，通用女声（精品）</li><li>101003-智美，客服女声（精品）</li><li>101004-智云，通用男声</li><li>101005-智莉，通用女声</li><li>101006-智言，助手女声</li><li>101007-智娜，客服女声</li><li>101008-智琪，客服女声</li><li>101009-智芸，知性女声</li><li>101010-智华，通用男声</li><li>101011-智燕，新闻女声</li><li>101012-智丹，新闻女声</li><li>101013-智辉，新闻男声</li><li>101014-智宁，新闻男声</li><li>101015-智萌，男童声</li><li>101016-智甜，女童声</li><li>101017-智蓉，情感女声</li><li>101018-智靖，情感男声</li><li>101019-智彤，粤语女声</li><li>101050-WeJack，英文男声（精品）</li><li>101051-WeRose，英文女声（精品）</li>
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 主语言类型：<li>1-中文（默认）</li><li>2-英文</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// 音频采样率：<li>16000：16k（默认）</li><li>8000：8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 返回音频格式，可取值：mp3（默认），wav，pcm
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 回调 URL，用户自行搭建的用于接收识别结果的服务URL。如果用户使用轮询方式获取识别结果，则无需提交该参数。[回调说明](https://cloud.tencent.com/document/product/1073/55746)
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 旁白与对白文本解析，分别合成相应风格（仅适用于旁对白音色），默认 false
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

type CreateTtsTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 id
		Data *CreateTtsTaskRespData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 失败原因说明。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

type DescribeTtsTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务状态返回
		Data *DescribeTtsTaskStatusRespData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type TextToVoiceRequest struct {
	*tchttp.BaseRequest

	// 合成语音的源文本，按UTF-8编码统一计算。
	// 中文最大支持150个汉字（全角标点符号算一个汉字）；英文最大支持500个字母（半角标点符号算一个字母）。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 一次请求对应一个SessionId，会原样返回，建议传入类似于uuid的字符串防止重复。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 模型类型，1-默认模型。
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// 音量大小，范围：[0，10]，分别对应11个等级的音量，默认为0，代表正常音量。没有静音选项。
	// 输入除以上整数之外的其他参数不生效，按默认值处理。
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`

	// 语速，范围：[-2，2]，分别对应不同语速：<li>-2代表0.6倍</li><li>-1代表0.8倍</li><li>0代表1.0倍（默认）</li><li>1代表1.2倍</li><li>2代表1.5倍</li>如果需要更细化的语速，可以保留小数点后一位，例如0.5 1.1 1.8等。<br>
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// 项目id，用户自定义，默认为0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 普通音色<li>0-云小宁，亲和女声(默认)</li><li>1-云小奇，亲和男声</li><li>2-云小晚，成熟男声</li><li>4-云小叶，温暖女声</li><li>5-云小欣，情感女声</li><li>6-云小龙，情感男声</li><li>7-云小曼，客服女声</li><li>1001-智瑜，情感女声</li><li>1002-智聆，通用女声</li><li>1003-智美，客服女声</li><li>1050-WeJack，英文男声</li><li>1051-WeRose，英文女声</li>精品音色<br>精品音色拟真度更高，价格不同于普通音色，查看[购买指南](https://cloud.tencent.com/product/tts/pricing)<br><li>101001-智瑜，情感女声（精品）</li><li>101002-智聆，通用女声（精品）</li><li>101003-智美，客服女声（精品）</li><li>101004-智云，通用男声</li><li>101005-智莉，通用女声</li><li>101006-智言，助手女声</li><li>101007-智娜，客服女声</li><li>101008-智琪，客服女声</li><li>101009-智芸，知性女声</li><li>101010-智华，通用男声</li><li>101011-智燕，新闻女声</li><li>101012-智丹，新闻女声</li><li>101013-智辉，新闻男声</li><li>101014-智宁，新闻男声</li><li>101015-智萌，男童声</li><li>101016-智甜，女童声</li><li>101017-智蓉，情感女声</li><li>101018-智靖，情感男声</li><li>101019-智彤，粤语女声</li><li>101050-WeJack，英文男声（精品）</li><li>101051-WeRose，英文女声（精品）</li><li>102000-贝蕾，客服女声</li><li>102001-贝果，客服女声</li><li>102002-贝紫，粤语女声</li><li>102003-贝雪，新闻女声</li>
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 主语言类型：<li>1-中文（默认）</li><li>2-英文</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// 音频采样率：<li>16000：16k（默认）</li><li>8000：8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 返回音频格式，可取值：wav（默认），mp3，pcm
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

type TextToVoiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// base64编码的wav/mp3音频数据
		Audio *string `json:"Audio,omitempty" name:"Audio"`

		// 一次请求对应一个SessionId
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
