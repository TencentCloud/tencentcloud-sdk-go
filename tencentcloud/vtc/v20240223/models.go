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

package v20240223

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AsrTimestamps struct {
	// 文本片段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartMs *int64 `json:"StartMs,omitnil,omitempty" name:"StartMs"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndMs *int64 `json:"EndMs,omitnil,omitempty" name:"EndMs"`
}

type AudioTranslateResult struct {
	// 原文本
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 目标文本
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`
}

// Predefined struct for user
type ConfirmVideoTranslateJobRequestParams struct {
	// 视频翻译任务 ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 待确认文本
	TranslateResults []*AudioTranslateResult `json:"TranslateResults,omitnil,omitempty" name:"TranslateResults"`
}

type ConfirmVideoTranslateJobRequest struct {
	*tchttp.BaseRequest
	
	// 视频翻译任务 ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 待确认文本
	TranslateResults []*AudioTranslateResult `json:"TranslateResults,omitnil,omitempty" name:"TranslateResults"`
}

func (r *ConfirmVideoTranslateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmVideoTranslateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TranslateResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfirmVideoTranslateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfirmVideoTranslateJobResponseParams struct {
	// 视频翻译任务 ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 音频转换任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 音频翻译结果确认 session
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 视频转译任务状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 视频转译任务信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfirmVideoTranslateJobResponse struct {
	*tchttp.BaseResponse
	Response *ConfirmVideoTranslateJobResponseParams `json:"Response"`
}

func (r *ConfirmVideoTranslateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmVideoTranslateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoTranslateJobRequestParams struct {
	// JobId。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeVideoTranslateJobRequest struct {
	*tchttp.BaseRequest
	
	// JobId。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeVideoTranslateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoTranslateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoTranslateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoTranslateJobResponseParams struct {
	// 任务状态。 1：音频翻译中。 2：音频翻译失败。 3：音频翻译成功。 4：音频结果待确认。 5：音频结果已确认完毕。6：视频翻译中。 7：视频翻译失败。 8：视频翻译成功。
	JobStatus *int64 `json:"JobStatus,omitnil,omitempty" name:"JobStatus"`

	// 任务错误码。
	JobErrorCode *string `json:"JobErrorCode,omitnil,omitempty" name:"JobErrorCode"`

	// 任务错误信息。
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`

	// 视频翻译结果。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 音频翻译结果。
	TranslateResults []*TranslateResult `json:"TranslateResults,omitnil,omitempty" name:"TranslateResults"`

	// 是否需要确认翻译结果。0：不需要，1：需要
	JobConfirm *int64 `json:"JobConfirm,omitnil,omitempty" name:"JobConfirm"`

	// 音频任务 ID
	JobAudioTaskId *string `json:"JobAudioTaskId,omitnil,omitempty" name:"JobAudioTaskId"`

	// 视频审核任务ID
	JobVideoModerationId *string `json:"JobVideoModerationId,omitnil,omitempty" name:"JobVideoModerationId"`

	// 视频生成任务 ID
	JobVideoId *string `json:"JobVideoId,omitnil,omitempty" name:"JobVideoId"`

	// 视频素材原始 URL
	OriginalVideoUrl *string `json:"OriginalVideoUrl,omitnil,omitempty" name:"OriginalVideoUrl"`

	// 文本片段及其时间戳
	AsrTimestamps []*AsrTimestamps `json:"AsrTimestamps,omitnil,omitempty" name:"AsrTimestamps"`

	// 提交视频翻译任务时的 requestId
	JobSubmitReqId *string `json:"JobSubmitReqId,omitnil,omitempty" name:"JobSubmitReqId"`

	// 音频审核任务 ID
	JobAudioModerationId *string `json:"JobAudioModerationId,omitnil,omitempty" name:"JobAudioModerationId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoTranslateJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoTranslateJobResponseParams `json:"Response"`
}

func (r *DescribeVideoTranslateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoTranslateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoTranslateJobRequestParams struct {
	// 视频地址URL。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 源语言：zh, en
	SrcLang *string `json:"SrcLang,omitnil,omitempty" name:"SrcLang"`

	// 目标语言：zh, en
	DstLang *string `json:"DstLang,omitnil,omitempty" name:"DstLang"`

	// 当音频 URL 不为空时，默认以音频驱动视频任务口型
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 是否需要确认翻译结果0：不需要，1：需要
	Confirm *int64 `json:"Confirm,omitnil,omitempty" name:"Confirm"`

	// 是否开启口型驱动，0：不开启，1：开启。默认开启。
	LipSync *int64 `json:"LipSync,omitnil,omitempty" name:"LipSync"`

	// 音色 ID
	VoiceType *string `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`
}

type SubmitVideoTranslateJobRequest struct {
	*tchttp.BaseRequest
	
	// 视频地址URL。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 源语言：zh, en
	SrcLang *string `json:"SrcLang,omitnil,omitempty" name:"SrcLang"`

	// 目标语言：zh, en
	DstLang *string `json:"DstLang,omitnil,omitempty" name:"DstLang"`

	// 当音频 URL 不为空时，默认以音频驱动视频任务口型
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 是否需要确认翻译结果0：不需要，1：需要
	Confirm *int64 `json:"Confirm,omitnil,omitempty" name:"Confirm"`

	// 是否开启口型驱动，0：不开启，1：开启。默认开启。
	LipSync *int64 `json:"LipSync,omitnil,omitempty" name:"LipSync"`

	// 音色 ID
	VoiceType *string `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`
}

func (r *SubmitVideoTranslateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoTranslateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoUrl")
	delete(f, "SrcLang")
	delete(f, "DstLang")
	delete(f, "AudioUrl")
	delete(f, "Confirm")
	delete(f, "LipSync")
	delete(f, "VoiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoTranslateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoTranslateJobResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitVideoTranslateJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitVideoTranslateJobResponseParams `json:"Response"`
}

func (r *SubmitVideoTranslateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoTranslateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TranslateResult struct {
	// 翻译源文字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 翻译后文字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`
}