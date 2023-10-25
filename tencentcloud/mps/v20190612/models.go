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

package v20190612

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AIAnalysisTemplateItem struct {
	// 智能分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 智能分析模板名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 智能分析模板描述信息。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 智能分类任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitnil" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitnil" name:"TagConfigure"`

	// 智能封面任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitnil" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitnil" name:"FrameTagConfigure"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 模板类型，取值范围：
	// * Preset：系统预置模板；
	// * Custom：用户自定义模板。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type AIRecognitionTemplateItem struct {
	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 视频内容识别模板名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 视频内容识别模板描述信息。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitnil" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitnil" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitnil" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitnil" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitnil" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranslateConfigure *TranslateConfigureInfo `json:"TranslateConfigure,omitnil" name:"TranslateConfigure"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 模板类型，取值范围：
	// * Preset：系统预置模板；
	// * Custom：用户自定义模板。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type ActionConfigInfo struct {
	// 动作识别任务开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type Activity struct {
	// 原子任务类型：
	// <li>input: 起始节点</li>
	// <li>output：终止节点</li>
	// <li>action-trans：转码</li>
	// <li>action-samplesnapshot：采样截图</li>
	// <li>action-AIAnalysis: 分析</li>
	// <li>action-AIRecognition：识别</li>
	// <li>action-aiReview：审核</li>
	// <li>action-animated-graphics：转动图</li>
	// <li>action-image-sprite：雪碧图</li>
	// <li>action-snapshotByTimeOffset: 时间点截图</li>
	// <li>action-adaptive-substream：自适应码流</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityType *string `json:"ActivityType,omitnil" name:"ActivityType"`

	// 后驱节点索引数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReardriveIndex []*int64 `json:"ReardriveIndex,omitnil" name:"ReardriveIndex"`

	// 原子任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityPara *ActivityPara `json:"ActivityPara,omitnil" name:"ActivityPara"`
}

type ActivityPara struct {
	// 视频转码任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTask *TranscodeTaskInput `json:"TranscodeTask,omitnil" name:"TranscodeTask"`

	// 视频转动图任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicTask *AnimatedGraphicTaskInput `json:"AnimatedGraphicTask,omitnil" name:"AnimatedGraphicTask"`

	// 视频按时间点截图任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTask *SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTask,omitnil" name:"SnapshotByTimeOffsetTask"`

	// 视频采样截图任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotTask *SampleSnapshotTaskInput `json:"SampleSnapshotTask,omitnil" name:"SampleSnapshotTask"`

	// 视频截雪碧图任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteTask *ImageSpriteTaskInput `json:"ImageSpriteTask,omitnil" name:"ImageSpriteTask"`

	// 转自适应码流任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingTask *AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTask,omitnil" name:"AdaptiveDynamicStreamingTask"`

	// 视频内容审核类型任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容分析类型任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容识别类型任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`
}

type ActivityResItem struct {
	// 转码任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitnil" name:"TranscodeTask"`

	// 转动图任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitnil" name:"AnimatedGraphicTask"`

	// 时间点截图任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitnil" name:"SnapshotByTimeOffsetTask"`

	// 采样截图任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitnil" name:"SampleSnapshotTask"`

	// 雪碧图任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitnil" name:"ImageSpriteTask"`

	// 自适应码流任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingTask *MediaProcessTaskAdaptiveDynamicStreamingResult `json:"AdaptiveDynamicStreamingTask,omitnil" name:"AdaptiveDynamicStreamingTask"`

	// 识别任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecognitionTask *ScheduleRecognitionTaskResult `json:"RecognitionTask,omitnil" name:"RecognitionTask"`

	// 审核任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewTask *ScheduleReviewTaskResult `json:"ReviewTask,omitnil" name:"ReviewTask"`

	// 分析任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisTask *ScheduleAnalysisTaskResult `json:"AnalysisTask,omitnil" name:"AnalysisTask"`
}

type ActivityResult struct {
	// 原子任务类型。
	// <li>Transcode：转码。</li>
	// <li>SampleSnapshot：采样截图。</li>
	// <li>AnimatedGraphics：转动图。</li>
	// <li>SnapshotByTimeOffset：时间点截图。</li>
	// <li>ImageSprites：雪碧图。</li>
	// <li>AdaptiveDynamicStreaming：自适应码流。</li>
	// <li>AiContentReview：内容审核。</li>
	// <li>AIRecognition：智能识别。</li>
	// <li>AIAnalysis：智能分析。</li>
	ActivityType *string `json:"ActivityType,omitnil" name:"ActivityType"`

	// 原子任务输出。
	ActivityResItem *ActivityResItem `json:"ActivityResItem,omitnil" name:"ActivityResItem"`
}

type AdaptiveDynamicStreamingInfoItem struct {
	// 转自适应码流规格。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 打包格式，可能为 HLS和 MPEG-DASH 两种。
	Package *string `json:"Package,omitnil" name:"Package"`

	// 播放路径。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 自适应码流文件的存储位置。
	Storage *TaskOutputStorage `json:"Storage,omitnil" name:"Storage"`
}

type AdaptiveDynamicStreamingTaskInput struct {
	// 转自适应码流模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil" name:"WatermarkSet"`

	// 转自适应码流后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 转自适应码流后，manifest 文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil" name:"OutputObjectPath"`

	// 转自适应码流后，子流文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}.{format}`。
	SubStreamObjectName *string `json:"SubStreamObjectName,omitnil" name:"SubStreamObjectName"`

	// 转自适应码流（仅 HLS）后，分片文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}_{segmentNumber}.{format}`。
	SegmentObjectName *string `json:"SegmentObjectName,omitnil" name:"SegmentObjectName"`

	// 要插入的字幕文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddOnSubtitles []*AddOnSubtitle `json:"AddOnSubtitles,omitnil" name:"AddOnSubtitles"`

	// Drm信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrmInfo *DrmInfo `json:"DrmInfo,omitnil" name:"DrmInfo"`
}

type AdaptiveDynamicStreamingTemplate struct {
	// 转自适应码流模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 转自适应码流模板名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 转自适应码流模板描述信息。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 转自适应码流格式，取值范围：
	// <li>HLS，</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil" name:"Format"`

	// 转自适应码流输入流参数信息，最多输入10路流。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil" name:"StreamInfos"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil" name:"DisableHigherVideoResolution"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type AdaptiveStreamTemplate struct {
	// 视频参数信息。
	Video *VideoTemplateInfo `json:"Video,omitnil" name:"Video"`

	// 音频参数信息。
	Audio *AudioTemplateInfo `json:"Audio,omitnil" name:"Audio"`

	// 是否移除音频流，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	RemoveAudio *uint64 `json:"RemoveAudio,omitnil" name:"RemoveAudio"`

	// 是否移除视频流，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	RemoveVideo *uint64 `json:"RemoveVideo,omitnil" name:"RemoveVideo"`
}

type AddOnSubtitle struct {
	// 插入形式，可选值：
	// <li>subtitle-stream：插入字幕轨道</li>
	// <li>close-caption-708：CEA-708字幕编码到SEI帧</li>
	// <li>close-caption-608：CEA-608字幕编码到SEI帧</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 字幕文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subtitle *MediaInputInfo `json:"Subtitle,omitnil" name:"Subtitle"`
}

type AiAnalysisResult struct {
	// 任务的类型，可以取的值有：
	// <li>Classification：智能分类</li>
	// <li>Cover：智能封面</li>
	// <li>Tag：智能标签</li>
	// <li>FrameTag：智能按帧标签</li>
	// <li>Highlight：智能精彩集锦</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 视频内容分析智能分类任务的查询结果，当任务类型为 Classification 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationTask *AiAnalysisTaskClassificationResult `json:"ClassificationTask,omitnil" name:"ClassificationTask"`

	// 视频内容分析智能封面任务的查询结果，当任务类型为 Cover 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverTask *AiAnalysisTaskCoverResult `json:"CoverTask,omitnil" name:"CoverTask"`

	// 视频内容分析智能标签任务的查询结果，当任务类型为 Tag 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagTask *AiAnalysisTaskTagResult `json:"TagTask,omitnil" name:"TagTask"`

	// 视频内容分析智能按帧标签任务的查询结果，当任务类型为 FrameTag 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagTask *AiAnalysisTaskFrameTagResult `json:"FrameTagTask,omitnil" name:"FrameTagTask"`

	// 视频内容分析集锦任务的查询结果，当任务类型为 Highlight时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighlightTask *AiAnalysisTaskHighlightResult `json:"HighlightTask,omitnil" name:"HighlightTask"`

	// 视频内容分析去水印任务的查询结果，当任务类型为 DeLogo 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeLogoTask *AiAnalysisTaskDelLogoResult `json:"DeLogoTask,omitnil" name:"DeLogoTask"`
}

type AiAnalysisTaskClassificationInput struct {
	// 视频智能分类模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiAnalysisTaskClassificationOutput struct {
	// 视频智能分类列表。
	ClassificationSet []*MediaAiAnalysisClassificationItem `json:"ClassificationSet,omitnil" name:"ClassificationSet"`
}

type AiAnalysisTaskClassificationResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 智能分类任务输入。
	Input *AiAnalysisTaskClassificationInput `json:"Input,omitnil" name:"Input"`

	// 智能分类任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskClassificationOutput `json:"Output,omitnil" name:"Output"`
}

type AiAnalysisTaskCoverInput struct {
	// 视频智能封面模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiAnalysisTaskCoverOutput struct {
	// 智能封面列表。
	CoverSet []*MediaAiAnalysisCoverItem `json:"CoverSet,omitnil" name:"CoverSet"`

	// 智能封面的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`
}

type AiAnalysisTaskCoverResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 智能封面任务输入。
	Input *AiAnalysisTaskCoverInput `json:"Input,omitnil" name:"Input"`

	// 智能封面任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskCoverOutput `json:"Output,omitnil" name:"Output"`
}

type AiAnalysisTaskDelLogoInput struct {
	// 视频智能去水印模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiAnalysisTaskDelLogoOutput struct {
	// 去水印后文件的路径。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 去水印后文件的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`
}

type AiAnalysisTaskDelLogoResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 智能去水印任务输入。
	Input *AiAnalysisTaskDelLogoInput `json:"Input,omitnil" name:"Input"`

	// 智能去水印任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskDelLogoOutput `json:"Output,omitnil" name:"Output"`
}

type AiAnalysisTaskFrameTagInput struct {
	// 视频智能按帧标签模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiAnalysisTaskFrameTagOutput struct {
	// 视频按帧标签列表。
	SegmentSet []*MediaAiAnalysisFrameTagSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiAnalysisTaskFrameTagResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 智能按帧标签任务输入。
	Input *AiAnalysisTaskFrameTagInput `json:"Input,omitnil" name:"Input"`

	// 智能按帧标签任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskFrameTagOutput `json:"Output,omitnil" name:"Output"`
}

type AiAnalysisTaskHighlightInput struct {
	// 视频智能精彩片段模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiAnalysisTaskHighlightOutput struct {
	// 视频智能精彩片段列表。
	HighlightSet []*MediaAiAnalysisHighlightItem `json:"HighlightSet,omitnil" name:"HighlightSet"`

	// 精彩片段的存储位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`
}

type AiAnalysisTaskHighlightResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 智能精彩片段任务输入。
	Input *AiAnalysisTaskHighlightInput `json:"Input,omitnil" name:"Input"`

	// 智能精彩片段任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskHighlightOutput `json:"Output,omitnil" name:"Output"`
}

type AiAnalysisTaskInput struct {
	// 视频内容分析模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 扩展参数，其值为序列化的 json字符串。
	// 注意：此参数为定制需求参数，需要线下对接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendedParameter *string `json:"ExtendedParameter,omitnil" name:"ExtendedParameter"`
}

type AiAnalysisTaskTagInput struct {
	// 视频智能标签模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiAnalysisTaskTagOutput struct {
	// 视频智能标签列表。
	TagSet []*MediaAiAnalysisTagItem `json:"TagSet,omitnil" name:"TagSet"`
}

type AiAnalysisTaskTagResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 智能标签任务输入。
	Input *AiAnalysisTaskTagInput `json:"Input,omitnil" name:"Input"`

	// 智能标签任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskTagOutput `json:"Output,omitnil" name:"Output"`
}

type AiContentReviewResult struct {
	// 任务的类型，可以取的值有：
	// <li>Porn：图片鉴黄</li>
	// <li>Terrorism：图片敏感</li>
	// <li>Political：图片敏感</li>
	// <li>Porn.Asr：Asr 文字鉴黄</li>
	// <li>Porn.Ocr：Ocr 文字鉴黄</li>
	// <li>Political.Asr：Asr 文字敏感</li>
	// <li>Political.Ocr：Ocr 文字敏感</li>
	// <li>Terrorism.Ocr：Ocr 文字敏感</li>
	// <li>Prohibited.Asr：Asr 文字鉴违禁</li>
	// <li>Prohibited.Ocr：Ocr 文字鉴违禁</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 采样频率，即对视频每秒截取进行审核的帧数。
	SampleRate *float64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// 审核的视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitnil" name:"Duration"`

	// 视频内容审核智能画面鉴黄任务的查询结果，当任务类型为 Porn 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornTask *AiReviewTaskPornResult `json:"PornTask,omitnil" name:"PornTask"`

	// 视频内容审核智能画面敏感任务的查询结果，当任务类型为 Terrorism 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismTask *AiReviewTaskTerrorismResult `json:"TerrorismTask,omitnil" name:"TerrorismTask"`

	// 视频内容审核智能画面敏感任务的查询结果，当任务类型为 Political 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalTask *AiReviewTaskPoliticalResult `json:"PoliticalTask,omitnil" name:"PoliticalTask"`

	// 视频内容审核 Asr 文字鉴黄任务的查询结果，当任务类型为 Porn.Asr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornAsrTask *AiReviewTaskPornAsrResult `json:"PornAsrTask,omitnil" name:"PornAsrTask"`

	// 视频内容审核 Ocr 文字鉴黄任务的查询结果，当任务类型为 Porn.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornOcrTask *AiReviewTaskPornOcrResult `json:"PornOcrTask,omitnil" name:"PornOcrTask"`

	// 视频内容审核 Asr 文字敏感任务的查询结果，当任务类型为 Political.Asr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalAsrTask *AiReviewTaskPoliticalAsrResult `json:"PoliticalAsrTask,omitnil" name:"PoliticalAsrTask"`

	// 视频内容审核 Ocr 文字敏感任务的查询结果，当任务类型为 Political.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalOcrTask *AiReviewTaskPoliticalOcrResult `json:"PoliticalOcrTask,omitnil" name:"PoliticalOcrTask"`

	// 视频内容审核 Ocr 文字敏感任务的查询结果，当任务类型为 Terrorism.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismOcrTask *AiReviewTaskTerrorismOcrResult `json:"TerrorismOcrTask,omitnil" name:"TerrorismOcrTask"`

	// 视频内容审核 Asr 文字鉴违禁任务的查询结果，当任务类型为 Prohibited.Asr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProhibitedAsrTask *AiReviewTaskProhibitedAsrResult `json:"ProhibitedAsrTask,omitnil" name:"ProhibitedAsrTask"`

	// 视频内容审核 Ocr 文字鉴违禁任务的查询结果，当任务类型为 Prohibited.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProhibitedOcrTask *AiReviewTaskProhibitedOcrResult `json:"ProhibitedOcrTask,omitnil" name:"ProhibitedOcrTask"`
}

type AiContentReviewTaskInput struct {
	// 视频内容审核模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiQualityControlTaskInput struct {
	// 视频质检模板 ID 。暂时可以直接使用 预设模板ID 10，后面控制台支持用户配置自定义模板。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 渠道扩展参数json序列化字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelExtPara *string `json:"ChannelExtPara,omitnil" name:"ChannelExtPara"`
}

type AiRecognitionResult struct {
	// 任务的类型，取值范围：
	// <li>FaceRecognition：人脸识别，</li>
	// <li>AsrWordsRecognition：语音关键词识别，</li>
	// <li>OcrWordsRecognition：文本关键词识别，</li>
	// <li>AsrFullTextRecognition：语音全文识别，</li>
	// <li>OcrFullTextRecognition：文本全文识别。</li>
	// <li>TransTextRecognition：语音翻译。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 人脸识别结果，当 Type 为 
	//  FaceRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceTask *AiRecognitionTaskFaceResult `json:"FaceTask,omitnil" name:"FaceTask"`

	// 语音关键词识别结果，当 Type 为
	//  AsrWordsRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrWordsTask *AiRecognitionTaskAsrWordsResult `json:"AsrWordsTask,omitnil" name:"AsrWordsTask"`

	// 语音全文识别结果，当 Type 为
	//  AsrFullTextRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrFullTextTask *AiRecognitionTaskAsrFullTextResult `json:"AsrFullTextTask,omitnil" name:"AsrFullTextTask"`

	// 文本关键词识别结果，当 Type 为
	//  OcrWordsRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrWordsTask *AiRecognitionTaskOcrWordsResult `json:"OcrWordsTask,omitnil" name:"OcrWordsTask"`

	// 文本全文识别结果，当 Type 为
	//  OcrFullTextRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrFullTextTask *AiRecognitionTaskOcrFullTextResult `json:"OcrFullTextTask,omitnil" name:"OcrFullTextTask"`

	// 翻译结果，当 Type 为
	// 
	// TransTextRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransTextTask *AiRecognitionTaskTransTextResult `json:"TransTextTask,omitnil" name:"TransTextTask"`

	// 物体识别结果，当Type 为
	// 
	// ObjectRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectTask *AiRecognitionTaskObjectResult `json:"ObjectTask,omitnil" name:"ObjectTask"`
}

type AiRecognitionTaskAsrFullTextResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 语音全文识别任务输入信息。
	Input *AiRecognitionTaskAsrFullTextResultInput `json:"Input,omitnil" name:"Input"`

	// 语音全文识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskAsrFullTextResultOutput `json:"Output,omitnil" name:"Output"`
}

type AiRecognitionTaskAsrFullTextResultInput struct {
	// 语音全文识别模板 ID。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type AiRecognitionTaskAsrFullTextResultOutput struct {
	// 语音全文识别片段列表。
	SegmentSet []*AiRecognitionTaskAsrFullTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`

	// 字幕文件地址。
	SubtitlePath *string `json:"SubtitlePath,omitnil" name:"SubtitlePath"`

	// 字幕文件存储位置。
	//
	// Deprecated: OutputStorage is deprecated.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`
}

type AiRecognitionTaskAsrFullTextSegmentItem struct {
	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 识别文本。
	Text *string `json:"Text,omitnil" name:"Text"`
}

type AiRecognitionTaskAsrWordsResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 语音关键词识别任务输入信息。
	Input *AiRecognitionTaskAsrWordsResultInput `json:"Input,omitnil" name:"Input"`

	// 语音关键词识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskAsrWordsResultOutput `json:"Output,omitnil" name:"Output"`
}

type AiRecognitionTaskAsrWordsResultInput struct {
	// 语音关键词识别模板 ID。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type AiRecognitionTaskAsrWordsResultItem struct {
	// 语音关键词。
	Word *string `json:"Word,omitnil" name:"Word"`

	// 语音关键词出现的时间片段列表。
	SegmentSet []*AiRecognitionTaskAsrWordsSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiRecognitionTaskAsrWordsResultOutput struct {
	// 语音关键词识别结果集。
	ResultSet []*AiRecognitionTaskAsrWordsResultItem `json:"ResultSet,omitnil" name:"ResultSet"`
}

type AiRecognitionTaskAsrWordsSegmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`
}

type AiRecognitionTaskFaceResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 人脸识别任务输入信息。
	Input *AiRecognitionTaskFaceResultInput `json:"Input,omitnil" name:"Input"`

	// 人脸识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskFaceResultOutput `json:"Output,omitnil" name:"Output"`
}

type AiRecognitionTaskFaceResultInput struct {
	// 人脸识别模板 ID。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type AiRecognitionTaskFaceResultItem struct {
	// 人物唯一标识 ID。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 人物库类型，表示识别出的人物来自哪个人物库：
	// <li>Default：默认人物库；</li>
	// <li>UserDefine：用户自定义人物库。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 人物名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 人物出现的片段结果集。
	SegmentSet []*AiRecognitionTaskFaceSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`

	// 人物性别：
	// <li>Male：男性；</li>
	// <li>Female：女性。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gender *string `json:"Gender,omitnil" name:"Gender"`

	// 人物出生日期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// 人物职业或者职务。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Profession *string `json:"Profession,omitnil" name:"Profession"`

	// 人物毕业院校。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchoolOfGraduation *string `json:"SchoolOfGraduation,omitnil" name:"SchoolOfGraduation"`

	// 人物简介。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abstract *string `json:"Abstract,omitnil" name:"Abstract"`

	// 人物出生地或者籍贯。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlaceOfBirth *string `json:"PlaceOfBirth,omitnil" name:"PlaceOfBirth"`

	// 人物类型：
	// <li>Politician：官员；</li>
	// <li>Artist：艺人。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonType *string `json:"PersonType,omitnil" name:"PersonType"`

	// 敏感度标注：
	// <li>Normal：正常；</li>
	// <li>Sensitive：敏感。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 截图链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type AiRecognitionTaskFaceResultOutput struct {
	// 智能人脸识别结果集。
	ResultSet []*AiRecognitionTaskFaceResultItem `json:"ResultSet,omitnil" name:"ResultSet"`
}

type AiRecognitionTaskFaceSegmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`
}

type AiRecognitionTaskInput struct {
	// 视频智能识别模板 ID 。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiRecognitionTaskObjectResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 物体识别任务输入信息。
	Input *AiRecognitionTaskObjectResultInput `json:"Input,omitnil" name:"Input"`

	// 物体识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskObjectResultOutput `json:"Output,omitnil" name:"Output"`
}

type AiRecognitionTaskObjectResultInput struct {
	// 物体识别模板 ID。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type AiRecognitionTaskObjectResultItem struct {
	// 识别的物体名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 物体出现的片段列表。
	SegmentSet []*AiRecognitionTaskObjectSeqmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiRecognitionTaskObjectResultOutput struct {
	// 智能物体识别结果集。
	ResultSet []*AiRecognitionTaskObjectResultItem `json:"ResultSet,omitnil" name:"ResultSet"`
}

type AiRecognitionTaskObjectSeqmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`
}

type AiRecognitionTaskOcrFullTextResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 文本全文识别任务输入信息。
	Input *AiRecognitionTaskOcrFullTextResultInput `json:"Input,omitnil" name:"Input"`

	// 文本全文识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskOcrFullTextResultOutput `json:"Output,omitnil" name:"Output"`
}

type AiRecognitionTaskOcrFullTextResultInput struct {
	// 文本全文识别模板 ID。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type AiRecognitionTaskOcrFullTextResultOutput struct {
	// 文本全文识别结果集。
	SegmentSet []*AiRecognitionTaskOcrFullTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiRecognitionTaskOcrFullTextSegmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 识别片段结果集。
	TextSet []*AiRecognitionTaskOcrFullTextSegmentTextItem `json:"TextSet,omitnil" name:"TextSet"`
}

type AiRecognitionTaskOcrFullTextSegmentTextItem struct {
	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`

	// 识别文本。
	Text *string `json:"Text,omitnil" name:"Text"`
}

type AiRecognitionTaskOcrWordsResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 文本关键词识别任务输入信息。
	Input *AiRecognitionTaskOcrWordsResultInput `json:"Input,omitnil" name:"Input"`

	// 文本关键词识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskOcrWordsResultOutput `json:"Output,omitnil" name:"Output"`
}

type AiRecognitionTaskOcrWordsResultInput struct {
	// 文本关键词识别模板 ID。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type AiRecognitionTaskOcrWordsResultItem struct {
	// 文本关键词。
	Word *string `json:"Word,omitnil" name:"Word"`

	// 文本关键出现的片段列表。
	SegmentSet []*AiRecognitionTaskOcrWordsSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiRecognitionTaskOcrWordsResultOutput struct {
	// 文本关键词识别结果集。
	ResultSet []*AiRecognitionTaskOcrWordsResultItem `json:"ResultSet,omitnil" name:"ResultSet"`
}

type AiRecognitionTaskOcrWordsSegmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`
}

type AiRecognitionTaskTransTextResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 翻译任务输入信息。
	Input *AiRecognitionTaskTransTextResultInput `json:"Input,omitnil" name:"Input"`

	// 翻译任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskTransTextResultOutput `json:"Output,omitnil" name:"Output"`
}

type AiRecognitionTaskTransTextResultInput struct {
	// 翻译模板 ID。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type AiRecognitionTaskTransTextResultOutput struct {
	// 翻译片段列表。
	SegmentSet []*AiRecognitionTaskTransTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`

	// 字幕文件地址。
	SubtitlePath *string `json:"SubtitlePath,omitnil" name:"SubtitlePath"`
}

type AiRecognitionTaskTransTextSegmentItem struct {
	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 识别文本。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 翻译文本。
	Trans *string `json:"Trans,omitnil" name:"Trans"`
}

type AiReviewPoliticalAsrTaskInput struct {
	// 模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewPoliticalAsrTaskOutput struct {
	// Asr 文字敏感评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// Asr 文字敏感结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// Asr 文字敏感嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiReviewPoliticalOcrTaskInput struct {
	// 模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewPoliticalOcrTaskOutput struct {
	// Ocr 文字敏感评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// Ocr 文字敏感结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// Ocr 文字有敏感嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiReviewPoliticalTaskInput struct {
	// 模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewPoliticalTaskOutput struct {
	// 视频涉敏评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 涉敏结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 视频涉敏结果标签。内容审核模板[画面涉敏任务控制参数](https://cloud.tencent.com/document/api/862/37615#AiReviewPoliticalTaskOutput)里 LabelSet 参数与此参数取值范围的对应关系：
	// violation_photo：
	// <li>violation_photo：违规图标。</li>
	// 其他（即 politician/entertainment/sport/entrepreneur/scholar/celebrity/military）：
	// <li>politician：涉敏人物。</li>
	Label *string `json:"Label,omitnil" name:"Label"`

	// 有涉敏嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewPoliticalSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiReviewPornAsrTaskInput struct {
	// 鉴黄模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewPornAsrTaskOutput struct {
	// Asr 文字涉黄评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// Asr 文字涉黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// Asr 文字有涉黄嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiReviewPornOcrTaskInput struct {
	// 鉴黄模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewPornOcrTaskOutput struct {
	// Ocr 文字涉黄评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// Ocr 文字涉黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// Ocr 文字有涉黄嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiReviewPornTaskInput struct {
	// 鉴黄模板 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewPornTaskOutput struct {
	// 视频鉴黄评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 鉴黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 视频鉴黄结果标签，取值范围：
	// <li>porn：色情。</li>
	// <li>sexy：性感。</li>
	// <li>vulgar：低俗。</li>
	// <li>intimacy：亲密行为。</li>
	Label *string `json:"Label,omitnil" name:"Label"`

	// 有涉黄嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiReviewProhibitedAsrTaskInput struct {
	// 鉴违禁模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewProhibitedAsrTaskOutput struct {
	// Asr 文字涉违禁评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// Asr 文字涉违禁结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// Asr 文字有涉违禁嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiReviewProhibitedOcrTaskInput struct {
	// 鉴违禁模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewProhibitedOcrTaskOutput struct {
	// Ocr 文字涉违禁评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// Ocr 文字涉违禁结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// Ocr 文字有涉违禁嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiReviewTaskPoliticalAsrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核 Asr 文字敏感任务输入。
	Input *AiReviewPoliticalAsrTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核 Asr 文字敏感任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPoliticalAsrTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTaskPoliticalOcrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS，FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核 Ocr 文字敏感任务输入。
	Input *AiReviewPoliticalOcrTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核 Ocr 文字敏感任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPoliticalOcrTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTaskPoliticalResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核涉敏任务输入。
	Input *AiReviewPoliticalTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核涉敏任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPoliticalTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTaskPornAsrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核 Asr 文字鉴黄任务输入。
	Input *AiReviewPornAsrTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核 Asr 文字鉴黄任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPornAsrTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTaskPornOcrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核 Ocr 文字鉴黄任务输入。
	Input *AiReviewPornOcrTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核 Ocr 文字鉴黄任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPornOcrTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTaskPornResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核鉴黄任务输入。
	Input *AiReviewPornTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核鉴黄任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPornTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTaskProhibitedAsrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核 Asr 文字鉴违禁任务输入。
	Input *AiReviewProhibitedAsrTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核 Asr 文字鉴违禁任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewProhibitedAsrTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTaskProhibitedOcrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核 Ocr 文字鉴违禁任务输入。
	Input *AiReviewProhibitedOcrTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核 Ocr 文字鉴违禁任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewProhibitedOcrTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTaskTerrorismOcrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核 Ocr 文字敏感任务输入。
	Input *AiReviewTerrorismOcrTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核 Ocr 文字敏感任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewTerrorismOcrTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTaskTerrorismResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 内容审核涉敏任务输入。
	Input *AiReviewTerrorismTaskInput `json:"Input,omitnil" name:"Input"`

	// 内容审核涉敏任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewTerrorismTaskOutput `json:"Output,omitnil" name:"Output"`
}

type AiReviewTerrorismOcrTaskInput struct {
	// 模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewTerrorismOcrTaskOutput struct {
	// Ocr 文字涉敏评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// Ocr 文字涉敏结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// Ocr 文字有涉敏嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiReviewTerrorismTaskInput struct {
	// 模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type AiReviewTerrorismTaskOutput struct {
	// 视频涉敏评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 涉敏结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 视频涉敏结果标签，取值范围：
	// <li>guns：武器枪支。</li>
	// <li>crowd：人群聚集。</li>
	// <li>police：警察部队。</li>
	// <li>bloody：血腥画面。</li>
	// <li>banners：涉敏旗帜。</li>
	// <li>militant：武装分子。</li>
	// <li>explosion：爆炸火灾。</li>
	// <li>terrorists：涉敏人物。</li>
	// <li>scenario：涉敏画面。</li>
	Label *string `json:"Label,omitnil" name:"Label"`

	// 有涉敏嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type AiSampleFaceInfo struct {
	// 人脸图片 ID。
	FaceId *string `json:"FaceId,omitnil" name:"FaceId"`

	// 人脸图片地址。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type AiSampleFaceOperation struct {
	// 操作类型，可选值：add（添加）、delete（删除）、reset（重置）。重置操作将清空该人物已有人脸数据，并添加 FaceContents 指定人脸数据。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 人脸 ID 集合，当 Type为delete 时，该字段必填。
	FaceIds []*string `json:"FaceIds,omitnil" name:"FaceIds"`

	// 人脸图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串集合。
	// <li>当 Type为add 或 reset 时，该字段必填；</li>
	// <li>数组长度限制：5 张图片。</li>
	// 注意：图片必须是单人像正面人脸较清晰的照片，像素不低于 200*200。
	FaceContents []*string `json:"FaceContents,omitnil" name:"FaceContents"`
}

type AiSampleFailFaceInfo struct {
	// 对应入参 FaceContents 中错误图片下标，从 0 开始。
	Index *uint64 `json:"Index,omitnil" name:"Index"`

	// 错误码，取值：
	// <li>0：成功；</li>
	// <li>其他：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误描述。
	Message *string `json:"Message,omitnil" name:"Message"`
}

type AiSamplePerson struct {
	// 人物 ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 人物名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 人物描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 人脸信息。
	FaceInfoSet []*AiSampleFaceInfo `json:"FaceInfoSet,omitnil" name:"FaceInfoSet"`

	// 人物标签。
	TagSet []*string `json:"TagSet,omitnil" name:"TagSet"`

	// 应用场景。
	UsageSet []*string `json:"UsageSet,omitnil" name:"UsageSet"`

	// 创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type AiSampleTagOperation struct {
	// 操作类型，可选值：add（添加）、delete（删除）、reset（重置）。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 标签，长度限制：128 个字符。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type AiSampleWord struct {
	// 关键词。
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 关键词标签。
	TagSet []*string `json:"TagSet,omitnil" name:"TagSet"`

	// 关键词应用场景。
	UsageSet []*string `json:"UsageSet,omitnil" name:"UsageSet"`

	// 创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type AiSampleWordInfo struct {
	// 关键词，长度限制：20 个字符。
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 关键词标签
	// <li>数组长度限制：20 个标签；</li>
	// <li>单个标签长度限制：128 个字符。</li>
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type AnimatedGraphicTaskInput struct {
	// 视频转动图模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 动图在视频中的开始时间，单位为秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 动图在视频中的结束时间，单位为秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 转动图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 转动图后文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_animatedGraphic_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil" name:"OutputObjectPath"`
}

