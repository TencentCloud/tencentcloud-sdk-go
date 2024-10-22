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

package v20210325

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type DisplayInfo struct {
	// 句子 ID
	SeId *string `json:"SeId,omitnil,omitempty" name:"SeId"`

	// 句子版本号
	SeVer *uint64 `json:"SeVer,omitnil,omitempty" name:"SeVer"`

	// 识别结果
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	//  翻译结果
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// 句子开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 句子结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	//  当前句子是否已结束
	IsEnd *bool `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// base64编码的wav/mp3音频数据
	Audio *string `json:"Audio,omitnil,omitempty" name:"Audio"`
}

type TTS struct {
	// 返回音频格式，可取值：wav，mp3，pcm
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 音色 ID，只包括标准音色（注，日文只有一个固定音色）。
	// 完整的音色 ID 列表请参见[音色列表](https://cloud.tencent.com/document/product/1073/92668)。
	VoiceType *uint64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// 音量大小，范围[-10，10]，对应音量大小。默认为0，代表正常音量，值越大音量越高。
	Volume *float64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 语速，范围：[-2，6]，分别对应不同语速：
	// 
	// - -2代表0.6倍
	// - -1代表0.8倍
	// - 0代表1.0倍（默认）
	// - 1代表1.2倍
	// - 2代表1.5倍
	// - 6代表2.5倍
	// 
	// 如果需要更细化的语速，可以保留小数点后 2 位，例如0.5/1.25/2.81等。
	// 参数值与实际语速转换，可参考[代码示例](https://sdk-1300466766.cos.ap-shanghai.myqcloud.com/sample/speed_sample.tar.gz)
	Speed *float64 `json:"Speed,omitnil,omitempty" name:"Speed"`

	// 音频采样率：
	// 
	// - 16000：16k（默认）
	// - 8000：8k
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`
}

// Predefined struct for user
type TongChuanDisplayRequestParams struct {
	// 一段完整的语音对应一个SessionUuid
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 句子排序方式，1-由新到旧
	IsNew *uint64 `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 最多返回几句，目前只支持 5 条数据
	SeMax *uint64 `json:"SeMax,omitnil,omitempty" name:"SeMax"`
}

type TongChuanDisplayRequest struct {
	*tchttp.BaseRequest
	
	// 一段完整的语音对应一个SessionUuid
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 句子排序方式，1-由新到旧
	IsNew *uint64 `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 最多返回几句，目前只支持 5 条数据
	SeMax *uint64 `json:"SeMax,omitnil,omitempty" name:"SeMax"`
}

