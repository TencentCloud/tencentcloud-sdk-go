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

package v20240523

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

// Predefined struct for user
type ConfirmVideoTranslateJobRequestParams struct {
	// 视频翻译任务 ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 待确认文本
	TranslateResults []*TranslateResult `json:"TranslateResults,omitnil,omitempty" name:"TranslateResults"`
}

type ConfirmVideoTranslateJobRequest struct {
	*tchttp.BaseRequest
	
	// 视频翻译任务 ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 待确认文本
	TranslateResults []*TranslateResult `json:"TranslateResults,omitnil,omitempty" name:"TranslateResults"`
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
type DescribeImageAnimateJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeImageAnimateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeImageAnimateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAnimateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageAnimateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAnimateJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 结果视频URL。有效期 24 小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 掩码视频链接
	MaskVideoUrl *string `json:"MaskVideoUrl,omitnil,omitempty" name:"MaskVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageAnimateJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageAnimateJobResponseParams `json:"Response"`
}

func (r *DescribeImageAnimateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAnimateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePortraitSingJobRequestParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribePortraitSingJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribePortraitSingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePortraitSingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePortraitSingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePortraitSingJobResponseParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务状态码
	// —RUN：处理中
	// —FAIL：处理失败
	// —STOP：处理终止
	// —DONE：处理完成
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 任务状态信息
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// 生成视频的URL地址
	// 有效期24小时
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePortraitSingJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribePortraitSingJobResponseParams `json:"Response"`
}

func (r *DescribePortraitSingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePortraitSingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoStylizationJobRequestParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeVideoStylizationJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeVideoStylizationJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoStylizationJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoStylizationJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoStylizationJobResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务状态码：
	// JobInit:  "初始化中"
	// JobModerationFailed: "审核失败",
	// JobRunning: "处理中",
	// JobFailed: "处理失败",
	// JobSuccess: "处理完成"。
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 任务状态描述。
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// 处理结果视频Url。URL有效期为24小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoStylizationJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoStylizationJobResponseParams `json:"Response"`
}

func (r *DescribeVideoStylizationJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoStylizationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoTranslateJobRequestParams struct {
	// 视频转译任务 ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeVideoTranslateJobRequest struct {
	*tchttp.BaseRequest
	
	// 视频转译任务 ID
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

	// 音频审核任务 ID	
	JobAudioModerationId *string `json:"JobAudioModerationId,omitnil,omitempty" name:"JobAudioModerationId"`

	// 口型驱动任务 ID	
	JobVideoId *string `json:"JobVideoId,omitnil,omitempty" name:"JobVideoId"`

	// 视频素材原始 URL	
	OriginalVideoUrl *string `json:"OriginalVideoUrl,omitnil,omitempty" name:"OriginalVideoUrl"`

	// 文本片段及其时间戳	
	AsrTimestamps []*AsrTimestamps `json:"AsrTimestamps,omitnil,omitempty" name:"AsrTimestamps"`

	// 提交视频翻译任务时的 requestId	
	JobSubmitReqId *string `json:"JobSubmitReqId,omitnil,omitempty" name:"JobSubmitReqId"`

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
type SubmitImageAnimateJobRequestParams struct {
	// 图片格式：支持PNG、JPG、JPEG格式；
	// 图片分辨率：长边分辨率不超过2056；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 图片base64数据。图片格式：支持PNG、JPG、JPEG格式；图片分辨率：长边分辨率不超过2056；图片大小：不超过10M；图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 动作模板ID。取值说明：ke3 科目三；tuziwu 兔子舞；huajiangwu 划桨舞。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 结果视频是否保留模板音频。默认为true
	EnableAudio *bool `json:"EnableAudio,omitnil,omitempty" name:"EnableAudio"`

	// 是否检测输入图人体12个身体部位（头部、颈部、右肩、右肘、右腕、左肩、左肘、左腕、右髋、左髋,、左膝、右膝）。默认不检测。
	EnableBodyJoins *bool `json:"EnableBodyJoins,omitnil,omitempty" name:"EnableBodyJoins"`

	// 最终视频是否保留原图的背景（该模式对于tuziwu、huajiangwu不生效）
	EnableSegment *bool `json:"EnableSegment,omitnil,omitempty" name:"EnableSegment"`
}

type SubmitImageAnimateJobRequest struct {
	*tchttp.BaseRequest
	
	// 图片格式：支持PNG、JPG、JPEG格式；
	// 图片分辨率：长边分辨率不超过2056；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 图片base64数据。图片格式：支持PNG、JPG、JPEG格式；图片分辨率：长边分辨率不超过2056；图片大小：不超过10M；图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 动作模板ID。取值说明：ke3 科目三；tuziwu 兔子舞；huajiangwu 划桨舞。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 结果视频是否保留模板音频。默认为true
	EnableAudio *bool `json:"EnableAudio,omitnil,omitempty" name:"EnableAudio"`

	// 是否检测输入图人体12个身体部位（头部、颈部、右肩、右肘、右腕、左肩、左肘、左腕、右髋、左髋,、左膝、右膝）。默认不检测。
	EnableBodyJoins *bool `json:"EnableBodyJoins,omitnil,omitempty" name:"EnableBodyJoins"`

	// 最终视频是否保留原图的背景（该模式对于tuziwu、huajiangwu不生效）
	EnableSegment *bool `json:"EnableSegment,omitnil,omitempty" name:"EnableSegment"`
}

func (r *SubmitImageAnimateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageAnimateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "TemplateId")
	delete(f, "EnableAudio")
	delete(f, "EnableBodyJoins")
	delete(f, "EnableSegment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitImageAnimateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitImageAnimateJobResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitImageAnimateJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitImageAnimateJobResponseParams `json:"Response"`
}

func (r *SubmitImageAnimateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageAnimateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitPortraitSingJobRequestParams struct {
	// 传入音频URL地址。音频要求：
	// —音频时长：不超过60秒
	// —音频格式：mp3、wav、m4a
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 传入图片URL地址，图片要求：
	// —图片格式：jpg、jpeg、png
	// —图片分辨率：长边不超过2560
	// —图片大小：不超过6M
	// —图片宽高比：图片【宽：高】在1:2到2:1范围内
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 传入图片Base64编码。
	// —图片Base64编码与URL地址必传其一
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`
}