type AnimatedGraphicsTemplate struct {
	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 转动图模板名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 转动图模板描述。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 动图格式。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 帧率。
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// 图片质量。
	Quality *float64 `json:"Quality,omitnil" name:"Quality"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type ArtifactRepairConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 类型，可选值：
	// <li>weak</li>
	// <li>strong</li>
	// 默认值：weak。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type AsrFullTextConfigureInfo struct {
	// 语音全文识别任务开关，可选值：
	// <li>ON：开启智能语音全文识别任务；</li>
	// <li>OFF：关闭智能语音全文识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 生成的字幕文件格式，不填或者填空字符串表示不生成字幕文件，可选值：
	// <li>vtt：生成 WebVTT 字幕文件。</li>
	SubtitleFormat *string `json:"SubtitleFormat,omitnil" name:"SubtitleFormat"`

	// 视频源语言。
	SourceLanguage *string `json:"SourceLanguage,omitnil" name:"SourceLanguage"`
}

type AsrFullTextConfigureInfoForUpdate struct {
	// 语音全文识别任务开关，可选值：
	// <li>ON：开启智能语音全文识别任务；</li>
	// <li>OFF：关闭智能语音全文识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 生成的字幕文件格式，填空字符串表示不生成字幕文件，可选值：
	// <li>vtt：生成 WebVTT 字幕文件。</li>
	SubtitleFormat *string `json:"SubtitleFormat,omitnil" name:"SubtitleFormat"`

	// 视频源语言。
	SourceLanguage *string `json:"SourceLanguage,omitnil" name:"SourceLanguage"`
}

type AsrWordsConfigureInfo struct {
	// 语音关键词识别任务开关，可选值：
	// <li>ON：开启语音关键词识别任务；</li>
	// <li>OFF：关闭语音关键词识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`
}

type AsrWordsConfigureInfoForUpdate struct {
	// 语音关键词识别任务开关，可选值：
	// <li>ON：开启语音关键词识别任务；</li>
	// <li>OFF：关闭语音关键词识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`
}

type AudioBeautifyConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 类型，可多选，可选值：
	// <li>declick：杂音去除</li>
	// <li>deesser：齿音压制</li>
	// 默认值：declick。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Types []*string `json:"Types,omitnil" name:"Types"`
}

type AudioDenoiseConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type AudioEnhanceConfig struct {
	// 音频降噪配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Denoise *AudioDenoiseConfig `json:"Denoise,omitnil" name:"Denoise"`

	// 音频分离配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Separate *AudioSeparateConfig `json:"Separate,omitnil" name:"Separate"`

	// 音量均衡配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeBalance *VolumeBalanceConfig `json:"VolumeBalance,omitnil" name:"VolumeBalance"`

	// 音频美化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Beautify *AudioBeautifyConfig `json:"Beautify,omitnil" name:"Beautify"`
}

type AudioSeparateConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 场景类型，可选值：
	// <li>normal：人声背景声场景</li>
	// <li>music：演唱伴奏场景</li>
	// 默认值：normal。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输出音轨，可选值：
	// <li>vocal：输出人声</li>
	// <li>background：应用场景为normal时输出背景声，应用场景为music时输出伴奏</li>
	// 默认值：vocal。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Track *string `json:"Track,omitnil" name:"Track"`
}

type AudioTemplateInfo struct {
	// 音频流的编码格式。
	// 当外层参数 Container 为 mp3 时，可选值为：
	// <li>libmp3lame。</li>
	// 当外层参数 Container 为 ogg 或 flac 时，可选值为：
	// <li>flac。</li>
	// 当外层参数 Container 为 m4a 时，可选值为：
	// <li>libfdk_aac；</li>
	// <li>libmp3lame；</li>
	// <li>ac3。</li>
	// 当外层参数 Container 为 mp4 或 flv 时，可选值为：
	// <li>libfdk_aac：更适合 mp4；</li>
	// <li>libmp3lame：更适合 flv。</li>
	// 当外层参数 Container 为 hls 时，可选值为：
	// <li>libfdk_aac；</li>
	// <li>libmp3lame。</li>
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。
	// 当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 音频流的采样率，可选值：
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	SampleRate *uint64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// 音频通道方式，可选值：
	// <li>1：单通道</li>
	// <li>2：双通道</li>
	// <li>6：立体声</li>
	// 当媒体的封装格式是音频格式时（flac，ogg，mp3，m4a）时，声道数不允许设为立体声。
	// 默认值：2。
	AudioChannel *int64 `json:"AudioChannel,omitnil" name:"AudioChannel"`
}

type AudioTemplateInfoForUpdate struct {
	// 音频流的编码格式。
	// 当外层参数 Container 为 mp3 时，可选值为：
	// <li>libmp3lame。</li>
	// 当外层参数 Container 为 ogg 或 flac 时，可选值为：
	// <li>flac。</li>
	// 当外层参数 Container 为 m4a 时，可选值为：
	// <li>libfdk_aac；</li>
	// <li>libmp3lame；</li>
	// <li>ac3。</li>
	// 当外层参数 Container 为 mp4 或 flv 时，可选值为：
	// <li>libfdk_aac：更适合 mp4；</li>
	// <li>libmp3lame：更适合 flv；</li>
	// <li>mp2。</li>
	// 当外层参数 Container 为 hls 时，可选值为：
	// <li>libfdk_aac；</li>
	// <li>libmp3lame。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。 当取值为 0，表示音频码率和原始音频保持一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 音频流的采样率，可选值：
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleRate *uint64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// 音频通道方式，可选值：
	// <li>1：单通道</li>
	// <li>2：双通道</li>
	// <li>6：立体声</li>
	// 当媒体的封装格式是音频格式时（flac，ogg，mp3，m4a）时，声道数不允许设为立体声。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioChannel *int64 `json:"AudioChannel,omitnil" name:"AudioChannel"`

	// 指定输出要保留的音频轨道。默认是全部保留源的。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamSelects []*int64 `json:"StreamSelects,omitnil" name:"StreamSelects"`
}

type AwsS3FileUploadTrigger struct {
	// 绑定的 AWS S3 存储桶。
	S3Bucket *string `json:"S3Bucket,omitnil" name:"S3Bucket"`

	// 绑定的桶所在 AWS 区域，目前支持：  
	// us-east-1  
	// eu-west-3
	S3Region *string `json:"S3Region,omitnil" name:"S3Region"`

	// 绑定的输入路径目录，必须为绝对路径，即以 `/` 开头和结尾。如`/movie/201907/`，不填代表根目录`/`。	
	Dir *string `json:"Dir,omitnil" name:"Dir"`

	// 允许触发的文件格式列表，如 ["mp4", "flv", "mov"]。不填代表所有格式的文件都可以触发工作流。	
	Formats []*string `json:"Formats,omitnil" name:"Formats"`

	// 绑定的 AWS S3 存储桶的秘钥ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3SecretId *string `json:"S3SecretId,omitnil" name:"S3SecretId"`

	// 绑定的 AWS S3 存储桶的秘钥Key。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3SecretKey *string `json:"S3SecretKey,omitnil" name:"S3SecretKey"`

	// 绑定的 AWS S3 存储桶对应的 SQS事件队列。
	// 注意：队列和桶需要在同一区域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AwsSQS *AwsSQS `json:"AwsSQS,omitnil" name:"AwsSQS"`
}

type AwsSQS struct {
	// SQS 队列区域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQSRegion *string `json:"SQSRegion,omitnil" name:"SQSRegion"`

	// SQS 队列名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQSQueueName *string `json:"SQSQueueName,omitnil" name:"SQSQueueName"`

	// 读写SQS的秘钥id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3SecretId *string `json:"S3SecretId,omitnil" name:"S3SecretId"`

	// 读写SQS的秘钥key。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3SecretKey *string `json:"S3SecretKey,omitnil" name:"S3SecretKey"`
}

// Predefined struct for user
type BatchDeleteStreamLinkFlowRequestParams struct {
	// EventId。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// Event关联的流Id数组，如果不传默认删除Event下面的所有媒体传输流。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`
}

type BatchDeleteStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// EventId。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// Event关联的流Id数组，如果不传默认删除Event下面的所有媒体传输流。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`
}

func (r *BatchDeleteStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "FlowIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteStreamLinkFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BatchDeleteStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteStreamLinkFlowResponseParams `json:"Response"`
}

func (r *BatchDeleteStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStartStreamLinkFlowRequestParams struct {
	// EventId。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// Event关联的流Id数组，如果不传默认启动Event下面的所有媒体传输流。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`
}

type BatchStartStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// EventId。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// Event关联的流Id数组，如果不传默认启动Event下面的所有媒体传输流。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`
}

func (r *BatchStartStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStartStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "FlowIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchStartStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStartStreamLinkFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BatchStartStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *BatchStartStreamLinkFlowResponseParams `json:"Response"`
}

func (r *BatchStartStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStartStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopStreamLinkFlowRequestParams struct {
	// EventId。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 流Id，如果不传默认停止Event下所有的媒体传输流。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`
}

type BatchStopStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// EventId。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 流Id，如果不传默认停止Event下所有的媒体传输流。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`
}

func (r *BatchStopStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "FlowIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchStopStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopStreamLinkFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BatchStopStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *BatchStopStreamLinkFlowResponseParams `json:"Response"`
}

