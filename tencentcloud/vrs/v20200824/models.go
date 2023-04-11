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

package v20200824

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateVRSTaskRequestParams struct {
	// 唯一请求 ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 音色名称
	VoiceName *string `json:"VoiceName,omitempty" name:"VoiceName"`

	// 音频采样率：
	// 
	// 16000：16k（默认）
	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 音色性别:
	// 
	// 1-male
	// 
	// 2-female
	VoiceGender *int64 `json:"VoiceGender,omitempty" name:"VoiceGender"`

	// 语言类型：
	// 
	// 1-中文（默认）
	VoiceLanguage *int64 `json:"VoiceLanguage,omitempty" name:"VoiceLanguage"`

	// 音频格式，音频类型(wav,mp3,aac,m4a)
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 音频ID集合
	AudioIdList []*string `json:"AudioIdList,omitempty" name:"AudioIdList"`

	// 回调 URL，用户自行搭建的用于接收识别结果的服务URL。如果用户使用轮询方式获取识别结果，则无需提交该参数。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

type CreateVRSTaskRequest struct {
	*tchttp.BaseRequest
	
	// 唯一请求 ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 音色名称
	VoiceName *string `json:"VoiceName,omitempty" name:"VoiceName"`

	// 音频采样率：
	// 
	// 16000：16k（默认）
	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 音色性别:
	// 
	// 1-male
	// 
	// 2-female
	VoiceGender *int64 `json:"VoiceGender,omitempty" name:"VoiceGender"`

	// 语言类型：
	// 
	// 1-中文（默认）
	VoiceLanguage *int64 `json:"VoiceLanguage,omitempty" name:"VoiceLanguage"`

	// 音频格式，音频类型(wav,mp3,aac,m4a)
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 音频ID集合
	AudioIdList []*string `json:"AudioIdList,omitempty" name:"AudioIdList"`

	// 回调 URL，用户自行搭建的用于接收识别结果的服务URL。如果用户使用轮询方式获取识别结果，则无需提交该参数。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

func (r *CreateVRSTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVRSTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "VoiceName")
	delete(f, "SampleRate")
	delete(f, "VoiceGender")
	delete(f, "VoiceLanguage")
	delete(f, "Codec")
	delete(f, "AudioIdList")
	delete(f, "CallbackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVRSTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateVRSTaskRespData struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CreateVRSTaskResponseParams struct {
	// 创建任务结果
	Data *CreateVRSTaskRespData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVRSTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateVRSTaskResponseParams `json:"Response"`
}

func (r *CreateVRSTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVRSTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVRSTaskStatusRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeVRSTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeVRSTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVRSTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVRSTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVRSTaskStatusRespData struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态码，0：任务等待，1：任务执行中，2：任务成功，3：任务失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 任务状态，waiting：任务等待，doing：任务执行中，success：任务成功，failed：任务失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusStr *string `json:"StatusStr,omitempty" name:"StatusStr"`

	// 音色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// 失败原因说明。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

// Predefined struct for user
type DescribeVRSTaskStatusResponseParams struct {
	// 声音复刻任务结果
	Data *DescribeVRSTaskStatusRespData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVRSTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVRSTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeVRSTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVRSTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectEnvAndSoundQualityRequestParams struct {
	// 标注文本信息 ID
	TextId *string `json:"TextId,omitempty" name:"TextId"`

	// 语音数据 要使用base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。
	AudioData *string `json:"AudioData,omitempty" name:"AudioData"`

	// 音频格式，音频类型(wav,mp3,aac,m4a)
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 1:环境检测 2:音质检测
	TypeId *int64 `json:"TypeId,omitempty" name:"TypeId"`

	// 音频采样率：
	// 
	// 16000：16k（默认）
	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
}

type DetectEnvAndSoundQualityRequest struct {
	*tchttp.BaseRequest
	
	// 标注文本信息 ID
	TextId *string `json:"TextId,omitempty" name:"TextId"`

	// 语音数据 要使用base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。
	AudioData *string `json:"AudioData,omitempty" name:"AudioData"`

	// 音频格式，音频类型(wav,mp3,aac,m4a)
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 1:环境检测 2:音质检测
	TypeId *int64 `json:"TypeId,omitempty" name:"TypeId"`

	// 音频采样率：
	// 
	// 16000：16k（默认）
	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
}

func (r *DetectEnvAndSoundQualityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectEnvAndSoundQualityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TextId")
	delete(f, "AudioData")
	delete(f, "Codec")
	delete(f, "TypeId")
	delete(f, "SampleRate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectEnvAndSoundQualityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectEnvAndSoundQualityResponseParams struct {
	// 检测结果
	Data *DetectionEnvAndSoundQualityRespData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectEnvAndSoundQualityResponse struct {
	*tchttp.BaseResponse
	Response *DetectEnvAndSoundQualityResponseParams `json:"Response"`
}

func (r *DetectEnvAndSoundQualityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectEnvAndSoundQualityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectionEnvAndSoundQualityRespData struct {
	// 音频ID （用于创建任务接口AudioIds）,环境检测该值为空，仅在音质检测情况下返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioId *string `json:"AudioId,omitempty" name:"AudioId"`

	// 检测code 
	// 
	// 0 表示当前语音通过
	// -1 表示检测失败，需要重试
	// -2 表示语音检测不通过，提示用户再重新录制一下（通常漏读，错读，或多读）
	// -3 表示语音中噪声较大，不通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionCode *int64 `json:"DetectionCode,omitempty" name:"DetectionCode"`

	// 检测提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionMsg *string `json:"DetectionMsg,omitempty" name:"DetectionMsg"`

	// 检测提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionTip []*Words `json:"DetectionTip,omitempty" name:"DetectionTip"`
}

// Predefined struct for user
type GetTrainingTextRequestParams struct {

}

type GetTrainingTextRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetTrainingTextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTrainingTextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTrainingTextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTrainingTextResponseParams struct {
	// 文本列表
	Data *TrainingTexts `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTrainingTextResponse struct {
	*tchttp.BaseResponse
	Response *GetTrainingTextResponseParams `json:"Response"`
}

func (r *GetTrainingTextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTrainingTextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrainingText struct {
	// 文本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextId *string `json:"TextId,omitempty" name:"TextId"`

	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type TrainingTexts struct {
	// 训练文本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingTextList []*TrainingText `json:"TrainingTextList,omitempty" name:"TrainingTextList"`
}

type Words struct {
	// 准确度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PronAccuracy *float64 `json:"PronAccuracy,omitempty" name:"PronAccuracy"`

	// 流畅度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PronFluency *float64 `json:"PronFluency,omitempty" name:"PronFluency"`

	// tag: 0: match, 1: insert, 2: delete, 3: replace, 4: oov, 5: unknown
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *int64 `json:"Tag,omitempty" name:"Tag"`

	// 字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Word *string `json:"Word,omitempty" name:"Word"`
}