type SubmitPortraitSingJobRequest struct {
	*tchttp.BaseRequest
	
	// 传入音频URL地址。音频要求：
	// —音频时长：不超过60秒
	// —音频格式：mp3、wav、m4a
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 传入图片URL地址，图片要求：
	// —图片格式：jpg、jpeg、png
	// —图片分辨率：长边不超过2560
	// —图片大小：不超过6M
	// —图片宽高比：图片【宽：高】在1:2到2:1范围内
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 传入图片Base64编码。
	// —图片Base64编码与URL地址必传其一
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`
}

func (r *SubmitPortraitSingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitPortraitSingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AudioUrl")
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitPortraitSingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitPortraitSingJobResponseParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitPortraitSingJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitPortraitSingJobResponseParams `json:"Response"`
}

func (r *SubmitPortraitSingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitPortraitSingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoStylizationJobRequestParams struct {
	// 风格ID，取值说明：2d_anime 2D动漫；3d_cartoon 3D卡通；3d_china 3D国潮；pixel_art	像素风。
	StyleId *string `json:"StyleId,omitnil,omitempty" name:"StyleId"`

	// 输入视频URL。视频要求：
	// - 视频格式：mp4、mov；
	// - 视频时长：1～60秒；
	// - 视频分辨率：540P~2056P，即长宽像素数均在540px～2056px范围内；
	// - 视频大小：不超过200M；
	// - 视频FPS：15～60fps。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 风格化强度 可选参数["low","medium","high"] 
	// "low":风格化强度弱,"medium":"风格化强度中等","high":"风格化强度强" 
	// 默认为medium
	StyleStrength *string `json:"StyleStrength,omitnil,omitempty" name:"StyleStrength"`
}

type SubmitVideoStylizationJobRequest struct {
	*tchttp.BaseRequest
	
	// 风格ID，取值说明：2d_anime 2D动漫；3d_cartoon 3D卡通；3d_china 3D国潮；pixel_art	像素风。
	StyleId *string `json:"StyleId,omitnil,omitempty" name:"StyleId"`

	// 输入视频URL。视频要求：
	// - 视频格式：mp4、mov；
	// - 视频时长：1～60秒；
	// - 视频分辨率：540P~2056P，即长宽像素数均在540px～2056px范围内；
	// - 视频大小：不超过200M；
	// - 视频FPS：15～60fps。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 风格化强度 可选参数["low","medium","high"] 
	// "low":风格化强度弱,"medium":"风格化强度中等","high":"风格化强度强" 
	// 默认为medium
	StyleStrength *string `json:"StyleStrength,omitnil,omitempty" name:"StyleStrength"`
}

func (r *SubmitVideoStylizationJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoStylizationJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StyleId")
	delete(f, "VideoUrl")
	delete(f, "StyleStrength")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoStylizationJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoStylizationJobResponseParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitVideoStylizationJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitVideoStylizationJobResponseParams `json:"Response"`
}

func (r *SubmitVideoStylizationJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoStylizationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoTranslateJobRequestParams struct {
	// 视频地址URL。
	// 格式要求：支持 mp4、mov 。
	// 时长要求：【10-300】秒。
	// fps 要求：【15-60】fps
	// 分辨率要求：单边像素要求在 【540~1920】 之间。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 源语言：zh(中文), en(英文)
	SrcLang *string `json:"SrcLang,omitnil,omitempty" name:"SrcLang"`

	// 目标语种：
	// zh(简体中文)、en(英语)、ar(阿拉伯语)、de(德语)、es(西班牙语)、fr(法语)、id(印尼语)、it(意大利语)、ja(日语)、ko(韩语)、ms(马来语)、pt(葡萄牙语)、ru(俄语)、th(泰语)、tr(土耳其语)、vi(越南语)
	DstLang *string `json:"DstLang,omitnil,omitempty" name:"DstLang"`

	// 当音频 URL 不为空时，默认以音频驱动视频任务口型。
	// 格式要求：支持 mp3、m4a、acc、wav 格式。
	// 时长要求：【10~300】秒
	// 大小要求：不超过 100M。
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 是否需要确认翻译结果0：不需要，1：需要
	Confirm *int64 `json:"Confirm,omitnil,omitempty" name:"Confirm"`

	// 是否开启口型驱动，0：不开启，1：开启。默认开启。
	LipSync *int64 `json:"LipSync,omitnil,omitempty" name:"LipSync"`

	// 音色种别：一种音色种别对应一种不同区域的音色
	// 1）目标语种为小语种(非zh,en)时，该项为必填
	// 2）目标语种为zh,en时，该项为非必填，若填入，则对应填入的音色
	// 
	// 具体音色包含请见“支持音色种别列表”
	VoiceType *string `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`
}

type SubmitVideoTranslateJobRequest struct {
	*tchttp.BaseRequest
	
	// 视频地址URL。
	// 格式要求：支持 mp4、mov 。
	// 时长要求：【10-300】秒。
	// fps 要求：【15-60】fps
	// 分辨率要求：单边像素要求在 【540~1920】 之间。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 源语言：zh(中文), en(英文)
	SrcLang *string `json:"SrcLang,omitnil,omitempty" name:"SrcLang"`

	// 目标语种：
	// zh(简体中文)、en(英语)、ar(阿拉伯语)、de(德语)、es(西班牙语)、fr(法语)、id(印尼语)、it(意大利语)、ja(日语)、ko(韩语)、ms(马来语)、pt(葡萄牙语)、ru(俄语)、th(泰语)、tr(土耳其语)、vi(越南语)
	DstLang *string `json:"DstLang,omitnil,omitempty" name:"DstLang"`

	// 当音频 URL 不为空时，默认以音频驱动视频任务口型。
	// 格式要求：支持 mp3、m4a、acc、wav 格式。
	// 时长要求：【10~300】秒
	// 大小要求：不超过 100M。
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 是否需要确认翻译结果0：不需要，1：需要
	Confirm *int64 `json:"Confirm,omitnil,omitempty" name:"Confirm"`

	// 是否开启口型驱动，0：不开启，1：开启。默认开启。
	LipSync *int64 `json:"LipSync,omitnil,omitempty" name:"LipSync"`

	// 音色种别：一种音色种别对应一种不同区域的音色
	// 1）目标语种为小语种(非zh,en)时，该项为必填
	// 2）目标语种为zh,en时，该项为非必填，若填入，则对应填入的音色
	// 
	// 具体音色包含请见“支持音色种别列表”
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
	// 视频转译任务的Job id
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
	// 翻译源文字
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 翻译后文字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`
}