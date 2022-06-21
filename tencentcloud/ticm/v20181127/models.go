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

package v20181127

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Candidate struct {
	// 识别出人脸对应的候选人数组。当前返回相似度最高的候选人。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 相似度，0-100之间。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`
}

// Predefined struct for user
type DescribeVideoTaskRequestParams struct {
	// 需要查询的视频审核的任务ID
	VodTaskId *string `json:"VodTaskId,omitempty" name:"VodTaskId"`
}

type DescribeVideoTaskRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的视频审核的任务ID
	VodTaskId *string `json:"VodTaskId,omitempty" name:"VodTaskId"`
}

func (r *DescribeVideoTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VodTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoTaskResponseParams struct {
	// 任务状态，取值：
	// WAITING：等待中；
	// PROCESSING：处理中；
	// FINISH：已完成。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务开始执行的时间，采用 ISO 日期格式。
	BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 ISO 日期格式。
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 视频内容审核智能画面鉴黄任务的查询结果。
	PornResult *VodPornReviewResult `json:"PornResult,omitempty" name:"PornResult"`

	// 视频内容审核智能画面鉴恐任务的查询结果。
	TerrorismResult *VodTerrorismReviewResult `json:"TerrorismResult,omitempty" name:"TerrorismResult"`

	// 视频内容审核智能画面鉴政任务的查询结果。
	PoliticalResult *VodPoliticalReviewResult `json:"PoliticalResult,omitempty" name:"PoliticalResult"`

	// 视频内容审核 Ocr 文字鉴政任务的查询结果。
	PoliticalOcrResult *VodPoliticalOcrReviewResult `json:"PoliticalOcrResult,omitempty" name:"PoliticalOcrResult"`

	// 视频内容审核 Asr 文字鉴黄任务的查询结果。
	PornAsrResult *VodPornAsrReviewResult `json:"PornAsrResult,omitempty" name:"PornAsrResult"`

	// 视频内容审核 Asr 文字鉴政任务的查询结果。
	PoliticalAsrResult *VodPoliticalAsrReviewResult `json:"PoliticalAsrResult,omitempty" name:"PoliticalAsrResult"`

	// 视频内容审核 Ocr 文字鉴黄任务的查询结果。
	PornOcrResult *VodPornOcrResult `json:"PornOcrResult,omitempty" name:"PornOcrResult"`

	// 原始视频的元信息。
	MetaData *VodMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVideoTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoTaskResponseParams `json:"Response"`
}

func (r *DescribeVideoTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisgustResult struct {
	// 该识别场景的错误码：
	// 0表示成功，
	// -1表示系统错误，
	// -2表示引擎错误。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误码描述信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 图像恶心的分数，0-100之间，分数越高恶心几率越大。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`
}

type FaceRect struct {
	// 人脸区域左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸区域左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸区域宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸区域高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type FaceResult struct {
	// 检测出的人脸框位置。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// 候选人列表。当前返回相似度最高的候选人。
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates"`
}

// Predefined struct for user
type ImageModerationRequestParams struct {
	// 本次调用支持的识别场景，可选值如下：
	// 1. PORN，即色情识别
	// 2. TERRORISM，即暴恐识别
	// 3. POLITICS，即政治敏感识别
	// 
	// 支持多场景（Scenes）一起检测。例如，使用 Scenes=["PORN", "TERRORISM"]，即对一张图片同时进行色情识别和暴恐识别。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// 图片URL地址。 
	// 图片限制： 
	//  • 图片格式：PNG、JPG、JPEG。 
	//  • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	//  • 图片像素：大于50*50像素，否则影响识别效果； 
	//  • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 预留字段，后期用于展示更多识别信息。
	Config *string `json:"Config,omitempty" name:"Config"`

	// 透传字段，透传简单信息。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type ImageModerationRequest struct {
	*tchttp.BaseRequest
	
	// 本次调用支持的识别场景，可选值如下：
	// 1. PORN，即色情识别
	// 2. TERRORISM，即暴恐识别
	// 3. POLITICS，即政治敏感识别
	// 
	// 支持多场景（Scenes）一起检测。例如，使用 Scenes=["PORN", "TERRORISM"]，即对一张图片同时进行色情识别和暴恐识别。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// 图片URL地址。 
	// 图片限制： 
	//  • 图片格式：PNG、JPG、JPEG。 
	//  • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	//  • 图片像素：大于50*50像素，否则影响识别效果； 
	//  • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 预留字段，后期用于展示更多识别信息。
	Config *string `json:"Config,omitempty" name:"Config"`

	// 透传字段，透传简单信息。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *ImageModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageModerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Scenes")
	delete(f, "ImageUrl")
	delete(f, "Config")
	delete(f, "Extra")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageModerationResponseParams struct {
	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 色情识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornResult *PornResult `json:"PornResult,omitempty" name:"PornResult"`

	// 暴恐识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismResult *TerrorismResult `json:"TerrorismResult,omitempty" name:"TerrorismResult"`

	// 政治敏感识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticsResult *PoliticsResult `json:"PoliticsResult,omitempty" name:"PoliticsResult"`

	// 透传字段，透传简单信息。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 恶心内容识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisgustResult *DisgustResult `json:"DisgustResult,omitempty" name:"DisgustResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImageModerationResponse struct {
	*tchttp.BaseResponse
	Response *ImageModerationResponseParams `json:"Response"`
}

func (r *ImageModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PoliticsResult struct {
	// 该识别场景的错误码：
	// 0表示成功，
	// -1表示系统错误，
	// -2表示引擎错误，
	// -1400表示图片解码失败，
	// -1401表示图片不符合规范。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误码描述信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 图像涉政的分数，0-100之间，分数越高涉政几率越大。
	// Type为DNA时：
	// 0到75，Suggestion建议为PASS
	// 75到90，Suggestion建议为REVIEW
	// 90到100，Suggestion建议为BLOCK
	// Type为FACE时：
	// 0到55，Suggestion建议为PASS
	// 55到60，Suggestion建议为REVIEW
	// 60到100，Suggestion建议为BLOCK
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// Type取值为‘FACE’时，人脸识别的结果列表。基于图片中实际检测到的人脸数，返回数组最大值不超过5个。
	FaceResults []*FaceResult `json:"FaceResults,omitempty" name:"FaceResults"`

	// 取值'DNA' 或‘FACE’。DNA表示结论和置信度来自图像指纹，FACE表示结论和置信度来自人脸识别。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 鉴政识别返回的详细标签后期开放。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type PornResult struct {
	// 该识别场景的错误码：
	// 0表示成功，
	// -1表示系统错误，
	// -2表示引擎错误，
	// -1400表示图片解码失败。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误码描述信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 算法对于Suggestion的置信度，0-100之间，值越高，表示对于Suggestion越确定。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 预留字段，后期用于展示更多识别信息。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// 取值'LABEL‘，LABEL表示结论和置信度来自标签分类。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type TerrorismResult struct {
	// 该识别场景的错误码：
	// 0表示成功，
	// -1表示系统错误，
	// -2表示引擎错误，
	// -1400表示图片解码失败。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误码描述信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 图像涉恐的分数，0-100之间，分数越高涉恐几率越大。
	// Type为LABEL时：
	// 0到86，Suggestion建议为PASS
	// 86到91，Suggestion建议为REVIEW
	// 91到100，Suggestion建议为BLOCK
	// Type为FACE时：
	// 0到55，Suggestion建议为PASS
	// 55到60，Suggestion建议为REVIEW
	// 60到100，Suggestion建议为BLOCK
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// Type取值为‘FACE’时，人脸识别的结果列表。基于图片中实际检测到的人脸数，返回数组最大值不超过5个。
	FaceResults []*FaceResult `json:"FaceResults,omitempty" name:"FaceResults"`

	// 暴恐识别返回的详细标签后期开放。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// 取值'LABEL' 或‘FACE’，LABEL表示结论和置信度来自标签分类，FACE表示结论和置信度来自人脸识别。
	Type *string `json:"Type,omitempty" name:"Type"`
}

// Predefined struct for user
type VideoModerationRequestParams struct {
	// 需要审核的视频的URL地址
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// 开发者标识
	DeveloperId *string `json:"DeveloperId,omitempty" name:"DeveloperId"`

	// 审核完成后回调地址
	CBUrl *string `json:"CBUrl,omitempty" name:"CBUrl"`

	// 透传字段，透传简单信息。
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

type VideoModerationRequest struct {
	*tchttp.BaseRequest
	
	// 需要审核的视频的URL地址
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// 开发者标识
	DeveloperId *string `json:"DeveloperId,omitempty" name:"DeveloperId"`

	// 审核完成后回调地址
	CBUrl *string `json:"CBUrl,omitempty" name:"CBUrl"`

	// 透传字段，透传简单信息。
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

func (r *VideoModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VideoModerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoUrl")
	delete(f, "DeveloperId")
	delete(f, "CBUrl")
	delete(f, "Extra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VideoModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VideoModerationResponseParams struct {
	// 视频审核任务ID
	VodTaskId *string `json:"VodTaskId,omitempty" name:"VodTaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VideoModerationResponse struct {
	*tchttp.BaseResponse
	Response *VideoModerationResponseParams `json:"Response"`
}

func (r *VideoModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VideoModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VodAsrTextSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段审核结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 嫌疑关键词列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet"`
}

type VodAudioStreamItem struct {
	// 音频流的码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频流的采样率，单位：hz。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SamplingRate *int64 `json:"SamplingRate,omitempty" name:"SamplingRate"`

	// 音频流的编码格式，例如 aac。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type VodMetaData struct {
	// 上传的媒体文件大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 容器类型，例如 m4a，mp4 等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 视频流码率平均值与音频流码率平均值之和，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 视频拍摄时的选择角度，单位：度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 视频流信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoStreamSet []*VodVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet"`

	// 音频流信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioStreamSet []*VodAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet"`

	// 视频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoDuration *float64 `json:"VideoDuration,omitempty" name:"VideoDuration"`

	// 音频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioDuration *float64 `json:"AudioDuration,omitempty" name:"AudioDuration"`
}

type VodOcrTextSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段审核结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 嫌疑关键词列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet"`

	// 嫌疑文字出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`
}

type VodPoliticalAsrReviewResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 嫌疑片段审核结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// Asr 文字涉政、敏感评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Asr 文字涉政、敏感结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Asr 文字有涉政、敏感嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*VodAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type VodPoliticalOcrReviewResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Ocr 文字涉政、敏感评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Ocr 文字涉政、敏感结果建议，取值范围：
	// pass。
	// review。
	// block。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Ocr 文字有涉政、敏感嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*VodOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type VodPoliticalReviewResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 视频涉政评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 涉政结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 视频鉴政结果标签，取值范围：
	// politician：政治人物。
	// violation_photo：违规图标。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 有涉政嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*VodPoliticalReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type VodPoliticalReviewSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段涉政分数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段鉴政结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 涉政人物、违规图标名字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 嫌疑片段鉴政结果标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	// PicUrlExpireTime 时间点后图片将被删除）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 ISO 日期格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`

	// 涉政人物、违规图标出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`
}

type VodPornAsrReviewResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Asr 文字涉黄评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Asr 文字涉黄结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Asr 文字有涉黄嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*VodAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type VodPornOcrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Ocr 文字涉黄评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Ocr 文字涉黄结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Ocr 文字有涉黄嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*VodOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type VodPornReviewResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 视频鉴黄评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 鉴黄结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 视频鉴黄结果标签，取值范围：
	// porn：色情。
	// sexy：性感。
	// vulgar：低俗。
	// intimacy：亲密行为。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 有涉黄嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*VodPornReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type VodPornReviewSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段涉黄分数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段鉴黄结果标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 嫌疑片段鉴黄结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	// PicUrlExpireTime 时间点后图片将被删除）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 ISO 日期格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`
}

type VodTerrorismReviewResult struct {
	// 视频暴恐评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 暴恐结果建议，取值范围：
	// pass。
	// review。
	// block。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 视频暴恐结果标签，取值范围：
	// guns：武器枪支。
	// crowd：人群聚集。
	// police：警察部队。
	// bloody：血腥画面。
	// banners：暴恐旗帜。
	// militant：武装分子。
	// explosion：爆炸火灾。
	// terrorists：暴恐人物。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 有暴恐嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*VodPornReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type VodVideoStreamItem struct {
	// 视频流的码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流的高度，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流的宽度，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频流的编码格式，例如 h264。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 帧率，单位：hz。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`
}