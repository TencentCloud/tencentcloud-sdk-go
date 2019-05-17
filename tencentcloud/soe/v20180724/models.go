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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 被评估语音对应的文本，句子模式下不超过个 20 单词或者中文文字，段落模式不超过 120 单词或者中文文字，中文评估使用 utf-8 编码，自由说模式该值传空。如需要在单词模式和句子模式下使用自定义音素，可以通过设置 TextMode 使用[音素标注](https://cloud.tencent.com/document/product/884/33698)。
	RefText *string `json:"RefText,omitempty" name:"RefText"`

	// 语音输入模式，0：流式分片，1：非流式一次性评估
	WorkMode *int64 `json:"WorkMode,omitempty" name:"WorkMode"`

	// 评估模式，0：词模式（中文评测模式下为文字模式），1：句子模式，2：段落模式，3：自由说模式，当为词模式评估时，能够提供每个音节的评估信息，当为句子模式时，能够提供完整度和流利度信息。
	EvalMode *int64 `json:"EvalMode,omitempty" name:"EvalMode"`

	// 评价苛刻指数，取值为[1.0 - 4.0]范围内的浮点数，用于平滑不同年龄段的分数，1.0为小年龄段，4.0为最高年龄段
	ScoreCoeff *float64 `json:"ScoreCoeff,omitempty" name:"ScoreCoeff"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数，新的 SoeAppId 可以在[控制台](https://console.cloud.tencent.com/soe)【应用管理】下新建。
	SoeAppId *string `json:"SoeAppId,omitempty" name:"SoeAppId"`

	// 长效session标识，当该参数为1时，session的持续时间为300s，但会一定程度上影响第一个数据包的返回速度，且TransmitOralProcess必须同时为1才可生效。
	IsLongLifeSession *int64 `json:"IsLongLifeSession,omitempty" name:"IsLongLifeSession"`

	// 音频存储模式，0：不存储，1：存储到公共对象存储，输出结果为该会话最后一个分片TransmitOralProcess 返回结果 AudioUrl 字段。
	StorageMode *int64 `json:"StorageMode,omitempty" name:"StorageMode"`

	// 输出断句中间结果标识，0：不输出，1：输出，通过设置该参数，可以在评估过程中的分片传输请求中，返回已经评估断句的中间结果，中间结果可用于客户端 UI 更新，输出结果为TransmitOralProcess请求返回结果 SentenceInfoSet 字段。
	SentenceInfoEnabled *int64 `json:"SentenceInfoEnabled,omitempty" name:"SentenceInfoEnabled"`

	// 评估语言，0：英文，1：中文。
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 异步模式标识，0：同步模式，1：异步模式，可选值参考[服务模式](https://cloud.tencent.com/document/product/884/33697)。
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`

	// 输入文本模式，0: 普通文本，1：[音素结构](https://cloud.tencent.com/document/product/884/33698)文本。
	TextMode *int64 `json:"TextMode,omitempty" name:"TextMode"`
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

		// 语音段唯一标识，一个完整语音一个SessionId
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	MemBeginTime *int64 `json:"MemBeginTime,omitempty" name:"MemBeginTime"`

	// 当前音节语音终止时间点，单位为ms
	MemEndTime *int64 `json:"MemEndTime,omitempty" name:"MemEndTime"`

	// 音节发音准确度，取值范围[-1, 100]，当取-1时指完全不匹配
	PronAccuracy *float64 `json:"PronAccuracy,omitempty" name:"PronAccuracy"`

	// 当前音节是否检测为重音
	DetectedStress *bool `json:"DetectedStress,omitempty" name:"DetectedStress"`

	// 当前音节
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 当前音节是否应为重音
	Stress *bool `json:"Stress,omitempty" name:"Stress"`
}

