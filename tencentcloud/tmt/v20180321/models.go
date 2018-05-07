// Copyright 1999-2018 Tencent Ltd.
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

package v20180321

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type LanguageDetectRequest struct {
	*tchttp.BaseRequest
	// 待识别的文本
	Text *string `json:"Text" name:"Text"`
	// 项目id
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
}

func (r *LanguageDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LanguageDetectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LanguageDetectResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 识别出的语言种类，参考语言列表
	// <li> zh : 中文 </li> <li> en : 英文 </li><li> jp : 日语 </li> <li> kr : 韩语 </li><li> de : 德语 </li><li> fr : 法语 </li><li> es : 西班牙文 </li> <li> it : 意大利文 </li><li> tr : 土耳其文 </li><li> ru : 俄文 </li><li> pt : 葡萄牙文 </li><li> vi : 越南文 </li><li> id : 印度尼西亚文 </li><li> ms : 马来西亚文 </li><li> th : 泰文 </li>
		Lang *string `json:"Lang" name:"Lang"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *LanguageDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LanguageDetectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SpeechTranslateRequest struct {
	*tchttp.BaseRequest
	// 一段完整的语⾳音对应⼀一个sessionUuid
	SessionUuid *string `json:"SessionUuid" name:"SessionUuid"`
	// 音频中的语⾔言类型，支持语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Source *string `json:"Source" name:"Source"`
	// 翻译⽬目标语⾔言类型 ，支持的语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Target *string `json:"Target" name:"Target"`
	// pcm : 146   amr : 33554432   mp3 : 83886080
	AudioFormat *int64 `json:"AudioFormat" name:"AudioFormat"`
	// 语⾳音分⽚片后的第⼏几⽚片
	Seq *int64 `json:"Seq" name:"Seq"`
	// 是否最后⼀片
	IsEnd *int64 `json:"IsEnd" name:"IsEnd"`
	// 语⾳音分⽚片内容的 base64字符串串
	Data *string `json:"Data" name:"Data"`
}

func (r *SpeechTranslateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SpeechTranslateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SpeechTranslateResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *SpeechTranslateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SpeechTranslateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextTranslateRequest struct {
	*tchttp.BaseRequest
	// 待翻译的文本
	SourceText *string `json:"SourceText" name:"SourceText"`
	// 源语言，参照Target支持语言列表
	Source *string `json:"Source" name:"Source"`
	// 目标语言，参照支持语言列表
	// <li> zh : 中文 </li> <li> en : 英文 </li><li> jp : 日语 </li> <li> kr : 韩语 </li><li> de : 德语 </li><li> fr : 法语 </li><li> es : 西班牙文 </li> <li> it : 意大利文 </li><li> tr : 土耳其文 </li><li> ru : 俄文 </li><li> pt : 葡萄牙文 </li><li> vi : 越南文 </li><li> id : 印度尼西亚文 </li><li> ms : 马来西亚文 </li><li> th : 泰文 </li><li> auto : 自动识别源语言，只能用于source字段 </li>
	Target *string `json:"Target" name:"Target"`
	// 项目id
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
}

func (r *TextTranslateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextTranslateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextTranslateResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 翻译后的文本
		TargetText *string `json:"TargetText" name:"TargetText"`
		// 源语言，详见入参Target
		Source *string `json:"Source" name:"Source"`
		// 目标语言，详见入参Target
		Target *string `json:"Target" name:"Target"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextTranslateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextTranslateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
