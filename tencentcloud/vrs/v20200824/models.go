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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CancelVRSTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CancelVRSTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *CancelVRSTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelVRSTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelVRSTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelVRSTaskResponseParams struct {
	// 任务ID
	Data *CancelVRSTaskRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelVRSTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelVRSTaskResponseParams `json:"Response"`
}

func (r *CancelVRSTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelVRSTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelVRSTaskRsp struct {

}

// Predefined struct for user
type CreateVRSTaskRequestParams struct {
	// 唯一请求 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 音色名称
	VoiceName *string `json:"VoiceName,omitnil,omitempty" name:"VoiceName"`

	// 音色性别:
	// 
	// 1-male
	// 
	// 2-female
	VoiceGender *int64 `json:"VoiceGender,omitnil,omitempty" name:"VoiceGender"`

	// 语言类型：
	// 
	// 1-中文
	VoiceLanguage *int64 `json:"VoiceLanguage,omitnil,omitempty" name:"VoiceLanguage"`

	// 音频ID集合。（一句话声音复刻仅需填写一个音质检测接口返回的AudioId）
	AudioIdList []*string `json:"AudioIdList,omitnil,omitempty" name:"AudioIdList"`

	// 音频采样率：
	// 
	// 16000：16k
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频格式，音频类型(wav,mp3,aac,m4a)
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 回调 URL，用户自行搭建的用于接收结果的服务URL。如果用户使用轮询方式获取识别结果，则无需提交该参数。（注意：回调方式目前仅支持轻量版声音复刻）
	// 回调采用POST请求方式，Content-Type为application/json，回调数据格式如下:{"TaskId":"xxxxxxxxxxxxxx","Status":2,"StatusStr":"success","VoiceType":xxxxx,"ErrorMsg":""}
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 模型类型 1:在线 2:离线  默认为1
	ModelType *int64 `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 复刻类型。
	// 0 - 轻量版声音复刻（默认）；
	// 5 - 一句话声音复刻。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 校验音频ID。（仅基础版声音复刻使用）
	VPRAudioId *string `json:"VPRAudioId,omitnil,omitempty" name:"VPRAudioId"`
}

type CreateVRSTaskRequest struct {
	*tchttp.BaseRequest
	
	// 唯一请求 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 音色名称
	VoiceName *string `json:"VoiceName,omitnil,omitempty" name:"VoiceName"`

	// 音色性别:
	// 
	// 1-male
	// 
	// 2-female
	VoiceGender *int64 `json:"VoiceGender,omitnil,omitempty" name:"VoiceGender"`

	// 语言类型：
	// 
	// 1-中文
	VoiceLanguage *int64 `json:"VoiceLanguage,omitnil,omitempty" name:"VoiceLanguage"`

	// 音频ID集合。（一句话声音复刻仅需填写一个音质检测接口返回的AudioId）
	AudioIdList []*string `json:"AudioIdList,omitnil,omitempty" name:"AudioIdList"`

	// 音频采样率：
	// 
	// 16000：16k
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频格式，音频类型(wav,mp3,aac,m4a)
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 回调 URL，用户自行搭建的用于接收结果的服务URL。如果用户使用轮询方式获取识别结果，则无需提交该参数。（注意：回调方式目前仅支持轻量版声音复刻）
	// 回调采用POST请求方式，Content-Type为application/json，回调数据格式如下:{"TaskId":"xxxxxxxxxxxxxx","Status":2,"StatusStr":"success","VoiceType":xxxxx,"ErrorMsg":""}
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 模型类型 1:在线 2:离线  默认为1
	ModelType *int64 `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 复刻类型。
	// 0 - 轻量版声音复刻（默认）；
	// 5 - 一句话声音复刻。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 校验音频ID。（仅基础版声音复刻使用）
	VPRAudioId *string `json:"VPRAudioId,omitnil,omitempty" name:"VPRAudioId"`
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
	delete(f, "VoiceGender")
	delete(f, "VoiceLanguage")
	delete(f, "AudioIdList")
	delete(f, "SampleRate")
	delete(f, "Codec")
	delete(f, "CallbackUrl")
	delete(f, "ModelType")
	delete(f, "TaskType")
	delete(f, "VPRAudioId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVRSTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateVRSTaskRespData struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CreateVRSTaskResponseParams struct {
	// 创建任务结果
	Data *CreateVRSTaskRespData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeVRSTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态码，0：任务等待，1：任务执行中，2：任务成功，3：任务失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务状态，waiting：任务等待，doing：任务执行中，success：任务成功，failed：任务失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusStr *string `json:"StatusStr,omitnil,omitempty" name:"StatusStr"`

	// 音色id。（若为一句话复刻时，该值为固定值“200000000”）
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceType *int64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// 失败原因说明。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 任务过期时间。（当复刻类型为一句话复刻时展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 快速复刻音色ID。（当复刻类型为一句话复刻时展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FastVoiceType *string `json:"FastVoiceType,omitnil,omitempty" name:"FastVoiceType"`
}

// Predefined struct for user
type DescribeVRSTaskStatusResponseParams struct {
	// 声音复刻任务结果
	Data *DescribeVRSTaskStatusRespData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TextId *string `json:"TextId,omitnil,omitempty" name:"TextId"`

	// 语音数据 要使用base64编码(采用python语言时注意读取文件时需要转成base64字符串编码，例如：str(base64.b64encode(open("input.aac", mode="rb").read()), encoding='utf-8') )。
	AudioData *string `json:"AudioData,omitnil,omitempty" name:"AudioData"`

	// 1:环境检测 2:音质检测
	TypeId *int64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 音频格式，音频类型(wav,mp3,aac,m4a)
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 音频采样率。
	// 16000：16k（默认）；
	// 24000：24k（仅一句话声音复刻支持）；
	// 48000：48k（仅一句话声音复刻支持）。
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 复刻类型。
	// 0 - 轻量版声音复刻（默认）;
	// 5 - 一句话声音复刻。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type DetectEnvAndSoundQualityRequest struct {
	*tchttp.BaseRequest
	
	// 标注文本信息 ID
	TextId *string `json:"TextId,omitnil,omitempty" name:"TextId"`

	// 语音数据 要使用base64编码(采用python语言时注意读取文件时需要转成base64字符串编码，例如：str(base64.b64encode(open("input.aac", mode="rb").read()), encoding='utf-8') )。
	AudioData *string `json:"AudioData,omitnil,omitempty" name:"AudioData"`

	// 1:环境检测 2:音质检测
	TypeId *int64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 音频格式，音频类型(wav,mp3,aac,m4a)
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 音频采样率。
	// 16000：16k（默认）；
	// 24000：24k（仅一句话声音复刻支持）；
	// 48000：48k（仅一句话声音复刻支持）。
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 复刻类型。
	// 0 - 轻量版声音复刻（默认）;
	// 5 - 一句话声音复刻。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
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
	delete(f, "TypeId")
	delete(f, "Codec")
	delete(f, "SampleRate")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectEnvAndSoundQualityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectEnvAndSoundQualityResponseParams struct {
	// 检测结果
	Data *DetectionEnvAndSoundQualityRespData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	AudioId *string `json:"AudioId,omitnil,omitempty" name:"AudioId"`

	// 检测code 
	// 
	// 0 表示当前语音通过
	// -1 表示检测失败，需要重试
	// -2 表示语音检测不通过，提示用户再重新录制一下（通常漏读，错读，或多读）
	// -3 表示语音中噪声较大，不通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionCode *int64 `json:"DetectionCode,omitnil,omitempty" name:"DetectionCode"`

	// 检测提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionMsg *string `json:"DetectionMsg,omitnil,omitempty" name:"DetectionMsg"`

	// 检测提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionTip []*Words `json:"DetectionTip,omitnil,omitempty" name:"DetectionTip"`
}

// Predefined struct for user
type DownloadVRSModelRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DownloadVRSModelRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DownloadVRSModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadVRSModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadVRSModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadVRSModelResponseParams struct {
	// 响应
	Data *DownloadVRSModelRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DownloadVRSModelResponse struct {
	*tchttp.BaseResponse
	Response *DownloadVRSModelResponseParams `json:"Response"`
}

func (r *DownloadVRSModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadVRSModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadVRSModelRsp struct {
	// 模型cos地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 音色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceName *string `json:"VoiceName,omitnil,omitempty" name:"VoiceName"`

	// 音色性别:
	// 1-male
	// 2-female
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceGender *int64 `json:"VoiceGender,omitnil,omitempty" name:"VoiceGender"`

	// 语言类型：
	// 1-中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceLanguage *int64 `json:"VoiceLanguage,omitnil,omitempty" name:"VoiceLanguage"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type GetTrainingTextRequestParams struct {
	// 复刻类型。
	// 0 - 轻量版声音复刻（默认）;
	// 5 - 一句话声音复刻。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 音色场景。（仅支持一句话声音复刻，其余复刻类型不生效） 
	// 0 - 通用场景（默认）； 
	// 1 - 聊天场景； 
	// 2 - 阅读场景； 
	// 3 - 资讯播报场景。
	Domain *int64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 文本语种。（仅支持一句话声音复刻，其余复刻类型不生效） 
	// 1 - 中文（默认）。
	TextLanguage *int64 `json:"TextLanguage,omitnil,omitempty" name:"TextLanguage"`
}

type GetTrainingTextRequest struct {
	*tchttp.BaseRequest
	
	// 复刻类型。
	// 0 - 轻量版声音复刻（默认）;
	// 5 - 一句话声音复刻。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 音色场景。（仅支持一句话声音复刻，其余复刻类型不生效） 
	// 0 - 通用场景（默认）； 
	// 1 - 聊天场景； 
	// 2 - 阅读场景； 
	// 3 - 资讯播报场景。
	Domain *int64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 文本语种。（仅支持一句话声音复刻，其余复刻类型不生效） 
	// 1 - 中文（默认）。
	TextLanguage *int64 `json:"TextLanguage,omitnil,omitempty" name:"TextLanguage"`
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
	delete(f, "TaskType")
	delete(f, "Domain")
	delete(f, "TextLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTrainingTextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTrainingTextResponseParams struct {
	// 文本列表
	Data *TrainingTexts `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type GetVRSVoiceTypeInfoRequestParams struct {
	// 音色id。
	VoiceType *int64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// 0 - 除快速声音复刻外其他复刻类型（默认）； 5 - 快速声音复刻。 默认为0。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 快速复刻音色id。
	FastVoiceType *string `json:"FastVoiceType,omitnil,omitempty" name:"FastVoiceType"`
}

type GetVRSVoiceTypeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 音色id。
	VoiceType *int64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// 0 - 除快速声音复刻外其他复刻类型（默认）； 5 - 快速声音复刻。 默认为0。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 快速复刻音色id。
	FastVoiceType *string `json:"FastVoiceType,omitnil,omitempty" name:"FastVoiceType"`
}

func (r *GetVRSVoiceTypeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVRSVoiceTypeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VoiceType")
	delete(f, "TaskType")
	delete(f, "FastVoiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetVRSVoiceTypeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVRSVoiceTypeInfoResponseParams struct {
	// 音色信息
	Data *VoiceTypeInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetVRSVoiceTypeInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetVRSVoiceTypeInfoResponseParams `json:"Response"`
}

func (r *GetVRSVoiceTypeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVRSVoiceTypeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVRSVoiceTypesRequestParams struct {
	// 复刻类型。
	// 0 - 除快速声音复刻外其他复刻类型（默认）；
	// 5 - 一句话声音复刻。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type GetVRSVoiceTypesRequest struct {
	*tchttp.BaseRequest
	
	// 复刻类型。
	// 0 - 除快速声音复刻外其他复刻类型（默认）；
	// 5 - 一句话声音复刻。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

func (r *GetVRSVoiceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVRSVoiceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetVRSVoiceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVRSVoiceTypesResponseParams struct {
	// 复刻音色信息
	Data *VoiceTypeListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetVRSVoiceTypesResponse struct {
	*tchttp.BaseResponse
	Response *GetVRSVoiceTypesResponseParams `json:"Response"`
}

func (r *GetVRSVoiceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVRSVoiceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrainingText struct {
	// 文本ID
	// 当复刻类型为一句话声音复刻时，生成的TextId有效期为7天，且在成功创建一次复刻任务后失效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextId *string `json:"TextId,omitnil,omitempty" name:"TextId"`

	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type TrainingTexts struct {
	// 训练文本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingTextList []*TrainingText `json:"TrainingTextList,omitnil,omitempty" name:"TrainingTextList"`
}

type VoiceTypeInfo struct {
	// 音色id。（若为一句话复刻时，该值为固定值“200000000”）
	VoiceType *int64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// 音色名称
	VoiceName *string `json:"VoiceName,omitnil,omitempty" name:"VoiceName"`

	// 音色性别: 1-male 2-female
	VoiceGender *int64 `json:"VoiceGender,omitnil,omitempty" name:"VoiceGender"`

	// 复刻类型: 0-轻量版复刻 1-基础版复刻
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 复刻任务 ID
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 创建时间
	DateCreated *string `json:"DateCreated,omitnil,omitempty" name:"DateCreated"`

	// 部署状态。若已部署，则可通过语音合成接口调用该音色
	IsDeployed *bool `json:"IsDeployed,omitnil,omitempty" name:"IsDeployed"`

	// 任务过期时间。（当复刻类型为一句话复刻时展示）
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 快速复刻音色ID。（当复刻类型为一句话复刻时展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FastVoiceType *string `json:"FastVoiceType,omitnil,omitempty" name:"FastVoiceType"`
}

type VoiceTypeListData struct {
	// 音色信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceTypeList []*VoiceTypeInfo `json:"VoiceTypeList,omitnil,omitempty" name:"VoiceTypeList"`
}

type Words struct {
	// 准确度 (小于75则认为不合格)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PronAccuracy *float64 `json:"PronAccuracy,omitnil,omitempty" name:"PronAccuracy"`

	// 流畅度 (小于0.95则认为不合格)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PronFluency *float64 `json:"PronFluency,omitnil,omitempty" name:"PronFluency"`

	// tag: 
	// 0: match  匹配
	// 1: insert   多读
	// 2: delete  少读
	// 3: replace 错读
	// 4: oov  待评估字不在发音评估的词库
	// 5: unknown 未知错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *int64 `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`
}