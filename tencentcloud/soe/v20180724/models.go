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

package v20180724

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type InitOralProcessRequest struct {
	*tchttp.BaseRequest
	// 语音段唯一标识，一段语音一个SessionId
	SessionId *string `json:"SessionId" name:"SessionId"`
	// 被评估语音对应的文本
	RefText *string `json:"RefText" name:"RefText"`
	// 语音输入模式，0流式分片，1非流式一次性评估
	WorkMode *int64 `json:"WorkMode" name:"WorkMode"`
	// 评估模式，0:词模式, 1:句子模式，当为词模式评估时，能够提供每个音节的评估信息，当为句子模式时，能够提供完整度和流利度信息。
	EvalMode *int64 `json:"EvalMode" name:"EvalMode"`
	// 评价苛刻指数，取值为[1.0 - 4.0]范围内的浮点数，用于平滑不同年龄段的分数，1.0为小年龄段，4.0为最高年龄段
	ScoreCoeff *float64 `json:"ScoreCoeff" name:"ScoreCoeff"`
}

func (r *InitOralProcessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitOralProcessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitOralProcessResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitOralProcessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitOralProcessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PhoneInfo struct {
	// 当前音节语音起始时间点，单位为ms
	MemBeginTime *int64 `json:"MemBeginTime" name:"MemBeginTime"`
	// 当前音节语音终止时间点，单位为ms
	MemEndTime *int64 `json:"MemEndTime" name:"MemEndTime"`
	// 音节发音准确度，取值范围[-1, 100]，当取-1时指完全不匹配
	PronAccuracy *float64 `json:"PronAccuracy" name:"PronAccuracy"`
	// 当前音节是否检测为重音
	DetectedStress *bool `json:"DetectedStress" name:"DetectedStress"`
	// 当前音节
	Phone *string `json:"Phone" name:"Phone"`
	// 当前音节是否应为重音
	Stress *bool `json:"Stress" name:"Stress"`
}

type TransmitOralProcessRequest struct {
	*tchttp.BaseRequest
	// 流式数据包的序号，从1开始，当IsEnd字段为1后后续序号无意义，非流式模式下无意义
	SeqId *int64 `json:"SeqId" name:"SeqId"`
	// 是否传输完毕标志，若为0表示未完毕，若为1则传输完毕开始评估，非流式模式下无意义
	IsEnd *int64 `json:"IsEnd" name:"IsEnd"`
	// 语音文件类型 	1:raw, 2:wav, 3:mp3(mp3格式目前仅支持16k采样率16bit编码单声道)
	VoiceFileType *int64 `json:"VoiceFileType" name:"VoiceFileType"`
	// 语音编码类型	1:pcm
	VoiceEncodeType *int64 `json:"VoiceEncodeType" name:"VoiceEncodeType"`
	// 当前数据包数据, 流式模式下数据包大小可以按需设置，数据包大小必须 >= 4K，编码格式要求为BASE64。
	UserVoiceData *string `json:"UserVoiceData" name:"UserVoiceData"`
	// 语音段唯一标识，一个完整语音一个SessionId
	SessionId *string `json:"SessionId" name:"SessionId"`
}

func (r *TransmitOralProcessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransmitOralProcessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransmitOralProcessResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 发音精准度，取值范围[-1, 100]，当取-1时指完全不匹配，当为句子模式时，是所有已识别单词准确度的加权平均值。当为流式模式且请求中IsEnd未置1时，取值无意义
		PronAccuracy *float64 `json:"PronAccuracy" name:"PronAccuracy"`
		// 发音流利度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
		PronFluency *float64 `json:"PronFluency" name:"PronFluency"`
		// 发音完整度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
		PronCompletion *float64 `json:"PronCompletion" name:"PronCompletion"`
		// 详细发音评估结果
		Words []*WordRsp `json:"Words" name:"Words" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransmitOralProcessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransmitOralProcessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WordRsp struct {
	// 当前单词语音起始时间点，单位为ms
	MemBeginTime *int64 `json:"MemBeginTime" name:"MemBeginTime"`
	// 当前单词语音终止时间点，单位为ms
	MemEndTime *int64 `json:"MemEndTime" name:"MemEndTime"`
	// 单词发音准确度，取值范围[-1, 100]，当取-1时指完全不匹配
	PronAccuracy *float64 `json:"PronAccuracy" name:"PronAccuracy"`
	// 单词发音流利度，取值范围[0, 1]
	PronFluency *float64 `json:"PronFluency" name:"PronFluency"`
	// 当前词
	Word *string `json:"Word" name:"Word"`
	// 当前词与输入语句的匹配情况，0:匹配单词、1：新增单词、2：缺少单词
	MatchTag *int64 `json:"MatchTag" name:"MatchTag"`
	// 音节评估详情
	PhoneInfos []*PhoneInfo `json:"PhoneInfos" name:"PhoneInfos" list`
}
