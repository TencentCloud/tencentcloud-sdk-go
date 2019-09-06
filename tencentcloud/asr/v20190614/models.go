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

package v20190614

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateRecTaskRequest struct {
	*tchttp.BaseRequest

	// 引擎类型。8k_0：电话 8k 通用模型；16k_0：16k 通用模型；8k_6: 电话场景下单声道话者分离模型。
	EngineModelType *string `json:"EngineModelType,omitempty" name:"EngineModelType"`

	// 语音声道数。1：单声道；2：双声道（仅在电话 8k 通用模型下支持）。
	ChannelNum *uint64 `json:"ChannelNum,omitempty" name:"ChannelNum"`

	// 识别结果文本编码方式。0：UTF-8。
	ResTextFormat *uint64 `json:"ResTextFormat,omitempty" name:"ResTextFormat"`

	// 语音数据来源。0：语音 URL；1：语音数据（post body）。
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 回调 URL，用户自行搭建的用于接收识别结果的服务器地址， 长度小于2048字节。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 语音的URL地址，需要公网可下载。长度小于2048字节，当 source_type 值为 0 时须填写该字段，为 1 时不需要填写。注意：请确保录音文件时长在一个小时之内，否则可能识别失败。请保证文件的下载速度，否则可能下载失败。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 语音数据，当SourceType 值为1时必须填写，为0可不写。要base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。音频数据要小于5MB。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据长度，当 SourceType 值为1时必须填写，为0可不写（此数据长度为数据未进行base64编码时的数据长度）。
	DataLen *uint64 `json:"DataLen,omitempty" name:"DataLen"`
}

func (r *CreateRecTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRecTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRecTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录音文件识别的请求返回结果。
		Data *Task `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRecTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 从CreateRecTask接口获取的TaskId，用于获取任务状态与结果。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录音文件识别的请求返回结果。
		Data *TaskStatus `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceRecognitionRequest struct {
	*tchttp.BaseRequest

	// 腾讯云项目 ID，可填 0，总长度不超过 1024 字节。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子服务类型。2： 一句话识别。
	SubServiceType *uint64 `json:"SubServiceType,omitempty" name:"SubServiceType"`

	// 引擎类型。8k：电话 8k 中文普通话通用；16k：16k 中文普通话通用；16k_en：16k 英语；16k_ca：16k 粤语。
	EngSerViceType *string `json:"EngSerViceType,omitempty" name:"EngSerViceType"`

	// 语音数据来源。0：语音 URL；1：语音数据（post body）。
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 识别音频的音频格式。mp3、wav。
	VoiceFormat *string `json:"VoiceFormat,omitempty" name:"VoiceFormat"`

	// 用户端对此任务的唯一标识，用户自助生成，用于用户查找识别结果。
	UsrAudioKey *string `json:"UsrAudioKey,omitempty" name:"UsrAudioKey"`

	// 语音 URL，公网可下载。当 SourceType 值为 0（语音 URL上传） 时须填写该字段，为 1 时不填；URL 的长度大于 0，小于 2048，需进行urlencode编码。音频时间长度要小于60s。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 语音数据，当SourceType 值为1（本地语音数据上传）时必须填写，当SourceType 值为0（语音 URL上传）可不写。要使用base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。音频数据要小于600KB。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据长度，单位为字节。当 SourceType 值为1（本地语音数据上传）时必须填写，当 SourceType 值为0（语音 URL上传）可不写（此数据长度为数据未进行base64编码时的数据长度）。
	DataLen *int64 `json:"DataLen,omitempty" name:"DataLen"`
}

func (r *SentenceRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceRecognitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SentenceRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Task struct {

	// 任务ID，可通过此ID在轮询接口获取识别状态与结果
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type TaskStatus struct {

	// 任务标识。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态码，0：任务等待，1：任务执行中，2：任务成功，3：任务失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 任务状态，waiting：任务等待，doing：任务执行中，success：任务成功，failed：任务失败。
	StatusStr *string `json:"StatusStr,omitempty" name:"StatusStr"`

	// 识别结果。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 失败原因说明。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}