type SentenceInfo struct {

	// 句子序号，在段落、自由说模式下有效，表示断句序号，最后的综合结果的为-1.
	SentenceId *int64 `json:"SentenceId,omitempty" name:"SentenceId"`

	// 详细发音评估结果
	Words []*WordRsp `json:"Words,omitempty" name:"Words" list`

	// 发音精准度，取值范围[-1, 100]，当取-1时指完全不匹配，当为句子模式时，是所有已识别单词准确度的加权平均值，在reftext中但未识别出来的词不计入分数中。
	PronAccuracy *float64 `json:"PronAccuracy,omitempty" name:"PronAccuracy"`

	// 发音流利度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
	PronFluency *float64 `json:"PronFluency,omitempty" name:"PronFluency"`

	// 发音完整度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
	PronCompletion *float64 `json:"PronCompletion,omitempty" name:"PronCompletion"`

	// 建议评分，取值范围[0,100]，评分方式为建议评分 = 准确度（PronAccuracyfloat）* 完整度（PronCompletionfloat）*（2 - 完整度（PronCompletionfloat）），如若评分策略不符合请参考Words数组中的详细分数自定义评分逻辑。
	SuggestedScore *float64 `json:"SuggestedScore,omitempty" name:"SuggestedScore"`
}

type TransmitOralProcessRequest struct {
	*tchttp.BaseRequest

	// 流式数据包的序号，从1开始，当IsEnd字段为1后后续序号无意义，当IsLongLifeSession不为1且为非流式模式时无意义。
	SeqId *int64 `json:"SeqId,omitempty" name:"SeqId"`

	// 是否传输完毕标志，若为0表示未完毕，若为1则传输完毕开始评估，非流式模式下无意义。
	IsEnd *int64 `json:"IsEnd,omitempty" name:"IsEnd"`

	// 语音文件类型 	1:raw, 2:wav, 3:mp3(三种格式目前仅支持16k采样率16bit编码单声道，如有不一致可能导致评估不准确或失败)。
	VoiceFileType *int64 `json:"VoiceFileType,omitempty" name:"VoiceFileType"`

	// 语音编码类型	1:pcm。
	VoiceEncodeType *int64 `json:"VoiceEncodeType,omitempty" name:"VoiceEncodeType"`

	// 当前数据包数据, 流式模式下数据包大小可以按需设置，在网络稳定时，分片大小建议设置0.5k，且必须保证分片帧完整（16bit的数据必须保证音频长度为偶数），编码格式要求为BASE64。
	UserVoiceData *string `json:"UserVoiceData,omitempty" name:"UserVoiceData"`

	// 语音段唯一标识，一个完整语音一个SessionId。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数，新的 SoeAppId 可以在[控制台](https://console.cloud.tencent.com/soe)【应用管理】下新建。
	SoeAppId *string `json:"SoeAppId,omitempty" name:"SoeAppId"`

	// 长效session标识，当该参数为1时，session的持续时间为300s，但会一定程度上影响第一个数据包的返回速度。当InitOralProcess接口调用时此项为1时，此项必填1才可生效。
	IsLongLifeSession *int64 `json:"IsLongLifeSession,omitempty" name:"IsLongLifeSession"`

	// 查询标识，当该参数为1时，该请求为查询请求，请求返回该 Session 的评估结果。
	IsQuery *int64 `json:"IsQuery,omitempty" name:"IsQuery"`
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

		// 发音精准度，取值范围[-1, 100]，当取-1时指完全不匹配，当为句子模式时，是所有已识别单词准确度的加权平均值，在reftext中但未识别出来的词不计入分数中。当为流式模式且请求中IsEnd未置1时，取值无意义。
		PronAccuracy *float64 `json:"PronAccuracy,omitempty" name:"PronAccuracy"`

		// 发音流利度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
		PronFluency *float64 `json:"PronFluency,omitempty" name:"PronFluency"`

		// 发音完整度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
		PronCompletion *float64 `json:"PronCompletion,omitempty" name:"PronCompletion"`

		// 详细发音评估结果
		Words []*WordRsp `json:"Words,omitempty" name:"Words" list`

		// 语音段唯一标识，一段语音一个SessionId
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// 保存语音音频文件下载地址
		AudioUrl *string `json:"AudioUrl,omitempty" name:"AudioUrl"`

		// 断句中间结果，中间结果是局部最优而非全局最优的结果，所以中间结果有可能和最终整体结果对应部分不一致；中间结果的输出便于客户端UI更新；待用户发音完全结束后，系统会给出一个综合所有句子的整体结果。
		SentenceInfoSet []*SentenceInfo `json:"SentenceInfoSet,omitempty" name:"SentenceInfoSet" list`

		// 评估 session 状态，“Evaluating"：评估中、"Failed"：评估失败、"Finished"：评估完成
		Status *string `json:"Status,omitempty" name:"Status"`

		// 建议评分，取值范围[0,100]，评分方式为建议评分 = 准确度（PronAccuracyfloat）× 完整度（PronCompletionfloat）×（2 - 完整度（PronCompletionfloat）），如若评分策略不符合请参考Words数组中的详细分数自定义评分逻辑。
		SuggestedScore *float64 `json:"SuggestedScore,omitempty" name:"SuggestedScore"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransmitOralProcessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransmitOralProcessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransmitOralProcessWithInitRequest struct {
	*tchttp.BaseRequest

	// 流式数据包的序号，从1开始，当IsEnd字段为1后后续序号无意义，当IsLongLifeSession不为1且为非流式模式时无意义。
	SeqId *int64 `json:"SeqId,omitempty" name:"SeqId"`

	// 是否传输完毕标志，若为0表示未完毕，若为1则传输完毕开始评估，非流式模式下无意义。
	IsEnd *int64 `json:"IsEnd,omitempty" name:"IsEnd"`

	// 语音文件类型 	1:raw, 2:wav, 3:mp3(三种格式目前仅支持16k采样率16bit编码单声道，如有不一致可能导致评估不准确或失败)。
	VoiceFileType *int64 `json:"VoiceFileType,omitempty" name:"VoiceFileType"`

	// 语音编码类型	1:pcm。
	VoiceEncodeType *int64 `json:"VoiceEncodeType,omitempty" name:"VoiceEncodeType"`

	// 当前数据包数据, 流式模式下数据包大小可以按需设置，在网络良好的情况下，建议设置为0.5k，且必须保证分片帧完整（16bit的数据必须保证音频长度为偶数），编码格式要求为BASE64。
	UserVoiceData *string `json:"UserVoiceData,omitempty" name:"UserVoiceData"`

	// 语音段唯一标识，一个完整语音一个SessionId。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 被评估语音对应的文本，句子模式下不超过个 20 单词或者中文文字，段落模式不超过 120 单词或者中文文字，中文评估使用 utf-8 编码，自由说模式该值无效。如需要在单词模式和句子模式下使用自定义音素，可以通过设置 TextMode 使用[音素标注](https://cloud.tencent.com/document/product/884/33698)。
	RefText *string `json:"RefText,omitempty" name:"RefText"`

	// 语音输入模式，0：流式分片，1：非流式一次性评估
	WorkMode *int64 `json:"WorkMode,omitempty" name:"WorkMode"`

	// 评估模式，0：词模式（中文评测模式下为文字模式），1：句子模式，2：段落模式，3：自由说模式，当为词模式评估时，能够提供每个音节的评估信息，当为句子模式时，能够提供完整度和流利度信息。
	EvalMode *int64 `json:"EvalMode,omitempty" name:"EvalMode"`

	// 评价苛刻指数，取值为[1.0 - 4.0]范围内的浮点数，用于平滑不同年龄段的分数，1.0为小年龄段，4.0为最高年龄段
	ScoreCoeff *float64 `json:"ScoreCoeff,omitempty" name:"ScoreCoeff"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数，新的 SoeAppId 可以在[控制台](https://console.cloud.tencent.com/soe)【应用管理】下新建。
	SoeAppId *string `json:"SoeAppId,omitempty" name:"SoeAppId"`

	// 音频存储模式，0：不存储，1：存储到公共对象存储，输出结果为该会话最后一个分片TransmitOralProcess 返回结果 AudioUrl 字段。
	StorageMode *int64 `json:"StorageMode,omitempty" name:"StorageMode"`

	// 输出断句中间结果标识，0：不输出，1：输出，通过设置该参数，可以在评估过程中的分片传输请求中，返回已经评估断句的中间结果，中间结果可用于客户端 UI 更新，输出结果为TransmitOralProcess请求返回结果 SentenceInfoSet 字段。
	SentenceInfoEnabled *int64 `json:"SentenceInfoEnabled,omitempty" name:"SentenceInfoEnabled"`

	// 评估语言，0：英文，1：中文。
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 异步模式标识，0：同步模式，1：异步模式，可选值参考[服务模式](https://cloud.tencent.com/document/product/884/33697)。
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`

	// 查询标识，当该参数为1时，该请求为查询请求，请求返回该 Session 评估结果。
	IsQuery *int64 `json:"IsQuery,omitempty" name:"IsQuery"`

	// 输入文本模式，0: 普通文本，1：[音素结构](https://cloud.tencent.com/document/product/884/33698)文本。
	TextMode *int64 `json:"TextMode,omitempty" name:"TextMode"`
}

func (r *TransmitOralProcessWithInitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransmitOralProcessWithInitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransmitOralProcessWithInitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发音精准度，取值范围[-1, 100]，当取-1时指完全不匹配，当为句子模式时，是所有已识别单词准确度的加权平均值，在reftext中但未识别出来的词不计入分数中。当为流式模式且请求中IsEnd未置1时，取值无意义。
		PronAccuracy *float64 `json:"PronAccuracy,omitempty" name:"PronAccuracy"`

		// 发音流利度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
		PronFluency *float64 `json:"PronFluency,omitempty" name:"PronFluency"`

		// 发音完整度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
		PronCompletion *float64 `json:"PronCompletion,omitempty" name:"PronCompletion"`

		// 详细发音评估结果
		Words []*WordRsp `json:"Words,omitempty" name:"Words" list`

		// 语音段唯一标识，一段语音一个SessionId
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// 保存语音音频文件下载地址
		AudioUrl *string `json:"AudioUrl,omitempty" name:"AudioUrl"`

		// 断句中间结果，中间结果是局部最优而非全局最优的结果，所以中间结果有可能和最终整体结果对应部分不一致；中间结果的输出便于客户端UI更新；待用户发音完全结束后，系统会给出一个综合所有句子的整体结果。
		SentenceInfoSet []*SentenceInfo `json:"SentenceInfoSet,omitempty" name:"SentenceInfoSet" list`

		// 评估 session 状态，“Evaluating"：评估中、"Failed"：评估失败、"Finished"：评估完成
		Status *string `json:"Status,omitempty" name:"Status"`

		// 建议评分，取值范围[0,100]，评分方式为建议评分 = 准确度（PronAccuracyfloat）× 完整度（PronCompletionfloat）×（2 - 完整度（PronCompletionfloat）），如若评分策略不符合请参考Words数组中的详细分数自定义评分逻辑。
		SuggestedScore *float64 `json:"SuggestedScore,omitempty" name:"SuggestedScore"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransmitOralProcessWithInitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransmitOralProcessWithInitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WordRsp struct {

	// 当前单词语音起始时间点，单位为ms
	MemBeginTime *int64 `json:"MemBeginTime,omitempty" name:"MemBeginTime"`

	// 当前单词语音终止时间点，单位为ms
	MemEndTime *int64 `json:"MemEndTime,omitempty" name:"MemEndTime"`

	// 单词发音准确度，取值范围[-1, 100]，当取-1时指完全不匹配
	PronAccuracy *float64 `json:"PronAccuracy,omitempty" name:"PronAccuracy"`

	// 单词发音流利度，取值范围[0, 1]
	PronFluency *float64 `json:"PronFluency,omitempty" name:"PronFluency"`

	// 当前词
	Word *string `json:"Word,omitempty" name:"Word"`

	// 当前词与输入语句的匹配情况，0：匹配单词、1：新增单词、2：缺少单词、3：错读的词、4：未录入单词。
	MatchTag *int64 `json:"MatchTag,omitempty" name:"MatchTag"`

	// 音节评估详情
	PhoneInfos []*PhoneInfo `json:"PhoneInfos,omitempty" name:"PhoneInfos" list`
}
