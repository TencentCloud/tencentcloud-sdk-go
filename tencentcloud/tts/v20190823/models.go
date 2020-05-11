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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type TextToVoiceRequest struct {
	*tchttp.BaseRequest

	// 合成语音的源文本，按UTF-8编码统一计算。
	// 中文最大支持110个汉字（全角标点符号算一个汉字）；英文最大支持350个字母（半角标点符号算一个字母）。
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

	// 普通音色<li>0-云小宁，亲和女声(默认)</li><li>1-云小奇，亲和男声</li><li>2-云小晚，成熟男声</li><li>4-云小叶，温暖女声</li><li>5-云小欣，情感女声</li><li>6-云小龙，情感男声</li><li>7-云小曼，客服女声</li><li>1000-智侠，情感男声</li><li>1001-智瑜，情感女声</li><li>1002-智聆，通用女声</li><li>1003-智美，客服女声</li><li>1050-WeJack，英文男声</li><li>1051-WeRose，英文女声</li>精品音色<br>精品音色拟真度更高，价格不同于普通音色，查看[购买指南](https://cloud.tencent.com/product/tts/pricing)<br><li>101000-智侠，情感男声（精品）</li><li>101001-智瑜，情感女声（精品）</li><li>101002-智聆，通用女声（精品）</li><li>101003-智美，客服女声（精品）</li><li>101050-WeJack，英文男声（精品）</li><li>101051-WeRose，英文女声（精品）</li>
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 主语言类型：<li>1-中文（默认）</li><li>2-英文</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// 音频采样率：<li>16000：16k（默认）</li><li>8000：8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 返回音频格式，可取值：wav（默认），mp3
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

func (r *TextToVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextToVoiceRequest) FromJsonString(s string) error {
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

func (r *TextToVoiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