func (r *BatchStopStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassificationConfigureInfo struct {
	// 智能分类任务开关，可选值：
	// <li>ON：开启智能分类任务；</li>
	// <li>OFF：关闭智能分类任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type ClassificationConfigureInfoForUpdate struct {
	// 智能分类任务开关，可选值：
	// <li>ON：开启智能分类任务；</li>
	// <li>OFF：关闭智能分类任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type ColorEnhanceConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 类型，可选值：
	// <li>weak</li>
	// <li>normal</li>
	// <li>strong</li>
	// 默认值：weak。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type ComposeAudioItem struct {
	// 元素对应媒体信息。
	SourceMedia *ComposeSourceMedia `json:"SourceMedia,omitnil" name:"SourceMedia"`

	// 元素在轨道时间轴上的时间信息，不填则紧跟上一个元素。
	TrackTime *ComposeTrackTime `json:"TrackTime,omitnil" name:"TrackTime"`

	// 对音频进行操作，如静音等。
	AudioOperations []*ComposeAudioOperation `json:"AudioOperations,omitnil" name:"AudioOperations"`
}

type ComposeAudioOperation struct {
	// 音频操作类型，取值有：
	// <li>Volume：音量调节。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	//  当 Type = Volume 时有效。音量调节参数，取值范围: 0~5。 
	// <li>0 表示静音。</li>
	// <li>小于1 表示降低音量。</li>
	// <li>1 表示不变。</li>
	// <li>大于1表示升高音量。</li>
	Volume *float64 `json:"Volume,omitnil" name:"Volume"`
}

type ComposeAudioStream struct {
	// 音频流的编码方式，可选值：
	// <li>AAC：AAC 编码（默认），用于容器为 mp4。</li>
	// <li>MP3：mp3 编码，用于容器为 mp3。</li>
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// 音频流的采样率，单位：Hz，可选值：
	// <li>16000（默认）</li>
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	SampleRate *int64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// 声道数，可选值：
	// <li>1：单声道 。</li>
	// <li>2：双声道（默认）。</li>
	AudioChannel *int64 `json:"AudioChannel,omitnil" name:"AudioChannel"`
}

type ComposeCanvas struct {
	// 背景颜色对应的 RGB 参考值，取值格式： #RRGGBB，如 #F0F0F0 。 
	// 默认值：#000000（黑色）。
	Color *string `json:"Color,omitnil" name:"Color"`

	// 画布宽度，即输出视频的宽度，取值范围：0~ 3840，单位：px。  
	// 默认值：0，表示和第一个视频宽度一致。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 画布高度，即输出视频的高度，取值范围：0~ 3840，单位：px。  
	// 默认值：0，表示和第一个视频高度一致。
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type ComposeEmptyItem struct {
	// 元素时长，时间支持：
	// <li>以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒。</li>
	Duration *string `json:"Duration,omitnil" name:"Duration"`
}

type ComposeImageItem struct {
	// 元素对应媒体信息。
	SourceMedia *ComposeSourceMedia `json:"SourceMedia,omitnil" name:"SourceMedia"`

	// 元素在轨道时间轴上的时间信息，不填则紧跟上一个元素。
	TrackTime *ComposeTrackTime `json:"TrackTime,omitnil" name:"TrackTime"`

	// 元素中心点距离画布原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示元素 XPos 为画布宽度指定百分比的位置，如 10% 表示 XPos 为画布宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示元素 XPos 单位为像素，如 100px 表示 XPos 为100像素。</li>
	// 默认：50%。
	XPos *string `json:"XPos,omitnil" name:"XPos"`

	// 元素中心点距离画布原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示元素 YPos 为画布高度指定百分比的位置，如 10% 表示 YPos 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示元素 YPos 单位为像素，如 100px 表示 YPos 为100像素。</li>
	// 默认：50%。
	YPos *string `json:"YPos,omitnil" name:"YPos"`

	// 视频片段的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示元素 Width 为画布宽度的百分比大小，如 10% 表示 Width 为画布宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示元素 Width 单位为像素，如 100px 表示 Width 为100像素。</li>
	// 为空（或0） 的场景：
	// <li>当 Width、Height 均为空，则 Width 和 Height 取源素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按源素材比例缩放。</li>
	// <li>当 Width 非空，Height 为空，则 Height 按源素材比例缩放。</li>
	Width *string `json:"Width,omitnil" name:"Width"`

	// 元素的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示元素 Height 为画布高度的百分比大小，如 10% 表示 Height 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示元素 Height 单位为像素，如 100px 表示 Height 为100像素。</li>
	// 为空（或0） 的场景：
	// <li>当 Width、Height 均为空，则 Width 和 Height 取源素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按源素材比例缩放。</li>
	// <li>当 Width 非空，Height 为空，则 Height 按源素材比例缩放。</li>
	Height *string `json:"Height,omitnil" name:"Height"`

	// 对图像画面进行的操作，如图像旋转等。
	ImageOperations []*ComposeImageOperation `json:"ImageOperations,omitnil" name:"ImageOperations"`
}

type ComposeImageOperation struct {
	// 类型，取值有：
	// <li>Rotate：图像旋转。</li>
	// <li>Flip：图像翻转。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 当 Type = Rotate 时有效。图像以中心点为原点进行旋转的角度，取值范围0~360。
	RotateAngle *float64 `json:"RotateAngle,omitnil" name:"RotateAngle"`

	// 当 Type = Flip 时有效。图像翻转动作，取值有： 
	// <li>Horizental：水平翻转，即左右镜像。</li>
	// <li>Vertical：垂直翻转，即上下镜像。</li>
	FlipType *string `json:"FlipType,omitnil" name:"FlipType"`
}

type ComposeMediaConfig struct {
	// 合成目标视频信息。
	TargetInfo *ComposeTargetInfo `json:"TargetInfo,omitnil" name:"TargetInfo"`

	// 合成目标视频的画布信息。
	Canvas *ComposeCanvas `json:"Canvas,omitnil" name:"Canvas"`

	// 全局样式，和轨道 Tracks 配合使用，用于定于样式，如字幕样式。
	Styles []*ComposeStyles `json:"Styles,omitnil" name:"Styles"`

	// 用于描述合成视频的轨道列表，包括：视频、音频、图片、文字等元素组成的多个轨道信息。关于轨道和时间：
	// <ul><li>轨道时间轴即为目标视频时间轴。</li><li>时间轴上相同时间点的不同轨道上的元素会重叠：</li><ul><li>视频、图片、文字：按轨道顺序进行图像的叠加，轨道顺序靠前的在上面。</li><li>音频 ：进行混音。</li></ul></ul>注意：同一轨道中各个元素（除字幕元素外）的轨道时间不能重叠。
	Tracks []*ComposeMediaTrack `json:"Tracks,omitnil" name:"Tracks"`
}

type ComposeMediaItem struct {
	// 元素类型。取值有：
	// <li>Video：视频元素。</li>
	// <li>Audio：音频元素。</li>
	// <li>Image：图片元素。</li>
	// <li>Transition：转场元素。</li>
	// <li>Subtitle：字幕元素。</li>
	// <li>Empty：空白元素。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 视频元素，当 Type = Video 时有效。
	Video *ComposeVideoItem `json:"Video,omitnil" name:"Video"`

	// 音频元素，当 Type = Audio 时有效。
	Audio *ComposeAudioItem `json:"Audio,omitnil" name:"Audio"`

	// 图片元素，当 Type = Image 时有效。
	Image *ComposeImageItem `json:"Image,omitnil" name:"Image"`

	// 转场元素，当 Type = Transition 时有效。
	Transition *ComposeTransitionItem `json:"Transition,omitnil" name:"Transition"`

	// 字幕元素，当 Type = Subtitle 是有效。
	Subtitle *ComposeSubtitleItem `json:"Subtitle,omitnil" name:"Subtitle"`

	// 空白元素，当 Type = Empty 时有效。用于时间轴的占位。
	Empty *ComposeEmptyItem `json:"Empty,omitnil" name:"Empty"`
}

type ComposeMediaTrack struct {
	// 轨道类型，取值有：<ul><li>Video ：视频轨道。视频轨道可由以下元素组成：</li><ul><li>Video 元素</li><li>Image 元素</li><li>Transition 元素</li><li>Empty 元素</li></ul><li>Audio ：音频轨道。音频轨道可由以下元素组成：</li><ul><li>Audio 元素</li><li>Transition 元素</li><li>Empty 元素</li></ul><li>Title：文字轨道。文字轨道可由以下元素组成：</li><ul><li>Subtitle 元素</li></ul>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 轨道上的元素列表。
	Items []*ComposeMediaItem `json:"Items,omitnil" name:"Items"`
}

type ComposeSourceMedia struct {
	// 媒体对应的素材ID，即 FileInfos 列表中对应素材的 ID。
	FileId *string `json:"FileId,omitnil" name:"FileId"`

	// 媒体位于素材的起始时间，时间点支持 s、% 两种格式：
	// <li>当字符串以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// <li>当字符串以 % 结尾，表示时间点为素材时长的百分比大小，如10%表示时间点为素材第10% 的时刻。</li>
	// 默认：0s
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 媒体位于素材的结束时间，和 StartTime 构成媒体在源素材的时间区间，时间点支持 s、% 两种格式：
	// <li>当字符串以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// <li>当字符串以 % 结尾，表示时间点为素材时长的百分比大小，如10%表示时间点为素材第10%的时间。</li>
	// 默认：如果对应轨道时长有设置，则默认轨道时长，否则为素材时长，无时长的素材默认为 1 秒。
	// 注意：至少需要大于 StartTime 0.02 秒。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type ComposeStyles struct {
	// 样式 Id，用于和轨道元素中的样式关联。
	// 注意：允许字母、数字、-、_ 组合，最长 32 字符。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 样式类型，取值有：
	// <li>Subtitle：字幕样式。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 字幕样式信息，当 Type = Subtitle 时有效。
	Subtitle *ComposeSubtitleStyle `json:"Subtitle,omitnil" name:"Subtitle"`
}

type ComposeSubtitleItem struct {
	// 字幕样式，Styles 列表中对应的 Subtitle样式的 ID。
	StyleId *string `json:"StyleId,omitnil" name:"StyleId"`

	// 字幕文本。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 元素在轨道时间轴上的时间信息，不填则紧跟上一个元素。	
	TrackTime *ComposeTrackTime `json:"TrackTime,omitnil" name:"TrackTime"`
}

type ComposeSubtitleStyle struct {
	// 字幕高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示为画布高度的百分比大小，如 10% 表示为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示单位为像素，如 100px 表示为100像素。</li>
	// 默认为 FontSize 大小。
	Height *string `json:"Height,omitnil" name:"Height"`

	// 字幕距离下边框距离，支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示为画布高度的百分比大小，如 10% 表示为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示单位为像素，如 100px 表示为100像素。</li>
	// 默认：0px
	MarginBottom *string `json:"MarginBottom,omitnil" name:"MarginBottom"`

	// 字体类型，支持：
	// <li>SimHei：黑体（默认）。</li>
	// <li>SimSun：宋体。</li>
	FontType *string `json:"FontType,omitnil" name:"FontType"`

	// 字体大小，支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示为画布高度的百分比大小，如 10% 表示为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示单位为像素，如 100px 表示为100像素。</li>
	// 默认：2%
	FontSize *string `json:"FontSize,omitnil" name:"FontSize"`

	// 是否使用粗体，和字体相关，可选值：
	// <li>0：否（默认）。</li>
	// <li>1：是。</li>
	FontBold *int64 `json:"FontBold,omitnil" name:"FontBold"`

	// 是否使用斜体，和字体相关，可选值：
	// <li>0：否（默认）。</li>
	// <li>1：是。</li>
	FontItalic *int64 `json:"FontItalic,omitnil" name:"FontItalic"`

	// 字体颜色，格式：#RRGGBBAA。  
	// 默认值：0x000000FF（黑色）。  
	// 注意：其中 AA 部分指的是透明度，为可选。
	FontColor *string `json:"FontColor,omitnil" name:"FontColor"`

	// 文字对齐方式：
	// <li>Center：居中（默认）。</li>
	// <li>Left：左对齐。</li>
	// <li>Right：右对齐。</li>
	FontAlign *string `json:"FontAlign,omitnil" name:"FontAlign"`

	// 用于字幕对齐留白：
	// <li>FontAlign=Left 时，表示距离左边距离。</li>
	// <li>FontAlign=Right时，表示距离右边距离。</li>
	// 支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示为画布宽度的百分比大小，如 10% 表示为画布宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示单位为像素，如 100px 表示为100像素。</li>
	FontAlignMargin *string `json:"FontAlignMargin,omitnil" name:"FontAlignMargin"`

	// 字体边框宽度，支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示为画布高度的百分比大小，如 10% 表示为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示单位为像素，如 100px 表示为100像素。</li>
	// 默认： 0，表示不需要边框。
	BorderWidth *string `json:"BorderWidth,omitnil" name:"BorderWidth"`

	// 边框颜色，当 BorderWidth 不为 0 时生效，其值格式和 FontColor 一致。
	BorderColor *string `json:"BorderColor,omitnil" name:"BorderColor"`

	// 文字底色，其值格式和 FontColor 一致。  
	// 默认为空， 表示不使用底色。
	BottomColor *string `json:"BottomColor,omitnil" name:"BottomColor"`
}

type ComposeTargetInfo struct {
	// 封装容器格式，可选值：
	// <li>mp4：视频文件（默认）。</li>
	// <li>mp3：纯音频文件。</li>
	Container *string `json:"Container,omitnil" name:"Container"`

	// 是否去除视频数据，可选值：
	// <li>0：保留（默认）。</li>
	// <li>1：去除。</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留（默认）。</li>
	// <li>1：去除。</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil" name:"RemoveAudio"`

	// 输出视频流信息。
	VideoStream *ComposeVideoStream `json:"VideoStream,omitnil" name:"VideoStream"`

	// 输出音频流信息。
	AudioStream *ComposeAudioStream `json:"AudioStream,omitnil" name:"AudioStream"`
}

type ComposeTrackTime struct {
	// 元素在轨道上的起始时间，时间点支持：
	// <li>以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// 注意：不填则默认为前一个元素的结束时间，此时可以通过 ComposeEmptyItem 元素来进行占位，实现轨道起始时间设置。
	Start *string `json:"Start,omitnil" name:"Start"`

	// 元素时长，时间支持：
	// <li>以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// 默认：取对应 ComposeSourceMedia 媒体的有效时长（即 EndTime-StartTime），没有 ComposeSourceMedia 则默认为 1 秒。
	Duration *string `json:"Duration,omitnil" name:"Duration"`
}

type ComposeTransitionItem struct {
	// 元素时长，时间支持：<li>以 s 结尾，表示时间点单位为秒，如 3s 表示时间点为第3秒。</li>
	// 默认：1s
	// 注意：
	// <li>必须是整数s，否则向下取整。</li>
	// <li>转场 前后必须紧挨着两个不为 Empty 的元素。</li>
	// <li>转场 Duration 必须小于前一个元素的 Duration，同时必须小于后一个元素的 Duration。</li>
	// <li>进行转场处理的两个元素，第二个元素在轨道上的起始时间会自动调整为前一个元素的结束时间减去转场的 Duration。</li>
	Duration *string `json:"Duration,omitnil" name:"Duration"`

	// 转场操作列表。
	// 默认：淡入淡出。
	// 注意：图像转场操作和音频转场操作各自最多支持一个。
	Transitions []*ComposeTransitionOperation `json:"Transitions,omitnil" name:"Transitions"`
}

type ComposeTransitionOperation struct {
	// 转场类型。
	// 
	// 图像的转场操作，用于两个视频片段图像间的转场处理：
	// <li>ImageFadeInFadeOut：图像淡入淡出。</li>
	// <li>BowTieHorizontal：水平蝴蝶结。</li>
	// <li>BowTieVertical：垂直蝴蝶结。</li>
	// <li>ButterflyWaveScrawler：晃动。</li>
	// <li>Cannabisleaf：枫叶。</li>
	// <li>Circle：弧形收放。</li>
	// <li>CircleCrop：圆环聚拢。</li>
	// <li>Circleopen：椭圆聚拢。</li>
	// <li>Crosswarp：横向翘曲。</li>
	// <li>Cube：立方体。</li>
	// <li>DoomScreenTransition：幕布。</li>
	// <li>Doorway：门廊。</li>
	// <li>Dreamy：波浪。</li>
	// <li>DreamyZoom：水平聚拢。</li>
	// <li>FilmBurn：火烧云。</li>
	// <li>GlitchMemories：抖动。</li>
	// <li>Heart：心形。</li>
	// <li>InvertedPageCurl：翻页。</li>
	// <li>Luma：腐蚀。</li>
	// <li>Mosaic：九宫格。</li>
	// <li>Pinwheel：风车。</li>
	// <li>PolarFunction：椭圆扩散。</li>
	// <li>PolkaDotsCurtain：弧形扩散。</li>
	// <li>Radial：雷达扫描。</li>
	// <li>RotateScaleFade：上下收放。</li>
	// <li>Squeeze：上下聚拢。</li>
	// <li>Swap：放大切换。</li>
	// <li>Swirl：螺旋。</li>
	// <li>UndulatingBurnOutSwirl：水流蔓延。</li>
	// <li>Windowblinds：百叶窗。</li>
	// <li>WipeDown：向下收起。</li>
	// <li>WipeLeft：向左收起。</li>
	// <li>WipeRight：向右收起。</li>
	// <li>WipeUp：向上收起。</li>
	// <li>ZoomInCircles：水波纹。</li> 
	// 音频的转场操作，用于两个音频片段间的转场处理：
	// <li>AudioFadeInFadeOut：声音淡入淡出。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

type ComposeVideoItem struct {
	// 元素对应媒体信息。
	SourceMedia *ComposeSourceMedia `json:"SourceMedia,omitnil" name:"SourceMedia"`

	// 元素在轨道时间轴上的时间信息，不填则紧跟上一个元素。
	TrackTime *ComposeTrackTime `json:"TrackTime,omitnil" name:"TrackTime"`

	// 元素中心点距离画布原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示元素 XPos 为画布宽度指定百分比的位置，如 10% 表示 XPos 为画布宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示元素 XPos 单位为像素，如 100px 表示 XPos 为100像素。</li>
	// 默认：50%。
	XPos *string `json:"XPos,omitnil" name:"XPos"`

	// 元素中心点距离画布原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示元素 YPos 为画布高度指定百分比的位置，如 10% 表示 YPos 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示元素 YPos 单位为像素，如 100px 表示 YPos 为100像素。</li>
	// 默认：50%。
	YPos *string `json:"YPos,omitnil" name:"YPos"`

	// 视频片段的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示元素 Width 为画布宽度的百分比大小，如 10% 表示 Width 为画布宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示元素 Width 单位为像素，如 100px 表示 Width 为100像素。</li>
	// 为空（或0） 的场景：
	// <li>当 Width、Height 均为空，则 Width 和 Height 取源素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按源素材比例缩放。</li>
	// <li>当 Width 非空，Height 为空，则 Height 按源素材比例缩放。</li>
	Width *string `json:"Width,omitnil" name:"Width"`

	// 元素的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示元素 Height 为画布高度的百分比大小，如 10% 表示 Height 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示元素 Height 单位为像素，如 100px 表示 Height 为100像素。</li>
	// 为空（或0） 的场景：
	// <li>当 Width、Height 均为空，则 Width 和 Height 取源素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按源素材比例缩放。</li>
	// <li>当 Width 非空，Height 为空，则 Height 按源素材比例缩放。</li>
	Height *string `json:"Height,omitnil" name:"Height"`

	// 对图像画面进行的操作，如图像旋转等。
	ImageOperations []*ComposeImageOperation `json:"ImageOperations,omitnil" name:"ImageOperations"`

	// 对音频进行操作，如静音等。
	AudioOperations []*ComposeAudioOperation `json:"AudioOperations,omitnil" name:"AudioOperations"`
}

type ComposeVideoStream struct {
	// 视频流的编码方式，可选值：
	// <li>H.264：H.264 编码（默认）。</li>
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// 视频帧率，取值范围：[0, 60]，单位：Hz。  
	// 默认值：0，表示和第一个视频帧率一致。
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`
}

type ContentReviewTemplateItem struct {
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 鉴黄控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitnil" name:"PornConfigure"`

	// 涉敏控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitnil" name:"TerrorismConfigure"`

	// 涉敏控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitnil" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitnil" name:"ProhibitedConfigure"`

	// 用户自定义内容审核控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitnil" name:"UserDefineConfigure"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 模板类型，取值范围：
	// * Preset：系统预置模板；
	// * Custom：用户自定义模板。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type CosFileUploadTrigger struct {
	// 工作流绑定的 COS Bucket 名，如 TopRankVideo-125xxx88。
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`

	// 工作流绑定的 COS Bucket 所属园区，如 ap-chongiqng。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 工作流绑定的输入路径目录，必须为绝对路径，即以 `/` 开头和结尾。如`/movie/201907/`，不填代表根目录`/`。
	Dir *string `json:"Dir,omitnil" name:"Dir"`

	// 工作流允许触发的文件格式列表，如 ["mp4", "flv", "mov"]。不填代表所有格式的文件都可以触发工作流。
	Formats []*string `json:"Formats,omitnil" name:"Formats"`
}

type CosInputInfo struct {
	// 媒体处理对象文件所在的 COS Bucket 名，如 TopRankVideo-125xxx88。
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`

	// 媒体处理对象文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 媒体处理对象文件的输入路径，如`/movie/201907/WildAnimal.mov`。
	Object *string `json:"Object,omitnil" name:"Object"`
}

type CosOutputStorage struct {
	// 媒体处理生成的文件输出的目标 Bucket 名，如 TopRankVideo-125xxx88。如果不填，表示继承上层。
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`

	// 媒体处理生成的文件输出的目标 Bucket 的园区，如 ap-chongqing。如果不填，表示继承上层。
	Region *string `json:"Region,omitnil" name:"Region"`
}

type CoverConfigureInfo struct {
	// 智能封面任务开关，可选值：
	// <li>ON：开启智能封面任务；</li>
	// <li>OFF：关闭智能封面任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type CoverConfigureInfoForUpdate struct {
	// 智能封面任务开关，可选值：
	// <li>ON：开启智能封面任务；</li>
	// <li>OFF：关闭智能封面任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

// Predefined struct for user
type CreateAIAnalysisTemplateRequestParams struct {
	// 视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitnil" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitnil" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitnil" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitnil" name:"FrameTagConfigure"`
}

type CreateAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitnil" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitnil" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitnil" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitnil" name:"FrameTagConfigure"`
}

func (r *CreateAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIAnalysisTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ClassificationConfigure")
	delete(f, "TagConfigure")
	delete(f, "CoverConfigure")
	delete(f, "FrameTagConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIAnalysisTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIAnalysisTemplateResponseParams struct {
	// 视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIAnalysisTemplateResponseParams `json:"Response"`
}

func (r *CreateAIAnalysisTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIAnalysisTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIRecognitionTemplateRequestParams struct {
	// 视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitnil" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitnil" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitnil" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitnil" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitnil" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	TranslateConfigure *TranslateConfigureInfo `json:"TranslateConfigure,omitnil" name:"TranslateConfigure"`
}

type CreateAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitnil" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitnil" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitnil" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitnil" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitnil" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	TranslateConfigure *TranslateConfigureInfo `json:"TranslateConfigure,omitnil" name:"TranslateConfigure"`
}

func (r *CreateAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIRecognitionTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "FaceConfigure")
	delete(f, "OcrFullTextConfigure")
	delete(f, "OcrWordsConfigure")
	delete(f, "AsrFullTextConfigure")
	delete(f, "AsrWordsConfigure")
	delete(f, "TranslateConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIRecognitionTemplateResponseParams struct {
	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIRecognitionTemplateResponseParams `json:"Response"`
}

func (r *CreateAIRecognitionTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIRecognitionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdaptiveDynamicStreamingTemplateRequestParams struct {
	// 自适应转码格式，取值范围：
	// <li>HLS，</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil" name:"Format"`

	// 转自适应码流输出子流参数信息，最多输出10路子流。
	// 注意：各个子流的帧率必须保持一致；如果不一致，采用第一个子流的帧率作为输出帧率。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil" name:"StreamInfos"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	// 默认为否。
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	// 默认为否。
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil" name:"DisableHigherVideoResolution"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`
}

type CreateAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 自适应转码格式，取值范围：
	// <li>HLS，</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil" name:"Format"`

	// 转自适应码流输出子流参数信息，最多输出10路子流。
	// 注意：各个子流的帧率必须保持一致；如果不一致，采用第一个子流的帧率作为输出帧率。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil" name:"StreamInfos"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	// 默认为否。
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	// 默认为否。
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil" name:"DisableHigherVideoResolution"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`
}

func (r *CreateAdaptiveDynamicStreamingTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdaptiveDynamicStreamingTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Format")
	delete(f, "StreamInfos")
	delete(f, "Name")
	delete(f, "DisableHigherVideoBitrate")
	delete(f, "DisableHigherVideoResolution")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdaptiveDynamicStreamingTemplateResponseParams struct {
	// 自适应转码模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAdaptiveDynamicStreamingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAdaptiveDynamicStreamingTemplateResponseParams `json:"Response"`
}

func (r *CreateAdaptiveDynamicStreamingTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdaptiveDynamicStreamingTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAnimatedGraphicsTemplateRequestParams struct {
	// 帧率，取值范围：[1, 30]，单位：Hz。
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 动图格式，取值为 gif 和 webp。默认为 gif。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitnil" name:"Quality"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`
}

type CreateAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 帧率，取值范围：[1, 30]，单位：Hz。
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 动图格式，取值为 gif 和 webp。默认为 gif。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitnil" name:"Quality"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`
}

func (r *CreateAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Fps")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Quality")
	delete(f, "Name")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAnimatedGraphicsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAnimatedGraphicsTemplateResponseParams struct {
	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAnimatedGraphicsTemplateResponseParams `json:"Response"`
}

func (r *CreateAnimatedGraphicsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContentReviewTemplateRequestParams struct {
	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 令人反感的信息的控制参数。
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitnil" name:"PornConfigure"`

	// 令人不安全的信息的控制参数。
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitnil" name:"TerrorismConfigure"`

	// 令人不适宜的信息的控制参数。
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitnil" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	// 注意：此参数尚未支持。
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitnil" name:"ProhibitedConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitnil" name:"UserDefineConfigure"`
}

type CreateContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 令人反感的信息的控制参数。
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitnil" name:"PornConfigure"`

	// 令人不安全的信息的控制参数。
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitnil" name:"TerrorismConfigure"`

	// 令人不适宜的信息的控制参数。
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitnil" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	// 注意：此参数尚未支持。
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitnil" name:"ProhibitedConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitnil" name:"UserDefineConfigure"`
}

func (r *CreateContentReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContentReviewTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "PornConfigure")
	delete(f, "TerrorismConfigure")
	delete(f, "PoliticalConfigure")
	delete(f, "ProhibitedConfigure")
	delete(f, "UserDefineConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContentReviewTemplateResponseParams struct {
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateContentReviewTemplateResponseParams `json:"Response"`
}

func (r *CreateContentReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContentReviewTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageSpriteTemplateRequestParams struct {
	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitnil" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitnil" name:"ColumnCount"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 图片格式，取值为 jpg、png、webp。默认为 jpg。
	Format *string `json:"Format,omitnil" name:"Format"`
}

type CreateImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitnil" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitnil" name:"ColumnCount"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 图片格式，取值为 jpg、png、webp。默认为 jpg。
	Format *string `json:"Format,omitnil" name:"Format"`
}

func (r *CreateImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageSpriteTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SampleType")
	delete(f, "SampleInterval")
	delete(f, "RowCount")
	delete(f, "ColumnCount")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "FillType")
	delete(f, "Comment")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageSpriteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageSpriteTemplateResponseParams struct {
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageSpriteTemplateResponseParams `json:"Response"`
}

func (r *CreateImageSpriteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageSpriteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInput struct {
	// 输入名称，可填大小写、数字和下划线，长度为[1, 32]。
	InputName *string `json:"InputName,omitnil" name:"InputName"`

	// 输入的协议，可选[SRT|RTP|RTMP|RTMP_PULL]。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 输入描述，长度为[0, 255]。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 输入的IP白名单，格式为CIDR。
	AllowIpList []*string `json:"AllowIpList,omitnil" name:"AllowIpList"`

	// 输入的SRT配置信息。
	SRTSettings *CreateInputSRTSettings `json:"SRTSettings,omitnil" name:"SRTSettings"`

	// 输入的RTP配置信息。
	RTPSettings *CreateInputRTPSettings `json:"RTPSettings,omitnil" name:"RTPSettings"`

	// 输入的主备开关，可选[OPEN|CLOSE]，默认为CLOSE。
	FailOver *string `json:"FailOver,omitnil" name:"FailOver"`

	// 输入的RTMP_PULL配置信息。
	RTMPPullSettings *CreateInputRTMPPullSettings `json:"RTMPPullSettings,omitnil" name:"RTMPPullSettings"`

	// 输入的RTSP_PULL配置信息。
	RTSPPullSettings *CreateInputRTSPPullSettings `json:"RTSPPullSettings,omitnil" name:"RTSPPullSettings"`

	// 输入的HLS_PULL配置信息。
	HLSPullSettings *CreateInputHLSPullSettings `json:"HLSPullSettings,omitnil" name:"HLSPullSettings"`

	// 延播平滑吐流配置信息。
	ResilientStream *ResilientStreamConf `json:"ResilientStream,omitnil" name:"ResilientStream"`

	// 绑定的输入安全组 ID。 
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type CreateInputHLSPullSettings struct {
	// HLS源站的源站地址，有且只能有一个。
	SourceAddresses []*HLSPullSourceAddress `json:"SourceAddresses,omitnil" name:"SourceAddresses"`
}

type CreateInputRTMPPullSettings struct {
	// RTMP源站的源站地址，有且只能有一个。
	SourceAddresses []*RTMPPullSourceAddress `json:"SourceAddresses,omitnil" name:"SourceAddresses"`
}

type CreateInputRTPSettings struct {
	// 默认为“none”，可选值['none']。
	FEC *string `json:"FEC,omitnil" name:"FEC"`

	// 空闲超时时间，默认5000，单位ms，范围为[1000, 10000]。
	IdleTimeout *int64 `json:"IdleTimeout,omitnil" name:"IdleTimeout"`
}

type CreateInputRTSPPullSettings struct {
	// RTSP源站的源站地址，有且只能有一个。
	SourceAddresses []*RTSPPullSourceAddress `json:"SourceAddresses,omitnil" name:"SourceAddresses"`
}

type CreateInputSRTSettings struct {
	// SRT模式，可选[LISTENER|CALLER]，默认为LISTENER。
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 流Id，可选大小写字母、数字和特殊字符（.#!:&,=_-），长度为0~512。具体格式可以参考：https://github.com/Haivision/srt/blob/master/docs/features/access-control.md#standard-keys。
	StreamId *string `json:"StreamId,omitnil" name:"StreamId"`

	// 延迟，默认0，单位ms，范围为[0, 3000]。
	Latency *int64 `json:"Latency,omitnil" name:"Latency"`

	// 接收延迟，默认120，单位ms，范围为[0, 3000]。
	RecvLatency *int64 `json:"RecvLatency,omitnil" name:"RecvLatency"`

	// 对端延迟，默认0，单位ms，范围为[0, 3000]。
	PeerLatency *int64 `json:"PeerLatency,omitnil" name:"PeerLatency"`

	// 对端超时时间，默认5000，单位ms，范围为[1000, 10000]。
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitnil" name:"PeerIdleTimeout"`

	// 解密密钥，默认为空，表示不加密。只可填ascii码值，长度为[10, 79]。
	Passphrase *string `json:"Passphrase,omitnil" name:"Passphrase"`

	// 密钥长度，默认为0，可选[0|16|24|32]。
	PbKeyLen *int64 `json:"PbKeyLen,omitnil" name:"PbKeyLen"`

	// SRT对端地址，当Mode为CALLER时必填，且只能填1组。
	SourceAddresses []*SRTSourceAddressReq `json:"SourceAddresses,omitnil" name:"SourceAddresses"`
}

type CreateOutputInfo struct {
	// 输出的名称。
	OutputName *string `json:"OutputName,omitnil" name:"OutputName"`

	// 输出描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 输出协议，可选[SRT|RTP|RTMP|RTMP_PULL]。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 输出地区。
	OutputRegion *string `json:"OutputRegion,omitnil" name:"OutputRegion"`

	// 输出的SRT的配置。
	SRTSettings *CreateOutputSRTSettings `json:"SRTSettings,omitnil" name:"SRTSettings"`

	// 输出的RTMP的配置。
	RTMPSettings *CreateOutputRTMPSettings `json:"RTMPSettings,omitnil" name:"RTMPSettings"`

	// 输出的RTP的配置。
	RTPSettings *CreateOutputInfoRTPSettings `json:"RTPSettings,omitnil" name:"RTPSettings"`

	// IP白名单列表，格式为CIDR，如0.0.0.0/0。
	// 当Protocol为RTMP_PULL有效，为空代表不限制客户端IP。
	AllowIpList []*string `json:"AllowIpList,omitnil" name:"AllowIpList"`

	// 最大拉流并发数，最大4，默认4。
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitnil" name:"MaxConcurrent"`

	// 绑定的输入安全组 ID。 
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type CreateOutputInfoRTPSettings struct {
	// 转推的目标地址，可填1~2个。
	Destinations []*CreateOutputRTPSettingsDestinations `json:"Destinations,omitnil" name:"Destinations"`

	// 只能填none。
	FEC *string `json:"FEC,omitnil" name:"FEC"`

	// 空闲超时时间，单位ms。
	IdleTimeout *int64 `json:"IdleTimeout,omitnil" name:"IdleTimeout"`
}

type CreateOutputRTMPSettings struct {
	// 转推的目标地址，可填1~2个。
	Destinations []*CreateOutputRtmpSettingsDestinations `json:"Destinations,omitnil" name:"Destinations"`

	// RTMP的Chunk大小，范围为[4096, 40960]。
	ChunkSize *int64 `json:"ChunkSize,omitnil" name:"ChunkSize"`
}

type CreateOutputRTPSettingsDestinations struct {
	// 转推的目标IP。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 转推的目标端口。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type CreateOutputRtmpSettingsDestinations struct {
	// 转推的URL，格式如：rtmp://domain/live。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 转推的StreamKey，格式如：stream?key=value。
	StreamKey *string `json:"StreamKey,omitnil" name:"StreamKey"`
}

type CreateOutputSRTSettings struct {
	// 转推的目标地址，当Mode为CALLER时必填，且只能填1组。
	Destinations []*CreateOutputSRTSettingsDestinations `json:"Destinations,omitnil" name:"Destinations"`

	// 转推SRT的流Id，可选大小写字母、数字和特殊字符（.#!:&,=_-），长度为0~512。
	StreamId *string `json:"StreamId,omitnil" name:"StreamId"`

	// 转推SRT的总延迟，默认0，单位ms，范围为[0, 3000]。
	Latency *int64 `json:"Latency,omitnil" name:"Latency"`

	// 转推SRT的接收延迟，默认120，单位ms，范围为[0, 3000]。
	RecvLatency *int64 `json:"RecvLatency,omitnil" name:"RecvLatency"`

	// 转推SRT的对端延迟，默认0，单位ms，范围为[0, 3000]。
	PeerLatency *int64 `json:"PeerLatency,omitnil" name:"PeerLatency"`

	// 转推SRT的对端空闲超时时间，默认5000，单位ms，范围为[1000, 10000]。
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitnil" name:"PeerIdleTimeout"`

	// 转推SRT的加密密钥，默认为空，表示不加密。只可填ascii码值，长度为[10, 79]。
	Passphrase *string `json:"Passphrase,omitnil" name:"Passphrase"`

	// 转推SRT的密钥长度，默认为0，可选[0|16|24|32]。
	PbKeyLen *int64 `json:"PbKeyLen,omitnil" name:"PbKeyLen"`

	// SRT模式，可选[LISTENER|CALLER]，默认为CALLER。
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

type CreateOutputSRTSettingsDestinations struct {
	// 输出的IP。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 输出的端口。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

// Predefined struct for user
type CreatePersonSampleRequestParams struct {
	// 素材名称，长度限制：20 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 素材应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于不适宜内容识别，等价于 Review.Face。
	// 3. All：包含以上全部，等价于 1+2。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 素材描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 素材图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 jpeg、png 图片格式。数组长度限制：5 张图片。
	// 注意：图片必须是单人像五官较清晰的照片，像素不低于 200*200。
	FaceContents []*string `json:"FaceContents,omitnil" name:"FaceContents"`

	// 素材标签
	// <li>数组长度限制：20 个标签；</li>
	// <li>单个标签长度限制：128 个字符。</li>
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type CreatePersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// 素材名称，长度限制：20 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 素材应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于不适宜内容识别，等价于 Review.Face。
	// 3. All：包含以上全部，等价于 1+2。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 素材描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 素材图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 jpeg、png 图片格式。数组长度限制：5 张图片。
	// 注意：图片必须是单人像五官较清晰的照片，像素不低于 200*200。
	FaceContents []*string `json:"FaceContents,omitnil" name:"FaceContents"`

	// 素材标签
	// <li>数组长度限制：20 个标签；</li>
	// <li>单个标签长度限制：128 个字符。</li>
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreatePersonSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Usages")
	delete(f, "Description")
	delete(f, "FaceContents")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonSampleResponseParams struct {
	// 素材信息。
	Person *AiSamplePerson `json:"Person,omitnil" name:"Person"`

	// 处理失败的五官定位信息。
	FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitnil" name:"FailFaceInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *CreatePersonSampleResponseParams `json:"Response"`
}

func (r *CreatePersonSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSampleSnapshotTemplateRequestParams struct {
	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 图片格式，取值为 jpg、png、webp。默认为 jpg。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

type CreateSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 图片格式，取值为 jpg、png、webp。默认为 jpg。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

func (r *CreateSampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSampleSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SampleType")
	delete(f, "SampleInterval")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Comment")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSampleSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSampleSnapshotTemplateResponseParams struct {
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateSampleSnapshotTemplateResponseParams `json:"Response"`
}

func (r *CreateSampleSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSampleSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScheduleRequestParams struct {
	// 编排名称，最多128字符。同一个用户该名称唯一。
	ScheduleName *string `json:"ScheduleName,omitnil" name:"ScheduleName"`

	// 编排绑定的触发规则，当上传视频命中该规则到该对象时即触发编排。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 编排任务列表。
	Activities []*Activity `json:"Activities,omitnil" name:"Activities"`

	// 媒体处理的文件输出存储位置。不填则继承 Trigger 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理生成的文件输出的目标目录，必选以 / 开头和结尾，如`/movie/201907/`。
	// 如果不填，表示与触发文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 任务的事件通知配置，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`
}

type CreateScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 编排名称，最多128字符。同一个用户该名称唯一。
	ScheduleName *string `json:"ScheduleName,omitnil" name:"ScheduleName"`

	// 编排绑定的触发规则，当上传视频命中该规则到该对象时即触发编排。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 编排任务列表。
	Activities []*Activity `json:"Activities,omitnil" name:"Activities"`

	// 媒体处理的文件输出存储位置。不填则继承 Trigger 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理生成的文件输出的目标目录，必选以 / 开头和结尾，如`/movie/201907/`。
	// 如果不填，表示与触发文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 任务的事件通知配置，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`
}

func (r *CreateScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleName")
	delete(f, "Trigger")
	delete(f, "Activities")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "TaskNotifyConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScheduleResponseParams struct {
	// 编排 ID。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateScheduleResponse struct {
	*tchttp.BaseResponse
	Response *CreateScheduleResponseParams `json:"Response"`
}

func (r *CreateScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotByTimeOffsetTemplateRequestParams struct {
	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 图片格式，取值可以为 jpg、png、webp。默认为 jpg。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

type CreateSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 图片格式，取值可以为 jpg、png、webp。默认为 jpg。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

func (r *CreateSnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Comment")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotByTimeOffsetTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotByTimeOffsetTemplateResponseParams struct {
	// 时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotByTimeOffsetTemplateResponseParams `json:"Response"`
}

func (r *CreateSnapshotByTimeOffsetTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkEventRequestParams struct {
	// 事件名称。
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// 事件描述。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CreateStreamLinkEventRequest struct {
	*tchttp.BaseRequest
	
	// 事件名称。
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// 事件描述。
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *CreateStreamLinkEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLinkEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkEventResponseParams struct {
	// 创建的Event信息。
	Info *DescribeEvent `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamLinkEventResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLinkEventResponseParams `json:"Response"`
}

func (r *CreateStreamLinkEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkFlowRequestParams struct {
	// 流名称。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 最大带宽，单位bps，可选[10000000, 20000000, 50000000]。
	MaxBandwidth *int64 `json:"MaxBandwidth,omitnil" name:"MaxBandwidth"`

	// 流的输入组。
	InputGroup []*CreateInput `json:"InputGroup,omitnil" name:"InputGroup"`

	// 该Flow关联的媒体传输事件ID，每个flow只能关联一个Event。
	EventId *string `json:"EventId,omitnil" name:"EventId"`
}

type CreateStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// 流名称。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 最大带宽，单位bps，可选[10000000, 20000000, 50000000]。
	MaxBandwidth *int64 `json:"MaxBandwidth,omitnil" name:"MaxBandwidth"`

	// 流的输入组。
	InputGroup []*CreateInput `json:"InputGroup,omitnil" name:"InputGroup"`

	// 该Flow关联的媒体传输事件ID，每个flow只能关联一个Event。
	EventId *string `json:"EventId,omitnil" name:"EventId"`
}

func (r *CreateStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowName")
	delete(f, "MaxBandwidth")
	delete(f, "InputGroup")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkFlowResponseParams struct {
	// 创建的Flow信息。
	Info *DescribeFlow `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLinkFlowResponseParams `json:"Response"`
}

func (r *CreateStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkInputRequestParams struct {
	// 媒体传输流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流的输入组。
	InputGroup []*CreateInput `json:"InputGroup,omitnil" name:"InputGroup"`
}

type CreateStreamLinkInputRequest struct {
	*tchttp.BaseRequest
	
	// 媒体传输流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流的输入组。
	InputGroup []*CreateInput `json:"InputGroup,omitnil" name:"InputGroup"`
}

func (r *CreateStreamLinkInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkInputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "InputGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLinkInputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkInputResponseParams struct {
	// 创建的Flow信息。
	Info *DescribeFlow `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamLinkInputResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLinkInputResponseParams `json:"Response"`
}

func (r *CreateStreamLinkInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkInputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkOutputInfoRequestParams struct {
	// 传输流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 传输流的Output配置。
	Output *CreateOutputInfo `json:"Output,omitnil" name:"Output"`
}

type CreateStreamLinkOutputInfoRequest struct {
	*tchttp.BaseRequest
	
	// 传输流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 传输流的Output配置。
	Output *CreateOutputInfo `json:"Output,omitnil" name:"Output"`
}

func (r *CreateStreamLinkOutputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkOutputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Output")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLinkOutputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkOutputInfoResponseParams struct {
	// 创建后的Output信息。
	Info *DescribeOutput `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamLinkOutputInfoResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLinkOutputInfoResponseParams `json:"Response"`
}

func (r *CreateStreamLinkOutputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkOutputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeTemplateRequestParams struct {
	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil" name:"RemoveAudio"`

	// 视频流配置参数，当 RemoveVideo 为 0，该字段必填。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil" name:"VideoTemplate"`

	// 音频流配置参数，当 RemoveAudio 为 0，该字段必填。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil" name:"AudioTemplate"`

	// 极速高清转码参数。
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil" name:"TEHDConfig"`

	// 音视频增强配置。
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil" name:"EnhanceConfig"`
}

type CreateTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil" name:"RemoveAudio"`

	// 视频流配置参数，当 RemoveVideo 为 0，该字段必填。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil" name:"VideoTemplate"`

	// 音频流配置参数，当 RemoveAudio 为 0，该字段必填。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil" name:"AudioTemplate"`

	// 极速高清转码参数。
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil" name:"TEHDConfig"`

	// 音视频增强配置。
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil" name:"EnhanceConfig"`
}

func (r *CreateTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Container")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	delete(f, "VideoTemplate")
	delete(f, "AudioTemplate")
	delete(f, "TEHDConfig")
	delete(f, "EnhanceConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeTemplateResponseParams struct {
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateTranscodeTemplateResponseParams `json:"Response"`
}

func (r *CreateTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWatermarkTemplateRequestParams struct {
	// 水印类型，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印；</li>
	// <li>svg：SVG 水印。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	// 默认值：TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil" name:"YPos"`

	// 图片水印模板，仅当 Type 为 image，该字段必填且有效。
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitnil" name:"ImageTemplate"`

	// 文字水印模板，仅当 Type 为 text，该字段必填且有效。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitnil" name:"TextTemplate"`

	// SVG 水印模板，仅当 Type 为 svg，该字段必填且有效。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitnil" name:"SvgTemplate"`
}

type CreateWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 水印类型，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印；</li>
	// <li>svg：SVG 水印。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	// 默认值：TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil" name:"YPos"`

	// 图片水印模板，仅当 Type 为 image，该字段必填且有效。
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitnil" name:"ImageTemplate"`

	// 文字水印模板，仅当 Type 为 text，该字段必填且有效。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitnil" name:"TextTemplate"`

	// SVG 水印模板，仅当 Type 为 svg，该字段必填且有效。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitnil" name:"SvgTemplate"`
}

func (r *CreateWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWatermarkTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "CoordinateOrigin")
	delete(f, "XPos")
	delete(f, "YPos")
	delete(f, "ImageTemplate")
	delete(f, "TextTemplate")
	delete(f, "SvgTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWatermarkTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWatermarkTemplateResponseParams struct {
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 水印图片地址，仅当 Type 为 image，该字段有效。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateWatermarkTemplateResponseParams `json:"Response"`
}

func (r *CreateWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWatermarkTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWordSamplesRequestParams struct {
	// <b>关键词应用场景，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜内容识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜内容识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜内容识别，等价于 3+4；
	// 7. All：通过光学字符识别技术、音频识别技术，进行内容识别、不适宜内容识别，等价于 1+2+3+4。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 关键词，数组长度限制：100。
	Words []*AiSampleWordInfo `json:"Words,omitnil" name:"Words"`
}

type CreateWordSamplesRequest struct {
	*tchttp.BaseRequest
	
	// <b>关键词应用场景，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜内容识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜内容识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜内容识别，等价于 3+4；
	// 7. All：通过光学字符识别技术、音频识别技术，进行内容识别、不适宜内容识别，等价于 1+2+3+4。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 关键词，数组长度限制：100。
	Words []*AiSampleWordInfo `json:"Words,omitnil" name:"Words"`
}

func (r *CreateWordSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWordSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Usages")
	delete(f, "Words")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWordSamplesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *CreateWordSamplesResponseParams `json:"Response"`
}

func (r *CreateWordSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWordSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowRequestParams struct {
	// 工作流名称，最多128字符。同一个用户该名称唯一。
	WorkflowName *string `json:"WorkflowName,omitnil" name:"WorkflowName"`

	// 工作流绑定的触发规则，当上传视频命中该规则到该对象时即触发工作流。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 媒体处理的文件输出存储位置。不填则继承 Trigger 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理生成的文件输出的目标目录，必选以 / 开头和结尾，如`/movie/201907/`。
	// 如果不填，表示与触发文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 媒体处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`

	// 任务的事件通知配置，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 工作流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TaskPriority *int64 `json:"TaskPriority,omitnil" name:"TaskPriority"`
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流名称，最多128字符。同一个用户该名称唯一。
	WorkflowName *string `json:"WorkflowName,omitnil" name:"WorkflowName"`

	// 工作流绑定的触发规则，当上传视频命中该规则到该对象时即触发工作流。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 媒体处理的文件输出存储位置。不填则继承 Trigger 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理生成的文件输出的目标目录，必选以 / 开头和结尾，如`/movie/201907/`。
	// 如果不填，表示与触发文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 媒体处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`

	// 任务的事件通知配置，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 工作流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TaskPriority *int64 `json:"TaskPriority,omitnil" name:"TaskPriority"`
}

func (r *CreateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowName")
	delete(f, "Trigger")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "TaskNotifyConfig")
	delete(f, "TaskPriority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowResponseParams struct {
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowResponseParams `json:"Response"`
}

func (r *CreateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIAnalysisTemplateRequestParams struct {
	// 视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIAnalysisTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIAnalysisTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIAnalysisTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAIAnalysisTemplateResponseParams `json:"Response"`
}

func (r *DeleteAIAnalysisTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIAnalysisTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIRecognitionTemplateRequestParams struct {
	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIRecognitionTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIRecognitionTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAIRecognitionTemplateResponseParams `json:"Response"`
}

func (r *DeleteAIRecognitionTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIRecognitionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAdaptiveDynamicStreamingTemplateRequestParams struct {
	// 自适应转码模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 自适应转码模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteAdaptiveDynamicStreamingTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAdaptiveDynamicStreamingTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAdaptiveDynamicStreamingTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAdaptiveDynamicStreamingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAdaptiveDynamicStreamingTemplateResponseParams `json:"Response"`
}

func (r *DeleteAdaptiveDynamicStreamingTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAdaptiveDynamicStreamingTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAnimatedGraphicsTemplateRequestParams struct {
	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAnimatedGraphicsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAnimatedGraphicsTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAnimatedGraphicsTemplateResponseParams `json:"Response"`
}

func (r *DeleteAnimatedGraphicsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContentReviewTemplateRequestParams struct {
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteContentReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContentReviewTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContentReviewTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteContentReviewTemplateResponseParams `json:"Response"`
}

func (r *DeleteContentReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContentReviewTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageSpriteTemplateRequestParams struct {
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageSpriteTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageSpriteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageSpriteTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageSpriteTemplateResponseParams `json:"Response"`
}

func (r *DeleteImageSpriteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageSpriteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonSampleRequestParams struct {
	// 素材 ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`
}

type DeletePersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// 素材 ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`
}

func (r *DeletePersonSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonSampleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeletePersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *DeletePersonSampleResponseParams `json:"Response"`
}

func (r *DeletePersonSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSampleSnapshotTemplateRequestParams struct {
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteSampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSampleSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSampleSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSampleSnapshotTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSampleSnapshotTemplateResponseParams `json:"Response"`
}

func (r *DeleteSampleSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSampleSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScheduleRequestParams struct {
	// 编排唯一标识。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`
}

type DeleteScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 编排唯一标识。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`
}

func (r *DeleteScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScheduleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScheduleResponseParams `json:"Response"`
}

func (r *DeleteScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotByTimeOffsetTemplateRequestParams struct {
	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteSnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotByTimeOffsetTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotByTimeOffsetTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotByTimeOffsetTemplateResponseParams `json:"Response"`
}

func (r *DeleteSnapshotByTimeOffsetTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkEventRequestParams struct {
	// 媒体传输事件Id，删除前需要保证该Event关联的所有Flow都已经删除。
	EventId *string `json:"EventId,omitnil" name:"EventId"`
}

type DeleteStreamLinkEventRequest struct {
	*tchttp.BaseRequest
	
	// 媒体传输事件Id，删除前需要保证该Event关联的所有Flow都已经删除。
	EventId *string `json:"EventId,omitnil" name:"EventId"`
}

func (r *DeleteStreamLinkEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLinkEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkEventResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamLinkEventResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLinkEventResponseParams `json:"Response"`
}

func (r *DeleteStreamLinkEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkFlowRequestParams struct {
	// 传输流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
}

type DeleteStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// 传输流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
}

func (r *DeleteStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLinkFlowResponseParams `json:"Response"`
}

func (r *DeleteStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkOutputRequestParams struct {
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输出Id。
	OutputId *string `json:"OutputId,omitnil" name:"OutputId"`
}

type DeleteStreamLinkOutputRequest struct {
	*tchttp.BaseRequest
	
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输出Id。
	OutputId *string `json:"OutputId,omitnil" name:"OutputId"`
}

func (r *DeleteStreamLinkOutputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkOutputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "OutputId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLinkOutputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkOutputResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamLinkOutputResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLinkOutputResponseParams `json:"Response"`
}

func (r *DeleteStreamLinkOutputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkOutputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTranscodeTemplateRequestParams struct {
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTranscodeTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTranscodeTemplateResponseParams `json:"Response"`
}

func (r *DeleteTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWatermarkTemplateRequestParams struct {
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

type DeleteWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`
}

func (r *DeleteWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWatermarkTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWatermarkTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWatermarkTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWatermarkTemplateResponseParams `json:"Response"`
}

func (r *DeleteWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWatermarkTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWordSamplesRequestParams struct {
	// 关键词，数组长度限制：100 个词。
	Keywords []*string `json:"Keywords,omitnil" name:"Keywords"`
}

type DeleteWordSamplesRequest struct {
	*tchttp.BaseRequest
	
	// 关键词，数组长度限制：100 个词。
	Keywords []*string `json:"Keywords,omitnil" name:"Keywords"`
}

func (r *DeleteWordSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWordSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keywords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWordSamplesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWordSamplesResponseParams `json:"Response"`
}

func (r *DeleteWordSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWordSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowRequestParams struct {
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`
}

type DeleteWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`
}

func (r *DeleteWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowResponseParams `json:"Response"`
}

func (r *DeleteWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAnalysisTemplatesRequestParams struct {
	// 视频内容分析模板唯一标识过滤条件，数组长度限制：10。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，不填则返回所有，可选值：
	// * Preset：系统预置模板；
	// * Custom：用户自定义模板。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeAIAnalysisTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 视频内容分析模板唯一标识过滤条件，数组长度限制：10。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，不填则返回所有，可选值：
	// * Preset：系统预置模板；
	// * Custom：用户自定义模板。
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeAIAnalysisTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAnalysisTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIAnalysisTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAnalysisTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 视频内容分析模板详情列表。
	AIAnalysisTemplateSet []*AIAnalysisTemplateItem `json:"AIAnalysisTemplateSet,omitnil" name:"AIAnalysisTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAIAnalysisTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIAnalysisTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAIAnalysisTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAnalysisTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIRecognitionTemplatesRequestParams struct {
	// 视频内容识别模板唯一标识过滤条件，数组长度限制：10。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，不填则返回所有，可选值：
	// * Preset：系统预置模板；
	// * Custom：用户自定义模板。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeAIRecognitionTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 视频内容识别模板唯一标识过滤条件，数组长度限制：10。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，不填则返回所有，可选值：
	// * Preset：系统预置模板；
	// * Custom：用户自定义模板。
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeAIRecognitionTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIRecognitionTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIRecognitionTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIRecognitionTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 视频内容识别模板详情列表。
	AIRecognitionTemplateSet []*AIRecognitionTemplateItem `json:"AIRecognitionTemplateSet,omitnil" name:"AIRecognitionTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAIRecognitionTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIRecognitionTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAIRecognitionTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIRecognitionTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdaptiveDynamicStreamingTemplatesRequestParams struct {
	// 转自适应码流模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeAdaptiveDynamicStreamingTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 转自适应码流模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeAdaptiveDynamicStreamingTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdaptiveDynamicStreamingTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdaptiveDynamicStreamingTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdaptiveDynamicStreamingTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 转自适应码流模板详情列表。
	AdaptiveDynamicStreamingTemplateSet []*AdaptiveDynamicStreamingTemplate `json:"AdaptiveDynamicStreamingTemplateSet,omitnil" name:"AdaptiveDynamicStreamingTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAdaptiveDynamicStreamingTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAdaptiveDynamicStreamingTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAdaptiveDynamicStreamingTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdaptiveDynamicStreamingTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnimatedGraphicsTemplatesRequestParams struct {
	// 转动图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeAnimatedGraphicsTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 转动图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeAnimatedGraphicsTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAnimatedGraphicsTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAnimatedGraphicsTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnimatedGraphicsTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 转动图模板详情列表。
	AnimatedGraphicsTemplateSet []*AnimatedGraphicsTemplate `json:"AnimatedGraphicsTemplateSet,omitnil" name:"AnimatedGraphicsTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAnimatedGraphicsTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAnimatedGraphicsTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAnimatedGraphicsTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAnimatedGraphicsTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentReviewTemplatesRequestParams struct {
	// 智能审核模板唯一标识过滤条件，数组长度限制：50。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，不填则返回所有，可选值：
	// * Preset：系统预置模板；
	// * Custom：用户自定义模板。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeContentReviewTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 智能审核模板唯一标识过滤条件，数组长度限制：50。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，不填则返回所有，可选值：
	// * Preset：系统预置模板；
	// * Custom：用户自定义模板。
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeContentReviewTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentReviewTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContentReviewTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentReviewTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 内容审核模板详情列表。
	ContentReviewTemplateSet []*ContentReviewTemplateItem `json:"ContentReviewTemplateSet,omitnil" name:"ContentReviewTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeContentReviewTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContentReviewTemplatesResponseParams `json:"Response"`
}

func (r *DescribeContentReviewTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentReviewTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEvent struct {
	// Event的名称。
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// Event的Id，唯一标识一个event。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// Event创建时间，格式为yyyy-MM-ddTHH:mm:ssZ。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Event的描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// Event的状态信息
	// 0：未运行
	// 1：运行中
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Event关联的Flow列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachedFlowGroup []*DescribeFlowId `json:"AttachedFlowGroup,omitnil" name:"AttachedFlowGroup"`
}

type DescribeFlow struct {
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流名称。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 流状态，目前有IDLE/RUNNING。
	State *string `json:"State,omitnil" name:"State"`

	// 最大带宽值。
	MaxBandwidth *int64 `json:"MaxBandwidth,omitnil" name:"MaxBandwidth"`

	// 输入组。
	InputGroup []*DescribeInput `json:"InputGroup,omitnil" name:"InputGroup"`

	// 输出组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputGroup []*DescribeOutput `json:"OutputGroup,omitnil" name:"OutputGroup"`

	// 该Flow关联的媒体传输事件EventId。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 媒体传输输入流所属的区域，取值和InputRegion相同。
	Region *string `json:"Region,omitnil" name:"Region"`
}

type DescribeFlowId struct {
	// FlowId，唯一标识一个flow。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// flow所在的区域名称。
	Region *string `json:"Region,omitnil" name:"Region"`
}

type DescribeHLSPullSourceAddress struct {
	// HLS源站的Url地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`
}

// Predefined struct for user
type DescribeImageSpriteTemplatesRequestParams struct {
	// 雪碧图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeImageSpriteTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 雪碧图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeImageSpriteTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSpriteTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageSpriteTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSpriteTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 雪碧图模板详情列表。
	ImageSpriteTemplateSet []*ImageSpriteTemplate `json:"ImageSpriteTemplateSet,omitnil" name:"ImageSpriteTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeImageSpriteTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageSpriteTemplatesResponseParams `json:"Response"`
}

func (r *DescribeImageSpriteTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSpriteTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInput struct {
	// 输入Id。
	InputId *string `json:"InputId,omitnil" name:"InputId"`

	// 输入名称。
	InputName *string `json:"InputName,omitnil" name:"InputName"`

	// 输入描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 输入协议。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 输入地址列表。
	InputAddressList []*InputAddress `json:"InputAddressList,omitnil" name:"InputAddressList"`

	// 输入IP白名单列表。
	AllowIpList []*string `json:"AllowIpList,omitnil" name:"AllowIpList"`

	// 输入的SRT配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SRTSettings *DescribeInputSRTSettings `json:"SRTSettings,omitnil" name:"SRTSettings"`

	// 输入的RTP配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTPSettings *DescribeInputRTPSettings `json:"RTPSettings,omitnil" name:"RTPSettings"`

	// 输入的地区。
	InputRegion *string `json:"InputRegion,omitnil" name:"InputRegion"`

	// 输入的RTMP配置信息。
	RTMPSettings *DescribeInputRTMPSettings `json:"RTMPSettings,omitnil" name:"RTMPSettings"`

	// 输入的主备开关。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailOver *string `json:"FailOver,omitnil" name:"FailOver"`

	// 输入的RTMP_PULL配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTMPPullSettings *DescribeInputRTMPPullSettings `json:"RTMPPullSettings,omitnil" name:"RTMPPullSettings"`

	// 输入的RTSP_PULL配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTSPPullSettings *DescribeInputRTSPPullSettings `json:"RTSPPullSettings,omitnil" name:"RTSPPullSettings"`

	// 输入的HLS_PULL配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HLSPullSettings *DescribeInputHLSPullSettings `json:"HLSPullSettings,omitnil" name:"HLSPullSettings"`

	// 延播平滑吐流配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResilientStream *ResilientStreamConf `json:"ResilientStream,omitnil" name:"ResilientStream"`

	// 绑定的输入安全组 ID。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type DescribeInputHLSPullSettings struct {
	// HLS源站地址信息。
	SourceAddresses []*DescribeHLSPullSourceAddress `json:"SourceAddresses,omitnil" name:"SourceAddresses"`
}

type DescribeInputRTMPPullSettings struct {
	// RTMP源站地址信息。
	SourceAddresses []*DescribeRTMPPullSourceAddress `json:"SourceAddresses,omitnil" name:"SourceAddresses"`
}

type DescribeInputRTMPSettings struct {
	// RTMP的推流路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// RTMP的推流StreamKey。
	// RTMP的推流地址拼接规则为：rtmp://Ip:1935/AppName/StreamKey
	StreamKey *string `json:"StreamKey,omitnil" name:"StreamKey"`
}

type DescribeInputRTPSettings struct {
	// 是否FEC。
	FEC *string `json:"FEC,omitnil" name:"FEC"`

	// 空闲超时时间。
	IdleTimeout *int64 `json:"IdleTimeout,omitnil" name:"IdleTimeout"`
}

type DescribeInputRTSPPullSettings struct {
	// RTSP源站地址信息。
	SourceAddresses []*DescribeRTSPPullSourceAddress `json:"SourceAddresses,omitnil" name:"SourceAddresses"`
}

type DescribeInputSRTSettings struct {
	// SRT模式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 流Id。
	StreamId *string `json:"StreamId,omitnil" name:"StreamId"`

	// 延迟。
	Latency *int64 `json:"Latency,omitnil" name:"Latency"`

	// 接收延迟。
	RecvLatency *int64 `json:"RecvLatency,omitnil" name:"RecvLatency"`

	// 对端延迟。
	PeerLatency *int64 `json:"PeerLatency,omitnil" name:"PeerLatency"`

	// 对端空闲超时时间。
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitnil" name:"PeerIdleTimeout"`

	// 解密密钥。
	Passphrase *string `json:"Passphrase,omitnil" name:"Passphrase"`

	// 密钥长度。
	PbKeyLen *int64 `json:"PbKeyLen,omitnil" name:"PbKeyLen"`

	// SRT对端地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceAddresses []*SRTSourceAddressResp `json:"SourceAddresses,omitnil" name:"SourceAddresses"`
}

// Predefined struct for user
type DescribeMediaMetaDataRequestParams struct {
	// 需要获取元信息的文件输入信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`
}

type DescribeMediaMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取元信息的文件输入信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`
}

func (r *DescribeMediaMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaMetaDataResponseParams struct {
	// 媒体元信息。
	MetaData *MediaMetaData `json:"MetaData,omitnil" name:"MetaData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMediaMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaMetaDataResponseParams `json:"Response"`
}

func (r *DescribeMediaMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOutput struct {
	// 输出Id。
	OutputId *string `json:"OutputId,omitnil" name:"OutputId"`

	// 输出名称。
	OutputName *string `json:"OutputName,omitnil" name:"OutputName"`

	// 输出类型。
	OutputType *string `json:"OutputType,omitnil" name:"OutputType"`

	// 输出描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 输出协议。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 输出的出口地址信息列表。
	OutputAddressList []*OutputAddress `json:"OutputAddressList,omitnil" name:"OutputAddressList"`

	// 输出的地区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputRegion *string `json:"OutputRegion,omitnil" name:"OutputRegion"`

	// 输出的SRT配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SRTSettings *DescribeOutputSRTSettings `json:"SRTSettings,omitnil" name:"SRTSettings"`

	// 输出的RTP配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTPSettings *DescribeOutputRTPSettings `json:"RTPSettings,omitnil" name:"RTPSettings"`

	// 输出的RTMP配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTMPSettings *DescribeOutputRTMPSettings `json:"RTMPSettings,omitnil" name:"RTMPSettings"`

	// 输出的RTMP拉流配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTMPPullSettings *DescribeOutputRTMPPullSettings `json:"RTMPPullSettings,omitnil" name:"RTMPPullSettings"`

	// CIDR白名单列表。
	// 当Protocol为RTMP_PULL有效，为空代表不限制客户端IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowIpList []*string `json:"AllowIpList,omitnil" name:"AllowIpList"`

	// 输出的RTSP拉流配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTSPPullSettings *DescribeOutputRTSPPullSettings `json:"RTSPPullSettings,omitnil" name:"RTSPPullSettings"`

	// 输出的HLS拉流配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HLSPullSettings *DescribeOutputHLSPullSettings `json:"HLSPullSettings,omitnil" name:"HLSPullSettings"`

	// 最大拉流并发数，最大为4，默认4。
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitnil" name:"MaxConcurrent"`

	// 绑定的安全组 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type DescribeOutputHLSPullServerUrl struct {
	// HLS拉流地址的Url。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type DescribeOutputHLSPullSettings struct {
	// HLS拉流地址列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerUrls []*DescribeOutputHLSPullServerUrl `json:"ServerUrls,omitnil" name:"ServerUrls"`
}

type DescribeOutputRTMPPullServerUrl struct {
	// RTMP拉流地址的tcUrl。
	TcUrl *string `json:"TcUrl,omitnil" name:"TcUrl"`

	// RTMP拉流地址的流key。
	StreamKey *string `json:"StreamKey,omitnil" name:"StreamKey"`
}

type DescribeOutputRTMPPullSettings struct {
	// 拉流地址列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerUrls []*DescribeOutputRTMPPullServerUrl `json:"ServerUrls,omitnil" name:"ServerUrls"`
}

type DescribeOutputRTMPSettings struct {
	// 空闲超时时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdleTimeout *int64 `json:"IdleTimeout,omitnil" name:"IdleTimeout"`

	// Chunk大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChunkSize *int64 `json:"ChunkSize,omitnil" name:"ChunkSize"`

	// 转推RTMP的目标地址信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destinations []*RTMPAddressDestination `json:"Destinations,omitnil" name:"Destinations"`
}

type DescribeOutputRTPSettings struct {
	// 转推RTP的目标地址信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destinations []*RTPAddressDestination `json:"Destinations,omitnil" name:"Destinations"`

	// 是否FEC。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FEC *string `json:"FEC,omitnil" name:"FEC"`

	// 空闲超时时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdleTimeout *int64 `json:"IdleTimeout,omitnil" name:"IdleTimeout"`
}

type DescribeOutputRTSPPullServerUrl struct {
	// RTSP拉流地址的Url。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type DescribeOutputRTSPPullSettings struct {
	// RTSP拉流地址列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerUrls []*DescribeOutputRTSPPullServerUrl `json:"ServerUrls,omitnil" name:"ServerUrls"`
}

type DescribeOutputSRTSettings struct {
	// 转推的目标的地址信息列表，SRT模式为CALLER时使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destinations []*SRTAddressDestination `json:"Destinations,omitnil" name:"Destinations"`

	// 流Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamId *string `json:"StreamId,omitnil" name:"StreamId"`

	// 延迟。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Latency *int64 `json:"Latency,omitnil" name:"Latency"`

	// 接收延迟。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecvLatency *int64 `json:"RecvLatency,omitnil" name:"RecvLatency"`

	// 对端延迟。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerLatency *int64 `json:"PeerLatency,omitnil" name:"PeerLatency"`

	// 对端空闲超时时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitnil" name:"PeerIdleTimeout"`

	// 加密密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Passphrase *string `json:"Passphrase,omitnil" name:"Passphrase"`

	// 加密密钥长度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PbKeyLen *int64 `json:"PbKeyLen,omitnil" name:"PbKeyLen"`

	// SRT模式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 服务器监听地址，SRT模式为LISTENER时使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceAddresses []*OutputSRTSourceAddressResp `json:"SourceAddresses,omitnil" name:"SourceAddresses"`
}

// Predefined struct for user
type DescribePersonSamplesRequestParams struct {
	// 拉取的素材类型，可选值：
	// <li>UserDefine：用户自定义素材库；</li>
	// <li>Default：系统默认素材库。</li>
	// 
	// 默认值：UserDefine，拉取用户自定义素材库素材。
	// 说明：如果是拉取系统默认素材库，只能使用素材名字或者素材 ID + 素材名字的方式进行拉取，且人脸图片只返回一张。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 素材 ID，数组长度限制：100。
	PersonIds []*string `json:"PersonIds,omitnil" name:"PersonIds"`

	// 素材名称，数组长度限制：20。
	Names []*string `json:"Names,omitnil" name:"Names"`

	// 素材标签，数组长度限制：20。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribePersonSamplesRequest struct {
	*tchttp.BaseRequest
	
	// 拉取的素材类型，可选值：
	// <li>UserDefine：用户自定义素材库；</li>
	// <li>Default：系统默认素材库。</li>
	// 
	// 默认值：UserDefine，拉取用户自定义素材库素材。
	// 说明：如果是拉取系统默认素材库，只能使用素材名字或者素材 ID + 素材名字的方式进行拉取，且人脸图片只返回一张。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 素材 ID，数组长度限制：100。
	PersonIds []*string `json:"PersonIds,omitnil" name:"PersonIds"`

	// 素材名称，数组长度限制：20。
	Names []*string `json:"Names,omitnil" name:"Names"`

	// 素材标签，数组长度限制：20。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribePersonSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "PersonIds")
	delete(f, "Names")
	delete(f, "Tags")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonSamplesResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 素材信息。
	PersonSet []*AiSamplePerson `json:"PersonSet,omitnil" name:"PersonSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePersonSamplesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonSamplesResponseParams `json:"Response"`
}

func (r *DescribePersonSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRTMPPullSourceAddress struct {
	// RTMP源站的TcUrl地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcUrl *string `json:"TcUrl,omitnil" name:"TcUrl"`

	// RTMP源站的StreamKey。
	// RTMP源站地址拼接规则为：$TcUrl/$StreamKey。
	StreamKey *string `json:"StreamKey,omitnil" name:"StreamKey"`
}

type DescribeRTSPPullSourceAddress struct {
	// RTSP源站的Url地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`
}

// Predefined struct for user
type DescribeSampleSnapshotTemplatesRequestParams struct {
	// 采样截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeSampleSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 采样截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeSampleSnapshotTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSampleSnapshotTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSampleSnapshotTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSampleSnapshotTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 采样截图模板详情列表。
	SampleSnapshotTemplateSet []*SampleSnapshotTemplate `json:"SampleSnapshotTemplateSet,omitnil" name:"SampleSnapshotTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSampleSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSampleSnapshotTemplatesResponseParams `json:"Response"`
}

func (r *DescribeSampleSnapshotTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSampleSnapshotTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulesRequestParams struct {
	// 编排 ID 过滤条件，数组长度限制：100。
	ScheduleIds []*int64 `json:"ScheduleIds,omitnil" name:"ScheduleIds"`

	// 编排触发类型，可选值：
	// <li>CosFileUpload： 腾讯云 COS 文件上传触发</li>
	// <li>AwsS3FileUpload：Aws S3 文件上传触发。</li>
	// 不填或者为空表示全部。
	TriggerType *string `json:"TriggerType,omitnil" name:"TriggerType"`

	// 状态，取值范围：
	// <li>Enabled：已启用，</li>
	// <li>Disabled：已禁用。</li>
	// 不填此参数，则不区编排状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSchedulesRequest struct {
	*tchttp.BaseRequest
	
	// 编排 ID 过滤条件，数组长度限制：100。
	ScheduleIds []*int64 `json:"ScheduleIds,omitnil" name:"ScheduleIds"`

	// 编排触发类型，可选值：
	// <li>CosFileUpload： 腾讯云 COS 文件上传触发</li>
	// <li>AwsS3FileUpload：Aws S3 文件上传触发。</li>
	// 不填或者为空表示全部。
	TriggerType *string `json:"TriggerType,omitnil" name:"TriggerType"`

	// 状态，取值范围：
	// <li>Enabled：已启用，</li>
	// <li>Disabled：已禁用。</li>
	// 不填此参数，则不区编排状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSchedulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleIds")
	delete(f, "TriggerType")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSchedulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 编排信息数组。
	ScheduleInfoSet []*SchedulesInfo `json:"ScheduleInfoSet,omitnil" name:"ScheduleInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSchedulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSchedulesResponseParams `json:"Response"`
}

func (r *DescribeSchedulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotByTimeOffsetTemplatesRequestParams struct {
	// 指定时间点截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeSnapshotByTimeOffsetTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 指定时间点截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeSnapshotByTimeOffsetTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotByTimeOffsetTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotByTimeOffsetTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotByTimeOffsetTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 指定时间点截图模板详情列表。
	SnapshotByTimeOffsetTemplateSet []*SnapshotByTimeOffsetTemplate `json:"SnapshotByTimeOffsetTemplateSet,omitnil" name:"SnapshotByTimeOffsetTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSnapshotByTimeOffsetTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotByTimeOffsetTemplatesResponseParams `json:"Response"`
}

func (r *DescribeSnapshotByTimeOffsetTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotByTimeOffsetTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkActivateStateRequestParams struct {

}

type DescribeStreamLinkActivateStateRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamLinkActivateStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkActivateStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkActivateStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkActivateStateResponseParams struct {
	// 用户已激活为0，否则为非0。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkActivateStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkActivateStateResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkActivateStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkActivateStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkEventAttachedFlowsRequestParams struct {
	// EventId。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 当前页数，默认1。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页大小，默认1000。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeStreamLinkEventAttachedFlowsRequest struct {
	*tchttp.BaseRequest
	
	// EventId。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 当前页数，默认1。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页大小，默认1000。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeStreamLinkEventAttachedFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkEventAttachedFlowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkEventAttachedFlowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkEventAttachedFlowsResponseParams struct {
	// 流的配置信息列表。
	Infos []*DescribeFlow `json:"Infos,omitnil" name:"Infos"`

	// 总数量。
	TotalNum *int64 `json:"TotalNum,omitnil" name:"TotalNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkEventAttachedFlowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkEventAttachedFlowsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkEventAttachedFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkEventAttachedFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkEventRequestParams struct {
	// 媒体传输事件ID。
	EventId *string `json:"EventId,omitnil" name:"EventId"`
}

type DescribeStreamLinkEventRequest struct {
	*tchttp.BaseRequest
	
	// 媒体传输事件ID。
	EventId *string `json:"EventId,omitnil" name:"EventId"`
}

func (r *DescribeStreamLinkEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkEventResponseParams struct {
	// 媒体传输事件的配置信息。
	Info *DescribeEvent `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkEventResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkEventsRequestParams struct {
	// 当前页数，默认1。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页大小，默认10。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeStreamLinkEventsRequest struct {
	*tchttp.BaseRequest
	
	// 当前页数，默认1。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页大小，默认10。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeStreamLinkEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkEventsResponseParams struct {
	// 媒体传输事件的配置信息列表。
	Infos []*DescribeEvent `json:"Infos,omitnil" name:"Infos"`

	// 当前页数。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页大小。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 总数量。
	TotalNum *int64 `json:"TotalNum,omitnil" name:"TotalNum"`

	// 总页数。
	TotalPage *int64 `json:"TotalPage,omitnil" name:"TotalPage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkEventsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowLogsRequestParams struct {
	// 传输流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 统计的开始时间，默认为前一小时，最多支持查询近7天。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 统计的结束时间，默认为StartTime后一小时，最多支持查询24小时的数据。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 输入或输出类型，可选[input|output]。
	Type []*string `json:"Type,omitnil" name:"Type"`

	// 主通道或备通道，可选[0|1]。
	Pipeline []*string `json:"Pipeline,omitnil" name:"Pipeline"`

	// 每页大小，默认100，范围为[1, 1000]。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 按Timestamp升序或降序排序，默认降序，可选[desc|asc]。
	SortType *string `json:"SortType,omitnil" name:"SortType"`

	// 页码，默认1，范围为[1, 1000]。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`
}

type DescribeStreamLinkFlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 传输流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 统计的开始时间，默认为前一小时，最多支持查询近7天。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 统计的结束时间，默认为StartTime后一小时，最多支持查询24小时的数据。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 输入或输出类型，可选[input|output]。
	Type []*string `json:"Type,omitnil" name:"Type"`

	// 主通道或备通道，可选[0|1]。
	Pipeline []*string `json:"Pipeline,omitnil" name:"Pipeline"`

	// 每页大小，默认100，范围为[1, 1000]。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 按Timestamp升序或降序排序，默认降序，可选[desc|asc]。
	SortType *string `json:"SortType,omitnil" name:"SortType"`

	// 页码，默认1，范围为[1, 1000]。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`
}

func (r *DescribeStreamLinkFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Pipeline")
	delete(f, "PageSize")
	delete(f, "SortType")
	delete(f, "PageNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowLogsResponseParams struct {
	// 日志信息列表。
	Infos []*FlowLogInfo `json:"Infos,omitnil" name:"Infos"`

	// 当前页码。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页大小。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 总数量。
	TotalNum *int64 `json:"TotalNum,omitnil" name:"TotalNum"`

	// 总页数。
	TotalPage *int64 `json:"TotalPage,omitnil" name:"TotalPage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowLogsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowMediaStatisticsRequestParams struct {
	// 传输流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输入或输出类型，可选[input|output]。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入或输出Id。
	InputOutputId *string `json:"InputOutputId,omitnil" name:"InputOutputId"`

	// 主通道或备通道，可选[0|1]。
	Pipeline *string `json:"Pipeline,omitnil" name:"Pipeline"`

	// 查询间隔，可选[5s|1min|5min|15min]。
	Period *string `json:"Period,omitnil" name:"Period"`

	// 统计的开始时间，默认为前一小时，最多支持查询近7天。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 统计的结束时间，默认为StartTime后一小时，最多支持查询24小时的数据。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeStreamLinkFlowMediaStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 传输流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输入或输出类型，可选[input|output]。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入或输出Id。
	InputOutputId *string `json:"InputOutputId,omitnil" name:"InputOutputId"`

	// 主通道或备通道，可选[0|1]。
	Pipeline *string `json:"Pipeline,omitnil" name:"Pipeline"`

	// 查询间隔，可选[5s|1min|5min|15min]。
	Period *string `json:"Period,omitnil" name:"Period"`

	// 统计的开始时间，默认为前一小时，最多支持查询近7天。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 统计的结束时间，默认为StartTime后一小时，最多支持查询24小时的数据。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeStreamLinkFlowMediaStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowMediaStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Type")
	delete(f, "InputOutputId")
	delete(f, "Pipeline")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowMediaStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowMediaStatisticsResponseParams struct {
	// 传输流的媒体数据列表。
	Infos []*FlowMediaInfo `json:"Infos,omitnil" name:"Infos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkFlowMediaStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowMediaStatisticsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowMediaStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowMediaStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowRealtimeStatusRequestParams struct {
	// 流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输入id数组，如果输入输出数组都为空，则代表全量查询。
	InputIds []*string `json:"InputIds,omitnil" name:"InputIds"`

	// 输出id数组，如果输入输出数组都为空，则代表全量查询。
	OutputIds []*string `json:"OutputIds,omitnil" name:"OutputIds"`
}

type DescribeStreamLinkFlowRealtimeStatusRequest struct {
	*tchttp.BaseRequest
	
	// 流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输入id数组，如果输入输出数组都为空，则代表全量查询。
	InputIds []*string `json:"InputIds,omitnil" name:"InputIds"`

	// 输出id数组，如果输入输出数组都为空，则代表全量查询。
	OutputIds []*string `json:"OutputIds,omitnil" name:"OutputIds"`
}

func (r *DescribeStreamLinkFlowRealtimeStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowRealtimeStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "InputIds")
	delete(f, "OutputIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowRealtimeStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowRealtimeStatusResponseParams struct {
	// 查询时间，单位s。
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 实时数据信息列表。
	Datas []*FlowRealtimeStatusItem `json:"Datas,omitnil" name:"Datas"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkFlowRealtimeStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowRealtimeStatusResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowRealtimeStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowRealtimeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowRequestParams struct {
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
}

type DescribeStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
}

func (r *DescribeStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowResponseParams struct {
	// 流的配置信息。
	Info *DescribeFlow `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowSRTStatisticsRequestParams struct {
	// 传输流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输入或输出类型，可选[input|output]。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入或输出Id。
	InputOutputId *string `json:"InputOutputId,omitnil" name:"InputOutputId"`

	// 主通道或备通道，可选[0|1]。
	Pipeline *string `json:"Pipeline,omitnil" name:"Pipeline"`

	// 统计的开始时间，默认为前一小时，最多支持查询近7天。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 统计的结束时间，默认为StartTime后一小时，最多支持查询24小时的数据。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询间隔，可选[5s|1min|5min|15min]。
	Period *string `json:"Period,omitnil" name:"Period"`
}

type DescribeStreamLinkFlowSRTStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 传输流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输入或输出类型，可选[input|output]。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入或输出Id。
	InputOutputId *string `json:"InputOutputId,omitnil" name:"InputOutputId"`

	// 主通道或备通道，可选[0|1]。
	Pipeline *string `json:"Pipeline,omitnil" name:"Pipeline"`

	// 统计的开始时间，默认为前一小时，最多支持查询近7天。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 统计的结束时间，默认为StartTime后一小时，最多支持查询24小时的数据。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询间隔，可选[5s|1min|5min|15min]。
	Period *string `json:"Period,omitnil" name:"Period"`
}

func (r *DescribeStreamLinkFlowSRTStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowSRTStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Type")
	delete(f, "InputOutputId")
	delete(f, "Pipeline")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowSRTStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowSRTStatisticsResponseParams struct {
	// 传输流的SRT质量数据列表。
	Infos []*FlowSRTInfo `json:"Infos,omitnil" name:"Infos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkFlowSRTStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowSRTStatisticsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowSRTStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowSRTStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowStatisticsRequestParams struct {
	// 传输流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输入或输出类型，可选[input|output]。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入或输出Id。
	InputOutputId *string `json:"InputOutputId,omitnil" name:"InputOutputId"`

	// 主通道或备通道，可选[0|1]。
	Pipeline *string `json:"Pipeline,omitnil" name:"Pipeline"`

	// 查询间隔，可选[5s|1min|5min|15min]。
	Period *string `json:"Period,omitnil" name:"Period"`

	// 统计的开始时间，默认为前一小时，最多支持查询近7天。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 统计的结束时间，默认为StartTime后一小时，最多支持查询24小时的数据。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeStreamLinkFlowStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 传输流ID。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 输入或输出类型，可选[input|output]。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入或输出Id。
	InputOutputId *string `json:"InputOutputId,omitnil" name:"InputOutputId"`

	// 主通道或备通道，可选[0|1]。
	Pipeline *string `json:"Pipeline,omitnil" name:"Pipeline"`

	// 查询间隔，可选[5s|1min|5min|15min]。
	Period *string `json:"Period,omitnil" name:"Period"`

	// 统计的开始时间，默认为前一小时，最多支持查询近7天。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 统计的结束时间，默认为StartTime后一小时，最多支持查询24小时的数据。
	// UTC时间，如'2020-01-01T12:00:00Z'。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeStreamLinkFlowStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Type")
	delete(f, "InputOutputId")
	delete(f, "Pipeline")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowStatisticsResponseParams struct {
	// 传输流的媒体数据列表。
	Infos []*FlowStatisticsArray `json:"Infos,omitnil" name:"Infos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkFlowStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowStatisticsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowsRequestParams struct {
	// 当前页数，默认1。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页大小，默认10。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeStreamLinkFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 当前页数，默认1。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页大小，默认10。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeStreamLinkFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowsResponseParams struct {
	// 流的配置信息列表。
	Infos []*DescribeFlow `json:"Infos,omitnil" name:"Infos"`

	// 当前页数。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页大小。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 总数量。
	TotalNum *int64 `json:"TotalNum,omitnil" name:"TotalNum"`

	// 总页数。
	TotalPage *int64 `json:"TotalPage,omitnil" name:"TotalPage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkFlowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkRegionsRequestParams struct {

}

type DescribeStreamLinkRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamLinkRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkRegionsResponseParams struct {
	// 媒体传输地区信息。
	Info *StreamLinkRegionInfo `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLinkRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkRegionsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// 视频处理任务的任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 视频处理任务的任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
	// 任务类型，目前取值有：
	// <li>WorkflowTask：视频工作流处理任务。</li>
	// <li>EditMediaTask：视频编辑任务。</li>
	// <li>LiveStreamProcessTask：直播流处理任务。</li>
	// <li>ScheduleTask：编排处理任务。</li>
	TaskType *string `json:"TaskType,omitnil" name:"TaskType"`

	// 任务状态，取值：
	// <li>WAITING：等待中；</li>
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务的创建时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`

	// 视频编辑任务信息，仅当 TaskType 为 EditMediaTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditMediaTask *EditMediaTask `json:"EditMediaTask,omitnil" name:"EditMediaTask"`

	// 视频处理任务信息，仅当 TaskType 为 WorkflowTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowTask *WorkflowTask `json:"WorkflowTask,omitnil" name:"WorkflowTask"`

	// 直播流处理任务信息，仅当 TaskType 为 LiveStreamProcessTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveStreamProcessTask *LiveStreamProcessTask `json:"LiveStreamProcessTask,omitnil" name:"LiveStreamProcessTask"`

	// 任务的事件通知信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 任务流的优先级，取值范围为 [-10, 10]。
	TasksPriority *int64 `json:"TasksPriority,omitnil" name:"TasksPriority"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长50个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长1000个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 扩展信息字段，仅用于特定场景。
	ExtInfo *string `json:"ExtInfo,omitnil" name:"ExtInfo"`

	// 编排处理任务信息，仅当 TaskType 为 ScheduleTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleTask *ScheduleTask `json:"ScheduleTask,omitnil" name:"ScheduleTask"`

	// 直播编排处理任务信息，仅当 TaskType 为 LiveScheduleTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveScheduleTask *LiveScheduleTask `json:"LiveScheduleTask,omitnil" name:"LiveScheduleTask"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// 过滤条件：任务状态，可选值：WAITING（等待中）、PROCESSING（处理中）、FINISH（已完成）。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil" name:"ScrollToken"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件：任务状态，可选值：WAITING（等待中）、PROCESSING（处理中）、FINISH（已完成）。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil" name:"ScrollToken"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Limit")
	delete(f, "ScrollToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 任务概要列表。
	TaskSet []*TaskSimpleInfo `json:"TaskSet,omitnil" name:"TaskSet"`

	// 翻页标识，当请求未返回所有数据，该字段表示下一条记录的 ID。当该字段为空字符串，说明已无更多数据。
	ScrollToken *string `json:"ScrollToken,omitnil" name:"ScrollToken"`

	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeTemplatesRequestParams struct {
	// 转码模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式板；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式。</li>
	ContainerType *string `json:"ContainerType,omitnil" name:"ContainerType"`

	// （建议使用TranscodeType代替）极速高清过滤条件，用于过滤普通转码或极速高清转码模板，可选值：
	// <li>Common：普通转码模板；</li>
	// <li>TEHD：极速高清模板。</li>
	TEHDType *string `json:"TEHDType,omitnil" name:"TEHDType"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型（替换旧版本 TEHDType），可选值：
	// <li>Common：普通转码模板；</li>
	// <li>TEHD：视频极速高清，老的类型（建议使用 TEHD-100） 。</li>
	// <li>TEHD-100：视频极速高清</li>
	// <li>TEHD-200：音频极速高清</li>
	// <li>Enhance：音视频增强模板。</li>
	// 默认空，不限制类型。
	TranscodeType *string `json:"TranscodeType,omitnil" name:"TranscodeType"`
}

type DescribeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 转码模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式板；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式。</li>
	ContainerType *string `json:"ContainerType,omitnil" name:"ContainerType"`

	// （建议使用TranscodeType代替）极速高清过滤条件，用于过滤普通转码或极速高清转码模板，可选值：
	// <li>Common：普通转码模板；</li>
	// <li>TEHD：极速高清模板。</li>
	TEHDType *string `json:"TEHDType,omitnil" name:"TEHDType"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模板类型（替换旧版本 TEHDType），可选值：
	// <li>Common：普通转码模板；</li>
	// <li>TEHD：视频极速高清，老的类型（建议使用 TEHD-100） 。</li>
	// <li>TEHD-100：视频极速高清</li>
	// <li>TEHD-200：音频极速高清</li>
	// <li>Enhance：音视频增强模板。</li>
	// 默认空，不限制类型。
	TranscodeType *string `json:"TranscodeType,omitnil" name:"TranscodeType"`
}

func (r *DescribeTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "ContainerType")
	delete(f, "TEHDType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TranscodeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 转码模板详情列表。
	TranscodeTemplateSet []*TranscodeTemplate `json:"TranscodeTemplateSet,omitnil" name:"TranscodeTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTranscodeTemplatesResponseParams `json:"Response"`
}

func (r *DescribeTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWatermarkTemplatesRequestParams struct {
	// 水印模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 水印类型过滤条件，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数
	// <li>默认值：10；</li>
	// <li>最大值：100。</li>
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeWatermarkTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 水印模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil" name:"Definitions"`

	// 水印类型过滤条件，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数
	// <li>默认值：10；</li>
	// <li>最大值：100。</li>
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeWatermarkTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWatermarkTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWatermarkTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWatermarkTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 水印模板详情列表。
	WatermarkTemplateSet []*WatermarkTemplate `json:"WatermarkTemplateSet,omitnil" name:"WatermarkTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWatermarkTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWatermarkTemplatesResponseParams `json:"Response"`
}

func (r *DescribeWatermarkTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWatermarkTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWordSamplesRequestParams struct {
	// 关键词过滤条件，数组长度限制：100 个词。
	Keywords []*string `json:"Keywords,omitnil" name:"Keywords"`

	// <b>关键词应用场景过滤条件，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜内容的识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜内容的识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜内容的识别，等价于 3+4；
	// 可多选，元素间关系为 or，即关键词的应用场景包含该字段集合中任意元素的记录，均符合该条件。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 标签过滤条件，数组长度限制：20 个词。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeWordSamplesRequest struct {
	*tchttp.BaseRequest
	
	// 关键词过滤条件，数组长度限制：100 个词。
	Keywords []*string `json:"Keywords,omitnil" name:"Keywords"`

	// <b>关键词应用场景过滤条件，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜内容的识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜内容的识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜内容的识别，等价于 3+4；
	// 可多选，元素间关系为 or，即关键词的应用场景包含该字段集合中任意元素的记录，均符合该条件。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 标签过滤条件，数组长度限制：20 个词。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeWordSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWordSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keywords")
	delete(f, "Usages")
	delete(f, "Tags")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWordSamplesResponseParams struct {
	// 符合条件的记录总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 关键词信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordSet []*AiSampleWord `json:"WordSet,omitnil" name:"WordSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWordSamplesResponseParams `json:"Response"`
}

func (r *DescribeWordSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWordSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowsRequestParams struct {
	// 工作流 ID 过滤条件，数组长度限制：100。
	WorkflowIds []*int64 `json:"WorkflowIds,omitnil" name:"WorkflowIds"`

	// 工作流状态，取值范围：
	// <li>Enabled：已启用，</li>
	// <li>Disabled：已禁用。</li>
	// 不填此参数，则不区分工作流状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeWorkflowsRequest struct {
	*tchttp.BaseRequest
	
	// 工作流 ID 过滤条件，数组长度限制：100。
	WorkflowIds []*int64 `json:"WorkflowIds,omitnil" name:"WorkflowIds"`

	// 工作流状态，取值范围：
	// <li>Enabled：已启用，</li>
	// <li>Disabled：已禁用。</li>
	// 不填此参数，则不区分工作流状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowIds")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowsResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 工作流信息数组。
	WorkflowInfoSet []*WorkflowInfo `json:"WorkflowInfoSet,omitnil" name:"WorkflowInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowsResponseParams `json:"Response"`
}

func (r *DescribeWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiagnoseResult struct {
	// 诊断出的异常类别。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitnil" name:"Category"`

	// 诊断出的具体异常类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 诊断出异常开始的PTS时间戳。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *float64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 诊断出的异常描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 诊断到异常的北京时间，采用 ISO 日期格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DateTime *string `json:"DateTime,omitnil" name:"DateTime"`

	// 诊断出的异常级别。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SeverityLevel *string `json:"SeverityLevel,omitnil" name:"SeverityLevel"`
}

// Predefined struct for user
type DisableScheduleRequestParams struct {
	// 编排唯一表示。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`
}

type DisableScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 编排唯一表示。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`
}

func (r *DisableScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableScheduleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DisableScheduleResponseParams `json:"Response"`
}

func (r *DisableScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWorkflowRequestParams struct {
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`
}

type DisableWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`
}

func (r *DisableWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWorkflowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *DisableWorkflowResponseParams `json:"Response"`
}

func (r *DisableWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrmInfo struct {
	// 加密类型：
	// <li> simpleaes: aes-128 加密</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// SimpleAes 加密信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimpleAesDrm *SimpleAesDrm `json:"SimpleAesDrm,omitnil" name:"SimpleAesDrm"`
}

type EditMediaFileInfo struct {
	// 视频的输入信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 【剪辑】任务生效，视频剪辑的起始时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 【剪辑】任务生效，视频剪辑的结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 【合成】任务必选，用于轨道元素中媒体关联源素材 ID。
	// 
	// 注意：允许字母、数字、-、_ ，最长 32 字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`
}

type EditMediaOutputConfig struct {
	// 封装格式，可选值：mp4、hls、mov、flv、avi。默认是 mp4。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 剪辑模式，可选值 normal、fast。默认是精确剪辑 normal
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

// Predefined struct for user
type EditMediaRequestParams struct {
	// 输入的视频文件信息。
	FileInfos []*EditMediaFileInfo `json:"FileInfos,omitnil" name:"FileInfos"`

	// 媒体处理输出文件的目标存储。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理输出文件的目标路径。
	// 
	// 注意：对于复杂合成任务，路径中的文件名只可为数字、字母、-、_ 的组合，最长 64 个字符。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil" name:"OutputObjectPath"`

	// 【剪辑】任务生成的文件配置。
	OutputConfig *EditMediaOutputConfig `json:"OutputConfig,omitnil" name:"OutputConfig"`

	// 【合成】任务配置。
	// 
	// 注意：当其不为空时，认为是合成任务，否则按剪辑任务处理。
	ComposeConfig *ComposeMediaConfig `json:"ComposeConfig,omitnil" name:"ComposeConfig"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 任务优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil" name:"TasksPriority"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`
}

type EditMediaRequest struct {
	*tchttp.BaseRequest
	
	// 输入的视频文件信息。
	FileInfos []*EditMediaFileInfo `json:"FileInfos,omitnil" name:"FileInfos"`

	// 媒体处理输出文件的目标存储。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理输出文件的目标路径。
	// 
	// 注意：对于复杂合成任务，路径中的文件名只可为数字、字母、-、_ 的组合，最长 64 个字符。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil" name:"OutputObjectPath"`

	// 【剪辑】任务生成的文件配置。
	OutputConfig *EditMediaOutputConfig `json:"OutputConfig,omitnil" name:"OutputConfig"`

	// 【合成】任务配置。
	// 
	// 注意：当其不为空时，认为是合成任务，否则按剪辑任务处理。
	ComposeConfig *ComposeMediaConfig `json:"ComposeConfig,omitnil" name:"ComposeConfig"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 任务优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil" name:"TasksPriority"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`
}

func (r *EditMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileInfos")
	delete(f, "OutputStorage")
	delete(f, "OutputObjectPath")
	delete(f, "OutputConfig")
	delete(f, "ComposeConfig")
	delete(f, "TaskNotifyConfig")
	delete(f, "TasksPriority")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditMediaResponseParams struct {
	// 编辑视频的任务 ID，可以通过该 ID 查询编辑任务的状态。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EditMediaResponse struct {
	*tchttp.BaseResponse
	Response *EditMediaResponseParams `json:"Response"`
}

func (r *EditMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditMediaTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 视频编辑任务的输入。
	Input *EditMediaTaskInput `json:"Input,omitnil" name:"Input"`

	// 视频编辑任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *EditMediaTaskOutput `json:"Output,omitnil" name:"Output"`
}

type EditMediaTaskInput struct {
	// 输入的视频文件信息。
	FileInfoSet []*EditMediaFileInfo `json:"FileInfoSet,omitnil" name:"FileInfoSet"`
}

type EditMediaTaskOutput struct {
	// 编辑后文件的目标存储。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 编辑后的视频文件路径。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 编辑后的视频文件元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil" name:"MetaData"`
}

// Predefined struct for user
type EnableScheduleRequestParams struct {
	// 编排唯一标识。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`
}

type EnableScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 编排唯一标识。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`
}

func (r *EnableScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableScheduleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableScheduleResponse struct {
	*tchttp.BaseResponse
	Response *EnableScheduleResponseParams `json:"Response"`
}

func (r *EnableScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWorkflowRequestParams struct {
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`
}

type EnableWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`
}

func (r *EnableWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWorkflowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *EnableWorkflowResponseParams `json:"Response"`
}

func (r *EnableWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhanceConfig struct {
	// 视频增强配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoEnhance *VideoEnhanceConfig `json:"VideoEnhance,omitnil" name:"VideoEnhance"`

	// 音频增强配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioEnhance *AudioEnhanceConfig `json:"AudioEnhance,omitnil" name:"AudioEnhance"`
}

// Predefined struct for user
type ExecuteFunctionRequestParams struct {
	// 调用后端接口名称。
	FunctionName *string `json:"FunctionName,omitnil" name:"FunctionName"`

	// 接口参数，具体参数格式调用时与后端协调。
	FunctionArg *string `json:"FunctionArg,omitnil" name:"FunctionArg"`
}

type ExecuteFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 调用后端接口名称。
	FunctionName *string `json:"FunctionName,omitnil" name:"FunctionName"`

	// 接口参数，具体参数格式调用时与后端协调。
	FunctionArg *string `json:"FunctionArg,omitnil" name:"FunctionArg"`
}

func (r *ExecuteFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionName")
	delete(f, "FunctionArg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteFunctionResponseParams struct {
	// 处理结果打包后的字符串，具体与后台一同协调。
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExecuteFunctionResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteFunctionResponseParams `json:"Response"`
}

func (r *ExecuteFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpressionConfigInfo struct {
	// 表情识别任务开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type FaceConfigureInfo struct {
	// 人脸识别任务开关，可选值：
	// <li>ON：开启智能人脸识别任务；</li>
	// <li>OFF：关闭智能人脸识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 人脸识别过滤分数，当识别结果达到该分数以上，返回识别结果。默认 95 分。取值范围：0 - 100。
	Score *float64 `json:"Score,omitnil" name:"Score"`

	// 默认人物过滤标签，指定需要返回的默认人物的标签。如果未填或者为空，则全部默认人物结果都返回。标签可选值：
	// <li>entertainment：娱乐明星；</li>
	// <li>sport：体育明星；</li>
	// <li>politician：敏感人物。</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitnil" name:"DefaultLibraryLabelSet"`

	// 用户自定义人物过滤标签，指定需要返回的用户自定义人物的标签。如果未填或者为空，则全部自定义人物结果都返回。
	// 标签个数最多 100 个，每个标签长度最多 16 个字符。
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitnil" name:"UserDefineLibraryLabelSet"`

	// 人物库选择，可选值：
	// <li>Default：使用默认人物库；</li>
	// <li>UserDefine：使用用户自定义人物库。</li>
	// <li>All：同时使用默认人物库和用户自定义人物库。</li>
	// 默认值：All，使用系统默认人物库及用户自定义人物库。
	FaceLibrary *string `json:"FaceLibrary,omitnil" name:"FaceLibrary"`
}

type FaceConfigureInfoForUpdate struct {
	// 人脸识别任务开关，可选值：
	// <li>ON：开启智能人脸识别任务；</li>
	// <li>OFF：关闭智能人脸识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 人脸识别过滤分数，当识别结果达到该分数以上，返回识别结果。取值范围：0-100。
	Score *float64 `json:"Score,omitnil" name:"Score"`

	// 默认人物过滤标签，指定需要返回的默认人物的标签。如果未填或者为空，则全部默认人物结果都返回。标签可选值：
	// <li>entertainment：娱乐明星；</li>
	// <li>sport：体育明星；</li>
	// <li>politician：敏感人物。</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitnil" name:"DefaultLibraryLabelSet"`

	// 用户自定义人物过滤标签，指定需要返回的用户自定义人物的标签。如果未填或者为空，则全部自定义人物结果都返回。
	// 标签个数最多 100 个，每个标签长度最多 16 个字符。
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitnil" name:"UserDefineLibraryLabelSet"`

	// 人物库选择，可选值：
	// <li>Default：使用默认人物库；</li>
	// <li>UserDefine：使用用户自定义人物库。</li>
	// <li>All：同时使用默认人物库和用户自定义人物库。</li>
	FaceLibrary *string `json:"FaceLibrary,omitnil" name:"FaceLibrary"`
}

type FaceEnhanceConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 强度，取值范围：0.0~1.0。
	// 默认：0.0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intensity *float64 `json:"Intensity,omitnil" name:"Intensity"`
}

type FlowAudio struct {
	// 帧率。
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`

	// 码率，单位是bps。
	Rate *int64 `json:"Rate,omitnil" name:"Rate"`

	// 音频Pid。
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`
}

type FlowLogInfo struct {
	// 时间戳，单位为秒。
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 输入输出类型（input/output）。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入或输出Id。
	InputOutputId *string `json:"InputOutputId,omitnil" name:"InputOutputId"`

	// 协议。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 事件代码。
	EventCode *string `json:"EventCode,omitnil" name:"EventCode"`

	// 事件信息。
	EventMessage *string `json:"EventMessage,omitnil" name:"EventMessage"`

	// 对端IP。
	RemoteIp *string `json:"RemoteIp,omitnil" name:"RemoteIp"`

	// 对端端口。
	RemotePort *string `json:"RemotePort,omitnil" name:"RemotePort"`

	// 主备通道，0为主通道，1为备通道。
	Pipeline *string `json:"Pipeline,omitnil" name:"Pipeline"`

	// 输入或输出的名称。
	InputOutputName *string `json:"InputOutputName,omitnil" name:"InputOutputName"`
}

type FlowMediaAudio struct {
	// 帧率。
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`

	// 码率，单位是bps。
	Rate *int64 `json:"Rate,omitnil" name:"Rate"`

	// 音频Pid。
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// 标志同一次推流。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
}

type FlowMediaInfo struct {
	// 时间戳，单位是秒。
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 总带宽。
	Network *int64 `json:"Network,omitnil" name:"Network"`

	// 传输流的视频数据。
	Video []*FlowMediaVideo `json:"Video,omitnil" name:"Video"`

	// 传输流的音频数据。
	Audio []*FlowMediaAudio `json:"Audio,omitnil" name:"Audio"`

	// 标志同一次推流。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 客户端IP。
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`
}

type FlowMediaVideo struct {
	// 帧率。
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`

	// 码率，单位是bps。
	Rate *int64 `json:"Rate,omitnil" name:"Rate"`

	// 视频Pid。
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// 标志同一次推流。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
}

type FlowRealtimeStatusCommon struct {
	// 当前连接状态，Connected|Waiting|Idle。
	State *string `json:"State,omitnil" name:"State"`

	// 连接模式，Listener|Caller。
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 已连接时长，单位为ms。
	ConnectedTime *int64 `json:"ConnectedTime,omitnil" name:"ConnectedTime"`

	// 实时码率，单位为bps。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 重试次数。
	Reconnections *int64 `json:"Reconnections,omitnil" name:"Reconnections"`
}

type FlowRealtimeStatusItem struct {
	// 类型，Input|Output。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入Id，如果Type为Input，此字段不为空。
	InputId *string `json:"InputId,omitnil" name:"InputId"`

	// 输出Id，如果Type为Output，此字段不为空。
	OutputId *string `json:"OutputId,omitnil" name:"OutputId"`

	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 协议， SRT | RTMP。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 共同状态信息。
	CommonStatus *FlowRealtimeStatusCommon `json:"CommonStatus,omitnil" name:"CommonStatus"`

	// 如果是SRT协议则有此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SRTStatus *FlowRealtimeStatusSRT `json:"SRTStatus,omitnil" name:"SRTStatus"`

	// 如果是RTMP协议则有此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTMPStatus *FlowRealtimeStatusRTMP `json:"RTMPStatus,omitnil" name:"RTMPStatus"`

	// 服务器IP。
	ConnectServerIP *string `json:"ConnectServerIP,omitnil" name:"ConnectServerIP"`

	// 如果是RTP协议则有此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTPStatus *FlowRealtimeStatusRTP `json:"RTPStatus,omitnil" name:"RTPStatus"`
}

type FlowRealtimeStatusRTMP struct {
	// 视频帧率。
	VideoFPS *int64 `json:"VideoFPS,omitnil" name:"VideoFPS"`

	// 音频帧率。
	AudioFPS *int64 `json:"AudioFPS,omitnil" name:"AudioFPS"`
}

type FlowRealtimeStatusRTP struct {
	// 传输的包个数
	Packets *int64 `json:"Packets,omitnil" name:"Packets"`
}

type FlowRealtimeStatusSRT struct {
	// 延迟，单位为ms。
	Latency *int64 `json:"Latency,omitnil" name:"Latency"`

	// RTT，单位为ms。
	RTT *int64 `json:"RTT,omitnil" name:"RTT"`

	// 实时发包数或者收包数。
	Packets *int64 `json:"Packets,omitnil" name:"Packets"`

	// 丢包率。
	PacketLossRate *float64 `json:"PacketLossRate,omitnil" name:"PacketLossRate"`

	// 重传率。
	RetransmitRate *float64 `json:"RetransmitRate,omitnil" name:"RetransmitRate"`

	// 实时丢包数。
	DroppedPackets *int64 `json:"DroppedPackets,omitnil" name:"DroppedPackets"`

	// 是否加密，On|Off。
	Encryption *string `json:"Encryption,omitnil" name:"Encryption"`
}

type FlowSRTInfo struct {
	// 时间戳，单位是秒。
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 发送丢包率。
	SendPacketLossRate *int64 `json:"SendPacketLossRate,omitnil" name:"SendPacketLossRate"`

	// 发送重传率。
	SendRetransmissionRate *int64 `json:"SendRetransmissionRate,omitnil" name:"SendRetransmissionRate"`

	// 接收丢包率。
	RecvPacketLossRate *int64 `json:"RecvPacketLossRate,omitnil" name:"RecvPacketLossRate"`

	// 接收重传率。
	RecvRetransmissionRate *int64 `json:"RecvRetransmissionRate,omitnil" name:"RecvRetransmissionRate"`

	// 与对端的RTT时延。
	RTT *int64 `json:"RTT,omitnil" name:"RTT"`

	// 标志同一次推流。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 发送弃包数。
	SendPacketDropNumber *int64 `json:"SendPacketDropNumber,omitnil" name:"SendPacketDropNumber"`

	// 接收弃包数。
	RecvPacketDropNumber *int64 `json:"RecvPacketDropNumber,omitnil" name:"RecvPacketDropNumber"`
}

type FlowStatistics struct {
	// 会话Id。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 对端IP。
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`

	// 总带宽。
	Network *int64 `json:"Network,omitnil" name:"Network"`

	// 视频数据。
	Video []*FlowVideo `json:"Video,omitnil" name:"Video"`

	// 音频数据。
	Audio []*FlowAudio `json:"Audio,omitnil" name:"Audio"`
}

type FlowStatisticsArray struct {
	// 时间戳。
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 每个会话的统计数据。
	FlowStatistics []*FlowStatistics `json:"FlowStatistics,omitnil" name:"FlowStatistics"`
}

type FlowVideo struct {
	// 帧率。
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`

	// 码率，单位是bps。
	Rate *int64 `json:"Rate,omitnil" name:"Rate"`

	// 音频Pid。
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`
}

type FrameRateConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 帧率，取值范围：[0, 100]，单位：Hz。
	// 默认值 0。
	// 注意：对于转码，该参数会覆盖 VideoTemplate 内部的 Fps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`
}

type FrameTagConfigureInfo struct {
	// 智能按帧标签任务开关，可选值：
	// <li>ON：开启智能按帧标签任务；</li>
	// <li>OFF：关闭智能按帧标签任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type FrameTagConfigureInfoForUpdate struct {
	// 智能按帧标签任务开关，可选值：
	// <li>ON：开启智能按帧标签任务；</li>
	// <li>OFF：关闭智能按帧标签任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type HLSPullSourceAddress struct {
	// HLS源站的Url地址。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type HdrConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 类型，可选值：
	// <li>HDR10</li>
	// <li>HLG</li>
	// 默认值：HDR10。
	// 注意：video的编码方式需要为libx265；
	// 注意：视频编码位深为10。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type HeadTailParameter struct {
	// 片头列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadSet []*MediaInputInfo `json:"HeadSet,omitnil" name:"HeadSet"`

	// 片尾列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TailSet []*MediaInputInfo `json:"TailSet,omitnil" name:"TailSet"`
}

type HighlightSegmentItem struct {
	// 置信度。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 片段起始时间偏移。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 片段结束时间偏移。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`
}

type ImageQualityEnhanceConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 类型，可选值：
	// <li>weak</li>
	// <li>normal</li>
	// <li>strong</li>
	// 默认值：weak。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type ImageSpriteTaskInput struct {
	// 雪碧图模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 截取雪碧图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 截取雪碧图后，雪碧图图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_imageSprite_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil" name:"OutputObjectPath"`

	// 截取雪碧图后，Web VTT 文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_imageSprite_{definition}.{format}`。
	WebVttObjectName *string `json:"WebVttObjectName,omitnil" name:"WebVttObjectName"`

	// 截取雪碧图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitnil" name:"ObjectNumberFormat"`
}

type ImageSpriteTemplate struct {
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 雪碧图模板名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 采样类型。
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitnil" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitnil" name:"ColumnCount"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 图片格式。
	Format *string `json:"Format,omitnil" name:"Format"`
}

type ImageWatermarkInput struct {
	// 水印图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串。支持 jpeg、png 图片格式。
	ImageContent *string `json:"ImageContent,omitnil" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。取值范围为[8, 4096]。</li>
	// 当宽高都不填或者为0时，默认为 10%。
	Width *string `json:"Width,omitnil" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素。取值范围为0或[8, 4096]。</li>
	// 默认值：0px，表示 Height 按照原始水印图片的宽高比缩放。
	Height *string `json:"Height,omitnil" name:"Height"`

	// 水印重复类型。使用场景：水印为动态图像。取值范围：
	// <li>once：动态水印播放完后，不再出现；</li>
	// <li>repeat_last_frame：水印播放完后，停留在最后一帧；</li>
	// <li>repeat：水印循环播放，直到视频结束（默认值）。</li>
	RepeatType *string `json:"RepeatType,omitnil" name:"RepeatType"`
}

type ImageWatermarkInputForUpdate struct {
	// 水印图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串。支持 jpeg、png 图片格式。
	ImageContent *string `json:"ImageContent,omitnil" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。取值范围为[8, 4096]。</li>
	Width *string `json:"Width,omitnil" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素。取值范围为0或[8, 4096]。</li>
	Height *string `json:"Height,omitnil" name:"Height"`

	// 水印重复类型。使用场景：水印为动态图像。取值范围：
	// <li>once：动态水印播放完后，不再出现；</li>
	// <li>repeat_last_frame：水印播放完后，停留在最后一帧；</li>
	// <li>repeat：水印循环播放，直到视频结束。</li>
	RepeatType *string `json:"RepeatType,omitnil" name:"RepeatType"`
}

type ImageWatermarkTemplate struct {
	// 水印图片地址。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	Width *string `json:"Width,omitnil" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素；</li>
	// 0px：表示 Height 按照 Width 对视频宽度的比例缩放。
	Height *string `json:"Height,omitnil" name:"Height"`

	// 水印重复类型。使用场景：水印为动态图像。取值范围：
	// <li>once：动态水印播放完后，不再出现；</li>
	// <li>repeat_last_frame：水印播放完后，停留在最后一帧；</li>
	// <li>repeat：水印循环播放，直到视频结束。</li>
	RepeatType *string `json:"RepeatType,omitnil" name:"RepeatType"`
}

type InputAddress struct {
	// 输入地址的IP。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 输入地址的端口。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type LiveActivityResItem struct {
	// 直播录制任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveRecordTask *LiveScheduleLiveRecordTaskResult `json:"LiveRecordTask,omitnil" name:"LiveRecordTask"`
}

type LiveActivityResult struct {
	// 原子任务类型。
	// <li>LiveRecord：直播录制。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityType *string `json:"ActivityType,omitnil" name:"ActivityType"`

	// 原子任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveActivityResItem *LiveActivityResItem `json:"LiveActivityResItem,omitnil" name:"LiveActivityResItem"`
}

type LiveRecordFile struct {
	// 直播录制文件地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 直播录制文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 直播录制文件时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// 直播录制文件开始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 直播录制文件结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type LiveRecordResult struct {
	// 直播录制文件的目标存储。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 直播录制文件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileList []*LiveRecordFile `json:"FileList,omitnil" name:"FileList"`
}

type LiveRecordTaskInput struct {
	// 直播录制模板 ID。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 直播录制后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 直播录制后文件的输出路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil" name:"OutputObjectPath"`
}

type LiveScheduleLiveRecordTaskResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 直播录制任务的输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *LiveRecordTaskInput `json:"Input,omitnil" name:"Input"`

	// 直播录制任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *LiveRecordResult `json:"Output,omitnil" name:"Output"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`
}

type LiveScheduleTask struct {
	// 直播编排任务 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 源异常时返回非0错误码，返回0 时请使用各个具体任务的 ErrCode。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 源异常时返回对应异常Message，否则请使用各个具体任务的 Message。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 直播流 URL。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 直播编排任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveActivityResultSet []*LiveActivityResult `json:"LiveActivityResultSet,omitnil" name:"LiveActivityResultSet"`
}

type LiveStreamAiAnalysisResultInfo struct {
	// 直播分析子任务结果，暂时只支持直播拆条。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultSet []*LiveStreamAiAnalysisResultItem `json:"ResultSet,omitnil" name:"ResultSet"`
}

type LiveStreamAiAnalysisResultItem struct {
	// 结果的类型，取值范围：
	// <li>SegmentRecognition：拆条。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 拆条结果，当 Type 为
	// SegmentRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentResultSet []*SegmentRecognitionItem `json:"SegmentResultSet,omitnil" name:"SegmentResultSet"`
}

type LiveStreamAiQualityControlResultInfo struct {
	// 质检结果列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityControlResults []*QualityControlResult `json:"QualityControlResults,omitnil" name:"QualityControlResults"`

	// 格式诊断结果列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnoseResults []*DiagnoseResult `json:"DiagnoseResults,omitnil" name:"DiagnoseResults"`
}

type LiveStreamAiRecognitionResultInfo struct {
	// 内容识别结果列表。
	ResultSet []*LiveStreamAiRecognitionResultItem `json:"ResultSet,omitnil" name:"ResultSet"`
}

type LiveStreamAiRecognitionResultItem struct {
	// 结果的类型，取值范围：
	// <li>FaceRecognition：人脸识别，</li>
	// <li>AsrWordsRecognition：语音关键词识别，</li>
	// <li>OcrWordsRecognition：文本关键词识别，</li>
	// <li>AsrFullTextRecognition：语音全文识别，</li>
	// <li>OcrFullTextRecognition：文本全文识别。</li>
	// <li>TransTextRecognition：语音翻译。</li>
	// <li>TagRecognition：精彩打点。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 人脸识别结果，当 Type 为
	// FaceRecognition 时有效。
	FaceRecognitionResultSet []*LiveStreamFaceRecognitionResult `json:"FaceRecognitionResultSet,omitnil" name:"FaceRecognitionResultSet"`

	// 语音关键词识别结果，当 Type 为
	// AsrWordsRecognition 时有效。
	AsrWordsRecognitionResultSet []*LiveStreamAsrWordsRecognitionResult `json:"AsrWordsRecognitionResultSet,omitnil" name:"AsrWordsRecognitionResultSet"`

	// 文本关键词识别结果，当 Type 为
	// OcrWordsRecognition 时有效。
	OcrWordsRecognitionResultSet []*LiveStreamOcrWordsRecognitionResult `json:"OcrWordsRecognitionResultSet,omitnil" name:"OcrWordsRecognitionResultSet"`

	// 语音全文识别结果，当 Type 为
	// AsrFullTextRecognition 时有效。
	AsrFullTextRecognitionResultSet []*LiveStreamAsrFullTextRecognitionResult `json:"AsrFullTextRecognitionResultSet,omitnil" name:"AsrFullTextRecognitionResultSet"`

	// 文本全文识别结果，当 Type 为
	// OcrFullTextRecognition 时有效。
	OcrFullTextRecognitionResultSet []*LiveStreamOcrFullTextRecognitionResult `json:"OcrFullTextRecognitionResultSet,omitnil" name:"OcrFullTextRecognitionResultSet"`

	// 翻译结果，当Type 为 TransTextRecognition 时有效。
	TransTextRecognitionResultSet []*LiveStreamTransTextRecognitionResult `json:"TransTextRecognitionResultSet,omitnil" name:"TransTextRecognitionResultSet"`

	// 打点结果，当Type 为 TagRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagRecognitionResultSet []*LiveStreamTagRecognitionResult `json:"TagRecognitionResultSet,omitnil" name:"TagRecognitionResultSet"`
}

type LiveStreamAiReviewImagePoliticalResult struct {
	// 嫌疑片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 嫌疑片段结束的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 嫌疑片段敏感分数。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 嫌疑片段鉴黄结果建议，取值范围：
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 视频敏感结果标签，取值范围：
	// <li>politician：敏感人物。</li>
	// <li>violation_photo：违规图标。</li>
	Label *string `json:"Label,omitnil" name:"Label"`

	// 敏感人物、违规图标名字。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 敏感人物、违规图标出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	// PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil" name:"PicUrlExpireTime"`
}

type LiveStreamAiReviewImagePornResult struct {
	// 嫌疑片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 嫌疑片段结束的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 嫌疑片段涉黄分数。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 嫌疑片段鉴黄结果建议，取值范围：
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 视频鉴黄结果标签，取值范围：
	// <li>porn：色情。</li>
	// <li>sexy：性感。</li>
	// <li>vulgar：低俗。</li>
	// <li>intimacy：亲密行为。</li>
	Label *string `json:"Label,omitnil" name:"Label"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	// PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil" name:"PicUrlExpireTime"`
}

type LiveStreamAiReviewImageTerrorismResult struct {
	// 嫌疑片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 嫌疑片段结束的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 嫌疑片段涉敏分数。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 嫌疑片段涉敏结果建议，取值范围：
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 视频涉敏结果标签，取值范围：
	// <li>guns：武器枪支。</li>
	// <li>crowd：人群聚集。</li>
	// <li>police：警察部队。</li>
	// <li>bloody：血腥画面。</li>
	// <li>banners：涉敏旗帜。</li>
	// <li>militant：武装分子。</li>
	// <li>explosion：爆炸火灾。</li>
	// <li>terrorists：涉敏人物。</li>
	Label *string `json:"Label,omitnil" name:"Label"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	// PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil" name:"PicUrlExpireTime"`
}

type LiveStreamAiReviewResultInfo struct {
	// 内容审核结果列表。
	ResultSet []*LiveStreamAiReviewResultItem `json:"ResultSet,omitnil" name:"ResultSet"`
}

type LiveStreamAiReviewResultItem struct {
	// 审核结果的类型，可以取的值有：
	// <li>ImagePorn：图片鉴黄</li>
	// <li>ImageTerrorism：图片涉敏</li>
	// <li>ImagePolitical：图片涉敏</li>
	// <li>VoicePorn：声音违规</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 图片鉴黄的结果，当 Type 为 ImagePorn 时有效。
	ImagePornResultSet []*LiveStreamAiReviewImagePornResult `json:"ImagePornResultSet,omitnil" name:"ImagePornResultSet"`

	// 图片涉敏的结果，当 Type 为 ImageTerrorism 时有效。
	ImageTerrorismResultSet []*LiveStreamAiReviewImageTerrorismResult `json:"ImageTerrorismResultSet,omitnil" name:"ImageTerrorismResultSet"`

	// 图片涉敏的结果，当 Type 为 ImagePolitical 时有效。
	ImagePoliticalResultSet []*LiveStreamAiReviewImagePoliticalResult `json:"ImagePoliticalResultSet,omitnil" name:"ImagePoliticalResultSet"`

	// 声音违规的结果，当 Type 为 VoicePorn 时有效。
	VoicePornResultSet []*LiveStreamAiReviewVoicePornResult `json:"VoicePornResultSet,omitnil" name:"VoicePornResultSet"`
}

type LiveStreamAiReviewVoicePornResult struct {
	// 嫌疑片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 嫌疑片段结束的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 嫌疑片段涉黄分数。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 嫌疑片段鉴黄结果建议，取值范围：
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 视频鉴黄结果标签，取值范围：
	// <li>sexual_moan：呻吟。</li>
	Label *string `json:"Label,omitnil" name:"Label"`
}

type LiveStreamAsrFullTextRecognitionResult struct {
	// 识别文本。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 识别片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 识别片段终止的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别开始UTC时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 识别结束UTC时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 稳态标记。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SteadyState *bool `json:"SteadyState,omitnil" name:"SteadyState"`
}

type LiveStreamAsrWordsRecognitionResult struct {
	// 语音关键词。
	Word *string `json:"Word,omitnil" name:"Word"`

	// 识别片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 识别片段终止的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`
}

type LiveStreamFaceRecognitionResult struct {
	// 人物唯一标识 ID。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 人物名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 人物库类型，表示识别出的人物来自哪个人物库：
	// <li>Default：默认人物库；</li><li>UserDefine：用户自定义人物库。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 识别片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 识别片段终止的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`
}

type LiveStreamOcrFullTextRecognitionResult struct {
	// 语音文本。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 识别片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 识别片段终止的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`
}

type LiveStreamOcrWordsRecognitionResult struct {
	// 文本关键词。
	Word *string `json:"Word,omitnil" name:"Word"`

	// 识别片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 识别片段终止的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoords []*int64 `json:"AreaCoords,omitnil" name:"AreaCoords"`
}

type LiveStreamProcessErrorInfo struct {
	// 错误码：
	// <li>0表示没有错误；</li>
	// <li>非0表示错误，请参考 Message 错误信息。</li>
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`
}

type LiveStreamProcessTask struct {
	// 媒体处理任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，0 表示成功，其他值表示失败。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 直播流 URL。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type LiveStreamTagRecognitionResult struct {
	// 打点事件。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 识别片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 识别片段终止的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`
}

type LiveStreamTaskNotifyConfig struct {
	// CMQ 的模型，有 Queue 和 Topic 两种，目前仅支持 Queue。
	CmqModel *string `json:"CmqModel,omitnil" name:"CmqModel"`

	// CMQ 的园区，如 sh，bj 等。
	CmqRegion *string `json:"CmqRegion,omitnil" name:"CmqRegion"`

	// 当模型为 Queue 时有效，表示接收事件通知的 CMQ 的队列名。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// 当模型为 Topic 时有效，表示接收事件通知的 CMQ 的主题名。
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 通知类型，默认CMQ，指定URL时HTTP回调推送到 NotifyUrl 指定的地址。
	// 
	// <font color="red"> 注：不填或为空时默认 CMQ，如需采用其他类型需填写对应类型值。 </font>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// HTTP回调地址，NotifyType为URL时必填。
	NotifyUrl *string `json:"NotifyUrl,omitnil" name:"NotifyUrl"`
}

type LiveStreamTransTextRecognitionResult struct {
	// 识别文本。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 翻译片段起始的 PTS 时间，单位：秒。
	StartPtsTime *float64 `json:"StartPtsTime,omitnil" name:"StartPtsTime"`

	// 翻译片段终止的 PTS 时间，单位：秒。
	EndPtsTime *float64 `json:"EndPtsTime,omitnil" name:"EndPtsTime"`

	// 翻译片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 翻译文本。
	Trans *string `json:"Trans,omitnil" name:"Trans"`

	// 翻译开始UTC时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 翻译结束UTC时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 稳态标记。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SteadyState *bool `json:"SteadyState,omitnil" name:"SteadyState"`
}

type LowLightEnhanceConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 类型，可选值：
	// <li>normal</li>
	// 默认值：normal。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

// Predefined struct for user
type ManageTaskRequestParams struct {
	// 操作类型，取值范围：
	// <ul>
	// <li>Abort：终止任务。使用说明：
	// <ul><li>若 [任务类型](/document/product/862/37614#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) 为直播流处理任务（LiveStreamProcessTask），支持终止 [任务状态](/document/product/862/37614#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) 为等待中（WAITING）或处理中（PROCESSING）的任务；</li>
	// <li>否则，对于其他 [任务类型](/document/product/862/37614#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0)，只支持终止 [任务状态](/document/product/862/37614#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) 为等待中（WAITING）的任务。</li></ul>
	// </li></ul>
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 视频处理的任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type ManageTaskRequest struct {
	*tchttp.BaseRequest
	
	// 操作类型，取值范围：
	// <ul>
	// <li>Abort：终止任务。使用说明：
	// <ul><li>若 [任务类型](/document/product/862/37614#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) 为直播流处理任务（LiveStreamProcessTask），支持终止 [任务状态](/document/product/862/37614#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) 为等待中（WAITING）或处理中（PROCESSING）的任务；</li>
	// <li>否则，对于其他 [任务类型](/document/product/862/37614#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0)，只支持终止 [任务状态](/document/product/862/37614#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) 为等待中（WAITING）的任务。</li></ul>
	// </li></ul>
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 视频处理的任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *ManageTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationType")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ManageTaskResponse struct {
	*tchttp.BaseResponse
	Response *ManageTaskResponseParams `json:"Response"`
}

func (r *ManageTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MediaAiAnalysisClassificationItem struct {
	// 智能分类的类别名称。
	Classification *string `json:"Classification,omitnil" name:"Classification"`

	// 智能分类的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`
}

type MediaAiAnalysisCoverItem struct {
	// 智能封面存储路径。
	CoverPath *string `json:"CoverPath,omitnil" name:"CoverPath"`

	// 智能封面的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`
}

type MediaAiAnalysisFrameTagItem struct {
	// 按帧标签名称。
	Tag *string `json:"Tag,omitnil" name:"Tag"`

	// 按帧标签名称的分类列表，CategorySet.N 表示第 N+1级分类。
	// 比如 Tag 为“塔楼”时，CategorySet 包含两个元素：CategorySet.0 为“场景”，CategorySet.1为 “建筑”，表示按帧标签为“塔楼”，且第1级分类是“场景”，第2级分类是“建筑”。
	CategorySet []*string `json:"CategorySet,omitnil" name:"CategorySet"`

	// 按帧标签的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`
}

type MediaAiAnalysisFrameTagSegmentItem struct {
	// 按帧标签起始的偏移时间。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 按帧标签结束的偏移时间。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 时间片段内的标签列表。
	TagSet []*MediaAiAnalysisFrameTagItem `json:"TagSet,omitnil" name:"TagSet"`
}

type MediaAiAnalysisHighlightItem struct {
	// 智能精彩集锦地址。
	HighlightPath *string `json:"HighlightPath,omitnil" name:"HighlightPath"`

	// 智能精彩集锦封面地址。
	CovImgPath *string `json:"CovImgPath,omitnil" name:"CovImgPath"`

	// 智能精彩集锦的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 智能精彩集锦持续时间。
	Duration *float64 `json:"Duration,omitnil" name:"Duration"`

	// 智能精彩集锦子片段列表。
	SegmentSet []*HighlightSegmentItem `json:"SegmentSet,omitnil" name:"SegmentSet"`
}

type MediaAiAnalysisTagItem struct {
	// 标签名称。
	Tag *string `json:"Tag,omitnil" name:"Tag"`

	// 标签的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`
}

type MediaAnimatedGraphicsItem struct {
	// 转动图文件的存储位置。
	Storage *TaskOutputStorage `json:"Storage,omitnil" name:"Storage"`

	// 转动图的文件路径。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 转动图模板 ID，参见[转动图参数模板](https://cloud.tencent.com/document/product/862/37042#.E9.A2.84.E7.BD.AE.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 动图格式，如 gif。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 动图的高度，单位：px。
	Height *int64 `json:"Height,omitnil" name:"Height"`

	// 动图的宽度，单位：px。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 动图码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 动图大小，单位：字节。
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 动图的md5值。
	Md5 *string `json:"Md5,omitnil" name:"Md5"`

	// 动图在视频中的起始时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 动图在视频中的结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`
}

type MediaAudioStreamItem struct {
	// 音频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 音频流的采样率，单位：hz。
	SamplingRate *int64 `json:"SamplingRate,omitnil" name:"SamplingRate"`

	// 音频流的编码格式，例如 aac。
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// 音频声道数，例如 2。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channel *int64 `json:"Channel,omitnil" name:"Channel"`

	// 音频Codecs。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codecs *string `json:"Codecs,omitnil" name:"Codecs"`

	// 音频响度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Loudness *float64 `json:"Loudness,omitnil" name:"Loudness"`
}

type MediaContentReviewAsrTextSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 嫌疑片段审核结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 嫌疑关键词列表。
	KeywordSet []*string `json:"KeywordSet,omitnil" name:"KeywordSet"`
}

type MediaContentReviewOcrTextSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 嫌疑片段审核结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 嫌疑关键词列表。
	KeywordSet []*string `json:"KeywordSet,omitnil" name:"KeywordSet"`

	// 嫌疑文字出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	// PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil" name:"PicUrlExpireTime"`
}

type MediaContentReviewPoliticalSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 嫌疑片段涉敏分数。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 嫌疑片段涉敏结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 涉敏人物、违规图标名字。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 嫌疑片段涉敏结果标签。内容审核模板[画面涉敏任务控制参数](https://cloud.tencent.com/document/api/862/37615#PoliticalImgReviewTemplateInfo)里 LabelSet 参数与此参数取值范围的对应关系：
	// violation_photo：
	// <li>violation_photo：违规图标。</li>
	// politician：
	// <li>nation_politician：国家领导人；</li>
	// <li>province_politician: 省部级领导人；</li>
	// <li>bureau_politician：厅局级领导人；</li>
	// <li>county_politician：县处级领导人；</li>
	// <li>rural_politician：乡科级领导人；</li>
	// <li>sensitive_politician：涉敏人物；</li>
	// <li>foreign_politician：国外领导人。</li>
	// entertainment：
	// <li>sensitive_entertainment：敏感娱乐人物。</li>
	// sport：
	// <li>sensitive_sport：敏感体育人物。</li>
	// entrepreneur：
	// <li>sensitive_entrepreneur：敏感商业人物。</li>
	// scholar：
	// <li>sensitive_scholar：敏感教育学者。</li>
	// celebrity：
	// <li>sensitive_celebrity：敏感知名人物；</li>
	// <li>historical_celebrity：历史知名人物。</li>
	// military：
	// <li>sensitive_military：敏感军事人物。</li>
	Label *string `json:"Label,omitnil" name:"Label"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	//  PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 涉敏人物、违规图标出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil" name:"PicUrlExpireTime"`
}

type MediaContentReviewSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 嫌疑片段涉黄分数。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 嫌疑片段鉴黄结果标签。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 嫌疑片段鉴黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	//  PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil" name:"PicUrlExpireTime"`
}

type MediaImageSpriteItem struct {
	// 雪碧图规格，参见[雪碧图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 雪碧图小图的高度。
	Height *int64 `json:"Height,omitnil" name:"Height"`

	// 雪碧图小图的宽度。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 每一张雪碧图大图里小图的数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 每一张雪碧图大图的路径。
	ImagePathSet []*string `json:"ImagePathSet,omitnil" name:"ImagePathSet"`

	// 雪碧图子图位置与时间关系的 WebVtt 文件路径。WebVtt 文件表明了各个雪碧图小图对应的时间点，以及在雪碧大图里的坐标位置，一般被播放器用于实现预览。
	WebVttPath *string `json:"WebVttPath,omitnil" name:"WebVttPath"`

	// 雪碧图文件的存储位置。
	Storage *TaskOutputStorage `json:"Storage,omitnil" name:"Storage"`
}

type MediaInputInfo struct {
	// 输入来源对象的类型，支持：
	// <li> COS：COS源</li>
	// <li> URL：URL源</li>
	// <li> AWS-S3：AWS 源，目前只支持转码任务 </li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 当 Type 为 COS 时有效，则该项为必填，表示媒体处理 COS 对象信息。
	CosInputInfo *CosInputInfo `json:"CosInputInfo,omitnil" name:"CosInputInfo"`

	// 当 Type 为 URL 时有效，则该项为必填，表示媒体处理 URL 对象信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlInputInfo *UrlInputInfo `json:"UrlInputInfo,omitnil" name:"UrlInputInfo"`

	// 当 Type 为 AWS-S3 时有效，则该项为必填，表示媒体处理 AWS S3 对象信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3InputInfo *S3InputInfo `json:"S3InputInfo,omitnil" name:"S3InputInfo"`
}

type MediaMetaData struct {
	// 上传的媒体文件大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 视频流码率平均值与音频流码率平均值之和，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitnil" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitnil" name:"Duration"`

	// 视频拍摄时的选择角度，单位：度。
	Rotate *int64 `json:"Rotate,omitnil" name:"Rotate"`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitnil" name:"VideoStreamSet"`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitnil" name:"AudioStreamSet"`

	// 视频时长，单位：秒。
	VideoDuration *float64 `json:"VideoDuration,omitnil" name:"VideoDuration"`

	// 音频时长，单位：秒。
	AudioDuration *float64 `json:"AudioDuration,omitnil" name:"AudioDuration"`
}

type MediaProcessTaskAdaptiveDynamicStreamingResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 对视频转自适应码流任务的输入。
	Input *AdaptiveDynamicStreamingTaskInput `json:"Input,omitnil" name:"Input"`

	// 对视频转自适应码流任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AdaptiveDynamicStreamingInfoItem `json:"Output,omitnil" name:"Output"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`
}

type MediaProcessTaskAnimatedGraphicResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 转动图任务的输入。
	Input *AnimatedGraphicTaskInput `json:"Input,omitnil" name:"Input"`

	// 转动图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaAnimatedGraphicsItem `json:"Output,omitnil" name:"Output"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`
}

type MediaProcessTaskImageSpriteResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 对视频截雪碧图任务的输入。
	Input *ImageSpriteTaskInput `json:"Input,omitnil" name:"Input"`

	// 对视频截雪碧图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaImageSpriteItem `json:"Output,omitnil" name:"Output"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`
}

type MediaProcessTaskInput struct {
	// 视频转码任务列表。
	TranscodeTaskSet []*TranscodeTaskInput `json:"TranscodeTaskSet,omitnil" name:"TranscodeTaskSet"`

	// 视频转动图任务列表。
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInput `json:"AnimatedGraphicTaskSet,omitnil" name:"AnimatedGraphicTaskSet"`

	// 对视频按时间点截图任务列表。
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTaskSet,omitnil" name:"SnapshotByTimeOffsetTaskSet"`

	// 对视频采样截图任务列表。
	SampleSnapshotTaskSet []*SampleSnapshotTaskInput `json:"SampleSnapshotTaskSet,omitnil" name:"SampleSnapshotTaskSet"`

	// 对视频截雪碧图任务列表。
	ImageSpriteTaskSet []*ImageSpriteTaskInput `json:"ImageSpriteTaskSet,omitnil" name:"ImageSpriteTaskSet"`

	// 转自适应码流任务列表。
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTaskSet,omitnil" name:"AdaptiveDynamicStreamingTaskSet"`
}

type MediaProcessTaskResult struct {
	// 任务的类型，可以取的值有：
	// <li>Transcode：转码</li>
	// <li>AnimatedGraphics：转动图</li>
	// <li>SnapshotByTimeOffset：时间点截图</li>
	// <li>SampleSnapshot：采样截图</li>
	// <li>ImageSprites：雪碧图</li>
	// <li>CoverBySnapshot：截图做封面</li>
	// <li>AdaptiveDynamicStreaming：自适应码流</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 视频转码任务的查询结果，当任务类型为 Transcode 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitnil" name:"TranscodeTask"`

	// 视频转动图任务的查询结果，当任务类型为 AnimatedGraphics 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitnil" name:"AnimatedGraphicTask"`

	// 对视频按时间点截图任务的查询结果，当任务类型为 SnapshotByTimeOffset 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitnil" name:"SnapshotByTimeOffsetTask"`

	// 对视频采样截图任务的查询结果，当任务类型为 SampleSnapshot 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitnil" name:"SampleSnapshotTask"`

	// 对视频截雪碧图任务的查询结果，当任务类型为 ImageSprite 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitnil" name:"ImageSpriteTask"`

	// 转自适应码流任务查询结果，当任务类型为 AdaptiveDynamicStreaming 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingTask *MediaProcessTaskAdaptiveDynamicStreamingResult `json:"AdaptiveDynamicStreamingTask,omitnil" name:"AdaptiveDynamicStreamingTask"`
}

type MediaProcessTaskSampleSnapshotResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 对视频做采样截图任务输入。
	Input *SampleSnapshotTaskInput `json:"Input,omitnil" name:"Input"`

	// 对视频做采样截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSampleSnapshotItem `json:"Output,omitnil" name:"Output"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`
}

type MediaProcessTaskSnapshotByTimeOffsetResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 对视频按指定时间点截图任务输入。
	Input *SnapshotByTimeOffsetTaskInput `json:"Input,omitnil" name:"Input"`

	// 对视频按指定时间点截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSnapshotByTimeOffsetItem `json:"Output,omitnil" name:"Output"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`
}

type MediaProcessTaskTranscodeResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 转码任务的输入。
	Input *TranscodeTaskInput `json:"Input,omitnil" name:"Input"`

	// 转码任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaTranscodeItem `json:"Output,omitnil" name:"Output"`

	// 转码进度，取值范围 [0-100]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitnil" name:"Progress"`
}

type MediaSampleSnapshotItem struct {
	// 采样截图规格 ID，参见[采样截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 采样方式，取值范围：
	// <li>Percent：根据百分比间隔采样。</li>
	// <li>Time：根据时间间隔采样。</li>
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔
	// <li>当 SampleType 为 Percent 时，该值表示多少百分比一张图。</li>
	// <li>当 SampleType 为 Time 时，该值表示多少时间间隔一张图，单位秒， 第一张图均为视频首帧。</li>
	Interval *int64 `json:"Interval,omitnil" name:"Interval"`

	// 截图后文件的存储位置。
	Storage *TaskOutputStorage `json:"Storage,omitnil" name:"Storage"`

	// 生成的截图 path 列表。
	ImagePathSet []*string `json:"ImagePathSet,omitnil" name:"ImagePathSet"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitnil" name:"WaterMarkDefinition"`
}

type MediaSnapshotByTimeOffsetItem struct {
	// 指定时间点截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 同一规格的截图信息集合，每个元素代表一张截图。
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet,omitnil" name:"PicInfoSet"`

	// 指定时间点截图文件的存储位置。
	Storage *TaskOutputStorage `json:"Storage,omitnil" name:"Storage"`
}

type MediaSnapshotByTimePicInfoItem struct {
	// 该张截图对应视频文件中的时间偏移，单位为秒。
	TimeOffset *float64 `json:"TimeOffset,omitnil" name:"TimeOffset"`

	// 该张截图的路径。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitnil" name:"WaterMarkDefinition"`
}

type MediaTranscodeItem struct {
	// 转码后文件的目标存储。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 转码后的视频文件路径。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/862/37042)。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitnil" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 媒体文件总大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitnil" name:"Duration"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 视频的 md5 值。
	Md5 *string `json:"Md5,omitnil" name:"Md5"`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitnil" name:"AudioStreamSet"`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitnil" name:"VideoStreamSet"`
}

type MediaVideoStreamItem struct {
	// 视频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 视频流的高度，单位：px。
	Height *int64 `json:"Height,omitnil" name:"Height"`

	// 视频流的宽度，单位：px。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 视频流的编码格式，例如 h264。
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// 帧率，单位：hz。
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`

	// 色彩空间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColorPrimaries *string `json:"ColorPrimaries,omitnil" name:"ColorPrimaries"`

	// 色彩空间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColorSpace *string `json:"ColorSpace,omitnil" name:"ColorSpace"`

	// 色彩空间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColorTransfer *string `json:"ColorTransfer,omitnil" name:"ColorTransfer"`

	// HDR类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HdrType *string `json:"HdrType,omitnil" name:"HdrType"`

	// 视频Codecs。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codecs *string `json:"Codecs,omitnil" name:"Codecs"`
}

// Predefined struct for user
type ModifyAIAnalysisTemplateRequestParams struct {
	// 视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfoForUpdate `json:"ClassificationConfigure,omitnil" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfoForUpdate `json:"TagConfigure,omitnil" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfoForUpdate `json:"CoverConfigure,omitnil" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfoForUpdate `json:"FrameTagConfigure,omitnil" name:"FrameTagConfigure"`
}

type ModifyAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfoForUpdate `json:"ClassificationConfigure,omitnil" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfoForUpdate `json:"TagConfigure,omitnil" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfoForUpdate `json:"CoverConfigure,omitnil" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfoForUpdate `json:"FrameTagConfigure,omitnil" name:"FrameTagConfigure"`
}

func (r *ModifyAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAIAnalysisTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ClassificationConfigure")
	delete(f, "TagConfigure")
	delete(f, "CoverConfigure")
	delete(f, "FrameTagConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAIAnalysisTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAIAnalysisTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAIAnalysisTemplateResponseParams `json:"Response"`
}

func (r *ModifyAIAnalysisTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAIAnalysisTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAIRecognitionTemplateRequestParams struct {
	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitnil" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitnil" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitnil" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitnil" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitnil" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	TranslateConfigure *TranslateConfigureInfoForUpdate `json:"TranslateConfigure,omitnil" name:"TranslateConfigure"`
}

type ModifyAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitnil" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitnil" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitnil" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitnil" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitnil" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	TranslateConfigure *TranslateConfigureInfoForUpdate `json:"TranslateConfigure,omitnil" name:"TranslateConfigure"`
}

func (r *ModifyAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAIRecognitionTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "FaceConfigure")
	delete(f, "OcrFullTextConfigure")
	delete(f, "OcrWordsConfigure")
	delete(f, "AsrFullTextConfigure")
	delete(f, "AsrWordsConfigure")
	delete(f, "TranslateConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAIRecognitionTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAIRecognitionTemplateResponseParams `json:"Response"`
}

func (r *ModifyAIRecognitionTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAIRecognitionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAdaptiveDynamicStreamingTemplateRequestParams struct {
	// 转自适应码流模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 转自适应码流格式，取值范围：
	// <li>HLS，</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil" name:"Format"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil" name:"DisableHigherVideoResolution"`

	// 转自适应码流输入流参数信息，最多输入10路流。
	// 注意：各个流的帧率必须保持一致；如果不一致，采用第一个流的帧率作为输出帧率。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil" name:"StreamInfos"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`
}

type ModifyAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 转自适应码流模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 转自适应码流格式，取值范围：
	// <li>HLS，</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil" name:"Format"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil" name:"DisableHigherVideoResolution"`

	// 转自适应码流输入流参数信息，最多输入10路流。
	// 注意：各个流的帧率必须保持一致；如果不一致，采用第一个流的帧率作为输出帧率。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil" name:"StreamInfos"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`
}

func (r *ModifyAdaptiveDynamicStreamingTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAdaptiveDynamicStreamingTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Format")
	delete(f, "DisableHigherVideoBitrate")
	delete(f, "DisableHigherVideoResolution")
	delete(f, "StreamInfos")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAdaptiveDynamicStreamingTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAdaptiveDynamicStreamingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAdaptiveDynamicStreamingTemplateResponseParams `json:"Response"`
}

func (r *ModifyAdaptiveDynamicStreamingTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAdaptiveDynamicStreamingTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAnimatedGraphicsTemplateRequestParams struct {
	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 动图格式，取值为 gif 和 webp。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 帧率，取值范围：[1, 30]，单位：Hz。
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitnil" name:"Quality"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`
}

type ModifyAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 动图格式，取值为 gif 和 webp。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 帧率，取值范围：[1, 30]，单位：Hz。
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitnil" name:"Quality"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`
}

func (r *ModifyAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Fps")
	delete(f, "Quality")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAnimatedGraphicsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAnimatedGraphicsTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAnimatedGraphicsTemplateResponseParams `json:"Response"`
}

func (r *ModifyAnimatedGraphicsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContentReviewTemplateRequestParams struct {
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 令人反感的信息的控制参数。
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitnil" name:"PornConfigure"`

	// 令人不安全的信息的控制参数。
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitnil" name:"TerrorismConfigure"`

	// 令人不适宜的控制参数。
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitnil" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	// 注意：此参数尚未支持。
	ProhibitedConfigure *ProhibitedConfigureInfoForUpdate `json:"ProhibitedConfigure,omitnil" name:"ProhibitedConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitnil" name:"UserDefineConfigure"`
}

type ModifyContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 令人反感的信息的控制参数。
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitnil" name:"PornConfigure"`

	// 令人不安全的信息的控制参数。
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitnil" name:"TerrorismConfigure"`

	// 令人不适宜的控制参数。
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitnil" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	// 注意：此参数尚未支持。
	ProhibitedConfigure *ProhibitedConfigureInfoForUpdate `json:"ProhibitedConfigure,omitnil" name:"ProhibitedConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitnil" name:"UserDefineConfigure"`
}

func (r *ModifyContentReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContentReviewTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "PornConfigure")
	delete(f, "TerrorismConfigure")
	delete(f, "PoliticalConfigure")
	delete(f, "ProhibitedConfigure")
	delete(f, "UserDefineConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContentReviewTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyContentReviewTemplateResponseParams `json:"Response"`
}

func (r *ModifyContentReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContentReviewTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSpriteTemplateRequestParams struct {
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitnil" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitnil" name:"ColumnCount"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 图片格式，取值可以为 jpg、png、webp。
	Format *string `json:"Format,omitnil" name:"Format"`
}

type ModifyImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitnil" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitnil" name:"ColumnCount"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 图片格式，取值可以为 jpg、png、webp。
	Format *string `json:"Format,omitnil" name:"Format"`
}

func (r *ModifyImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSpriteTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "SampleType")
	delete(f, "SampleInterval")
	delete(f, "RowCount")
	delete(f, "ColumnCount")
	delete(f, "FillType")
	delete(f, "Comment")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageSpriteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSpriteTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageSpriteTemplateResponseParams `json:"Response"`
}

func (r *ModifyImageSpriteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSpriteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInput struct {
	// 输入Id。
	InputId *string `json:"InputId,omitnil" name:"InputId"`

	// 输入名称。
	InputName *string `json:"InputName,omitnil" name:"InputName"`

	// 输入描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 允许的推流的IP，CIDR格式。
	AllowIpList []*string `json:"AllowIpList,omitnil" name:"AllowIpList"`

	// SRT的配置信息。
	SRTSettings *CreateInputSRTSettings `json:"SRTSettings,omitnil" name:"SRTSettings"`

	// RTP的配置信息。
	RTPSettings *CreateInputRTPSettings `json:"RTPSettings,omitnil" name:"RTPSettings"`

	// 输入的协议，可选[SRT|RTP|RTMP]。
	// 当输出包含RTP时，输入只能是RTP。
	// 当输出包含RTMP时，输入可以是SRT/RTMP。
	// 当输出包含SRT时，输入只能是SRT。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 输入的主备开关，可选[OPEN|CLOSE]。
	FailOver *string `json:"FailOver,omitnil" name:"FailOver"`

	// RTMP_PULL的配置信息。
	RTMPPullSettings *CreateInputRTMPPullSettings `json:"RTMPPullSettings,omitnil" name:"RTMPPullSettings"`

	// RTSP_PULL的配置信息。
	RTSPPullSettings *CreateInputRTSPPullSettings `json:"RTSPPullSettings,omitnil" name:"RTSPPullSettings"`

	// HLS_PULL的配置信息。
	HLSPullSettings *CreateInputHLSPullSettings `json:"HLSPullSettings,omitnil" name:"HLSPullSettings"`

	// 延播平滑吐流配置信息。
	ResilientStream *ResilientStreamConf `json:"ResilientStream,omitnil" name:"ResilientStream"`

	// 绑定的输入安全组 ID。 仅支持关联一组安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type ModifyOutputInfo struct {
	// 需要修改的Output的Id。
	OutputId *string `json:"OutputId,omitnil" name:"OutputId"`

	// 输出的名称。
	OutputName *string `json:"OutputName,omitnil" name:"OutputName"`

	// 输出的描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 输出的转推协议，支持SRT|RTP|RTMP。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 转推SRT的配置。
	SRTSettings *CreateOutputSRTSettings `json:"SRTSettings,omitnil" name:"SRTSettings"`

	// 转推RTP的配置。
	RTPSettings *CreateOutputInfoRTPSettings `json:"RTPSettings,omitnil" name:"RTPSettings"`

	// 转推RTMP的配置。
	RTMPSettings *CreateOutputRTMPSettings `json:"RTMPSettings,omitnil" name:"RTMPSettings"`

	// IP白名单列表，格式为CIDR，如0.0.0.0/0。
	// 当Protocol为RTMP_PULL有效，为空代表不限制客户端IP。
	AllowIpList []*string `json:"AllowIpList,omitnil" name:"AllowIpList"`

	// 最大拉流并发数，最大4，默认4。
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitnil" name:"MaxConcurrent"`

	// 绑定的安全组 ID。 仅支持关联一组安全组。	
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

// Predefined struct for user
type ModifyPersonSampleRequestParams struct {
	// 素材 ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 名称，长度限制：128 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 素材应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于不适宜的内容识别，等价于 Review.Face。
	// 3. All：用于内容识别、不适宜的内容识别，等价于 1+2。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 五官操作信息。
	FaceOperationInfo *AiSampleFaceOperation `json:"FaceOperationInfo,omitnil" name:"FaceOperationInfo"`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil" name:"TagOperationInfo"`
}

type ModifyPersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// 素材 ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 名称，长度限制：128 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 素材应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于不适宜的内容识别，等价于 Review.Face。
	// 3. All：用于内容识别、不适宜的内容识别，等价于 1+2。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 五官操作信息。
	FaceOperationInfo *AiSampleFaceOperation `json:"FaceOperationInfo,omitnil" name:"FaceOperationInfo"`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil" name:"TagOperationInfo"`
}

func (r *ModifyPersonSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Usages")
	delete(f, "FaceOperationInfo")
	delete(f, "TagOperationInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPersonSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonSampleResponseParams struct {
	// 素材信息。
	Person *AiSamplePerson `json:"Person,omitnil" name:"Person"`

	// 处理失败的五官信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitnil" name:"FailFaceInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyPersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPersonSampleResponseParams `json:"Response"`
}

func (r *ModifyPersonSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySampleSnapshotTemplateRequestParams struct {
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 图片格式，取值为 jpg、png、webp。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

type ModifySampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 图片格式，取值为 jpg、png、webp。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

func (r *ModifySampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySampleSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "SampleType")
	delete(f, "SampleInterval")
	delete(f, "Format")
	delete(f, "Comment")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySampleSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySampleSnapshotTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifySampleSnapshotTemplateResponseParams `json:"Response"`
}

func (r *ModifySampleSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySampleSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScheduleRequestParams struct {
	// 编排唯一标识。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`

	// 编排名称。
	ScheduleName *string `json:"ScheduleName,omitnil" name:"ScheduleName"`

	// 编排绑定的触发规则。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 编排任务列表。
	// 注意：内部不允许部分更新，如果需要更新需全量提交编排任务列表。
	Activities []*Activity `json:"Activities,omitnil" name:"Activities"`

	// 媒体处理的文件输出存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理生成的文件输出的目标目录，必选以 / 开头和结尾。
	// 注意：如果设置为空，则表示取消老配置的OutputDir值。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 任务的事件通知配置。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`
}

type ModifyScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 编排唯一标识。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`

	// 编排名称。
	ScheduleName *string `json:"ScheduleName,omitnil" name:"ScheduleName"`

	// 编排绑定的触发规则。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 编排任务列表。
	// 注意：内部不允许部分更新，如果需要更新需全量提交编排任务列表。
	Activities []*Activity `json:"Activities,omitnil" name:"Activities"`

	// 媒体处理的文件输出存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理生成的文件输出的目标目录，必选以 / 开头和结尾。
	// 注意：如果设置为空，则表示取消老配置的OutputDir值。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 任务的事件通知配置。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`
}

func (r *ModifyScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	delete(f, "ScheduleName")
	delete(f, "Trigger")
	delete(f, "Activities")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "TaskNotifyConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScheduleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyScheduleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyScheduleResponseParams `json:"Response"`
}

func (r *ModifyScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotByTimeOffsetTemplateRequestParams struct {
	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 图片格式，取值可以为 jpg、png、webp。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

type ModifySnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 图片格式，取值可以为 jpg、png、webp。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

func (r *ModifySnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Comment")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotByTimeOffsetTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotByTimeOffsetTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotByTimeOffsetTemplateResponseParams `json:"Response"`
}

func (r *ModifySnapshotByTimeOffsetTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkEventRequestParams struct {
	// 媒体传输事件Event Id。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 需要修改的事件名称。
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// Event的描述信息。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type ModifyStreamLinkEventRequest struct {
	*tchttp.BaseRequest
	
	// 媒体传输事件Event Id。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 需要修改的事件名称。
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// Event的描述信息。
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *ModifyStreamLinkEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "EventName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLinkEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkEventResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamLinkEventResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLinkEventResponseParams `json:"Response"`
}

func (r *ModifyStreamLinkEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkFlowRequestParams struct {
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 需要修改的流名称。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`
}

type ModifyStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 需要修改的流名称。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`
}

func (r *ModifyStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "FlowName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLinkFlowResponseParams `json:"Response"`
}

func (r *ModifyStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkInputRequestParams struct {
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 需要修改的Input信息。
	Input *ModifyInput `json:"Input,omitnil" name:"Input"`
}

type ModifyStreamLinkInputRequest struct {
	*tchttp.BaseRequest
	
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 需要修改的Input信息。
	Input *ModifyInput `json:"Input,omitnil" name:"Input"`
}

func (r *ModifyStreamLinkInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkInputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Input")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLinkInputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkInputResponseParams struct {
	// 修改后的Input信息。
	Info *DescribeInput `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamLinkInputResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLinkInputResponseParams `json:"Response"`
}

func (r *ModifyStreamLinkInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkInputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkOutputInfoRequestParams struct {
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 需要修改的Output配置。
	Output *ModifyOutputInfo `json:"Output,omitnil" name:"Output"`
}

type ModifyStreamLinkOutputInfoRequest struct {
	*tchttp.BaseRequest
	
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 需要修改的Output配置。
	Output *ModifyOutputInfo `json:"Output,omitnil" name:"Output"`
}

func (r *ModifyStreamLinkOutputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkOutputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Output")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLinkOutputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkOutputInfoResponseParams struct {
	// 修改后的Output配置。
	Info *DescribeOutput `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamLinkOutputInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLinkOutputInfoResponseParams `json:"Response"`
}

func (r *ModifyStreamLinkOutputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkOutputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTranscodeTemplateRequestParams struct {
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil" name:"RemoveAudio"`

	// 视频流配置参数。
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitnil" name:"VideoTemplate"`

	// 音频流配置参数。
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitnil" name:"AudioTemplate"`

	// 极速高清转码参数。
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitnil" name:"TEHDConfig"`

	// 音视频增强参数。
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil" name:"EnhanceConfig"`
}

type ModifyTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil" name:"RemoveAudio"`

	// 视频流配置参数。
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitnil" name:"VideoTemplate"`

	// 音频流配置参数。
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitnil" name:"AudioTemplate"`

	// 极速高清转码参数。
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitnil" name:"TEHDConfig"`

	// 音视频增强参数。
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil" name:"EnhanceConfig"`
}

func (r *ModifyTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Container")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	delete(f, "VideoTemplate")
	delete(f, "AudioTemplate")
	delete(f, "TEHDConfig")
	delete(f, "EnhanceConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTranscodeTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTranscodeTemplateResponseParams `json:"Response"`
}

func (r *ModifyTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWatermarkTemplateRequestParams struct {
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	XPos *string `json:"XPos,omitnil" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	YPos *string `json:"YPos,omitnil" name:"YPos"`

	// 图片水印模板，该字段仅对图片水印模板有效。
	ImageTemplate *ImageWatermarkInputForUpdate `json:"ImageTemplate,omitnil" name:"ImageTemplate"`

	// 文字水印模板，该字段仅对文字水印模板有效。
	TextTemplate *TextWatermarkTemplateInputForUpdate `json:"TextTemplate,omitnil" name:"TextTemplate"`

	// SVG水印模板，当 Type 为 svg，该字段必填。当 Type 为 image 或 text，该字段无效。
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitnil" name:"SvgTemplate"`
}

type ModifyWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	XPos *string `json:"XPos,omitnil" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	YPos *string `json:"YPos,omitnil" name:"YPos"`

	// 图片水印模板，该字段仅对图片水印模板有效。
	ImageTemplate *ImageWatermarkInputForUpdate `json:"ImageTemplate,omitnil" name:"ImageTemplate"`

	// 文字水印模板，该字段仅对文字水印模板有效。
	TextTemplate *TextWatermarkTemplateInputForUpdate `json:"TextTemplate,omitnil" name:"TextTemplate"`

	// SVG水印模板，当 Type 为 svg，该字段必填。当 Type 为 image 或 text，该字段无效。
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitnil" name:"SvgTemplate"`
}

func (r *ModifyWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWatermarkTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "CoordinateOrigin")
	delete(f, "XPos")
	delete(f, "YPos")
	delete(f, "ImageTemplate")
	delete(f, "TextTemplate")
	delete(f, "SvgTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWatermarkTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWatermarkTemplateResponseParams struct {
	// 图片水印地址，仅当 ImageTemplate.ImageContent 非空，该字段有效。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWatermarkTemplateResponseParams `json:"Response"`
}

func (r *ModifyWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWatermarkTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWordSampleRequestParams struct {
	// 关键词，长度限制：128 个字符。
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// <b>关键词应用场景，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜的内容识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜的音频识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜的内容识别，等价于 3+4；
	// 7. All：包含以上全部，等价于 1+2+3+4。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil" name:"TagOperationInfo"`
}

type ModifyWordSampleRequest struct {
	*tchttp.BaseRequest
	
	// 关键词，长度限制：128 个字符。
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// <b>关键词应用场景，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜的内容识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜的音频识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜的内容识别，等价于 3+4；
	// 7. All：包含以上全部，等价于 1+2+3+4。
	Usages []*string `json:"Usages,omitnil" name:"Usages"`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil" name:"TagOperationInfo"`
}

func (r *ModifyWordSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWordSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keyword")
	delete(f, "Usages")
	delete(f, "TagOperationInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWordSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWordSampleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWordSampleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWordSampleResponseParams `json:"Response"`
}

func (r *ModifyWordSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWordSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MosaicInput struct {
	// 原点位置，目前仅支持：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，马赛克原点为图片或文字的左上角。</li>
	// 默认值：TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil" name:"CoordinateOrigin"`

	// 马赛克原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示马赛克 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示马赛克 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil" name:"XPos"`

	// 马赛克原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示马赛克 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示马赛克 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil" name:"YPos"`

	// 马赛克的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示马赛克 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示马赛克 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 默认值：10%。
	Width *string `json:"Width,omitnil" name:"Width"`

	// 马赛克的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示马赛克 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示马赛克 Height 单位为像素，如 100px 表示 Height 为 100 像素。</li>
	// 默认值：10%。
	Height *string `json:"Height,omitnil" name:"Height"`

	// 马赛克的起始时间偏移，单位：秒。不填或填0，表示马赛克从画面出现时开始显现。
	// <li>不填或填0，表示马赛克从画面开始就出现；</li>
	// <li>当数值大于0时（假设为 n），表示马赛克从画面开始的第 n 秒出现；</li>
	// <li>当数值小于0时（假设为 -n），表示马赛克从离画面结束 n 秒前开始出现。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 马赛克的结束时间偏移，单位：秒。
	// <li>不填或填0，表示马赛克持续到画面结束；</li>
	// <li>当数值大于0时（假设为 n），表示马赛克持续到第 n 秒时消失；</li>
	// <li>当数值小于0时（假设为 -n），表示马赛克持续到离画面结束 n 秒前消失。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`
}

type NumberFormat struct {
	// `{number}`变量的起始值，默认为0。
	InitialValue *uint64 `json:"InitialValue,omitnil" name:"InitialValue"`

	// `{number}`变量的增长步长，默认为1。
	Increment *uint64 `json:"Increment,omitnil" name:"Increment"`

	// `{number}`变量的最小长度，不足时补占位符。默认为1。
	MinLength *uint64 `json:"MinLength,omitnil" name:"MinLength"`

	// `{number}`变量的长度不足时，补充的占位符。默认为"0"。
	PlaceHolder *string `json:"PlaceHolder,omitnil" name:"PlaceHolder"`
}

type OcrFullTextConfigureInfo struct {
	// 文本全文识别任务开关，可选值：
	// <li>ON：开启智能文本全文识别任务；</li>
	// <li>OFF：关闭智能文本全文识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type OcrFullTextConfigureInfoForUpdate struct {
	// 文本全文识别任务开关，可选值：
	// <li>ON：开启智能文本全文识别任务；</li>
	// <li>OFF：关闭智能文本全文识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type OcrWordsConfigureInfo struct {
	// 文本关键词识别任务开关，可选值：
	// <li>ON：开启文本关键词识别任务；</li>
	// <li>OFF：关闭文本关键词识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`
}

type OcrWordsConfigureInfoForUpdate struct {
	// 文本关键词识别任务开关，可选值：
	// <li>ON：开启文本关键词识别任务；</li>
	// <li>OFF：关闭文本关键词识别任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`
}

type OutputAddress struct {
	// 出口IP。
	Ip *string `json:"Ip,omitnil" name:"Ip"`
}

type OutputSRTSourceAddressResp struct {
	// 监听IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 监听端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type OverrideTranscodeParameter struct {
	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 是否去除视频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveVideo *uint64 `json:"RemoveVideo,omitnil" name:"RemoveVideo"`

	// 是否去除音频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveAudio *uint64 `json:"RemoveAudio,omitnil" name:"RemoveAudio"`

	// 视频流配置参数。
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitnil" name:"VideoTemplate"`

	// 音频流配置参数。
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitnil" name:"AudioTemplate"`

	// 极速高清转码参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitnil" name:"TEHDConfig"`

	// 字幕流配置参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubtitleTemplate *SubtitleTemplate `json:"SubtitleTemplate,omitnil" name:"SubtitleTemplate"`

	// 外挂音轨参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddonAudioStream []*MediaInputInfo `json:"AddonAudioStream,omitnil" name:"AddonAudioStream"`

	// 转码扩展字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StdExtInfo *string `json:"StdExtInfo,omitnil" name:"StdExtInfo"`

	// 要插入的字幕文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddOnSubtitles []*AddOnSubtitle `json:"AddOnSubtitles,omitnil" name:"AddOnSubtitles"`
}

// Predefined struct for user
type ParseLiveStreamProcessNotificationRequestParams struct {
	// 从 CMQ 获取到的直播流事件通知内容。
	Content *string `json:"Content,omitnil" name:"Content"`
}

type ParseLiveStreamProcessNotificationRequest struct {
	*tchttp.BaseRequest
	
	// 从 CMQ 获取到的直播流事件通知内容。
	Content *string `json:"Content,omitnil" name:"Content"`
}

func (r *ParseLiveStreamProcessNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseLiveStreamProcessNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseLiveStreamProcessNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseLiveStreamProcessNotificationResponseParams struct {
	// 直播流处理结果类型，包含：
	// <li>AiReviewResult：内容审核结果；</li>
	// <li>AiRecognitionResult：内容识别结果；</li>
	// <li>ProcessEof：直播流处理结束。</li>
	NotificationType *string `json:"NotificationType,omitnil" name:"NotificationType"`

	// 视频处理任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 直播流处理错误信息，当 NotificationType 为 ProcessEof 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessEofInfo *LiveStreamProcessErrorInfo `json:"ProcessEofInfo,omitnil" name:"ProcessEofInfo"`

	// 内容审核结果，当 NotificationType 为 AiReviewResult 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiReviewResultInfo *LiveStreamAiReviewResultInfo `json:"AiReviewResultInfo,omitnil" name:"AiReviewResultInfo"`

	// 内容识别结果，当 NotificationType 为 AiRecognitionResult 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiRecognitionResultInfo *LiveStreamAiRecognitionResultInfo `json:"AiRecognitionResultInfo,omitnil" name:"AiRecognitionResultInfo"`

	// 内容分析结果，当 NotificationType 为 AiAnalysisResult 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiAnalysisResultInfo *LiveStreamAiAnalysisResultInfo `json:"AiAnalysisResultInfo,omitnil" name:"AiAnalysisResultInfo"`

	// 媒体质检结果，当 NotificationType 为 AiQualityControlResult 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiQualityControlResultInfo *LiveStreamAiQualityControlResultInfo `json:"AiQualityControlResultInfo,omitnil" name:"AiQualityControlResultInfo"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长50个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长1000个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ParseLiveStreamProcessNotificationResponse struct {
	*tchttp.BaseResponse
	Response *ParseLiveStreamProcessNotificationResponseParams `json:"Response"`
}

func (r *ParseLiveStreamProcessNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseLiveStreamProcessNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseNotificationRequestParams struct {
	// 从 CMQ 获取到的事件通知内容。
	Content *string `json:"Content,omitnil" name:"Content"`
}

type ParseNotificationRequest struct {
	*tchttp.BaseRequest
	
	// 从 CMQ 获取到的事件通知内容。
	Content *string `json:"Content,omitnil" name:"Content"`
}

func (r *ParseNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseNotificationResponseParams struct {
	// 支持事件类型，目前取值有：
	// <li>WorkflowTask：视频工作流处理任务。</li>
	// <li>EditMediaTask：视频编辑任务。</li>
	// <li>ScheduleTask：编排任务。</li>
	EventType *string `json:"EventType,omitnil" name:"EventType"`

	// 视频处理任务信息，仅当 EventType 为 WorkflowTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowTaskEvent *WorkflowTask `json:"WorkflowTaskEvent,omitnil" name:"WorkflowTaskEvent"`

	// 视频编辑任务信息，仅当 EventType 为 EditMediaTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditMediaTaskEvent *EditMediaTask `json:"EditMediaTaskEvent,omitnil" name:"EditMediaTaskEvent"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长50个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长1000个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 编排任务信息，仅当 EventType 为 ScheduleTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleTaskEvent *ScheduleTask `json:"ScheduleTaskEvent,omitnil" name:"ScheduleTaskEvent"`

	// - 过期时间，事件通知签名过期 UNIX 时间戳。
	// - 来自媒体处理的消息通知默认过期时间是10分钟，如果一条消息通知中的 Timestamp 值所指定的时间已经过期，则可以判定这条通知无效，进而可以防止网络重放攻击。
	// - Timestamp 的格式为十进制 UNIX 时间戳，即从1970年01月01日（UTC/GMT 的午夜）开始所经过的秒数。
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 事件通知安全签名 Sign = MD5（Timestamp + NotifyKey）。说明：媒体处理把Timestamp 和 TaskNotifyConfig 里面的NotifyKey 进行字符串拼接后通过 MD5 计算得出 Sign 值，并将其放在通知消息里，您的后台服务器在收到通知消息后可以根据同样的算法确认 Sign 是否正确，进而确认消息是否确实来自媒体处理后台。
	Sign *string `json:"Sign,omitnil" name:"Sign"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ParseNotificationResponse struct {
	*tchttp.BaseResponse
	Response *ParseNotificationResponseParams `json:"Response"`
}

func (r *ParseNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PoliticalAsrReviewTemplateInfo struct {
	// 语音涉敏任务开关，可选值：
	// <li>ON：开启语音涉敏任务；</li>
	// <li>OFF：关闭语音涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PoliticalAsrReviewTemplateInfoForUpdate struct {
	// 语音涉敏任务开关，可选值：
	// <li>ON：开启语音涉敏任务；</li>
	// <li>OFF：关闭语音涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PoliticalConfigureInfo struct {
	// 画面涉敏控制参数。
	ImgReviewInfo *PoliticalImgReviewTemplateInfo `json:"ImgReviewInfo,omitnil" name:"ImgReviewInfo"`

	// 语音涉敏控制参数。
	AsrReviewInfo *PoliticalAsrReviewTemplateInfo `json:"AsrReviewInfo,omitnil" name:"AsrReviewInfo"`

	// 文本涉敏控制参数。
	OcrReviewInfo *PoliticalOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type PoliticalConfigureInfoForUpdate struct {
	// 画面涉敏控制参数。
	ImgReviewInfo *PoliticalImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitnil" name:"ImgReviewInfo"`

	// 语音涉敏控制参数。
	AsrReviewInfo *PoliticalAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil" name:"AsrReviewInfo"`

	// 文本涉敏控制参数。
	OcrReviewInfo *PoliticalOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type PoliticalImgReviewTemplateInfo struct {
	// 画面涉敏任务开关，可选值：
	// <li>ON：开启画面涉敏任务；</li>
	// <li>OFF：关闭画面涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 画面涉敏过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>violation_photo：违规图标；</li>
	// <li>politician：涉敏人物；</li>
	// <li>entertainment：娱乐人物；</li>
	// <li>sport：体育人物；</li>
	// <li>entrepreneur：商业人物；</li>
	// <li>scholar：教育学者；</li>
	// <li>celebrity：知名人物；</li>
	// <li>military：军事人物。</li>
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 97 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 95 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PoliticalImgReviewTemplateInfoForUpdate struct {
	// 画面涉敏任务开关，可选值：
	// <li>ON：开启画面涉敏任务；</li>
	// <li>OFF：关闭画面涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 画面涉敏过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>violation_photo：违规图标；</li>
	// <li>politician：涉敏人物；</li>
	// <li>entertainment：娱乐人物；</li>
	// <li>sport：体育人物；</li>
	// <li>entrepreneur：商业人物；</li>
	// <li>scholar：教育学者；</li>
	// <li>celebrity：知名人物；</li>
	// <li>military：军事人物。</li>
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfo struct {
	// 文本涉敏任务开关，可选值：
	// <li>ON：开启文本涉敏任务；</li>
	// <li>OFF：关闭文本涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfoForUpdate struct {
	// 文本涉敏任务开关，可选值：
	// <li>ON：开启文本涉敏任务；</li>
	// <li>OFF：关闭文本涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfo struct {
	// 语音鉴黄任务开关，可选值：
	// <li>ON：开启语音鉴黄任务；</li>
	// <li>OFF：关闭语音鉴黄任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfoForUpdate struct {
	// 语音鉴黄任务开关，可选值：
	// <li>ON：开启语音鉴黄任务；</li>
	// <li>OFF：关闭语音鉴黄任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PornConfigureInfo struct {
	// 画面鉴黄控制参数。
	ImgReviewInfo *PornImgReviewTemplateInfo `json:"ImgReviewInfo,omitnil" name:"ImgReviewInfo"`

	// 语音鉴黄控制参数。
	AsrReviewInfo *PornAsrReviewTemplateInfo `json:"AsrReviewInfo,omitnil" name:"AsrReviewInfo"`

	// 文本鉴黄控制参数。
	OcrReviewInfo *PornOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type PornConfigureInfoForUpdate struct {
	// 画面鉴黄控制参数。
	ImgReviewInfo *PornImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitnil" name:"ImgReviewInfo"`

	// 语音鉴黄控制参数。
	AsrReviewInfo *PornAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil" name:"AsrReviewInfo"`

	// 文本鉴黄控制参数。
	OcrReviewInfo *PornOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type PornImgReviewTemplateInfo struct {
	// 画面鉴黄任务开关，可选值：
	// <li>ON：开启画面鉴黄任务；</li>
	// <li>OFF：关闭画面鉴黄任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 画面鉴黄过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>porn：色情；</li>
	// <li>vulgar：低俗；</li>
	// <li>intimacy：亲密行为；</li>
	// <li>sexy：性感。</li>
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 90 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 0 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PornImgReviewTemplateInfoForUpdate struct {
	// 画面鉴黄任务开关，可选值：
	// <li>ON：开启画面鉴黄任务；</li>
	// <li>OFF：关闭画面鉴黄任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 画面鉴黄过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>porn：色情；</li>
	// <li>vulgar：低俗；</li>
	// <li>intimacy：亲密行为；</li>
	// <li>sexy：性感。</li>
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfo struct {
	// 文本鉴黄任务开关，可选值：
	// <li>ON：开启文本鉴黄任务；</li>
	// <li>OFF：关闭文本鉴黄任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfoForUpdate struct {
	// 文本鉴黄任务开关，可选值：
	// <li>ON：开启文本鉴黄任务；</li>
	// <li>OFF：关闭文本鉴黄任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

// Predefined struct for user
type ProcessLiveStreamRequestParams struct {
	// 直播流 URL（必须是直播文件地址，支持 rtmp，hls 和 flv 等）。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 任务的事件通知信息，用于指定直播流处理的结果。
	TaskNotifyConfig *LiveStreamTaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 直播流处理输出文件的目标存储。如处理有文件输出，该参数为必填项。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 直播流处理生成的文件输出的目标目录，如`/movie/201909/`，如果不填为 `/` 目录。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容质检类型任务参数。
	AiQualityControlTask *AiQualityControlTaskInput `json:"AiQualityControlTask,omitnil" name:"AiQualityControlTask"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 直播编排ID。
	// 注意1：对于OutputStorage、OutputDir参数：
	// <li>当服务编排中子任务节点配置了OutputStorage、OutputDir时，该子任务节点中配置的输出作为子任务的输出。</li>
	// <li>当服务编排中子任务节点没有配置OutputStorage、OutputDir时，若对直播流发起处理（ProcessLiveStream）有输出，将覆盖原有编排的默认输出。</li>
	// 注意2：对于TaskNotifyConfig参数，若创建任务接口（ProcessLiveStream）有设置，将覆盖原有编排的默认回调。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`
}

type ProcessLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 直播流 URL（必须是直播文件地址，支持 rtmp，hls 和 flv 等）。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 任务的事件通知信息，用于指定直播流处理的结果。
	TaskNotifyConfig *LiveStreamTaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 直播流处理输出文件的目标存储。如处理有文件输出，该参数为必填项。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 直播流处理生成的文件输出的目标目录，如`/movie/201909/`，如果不填为 `/` 目录。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容质检类型任务参数。
	AiQualityControlTask *AiQualityControlTaskInput `json:"AiQualityControlTask,omitnil" name:"AiQualityControlTask"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 直播编排ID。
	// 注意1：对于OutputStorage、OutputDir参数：
	// <li>当服务编排中子任务节点配置了OutputStorage、OutputDir时，该子任务节点中配置的输出作为子任务的输出。</li>
	// <li>当服务编排中子任务节点没有配置OutputStorage、OutputDir时，若对直播流发起处理（ProcessLiveStream）有输出，将覆盖原有编排的默认输出。</li>
	// 注意2：对于TaskNotifyConfig参数，若创建任务接口（ProcessLiveStream）有设置，将覆盖原有编排的默认回调。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`
}

func (r *ProcessLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "TaskNotifyConfig")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "AiContentReviewTask")
	delete(f, "AiRecognitionTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiQualityControlTask")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "ScheduleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessLiveStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessLiveStreamResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ProcessLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *ProcessLiveStreamResponseParams `json:"Response"`
}

func (r *ProcessLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaRequestParams struct {
	// 媒体处理的文件输入信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 媒体处理输出文件的目标存储。不填则继承 InputInfo 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理生成的文件输出的目标目录，必选以 / 开头和结尾，如`/movie/201907/`。
	// 如果不填，表示与 InputInfo 中文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 编排ID。
	// 注意1：对于OutputStorage、OutputDir参数：
	// <li>当服务编排中子任务节点配置了OutputStorage、OutputDir时，该子任务节点中配置的输出作为子任务的输出。</li>
	// <li>当服务编排中子任务节点没有配置OutputStorage、OutputDir时，若创建任务接口（ProcessMedia）有输出，将覆盖原有编排的默认输出。</li>
	// 注意2：对于TaskNotifyConfig参数，若创建任务接口（ProcessMedia）有设置，将覆盖原有编排的默认回调。
	// 
	// 注意3：编排的 Trigger 只是用来自动化触发场景，在手动发起的请求中已经配置的 Trigger 无意义。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`

	// 媒体处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`

	// 视频质检类型任务参数。
	AiQualityControlTask *AiQualityControlTaskInput `json:"AiQualityControlTask,omitnil" name:"AiQualityControlTask"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 任务流的优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil" name:"TasksPriority"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 任务类型，默认Online
	// <li> Online：实时任务</li>
	// <li> Offline：闲时任务，不保证实效性，默认3天内处理完</li>
	TaskType *string `json:"TaskType,omitnil" name:"TaskType"`
}

type ProcessMediaRequest struct {
	*tchttp.BaseRequest
	
	// 媒体处理的文件输入信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 媒体处理输出文件的目标存储。不填则继承 InputInfo 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理生成的文件输出的目标目录，必选以 / 开头和结尾，如`/movie/201907/`。
	// 如果不填，表示与 InputInfo 中文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 编排ID。
	// 注意1：对于OutputStorage、OutputDir参数：
	// <li>当服务编排中子任务节点配置了OutputStorage、OutputDir时，该子任务节点中配置的输出作为子任务的输出。</li>
	// <li>当服务编排中子任务节点没有配置OutputStorage、OutputDir时，若创建任务接口（ProcessMedia）有输出，将覆盖原有编排的默认输出。</li>
	// 注意2：对于TaskNotifyConfig参数，若创建任务接口（ProcessMedia）有设置，将覆盖原有编排的默认回调。
	// 
	// 注意3：编排的 Trigger 只是用来自动化触发场景，在手动发起的请求中已经配置的 Trigger 无意义。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`

	// 媒体处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`

	// 视频质检类型任务参数。
	AiQualityControlTask *AiQualityControlTaskInput `json:"AiQualityControlTask,omitnil" name:"AiQualityControlTask"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 任务流的优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil" name:"TasksPriority"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 任务类型，默认Online
	// <li> Online：实时任务</li>
	// <li> Offline：闲时任务，不保证实效性，默认3天内处理完</li>
	TaskType *string `json:"TaskType,omitnil" name:"TaskType"`
}

func (r *ProcessMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "ScheduleId")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "AiQualityControlTask")
	delete(f, "TaskNotifyConfig")
	delete(f, "TasksPriority")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaResponseParams struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ProcessMediaResponse struct {
	*tchttp.BaseResponse
	Response *ProcessMediaResponseParams `json:"Response"`
}

func (r *ProcessMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProhibitedAsrReviewTemplateInfo struct {
	// 语音违禁任务开关，可选值：
	// <li>ON：开启语音违禁任务；</li>
	// <li>OFF：关闭语音违禁任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type ProhibitedAsrReviewTemplateInfoForUpdate struct {
	// 语音违禁任务开关，可选值：
	// <li>ON：开启语音违禁任务；</li>
	// <li>OFF：关闭语音违禁任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type ProhibitedConfigureInfo struct {
	// 语音违禁控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrReviewInfo *ProhibitedAsrReviewTemplateInfo `json:"AsrReviewInfo,omitnil" name:"AsrReviewInfo"`

	// 文本违禁控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrReviewInfo *ProhibitedOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type ProhibitedConfigureInfoForUpdate struct {
	// 语音违禁控制参数。
	AsrReviewInfo *ProhibitedAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil" name:"AsrReviewInfo"`

	// 文本违禁控制参数。
	OcrReviewInfo *ProhibitedOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type ProhibitedOcrReviewTemplateInfo struct {
	// 文本违禁任务开关，可选值：
	// <li>ON：开启文本违禁任务；</li>
	// <li>OFF：关闭文本违禁任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type ProhibitedOcrReviewTemplateInfoForUpdate struct {
	// 文本违禁任务开关，可选值：
	// <li>ON：开启文本违禁任务；</li>
	// <li>OFF：关闭文本违禁任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type QualityControlData struct {
	// 为true时表示视频无音频轨。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoAudio *bool `json:"NoAudio,omitnil" name:"NoAudio"`

	// 为true时表示视频无视频轨。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoVideo *bool `json:"NoVideo,omitnil" name:"NoVideo"`

	// 视频无参考质量打分，百分制。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityEvaluationScore *int64 `json:"QualityEvaluationScore,omitnil" name:"QualityEvaluationScore"`

	// 质检检出异常项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityControlResultSet []*QualityControlResult `json:"QualityControlResultSet,omitnil" name:"QualityControlResultSet"`
}

type QualityControlItem struct {
	// 置信度，取值范围是 0 到 100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 出现的起始时间戳，秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 出现的结束时间戳，秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 区域坐标(px)，即左上角坐标、右下角坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil" name:"AreaCoordSet"`
}

type QualityControlResult struct {
	// 异常类型，取值范围：
	// Jitter：抖动，
	// Blur：模糊，
	// LowLighting：低光照，
	// HighLighting：过曝，
	// CrashScreen：花屏，
	// BlackWhiteEdge：黑白边，
	// SolidColorScreen：纯色屏，
	// Noise：噪点，
	// Mosaic：马赛克，
	// QRCode：二维码，
	// AppletCode：小程序码，
	// BarCode：条形码，
	// LowVoice：低音，
	// HighVoice：爆音，
	// NoVoice：静音，
	// LowEvaluation：无参考打分低于阈值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 质检结果项。
	QualityControlItems []*QualityControlItem `json:"QualityControlItems,omitnil" name:"QualityControlItems"`
}

type RTMPAddressDestination struct {
	// 转推RTMP的目标Url，格式如'rtmp://domain/live'。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 转推RTMP的目标StreamKey，格式如'steamid?key=value'。
	StreamKey *string `json:"StreamKey,omitnil" name:"StreamKey"`
}

type RTMPPullSourceAddress struct {
	// RTMP源站的TcUrl地址。
	TcUrl *string `json:"TcUrl,omitnil" name:"TcUrl"`

	// RTMP源站的StreamKey信息。
	StreamKey *string `json:"StreamKey,omitnil" name:"StreamKey"`
}

type RTPAddressDestination struct {
	// 转推的目标地址的IP。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 转推的目标地址的端口。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type RTSPPullSourceAddress struct {
	// RTSP源站的Url地址。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type RawImageWatermarkInput struct {
	// 水印图片的输入内容。支持 jpeg、png 图片格式。
	ImageContent *MediaInputInfo `json:"ImageContent,omitnil" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 默认值：10%。
	Width *string `json:"Width,omitnil" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素。</li>
	// 默认值：0px，表示 Height 按照原始水印图片的宽高比缩放。
	Height *string `json:"Height,omitnil" name:"Height"`

	// 水印重复类型。使用场景：水印为动态图像。取值范围：
	// <li>once：动态水印播放完后，不再出现；</li>
	// <li>repeat_last_frame：水印播放完后，停留在最后一帧；</li>
	// <li>repeat：水印循环播放，直到视频结束（默认值）。</li>
	RepeatType *string `json:"RepeatType,omitnil" name:"RepeatType"`
}

type RawTranscodeParameter struct {
	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 是否去除视频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil" name:"RemoveVideo"`

	// 是否去除音频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil" name:"RemoveAudio"`

	// 视频流配置参数，当 RemoveVideo 为 0，该字段必填。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil" name:"VideoTemplate"`

	// 音频流配置参数，当 RemoveAudio 为 0，该字段必填。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil" name:"AudioTemplate"`

	// 极速高清转码参数。
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil" name:"TEHDConfig"`
}

type RawWatermarkParameter struct {
	// 水印类型，可选值：
	// <li>image：图片水印。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 原点位置，目前仅支持：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角。</li>
	// 默认值：TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil" name:"YPos"`

	// 图片水印模板，当 Type 为 image，该字段必填。当 Type 为 text，该字段无效。
	ImageTemplate *RawImageWatermarkInput `json:"ImageTemplate,omitnil" name:"ImageTemplate"`
}

// Predefined struct for user
type RecognizeMediaForZhiXueRequestParams struct {
	// 输入媒体文件存储信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 表情识别参数配置。默认开启。
	ExpressionConfig *ExpressionConfigInfo `json:"ExpressionConfig,omitnil" name:"ExpressionConfig"`

	// 动作识别参数配置。默认开启。
	ActionConfig *ActionConfigInfo `json:"ActionConfig,omitnil" name:"ActionConfig"`
}

type RecognizeMediaForZhiXueRequest struct {
	*tchttp.BaseRequest
	
	// 输入媒体文件存储信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 表情识别参数配置。默认开启。
	ExpressionConfig *ExpressionConfigInfo `json:"ExpressionConfig,omitnil" name:"ExpressionConfig"`

	// 动作识别参数配置。默认开启。
	ActionConfig *ActionConfigInfo `json:"ActionConfig,omitnil" name:"ActionConfig"`
}

func (r *RecognizeMediaForZhiXueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMediaForZhiXueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	delete(f, "ExpressionConfig")
	delete(f, "ActionConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeMediaForZhiXueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeMediaForZhiXueResponseParams struct {
	// 任务 ID，可以通过该 ID 查询任务状态和结果。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeMediaForZhiXueResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeMediaForZhiXueResponseParams `json:"Response"`
}

func (r *RecognizeMediaForZhiXueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMediaForZhiXueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地区名称。
	Name *string `json:"Name,omitnil" name:"Name"`
}

// Predefined struct for user
type ResetWorkflowRequestParams struct {
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`

	// 工作流名称，最多128字符。同一个用户该名称唯一。
	WorkflowName *string `json:"WorkflowName,omitnil" name:"WorkflowName"`

	// 工作流绑定的触发规则，当上传视频命中该规则到该对象时即触发工作流。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 视频处理的文件输出配置。不填则继承 Trigger 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 视频处理生成的文件输出的目标目录，如`/movie/201907/`。如果不填，表示与触发文件所在的目录一致，即`{inputDir}`。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`

	// 工作流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TaskPriority *int64 `json:"TaskPriority,omitnil" name:"TaskPriority"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`
}

type ResetWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`

	// 工作流名称，最多128字符。同一个用户该名称唯一。
	WorkflowName *string `json:"WorkflowName,omitnil" name:"WorkflowName"`

	// 工作流绑定的触发规则，当上传视频命中该规则到该对象时即触发工作流。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 视频处理的文件输出配置。不填则继承 Trigger 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 视频处理生成的文件输出的目标目录，如`/movie/201907/`。如果不填，表示与触发文件所在的目录一致，即`{inputDir}`。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`

	// 工作流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TaskPriority *int64 `json:"TaskPriority,omitnil" name:"TaskPriority"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`
}

func (r *ResetWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	delete(f, "WorkflowName")
	delete(f, "Trigger")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "TaskPriority")
	delete(f, "TaskNotifyConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetWorkflowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *ResetWorkflowResponseParams `json:"Response"`
}

func (r *ResetWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResilientStreamConf struct {
	// 是否开启延播平滑吐流，true开启，false不开启，默认不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 延播时间，单位秒，目前支持的范围为10~300秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BufferTime *uint64 `json:"BufferTime,omitnil" name:"BufferTime"`
}

type S3InputInfo struct {
	// S3 bucket。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3Bucket *string `json:"S3Bucket,omitnil" name:"S3Bucket"`

	// S3 bucket 对应的区域，目前支持：  
	// us-east-1  
	// eu-west-3
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3Region *string `json:"S3Region,omitnil" name:"S3Region"`

	// S3 bucket 中的媒体资源路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3Object *string `json:"S3Object,omitnil" name:"S3Object"`

	// AWS 内网访问 媒体资源的秘钥id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3SecretId *string `json:"S3SecretId,omitnil" name:"S3SecretId"`

	// AWS 内网访问 媒体资源的秘钥key。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3SecretKey *string `json:"S3SecretKey,omitnil" name:"S3SecretKey"`
}

type S3OutputStorage struct {
	// S3 bucket。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3Bucket *string `json:"S3Bucket,omitnil" name:"S3Bucket"`

	// S3 bucket 对应的区域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3Region *string `json:"S3Region,omitnil" name:"S3Region"`

	// AWS 内网上传 媒体资源的秘钥id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3SecretId *string `json:"S3SecretId,omitnil" name:"S3SecretId"`

	// AWS 内网上传 媒体资源的秘钥key。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3SecretKey *string `json:"S3SecretKey,omitnil" name:"S3SecretKey"`
}

type SRTAddressDestination struct {
	// 目标地址的IP。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 目标地址的端口。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type SRTSourceAddressReq struct {
	// 对端IP。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 对端端口。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type SRTSourceAddressResp struct {
	// 对端IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 对端端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type SampleSnapshotTaskInput struct {
	// 采样截图模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil" name:"WatermarkSet"`

	// 采样截图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 采样截图后图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_sampleSnapshot_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil" name:"OutputObjectPath"`

	// 采样截图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitnil" name:"ObjectNumberFormat"`
}

type SampleSnapshotTemplate struct {
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 采样截图模板名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 图片格式。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 采样截图类型。
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 采样间隔。
	SampleInterval *uint64 `json:"SampleInterval,omitnil" name:"SampleInterval"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

type ScheduleAnalysisTaskResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 分析任务的输入。
	Input *AiAnalysisTaskInput `json:"Input,omitnil" name:"Input"`

	// 分析任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output []*AiAnalysisResult `json:"Output,omitnil" name:"Output"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`
}

type ScheduleQualityControlTaskResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 质检任务的输入。
	Input *AiQualityControlTaskInput `json:"Input,omitnil" name:"Input"`

	// 质检任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *QualityControlData `json:"Output,omitnil" name:"Output"`
}

type ScheduleRecognitionTaskResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 识别任务的输入。
	Input *AiRecognitionTaskInput `json:"Input,omitnil" name:"Input"`

	// 识别任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output []*AiRecognitionResult `json:"Output,omitnil" name:"Output"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`
}

type ScheduleReviewTaskResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [媒体处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 审核任务的输入。
	Input *AiContentReviewTaskInput `json:"Input,omitnil" name:"Input"`

	// 审核任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output []*AiContentReviewResult `json:"Output,omitnil" name:"Output"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`
}

type ScheduleTask struct {
	// 编排任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 源异常时返回非0错误码，返回0 时请使用各个具体任务的 ErrCode。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 源异常时返回对应异常Message，否则请使用各个具体任务的 Message。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 媒体处理的目标文件信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 原始视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil" name:"MetaData"`

	// 编排任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityResultSet []*ActivityResult `json:"ActivityResultSet,omitnil" name:"ActivityResultSet"`
}

type SchedulesInfo struct {
	// 编排唯一标识。
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`

	// 编排名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleName *string `json:"ScheduleName,omitnil" name:"ScheduleName"`

	// 编排类型，可选值：
	//  <li>Preset：系统预置编排；</li>
	// <li>Custom：用户自定义编排。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 编排状态，取值范围：
	// Enabled：已启用，
	// Disabled：已禁用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 编排绑定的触发规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 编排任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Activities []*Activity `json:"Activities,omitnil" name:"Activities"`

	// 媒体处理的文件输出存储位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理生成的文件输出的目标目录。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 任务的事件通知配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最后编辑时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type ScratchRepairConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 强度，取值范围：0.0~1.0。
	// 默认：0.0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intensity *float64 `json:"Intensity,omitnil" name:"Intensity"`
}

type SegmentRecognitionItem struct {
	// 置信度。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 片段起始时间偏移。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 片段结束时间偏移。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 拆条片段URL。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentUrl *string `json:"SegmentUrl,omitnil" name:"SegmentUrl"`
}

type SharpEnhanceConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 强度，取值范围：0.0~1.0。
	// 默认：0.0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intensity *float64 `json:"Intensity,omitnil" name:"Intensity"`
}

type SimpleAesDrm struct {
	// 请求解密秘钥uri地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uri *string `json:"Uri,omitnil" name:"Uri"`

	// 加密key(32字节字符串)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 加密初始化向量(32字节字符串)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vector *string `json:"Vector,omitnil" name:"Vector"`
}

type SnapshotByTimeOffsetTaskInput struct {
	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 截图时间点列表，时间点支持 s、% 两种格式：
	// <li>当字符串以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// <li>当字符串以 % 结尾，表示时间点为视频时长的百分比大小，如10%表示时间点为视频前第10%的时间。</li>
	ExtTimeOffsetSet []*string `json:"ExtTimeOffsetSet,omitnil" name:"ExtTimeOffsetSet"`

	// 截图时间点列表，单位为<font color=red>秒</font>。此参数已不再建议使用，建议您使用 ExtTimeOffsetSet 参数。
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitnil" name:"TimeOffsetSet"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil" name:"WatermarkSet"`

	// 时间点截图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 时间点截图后图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_snapshotByTimeOffset_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil" name:"OutputObjectPath"`

	// 时间点截图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitnil" name:"ObjectNumberFormat"`
}

type SnapshotByTimeOffsetTemplate struct {
	// 时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 时间点截图模板名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 图片格式。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>black：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>black：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil" name:"FillType"`
}

// Predefined struct for user
type StartStreamLinkFlowRequestParams struct {
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
}

type StartStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
}

func (r *StartStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStreamLinkFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *StartStreamLinkFlowResponseParams `json:"Response"`
}

func (r *StartStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamLinkFlowRequestParams struct {
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
}

type StopStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// 流Id。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
}

func (r *StopStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamLinkFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *StopStreamLinkFlowResponseParams `json:"Response"`
}

func (r *StopStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StreamLinkRegionInfo struct {
	// 媒体直传输的地区信息列表。
	Regions []*RegionInfo `json:"Regions,omitnil" name:"Regions"`
}

type SubtitleTemplate struct {
	// 要压制到视频中的字幕文件地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 指定要压制到视频中的字幕轨道，如果有指定Path，则Path 优先级更高。Path 和 StreamIndex 至少指定一个。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamIndex *int64 `json:"StreamIndex,omitnil" name:"StreamIndex"`

	// 字体类型，
	// <li>hei.ttf：黑体</li>
	// <li>song.ttf：宋体</li>
	// <li>simkai.ttf：楷体</li>
	// <li>arial.ttf：仅支持英文</li>
	// 默认hei.ttf
	// 注意：此字段可能返回 null，表示取不到有效值。
	FontType *string `json:"FontType,omitnil" name:"FontType"`

	// 字体大小，格式：Npx，N 为数值，不指定则以字幕文件中为准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FontSize *string `json:"FontSize,omitnil" name:"FontSize"`

	// 字体颜色，格式：0xRRGGBB，默认值：0xFFFFFF（白色）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FontColor *string `json:"FontColor,omitnil" name:"FontColor"`

	// 文字透明度，取值范围：(0, 1]
	// <li>0：完全透明</li>
	// <li>1：完全不透明</li>
	// 默认值：1。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FontAlpha *float64 `json:"FontAlpha,omitnil" name:"FontAlpha"`
}

type SuperResolutionConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 类型，可选值：
	// <li>lq：针对低清晰度有较多噪声视频的超分；</li>
	// <li>hq：针对高清晰度视频超分。</li>
	// 默认值：lq。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 超分倍数，可选值：
	// <li>2：目前只支持 2 倍超分。</li>
	// 默认值：2。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil" name:"Size"`
}

type SvgWatermarkInput struct {
	// 水印的宽度，支持 px，%，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素；当填 0px 且
	//  Height 不为 0px 时，表示水印的宽度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的宽度取原始 SVG 图像的宽度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Width 为视频宽度的百分比大小，如 10W% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Width 为视频高度的百分比大小，如 10H% 表示 Width 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Width 为视频短边的百分比大小，如 10S% 表示 Width 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Width 为视频长边的百分比大小，如 10L% 表示 Width 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 W%。</li>
	// 默认值为 10W%。
	Width *string `json:"Width,omitnil" name:"Width"`

	// 水印的高度，支持 px，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素；当填 0px 且
	//  Width 不为 0px 时，表示水印的高度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的高度取原始 SVG 图像的高度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Height 为视频宽度的百分比大小，如 10W% 表示 Height 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Height 为视频高度的百分比大小，如 10H% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Height 为视频短边的百分比大小，如 10S% 表示 Height 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Height 为视频长边的百分比大小，如 10L% 表示 Height 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 H%。</li>
	// 默认值为 0px。
	Height *string `json:"Height,omitnil" name:"Height"`
}

type SvgWatermarkInputForUpdate struct {
	// 水印的宽度，支持 px，%，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素；当填 0px 且
	//  Height 不为 0px 时，表示水印的宽度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的宽度取原始 SVG 图像的宽度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Width 为视频宽度的百分比大小，如 10W% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Width 为视频高度的百分比大小，如 10H% 表示 Width 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Width 为视频短边的百分比大小，如 10S% 表示 Width 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Width 为视频长边的百分比大小，如 10L% 表示 Width 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 W%。</li>
	// 默认值为 10W%。
	Width *string `json:"Width,omitnil" name:"Width"`

	// 水印的高度，支持 px，%，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素；当填 0px 且
	//  Width 不为 0px 时，表示水印的高度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的高度取原始 SVG 图像的高度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Height 为视频宽度的百分比大小，如 10W% 表示 Height 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Height 为视频高度的百分比大小，如 10H% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Height 为视频短边的百分比大小，如 10S% 表示 Height 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Height 为视频长边的百分比大小，如 10L% 表示 Height 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 H%。
	// 默认值为 0px。
	Height *string `json:"Height,omitnil" name:"Height"`
}

type TEHDConfig struct {
	// 极速高清类型，可选值：
	// <li>TEHD-100：极速高清-100（视频极速高清）。</li>
	// <li>TEHD-200：极速高清-200（音频极速高清）。</li>
	// 不填代表不启用极速高清。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 视频码率上限，当 Type 指定了极速高清类型时有效。
	// 不填或填0表示不设视频码率上限。
	MaxVideoBitrate *int64 `json:"MaxVideoBitrate,omitnil" name:"MaxVideoBitrate"`
}

type TEHDConfigForUpdate struct {
	// 极速高清类型，可选值：
	// <li>TEHD-100：极速高清-100（视频极速高清）。</li>
	// <li>TEHD-200：极速高清-200（音频极速高清）。</li>
	// 不填代表不修改。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 视频码率上限，不填代表不修改。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxVideoBitrate *int64 `json:"MaxVideoBitrate,omitnil" name:"MaxVideoBitrate"`
}

type TagConfigureInfo struct {
	// 智能标签任务开关，可选值：
	// <li>ON：开启智能标签任务；</li>
	// <li>OFF：关闭智能标签任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type TagConfigureInfoForUpdate struct {
	// 智能标签任务开关，可选值：
	// <li>ON：开启智能标签任务；</li>
	// <li>OFF：关闭智能标签任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type TaskNotifyConfig struct {
	// 通知类型，可选值：
	// <li>CMQ：已下线，建议切换到TDMQ-CMQ</li>
	// <li>TDMQ-CMQ：消息队列</li>
	// <li>URL：指定URL时HTTP回调推送到 NotifyUrl 指定的地址，回调协议http+json，包体内容同解析事件通知接口的输出参数 </li>
	// <li>SCF：不推荐使用，需要在控制台额外配置SCF</li>
	// <li>AWS-SQS：AWS 队列，只适用于 AWS 任务，且要求同区域</li>
	// <font color="red"> 注：不填或为空时默认 CMQ，如需采用其他类型需填写对应类型值。 </font>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 工作流通知的模式，可取值有 Finish 和 Change，不填代表 Finish。
	NotifyMode *string `json:"NotifyMode,omitnil" name:"NotifyMode"`

	// HTTP回调地址，NotifyType为URL时必填。
	NotifyUrl *string `json:"NotifyUrl,omitnil" name:"NotifyUrl"`

	// CMQ或TDMQ-CMQ 的模型，有 Queue 和 Topic 两种。
	CmqModel *string `json:"CmqModel,omitnil" name:"CmqModel"`

	// CMQ或TDMQ-CMQ 的园区，如 sh，bj 等。
	CmqRegion *string `json:"CmqRegion,omitnil" name:"CmqRegion"`

	// 当模型为 Topic 时有效，表示接收事件通知的 CMQ 或 TDMQ-CMQ 的主题名。
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 当模型为 Queue 时有效，表示接收事件通知的 CMQ 或 TDMQ-CMQ 的队列名。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// AWS SQS 回调，NotifyType为 AWS-SQS 时必填。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AwsSQS *AwsSQS `json:"AwsSQS,omitnil" name:"AwsSQS"`

	// 用于生成回调签名的key。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyKey *string `json:"NotifyKey,omitnil" name:"NotifyKey"`
}

type TaskOutputStorage struct {
	// 媒体处理输出对象存储位置的类型，支持：
	// <li>COS：COS存储</li>
	// <li>AWS-S3：AWS 存储，只适用于AWS任务，且要求同区域</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 当 Type 为 COS 时有效，则该项为必填，表示媒体处理 COS 输出位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosOutputStorage *CosOutputStorage `json:"CosOutputStorage,omitnil" name:"CosOutputStorage"`

	// 当 Type 为 AWS-S3 时有效，则该项为必填，表示媒体处理 AWS S3 输出位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3OutputStorage *S3OutputStorage `json:"S3OutputStorage,omitnil" name:"S3OutputStorage"`
}

type TaskSimpleInfo struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务类型，包含：
	// <li> WorkflowTask：工作流处理任务；</li>
	// <li> EditMediaTask：视频编辑任务；</li>
	// <li> LiveProcessTask：直播处理任务。</li>
	TaskType *string `json:"TaskType,omitnil" name:"TaskType"`

	// 任务创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 任务开始执行时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。若任务尚未开始，该字段为：0000-00-00T00:00:00Z。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil" name:"BeginProcessTime"`

	// 任务结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。若任务尚未完成，该字段为：0000-00-00T00:00:00Z。
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`

	// 子任务类型。
	SubTaskTypes []*string `json:"SubTaskTypes,omitnil" name:"SubTaskTypes"`
}

type TerrorismConfigureInfo struct {
	// 画面涉敏任务控制参数。
	ImgReviewInfo *TerrorismImgReviewTemplateInfo `json:"ImgReviewInfo,omitnil" name:"ImgReviewInfo"`

	// 文本涉敏任务控制参数。
	OcrReviewInfo *TerrorismOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type TerrorismConfigureInfoForUpdate struct {
	// 画面涉敏任务控制参数。
	ImgReviewInfo *TerrorismImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitnil" name:"ImgReviewInfo"`

	// 文本涉敏任务控制参数。
	OcrReviewInfo *TerrorismOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type TerrorismImgReviewTemplateInfo struct {
	// 画面涉敏任务开关，可选值：
	// <li>ON：开启画面涉敏任务；</li>
	// <li>OFF：关闭画面涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 画面涉敏过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>guns：武器枪支；</li>
	// <li>crowd：人群聚集；</li>
	// <li>bloody：血腥画面；</li>
	// <li>police：警察部队；</li>
	// <li>banners：涉敏旗帜；</li>
	// <li>militant：武装分子；</li>
	// <li>explosion：爆炸火灾；</li>
	// <li>terrorists：涉敏人物；</li>
	// <li>scenario：涉敏画面。</li>
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 90 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 80 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type TerrorismImgReviewTemplateInfoForUpdate struct {
	// 画面涉敏任务开关，可选值：
	// <li>ON：开启画面涉敏任务；</li>
	// <li>OFF：关闭画面涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 画面涉敏过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>guns：武器枪支；</li>
	// <li>crowd：人群聚集；</li>
	// <li>bloody：血腥画面；</li>
	// <li>police：警察部队；</li>
	// <li>banners：涉敏旗帜；</li>
	// <li>militant：武装分子；</li>
	// <li>explosion：爆炸火灾；</li>
	// <li>terrorists：涉敏人物；</li>
	// <li>scenario：涉敏画面。</li>
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type TerrorismOcrReviewTemplateInfo struct {
	// 文本涉敏任务开关，可选值：
	// <li>ON：开启文本涉敏任务；</li>
	// <li>OFF：关闭文本涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type TerrorismOcrReviewTemplateInfoForUpdate struct {
	// 文本涉敏任务开关，可选值：
	// <li>ON：开启文本涉敏任务；</li>
	// <li>OFF：关闭文本涉敏任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type TextWatermarkTemplateInput struct {
	// 字体类型，目前可以支持两种：
	// <li>simkai.ttf：可以支持中文和英文；</li>
	// <li>arial.ttf：仅支持英文。</li>
	FontType *string `json:"FontType,omitnil" name:"FontType"`

	// 字体大小，格式：Npx，N 为数值。
	FontSize *string `json:"FontSize,omitnil" name:"FontSize"`

	// 字体颜色，格式：0xRRGGBB，默认值：0xFFFFFF（白色）。
	FontColor *string `json:"FontColor,omitnil" name:"FontColor"`

	// 文字透明度，取值范围：(0, 1]
	// <li>0：完全透明</li>
	// <li>1：完全不透明</li>
	// 默认值：1。
	FontAlpha *float64 `json:"FontAlpha,omitnil" name:"FontAlpha"`
}

type TextWatermarkTemplateInputForUpdate struct {
	// 字体类型，目前可以支持两种：
	// <li>simkai.ttf：可以支持中文和英文；</li>
	// <li>arial.ttf：仅支持英文。</li>
	FontType *string `json:"FontType,omitnil" name:"FontType"`

	// 字体大小，格式：Npx，N 为数值。
	FontSize *string `json:"FontSize,omitnil" name:"FontSize"`

	// 字体颜色，格式：0xRRGGBB，默认值：0xFFFFFF（白色）。
	FontColor *string `json:"FontColor,omitnil" name:"FontColor"`

	// 文字透明度，取值范围：(0, 1]
	// <li>0：完全透明</li>
	// <li>1：完全不透明</li>
	FontAlpha *float64 `json:"FontAlpha,omitnil" name:"FontAlpha"`
}

type TranscodeTaskInput struct {
	// 视频转码模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 视频转码自定义参数，当 Definition 填 0 时有效。
	// 该参数用于高度定制场景，建议您优先使用 Definition 指定转码参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawParameter *RawTranscodeParameter `json:"RawParameter,omitnil" name:"RawParameter"`

	// 视频转码自定义参数，当 Definition 不填 0 时有效。
	// 当填写了该结构中的部分转码参数时，将使用填写的参数覆盖转码模板中的参数。
	// 该参数用于高度定制场景，建议您仅使用 Definition 指定转码参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverrideParameter *OverrideTranscodeParameter `json:"OverrideParameter,omitnil" name:"OverrideParameter"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil" name:"WatermarkSet"`

	// 马赛克列表，最大可支持 10 张。
	MosaicSet []*MosaicInput `json:"MosaicSet,omitnil" name:"MosaicSet"`

	// 转码后的视频的起始时间偏移，单位：秒。
	// <li>不填或填0，表示转码后的视频从原始视频的起始位置开始；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频从原始视频的第 n 秒位置开始；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频从原始视频结束 n 秒前的位置开始。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 转码后视频的终止时间偏移，单位：秒。
	// <li>不填或填0，表示转码后的视频持续到原始视频的末尾终止；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频持续到原始视频第 n 秒时终止；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频持续到原始视频结束 n 秒前终止。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`

	// 转码后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 转码后主文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_transcode_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil" name:"OutputObjectPath"`

	// 转码后分片文件的输出路径（转码 HLS 时 ts 的路径），只能为相对路径。如果不填，则默认为：`{inputName}_transcode_{definition}_{number}.{format}`。
	SegmentObjectName *string `json:"SegmentObjectName,omitnil" name:"SegmentObjectName"`

	// 转码后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitnil" name:"ObjectNumberFormat"`

	// 片头片尾参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadTailParameter *HeadTailParameter `json:"HeadTailParameter,omitnil" name:"HeadTailParameter"`
}

type TranscodeTemplate struct {
	// 转码模板唯一标识。
	Definition *string `json:"Definition,omitnil" name:"Definition"`

	// 封装格式，取值：mp4、flv、hls、mp3、flac、ogg。
	Container *string `json:"Container,omitnil" name:"Container"`

	// 转码模板名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 模板类型，取值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 是否去除视频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil" name:"RemoveVideo"`

	// 是否去除音频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil" name:"RemoveAudio"`

	// 视频流配置参数，仅当 RemoveVideo 为 0，该字段有效。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil" name:"VideoTemplate"`

	// 音频流配置参数，仅当 RemoveAudio 为 0，该字段有效 。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil" name:"AudioTemplate"`

	// 极速高清转码参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil" name:"TEHDConfig"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式板。</li>
	ContainerType *string `json:"ContainerType,omitnil" name:"ContainerType"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 音视频增强配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil" name:"EnhanceConfig"`
}

type TranslateConfigureInfo struct {
	// 语音翻译任务开关，可选值：
	// <li>ON：开启智能语音翻译任务；</li>
	// <li>OFF：关闭智能语音翻译任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 视频源语言。
	SourceLanguage *string `json:"SourceLanguage,omitnil" name:"SourceLanguage"`

	// 翻译目标语言。
	DestinationLanguage *string `json:"DestinationLanguage,omitnil" name:"DestinationLanguage"`
}

type TranslateConfigureInfoForUpdate struct {
	// 语音翻译任务开关，可选值：
	// <li>ON：开启智能语音翻译任务；</li>
	// <li>OFF：关闭智能语音翻译任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 视频源语言。
	SourceLanguage *string `json:"SourceLanguage,omitnil" name:"SourceLanguage"`

	// 翻译目标语言。
	DestinationLanguage *string `json:"DestinationLanguage,omitnil" name:"DestinationLanguage"`
}

type UrlInputInfo struct {
	// 视频的 URL。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type UserDefineAsrTextReviewTemplateInfo struct {
	// 用户自定语音审核任务开关，可选值：
	// <li>ON：开启自定义语音审核任务；</li>
	// <li>OFF：关闭自定义语音审核任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 用户自定义语音过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义语音关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type UserDefineAsrTextReviewTemplateInfoForUpdate struct {
	// 用户自定语音审核任务开关，可选值：
	// <li>ON：开启自定义语音审核任务；</li>
	// <li>OFF：关闭自定义语音审核任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 用户自定义语音过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义语音关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type UserDefineConfigureInfo struct {
	// 用户自定义人物审核控制参数。
	FaceReviewInfo *UserDefineFaceReviewTemplateInfo `json:"FaceReviewInfo,omitnil" name:"FaceReviewInfo"`

	// 用户自定义语音审核控制参数。
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfo `json:"AsrReviewInfo,omitnil" name:"AsrReviewInfo"`

	// 用户自定义文本审核控制参数。
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfo `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type UserDefineConfigureInfoForUpdate struct {
	// 用户自定义人物审核控制参数。
	FaceReviewInfo *UserDefineFaceReviewTemplateInfoForUpdate `json:"FaceReviewInfo,omitnil" name:"FaceReviewInfo"`

	// 用户自定义语音审核控制参数。
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil" name:"AsrReviewInfo"`

	// 用户自定义文本审核控制参数。
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil" name:"OcrReviewInfo"`
}

type UserDefineFaceReviewTemplateInfo struct {
	// 用户自定义人物审核任务开关，可选值：
	// <li>ON：开启自定义人物审核任务；</li>
	// <li>OFF：关闭自定义人物审核任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 用户自定义人物过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义人物库的时，需要添加对应人物标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 97 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 95 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type UserDefineFaceReviewTemplateInfoForUpdate struct {
	// 用户自定义人物审核任务开关，可选值：
	// <li>ON：开启自定义人物审核任务；</li>
	// <li>OFF：关闭自定义人物审核任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 用户自定义人物过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义人物库的时，需要添加对应人物标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfo struct {
	// 用户自定文本审核任务开关，可选值：
	// <li>ON：开启自定义文本审核任务；</li>
	// <li>OFF：关闭自定义文本审核任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 用户自定义文本过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义文本关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfoForUpdate struct {
	// 用户自定文本审核任务开关，可选值：
	// <li>ON：开启自定义文本审核任务；</li>
	// <li>OFF：关闭自定义文本审核任务。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 用户自定义文本过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义文本关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet *string `json:"LabelSet,omitnil" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil" name:"ReviewConfidence"`
}

type VideoDenoiseConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 类型，可选值：
	// <li>weak</li>
	// <li>strong</li>
	// 默认值：weak。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type VideoEnhanceConfig struct {
	// 插帧帧率配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameRate *FrameRateConfig `json:"FrameRate,omitnil" name:"FrameRate"`

	// 超分配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuperResolution *SuperResolutionConfig `json:"SuperResolution,omitnil" name:"SuperResolution"`

	// HDR配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hdr *HdrConfig `json:"Hdr,omitnil" name:"Hdr"`

	// 视频降噪配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Denoise *VideoDenoiseConfig `json:"Denoise,omitnil" name:"Denoise"`

	// 综合增强配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageQualityEnhance *ImageQualityEnhanceConfig `json:"ImageQualityEnhance,omitnil" name:"ImageQualityEnhance"`

	// 色彩增强配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColorEnhance *ColorEnhanceConfig `json:"ColorEnhance,omitnil" name:"ColorEnhance"`

	// 细节增强配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SharpEnhance *SharpEnhanceConfig `json:"SharpEnhance,omitnil" name:"SharpEnhance"`

	// 人脸增强配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceEnhance *FaceEnhanceConfig `json:"FaceEnhance,omitnil" name:"FaceEnhance"`

	// 低光照增强配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LowLightEnhance *LowLightEnhanceConfig `json:"LowLightEnhance,omitnil" name:"LowLightEnhance"`

	// 去划痕配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScratchRepair *ScratchRepairConfig `json:"ScratchRepair,omitnil" name:"ScratchRepair"`

	// 去伪影（毛刺）配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArtifactRepair *ArtifactRepairConfig `json:"ArtifactRepair,omitnil" name:"ArtifactRepair"`
}

type VideoTemplateInfo struct {
	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码</li>
	// <li>libx265：H.265 编码</li>
	// <li>av1：AOMedia Video 1 编码</li>
	// 注意：目前 H.265 编码必须指定分辨率，并且需要在 640*480 以内。
	// 注意：av1 编码容器目前只支持 mp4 。
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// 视频帧率，取值范围：[0, 120]，单位：Hz。
	// 当取值为 0，表示帧率和原始视频保持一致。
	// 注意：自适应码率时取值范围是 [0, 60]
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	// 注意：自适应模式时，Width不能小于Height。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 关键帧 I 帧之间的间隔，取值范围：0 和 [1, 100000]，单位：帧数。
	// 当填 0 或不填时，系统将自动设置 gop 长度。
	Gop *uint64 `json:"Gop,omitnil" name:"Gop"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊填充。</li>
	// 默认值：black 。
	// 注意：自适应码流只支持 stretch、black。
	FillType *string `json:"FillType,omitnil" name:"FillType"`

	// 视频恒定码率控制因子，取值范围为[1, 51]。
	// 如果指定该参数，将使用 CRF 的码率控制方式做转码（视频码率将不再生效）。
	// 如果没有特殊需求，不建议指定该参数。
	Vcrf *uint64 `json:"Vcrf,omitnil" name:"Vcrf"`
}

type VideoTemplateInfoForUpdate struct {
	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码</li>
	// <li>libx265：H.265 编码</li>
	// <li>av1：AOMedia Video 1 编码</li>
	// 注意：目前 H.265 编码必须指定分辨率，并且需要在 640*480 以内。
	// 注意：av1 编码容器目前只支持 mp4 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// 视频帧率，取值范围：[0, 120]，单位：Hz。
	// 当取值为 0，表示帧率和原始视频保持一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 注意：自适应模式时，Width不能小于Height。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 关键帧 I 帧之间的间隔，取值范围：0 和 [1, 100000]，单位：帧数。当填 0 时，系统将自动设置 gop 长度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gop *uint64 `json:"Gop,omitnil" name:"Gop"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊填充。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FillType *string `json:"FillType,omitnil" name:"FillType"`

	// 视频恒定码率控制因子。取值范围为[0, 51]，填0表示禁用该参数。
	// 如果没有特殊需求，不建议指定该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vcrf *uint64 `json:"Vcrf,omitnil" name:"Vcrf"`

	// 内容自适应编码。可选值：
	// <li>0：不开启</li>
	// <li>1：开启</li>
	// 默认值: 0.   当开启该参数时，将会自适应生成多个不同分辨率，不同码率的码流， 其中VideoTemplate的宽和高为多个码流中的最大分辨率，VideoTemplate中的码率为多个码流中的最高码率， VideoTemplate中的vcrf为多个码流中的最高质量。 当不设置分辨率、码率和vcrf时， ContentAdaptStream 参数生成的最高分辨率为视频源的分辨率，视频质量为接近vmaf95分。 若要开启该参数或了解计费细节, 请联系您的腾讯云商务。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentAdaptStream *uint64 `json:"ContentAdaptStream,omitnil" name:"ContentAdaptStream"`
}

type VolumeBalanceConfig struct {
	// 能力配置开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	// 默认值：ON。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 类型，可选值：
	// <li>loudNorm：响度标准化</li>
	// <li>gainControl：减小突变</li>
	// 默认值：loudNorm。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type WatermarkInput struct {
	// 水印模板 ID。
	Definition *uint64 `json:"Definition,omitnil" name:"Definition"`

	// 水印自定义参数，当 Definition 填 0 时有效。
	// 该参数用于高度定制场景，建议您优先使用 Definition 指定水印参数。
	// 水印自定义参数不支持截图打水印。
	RawParameter *RawWatermarkParameter `json:"RawParameter,omitnil" name:"RawParameter"`

	// 文字内容，长度不超过100个字符。仅当水印类型为文字水印时填写。
	// 文字水印不支持截图打水印。
	TextContent *string `json:"TextContent,omitnil" name:"TextContent"`

	// SVG 内容。长度不超过 2000000 个字符。仅当水印类型为 SVG 水印时填写。
	// SVG 水印不支持截图打水印。
	SvgContent *string `json:"SvgContent,omitnil" name:"SvgContent"`

	// 水印的起始时间偏移，单位：秒。不填或填0，表示水印从画面出现时开始显现。
	// <li>不填或填0，表示水印从画面开始就出现；</li>
	// <li>当数值大于0时（假设为 n），表示水印从画面开始的第 n 秒出现；</li>
	// <li>当数值小于0时（假设为 -n），表示水印从离画面结束 n 秒前开始出现。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil" name:"StartTimeOffset"`

	// 水印的结束时间偏移，单位：秒。
	// <li>不填或填0，表示水印持续到画面结束；</li>
	// <li>当数值大于0时（假设为 n），表示水印持续到第 n 秒时消失；</li>
	// <li>当数值小于0时（假设为 -n），表示水印持续到离画面结束 n 秒前消失。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil" name:"EndTimeOffset"`
}

type WatermarkTemplate struct {
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil" name:"Definition"`

	// 水印类型，取值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 水印模板名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 水印图片原点距离视频图像原点的水平位置。
	// <li>当字符串以 % 结尾，表示水印 Left 为视频宽度指定百分比的位置，如 10% 表示 Left 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Left 为视频宽度指定像素的位置，如 100px 表示 Left 为 100 像素。</li>
	XPos *string `json:"XPos,omitnil" name:"XPos"`

	// 水印图片原点距离视频图像原点的垂直位置。
	// <li>当字符串以 % 结尾，表示水印 Top 为视频高度指定百分比的位置，如 10% 表示 Top 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Top 为视频高度指定像素的位置，如 100px 表示 Top 为 100 像素。</li>
	YPos *string `json:"YPos,omitnil" name:"YPos"`

	// 图片水印模板，仅当 Type 为 image，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTemplate *ImageWatermarkTemplate `json:"ImageTemplate,omitnil" name:"ImageTemplate"`

	// 文字水印模板，仅当 Type 为 text，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitnil" name:"TextTemplate"`

	// SVG 水印模板，当 Type 为 svg，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitnil" name:"SvgTemplate"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 原点位置，可选值：
	// <li>topLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>topRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>bottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>bottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下。；</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil" name:"CoordinateOrigin"`
}

// Predefined struct for user
type WithdrawsWatermarkRequestParams struct {
	// 输入媒体文件存储信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`
}

type WithdrawsWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 输入媒体文件存储信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`
}

func (r *WithdrawsWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WithdrawsWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	delete(f, "TaskNotifyConfig")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WithdrawsWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WithdrawsWatermarkResponseParams struct {
	// 任务 ID，可以通过该 ID 查询任务状态和结果。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type WithdrawsWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *WithdrawsWatermarkResponseParams `json:"Response"`
}

func (r *WithdrawsWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WithdrawsWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WorkflowInfo struct {
	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitnil" name:"WorkflowId"`

	// 工作流名称。
	WorkflowName *string `json:"WorkflowName,omitnil" name:"WorkflowName"`

	// 工作流状态，取值范围：
	// <li>Enabled：已启用，</li>
	// <li>Disabled：已禁用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 工作流绑定的输入规则，当上传视频命中该规则到该对象时即触发工作流。
	Trigger *WorkflowTrigger `json:"Trigger,omitnil" name:"Trigger"`

	// 媒体处理的文件输出存储位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil" name:"OutputStorage"`

	// 媒体处理类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil" name:"AiRecognitionTask"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil" name:"TaskNotifyConfig"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TaskPriority *int64 `json:"TaskPriority,omitnil" name:"TaskPriority"`

	// 媒体处理生成的文件输出的目标目录，如`/movie/201907/`。
	OutputDir *string `json:"OutputDir,omitnil" name:"OutputDir"`

	// 工作流创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 工作流最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type WorkflowTask struct {
	// 媒体处理任务 ID。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 源异常时返回非0错误码，返回0 时请使用各个具体任务的 ErrCode。
	ErrCode *int64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 源异常时返回对应异常Message，否则请使用各个具体任务的 Message。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 媒体处理的目标文件信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 原始视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil" name:"MetaData"`

	// 媒体处理任务的执行状态与结果。
	MediaProcessResultSet []*MediaProcessTaskResult `json:"MediaProcessResultSet,omitnil" name:"MediaProcessResultSet"`

	// 视频内容审核任务的执行状态与结果。
	AiContentReviewResultSet []*AiContentReviewResult `json:"AiContentReviewResultSet,omitnil" name:"AiContentReviewResultSet"`

	// 视频内容分析任务的执行状态与结果。
	AiAnalysisResultSet []*AiAnalysisResult `json:"AiAnalysisResultSet,omitnil" name:"AiAnalysisResultSet"`

	// 视频内容识别任务的执行状态与结果。
	AiRecognitionResultSet []*AiRecognitionResult `json:"AiRecognitionResultSet,omitnil" name:"AiRecognitionResultSet"`

	// 视频质检任务的执行状态与结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiQualityControlTaskResult *ScheduleQualityControlTaskResult `json:"AiQualityControlTaskResult,omitnil" name:"AiQualityControlTaskResult"`
}

type WorkflowTrigger struct {
	// 触发器的类型，可选值：
	// <li>CosFileUpload：COS触发</li>
	// <li>AwsS3FileUpload：AWS触发，目前只支持转码任务。只有编排支持，工作流不支持。  </li>
	// 
	Type *string `json:"Type,omitnil" name:"Type"`

	// 当 Type 为 CosFileUpload 时必填且有效，为 COS 触发规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosFileUploadTrigger *CosFileUploadTrigger `json:"CosFileUploadTrigger,omitnil" name:"CosFileUploadTrigger"`

	// 当 Type 为 AwsS3FileUpload 时必填且有效，为 AWS S3 触发规则。
	// 
	// 注意：目前AWS的S3、对应触发队列SQS、回调队列SQS的秘钥需要一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AwsS3FileUploadTrigger *AwsS3FileUploadTrigger `json:"AwsS3FileUploadTrigger,omitnil" name:"AwsS3FileUploadTrigger"`
}