func (r *TongChuanDisplayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TongChuanDisplayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionUuid")
	delete(f, "IsNew")
	delete(f, "SeMax")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TongChuanDisplayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TongChuanDisplayResponseParams struct {
	// 同传结果数组
	List []*DisplayInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TongChuanDisplayResponse struct {
	*tchttp.BaseResponse
	Response *TongChuanDisplayResponseParams `json:"Response"`
}

func (r *TongChuanDisplayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TongChuanDisplayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TongChuanRecognizeRequestParams struct {
	// 一段完整的语音对应一个SessionUuid
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 源语言，支持：
	// zh：中文
	// en：英语
	// ja：日语
	// ko：韩语
	// yue：粤语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// <li>zh（中文）：en（英语）、ja（日语）、ko（韩语）、yue（粤语）</li>
	// <li>en（英语）：zh（中文）</li>
	// <li>ja（日语）：zh（中文）</li>
	// <li>ko（韩语）：zh（中文）</li>
	// <li>yue（粤语）：zh（中文）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 语音编码类型，1-pcm
	AudioFormat *uint64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 语音分片的序号，从0开始
	Seq *uint64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 语音开始的时间戳
	Utc *uint64 `json:"Utc,omitnil,omitempty" name:"Utc"`

	// 是否最后一片语音分片，0-否，1-是
	IsEnd *uint64 `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// 翻译时机，0-不翻译 2-句子实时翻译
	TranslateTime *uint64 `json:"TranslateTime,omitnil,omitempty" name:"TranslateTime"`

	// 语音分片内容进行 Base64 编码后的字符串。音频内容需包含有效并可识别的文本信息。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// TTS播报控制参数	
	TTS *TTS `json:"TTS,omitnil,omitempty" name:"TTS"`
}

type TongChuanRecognizeRequest struct {
	*tchttp.BaseRequest
	
	// 一段完整的语音对应一个SessionUuid
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 源语言，支持：
	// zh：中文
	// en：英语
	// ja：日语
	// ko：韩语
	// yue：粤语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// <li>zh（中文）：en（英语）、ja（日语）、ko（韩语）、yue（粤语）</li>
	// <li>en（英语）：zh（中文）</li>
	// <li>ja（日语）：zh（中文）</li>
	// <li>ko（韩语）：zh（中文）</li>
	// <li>yue（粤语）：zh（中文）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 语音编码类型，1-pcm
	AudioFormat *uint64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 语音分片的序号，从0开始
	Seq *uint64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 语音开始的时间戳
	Utc *uint64 `json:"Utc,omitnil,omitempty" name:"Utc"`

	// 是否最后一片语音分片，0-否，1-是
	IsEnd *uint64 `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// 翻译时机，0-不翻译 2-句子实时翻译
	TranslateTime *uint64 `json:"TranslateTime,omitnil,omitempty" name:"TranslateTime"`

	// 语音分片内容进行 Base64 编码后的字符串。音频内容需包含有效并可识别的文本信息。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// TTS播报控制参数	
	TTS *TTS `json:"TTS,omitnil,omitempty" name:"TTS"`
}

func (r *TongChuanRecognizeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TongChuanRecognizeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionUuid")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "AudioFormat")
	delete(f, "Seq")
	delete(f, "Utc")
	delete(f, "IsEnd")
	delete(f, "TranslateTime")
	delete(f, "Data")
	delete(f, "TTS")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TongChuanRecognizeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TongChuanRecognizeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TongChuanRecognizeResponse struct {
	*tchttp.BaseResponse
	Response *TongChuanRecognizeResponseParams `json:"Response"`
}

func (r *TongChuanRecognizeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TongChuanRecognizeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TongChuanSyncRequestParams struct {
	// 一段完整的语音对应一个SessionUuid
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 源语言，支持：
	// zh：中文
	// en：英语
	// ja：日语
	// ko：韩语
	// yue：粤语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// <li>zh（中文）：en（英语）、ja（日语）、ko（韩语）、yue（粤语）</li>
	// <li>en（英语）：zh（中文）</li>
	// <li>ja（日语）：zh（中文）</li>
	// <li>ko（韩语）：zh（中文）</li>
	// <li>yue（粤语）：zh（中文）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 语音编码类型，1-pcm
	AudioFormat *uint64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 语音分片的序号，从0开始
	Seq *uint64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 语音开始的时间戳
	Utc *uint64 `json:"Utc,omitnil,omitempty" name:"Utc"`

	// 是否最后一片语音分片，0-否，1-是
	IsEnd *uint64 `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// 翻译时机，0-不翻译 2-句子实时翻译
	TranslateTime *uint64 `json:"TranslateTime,omitnil,omitempty" name:"TranslateTime"`

	// 语音分片内容进行 Base64 编码后的字符串。音频内容需包含有效并可识别的文本信息。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// TTS播报控制参数
	TTS *TTS `json:"TTS,omitnil,omitempty" name:"TTS"`
}

type TongChuanSyncRequest struct {
	*tchttp.BaseRequest
	
	// 一段完整的语音对应一个SessionUuid
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 源语言，支持：
	// zh：中文
	// en：英语
	// ja：日语
	// ko：韩语
	// yue：粤语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// <li>zh（中文）：en（英语）、ja（日语）、ko（韩语）、yue（粤语）</li>
	// <li>en（英语）：zh（中文）</li>
	// <li>ja（日语）：zh（中文）</li>
	// <li>ko（韩语）：zh（中文）</li>
	// <li>yue（粤语）：zh（中文）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 语音编码类型，1-pcm
	AudioFormat *uint64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 语音分片的序号，从0开始
	Seq *uint64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 语音开始的时间戳
	Utc *uint64 `json:"Utc,omitnil,omitempty" name:"Utc"`

	// 是否最后一片语音分片，0-否，1-是
	IsEnd *uint64 `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// 翻译时机，0-不翻译 2-句子实时翻译
	TranslateTime *uint64 `json:"TranslateTime,omitnil,omitempty" name:"TranslateTime"`

	// 语音分片内容进行 Base64 编码后的字符串。音频内容需包含有效并可识别的文本信息。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// TTS播报控制参数
	TTS *TTS `json:"TTS,omitnil,omitempty" name:"TTS"`
}

func (r *TongChuanSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TongChuanSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionUuid")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "AudioFormat")
	delete(f, "Seq")
	delete(f, "Utc")
	delete(f, "IsEnd")
	delete(f, "TranslateTime")
	delete(f, "Data")
	delete(f, "TTS")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TongChuanSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TongChuanSyncResponseParams struct {
	// 同传结果数组
	List []*DisplayInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TongChuanSyncResponse struct {
	*tchttp.BaseResponse
	Response *TongChuanSyncResponseParams `json:"Response"`
}

func (r *TongChuanSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TongChuanSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}