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

package v20180717

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AIAnalysisTemplateItem struct {
	// 智能分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 智能分析模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 智能分析模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 智能分类任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// 智能封面任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`

	// 智能精彩集锦任务控制参数。
	HighlightConfigure *HighlightsConfigureInfo `json:"HighlightConfigure,omitnil,omitempty" name:"HighlightConfigure"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AIRecognitionTemplateItem struct {
	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 视频内容识别模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 视频内容识别模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 模板类型，取值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 头尾识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadTailConfigure *HeadTailConfigureInfo `json:"HeadTailConfigure,omitnil,omitempty" name:"HeadTailConfigure"`

	// 拆条识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentConfigure *SegmentConfigureInfo `json:"SegmentConfigure,omitnil,omitempty" name:"SegmentConfigure"`

	// 人脸识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrTranslateConfigure *AsrTranslateConfigureInfo `json:"AsrTranslateConfigure,omitnil,omitempty" name:"AsrTranslateConfigure"`

	// 物体识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectConfigure *ObjectConfigureInfo `json:"ObjectConfigure,omitnil,omitempty" name:"ObjectConfigure"`

	// 截图时间间隔，单位：秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AbnormalLightingConfigureInfo struct {
	// 视频画面低光、过曝检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type AbnormalLightingConfigureInfoForUpdate struct {
	// 视频画面低光、过曝检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type AccelerateAreaInfo struct {
	// 加速地区，可选值：
	// <li>Chinese Mainland：中国境内（不包含港澳台）。</li>
	// <li>Outside Chinese Mainland：中国境外。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 腾讯禁用原因，可选值：
	// <li>ForLegalReasons：因法律原因导致关闭加速；</li>
	// <li>ForOverdueBills：因欠费停服导致关闭加速。</li>
	TencentDisableReason *string `json:"TencentDisableReason,omitnil,omitempty" name:"TencentDisableReason"`

	// 加速域名对应的 CNAME 域名。
	TencentEdgeDomain *string `json:"TencentEdgeDomain,omitnil,omitempty" name:"TencentEdgeDomain"`
}

type AdaptiveDynamicStreamingInfoItem struct {
	// 转自适应码流规格。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 打包格式，取值范围：
	// <li>HLS；</li>
	// <li>DASH。</li>
	Package *string `json:"Package,omitnil,omitempty" name:"Package"`

	// 加密类型。
	DrmType *string `json:"DrmType,omitnil,omitempty" name:"DrmType"`

	// 播放地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 媒体文件大小，单位：字节。
	// <li>当媒体文件为 HLS 时，大小是 m3u8 和 ts 文件大小的总和；</li>
	// <li>当媒体文件为 DASH 时，大小是 mpd 和分片文件大小的总和；</li>
	// <li><font color=red>注意</font>：在 2022-01-10T16:00:00Z 前处理生成的自适应码流文件此字段为0。</li>
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 数字水印类型。可选值：
	// <li>Trace 表示经过溯源水印处理；</li>
	// <li>CopyRight 表示经过版权水印处理；</li>
	// <li>None 表示没有经过数字水印处理。</li>
	DigitalWatermarkType *string `json:"DigitalWatermarkType,omitnil,omitempty" name:"DigitalWatermarkType"`

	// 子流信息列表。
	SubStreamSet []*MediaSubStreamInfoItem `json:"SubStreamSet,omitnil,omitempty" name:"SubStreamSet"`

	// 版权信息。
	CopyRightWatermarkText *string `json:"CopyRightWatermarkText,omitnil,omitempty" name:"CopyRightWatermarkText"`

	// 字幕信息列表。
	SubtitleSet []*MediaSubtitleItem `json:"SubtitleSet,omitnil,omitempty" name:"SubtitleSet"`

	// 默认字幕的唯一标识。
	DefaultSubtitleId *string `json:"DefaultSubtitleId,omitnil,omitempty" name:"DefaultSubtitleId"`
}

type AdaptiveDynamicStreamingTaskInput struct {
	// 转自适应码流模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// 溯源水印，该功能已不再支持使用。
	TraceWatermark *TraceWatermarkInput `json:"TraceWatermark,omitnil,omitempty" name:"TraceWatermark"`

	// 版权水印，该功能已不再支持使用。
	CopyRightWatermark *CopyRightWatermarkInput `json:"CopyRightWatermark,omitnil,omitempty" name:"CopyRightWatermark"`

	// 字幕列表，元素为字幕 ID，支持多个字幕，最大可支持16个。
	SubtitleSet []*string `json:"SubtitleSet,omitnil,omitempty" name:"SubtitleSet"`
}

type AdaptiveDynamicStreamingTemplate struct {
	// 转自适应码流模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 转自适应码流模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 转自适应码流模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 自适应转码格式，取值范围：
	// <li>HLS。</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// DRM 类型，取值范围：
	// <li>SimpleAES</li>
	// <li>Widevine</li>
	// <li>FairPlay</li>
	// 如果取值为空字符串，代表不对视频做 DRM 保护。
	DrmType *string `json:"DrmType,omitnil,omitempty" name:"DrmType"`

	// DRM 的密钥提供商，取值范围：
	// <li>SDMC：华曦达；</li>
	// <li>VOD：云点播。</li>
	// 默认值为 VOD 。
	DrmKeyProvider *string `json:"DrmKeyProvider,omitnil,omitempty" name:"DrmKeyProvider"`

	// 自适应转码输入流参数信息，最多输入10路流。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 切片类型，仅当 Format 为 HLS 时有效。
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
}

type AdaptiveStreamTemplate struct {
	// 视频参数信息。
	Video *VideoTemplateInfo `json:"Video,omitnil,omitempty" name:"Video"`

	// 音频参数信息。
	Audio *AudioTemplateInfo `json:"Audio,omitnil,omitempty" name:"Audio"`

	// 是否移除音频流，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	RemoveAudio *uint64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// 是否移除视频流，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	RemoveVideo *uint64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 极速高清转码参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`
}

type AiAnalysisResult struct {
	// 任务的类型，可以取的值有：
	// <li>Classification：智能分类</li>
	// <li>Cover：智能封面</li>
	// <li>Tag：智能标签</li>
	// <li>FrameTag：智能按帧标签</li>
	// <li>Highlight：智能精彩集锦</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 视频内容分析智能分类任务的查询结果，当任务类型为 Classification 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationTask *AiAnalysisTaskClassificationResult `json:"ClassificationTask,omitnil,omitempty" name:"ClassificationTask"`

	// 视频内容分析智能封面任务的查询结果，当任务类型为 Cover 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverTask *AiAnalysisTaskCoverResult `json:"CoverTask,omitnil,omitempty" name:"CoverTask"`

	// 视频内容分析智能标签任务的查询结果，当任务类型为 Tag 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagTask *AiAnalysisTaskTagResult `json:"TagTask,omitnil,omitempty" name:"TagTask"`

	// 视频内容分析智能按帧标签任务的查询结果，当任务类型为 FrameTag 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagTask *AiAnalysisTaskFrameTagResult `json:"FrameTagTask,omitnil,omitempty" name:"FrameTagTask"`

	// 视频内容分析智能精彩集锦任务的查询结果，当任务类型为 Highlight 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighlightTask *AiAnalysisTaskHighlightResult `json:"HighlightTask,omitnil,omitempty" name:"HighlightTask"`
}

type AiAnalysisTaskClassificationInput struct {
	// 视频智能分类模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskClassificationOutput struct {
	// 视频智能分类列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 ClassificationSetFileUrl 对应的文件中获取。
	ClassificationSet []*MediaAiAnalysisClassificationItem `json:"ClassificationSet,omitnil,omitempty" name:"ClassificationSet"`

	// 视频智能分类列表文件 URL。文件的内容为 JSON，数据结构与 ClassificationSet 字段一致。 （文件不会永久存储，到达 ClassificationSetFileUrlExpireTime 时间点后文件将被删除）。
	ClassificationSetFileUrl *string `json:"ClassificationSetFileUrl,omitnil,omitempty" name:"ClassificationSetFileUrl"`

	// 视频智能分类列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	ClassificationSetFileUrlExpireTime *string `json:"ClassificationSetFileUrlExpireTime,omitnil,omitempty" name:"ClassificationSetFileUrlExpireTime"`
}

type AiAnalysisTaskClassificationResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 智能分类任务输入。
	Input *AiAnalysisTaskClassificationInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 智能分类任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskClassificationOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 智能分类任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 智能分类任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 智能分类任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiAnalysisTaskCoverInput struct {
	// 视频智能封面模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskCoverOutput struct {
	// 智能封面列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 CoverSetFileUrl 对应的文件中获取。
	CoverSet []*MediaAiAnalysisCoverItem `json:"CoverSet,omitnil,omitempty" name:"CoverSet"`

	// 智能封面列表文件 URL。文件的内容为 JSON，数据结构与 CoverSet 字段一致。 （文件不会永久存储，到达 CoverSetFileUrlExpireTime 时间点后文件将被删除）。
	CoverSetFileUrl *string `json:"CoverSetFileUrl,omitnil,omitempty" name:"CoverSetFileUrl"`

	// 智能封面列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CoverSetFileUrlExpireTime *string `json:"CoverSetFileUrlExpireTime,omitnil,omitempty" name:"CoverSetFileUrlExpireTime"`
}

type AiAnalysisTaskCoverResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 智能封面任务输入。
	Input *AiAnalysisTaskCoverInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 智能封面任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskCoverOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 智能封面任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 智能封面任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 智能封面任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiAnalysisTaskFrameTagInput struct {
	// 视频智能按帧标签模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskFrameTagOutput struct {
	// 视频按帧标签列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaAiAnalysisFrameTagSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 视频按帧标签列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 视频按帧标签列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiAnalysisTaskFrameTagResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 智能按帧标签任务输入。
	Input *AiAnalysisTaskFrameTagInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 智能按帧标签任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskFrameTagOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 智能按帧标签任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 智能按帧标签任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 智能按帧标签任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiAnalysisTaskHighlightInput struct {
	// 视频智能精彩片段模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskHighlightOutput struct {
	// 视频智能精彩片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 HighlightSetFileUrl 对应的文件中获取。
	HighlightSet []*MediaAiAnalysisHighlightItem `json:"HighlightSet,omitnil,omitempty" name:"HighlightSet"`

	// 视频智能精彩片段列表文件 URL。文件的内容为 JSON，数据结构与 HighlightSet 字段一致。 （文件不会永久存储，到达 HighlightSetFileUrlExpireTime 时间点后文件将被删除）。
	HighlightSetFileUrl *string `json:"HighlightSetFileUrl,omitnil,omitempty" name:"HighlightSetFileUrl"`

	// 视频智能精彩片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	HighlightSetFileUrlExpireTime *string `json:"HighlightSetFileUrlExpireTime,omitnil,omitempty" name:"HighlightSetFileUrlExpireTime"`
}

type AiAnalysisTaskHighlightResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 智能精彩片段任务输入。
	Input *AiAnalysisTaskHighlightInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 智能精彩片段任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskHighlightOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 智能精彩片段任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 智能精彩片段任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 智能精彩片段任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiAnalysisTaskInput struct {
	// 视频内容分析模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskTagInput struct {
	// 视频智能标签模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskTagOutput struct {
	// 视频智能标签列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 TagSetFileUrl 对应的文件中获取。
	TagSet []*MediaAiAnalysisTagItem `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 视频智能标签列表文件 URL。文件的内容为 JSON，数据结构与 TagSet 字段一致。 （文件不会永久存储，到达 TagSetFileUrlExpireTime 时间点后文件将被删除）。
	TagSetFileUrl *string `json:"TagSetFileUrl,omitnil,omitempty" name:"TagSetFileUrl"`

	// 视频智能标签列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	TagSetFileUrlExpireTime *string `json:"TagSetFileUrlExpireTime,omitnil,omitempty" name:"TagSetFileUrlExpireTime"`
}

type AiAnalysisTaskTagResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 智能标签任务输入。
	Input *AiAnalysisTaskTagInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 智能标签任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskTagOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 智能标签任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 智能标签任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 智能标签任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiContentReviewResult struct {
	// 任务的类型，可以取的值有：
	// <li>Porn：图片鉴别是否涉及令人反感的信息</li>
	// <li>Terrorism：图片鉴别是否涉及令人不安全的信息</li>
	// <li>Political：图片鉴别是否涉及令人不适宜的信息</li>
	// <li>Porn.Asr：Asr 文字（ 音频中的文字）鉴别是否涉及令人反感的信息</li>
	// <li>Porn.Ocr：Ocr 文字鉴别是否涉及令人反感的信息</li>
	// <li>Political.Asr：Asr 文字（ 音频中的文字）鉴别是否涉及令人不适宜的信息</li>
	// <li>Political.Ocr：Ocr 文字鉴别是否涉及令人不适宜的信息</li>
	// <li>Terrorism.Ocr：Ocr 文字鉴别是否涉及令人不安全的信息</li>
	// <li>Prohibited.Asr：Asr 文字（ 音频中的文字）鉴违禁</li>
	// <li>Prohibited.Ocr：Ocr 文字鉴违禁</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 视频音视频审核任务（画面涉及令人反感的信息）的查询结果，当任务类型为 Porn 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornTask *AiReviewTaskPornResult `json:"PornTask,omitnil,omitempty" name:"PornTask"`

	// 视频音视频审核任务（画面涉及令人不安全的信息）的查询结果，当任务类型为 Terrorism 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismTask *AiReviewTaskTerrorismResult `json:"TerrorismTask,omitnil,omitempty" name:"TerrorismTask"`

	// 视频音视频审核任务（画面涉及令人不适宜的信息）的查询结果，当任务类型为 Political 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalTask *AiReviewTaskPoliticalResult `json:"PoliticalTask,omitnil,omitempty" name:"PoliticalTask"`

	// 视频音视频审核任务（Asr 文字涉及令人反感的信息）的查询结果，当任务类型为 Porn.Asr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornAsrTask *AiReviewTaskPornAsrResult `json:"PornAsrTask,omitnil,omitempty" name:"PornAsrTask"`

	// 视频音视频审核任务（Ocr 文字涉及令人反感的信息）的查询结果，当任务类型为 Porn.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornOcrTask *AiReviewTaskPornOcrResult `json:"PornOcrTask,omitnil,omitempty" name:"PornOcrTask"`

	// 视频音视频审核任务（Asr 文字涉及令人不适宜的信息）的查询结果，当任务类型为 Political.Asr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalAsrTask *AiReviewTaskPoliticalAsrResult `json:"PoliticalAsrTask,omitnil,omitempty" name:"PoliticalAsrTask"`

	// 视频音视频审核任务（Ocr 文字涉及令人不适宜的信息）的查询结果，当任务类型为 Political.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalOcrTask *AiReviewTaskPoliticalOcrResult `json:"PoliticalOcrTask,omitnil,omitempty" name:"PoliticalOcrTask"`

	// 视频音视频审核任务（ Ocr 文字涉及令人不安全的信息）的查询结果，当任务类型为 Terrorism.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismOcrTask *AiReviewTaskTerrorismOcrResult `json:"TerrorismOcrTask,omitnil,omitempty" name:"TerrorismOcrTask"`

	// 视频音视频审核 Ocr 文字鉴违禁任务的查询结果，当任务类型为 Prohibited.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProhibitedOcrTask *AiReviewTaskProhibitedOcrResult `json:"ProhibitedOcrTask,omitnil,omitempty" name:"ProhibitedOcrTask"`

	// 视频音视频审核 Asr 文字鉴违禁任务的查询结果，当任务类型为 Prohibited.Asr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProhibitedAsrTask *AiReviewTaskProhibitedAsrResult `json:"ProhibitedAsrTask,omitnil,omitempty" name:"ProhibitedAsrTask"`
}

type AiContentReviewTaskInput struct {
	// 音视频审核模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionResult struct {
	// 任务的类型，取值范围：
	// <li>FaceRecognition：人脸识别，</li>
	// <li>AsrWordsRecognition：语音关键词识别，</li>
	// <li>OcrWordsRecognition：文本关键词识别，</li>
	// <li>AsrFullTextRecognition：语音全文识别，</li>
	// <li>AsrTranslateRecognition：语音翻译识别，</li>
	// <li>OcrFullTextRecognition：文本全文识别，</li>
	// <li>HeadTailRecognition：视频片头片尾识别，</li>
	// <li>ObjectRecognition：物体识别。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 视频片头片尾识别结果，当 Type 为
	//  HeadTailRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadTailTask *AiRecognitionTaskHeadTailResult `json:"HeadTailTask,omitnil,omitempty" name:"HeadTailTask"`

	// 视频拆条识别结果，当 Type 为
	//  SegmentRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentTask *AiRecognitionTaskSegmentResult `json:"SegmentTask,omitnil,omitempty" name:"SegmentTask"`

	// 人脸识别结果，当 Type 为 
	//  FaceRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceTask *AiRecognitionTaskFaceResult `json:"FaceTask,omitnil,omitempty" name:"FaceTask"`

	// 语音关键词识别结果，当 Type 为
	//  AsrWordsRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrWordsTask *AiRecognitionTaskAsrWordsResult `json:"AsrWordsTask,omitnil,omitempty" name:"AsrWordsTask"`

	// 语音全文识别结果，当 Type 为
	//  AsrFullTextRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrFullTextTask *AiRecognitionTaskAsrFullTextResult `json:"AsrFullTextTask,omitnil,omitempty" name:"AsrFullTextTask"`

	// 语音翻译结果，当 Type 为 AsrTranslateRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrTranslateTask *AiRecognitionTaskAsrTranslateResult `json:"AsrTranslateTask,omitnil,omitempty" name:"AsrTranslateTask"`

	// 文本关键词识别结果，当 Type 为
	//  OcrWordsRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrWordsTask *AiRecognitionTaskOcrWordsResult `json:"OcrWordsTask,omitnil,omitempty" name:"OcrWordsTask"`

	// 文本全文识别结果，当 Type 为
	//  OcrFullTextRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrFullTextTask *AiRecognitionTaskOcrFullTextResult `json:"OcrFullTextTask,omitnil,omitempty" name:"OcrFullTextTask"`

	// 物体识别结果，当 Type 为
	//  ObjectRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectTask *AiRecognitionTaskObjectResult `json:"ObjectTask,omitnil,omitempty" name:"ObjectTask"`
}

type AiRecognitionTaskAsrFullTextResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 语音全文识别任务输入信息。
	Input *AiRecognitionTaskAsrFullTextResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 语音全文识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskAsrFullTextResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 语音全文识别任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 语音全文识别任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 语音全文识别任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiRecognitionTaskAsrFullTextResultInput struct {
	// 语音全文识别模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrFullTextResultOutput struct {
	// 语音全文识别片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*AiRecognitionTaskAsrFullTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 语音全文识别片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 语音全文识别片段列表文件 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`

	// 生成的字幕列表，对应 [语音全文识别任务控制参数](https://cloud.tencent.com/document/api/266/31773#AsrFullTextConfigureInfo) SubtitleFormats。
	SubtitleSet []*AiRecognitionTaskAsrFullTextResultOutputSubtitleItem `json:"SubtitleSet,omitnil,omitempty" name:"SubtitleSet"`

	// 生成的字幕文件 Url，对应 [语音全文识别任务控制参数](https://cloud.tencent.com/document/api/266/31773#AsrFullTextConfigureInfo) SubtitleFormat。
	SubtitleUrl *string `json:"SubtitleUrl,omitnil,omitempty" name:"SubtitleUrl"`
}

type AiRecognitionTaskAsrFullTextResultOutputSubtitleItem struct {
	// 媒资字幕 ID，用于媒资字幕管理，仅当 Format 为 vtt 时有效。
	// <font color=red>注意：</font>2024-11-01T10:00:00Z 之前的任务返回此字段无效。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 媒资字幕名字，用于播放器展示，仅当 Format 为 vtt 时有效。
	// <font color=red>注意：</font>2024-11-01T10:00:00Z 之前的任务返回此字段无效。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字幕语言。
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 字幕文件格式，取值范围：
	// <li>vtt：WebVTT 字幕文件；</li>
	// <li>srt：SRT 字幕文件。</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 字幕文件 Url。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type AiRecognitionTaskAsrFullTextSegmentItem struct {
	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 识别文本。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type AiRecognitionTaskAsrTranslateResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 语音翻译任务输入信息。
	Input *AiRecognitionTaskAsrTranslateResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 语音翻译任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskAsrTranslateResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 语音翻译任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 语音翻译任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 语音翻译任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiRecognitionTaskAsrTranslateResultInput struct {
	// 语音翻译模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrTranslateResultOutput struct {
	// 语音翻译片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*AiRecognitionTaskAsrTranslateSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 语音翻译片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 语音翻译片段列表文件 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`

	// 生成的字幕列表。
	SubtitleSet []*AiRecognitionTaskAsrFullTextResultOutputSubtitleItem `json:"SubtitleSet,omitnil,omitempty" name:"SubtitleSet"`
}

type AiRecognitionTaskAsrTranslateSegmentItem struct {
	// 语音翻译片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 语音翻译片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 语音翻译片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 识别文本。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 翻译文本。
	Translation *string `json:"Translation,omitnil,omitempty" name:"Translation"`
}

type AiRecognitionTaskAsrWordsResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 语音关键词识别任务输入信息。
	Input *AiRecognitionTaskAsrWordsResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 语音关键词识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskAsrWordsResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 语音关键词识别任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 语音关键词识别任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 语音关键词识别任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiRecognitionTaskAsrWordsResultInput struct {
	// 语音关键词识别模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrWordsResultItem struct {
	// 语音关键词。
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 语音关键词出现的时间片段列表。
	SegmentSet []*AiRecognitionTaskAsrWordsSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskAsrWordsResultOutput struct {
	// 语音关键词识别结果集。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 ResultSetFileUrl 对应的文件中获取。
	ResultSet []*AiRecognitionTaskAsrWordsResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`

	// 语音关键词识别结果集文件 URL。文件的内容为 JSON，数据结构与 ResultSet 字段一致。 （文件不会永久存储，到达ResultSetFileUrlExpireTime 时间点后文件将被删除）。
	ResultSetFileUrl *string `json:"ResultSetFileUrl,omitnil,omitempty" name:"ResultSetFileUrl"`

	// 语音关键词识别结果集文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	ResultSetFileUrlExpireTime *string `json:"ResultSetFileUrlExpireTime,omitnil,omitempty" name:"ResultSetFileUrlExpireTime"`
}

type AiRecognitionTaskAsrWordsSegmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type AiRecognitionTaskFaceResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 人脸识别任务输入信息。
	Input *AiRecognitionTaskFaceResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 人脸识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskFaceResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 人脸识别任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 人脸识别任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 人脸识别任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiRecognitionTaskFaceResultInput struct {
	// 人脸识别模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskFaceResultItem struct {
	// 人物唯一标识 ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 人物库类型，表示识别出的人物来自哪个人物库：
	// <li>Default：默认人物库；</li>
	// <li>UserDefine：用户自定义人物库。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 人物名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 人物出现的片段结果集。
	SegmentSet []*AiRecognitionTaskFaceSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskFaceResultOutput struct {
	// 智能人脸识别结果集。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 ResultSetFileUrl 对应的文件中获取。
	ResultSet []*AiRecognitionTaskFaceResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`

	// 智能人脸识别结果集文件 URL。文件的内容为 JSON，数据结构与 ResultSet 字段一致。 （文件不会永久存储，到达ResultSetFileUrlExpireTime 时间点后文件将被删除）。
	ResultSetFileUrl *string `json:"ResultSetFileUrl,omitnil,omitempty" name:"ResultSetFileUrl"`

	// 智能人脸识别结果集文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	ResultSetFileUrlExpireTime *string `json:"ResultSetFileUrlExpireTime,omitnil,omitempty" name:"ResultSetFileUrlExpireTime"`
}

type AiRecognitionTaskFaceSegmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskHeadTailResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 视频片头片尾识别任务输入信息。
	Input *AiRecognitionTaskHeadTailResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 视频片头片尾识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskHeadTailResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 视频片头片尾识别任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 视频片头片尾识别任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 视频片头片尾识别任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiRecognitionTaskHeadTailResultInput struct {
	// 视频片头片尾识别模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskHeadTailResultOutput struct {
	// 片头识别置信度。取值：0~100。
	HeadConfidence *float64 `json:"HeadConfidence,omitnil,omitempty" name:"HeadConfidence"`

	// 视频片头的结束时间点，单位：秒。
	HeadTimeOffset *float64 `json:"HeadTimeOffset,omitnil,omitempty" name:"HeadTimeOffset"`

	// 片尾识别置信度。取值：0~100。
	TailConfidence *float64 `json:"TailConfidence,omitnil,omitempty" name:"TailConfidence"`

	// 视频片尾的开始时间点，单位：秒。
	TailTimeOffset *float64 `json:"TailTimeOffset,omitnil,omitempty" name:"TailTimeOffset"`
}

type AiRecognitionTaskInput struct {
	// 视频智能识别模板 ID 。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskObjectResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 物体识别任务输入信息。
	Input *AiRecognitionTaskObjectResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 物体识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskObjectResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 物体识别任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 物体识别任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 物体识别任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiRecognitionTaskObjectResultInput struct {
	// 物体识别模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskObjectResultItem struct {
	// 识别的物体名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 物体出现的片段列表。
	//
	// Deprecated: SegmentSet is deprecated.
	SegmentSet []*AiRecognitionTaskObjectSeqmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 物体出现的片段列表。
	RecognitionSegmentSet []*AiRecognitionTaskObjectSegmentItem `json:"RecognitionSegmentSet,omitnil,omitempty" name:"RecognitionSegmentSet"`
}

type AiRecognitionTaskObjectResultOutput struct {
	// 智能物体识别结果集。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 ResultSetFileUrl 对应的文件中获取。
	ResultSet []*AiRecognitionTaskObjectResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`

	// 智能物体识别结果集文件 URL。文件的内容为 JSON，数据结构与 ResultSet 字段一致。 （文件不会永久存储，到达ResultSetFileUrlExpireTime 时间点后文件将被删除）。
	ResultSetFileUrl *string `json:"ResultSetFileUrl,omitnil,omitempty" name:"ResultSetFileUrl"`

	// 智能物体识别结果集文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	ResultSetFileUrlExpireTime *string `json:"ResultSetFileUrlExpireTime,omitnil,omitempty" name:"ResultSetFileUrlExpireTime"`
}

type AiRecognitionTaskObjectSegmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskObjectSeqmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskOcrFullTextResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 文本全文识别任务输入信息。
	Input *AiRecognitionTaskOcrFullTextResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 文本全文识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskOcrFullTextResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 文本全文识别任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 文本全文识别任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 文本全文识别任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiRecognitionTaskOcrFullTextResultInput struct {
	// 文本全文识别模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrFullTextResultOutput struct {
	// 文本全文识别结果集。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*AiRecognitionTaskOcrFullTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 文本全文识别结果集文件 URL。文件的内容为 JSON，数据结构与 ResultSet 字段一致。 （文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 文本全文识别结果集文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiRecognitionTaskOcrFullTextSegmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 识别片段结果集。
	TextSet []*AiRecognitionTaskOcrFullTextSegmentTextItem `json:"TextSet,omitnil,omitempty" name:"TextSet"`
}

type AiRecognitionTaskOcrFullTextSegmentTextItem struct {
	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// 识别文本。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type AiRecognitionTaskOcrWordsResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 文本关键词识别任务输入信息。
	Input *AiRecognitionTaskOcrWordsResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 文本关键词识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskOcrWordsResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 文本关键词识别任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 文本关键词识别任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 文本关键词识别任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiRecognitionTaskOcrWordsResultInput struct {
	// 文本关键词识别模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrWordsResultItem struct {
	// 文本关键词。
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 文本关键出现的片段列表。
	SegmentSet []*AiRecognitionTaskOcrWordsSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskOcrWordsResultOutput struct {
	// 文本关键词识别结果集。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 ResultSetFileUrl 对应的文件中获取。
	ResultSet []*AiRecognitionTaskOcrWordsResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`

	// 文本关键词识别结果集文件 URL。文件的内容为 JSON，数据结构与 ResultSet 字段一致。 （文件不会永久存储，到达ResultSetFileUrlExpireTime 时间点后文件将被删除）。
	ResultSetFileUrl *string `json:"ResultSetFileUrl,omitnil,omitempty" name:"ResultSetFileUrl"`

	// 文本关键词识别结果集文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	ResultSetFileUrlExpireTime *string `json:"ResultSetFileUrlExpireTime,omitnil,omitempty" name:"ResultSetFileUrlExpireTime"`
}

type AiRecognitionTaskOcrWordsSegmentItem struct {
	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskSegmentResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 视频拆条任务输入信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *AiRecognitionTaskSegmentResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 视频拆条任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskSegmentResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 视频拆条任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 视频拆条任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 视频拆条任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AiRecognitionTaskSegmentResultInput struct {
	// 视频拆条模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskSegmentResultOutput struct {
	// 视频拆条片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*AiRecognitionTaskSegmentSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 视频拆条片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 视频拆条片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiRecognitionTaskSegmentSegmentItem struct {
	// 文件 ID。仅当处理的是点播文件并且拆条生成的子片段为点播文件时有效。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 视频拆条片段 Url。
	SegmentUrl *string `json:"SegmentUrl,omitnil,omitempty" name:"SegmentUrl"`

	// 拆条片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 拆条片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 拆条片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 拆条封面图片 Url。
	CovImgUrl *string `json:"CovImgUrl,omitnil,omitempty" name:"CovImgUrl"`

	// 特殊字段，请忽略。
	SpecialInfo *string `json:"SpecialInfo,omitnil,omitempty" name:"SpecialInfo"`
}

type AiReviewPoliticalAsrTaskInput struct {
	// 鉴别涉及令人不适宜信息的模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPoliticalAsrTaskOutput struct {
	// Asr 文字涉及令人不适宜的信息、违规评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Asr 文字涉及令人不适宜的信息、违规结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Asr 文字有涉及令人不适宜的信息、违规嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Asr 文字有涉及令人不适宜的信息、违规嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达 SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// Asr 文字有涉及令人不适宜的信息、违规嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPoliticalOcrTaskInput struct {
	// 鉴别涉及令人不适宜信息的模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPoliticalOcrTaskOutput struct {
	// Ocr 文字涉及令人不适宜的信息、违规评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Ocr 文字涉及令人不适宜的信息、违规结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Ocr 文字有涉及令人不适宜的信息、违规嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Ocr 文字有涉及令人不适宜的信息、违规嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达 SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// Ocr 文字有涉及令人不适宜的信息、违规嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPoliticalTaskInput struct {
	// 鉴别涉及令人不适宜信息的模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPoliticalTaskOutput struct {
	// 视频涉及令人不适宜信息的评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 涉及令人不适宜信息的结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 视频涉及令人不适宜信息的结果标签。音视频审核模板[画面鉴政任务控制参数](https://cloud.tencent.com/document/api/266/31773#PoliticalImgReviewTemplateInfo)里 LabelSet 参数与此参数取值范围的对应关系：
	// violation_photo：
	// <li>violation_photo：违规图标。</li>
	// 其他（即 politician/entertainment/sport/entrepreneur/scholar/celebrity/military）：
	// <li>politician：相关人物。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 有涉及令人不适宜信息嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewPoliticalSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 有涉及令人不适宜的信息嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达 SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 有涉及令人不适宜的信息嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPornAsrTaskInput struct {
	// 鉴别涉及令人反感的信息的模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPornAsrTaskOutput struct {
	// Asr 文字涉及令人反感的信息的评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Asr 文字涉及令人反感的信息的结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Asr 文字有涉及令人反感的信息的嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Asr 文字有涉及令人反感的信息的嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达 SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// Asr 文字有涉及令人反感的信息的嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPornOcrTaskInput struct {
	// 鉴别涉及令人反感的信息的模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPornOcrTaskOutput struct {
	// Ocr 文字涉及令人反感的信息的评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Ocr 文字涉及令人反感的信息的结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Ocr 文字有涉及令人反感的信息的嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Ocr 文字有涉及令人反感的信息的嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达 SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// Ocr 文字有涉及令人反感的信息的嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPornTaskInput struct {
	// 鉴别涉及令人反感的信息的模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPornTaskOutput struct {
	// 视频鉴别涉及令人反感的信息的评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 鉴别涉及令人反感的信息的结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 视频鉴别涉及令人反感的信息的结果标签，取值范围：
	// <li>porn：色情。</li>
	// <li>sexy：性感。</li>
	// <li>vulgar：低俗。</li>
	// <li>intimacy：亲密行为。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 有涉及令人反感的信息的嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 涉及令人反感的信息的嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 涉及令人反感的信息的嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewProhibitedAsrTaskInput struct {
	// 鉴违禁模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewProhibitedAsrTaskOutput struct {
	// Asr 文字涉违禁评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Asr 文字涉违禁结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Asr 文字有涉违禁嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Asr 文字有涉违禁嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达 SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// Asr 文字有涉违禁嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewProhibitedOcrTaskInput struct {
	// 鉴违禁模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewProhibitedOcrTaskOutput struct {
	// Ocr 文字涉违禁评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Ocr 文字涉违禁结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Ocr 文字有涉违禁嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Ocr 文字有涉违禁嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达 SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// Ocr 文字有涉违禁嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewTaskPoliticalAsrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核 Asr 文字涉及令人不适宜信息的任务输入。
	Input *AiReviewPoliticalAsrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核 Asr 文字涉及令人不适宜信息的任务输出。
	Output *AiReviewPoliticalAsrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音视频审核 Asr 文字涉及令人不适宜信息的任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTaskPoliticalOcrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核 Ocr 文字涉及令人不适宜信息的任务输入。
	Input *AiReviewPoliticalOcrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核 Ocr 文字涉及令人不适宜信息的任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPoliticalOcrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音视频审核 Ocr 文字涉及令人不适宜信息的任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTaskPoliticalResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核涉及令人不适宜信息的任务输入。
	Input *AiReviewPoliticalTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核涉及令人不适宜信息的任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPoliticalTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音视频审核涉及令人不适宜信息的任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTaskPornAsrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核 Asr 文字涉及令人反感的信息的任务输入。
	Input *AiReviewPornAsrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核 Asr 文字涉及令人反感的信息的任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPornAsrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音视频审核 Asr 文字涉及令人反感的信息的任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTaskPornOcrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核 Ocr 文字涉及令人反感的信息的任务输入。
	Input *AiReviewPornOcrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Ocr 文字音视频审核涉及令人反感的信息的任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPornOcrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// Ocr 文字音视频审核涉及令人反感的信息的任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTaskPornResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核涉及令人反感的信息的任务输入。
	Input *AiReviewPornTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核涉及令人反感的信息的任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPornTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音视频审核涉及令人反感的信息的任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTaskProhibitedAsrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核 Asr 文字鉴违禁任务输入。
	Input *AiReviewProhibitedAsrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核 Asr 文字鉴违禁任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewProhibitedAsrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音视频审核 Asr 文字鉴违禁任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTaskProhibitedOcrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核 Ocr 文字鉴违禁任务输入。
	Input *AiReviewProhibitedOcrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核 Ocr 文字鉴违禁任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewProhibitedOcrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音视频审核 Ocr 文字鉴违禁任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTaskTerrorismOcrResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核 Ocr 文字涉及令人不安全的信息的任务输入。
	Input *AiReviewTerrorismOcrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核 Ocr 文字涉及令人不安全的信息的任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewTerrorismOcrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音视频审核 Ocr 文字涉及令人不安全的信息的任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTaskTerrorismResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核涉及令人不安全的信息的任务输入。
	Input *AiReviewTerrorismTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核涉及令人不安全的信息的任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewTerrorismTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音视频审核涉及令人不安全的信息的任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiReviewTerrorismOcrTaskInput struct {
	// 鉴别涉及令人不安全的信息的模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewTerrorismOcrTaskOutput struct {
	// Ocr 文字有涉及令人不安全信息的评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Ocr 文字有涉及令人不安全信息的结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Ocr 文字有涉及令人不安全信息嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Ocr 文字有涉及令人不安全信息嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达 SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// Ocr 文字有涉及令人不安全信息嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewTerrorismTaskInput struct {
	// 鉴别涉及令人不安全的信息的模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewTerrorismTaskOutput struct {
	// 视频暴恐评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 暴恐结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 视频暴恐结果标签，取值范围：
	// <li>guns：武器枪支。</li>
	// <li>crowd：人群聚集。</li>
	// <li>police：警察部队。</li>
	// <li>bloody：血腥画面。</li>
	// <li>banners：暴恐旗帜。</li>
	// <li>militant：武装分子。</li>
	// <li>explosion：爆炸火灾。</li>
	// <li>terrorists：暴恐人物。</li>
	// <li>scenario：暴恐画面。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 有暴恐嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 暴恐嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 暴恐嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiSampleFaceInfo struct {
	// 人脸图片 ID。
	FaceId *string `json:"FaceId,omitnil,omitempty" name:"FaceId"`

	// 人脸图片地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type AiSampleFaceOperation struct {
	// 操作类型，可选值：add（添加）、delete（删除）、reset（重置）。重置操作将清空该人物已有人脸数据，并添加 FaceContents 指定人脸数据。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 人脸 ID 集合，当 Type为delete 时，该字段必填。
	FaceIds []*string `json:"FaceIds,omitnil,omitempty" name:"FaceIds"`

	// 人脸图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串集合，仅支持 jpeg、png 图片格式。
	// <li>当 Type为add 或 reset 时，该字段必填；</li>
	// <li>数组长度限制：5 张图片。</li>
	// 注意：图片必须是单人像正面人脸较清晰的照片，像素不低于 200*200。
	FaceContents []*string `json:"FaceContents,omitnil,omitempty" name:"FaceContents"`
}

type AiSampleFailFaceInfo struct {
	// 对应入参 FaceContents 中错误图片下标，从 0 开始。
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 错误码，取值：
	// <li>0：成功；</li>
	// <li>其他：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误描述。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type AiSamplePerson struct {
	// 人物 ID。
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 人物名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 人物描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 人脸信息。
	FaceInfoSet []*AiSampleFaceInfo `json:"FaceInfoSet,omitnil,omitempty" name:"FaceInfoSet"`

	// 人物标签。
	TagSet []*string `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 应用场景。
	UsageSet []*string `json:"UsageSet,omitnil,omitempty" name:"UsageSet"`

	// 创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AiSampleTagOperation struct {
	// 操作类型，可选值：add（添加）、delete（删除）、reset（重置）。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 标签，长度限制：128 个字符。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type AiSampleWord struct {
	// 关键词。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 关键词标签。
	TagSet []*string `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 关键词应用场景。
	UsageSet []*string `json:"UsageSet,omitnil,omitempty" name:"UsageSet"`

	// 创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AiSampleWordInfo struct {
	// 关键词，长度限制：20 个字符。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 关键词标签
	// <li>数组长度限制：20 个标签；</li>
	// <li>单个标签长度限制：128 个字符。</li>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type AnimatedGraphicTaskInput struct {
	// 视频转动图模板 ID
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 动图在视频中的起始时间偏移，单位为秒。
	// <li>不填或填0，表示从视频的起始位置开始；</li>
	// <li>当数值大于0时（假设为 n），表示从视频的第 n 秒位置开始；</li>
	// <li>当数值小于0时（假设为 -n），表示从视频结束 n 秒前的位置开始。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的终止时间偏移，单位为秒。
	// <li>不填或填0，表示持续到视频的末尾终止；</li>
	// <li>当数值大于0时（假设为 n），表示持续到视频第 n 秒时终止；</li>
	// <li>当数值小于0时（假设为 -n），表示持续到视频结束 n 秒前终止。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type AnimatedGraphicsTemplate struct {
	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 转动图模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 转动图模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 动图格式。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 帧率。
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 图片质量。
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ApplyUploadRequestParams struct {
	// 媒体类型，可选值请参考 [上传能力综述](/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	MediaType *string `json:"MediaType,omitnil,omitempty" name:"MediaType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 媒体名称。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 封面类型，可选值请参考 [上传能力综述](/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	CoverType *string `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// 媒体后续任务处理操作，即完成媒体上传后，可自动发起任务流操作。参数值为任务流模板名，云点播支持 [创建任务流模板](/document/product/266/33819) 并为模板命名。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 媒体文件过期时间，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 指定上传园区，仅适用于对上传地域有特殊需求的用户。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，[上传完成回调](/document/product/266/7830) 将返回该字段值，最长 250 个字符。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`

	// 会话上下文，用于透传用户请求信息，当指定 Procedure 参数后，[任务流状态变更回调](/document/product/266/9636) 将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type ApplyUploadRequest struct {
	*tchttp.BaseRequest
	
	// 媒体类型，可选值请参考 [上传能力综述](/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	MediaType *string `json:"MediaType,omitnil,omitempty" name:"MediaType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 媒体名称。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 封面类型，可选值请参考 [上传能力综述](/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	CoverType *string `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// 媒体后续任务处理操作，即完成媒体上传后，可自动发起任务流操作。参数值为任务流模板名，云点播支持 [创建任务流模板](/document/product/266/33819) 并为模板命名。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 媒体文件过期时间，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 指定上传园区，仅适用于对上传地域有特殊需求的用户。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，[上传完成回调](/document/product/266/7830) 将返回该字段值，最长 250 个字符。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`

	// 会话上下文，用于透传用户请求信息，当指定 Procedure 参数后，[任务流状态变更回调](/document/product/266/9636) 将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *ApplyUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaType")
	delete(f, "SubAppId")
	delete(f, "MediaName")
	delete(f, "CoverType")
	delete(f, "Procedure")
	delete(f, "ExpireTime")
	delete(f, "StorageRegion")
	delete(f, "ClassId")
	delete(f, "SourceContext")
	delete(f, "SessionContext")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyUploadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyUploadResponseParams struct {
	// 存储桶，用于上传接口 URL 的 bucket_name。
	StorageBucket *string `json:"StorageBucket,omitnil,omitempty" name:"StorageBucket"`

	// 存储园区，用于上传接口 Host 的 Region。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 点播会话，用于确认上传接口的参数 VodSessionKey。
	VodSessionKey *string `json:"VodSessionKey,omitnil,omitempty" name:"VodSessionKey"`

	// 媒体存储路径，用于上传接口存储媒体的对象键（Key）。
	MediaStoragePath *string `json:"MediaStoragePath,omitnil,omitempty" name:"MediaStoragePath"`

	// 封面存储路径，用于上传接口存储封面的对象键（Key）。
	CoverStoragePath *string `json:"CoverStoragePath,omitnil,omitempty" name:"CoverStoragePath"`

	// 临时凭证，用于上传接口的权限验证。
	TempCertificate *TempCertificate `json:"TempCertificate,omitnil,omitempty" name:"TempCertificate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyUploadResponse struct {
	*tchttp.BaseResponse
	Response *ApplyUploadResponseParams `json:"Response"`
}

func (r *ApplyUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AreaCLSTargetInfo struct {
	// 日志集所属地区：
	// <li>ap-guangzhou：广州；</li>
	// <li>ap-beijing：北京；</li>
	// <li>ap-chengdu：成都；</li>
	// <li>ap-chongqing：重庆；</li>
	// <li>ap-nanjing：南京；</li>
	// <li>ap-shanghai：上海；</li>
	// <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`

	// 投递的目标主题 ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 投递的目标集 ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志投递状态。
	//  ON：启用；
	//  OFF：停用。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ArtifactRepairInfo struct {
	// 去伪影（毛刺）控制开关，可选值：
	// <li>ON：开启去伪影（毛刺）；</li>
	// <li>OFF：关闭去伪影（毛刺）。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 去伪影（毛刺）类型，仅当去伪影（毛刺）控制开关为 ON 时有效，可选值：
	// <li>weak：轻去伪影（毛刺）；</li>
	// <li>strong：强去伪影（毛刺）。</li>
	// 默认值：weak。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type AsrFullTextConfigureInfo struct {
	// 语音全文识别任务开关，可选值：
	// <li>ON：开启智能语音全文识别任务；</li>
	// <li>OFF：关闭智能语音全文识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 生成的字幕文件格式列表，不填或者填空数组表示不生成字幕文件，可选值：
	// <li>vtt：生成 WebVTT 字幕文件；</li>
	// <li>srt：生成 SRT 字幕文件。</li><font color=red>注意：</font>云点播媒资信息仅支持添加 vtt 字幕，因此当且仅当 SubtitleFormats 包含 vtt 时，云点播将生成的字幕添加到媒资。
	SubtitleFormats []*string `json:"SubtitleFormats,omitnil,omitempty" name:"SubtitleFormats"`

	// 生成的字幕文件格式，不填或者填空字符串表示不生成字幕文件，可选值：
	// <li>vtt：生成 WebVTT 字幕文件；</li>
	// <li>srt：生成 SRT 字幕文件。</li>
	// <font color='red'>注意：此字段已废弃，建议使用 SubtitleFormats。</font>
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`

	// 媒体源语言，取值范围：
	// <li>zh：中文普通话；</li>
	// <li>en：英语；</li>
	// <li>ja：日语；</li>
	// <li>zh-ca：粤语。</li>
	// <font color=red>注意：</font> 填空字符串，或者不填该参数，则自动识别（效果较难保证，推荐填写原始媒体对应的语言，以提高识别的准确率）。
	SrcLanguage *string `json:"SrcLanguage,omitnil,omitempty" name:"SrcLanguage"`

	// 指定字幕名称，长度限制：64 个字符。该值将用于播放器展示，若不填则云点播自动生成。
	// <font color=red>注意：</font>仅当 SubtitleFormats 包含 vtt 时，该字段有效。
	SubtitleName *string `json:"SubtitleName,omitnil,omitempty" name:"SubtitleName"`
}

type AsrFullTextConfigureInfoForUpdate struct {
	// 语音全文识别任务开关，可选值：
	// <li>ON：开启智能语音全文识别任务；</li>
	// <li>OFF：关闭智能语音全文识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 字幕格式列表操作信息。
	SubtitleFormatsOperation *SubtitleFormatsOperation `json:"SubtitleFormatsOperation,omitnil,omitempty" name:"SubtitleFormatsOperation"`

	// 生成的字幕文件格式，<font color='red'>填空字符串</font>表示不生成字幕文件，可选值：
	// <li>vtt：生成 WebVTT 字幕文件；</li>
	// <li>srt：生成 SRT 字幕文件。</li>
	// <font color='red'>注意：此字段已废弃，建议使用 SubtitleFormatsOperation。</font>
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`

	// 媒体源语言，取值范围：
	// <li>zh：中文普通话；</li>
	// <li>en：英语；</li>
	// <li>ja：日语；</li>
	// <li>zh-ca：粤语。</li>
	SrcLanguage *string `json:"SrcLanguage,omitnil,omitempty" name:"SrcLanguage"`

	// 指定字幕名称，长度限制：64 个字符。该值将用于播放器展示。
	SubtitleName *string `json:"SubtitleName,omitnil,omitempty" name:"SubtitleName"`
}

type AsrTranslateConfigureInfo struct {
	// 语音翻译任务开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li><font color=red>注意：</font>语音翻译任务本身会返回 ASR 全文识别结果，为避免重复收费，因此禁止同时开启语音翻译和 ASR 全文识别功能项。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 媒体源语言，当 Switch 为 ON 时，此参数必填。取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>ja：日文；</li>
	// <li>ko：韩文；</li>
	// <li>vi：越南语；</li>
	// <li>ms：马来语；</li>
	// <li>th：泰语；</li>
	// <li>pt：葡萄牙语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ar：阿拉伯语；</li>
	// <li>es：西班牙语；</li>
	// <li>hi：印地语；</li>
	// <li>fr：法语。</li>
	SrcLanguage *string `json:"SrcLanguage,omitnil,omitempty" name:"SrcLanguage"`

	// 翻译目标语言，当 Switch 为 ON 时，此参数必填。
	// 当 SrcLanguage 为 zh（中文）时，取值范围：
	// <li>en：英文；</li>
	// <li>ja：日文；</li>
	// <li>ko：韩文；</li>
	// <li>fr：法语；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语；</li>
	// <li>vi：越南语；</li>
	// <li>id：印尼语；</li>
	// <li>th：泰语；</li>
	// <li>ms：马来语。</li>
	// 当 SrcLanguage 为 en（英文）时，取值范围：
	// <li>zh：中文；</li>
	// <li>ja：日文；</li>
	// <li>ko：韩文；</li>
	// <li>fr：法语；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语；</li>
	// <li>vi：越南语；</li>
	// <li>id：印尼语；</li>
	// <li>th：泰语；</li>
	// <li>ms：马来语；</li>
	// <li>ar：阿拉伯语；</li>
	// <li>hi：印地语。</li>
	// 当 SrcLanguage 为 ja（日文）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>ko：韩文。</li>
	// 当 SrcLanguage 为 ko（韩文）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>ja：日文。</li>
	// 当 SrcLanguage 为 vi（越南语）或 ms（马来语）或 th（泰语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文。</li>
	// 当 SrcLanguage 为 pt（葡萄牙语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>fr：法语；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语。</li>
	// 当 SrcLanguage 为 tr（土耳其语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>fr：法语；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语。</li>
	// 当 SrcLanguage 为 es（西班牙语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>fr：法语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语。</li>
	// 当 SrcLanguage 为 ar（阿拉伯语）或 hi（印地语）时，取值范围：
	// <li>en：英文。</li>
	// 当 SrcLanguage 为 fr（法语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语。</li>
	DstLanguage *string `json:"DstLanguage,omitnil,omitempty" name:"DstLanguage"`

	// 生成的字幕文件格式列表，不填或者填空数组表示不生成字幕文件，可选值：
	// <li>vtt：生成 WebVTT 字幕文件；</li>
	// <li>srt：生成 SRT 字幕文件。</li><font color=red>注意：</font> 云点播媒资信息仅支持添加 vtt 字幕，因此当且仅当 SubtitleFormats 包含 vtt 时，云点播将生成的字幕添加到媒资。
	SubtitleFormats []*string `json:"SubtitleFormats,omitnil,omitempty" name:"SubtitleFormats"`

	// 指定字幕名称，长度限制：64 个字符。该值将用于播放器展示，若不填则云点播自动生成。
	// <font color=red>注意：</font>仅当 SubtitleFormats 包含 vtt 时，该字段有效。
	SubtitleName *string `json:"SubtitleName,omitnil,omitempty" name:"SubtitleName"`
}

type AsrTranslateConfigureInfoForUpdate struct {
	// 语音翻译任务开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 媒体源语言，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>ja：日文；</li>
	// <li>ko：韩文；</li>
	// <li>vi：越南语；</li>
	// <li>ms：马来语；</li>
	// <li>th：泰语；</li>
	// <li>pt：葡萄牙语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ar：阿拉伯语；</li>
	// <li>es：西班牙语；</li>
	// <li>hi：印地语；</li>
	// <li>fr：法语。</li>
	SrcLanguage *string `json:"SrcLanguage,omitnil,omitempty" name:"SrcLanguage"`

	// 翻译目标语言。
	// 当 SrcLanguage 为 zh（中文）时，取值范围：
	// <li>en：英文；</li>
	// <li>ja：日文；</li>
	// <li>ko：韩文；</li>
	// <li>fr：法语；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语；</li>
	// <li>vi：越南语；</li>
	// <li>id：印尼语；</li>
	// <li>th：泰语；</li>
	// <li>ms：马来语。</li>
	// 当 SrcLanguage 为 en（英文）时，取值范围：
	// <li>zh：中文；</li>
	// <li>ja：日文；</li>
	// <li>ko：韩文；</li>
	// <li>fr：法语；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语；</li>
	// <li>vi：越南语；</li>
	// <li>id：印尼语；</li>
	// <li>th：泰语；</li>
	// <li>ms：马来语；</li>
	// <li>ar：阿拉伯语；</li>
	// <li>hi：印地语。</li>
	// 当 SrcLanguage 为 ja（日文）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>ko：韩文。</li>
	// 当 SrcLanguage 为 ko（韩文）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>ja：日文。</li>
	// 当 SrcLanguage 为 vi（越南语）或 ms（马来语）或 th（泰语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文。</li>
	// 当 SrcLanguage 为 pt（葡萄牙语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>fr：法语；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语。</li>
	// 当 SrcLanguage 为 tr（土耳其语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>fr：法语；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语。</li>
	// 当 SrcLanguage 为 es（西班牙语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>fr：法语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语。</li>
	// 当 SrcLanguage 为 ar（阿拉伯语）或 hi（印地语）时，取值范围：
	// <li>en：英文。</li>
	// 当 SrcLanguage 为 fr（法语）时，取值范围：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>es：西班牙语；</li>
	// <li>it：意大利语；</li>
	// <li>de：德语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ru：俄语；</li>
	// <li>pt：葡萄牙语。</li>
	DstLanguage *string `json:"DstLanguage,omitnil,omitempty" name:"DstLanguage"`

	// 字幕格式列表操作信息。
	SubtitleFormatsOperation *SubtitleFormatsOperation `json:"SubtitleFormatsOperation,omitnil,omitempty" name:"SubtitleFormatsOperation"`

	// 指定字幕名称，长度限制：64 个字符。该值将用于播放器展示。
	SubtitleName *string `json:"SubtitleName,omitnil,omitempty" name:"SubtitleName"`
}

type AsrWordsConfigureInfo struct {
	// 语音关键词识别任务开关，可选值：
	// <li>ON：开启语音关键词识别任务；</li>
	// <li>OFF：关闭语音关键词识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`
}

type AsrWordsConfigureInfoForUpdate struct {
	// 语音关键词识别任务开关，可选值：
	// <li>ON：开启语音关键词识别任务；</li>
	// <li>OFF：关闭语音关键词识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`
}

// Predefined struct for user
type AttachMediaSubtitlesRequestParams struct {
	// 媒体文件唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 操作。取值如下：
	// <li>Attach：关联字幕。</li>
	// <li>Detach：解除关联字幕。</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// [转自适应码流模板号](https://cloud.tencent.com/document/product/266/34071#zsy)。
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitnil,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// 字幕的唯一标识。
	SubtitleIds []*string `json:"SubtitleIds,omitnil,omitempty" name:"SubtitleIds"`

	// 默认字幕的唯一标识。不填则不设置默认字幕。
	DefaultSubtitleId *string `json:"DefaultSubtitleId,omitnil,omitempty" name:"DefaultSubtitleId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type AttachMediaSubtitlesRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 操作。取值如下：
	// <li>Attach：关联字幕。</li>
	// <li>Detach：解除关联字幕。</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// [转自适应码流模板号](https://cloud.tencent.com/document/product/266/34071#zsy)。
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitnil,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// 字幕的唯一标识。
	SubtitleIds []*string `json:"SubtitleIds,omitnil,omitempty" name:"SubtitleIds"`

	// 默认字幕的唯一标识。不填则不设置默认字幕。
	DefaultSubtitleId *string `json:"DefaultSubtitleId,omitnil,omitempty" name:"DefaultSubtitleId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *AttachMediaSubtitlesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachMediaSubtitlesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "Operation")
	delete(f, "AdaptiveDynamicStreamingDefinition")
	delete(f, "SubtitleIds")
	delete(f, "DefaultSubtitleId")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachMediaSubtitlesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachMediaSubtitlesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachMediaSubtitlesResponse struct {
	*tchttp.BaseResponse
	Response *AttachMediaSubtitlesResponseParams `json:"Response"`
}

func (r *AttachMediaSubtitlesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachMediaSubtitlesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AudioDenoiseInfo struct {
	// 音频降噪控制开关，可选值：
	// <li>ON：开启音频降噪；</li>
	// <li>OFF：关闭音频降噪。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 音频降噪类型，仅当音频降噪控制开关为 ON 时有效，可选值：
	// <li>weak：轻音频降噪；</li>
	// <li>normal：正常音频降噪；</li>
	// <li>strong：强音频降噪。</li>
	// 默认值：weak。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	// <li>libmp3lame：更适合 flv；</li>
	// <li>mp2。</li>
	// 当外层参数 Container 为 hls 时，可选值为：
	// <li>libfdk_aac。</li>
	// 当外层参数 Format 为 HLS 或 MPEG-DASH 时，可选值为：
	// <li>libfdk_aac。</li>
	// 当外层参数 Container 为 wav 时，可选值为：
	// <li>pcm16。</li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。
	// 当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 音频流的采样率，可选值：
	// <li>16000，仅当 Codec 为 pcm16 时可选。</li>
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频通道，可选值：
	// <li>1：单通道</li>
	// <li>2：双通道</li>
	// <li>6：立体声</li>
	// <li>0：音频声道数和原始音频保持一致</li>
	// 当媒体的封装格式是音频格式时（flac，ogg，mp3，m4a）时，声道数不允许设为立体声。
	// 默认值：2。
	AudioChannel *int64 `json:"AudioChannel,omitnil,omitempty" name:"AudioChannel"`
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
	// <li>libfdk_aac。</li>
	// 当外层参数 Format 为 HLS 或 MPEG-DASH 时，可选值为：
	// <li>libfdk_aac。</li>
	// 当外层参数 Container 为 wav 时，可选值为：
	// <li>pcm16。</li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。 当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 音频流的采样率，可选值：
	// <li>16000，仅当 Codec 为 pcm16 时可选。</li>
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频通道，可选值：
	// <li>1：单通道</li>
	// <li>2：双通道</li>
	// <li>6：立体声</li>
	// <li>0：音频声道数和原始音频保持一致</li>
	// 当媒体的封装格式是音频格式时（flac，ogg，mp3，m4a）时，声道数不允许设为立体声。
	AudioChannel *int64 `json:"AudioChannel,omitnil,omitempty" name:"AudioChannel"`
}

type AudioTrackItem struct {
	// 音频片段的媒体素材来源，可以是：
	// <li>点播的媒体文件 ID；</li>
	// <li>其他媒体文件的下载 URL。</li>
	// 注意：当使用其他媒体文件的下载 URL 作为素材来源，且开启了访问控制（如防盗链）时，需要在 URL 携带访问控制参数（如防盗链签名）。
	SourceMedia *string `json:"SourceMedia,omitnil,omitempty" name:"SourceMedia"`

	// 音频片段取自素材文件的起始时间，单位为秒。0 表示从素材开始位置截取。默认为0。
	SourceMediaStartTime *float64 `json:"SourceMediaStartTime,omitnil,omitempty" name:"SourceMediaStartTime"`

	// 音频片段的时长，单位为秒。默认和素材本身长度一致，表示截取全部素材。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 音频片段目标时长，单位为秒。
	// <li>当 TargetDuration 不填或填0时，表示目标时长和 Duration 一致；</li>
	// <li>当 TargetDuration 取大于0的值时，将对音频片段做快进或慢放等处理，使得输出片段的时长等于 TargetDuration。</li>
	TargetDuration *float64 `json:"TargetDuration,omitnil,omitempty" name:"TargetDuration"`

	// 对音频片段进行的操作，如音量调节等。
	AudioOperations []*AudioTransform `json:"AudioOperations,omitnil,omitempty" name:"AudioOperations"`
}

type AudioTransform struct {
	// 音频操作类型，取值有：
	// <li>Volume：音量调节。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 音量调节参数， 当 Type = Volume 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeParam *AudioVolumeParam `json:"VolumeParam,omitnil,omitempty" name:"VolumeParam"`
}

type AudioVolumeParam struct {
	// 是否静音，取值范围0或1。
	// <li>0表示不静音。</li>
	// <li>1表示静音。</li>
	// 默认是0。
	Mute *int64 `json:"Mute,omitnil,omitempty" name:"Mute"`

	// 音频增益，取值范围0~10。
	// <li>大于1表示增加音量。</li>
	// <li>小于1表示降低音量。</li>
	// <li>0和1：表示不改变。</li>
	// 默认是0。
	Gain *float64 `json:"Gain,omitnil,omitempty" name:"Gain"`
}

type BlackWhiteEdgeConfigureInfo struct {
	// 视频画面黑边、白边、黑屏、白屏检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type BlackWhiteEdgeConfigureInfoForUpdate struct {
	// 视频画面黑边、白边、黑屏、白屏检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type BlurConfigureInfo struct {
	// 视频画面模糊检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type BlurConfigureInfoForUpdate struct {
	// 视频画面模糊检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type CLSLogsetInfo struct {
	// 日志集 ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`
}

type CLSTopicInfo struct {
	// 日志主题 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志集 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

type Canvas struct {
	// 背景颜色，取值有：
	// <li>Black：黑色背景</li>
	// <li>White：白色背景</li>
	// 默认值：Black。
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`

	// 画布宽度，即输出视频的宽度，取值范围：0~ 3840，单位：px。
	// 默认值：0，表示和第一个视频轨的第一个视频片段的视频宽度一致。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 画布高度，即输出视频的高度（或长边），取值范围：0~ 3840，单位：px。
	// 默认值：0，表示和第一个视频轨的第一个视频片段的视频高度一致。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type CdnLogInfo struct {
	// 日志所属日期， 格式为：yyyy-MM-dd ，如2018-03-01。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 日志名称，格式为：日期小时-域名
	// 如 2018120101-test.vod2.mqcloud.com。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志下载链接，24小时内下载有效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 日志起始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ClassificationConfigureInfo struct {
	// 智能分类任务开关，可选值：
	// <li>ON：开启智能分类任务；</li>
	// <li>OFF：关闭智能分类任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ClassificationConfigureInfoForUpdate struct {
	// 智能分类任务开关，可选值：
	// <li>ON：开启智能分类任务；</li>
	// <li>OFF：关闭智能分类任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ClipFileInfo2017 struct {
	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误描述。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 输出目标文件的文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 输出目标文件的文件地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 输出目标文件的文件类型。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type ClipTask2017 struct {
	// 视频剪辑任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 视频剪辑任务源文件 ID。
	SrcFileId *string `json:"SrcFileId,omitnil,omitempty" name:"SrcFileId"`

	// 视频剪辑输出的文件信息。
	FileInfo *ClipFileInfo2017 `json:"FileInfo,omitnil,omitempty" name:"FileInfo"`
}

type ColorEnhanceInfo struct {
	// 色彩增强控制开关，可选值：
	// <li>ON：开启综合增强；</li>
	// <li>OFF：关闭综合增强。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 色彩增强类型，仅当色彩增强控制开关为 ON 时有效，可选值：
	// <li>weak：轻色彩增强；</li>
	// <li>normal：正常色彩增强；</li>
	// <li>strong：强色彩增强。</li>
	// 默认值：weak。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type CommitUploadRequestParams struct {
	// 点播会话，取申请上传接口的返回值 VodSessionKey。
	VodSessionKey *string `json:"VodSessionKey,omitnil,omitempty" name:"VodSessionKey"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type CommitUploadRequest struct {
	*tchttp.BaseRequest
	
	// 点播会话，取申请上传接口的返回值 VodSessionKey。
	VodSessionKey *string `json:"VodSessionKey,omitnil,omitempty" name:"VodSessionKey"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *CommitUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VodSessionKey")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CommitUploadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommitUploadResponseParams struct {
	// 媒体文件的唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 媒体播放地址。
	MediaUrl *string `json:"MediaUrl,omitnil,omitempty" name:"MediaUrl"`

	// 媒体封面地址。
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CommitUploadResponse struct {
	*tchttp.BaseResponse
	Response *CommitUploadResponseParams `json:"Response"`
}

func (r *CommitUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplexAdaptiveDynamicStreamingTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 自适应码流任务的执行状态与结果，每个元素对应一个自适应码流模版。
	ComplexAdaptiveDynamicStreamingTaskResultSet []*ComplexAdaptiveDynamicStreamingTaskResult `json:"ComplexAdaptiveDynamicStreamingTaskResultSet,omitnil,omitempty" name:"ComplexAdaptiveDynamicStreamingTaskResultSet"`
}

type ComplexAdaptiveDynamicStreamingTaskInput struct {
	// 自适应码流参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamPara *ComplexAdaptiveDynamicStreamingTaskStreamPara `json:"StreamPara,omitnil,omitempty" name:"StreamPara"`
}

type ComplexAdaptiveDynamicStreamingTaskOutput struct {
	// 自适应码流模版 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 自适应码流打包格式。可选值：
	// <li>HLS；</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// DRM 方案类型。可选值：
	// <li>空字符串：无加密；</li>
	// <li>SimpleAES；</li>
	// <li>Widevine；</li>
	// <li>FairPlay。</li>
	DrmType *string `json:"DrmType,omitnil,omitempty" name:"DrmType"`

	// 自适应码流的播放地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ComplexAdaptiveDynamicStreamingTaskResult struct {
	// 任务状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>SUCCESS：已完成；</li>
	// <li>FAIL：失败。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 转码进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 自适应码流任务的输入。
	Input *ComplexAdaptiveDynamicStreamingTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 自适应码流任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *ComplexAdaptiveDynamicStreamingTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type ComplexAdaptiveDynamicStreamingTaskStreamPara struct {
	// 自适应码流模版 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type ComposeMediaOutput struct {
	// 文件名称，最长 64 个字符。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 描述信息，最长 128 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 封装格式，可选值：mp4、mp3。其中，mp3 为纯音频文件。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 输出的视频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoStream *OutputVideoStream `json:"VideoStream,omitnil,omitempty" name:"VideoStream"`

	// 输出的音频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioStream *OutputAudioStream `json:"AudioStream,omitnil,omitempty" name:"AudioStream"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

// Predefined struct for user
type ComposeMediaRequestParams struct {
	// 输入的媒体轨道列表，包括视频、音频、图片等素材组成的多个轨道信息，其中：<li>输入的多个轨道在时间轴上和输出媒体文件的时间轴对齐；</li><li>时间轴上相同时间点的各个轨道的素材进行重叠，视频或者图片按轨道顺序进行图像的叠加，轨道顺序高的素材叠加在上面，音频素材进行混音；</li><li>视频、音频、图片，每一种类型的轨道最多支持 10 个。</li><li>所有类型的轨道上放置的媒体片段数量总和最多支持 500 个。</li>
	Tracks []*MediaTrack `json:"Tracks,omitnil,omitempty" name:"Tracks"`

	// 输出的媒体文件信息。
	Output *ComposeMediaOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 制作视频文件时使用的画布。
	Canvas *Canvas `json:"Canvas,omitnil,omitempty" name:"Canvas"`

	// 标识来源上下文，用于透传用户请求信息，在ComposeMediaComplete回调将返回该字段值，最长 1000个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type ComposeMediaRequest struct {
	*tchttp.BaseRequest
	
	// 输入的媒体轨道列表，包括视频、音频、图片等素材组成的多个轨道信息，其中：<li>输入的多个轨道在时间轴上和输出媒体文件的时间轴对齐；</li><li>时间轴上相同时间点的各个轨道的素材进行重叠，视频或者图片按轨道顺序进行图像的叠加，轨道顺序高的素材叠加在上面，音频素材进行混音；</li><li>视频、音频、图片，每一种类型的轨道最多支持 10 个。</li><li>所有类型的轨道上放置的媒体片段数量总和最多支持 500 个。</li>
	Tracks []*MediaTrack `json:"Tracks,omitnil,omitempty" name:"Tracks"`

	// 输出的媒体文件信息。
	Output *ComposeMediaOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 制作视频文件时使用的画布。
	Canvas *Canvas `json:"Canvas,omitnil,omitempty" name:"Canvas"`

	// 标识来源上下文，用于透传用户请求信息，在ComposeMediaComplete回调将返回该字段值，最长 1000个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *ComposeMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ComposeMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tracks")
	delete(f, "Output")
	delete(f, "SubAppId")
	delete(f, "Canvas")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ComposeMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ComposeMediaResponseParams struct {
	// 制作媒体文件的任务 ID，可以通过该 ID 查询制作任务（任务类型为 MakeMedia）的状态。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ComposeMediaResponse struct {
	*tchttp.BaseResponse
	Response *ComposeMediaResponseParams `json:"Response"`
}

func (r *ComposeMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ComposeMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComposeMediaTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 制作媒体文件任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 制作媒体文件任务的输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *ComposeMediaTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 制作媒体文件任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *ComposeMediaTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 输出视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type ComposeMediaTaskInput struct {
	// 输入的媒体轨道列表，包括视频、音频、图片等素材组成的多个轨道信息。
	Tracks []*MediaTrack `json:"Tracks,omitnil,omitempty" name:"Tracks"`

	// 制作视频文件时使用的画布。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Canvas *Canvas `json:"Canvas,omitnil,omitempty" name:"Canvas"`

	// 输出的媒体文件信息。
	Output *ComposeMediaOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type ComposeMediaTaskOutput struct {
	// 文件类型，例如 mp4、mp3 等。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 媒体文件播放地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件名称，最长 64 个字符。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type ConcatFileInfo2017 struct {
	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 视频拼接源文件的 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 视频拼接源文件的地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 视频拼接源文件的格式。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type ConcatTask2017 struct {
	// 视频拼接任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 视频拼接源文件信息。
	FileInfoSet []*ConcatFileInfo2017 `json:"FileInfoSet,omitnil,omitempty" name:"FileInfoSet"`
}

// Predefined struct for user
type ConfirmEventsRequestParams struct {
	// 事件句柄，即 [拉取事件通知](/document/product/266/33433) 接口输出参数中的 EventSet. EventHandle 字段。
	// 数组长度限制：16。
	EventHandles []*string `json:"EventHandles,omitnil,omitempty" name:"EventHandles"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type ConfirmEventsRequest struct {
	*tchttp.BaseRequest
	
	// 事件句柄，即 [拉取事件通知](/document/product/266/33433) 接口输出参数中的 EventSet. EventHandle 字段。
	// 数组长度限制：16。
	EventHandles []*string `json:"EventHandles,omitnil,omitempty" name:"EventHandles"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *ConfirmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventHandles")
	delete(f, "ExtInfo")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfirmEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfirmEventsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfirmEventsResponse struct {
	*tchttp.BaseResponse
	Response *ConfirmEventsResponseParams `json:"Response"`
}

func (r *ConfirmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContentReviewOcrResult struct {
	// Ocr 文字鉴别结果的评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Ocr 文字鉴别的结果建议，取值范围：
	// <li>pass；</li>
	// <li>review；</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Ocr 文字鉴别的嫌疑关键词列表。
	KeywordSet []*string `json:"KeywordSet,omitnil,omitempty" name:"KeywordSet"`

	// Ocr 文字鉴别的嫌疑文字出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type ContentReviewResult struct {
	// 结果类型，取值范围：
	// <li>Porn.Image：图片画面中的鉴别令人反感的信息结果；</li>
	// <li>Terrorism.Image：图片画面中的鉴别令人不安全的信息结果；</li>
	// <li>Political.Image：图片画面中的鉴别令人不适宜信息结果；</li>
	// <li>Porn.Ocr：图片 OCR 文字中的鉴别令人反感的信息结果；</li>
	// <li>Terrorism.Ocr：图片 OCR 文字中的鉴别令人不安全的信息结果；</li>
	// <li>Political.Ocr：图片 OCR 文字中的鉴别令人不适宜信息结果。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 图片画面中的鉴别令人反感的信息结果，当 Type 为 Porn.Image 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornImageResult *PornImageResult `json:"PornImageResult,omitnil,omitempty" name:"PornImageResult"`

	// 图片画面中的鉴别令人不安全的信息结果，当 Type 为 Terrorism.Image 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismImageResult *TerrorismImageResult `json:"TerrorismImageResult,omitnil,omitempty" name:"TerrorismImageResult"`

	// 图片画面中的鉴别令人不适宜信息结果，当 Type 为 Political.Image 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalImageResult *PoliticalImageResult `json:"PoliticalImageResult,omitnil,omitempty" name:"PoliticalImageResult"`

	// 图片 OCR 文字中的鉴别令人反感的信息结果，当 Type 为 Porn.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornOcrResult *ContentReviewOcrResult `json:"PornOcrResult,omitnil,omitempty" name:"PornOcrResult"`

	// 图片 OCR 中的鉴别令人不安全的信息结果，当 Type 为 Terrorism.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismOcrResult *ContentReviewOcrResult `json:"TerrorismOcrResult,omitnil,omitempty" name:"TerrorismOcrResult"`

	// 图片 OCR 文字中的鉴别令人不适宜信息结果，当 Type 为 Political.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalOcrResult *ContentReviewOcrResult `json:"PoliticalOcrResult,omitnil,omitempty" name:"PoliticalOcrResult"`
}

type ContentReviewTemplateItem struct {
	// 音视频审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 音视频审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音视频审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 鉴别涉及令人反感的信息的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// 鉴别涉及令人不安全的信息的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// 鉴别涉及令人不适宜的信息的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// 用户自定义音视频审核控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`

	// 音视频审核结果是否进入音视频审核墙（对音视频审核结果进行人工复核）的开关。
	// <li>ON：是；</li>
	// <li>OFF：否。</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitnil,omitempty" name:"ReviewWallSwitch"`

	// 截帧间隔，单位为秒。当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type CopyRightWatermarkInput struct {
	// 版权信息，最大长度为 200 个字符。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type CoverBySnapshotTaskInput struct {
	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 截图方式。包含：
	// <li>Time：依照时间点截图</li>
	// <li>Percent：依照百分比截图</li>
	PositionType *string `json:"PositionType,omitnil,omitempty" name:"PositionType"`

	// 截图位置：
	// <li>对于依照时间点截图，该值表示指定视频第几秒的截图作为封面</li>
	// <li>对于依照百分比截图，该值表示使用视频百分之多少的截图作为封面</li>
	PositionValue *float64 `json:"PositionValue,omitnil,omitempty" name:"PositionValue"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`
}

type CoverBySnapshotTaskOutput struct {
	// 封面 URL。
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`
}

type CoverConfigureInfo struct {
	// 智能封面任务开关，可选值：
	// <li>ON：开启智能封面任务；</li>
	// <li>OFF：关闭智能封面任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type CoverConfigureInfoForUpdate struct {
	// 智能封面任务开关，可选值：
	// <li>ON：开启智能封面任务；</li>
	// <li>OFF：关闭智能封面任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type CrashScreenConfigureInfo struct {
	// 视频画面花屏检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type CrashScreenConfigureInfoForUpdate struct {
	// 视频画面花屏检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type CreateAIAnalysisTemplateRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`

	// 智能精彩集锦任务控制参数。
	HighlightConfigure *HighlightsConfigureInfo `json:"HighlightConfigure,omitnil,omitempty" name:"HighlightConfigure"`
}

type CreateAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`

	// 智能精彩集锦任务控制参数。
	HighlightConfigure *HighlightsConfigureInfo `json:"HighlightConfigure,omitnil,omitempty" name:"HighlightConfigure"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ClassificationConfigure")
	delete(f, "TagConfigure")
	delete(f, "CoverConfigure")
	delete(f, "FrameTagConfigure")
	delete(f, "HighlightConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIAnalysisTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIAnalysisTemplateResponseParams struct {
	// 音视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频片头片尾识别控制参数。
	HeadTailConfigure *HeadTailConfigureInfo `json:"HeadTailConfigure,omitnil,omitempty" name:"HeadTailConfigure"`

	// 视频拆条识别控制参数。
	SegmentConfigure *SegmentConfigureInfo `json:"SegmentConfigure,omitnil,omitempty" name:"SegmentConfigure"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	AsrTranslateConfigure *AsrTranslateConfigureInfo `json:"AsrTranslateConfigure,omitnil,omitempty" name:"AsrTranslateConfigure"`

	// 物体识别控制参数。
	ObjectConfigure *ObjectConfigureInfo `json:"ObjectConfigure,omitnil,omitempty" name:"ObjectConfigure"`

	// 截帧间隔，单位为秒。当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`
}

type CreateAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频片头片尾识别控制参数。
	HeadTailConfigure *HeadTailConfigureInfo `json:"HeadTailConfigure,omitnil,omitempty" name:"HeadTailConfigure"`

	// 视频拆条识别控制参数。
	SegmentConfigure *SegmentConfigureInfo `json:"SegmentConfigure,omitnil,omitempty" name:"SegmentConfigure"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	AsrTranslateConfigure *AsrTranslateConfigureInfo `json:"AsrTranslateConfigure,omitnil,omitempty" name:"AsrTranslateConfigure"`

	// 物体识别控制参数。
	ObjectConfigure *ObjectConfigureInfo `json:"ObjectConfigure,omitnil,omitempty" name:"ObjectConfigure"`

	// 截帧间隔，单位为秒。当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "HeadTailConfigure")
	delete(f, "SegmentConfigure")
	delete(f, "FaceConfigure")
	delete(f, "OcrFullTextConfigure")
	delete(f, "OcrWordsConfigure")
	delete(f, "AsrFullTextConfigure")
	delete(f, "AsrWordsConfigure")
	delete(f, "AsrTranslateConfigure")
	delete(f, "ObjectConfigure")
	delete(f, "ScreenshotInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIRecognitionTemplateResponseParams struct {
	// 音视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <li>HLS；</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 自适应转码输出子流参数信息，最多输出10路子流。
	// 注意：各个子流的帧率必须保持一致；如果不一致，采用第一个子流的帧率作为输出帧率。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// DRM 方案类型，取值范围：
	// <li>SimpleAES</li>
	// <li>Widevine</li>
	// <li>FairPlay</li>
	// 默认值为空字符串，如果取值为空字符串，代表不对视频做 DRM 保护。
	DrmType *string `json:"DrmType,omitnil,omitempty" name:"DrmType"`

	// DRM 的密钥提供商，取值范围：
	// <li>SDMC：华曦达；</li>
	// <li>VOD：云点播。</li>
	// 默认为 VOD 。
	DrmKeyProvider *string `json:"DrmKeyProvider,omitnil,omitempty" name:"DrmKeyProvider"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	// 默认为否。
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	// 默认为否。
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 切片类型，当 Format 为 HLS 时有效，可选值：
	// <li>ts：ts 切片；</li>
	// <li>fmp4：fmp4 切片。</li>
	// 默认值：ts。
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
}

type CreateAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 自适应转码格式，取值范围：
	// <li>HLS；</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 自适应转码输出子流参数信息，最多输出10路子流。
	// 注意：各个子流的帧率必须保持一致；如果不一致，采用第一个子流的帧率作为输出帧率。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// DRM 方案类型，取值范围：
	// <li>SimpleAES</li>
	// <li>Widevine</li>
	// <li>FairPlay</li>
	// 默认值为空字符串，如果取值为空字符串，代表不对视频做 DRM 保护。
	DrmType *string `json:"DrmType,omitnil,omitempty" name:"DrmType"`

	// DRM 的密钥提供商，取值范围：
	// <li>SDMC：华曦达；</li>
	// <li>VOD：云点播。</li>
	// 默认为 VOD 。
	DrmKeyProvider *string `json:"DrmKeyProvider,omitnil,omitempty" name:"DrmKeyProvider"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	// 默认为否。
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	// 默认为否。
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 切片类型，当 Format 为 HLS 时有效，可选值：
	// <li>ts：ts 切片；</li>
	// <li>fmp4：fmp4 切片。</li>
	// 默认值：ts。
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "DrmType")
	delete(f, "DrmKeyProvider")
	delete(f, "DisableHigherVideoBitrate")
	delete(f, "DisableHigherVideoResolution")
	delete(f, "Comment")
	delete(f, "SegmentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdaptiveDynamicStreamingTemplateResponseParams struct {
	// 自适应转码模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 动图格式，取值为 gif 和 webp。默认为 gif。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 帧率，取值范围：[1, 30]，单位：Hz。
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 动图格式，取值为 gif 和 webp。默认为 gif。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
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
	delete(f, "SubAppId")
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
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateCLSLogsetRequestParams struct {
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`
}

type CreateCLSLogsetRequest struct {
	*tchttp.BaseRequest
	
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`
}

func (r *CreateCLSLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CLSRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCLSLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCLSLogsetResponseParams struct {
	// 日志集 ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCLSLogsetResponse struct {
	*tchttp.BaseResponse
	Response *CreateCLSLogsetResponseParams `json:"Response"`
}

func (r *CreateCLSLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCLSTopicRequestParams struct {
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`

	// 日志主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志集 ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

type CreateCLSTopicRequest struct {
	*tchttp.BaseRequest
	
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`

	// 日志主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志集 ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

func (r *CreateCLSTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CLSRegion")
	delete(f, "TopicName")
	delete(f, "LogsetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCLSTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCLSTopicResponseParams struct {
	// 日志主题 ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCLSTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateCLSTopicResponseParams `json:"Response"`
}

func (r *CreateCLSTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClassRequestParams struct {
	// 父类 ID，一级分类填写 -1。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 分类名称，长度限制：1-64 个字符。
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type CreateClassRequest struct {
	*tchttp.BaseRequest
	
	// 父类 ID，一级分类填写 -1。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 分类名称，长度限制：1-64 个字符。
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *CreateClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ParentId")
	delete(f, "ClassName")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClassResponseParams struct {
	// 分类 ID
	ClassId *uint64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClassResponse struct {
	*tchttp.BaseResponse
	Response *CreateClassResponseParams `json:"Response"`
}

func (r *CreateClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContentReviewTemplateRequestParams struct {
	// 音视频审核结果是否进入音视频审核墙（对识别结果进行人工复核）的开关。
	// <li>ON：是；</li>
	// <li>OFF：否。</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitnil,omitempty" name:"ReviewWallSwitch"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 令人反感的信息的控制参数。
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// 令人不安全的信息的控制参数。
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// 令人不适宜的控制参数。
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`

	// 截帧间隔，单位为秒。当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`
}

type CreateContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 音视频审核结果是否进入音视频审核墙（对识别结果进行人工复核）的开关。
	// <li>ON：是；</li>
	// <li>OFF：否。</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitnil,omitempty" name:"ReviewWallSwitch"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 令人反感的信息的控制参数。
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// 令人不安全的信息的控制参数。
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// 令人不适宜的控制参数。
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`

	// 截帧间隔，单位为秒。当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`
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
	delete(f, "ReviewWallSwitch")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "PornConfigure")
	delete(f, "TerrorismConfigure")
	delete(f, "PoliticalConfigure")
	delete(f, "ProhibitedConfigure")
	delete(f, "UserDefineConfigure")
	delete(f, "ScreenshotInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContentReviewTemplateResponseParams struct {
	// 音视频内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateDomainVerifyRecordRequestParams struct {
	// 需要接入点播的加速域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type CreateDomainVerifyRecordRequest struct {
	*tchttp.BaseRequest
	
	// 需要接入点播的加速域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *CreateDomainVerifyRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainVerifyRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainVerifyRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainVerifyRecordResponseParams struct {
	// DNS解析信息
	DNSVerifyInfo *DNSVerifyInfo `json:"DNSVerifyInfo,omitnil,omitempty" name:"DNSVerifyInfo"`

	// 文件验证信息
	FileVerifyInfo *FileVerifyInfo `json:"FileVerifyInfo,omitnil,omitempty" name:"FileVerifyInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainVerifyRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainVerifyRecordResponseParams `json:"Response"`
}

func (r *CreateDomainVerifyRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainVerifyRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnhanceMediaTemplateRequestParams struct {
	// 输出文件封装格式，可选值：mp4、flv、hls。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 音画质重生视频控制控制信息。
	RebuildVideoInfo *RebuildVideoInfo `json:"RebuildVideoInfo,omitnil,omitempty" name:"RebuildVideoInfo"`

	// 音画质重生音频控制控制信息。
	RebuildAudioInfo *RebuildAudioInfo `json:"RebuildAudioInfo,omitnil,omitempty" name:"RebuildAudioInfo"`

	// 输出目标视频控制信息。
	TargetVideoInfo *RebuildMediaTargetVideoStream `json:"TargetVideoInfo,omitnil,omitempty" name:"TargetVideoInfo"`

	// 输出目标音频控制信息。
	TargetAudioInfo *RebuildMediaTargetAudioStream `json:"TargetAudioInfo,omitnil,omitempty" name:"TargetAudioInfo"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

type CreateEnhanceMediaTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 输出文件封装格式，可选值：mp4、flv、hls。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 音画质重生视频控制控制信息。
	RebuildVideoInfo *RebuildVideoInfo `json:"RebuildVideoInfo,omitnil,omitempty" name:"RebuildVideoInfo"`

	// 音画质重生音频控制控制信息。
	RebuildAudioInfo *RebuildAudioInfo `json:"RebuildAudioInfo,omitnil,omitempty" name:"RebuildAudioInfo"`

	// 输出目标视频控制信息。
	TargetVideoInfo *RebuildMediaTargetVideoStream `json:"TargetVideoInfo,omitnil,omitempty" name:"TargetVideoInfo"`

	// 输出目标音频控制信息。
	TargetAudioInfo *RebuildMediaTargetAudioStream `json:"TargetAudioInfo,omitnil,omitempty" name:"TargetAudioInfo"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

func (r *CreateEnhanceMediaTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnhanceMediaTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Container")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RebuildVideoInfo")
	delete(f, "RebuildAudioInfo")
	delete(f, "TargetVideoInfo")
	delete(f, "TargetAudioInfo")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnhanceMediaTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnhanceMediaTemplateResponseParams struct {
	// 音画质重生模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEnhanceMediaTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnhanceMediaTemplateResponseParams `json:"Response"`
}

func (r *CreateEnhanceMediaTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnhanceMediaTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHeadTailTemplateRequestParams struct {
	// 模板名，长度限制 64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板描述信息，长度限制 256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 片头候选列表，填写视频的 FileId。转码时将自动选择与正片宽高比最接近的一个片头（相同宽高比时，靠前的候选项优先）。最多支持 5 个候选片头。
	HeadCandidateSet []*string `json:"HeadCandidateSet,omitnil,omitempty" name:"HeadCandidateSet"`

	// 片尾候选列表，填写视频的 FileId。转码时将自动选择与正片宽高比最接近的一个片尾（相同宽高比时，靠前的候选项优先）。最多支持 5 个候选片尾。
	TailCandidateSet []*string `json:"TailCandidateSet,omitnil,omitempty" name:"TailCandidateSet"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li> gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊；</li>
	// <li> white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充；</li>
	// <li> black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：stretch 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type CreateHeadTailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名，长度限制 64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板描述信息，长度限制 256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 片头候选列表，填写视频的 FileId。转码时将自动选择与正片宽高比最接近的一个片头（相同宽高比时，靠前的候选项优先）。最多支持 5 个候选片头。
	HeadCandidateSet []*string `json:"HeadCandidateSet,omitnil,omitempty" name:"HeadCandidateSet"`

	// 片尾候选列表，填写视频的 FileId。转码时将自动选择与正片宽高比最接近的一个片尾（相同宽高比时，靠前的候选项优先）。最多支持 5 个候选片尾。
	TailCandidateSet []*string `json:"TailCandidateSet,omitnil,omitempty" name:"TailCandidateSet"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li> gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊；</li>
	// <li> white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充；</li>
	// <li> black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：stretch 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

func (r *CreateHeadTailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHeadTailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	delete(f, "Comment")
	delete(f, "HeadCandidateSet")
	delete(f, "TailCandidateSet")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHeadTailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHeadTailTemplateResponseParams struct {
	// 片头片尾模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateHeadTailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateHeadTailTemplateResponseParams `json:"Response"`
}

func (r *CreateHeadTailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHeadTailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageProcessingTemplateRequestParams struct {
	// 图片处理操作数组，操作将以其在数组中的顺序执行。
	// <li>长度限制：10。</li>
	Operations []*ImageOperation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 图片处理模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateImageProcessingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 图片处理操作数组，操作将以其在数组中的顺序执行。
	// <li>长度限制：10。</li>
	Operations []*ImageOperation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 图片处理模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateImageProcessingTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageProcessingTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operations")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageProcessingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageProcessingTemplateResponseParams struct {
	// 图片处理模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateImageProcessingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageProcessingTemplateResponseParams `json:"Response"`
}

func (r *CreateImageProcessingTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageProcessingTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageSpriteTask2017 struct {
	// 截图雪碧图任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 截取雪碧图文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 雪碧图规格，参见[雪碧图截图模板](https://cloud.tencent.com/document/product/266/33480#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 雪碧图小图总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 截取雪碧图输出的地址。
	ImageSpriteUrlSet []*string `json:"ImageSpriteUrlSet,omitnil,omitempty" name:"ImageSpriteUrlSet"`

	// 雪碧图子图位置与时间关系 WebVtt 文件地址。
	WebVttUrl *string `json:"WebVttUrl,omitnil,omitempty" name:"WebVttUrl"`
}

// Predefined struct for user
type CreateImageSpriteTemplateRequestParams struct {
	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	// 注意：小图的行数会影响最终大图的高度，大图的高度最大为15000像素，其中大图的高度为小图行数与小图高度的乘积。
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// 雪碧图中小图的列数。
	// 注意：小图的列数会影响最终大图的宽度，大图的宽度最大为15000像素，其中大图的宽度为小图列数与小图宽度的乘积。
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	// 注意：小图的宽度会影响最终大图的宽度，大图的宽度最大为15000像素，其中大图的宽度为小图列数与小图宽度的乘积。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	// 注意：小图的高度会影响最终大图的高度，大图的高度最大为15000像素，其中大图的高度为小图行数与小图高度的乘积。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式，取值：
	// <li> jpg：jpg 格式；</li>
	// <li> png：png 格式；</li>
	// <li> webp：webp 格式。</li>
	// 默认值：jpg。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type CreateImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	// 注意：小图的行数会影响最终大图的高度，大图的高度最大为15000像素，其中大图的高度为小图行数与小图高度的乘积。
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// 雪碧图中小图的列数。
	// 注意：小图的列数会影响最终大图的宽度，大图的宽度最大为15000像素，其中大图的宽度为小图列数与小图宽度的乘积。
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	// 注意：小图的宽度会影响最终大图的宽度，大图的宽度最大为15000像素，其中大图的宽度为小图列数与小图宽度的乘积。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	// 注意：小图的高度会影响最终大图的高度，大图的高度最大为15000像素，其中大图的高度为小图行数与小图高度的乘积。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式，取值：
	// <li> jpg：jpg 格式；</li>
	// <li> png：png 格式；</li>
	// <li> webp：webp 格式。</li>
	// 默认值：jpg。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "FillType")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageSpriteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageSpriteTemplateResponseParams struct {
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateJustInTimeTranscodeTemplateRequestParams struct {
	// 模板名字，长度限制64个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 视频参数配置。
	VideoConfigure *VideoConfigureInfo `json:"VideoConfigure,omitnil,omitempty" name:"VideoConfigure"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 水印参数配置。
	WatermarkConfigure *WatermarkConfigureInfo `json:"WatermarkConfigure,omitnil,omitempty" name:"WatermarkConfigure"`

	// 模板描述，长度限制256个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateJustInTimeTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名字，长度限制64个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 视频参数配置。
	VideoConfigure *VideoConfigureInfo `json:"VideoConfigure,omitnil,omitempty" name:"VideoConfigure"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 水印参数配置。
	WatermarkConfigure *WatermarkConfigureInfo `json:"WatermarkConfigure,omitnil,omitempty" name:"WatermarkConfigure"`

	// 模板描述，长度限制256个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateJustInTimeTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJustInTimeTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "VideoConfigure")
	delete(f, "SubAppId")
	delete(f, "WatermarkConfigure")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJustInTimeTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJustInTimeTranscodeTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateJustInTimeTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateJustInTimeTranscodeTemplateResponseParams `json:"Response"`
}

func (r *CreateJustInTimeTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJustInTimeTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonSampleRequestParams struct {
	// 素材名称，长度限制：20 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 素材应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于内容不适宜，等价于 Review.Face。
	// 3. All：包含以上全部，等价于 1+2。
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 素材描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 素材图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 jpeg、png 图片格式。数组长度限制：5 张图片。
	// 注意：图片必须是单人像五官较清晰的照片，像素不低于 200*200。
	FaceContents []*string `json:"FaceContents,omitnil,omitempty" name:"FaceContents"`

	// 素材标签
	// <li>数组长度限制：20 个标签；</li>
	// <li>单个标签长度限制：128 个字符。</li>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreatePersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// 素材名称，长度限制：20 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 素材应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于内容不适宜，等价于 Review.Face。
	// 3. All：包含以上全部，等价于 1+2。
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 素材描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 素材图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 jpeg、png 图片格式。数组长度限制：5 张图片。
	// 注意：图片必须是单人像五官较清晰的照片，像素不低于 200*200。
	FaceContents []*string `json:"FaceContents,omitnil,omitempty" name:"FaceContents"`

	// 素材标签
	// <li>数组长度限制：20 个标签；</li>
	// <li>单个标签长度限制：128 个字符。</li>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "SubAppId")
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
	Person *AiSamplePerson `json:"Person,omitnil,omitempty" name:"Person"`

	// 处理失败的五官定位信息。
	FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitnil,omitempty" name:"FailFaceInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateProcedureTemplateRequestParams struct {
	// 任务流名字（支持中文，不超过20个字）。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// AI 内容审核类型任务参数 \*。
	// <font color=red>\*：该参数用于发起旧版审核，不建议使用。推荐使用 ReviewAudioVideoTask 参数发起审核。</font> 
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// AI 内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// AI 内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// 音视频审核类型任务参数。
	ReviewAudioVideoTask *ProcedureReviewAudioVideoTaskInput `json:"ReviewAudioVideoTask,omitnil,omitempty" name:"ReviewAudioVideoTask"`
}

type CreateProcedureTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 任务流名字（支持中文，不超过20个字）。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// AI 内容审核类型任务参数 \*。
	// <font color=red>\*：该参数用于发起旧版审核，不建议使用。推荐使用 ReviewAudioVideoTask 参数发起审核。</font> 
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// AI 内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// AI 内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// 音视频审核类型任务参数。
	ReviewAudioVideoTask *ProcedureReviewAudioVideoTaskInput `json:"ReviewAudioVideoTask,omitnil,omitempty" name:"ReviewAudioVideoTask"`
}

func (r *CreateProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcedureTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	delete(f, "Comment")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "ReviewAudioVideoTask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProcedureTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProcedureTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateProcedureTemplateResponseParams `json:"Response"`
}

func (r *CreateProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcedureTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQualityInspectTemplateRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质检测模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音画质检测模板描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 截帧间隔，单位为秒，最小值为 1。当不填时，默认截帧间隔为 1 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`

	// 视频画面抖动重影检测的控制参数。
	JitterConfigure *JitterConfigureInfo `json:"JitterConfigure,omitnil,omitempty" name:"JitterConfigure"`

	// 视频画面模糊检测的控制参数。
	BlurConfigure *BlurConfigureInfo `json:"BlurConfigure,omitnil,omitempty" name:"BlurConfigure"`

	// 视频画面低光、过曝检测的控制参数。
	AbnormalLightingConfigure *AbnormalLightingConfigureInfo `json:"AbnormalLightingConfigure,omitnil,omitempty" name:"AbnormalLightingConfigure"`

	// 视频画面花屏检测的控制参数。
	CrashScreenConfigure *CrashScreenConfigureInfo `json:"CrashScreenConfigure,omitnil,omitempty" name:"CrashScreenConfigure"`

	// 视频画面黑边、白边、黑屏、白屏检测的控制参数。
	BlackWhiteEdgeConfigure *BlackWhiteEdgeConfigureInfo `json:"BlackWhiteEdgeConfigure,omitnil,omitempty" name:"BlackWhiteEdgeConfigure"`

	// 视频画面噪点检测的控制参数。
	NoiseConfigure *NoiseConfigureInfo `json:"NoiseConfigure,omitnil,omitempty" name:"NoiseConfigure"`

	// 视频画面马赛克检测的控制参数。
	MosaicConfigure *MosaicConfigureInfo `json:"MosaicConfigure,omitnil,omitempty" name:"MosaicConfigure"`

	// 视频画面二维码检测的控制参数。
	QRCodeConfigure *QRCodeConfigureInfo `json:"QRCodeConfigure,omitnil,omitempty" name:"QRCodeConfigure"`

	// 音频（静音、低音、爆音）检测的控制参数。
	VoiceConfigure *VoiceConfigureInfo `json:"VoiceConfigure,omitnil,omitempty" name:"VoiceConfigure"`

	// 视频画面质量评价的控制参数。
	QualityEvaluationConfigure *QualityEvaluationConfigureInfo `json:"QualityEvaluationConfigure,omitnil,omitempty" name:"QualityEvaluationConfigure"`
}

type CreateQualityInspectTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质检测模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音画质检测模板描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 截帧间隔，单位为秒，最小值为 1。当不填时，默认截帧间隔为 1 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`

	// 视频画面抖动重影检测的控制参数。
	JitterConfigure *JitterConfigureInfo `json:"JitterConfigure,omitnil,omitempty" name:"JitterConfigure"`

	// 视频画面模糊检测的控制参数。
	BlurConfigure *BlurConfigureInfo `json:"BlurConfigure,omitnil,omitempty" name:"BlurConfigure"`

	// 视频画面低光、过曝检测的控制参数。
	AbnormalLightingConfigure *AbnormalLightingConfigureInfo `json:"AbnormalLightingConfigure,omitnil,omitempty" name:"AbnormalLightingConfigure"`

	// 视频画面花屏检测的控制参数。
	CrashScreenConfigure *CrashScreenConfigureInfo `json:"CrashScreenConfigure,omitnil,omitempty" name:"CrashScreenConfigure"`

	// 视频画面黑边、白边、黑屏、白屏检测的控制参数。
	BlackWhiteEdgeConfigure *BlackWhiteEdgeConfigureInfo `json:"BlackWhiteEdgeConfigure,omitnil,omitempty" name:"BlackWhiteEdgeConfigure"`

	// 视频画面噪点检测的控制参数。
	NoiseConfigure *NoiseConfigureInfo `json:"NoiseConfigure,omitnil,omitempty" name:"NoiseConfigure"`

	// 视频画面马赛克检测的控制参数。
	MosaicConfigure *MosaicConfigureInfo `json:"MosaicConfigure,omitnil,omitempty" name:"MosaicConfigure"`

	// 视频画面二维码检测的控制参数。
	QRCodeConfigure *QRCodeConfigureInfo `json:"QRCodeConfigure,omitnil,omitempty" name:"QRCodeConfigure"`

	// 音频（静音、低音、爆音）检测的控制参数。
	VoiceConfigure *VoiceConfigureInfo `json:"VoiceConfigure,omitnil,omitempty" name:"VoiceConfigure"`

	// 视频画面质量评价的控制参数。
	QualityEvaluationConfigure *QualityEvaluationConfigureInfo `json:"QualityEvaluationConfigure,omitnil,omitempty" name:"QualityEvaluationConfigure"`
}

func (r *CreateQualityInspectTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQualityInspectTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ScreenshotInterval")
	delete(f, "JitterConfigure")
	delete(f, "BlurConfigure")
	delete(f, "AbnormalLightingConfigure")
	delete(f, "CrashScreenConfigure")
	delete(f, "BlackWhiteEdgeConfigure")
	delete(f, "NoiseConfigure")
	delete(f, "MosaicConfigure")
	delete(f, "QRCodeConfigure")
	delete(f, "VoiceConfigure")
	delete(f, "QualityEvaluationConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQualityInspectTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQualityInspectTemplateResponseParams struct {
	// 音画质检测模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQualityInspectTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateQualityInspectTemplateResponseParams `json:"Response"`
}

func (r *CreateQualityInspectTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQualityInspectTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRebuildMediaTemplateRequestParams struct {
	// 输出文件封装格式，可选值：mp4、flv、hls。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 音画质重生视频控制控制信息。
	RebuildVideoInfo *RebuildVideoInfo `json:"RebuildVideoInfo,omitnil,omitempty" name:"RebuildVideoInfo"`

	// 音画质重生音频控制控制信息。
	RebuildAudioInfo *RebuildAudioInfo `json:"RebuildAudioInfo,omitnil,omitempty" name:"RebuildAudioInfo"`

	// 输出目标视频控制信息。
	TargetVideoInfo *RebuildMediaTargetVideoStream `json:"TargetVideoInfo,omitnil,omitempty" name:"TargetVideoInfo"`

	// 输出目标音频控制信息。
	TargetAudioInfo *RebuildMediaTargetAudioStream `json:"TargetAudioInfo,omitnil,omitempty" name:"TargetAudioInfo"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveAudio *string `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

type CreateRebuildMediaTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 输出文件封装格式，可选值：mp4、flv、hls。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 音画质重生视频控制控制信息。
	RebuildVideoInfo *RebuildVideoInfo `json:"RebuildVideoInfo,omitnil,omitempty" name:"RebuildVideoInfo"`

	// 音画质重生音频控制控制信息。
	RebuildAudioInfo *RebuildAudioInfo `json:"RebuildAudioInfo,omitnil,omitempty" name:"RebuildAudioInfo"`

	// 输出目标视频控制信息。
	TargetVideoInfo *RebuildMediaTargetVideoStream `json:"TargetVideoInfo,omitnil,omitempty" name:"TargetVideoInfo"`

	// 输出目标音频控制信息。
	TargetAudioInfo *RebuildMediaTargetAudioStream `json:"TargetAudioInfo,omitnil,omitempty" name:"TargetAudioInfo"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveAudio *string `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

func (r *CreateRebuildMediaTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRebuildMediaTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Container")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RebuildVideoInfo")
	delete(f, "RebuildAudioInfo")
	delete(f, "TargetVideoInfo")
	delete(f, "TargetAudioInfo")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRebuildMediaTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRebuildMediaTemplateResponseParams struct {
	// 视频重生模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRebuildMediaTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateRebuildMediaTemplateResponseParams `json:"Response"`
}

func (r *CreateRebuildMediaTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRebuildMediaTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReviewTemplateRequestParams struct {
	// 需要返回的违规标签列表，可选值为：
	// <li>Porn：色情；</li>
	// <li>Terror：暴力；</li>
	// <li>Polity：不适宜的信息；</li>
	// <li>Illegal：违法；</li>
	// <li>Abuse：谩骂；</li>
	// <li>Ad：广告；</li>
	// <li>Moan：娇喘。</li>
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *string `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的违规标签列表，可选值为：
	// <li>Porn：色情；</li>
	// <li>Terror：暴力；</li>
	// <li>Polity：不适宜的信息；</li>
	// <li>Illegal：违法；</li>
	// <li>Abuse：谩骂；</li>
	// <li>Ad：广告；</li>
	// <li>Moan：娇喘。</li>
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *string `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReviewTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Labels")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReviewTemplateResponseParams struct {
	// 审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateReviewTemplateResponseParams `json:"Response"`
}

func (r *CreateReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReviewTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoundPlayRequestParams struct {
	// 启播时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 轮播列表。
	// <li>数组长度限制：100。</li>
	RoundPlaylist []*RoundPlayListItemInfo `json:"RoundPlaylist,omitnil,omitempty" name:"RoundPlaylist"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 轮播播单名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 轮播播单描述信息，长度限制：256 个字符。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 播放模式，可选值：
	// <li>Loop：循环播放播单；</li>
	// <li>Linear：单次播放，播单播放完停止播放。</li>
	// 默认值：Loop。
	PlayBackMode *string `json:"PlayBackMode,omitnil,omitempty" name:"PlayBackMode"`
}

type CreateRoundPlayRequest struct {
	*tchttp.BaseRequest
	
	// 启播时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 轮播列表。
	// <li>数组长度限制：100。</li>
	RoundPlaylist []*RoundPlayListItemInfo `json:"RoundPlaylist,omitnil,omitempty" name:"RoundPlaylist"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 轮播播单名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 轮播播单描述信息，长度限制：256 个字符。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 播放模式，可选值：
	// <li>Loop：循环播放播单；</li>
	// <li>Linear：单次播放，播单播放完停止播放。</li>
	// 默认值：Loop。
	PlayBackMode *string `json:"PlayBackMode,omitnil,omitempty" name:"PlayBackMode"`
}

func (r *CreateRoundPlayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoundPlayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "RoundPlaylist")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Desc")
	delete(f, "PlayBackMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoundPlayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoundPlayResponseParams struct {
	// 轮播播单唯一标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// 轮播播放地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRoundPlayResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoundPlayResponseParams `json:"Response"`
}

func (r *CreateRoundPlayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoundPlayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSampleSnapshotTemplateRequestParams struct {
	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式，取值为 jpg 和 png。默认为 jpg。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type CreateSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式，取值为 jpg 和 png。默认为 jpg。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
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
	delete(f, "SubAppId")
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
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateSnapshotByTimeOffsetTemplateRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式，取值可以为 jpg 和 png。默认为 jpg。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type CreateSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式，取值可以为 jpg 和 png。默认为 jpg。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
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
	delete(f, "SubAppId")
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
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateStorageRegionRequestParams struct {
	// 待开通的存储地域，必须是系统支持的地域。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type CreateStorageRegionRequest struct {
	*tchttp.BaseRequest
	
	// 待开通的存储地域，必须是系统支持的地域。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *CreateStorageRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StorageRegion")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStorageRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageRegionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStorageRegionResponse struct {
	*tchttp.BaseResponse
	Response *CreateStorageRegionResponseParams `json:"Response"`
}

func (r *CreateStorageRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubAppIdRequestParams struct {
	// 应用名称，长度限制：40个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用简介，长度限制： 300个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateSubAppIdRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称，长度限制：40个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用简介，长度限制： 300个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateSubAppIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubAppIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubAppIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubAppIdResponseParams struct {
	// 新创建的应用 ID。
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSubAppIdResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubAppIdResponseParams `json:"Response"`
}

func (r *CreateSubAppIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSuperPlayerConfigRequestParams struct {
	// 播放器配置名称，长度限制：64 个字符。只允许出现 [0-9a-zA-Z] 及 _- 字符（如 test_ABC-123），同一个用户该名称唯一。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 播放的音视频类型，可选值：
	// <li>AdaptiveDynamicStream：自适应码流输出；</li>
	// <li>Transcode：转码输出；</li>
	// <li>Original：原始音视频。</li>
	// 默认为 AdaptiveDynamicStream。
	AudioVideoType *string `json:"AudioVideoType,omitnil,omitempty" name:"AudioVideoType"`

	// 播放 DRM 保护的自适应码流开关：
	// <li>ON：开启，表示仅播放 DRM  保护的自适应码流输出；</li>
	// <li>OFF：关闭，表示播放未加密的自适应码流输出。</li>
	// 默认为 OFF。
	// 当 AudioVideoType 为 AdaptiveDynamicStream 时，此参数有效。
	DrmSwitch *string `json:"DrmSwitch,omitnil,omitempty" name:"DrmSwitch"`

	// 允许输出的未加密的自适应码流模板 ID。
	// 
	// 当 AudioVideoType 为 AdaptiveDynamicStream 并且 DrmSwitch 为 OFF 时，此参数为必填。
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitnil,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// 允许输出的 DRM 自适应码流模板内容。
	// 
	// 当 AudioVideoType 为 AdaptiveDynamicStream 并且 DrmSwitch 为 ON 时，此参数为必填。
	DrmStreamingsInfo *DrmStreamingsInfo `json:"DrmStreamingsInfo,omitnil,omitempty" name:"DrmStreamingsInfo"`

	// 允许输出的转码模板 ID。
	// 
	// 当 AudioVideoType 为 Transcode 时必填。
	TranscodeDefinition *uint64 `json:"TranscodeDefinition,omitnil,omitempty" name:"TranscodeDefinition"`

	// 允许输出的雪碧图模板 ID。
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitnil,omitempty" name:"ImageSpriteDefinition"`

	// 播放器对于不同分辨率的子流展示名字，不填或者填空数组则使用默认配置：
	// <li>MinEdgeLength：240，Name：流畅；</li>
	// <li>MinEdgeLength：480，Name：标清；</li>
	// <li>MinEdgeLength：720，Name：高清；</li>
	// <li>MinEdgeLength：1080，Name：全高清；</li>
	// <li>MinEdgeLength：1440，Name：2K；</li>
	// <li>MinEdgeLength：2160，Name：4K；</li>
	// <li>MinEdgeLength：4320，Name：8K。</li>
	ResolutionNames []*ResolutionNameInfo `json:"ResolutionNames,omitnil,omitempty" name:"ResolutionNames"`

	// 播放时使用的域名。不填或者填 Default，表示使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 播放时使用的 Scheme。不填或者填 Default，表示使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的 Scheme。其他可选值：
	// <li>HTTP；</li>
	// <li>HTTPS。</li>
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateSuperPlayerConfigRequest struct {
	*tchttp.BaseRequest
	
	// 播放器配置名称，长度限制：64 个字符。只允许出现 [0-9a-zA-Z] 及 _- 字符（如 test_ABC-123），同一个用户该名称唯一。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 播放的音视频类型，可选值：
	// <li>AdaptiveDynamicStream：自适应码流输出；</li>
	// <li>Transcode：转码输出；</li>
	// <li>Original：原始音视频。</li>
	// 默认为 AdaptiveDynamicStream。
	AudioVideoType *string `json:"AudioVideoType,omitnil,omitempty" name:"AudioVideoType"`

	// 播放 DRM 保护的自适应码流开关：
	// <li>ON：开启，表示仅播放 DRM  保护的自适应码流输出；</li>
	// <li>OFF：关闭，表示播放未加密的自适应码流输出。</li>
	// 默认为 OFF。
	// 当 AudioVideoType 为 AdaptiveDynamicStream 时，此参数有效。
	DrmSwitch *string `json:"DrmSwitch,omitnil,omitempty" name:"DrmSwitch"`

	// 允许输出的未加密的自适应码流模板 ID。
	// 
	// 当 AudioVideoType 为 AdaptiveDynamicStream 并且 DrmSwitch 为 OFF 时，此参数为必填。
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitnil,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// 允许输出的 DRM 自适应码流模板内容。
	// 
	// 当 AudioVideoType 为 AdaptiveDynamicStream 并且 DrmSwitch 为 ON 时，此参数为必填。
	DrmStreamingsInfo *DrmStreamingsInfo `json:"DrmStreamingsInfo,omitnil,omitempty" name:"DrmStreamingsInfo"`

	// 允许输出的转码模板 ID。
	// 
	// 当 AudioVideoType 为 Transcode 时必填。
	TranscodeDefinition *uint64 `json:"TranscodeDefinition,omitnil,omitempty" name:"TranscodeDefinition"`

	// 允许输出的雪碧图模板 ID。
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitnil,omitempty" name:"ImageSpriteDefinition"`

	// 播放器对于不同分辨率的子流展示名字，不填或者填空数组则使用默认配置：
	// <li>MinEdgeLength：240，Name：流畅；</li>
	// <li>MinEdgeLength：480，Name：标清；</li>
	// <li>MinEdgeLength：720，Name：高清；</li>
	// <li>MinEdgeLength：1080，Name：全高清；</li>
	// <li>MinEdgeLength：1440，Name：2K；</li>
	// <li>MinEdgeLength：2160，Name：4K；</li>
	// <li>MinEdgeLength：4320，Name：8K。</li>
	ResolutionNames []*ResolutionNameInfo `json:"ResolutionNames,omitnil,omitempty" name:"ResolutionNames"`

	// 播放时使用的域名。不填或者填 Default，表示使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 播放时使用的 Scheme。不填或者填 Default，表示使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的 Scheme。其他可选值：
	// <li>HTTP；</li>
	// <li>HTTPS。</li>
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateSuperPlayerConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSuperPlayerConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	delete(f, "AudioVideoType")
	delete(f, "DrmSwitch")
	delete(f, "AdaptiveDynamicStreamingDefinition")
	delete(f, "DrmStreamingsInfo")
	delete(f, "TranscodeDefinition")
	delete(f, "ImageSpriteDefinition")
	delete(f, "ResolutionNames")
	delete(f, "Domain")
	delete(f, "Scheme")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSuperPlayerConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSuperPlayerConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSuperPlayerConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateSuperPlayerConfigResponseParams `json:"Response"`
}

func (r *CreateSuperPlayerConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSuperPlayerConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeTemplateRequestParams struct {
	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a、wav。其中，mp3、flac、ogg、m4a、wav 为纯音频文件。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// 视频流配置参数，当 RemoveVideo 为 0，该字段必填。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，当 RemoveAudio 为 0，该字段必填。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// 极速高清转码参数。
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// 切片类型，当 Container 为 hls 时有效，可选值：
	// <li>ts：ts 切片；</li>
	// <li>fmp4：fmp4 切片。</li>
	// 默认值：ts。
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
}

type CreateTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a、wav。其中，mp3、flac、ogg、m4a、wav 为纯音频文件。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// 视频流配置参数，当 RemoveVideo 为 0，该字段必填。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，当 RemoveAudio 为 0，该字段必填。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// 极速高清转码参数。
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// 切片类型，当 Container 为 hls 时有效，可选值：
	// <li>ts：ts 切片；</li>
	// <li>fmp4：fmp4 切片。</li>
	// 默认值：ts。
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	delete(f, "VideoTemplate")
	delete(f, "AudioTemplate")
	delete(f, "TEHDConfig")
	delete(f, "SegmentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeTemplateResponseParams struct {
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateVodDomainRequestParams struct {
	// 需要接入点播的加速域名。注意：不支持填写泛域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 需要开启 CDN 加速的区域：
	// <li>Chinese Mainland：中国境内（不包含港澳台）。</li>
	// <li>Outside Chinese Mainland: 中国境外。</li>
	// <li>Global: 全球范围。</li>
	// 如果没有设置 AccelerateArea， 点播会根据用户在腾讯云设置的地域信息自动开通中国境内或者中国境外的 CDN 加速。开启中国境内加速的域名，需要先[备案域名](/document/product/243/18905)。
	AccelerateArea *string `json:"AccelerateArea,omitnil,omitempty" name:"AccelerateArea"`
}

type CreateVodDomainRequest struct {
	*tchttp.BaseRequest
	
	// 需要接入点播的加速域名。注意：不支持填写泛域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 需要开启 CDN 加速的区域：
	// <li>Chinese Mainland：中国境内（不包含港澳台）。</li>
	// <li>Outside Chinese Mainland: 中国境外。</li>
	// <li>Global: 全球范围。</li>
	// 如果没有设置 AccelerateArea， 点播会根据用户在腾讯云设置的地域信息自动开通中国境内或者中国境外的 CDN 加速。开启中国境内加速的域名，需要先[备案域名](/document/product/243/18905)。
	AccelerateArea *string `json:"AccelerateArea,omitnil,omitempty" name:"AccelerateArea"`
}

func (r *CreateVodDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVodDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SubAppId")
	delete(f, "AccelerateArea")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVodDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVodDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVodDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateVodDomainResponseParams `json:"Response"`
}

func (r *CreateVodDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVodDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWatermarkTemplateRequestParams struct {
	// 水印类型，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印；</li>
	// <li>svg：SVG 水印。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	// 默认值：TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// 图片水印模板，当 Type 为 image，该字段必填。当 Type 为 text，该字段无效。
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// 文字水印模板，当 Type 为 text，该字段必填。当 Type 为 image，该字段无效。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG水印模板，当 Type 为 svg，该字段必填。当 Type 为 image 或 text，该字段无效。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`
}

type CreateWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 水印类型，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印；</li>
	// <li>svg：SVG 水印。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	// 默认值：TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// 图片水印模板，当 Type 为 image，该字段必填。当 Type 为 text，该字段无效。
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// 文字水印模板，当 Type 为 text，该字段必填。当 Type 为 image，该字段无效。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG水印模板，当 Type 为 svg，该字段必填。当 Type 为 image 或 text，该字段无效。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`
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
	delete(f, "SubAppId")
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
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 水印图片地址，仅当 Type 为 image，该字段有效。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 关键词，数组长度限制：100。
	Words []*AiSampleWordInfo `json:"Words,omitnil,omitempty" name:"Words"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 关键词，数组长度限制：100。
	Words []*AiSampleWordInfo `json:"Words,omitnil,omitempty" name:"Words"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWordSamplesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DNSVerifyInfo struct {
	// 子解析。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 解析值。
	Record *string `json:"Record,omitnil,omitempty" name:"Record"`

	// 解析类型。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`
}

type DailyPlayStatInfo struct {
	// 播放媒体文件的日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 媒体文件ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 播放次数。
	PlayTimes *uint64 `json:"PlayTimes,omitnil,omitempty" name:"PlayTimes"`

	// 播放流量，单位：字节。
	Traffic *uint64 `json:"Traffic,omitnil,omitempty" name:"Traffic"`
}

// Predefined struct for user
type DeleteAIAnalysisTemplateRequestParams struct {
	// 音视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 音视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIAnalysisTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIAnalysisTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 音视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 音视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIRecognitionTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 自适应转码模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAdaptiveDynamicStreamingTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAnimatedGraphicsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAnimatedGraphicsTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteCLSTopicRequestParams struct {
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`

	// 日志主题 ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteCLSTopicRequest struct {
	*tchttp.BaseRequest
	
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`

	// 日志主题 ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteCLSTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCLSTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CLSRegion")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCLSTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCLSTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCLSTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCLSTopicResponseParams `json:"Response"`
}

func (r *DeleteCLSTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCLSTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClassRequestParams struct {
	// 分类 ID
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteClassRequest struct {
	*tchttp.BaseRequest
	
	// 分类 ID
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClassId")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClassResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClassResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClassResponseParams `json:"Response"`
}

func (r *DeleteClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContentReviewTemplateRequestParams struct {
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContentReviewTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteEnhanceMediaTemplateRequestParams struct {
	// 音画质重生模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteEnhanceMediaTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 音画质重生模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteEnhanceMediaTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnhanceMediaTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEnhanceMediaTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEnhanceMediaTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEnhanceMediaTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEnhanceMediaTemplateResponseParams `json:"Response"`
}

func (r *DeleteEnhanceMediaTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnhanceMediaTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHeadTailTemplateRequestParams struct {
	// 片头片尾模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteHeadTailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 片头片尾模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteHeadTailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHeadTailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHeadTailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHeadTailTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteHeadTailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHeadTailTemplateResponseParams `json:"Response"`
}

func (r *DeleteHeadTailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHeadTailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageProcessingTemplateRequestParams struct {
	// 图片处理模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteImageProcessingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 图片处理模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteImageProcessingTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageProcessingTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageProcessingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageProcessingTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImageProcessingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageProcessingTemplateResponseParams `json:"Response"`
}

func (r *DeleteImageProcessingTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageProcessingTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageSpriteTemplateRequestParams struct {
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageSpriteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageSpriteTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteJustInTimeTranscodeTemplateRequestParams struct {
	// 模板名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteJustInTimeTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteJustInTimeTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJustInTimeTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteJustInTimeTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJustInTimeTranscodeTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteJustInTimeTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteJustInTimeTranscodeTemplateResponseParams `json:"Response"`
}

func (r *DeleteJustInTimeTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJustInTimeTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMediaRequestParams struct {
	// 媒体文件的唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定本次需要删除的部分。默认值为 "[]", 表示删除媒体及其对应的全部视频处理文件。
	DeleteParts []*MediaDeleteItem `json:"DeleteParts,omitnil,omitempty" name:"DeleteParts"`
}

type DeleteMediaRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件的唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定本次需要删除的部分。默认值为 "[]", 表示删除媒体及其对应的全部视频处理文件。
	DeleteParts []*MediaDeleteItem `json:"DeleteParts,omitnil,omitempty" name:"DeleteParts"`
}

func (r *DeleteMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "DeleteParts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMediaResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMediaResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMediaResponseParams `json:"Response"`
}

func (r *DeleteMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonSampleRequestParams struct {
	// 素材 ID。
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeletePersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// 素材 ID。
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonSampleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteProcedureTemplateRequestParams struct {
	// 任务流名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteProcedureTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 任务流名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProcedureTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProcedureTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProcedureTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProcedureTemplateResponseParams `json:"Response"`
}

func (r *DeleteProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProcedureTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQualityInspectTemplateRequestParams struct {
	// 音画质检测模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteQualityInspectTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 音画质检测模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteQualityInspectTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQualityInspectTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQualityInspectTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQualityInspectTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQualityInspectTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQualityInspectTemplateResponseParams `json:"Response"`
}

func (r *DeleteQualityInspectTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQualityInspectTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRebuildMediaTemplateRequestParams struct {
	// 视频重生模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteRebuildMediaTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 视频重生模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteRebuildMediaTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRebuildMediaTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRebuildMediaTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRebuildMediaTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRebuildMediaTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRebuildMediaTemplateResponseParams `json:"Response"`
}

func (r *DeleteRebuildMediaTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRebuildMediaTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReviewTemplateRequestParams struct {
	// 审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReviewTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReviewTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReviewTemplateResponseParams `json:"Response"`
}

func (r *DeleteReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReviewTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoundPlayRequestParams struct {
	// 轮播播单唯一标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteRoundPlayRequest struct {
	*tchttp.BaseRequest
	
	// 轮播播单唯一标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteRoundPlayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoundPlayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoundPlayId")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoundPlayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoundPlayResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRoundPlayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoundPlayResponseParams `json:"Response"`
}

func (r *DeleteRoundPlayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoundPlayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSampleSnapshotTemplateRequestParams struct {
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSampleSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSampleSnapshotTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteSnapshotByTimeOffsetTemplateRequestParams struct {
	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotByTimeOffsetTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotByTimeOffsetTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteSuperPlayerConfigRequestParams struct {
	// 播放器配置名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteSuperPlayerConfigRequest struct {
	*tchttp.BaseRequest
	
	// 播放器配置名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteSuperPlayerConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSuperPlayerConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSuperPlayerConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSuperPlayerConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSuperPlayerConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSuperPlayerConfigResponseParams `json:"Response"`
}

func (r *DeleteSuperPlayerConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSuperPlayerConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTranscodeTemplateRequestParams struct {
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTranscodeTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteVodDomainRequestParams struct {
	// 要删除的点播加速域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteVodDomainRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的点播加速域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DeleteVodDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVodDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVodDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVodDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVodDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVodDomainResponseParams `json:"Response"`
}

func (r *DeleteVodDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVodDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWatermarkTemplateRequestParams struct {
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWatermarkTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWatermarkTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DeleteWordSamplesRequest struct {
	*tchttp.BaseRequest
	
	// 关键词，数组长度限制：100 个词。
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWordSamplesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAIAnalysisTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容分析模板唯一标识过滤条件，数组长度最大值：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAIAnalysisTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容分析模板唯一标识过滤条件，数组长度最大值：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIAnalysisTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAnalysisTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 音视频内容分析模板详情列表。
	AIAnalysisTemplateSet []*AIAnalysisTemplateItem `json:"AIAnalysisTemplateSet,omitnil,omitempty" name:"AIAnalysisTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容识别模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：<li>Preset：系统预置模板；</li><li>Custom：用户自定义模板。</li>不填默认为空，即不对模板类型过滤。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAIRecognitionTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容识别模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：<li>Preset：系统预置模板；</li><li>Custom：用户自定义模板。</li>不填默认为空，即不对模板类型过滤。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIRecognitionTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIRecognitionTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 音视频内容识别模板详情列表。
	AIRecognitionTemplateSet []*AIRecognitionTemplateItem `json:"AIRecognitionTemplateSet,omitnil,omitempty" name:"AIRecognitionTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转自适应码流模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeAdaptiveDynamicStreamingTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转自适应码流模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	delete(f, "SubAppId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 转自适应码流模板详情列表。
	AdaptiveDynamicStreamingTemplateSet []*AdaptiveDynamicStreamingTemplate `json:"AdaptiveDynamicStreamingTemplateSet,omitnil,omitempty" name:"AdaptiveDynamicStreamingTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAllClassRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeAllClassRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeAllClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllClassResponseParams struct {
	// 分类信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassInfoSet []*MediaClassInfo `json:"ClassInfoSet,omitnil,omitempty" name:"ClassInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllClassResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllClassResponseParams `json:"Response"`
}

func (r *DescribeAllClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnimatedGraphicsTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转动图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeAnimatedGraphicsTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转动图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	delete(f, "SubAppId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 转动图模板详情列表。
	AnimatedGraphicsTemplateSet []*AnimatedGraphicsTemplate `json:"AnimatedGraphicsTemplateSet,omitnil,omitempty" name:"AnimatedGraphicsTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCDNStatDetailsRequestParams struct {
	// 查询指标，取值有：
	// <li>Traffic：流量，单位为 Byte。</li>
	// <li>Bandwidth：带宽，单位为 Bps。</li>
	// <li>Requests：请求数。</li>
	// <li>QUICRequests：QUIC 请求数。</li>
	// <li>IpVisit：IP 访问数。</li>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 起始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 域名列表。一次最多查询20个域名的数据。默认返回所有域名叠加的用量数据。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 服务区域，取值有：
	// <li>Chinese Mainland：中国大陆。 </li>
	// <li>Asia Pacific Region 1：亚太一区，包括中国香港、中国澳门、新加坡、越南、泰国。 </li>
	// <li>Asia Pacific Region 2：亚太二区，包括中国台湾、日本、马来西亚、印度尼西亚、韩国。 </li>
	// <li>Asia Pacific Region 3：亚太三区，包括菲律宾、印度、澳大利亚和亚太其它国家和地区。 </li>
	// <li>Middle East：中东。 </li>
	// <li>Europe：欧洲。</li>
	// <li>North America：北美。</li>
	// <li>South America：南美。</li>
	// <li>Africa：非洲。</li>
	// 默认为中国大陆。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 用户所在地区，Area 为 Chinese Mainland 时，取值为以下地区信息，当 Area 为其它值时， 忽略 Districts 参数。
	// <li>Beijing：北京。</li>
	// <li>Inner Mongolia：内蒙古。</li>
	// <li>Shanxi：山西。</li>
	// <li>Hebei：河北。</li>
	// <li>Tianjin：天津。</li>
	// <li>Ningxia：宁夏。</li>
	// <li>Shaanxi：陕西。</li>
	// <li>Gansu：甘肃。</li>
	// <li>Qinghai：青海。</li>
	// <li>Xinjiang：新疆。</li>
	// <li>Heilongjiang：黑龙江。</li>
	// <li>Jilin：吉林。</li>
	// <li>Liaoning：辽宁。</li>
	// <li>Fujian：福建。</li>
	// <li>Jiangsu：江苏。</li>
	// <li>Anhui：安徽。</li>
	// <li>Shandong：山东。</li>
	// <li>Shanghai：上海。</li>
	// <li>Zhejiang：浙江。</li>
	// <li>Henan：河南。</li>
	// <li>Hubei：湖北。</li>
	// <li>Jiangxi：江西。</li>
	// <li>Hunan：湖南。</li>
	// <li>Guizhou：贵州。</li>
	// <li>Yunnan：云南。</li>
	// <li>Chongqing：重庆。</li>
	// <li>Sichuan：四川。</li>
	// <li>Tibet：西藏。</li>
	// <li>Guangdong：广东。</li>
	// <li>Guangxi：广西。</li>
	// <li>Hainan：海南。</li>
	// <li>Hong Kong, Macao and Taiwan：港澳台。</li>
	// <li>Outside Chinese Mainland：海外。</li>
	// <li>Other：其他 。</li>
	Districts []*string `json:"Districts,omitnil,omitempty" name:"Districts"`

	// 用户所属运营商信息，Area 为 Chinese Mainland 时，取值为以下运营商信息。当 Area 为其它值时忽略 Isps 参数。
	// <li>China Telecom：中国电信。 </li>
	// <li>China Unicom：中国联通。 </li>
	// <li>CERNET：教育网。</li>
	// <li>Great Wall Broadband Network：长城宽带。</li>
	// <li>China Mobile：中国移动。</li>
	// <li>China Mobile Tietong：中国铁通。</li>
	// <li>ISPs outside Chinese Mainland：海外运营商。</li>
	// <li>Other ISPs：其他运营商。</li>
	Isps []*string `json:"Isps,omitnil,omitempty" name:"Isps"`

	// 每条数据的时间粒度，单位：分钟，取值有：
	// <li>5：5 分钟粒度，返回指定查询时间内5分钟粒度的明细数据。</li>
	// <li>1440：天粒度，返回指定查询时间内1天粒度的数据。起始时间和结束时间跨度大于24小时，只支持天粒度的数据。</li>
	// 当 StartTime 和 EndTime 时间跨度大于24小时时，DataInterval 默认为 1440。
	DataInterval *uint64 `json:"DataInterval,omitnil,omitempty" name:"DataInterval"`
}

type DescribeCDNStatDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 查询指标，取值有：
	// <li>Traffic：流量，单位为 Byte。</li>
	// <li>Bandwidth：带宽，单位为 Bps。</li>
	// <li>Requests：请求数。</li>
	// <li>QUICRequests：QUIC 请求数。</li>
	// <li>IpVisit：IP 访问数。</li>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 起始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 域名列表。一次最多查询20个域名的数据。默认返回所有域名叠加的用量数据。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 服务区域，取值有：
	// <li>Chinese Mainland：中国大陆。 </li>
	// <li>Asia Pacific Region 1：亚太一区，包括中国香港、中国澳门、新加坡、越南、泰国。 </li>
	// <li>Asia Pacific Region 2：亚太二区，包括中国台湾、日本、马来西亚、印度尼西亚、韩国。 </li>
	// <li>Asia Pacific Region 3：亚太三区，包括菲律宾、印度、澳大利亚和亚太其它国家和地区。 </li>
	// <li>Middle East：中东。 </li>
	// <li>Europe：欧洲。</li>
	// <li>North America：北美。</li>
	// <li>South America：南美。</li>
	// <li>Africa：非洲。</li>
	// 默认为中国大陆。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 用户所在地区，Area 为 Chinese Mainland 时，取值为以下地区信息，当 Area 为其它值时， 忽略 Districts 参数。
	// <li>Beijing：北京。</li>
	// <li>Inner Mongolia：内蒙古。</li>
	// <li>Shanxi：山西。</li>
	// <li>Hebei：河北。</li>
	// <li>Tianjin：天津。</li>
	// <li>Ningxia：宁夏。</li>
	// <li>Shaanxi：陕西。</li>
	// <li>Gansu：甘肃。</li>
	// <li>Qinghai：青海。</li>
	// <li>Xinjiang：新疆。</li>
	// <li>Heilongjiang：黑龙江。</li>
	// <li>Jilin：吉林。</li>
	// <li>Liaoning：辽宁。</li>
	// <li>Fujian：福建。</li>
	// <li>Jiangsu：江苏。</li>
	// <li>Anhui：安徽。</li>
	// <li>Shandong：山东。</li>
	// <li>Shanghai：上海。</li>
	// <li>Zhejiang：浙江。</li>
	// <li>Henan：河南。</li>
	// <li>Hubei：湖北。</li>
	// <li>Jiangxi：江西。</li>
	// <li>Hunan：湖南。</li>
	// <li>Guizhou：贵州。</li>
	// <li>Yunnan：云南。</li>
	// <li>Chongqing：重庆。</li>
	// <li>Sichuan：四川。</li>
	// <li>Tibet：西藏。</li>
	// <li>Guangdong：广东。</li>
	// <li>Guangxi：广西。</li>
	// <li>Hainan：海南。</li>
	// <li>Hong Kong, Macao and Taiwan：港澳台。</li>
	// <li>Outside Chinese Mainland：海外。</li>
	// <li>Other：其他 。</li>
	Districts []*string `json:"Districts,omitnil,omitempty" name:"Districts"`

	// 用户所属运营商信息，Area 为 Chinese Mainland 时，取值为以下运营商信息。当 Area 为其它值时忽略 Isps 参数。
	// <li>China Telecom：中国电信。 </li>
	// <li>China Unicom：中国联通。 </li>
	// <li>CERNET：教育网。</li>
	// <li>Great Wall Broadband Network：长城宽带。</li>
	// <li>China Mobile：中国移动。</li>
	// <li>China Mobile Tietong：中国铁通。</li>
	// <li>ISPs outside Chinese Mainland：海外运营商。</li>
	// <li>Other ISPs：其他运营商。</li>
	Isps []*string `json:"Isps,omitnil,omitempty" name:"Isps"`

	// 每条数据的时间粒度，单位：分钟，取值有：
	// <li>5：5 分钟粒度，返回指定查询时间内5分钟粒度的明细数据。</li>
	// <li>1440：天粒度，返回指定查询时间内1天粒度的数据。起始时间和结束时间跨度大于24小时，只支持天粒度的数据。</li>
	// 当 StartTime 和 EndTime 时间跨度大于24小时时，DataInterval 默认为 1440。
	DataInterval *uint64 `json:"DataInterval,omitnil,omitempty" name:"DataInterval"`
}

func (r *DescribeCDNStatDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDNStatDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Metric")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "DomainNames")
	delete(f, "Area")
	delete(f, "Districts")
	delete(f, "Isps")
	delete(f, "DataInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCDNStatDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCDNStatDetailsResponseParams struct {
	// 每条数据的时间粒度，单位：分钟。
	DataInterval *uint64 `json:"DataInterval,omitnil,omitempty" name:"DataInterval"`

	// CDN 用量数据。
	Data []*StatDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCDNStatDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCDNStatDetailsResponseParams `json:"Response"`
}

func (r *DescribeCDNStatDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDNStatDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCDNUsageDataRequestParams struct {
	// 起始日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于开始日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// CDN 统计数据类型，有效值：
	// <li>Flux：流量，单位为 byte。</li>
	// <li>Bandwidth：带宽，单位为 bps。</li>
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 用量数据的时间粒度，单位：分钟，取值有：
	// <li>5：5 分钟粒度，返回指定查询时间内5分钟粒度的明细数据。</li>
	// <li>60：小时粒度，返回指定查询时间内1小时粒度的数据。</li>
	// <li>1440：天粒度，返回指定查询时间内1天粒度的数据。</li>
	// 默认值为1440，返回天粒度的数据。
	DataInterval *uint64 `json:"DataInterval,omitnil,omitempty" name:"DataInterval"`

	// 域名列表。一次最多查询20个域名的用量数据。可以指定多个域名，查询这些域名叠加的用量数据。默认返回所有域名叠加的用量数据。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`
}

type DescribeCDNUsageDataRequest struct {
	*tchttp.BaseRequest
	
	// 起始日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于开始日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// CDN 统计数据类型，有效值：
	// <li>Flux：流量，单位为 byte。</li>
	// <li>Bandwidth：带宽，单位为 bps。</li>
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 用量数据的时间粒度，单位：分钟，取值有：
	// <li>5：5 分钟粒度，返回指定查询时间内5分钟粒度的明细数据。</li>
	// <li>60：小时粒度，返回指定查询时间内1小时粒度的数据。</li>
	// <li>1440：天粒度，返回指定查询时间内1天粒度的数据。</li>
	// 默认值为1440，返回天粒度的数据。
	DataInterval *uint64 `json:"DataInterval,omitnil,omitempty" name:"DataInterval"`

	// 域名列表。一次最多查询20个域名的用量数据。可以指定多个域名，查询这些域名叠加的用量数据。默认返回所有域名叠加的用量数据。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`
}

func (r *DescribeCDNUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDNUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DataType")
	delete(f, "SubAppId")
	delete(f, "DataInterval")
	delete(f, "DomainNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCDNUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCDNUsageDataResponseParams struct {
	// 时间粒度，单位：分钟。
	DataInterval *int64 `json:"DataInterval,omitnil,omitempty" name:"DataInterval"`

	// CDN 统计数据。
	Data []*StatDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCDNUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCDNUsageDataResponseParams `json:"Response"`
}

func (r *DescribeCDNUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDNUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCLSLogsetsRequestParams struct {
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`
}

type DescribeCLSLogsetsRequest struct {
	*tchttp.BaseRequest
	
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`
}

func (r *DescribeCLSLogsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSLogsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CLSRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCLSLogsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCLSLogsetsResponseParams struct {
	// 查询到的日志集列表。
	Logsets []*CLSLogsetInfo `json:"Logsets,omitnil,omitempty" name:"Logsets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCLSLogsetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCLSLogsetsResponseParams `json:"Response"`
}

func (r *DescribeCLSLogsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSLogsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCLSPushTargetsRequestParams struct {
	// 点播域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeCLSPushTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 点播域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeCLSPushTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSPushTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCLSPushTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCLSPushTargetsResponseParams struct {
	// 域名推送总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 域名推送 CLS 目标列表。
	DomainCLSTargets []*DomainCLSTargetInfo `json:"DomainCLSTargets,omitnil,omitempty" name:"DomainCLSTargets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCLSPushTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCLSPushTargetsResponseParams `json:"Response"`
}

func (r *DescribeCLSPushTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSPushTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCLSTopicsRequestParams struct {
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`

	// 日志主题所属日志集 ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志主题 ID 列表。如果不填，表示查询所有的日志主题。
	TopicIds []*string `json:"TopicIds,omitnil,omitempty" name:"TopicIds"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCLSTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 日志集所属的地域，取值有： <li>ap-guangzhou：广州；</li> <li>ap-beijing：北京；</li> <li>ap-chengdu：成都；</li> <li>ap-chongqing：重庆；</li> <li>ap-nanjing：南京；</li> <li>ap-shanghai：上海；</li> <li>ap-singapore：新加坡。</li>
	CLSRegion *string `json:"CLSRegion,omitnil,omitempty" name:"CLSRegion"`

	// 日志主题所属日志集 ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志主题 ID 列表。如果不填，表示查询所有的日志主题。
	TopicIds []*string `json:"TopicIds,omitnil,omitempty" name:"TopicIds"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCLSTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CLSRegion")
	delete(f, "LogsetId")
	delete(f, "TopicIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCLSTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCLSTopicsResponseParams struct {
	// 日志主题总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志主题列表。
	Topics []*CLSTopicInfo `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCLSTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCLSTopicsResponseParams `json:"Response"`
}

func (r *DescribeCLSTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnLogsRequestParams struct {
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 获取日志起始时间点，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间需大于起始时间；使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 分页拉取的最大返回结果数。默认值：100；最大值：1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取的起始偏移量。默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCdnLogsRequest struct {
	*tchttp.BaseRequest
	
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 获取日志起始时间点，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间需大于起始时间；使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 分页拉取的最大返回结果数。默认值：100；最大值：1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取的起始偏移量。默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeCdnLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnLogsResponseParams struct {
	// 日志下载链接总数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 海外CDN节点的日志下载列表。如果域名没有开启海外加速，忽略该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverseaCdnLogs []*CdnLogInfo `json:"OverseaCdnLogs,omitnil,omitempty" name:"OverseaCdnLogs"`

	// 国内CDN节点的日志下载列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomesticCdnLogs []*CdnLogInfo `json:"DomesticCdnLogs,omitnil,omitempty" name:"DomesticCdnLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCdnLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnLogsResponseParams `json:"Response"`
}

func (r *DescribeCdnLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientUploadAccelerationUsageDataRequestParams struct {
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于等于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 客户端上传加速类型，取值有：
	// <li> AccelerationWithHTTP：HTTP 传输方式的上传加速。</li>
	// <li> AccelerationWithQUIC：QUIC 传输方式的上传加速。</li>
	// 默认查询所有加速类型的用量 。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeClientUploadAccelerationUsageDataRequest struct {
	*tchttp.BaseRequest
	
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于等于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 客户端上传加速类型，取值有：
	// <li> AccelerationWithHTTP：HTTP 传输方式的上传加速。</li>
	// <li> AccelerationWithQUIC：QUIC 传输方式的上传加速。</li>
	// 默认查询所有加速类型的用量 。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeClientUploadAccelerationUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientUploadAccelerationUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientUploadAccelerationUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientUploadAccelerationUsageDataResponseParams struct {
	// 客户端上传加速统计数据。
	ClientUploadAccelerationUsageDataSet []*StatDataItem `json:"ClientUploadAccelerationUsageDataSet,omitnil,omitempty" name:"ClientUploadAccelerationUsageDataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClientUploadAccelerationUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientUploadAccelerationUsageDataResponseParams `json:"Response"`
}

func (r *DescribeClientUploadAccelerationUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientUploadAccelerationUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentReviewTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 内容审核模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeContentReviewTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 内容审核模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContentReviewTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentReviewTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 内容审核模板详情列表。
	ContentReviewTemplateSet []*ContentReviewTemplateItem `json:"ContentReviewTemplateSet,omitnil,omitempty" name:"ContentReviewTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeCurrentPlaylistRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 轮播播单唯一标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// 返回的播放列表的长度。最大10，默认值为5。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCurrentPlaylistRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 轮播播单唯一标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// 返回的播放列表的长度。最大10，默认值为5。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCurrentPlaylistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentPlaylistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "RoundPlayId")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCurrentPlaylistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentPlaylistResponseParams struct {
	// 当前播放列表信息。
	CurrentPlaylist []*RoundPlayFilePlayInfo `json:"CurrentPlaylist,omitnil,omitempty" name:"CurrentPlaylist"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCurrentPlaylistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCurrentPlaylistResponseParams `json:"Response"`
}

func (r *DescribeCurrentPlaylistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentPlaylistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDailyMediaPlayStatRequestParams struct {
	// 媒体文件 ID 。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 起始日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。该参数仅日期部分有效。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。该参数仅日期部分有效。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeDailyMediaPlayStatRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID 。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 起始日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。该参数仅日期部分有效。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。该参数仅日期部分有效。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeDailyMediaPlayStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDailyMediaPlayStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDailyMediaPlayStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDailyMediaPlayStatResponseParams struct {
	// 播放统计数据。
	DailyPlayStatInfoSet []*DailyPlayStatInfo `json:"DailyPlayStatInfoSet,omitnil,omitempty" name:"DailyPlayStatInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDailyMediaPlayStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDailyMediaPlayStatResponseParams `json:"Response"`
}

func (r *DescribeDailyMediaPlayStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDailyMediaPlayStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDailyMostPlayedStatRequestParams struct {
	// 查询日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。该参数仅日期部分有效。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 域名。查询该域名播放 Top100 的媒体文件的统计数据。默认查询所有域名的播放统计数据。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Top 数据的统计指标，取值有：
	// <li>Traffic：播放流量，按播放流量统计 Top100 的数据。</li>
	// <li>PlayTimes：播放次数，按播放次数统计播放 Top100 的数据。</li>
	// 默认值为Traffic。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeDailyMostPlayedStatRequest struct {
	*tchttp.BaseRequest
	
	// 查询日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。该参数仅日期部分有效。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 域名。查询该域名播放 Top100 的媒体文件的统计数据。默认查询所有域名的播放统计数据。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Top 数据的统计指标，取值有：
	// <li>Traffic：播放流量，按播放流量统计 Top100 的数据。</li>
	// <li>PlayTimes：播放次数，按播放次数统计播放 Top100 的数据。</li>
	// 默认值为Traffic。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeDailyMostPlayedStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDailyMostPlayedStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Date")
	delete(f, "DomainName")
	delete(f, "Metric")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDailyMostPlayedStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDailyMostPlayedStatResponseParams struct {
	// 媒体文件播放统计信息。
	DailyPlayStatInfoSet []*DailyPlayStatInfo `json:"DailyPlayStatInfoSet,omitnil,omitempty" name:"DailyPlayStatInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDailyMostPlayedStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDailyMostPlayedStatResponseParams `json:"Response"`
}

func (r *DescribeDailyMostPlayedStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDailyMostPlayedStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDailyPlayStatFileListRequestParams struct {
	// 起始日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeDailyPlayStatFileListRequest struct {
	*tchttp.BaseRequest
	
	// 起始日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeDailyPlayStatFileListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDailyPlayStatFileListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDailyPlayStatFileListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDailyPlayStatFileListResponseParams struct {
	// 播放统计文件列表。
	PlayStatFileSet []*PlayStatFileInfo `json:"PlayStatFileSet,omitnil,omitempty" name:"PlayStatFileSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDailyPlayStatFileListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDailyPlayStatFileListResponseParams `json:"Response"`
}

func (r *DescribeDailyPlayStatFileListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDailyPlayStatFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultDistributionConfigRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeDefaultDistributionConfigRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeDefaultDistributionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultDistributionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultDistributionConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultDistributionConfigResponseParams struct {
	// 分发配置的域名(已废弃）。
	//
	// Deprecated: DomainName is deprecated.
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 分发配置的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分发配置的协议，为 HTTP 或 HTTPS。
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// 播放密钥，由大小写字母（a - Z）或者数字（0 - 9）组成，长度在8 - 20个字符之间。
	PlayKey *string `json:"PlayKey,omitnil,omitempty" name:"PlayKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDefaultDistributionConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultDistributionConfigResponseParams `json:"Response"`
}

func (r *DescribeDefaultDistributionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultDistributionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrmDataKeyRequestParams struct {
	// 加密后的数据密钥列表，最大支持10个。
	EdkList []*string `json:"EdkList,omitnil,omitempty" name:"EdkList"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeDrmDataKeyRequest struct {
	*tchttp.BaseRequest
	
	// 加密后的数据密钥列表，最大支持10个。
	EdkList []*string `json:"EdkList,omitnil,omitempty" name:"EdkList"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeDrmDataKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrmDataKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdkList")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDrmDataKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrmDataKeyResponseParams struct {
	// 密钥列表，包含加密的数据密钥。
	KeyList []*SimpleAesEdkPair `json:"KeyList,omitnil,omitempty" name:"KeyList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDrmDataKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDrmDataKeyResponseParams `json:"Response"`
}

func (r *DescribeDrmDataKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrmDataKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrmKeyProviderInfoRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeDrmKeyProviderInfoRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeDrmKeyProviderInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrmKeyProviderInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDrmKeyProviderInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrmKeyProviderInfoResponseParams struct {
	// 华曦达（SDMC）相关的 DRM 密钥提供商信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SDMCInfo *SDMCDrmKeyProviderInfo `json:"SDMCInfo,omitnil,omitempty" name:"SDMCInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDrmKeyProviderInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDrmKeyProviderInfoResponseParams `json:"Response"`
}

func (r *DescribeDrmKeyProviderInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrmKeyProviderInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnhanceMediaTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生模板列表。长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeEnhanceMediaTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生模板列表。长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeEnhanceMediaTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnhanceMediaTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnhanceMediaTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnhanceMediaTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 音画质重生模板详情列表。
	RebuildMediaTemplateSet []*RebuildMediaTemplate `json:"RebuildMediaTemplateSet,omitnil,omitempty" name:"RebuildMediaTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnhanceMediaTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnhanceMediaTemplatesResponseParams `json:"Response"`
}

func (r *DescribeEnhanceMediaTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnhanceMediaTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventConfigRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeEventConfigRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeEventConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventConfigResponseParams struct {
	// 接收事件通知的方式。"PUSH" 为 [HTTP 回调通知](https://cloud.tencent.com/document/product/266/7829#http.E5.9B.9E.E8.B0.83)，"PULL" 为 [基于消息队列的可靠通知](https://cloud.tencent.com/document/product/266/7829#.E5.9F.BA.E4.BA.8E.E6.B6.88.E6.81.AF.E9.98.9F.E5.88.97.E7.9A.84.E5.8F.AF.E9.9D.A0.E9.80.9A.E7.9F.A5)。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 采用 [HTTP 回调通知](https://cloud.tencent.com/document/product/266/7829#http.E5.9B.9E.E8.B0.83) 接收方式时，用于接收 V3 版本事件通知的地址。
	NotificationUrl *string `json:"NotificationUrl,omitnil,omitempty" name:"NotificationUrl"`

	// 是否接收 [视频上传完成](https://cloud.tencent.com/document/product/266/7830) 事件通知，"OFF" 为忽略该事件通知，"ON" 为接收事件通知。
	UploadMediaCompleteEventSwitch *string `json:"UploadMediaCompleteEventSwitch,omitnil,omitempty" name:"UploadMediaCompleteEventSwitch"`

	// 是否接收 [视频删除完成](https://cloud.tencent.com/document/product/266/13434) 事件通知，"OFF" 为忽略该事件通知，"ON" 为接收事件通知。
	DeleteMediaCompleteEventSwitch *string `json:"DeleteMediaCompleteEventSwitch,omitnil,omitempty" name:"DeleteMediaCompleteEventSwitch"`

	// 是否接收剪辑固化完成事件通知，"OFF" 为忽略该事件通知，"ON" 为接收事件通知。
	PersistenceCompleteEventSwitch *string `json:"PersistenceCompleteEventSwitch,omitnil,omitempty" name:"PersistenceCompleteEventSwitch"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventConfigResponseParams `json:"Response"`
}

func (r *DescribeEventConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsStateRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeEventsStateRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeEventsStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventsStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsStateResponseParams struct {
	// 待进行拉取的事件通知数，为近似值，约5秒延迟。
	CountOfEventsToPull *uint64 `json:"CountOfEventsToPull,omitnil,omitempty" name:"CountOfEventsToPull"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventsStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventsStateResponseParams `json:"Response"`
}

func (r *DescribeEventsStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileAttributesRequestParams struct {
	// 媒体文件 ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type DescribeFileAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *DescribeFileAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "TasksPriority")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileAttributesResponseParams struct {
	// 任务 ID 。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileAttributesResponseParams `json:"Response"`
}

func (r *DescribeFileAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileAttributesTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，0 表示成功，其他值表示失败：
	// <li>40000：输入参数不合法，请检查输入参数；</li>
	// <li>60000：源文件错误（如视频数据损坏），请确认源文件是否正常；</li>
	// <li>70000：内部服务错误，建议重试。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 获取媒体文件属性任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *DescribeFileAttributesTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type DescribeFileAttributesTaskOutput struct {
	// 媒体文件的 Md5 值。
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 媒体文件的 Sha1 值。
	Sha1 *string `json:"Sha1,omitnil,omitempty" name:"Sha1"`
}

// Predefined struct for user
type DescribeHeadTailTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 片头片尾模板号，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeHeadTailTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 片头片尾模板号，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeHeadTailTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHeadTailTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHeadTailTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHeadTailTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 片头片尾模板详情列表。
	HeadTailTemplateSet []*HeadTailTemplate `json:"HeadTailTemplateSet,omitnil,omitempty" name:"HeadTailTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHeadTailTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHeadTailTemplatesResponseParams `json:"Response"`
}

func (r *DescribeHeadTailTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHeadTailTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageProcessingTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 图片处理模板标识列表。长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeImageProcessingTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 图片处理模板标识列表。长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeImageProcessingTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageProcessingTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageProcessingTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageProcessingTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 图片处理模板详情列表。
	ImageProcessingTemplateSet []*ImageProcessingTemplate `json:"ImageProcessingTemplateSet,omitnil,omitempty" name:"ImageProcessingTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageProcessingTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageProcessingTemplatesResponseParams `json:"Response"`
}

func (r *DescribeImageProcessingTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageProcessingTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageReviewUsageDataRequestParams struct {
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于等于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeImageReviewUsageDataRequest struct {
	*tchttp.BaseRequest
	
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于等于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeImageReviewUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageReviewUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageReviewUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageReviewUsageDataResponseParams struct {
	// 图片审核次数统计数据，展示查询时间范围内的图片审核次数的概览数据。
	ImageReviewUsageDataSet []*ImageReviewUsageDataItem `json:"ImageReviewUsageDataSet,omitnil,omitempty" name:"ImageReviewUsageDataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageReviewUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageReviewUsageDataResponseParams `json:"Response"`
}

func (r *DescribeImageReviewUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageReviewUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSpriteTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 雪碧图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeImageSpriteTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 雪碧图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	delete(f, "SubAppId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 雪碧图模板详情列表。
	ImageSpriteTemplateSet []*ImageSpriteTemplate `json:"ImageSpriteTemplateSet,omitnil,omitempty" name:"ImageSpriteTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeJustInTimeTranscodeTemplatesRequestParams struct {
	// 模板名过滤条件，数组长度限制：100。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置任务流模板；</li>
	// <li>Custom：用户自定义任务流模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeJustInTimeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 模板名过滤条件，数组长度限制：100。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置任务流模板；</li>
	// <li>Custom：用户自定义任务流模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeJustInTimeTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJustInTimeTranscodeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Names")
	delete(f, "SubAppId")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJustInTimeTranscodeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJustInTimeTranscodeTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 模板详情列表。
	JustInTimeTranscodeTemplateSet []*JustInTimeTranscodeTemplate `json:"JustInTimeTranscodeTemplateSet,omitnil,omitempty" name:"JustInTimeTranscodeTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJustInTimeTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJustInTimeTranscodeTemplatesResponseParams `json:"Response"`
}

func (r *DescribeJustInTimeTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJustInTimeTranscodeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseUsageDataRequestParams struct {
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于等于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// License 类型，默认为 DRM 。目前支持的 License 类型包括：
	// <li> DRM: DRM 加密播放 License</li>
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeLicenseUsageDataRequest struct {
	*tchttp.BaseRequest
	
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于等于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// License 类型，默认为 DRM 。目前支持的 License 类型包括：
	// <li> DRM: DRM 加密播放 License</li>
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeLicenseUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "LicenseType")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLicenseUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseUsageDataResponseParams struct {
	// License 查询次数统计数据，展示所查询 License 次数的明细数据。
	LicenseUsageDataSet []*LicenseUsageDataItem `json:"LicenseUsageDataSet,omitnil,omitempty" name:"LicenseUsageDataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLicenseUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLicenseUsageDataResponseParams `json:"Response"`
}

func (r *DescribeLicenseUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaInfosRequestParams struct {
	// 媒体文件 ID 列表，N 从 0 开始取值，最大 19。
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定所有媒体文件需要返回的信息，可同时指定多个信息，N 从 0 开始递增。如果未填写该字段，默认返回所有信息。选项有：
	// <li>basicInfo（视频基础信息）。</li>
	// <li>metaData（视频元信息）。</li>
	// <li>transcodeInfo（视频转码结果信息）。</li>
	// <li>animatedGraphicsInfo（视频转动图结果信息）。</li>
	// <li>imageSpriteInfo（视频雪碧图信息）。</li>
	// <li>snapshotByTimeOffsetInfo（视频指定时间点截图信息）。</li>
	// <li>sampleSnapshotInfo（采样截图信息）。</li>
	// <li>keyFrameDescInfo（打点信息）。</li>
	// <li>adaptiveDynamicStreamingInfo（转自适应码流信息）。</li>
	// <li>miniProgramReviewInfo（小程序审核信息）。</li>
	// <li>subtitleInfo（字幕信息）。</li>
	// <li>reviewInfo（审核信息）。</li>
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeMediaInfosRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID 列表，N 从 0 开始取值，最大 19。
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定所有媒体文件需要返回的信息，可同时指定多个信息，N 从 0 开始递增。如果未填写该字段，默认返回所有信息。选项有：
	// <li>basicInfo（视频基础信息）。</li>
	// <li>metaData（视频元信息）。</li>
	// <li>transcodeInfo（视频转码结果信息）。</li>
	// <li>animatedGraphicsInfo（视频转动图结果信息）。</li>
	// <li>imageSpriteInfo（视频雪碧图信息）。</li>
	// <li>snapshotByTimeOffsetInfo（视频指定时间点截图信息）。</li>
	// <li>sampleSnapshotInfo（采样截图信息）。</li>
	// <li>keyFrameDescInfo（打点信息）。</li>
	// <li>adaptiveDynamicStreamingInfo（转自适应码流信息）。</li>
	// <li>miniProgramReviewInfo（小程序审核信息）。</li>
	// <li>subtitleInfo（字幕信息）。</li>
	// <li>reviewInfo（审核信息）。</li>
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeMediaInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileIds")
	delete(f, "SubAppId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaInfosResponseParams struct {
	// 媒体文件信息列表。
	MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitnil,omitempty" name:"MediaInfoSet"`

	// 不存在的文件 ID 列表。
	NotExistFileIdSet []*string `json:"NotExistFileIdSet,omitnil,omitempty" name:"NotExistFileIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMediaInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaInfosResponseParams `json:"Response"`
}

func (r *DescribeMediaInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaPlayStatDetailsRequestParams struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 起始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 统计时间粒度，有效值：
	// <li>Hour：以小时为粒度。</li>
	// <li>Day：以天为粒度。</li>
	// 默认按时间跨度决定，小于1天以小时为粒度，大于等于1天则以天为粒度。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`
}

type DescribeMediaPlayStatDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 起始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 统计时间粒度，有效值：
	// <li>Hour：以小时为粒度。</li>
	// <li>Day：以天为粒度。</li>
	// 默认按时间跨度决定，小于1天以小时为粒度，大于等于1天则以天为粒度。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`
}

func (r *DescribeMediaPlayStatDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaPlayStatDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaPlayStatDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaPlayStatDetailsResponseParams struct {
	// 播放统计数据。
	PlayStatInfoSet []*PlayStatInfo `json:"PlayStatInfoSet,omitnil,omitempty" name:"PlayStatInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMediaPlayStatDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaPlayStatDetailsResponseParams `json:"Response"`
}

func (r *DescribeMediaPlayStatDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaPlayStatDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaProcessUsageDataRequestParams struct {
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于等于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 查询视频处理任务类型，目前支持的任务类型包括：
	// <li> Transcoding: 普通转码</li>
	// <li> Transcoding-TESHD: 极速高清转码</li>
	// <li> Editing: 视频编辑</li>
	// <li> Editing-TESHD: 极速高清视频编辑</li>
	// <li> AdaptiveBitrateStreaming: 自适应码流</li>
	// <li> ContentAudit: 内容审核</li>
	// <li> ContentRecognition: 内容识别</li>
	// <li> RemoveWatermark: 去除水印</li>
	// <li> ExtractTraceWatermark: 提取水印</li>
	// <li> AddTraceWatermark: 添加水印</li>
	// <li> RebuildMedia: 音画质重生</li>
	// <li> QualityInspect: 音画质检测</li>
	// <li> VideoHighlight: 视频智能集锦</li>
	// <li> VideoTag: 视频智能标签</li>
	// <li> VideoClassification:  视频智能分类</li>
	// <li> VideoCover: 视频智能封面</li>
	// <li> VideoSegment: 视频智能拆条</li>
	// <li> VideoProduce: 视频制作</li>
	// <li> MediaCast: 媒体转推</li>
	// <li>Transcode: 转码，包含普通转码、极速高清和视频编辑（不推荐使用）</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeMediaProcessUsageDataRequest struct {
	*tchttp.BaseRequest
	
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于等于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 查询视频处理任务类型，目前支持的任务类型包括：
	// <li> Transcoding: 普通转码</li>
	// <li> Transcoding-TESHD: 极速高清转码</li>
	// <li> Editing: 视频编辑</li>
	// <li> Editing-TESHD: 极速高清视频编辑</li>
	// <li> AdaptiveBitrateStreaming: 自适应码流</li>
	// <li> ContentAudit: 内容审核</li>
	// <li> ContentRecognition: 内容识别</li>
	// <li> RemoveWatermark: 去除水印</li>
	// <li> ExtractTraceWatermark: 提取水印</li>
	// <li> AddTraceWatermark: 添加水印</li>
	// <li> RebuildMedia: 音画质重生</li>
	// <li> QualityInspect: 音画质检测</li>
	// <li> VideoHighlight: 视频智能集锦</li>
	// <li> VideoTag: 视频智能标签</li>
	// <li> VideoClassification:  视频智能分类</li>
	// <li> VideoCover: 视频智能封面</li>
	// <li> VideoSegment: 视频智能拆条</li>
	// <li> VideoProduce: 视频制作</li>
	// <li> MediaCast: 媒体转推</li>
	// <li>Transcode: 转码，包含普通转码、极速高清和视频编辑（不推荐使用）</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeMediaProcessUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaProcessUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaProcessUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaProcessUsageDataResponseParams struct {
	// 视频处理统计数据概览，展示所查询任务的概览以及详细数据。
	MediaProcessDataSet []*TaskStatData `json:"MediaProcessDataSet,omitnil,omitempty" name:"MediaProcessDataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMediaProcessUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaProcessUsageDataResponseParams `json:"Response"`
}

func (r *DescribeMediaProcessUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaProcessUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonSamplesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 拉取的素材类型，可选值：
	// <li>UserDefine：用户自定义素材库；</li>
	// <li>Default：系统默认素材库。</li>
	// 
	// 默认值：UserDefine，拉取用户自定义素材库素材。
	// 说明：如果是拉取系统默认素材库，只能使用素材名字或者素材 ID + 素材名字的方式进行拉取，且五官图片只返回一张。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 素材 ID，数组长度限制：100。
	PersonIds []*string `json:"PersonIds,omitnil,omitempty" name:"PersonIds"`

	// 素材名称，数组长度限制：20。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 素材标签，数组长度限制：20。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePersonSamplesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 拉取的素材类型，可选值：
	// <li>UserDefine：用户自定义素材库；</li>
	// <li>Default：系统默认素材库。</li>
	// 
	// 默认值：UserDefine，拉取用户自定义素材库素材。
	// 说明：如果是拉取系统默认素材库，只能使用素材名字或者素材 ID + 素材名字的方式进行拉取，且五官图片只返回一张。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 素材 ID，数组长度限制：100。
	PersonIds []*string `json:"PersonIds,omitnil,omitempty" name:"PersonIds"`

	// 素材名称，数组长度限制：20。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 素材标签，数组长度限制：20。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "SubAppId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 人物信息。
	PersonSet []*AiSamplePerson `json:"PersonSet,omitnil,omitempty" name:"PersonSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribePrepaidProductsRequestParams struct {

}

type DescribePrepaidProductsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePrepaidProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrepaidProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrepaidProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrepaidProductsResponseParams struct {
	// 购买的预付费商品实例列表。
	ProductInstanceSet []*ProductInstance `json:"ProductInstanceSet,omitnil,omitempty" name:"ProductInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrepaidProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrepaidProductsResponseParams `json:"Response"`
}

func (r *DescribePrepaidProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrepaidProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcedureTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 任务流模板名字过滤条件，数组长度限制：100。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 任务流模板类型过滤条件，可选值：
	// <li>Preset：系统预置任务流模板；</li>
	// <li>Custom：用户自定义任务流模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeProcedureTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 任务流模板名字过滤条件，数组长度限制：100。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 任务流模板类型过滤条件，可选值：
	// <li>Preset：系统预置任务流模板；</li>
	// <li>Custom：用户自定义任务流模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeProcedureTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcedureTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Names")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcedureTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcedureTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务流模板详情列表。
	ProcedureTemplateSet []*ProcedureTemplate `json:"ProcedureTemplateSet,omitnil,omitempty" name:"ProcedureTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProcedureTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProcedureTemplatesResponseParams `json:"Response"`
}

func (r *DescribeProcedureTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcedureTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityInspectTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质检测模板列表。长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeQualityInspectTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质检测模板列表。长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeQualityInspectTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityInspectTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQualityInspectTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityInspectTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 音画质检测模板详情列表。
	QualityInspectTemplateSet []*QualityInspectTemplateItem `json:"QualityInspectTemplateSet,omitnil,omitempty" name:"QualityInspectTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQualityInspectTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQualityInspectTemplatesResponseParams `json:"Response"`
}

func (r *DescribeQualityInspectTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityInspectTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRebuildMediaTemplatesRequestParams struct {
	// 视频重生模板列表。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRebuildMediaTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 视频重生模板列表。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRebuildMediaTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRebuildMediaTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "SubAppId")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRebuildMediaTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRebuildMediaTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 视频重生模板详情列表。
	RebuildMediaTemplateSet []*RebuildMediaTemplate `json:"RebuildMediaTemplateSet,omitnil,omitempty" name:"RebuildMediaTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRebuildMediaTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRebuildMediaTemplatesResponseParams `json:"Response"`
}

func (r *DescribeRebuildMediaTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRebuildMediaTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReviewDetailsRequestParams struct {
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeReviewDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，需大于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeReviewDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReviewDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReviewDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReviewDetailsResponseParams struct {
	// 发起内容智能识别次数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 内容智能识别总时长。
	TotalDuration *int64 `json:"TotalDuration,omitnil,omitempty" name:"TotalDuration"`

	// 内容智能识别时长统计数据，每天一个数据。
	Data []*StatDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReviewDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReviewDetailsResponseParams `json:"Response"`
}

func (r *DescribeReviewDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReviewDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReviewTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 审核模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeReviewTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 审核模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeReviewTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReviewTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReviewTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReviewTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审核模板详情列表。
	ReviewTemplateSet []*ReviewTemplate `json:"ReviewTemplateSet,omitnil,omitempty" name:"ReviewTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReviewTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReviewTemplatesResponseParams `json:"Response"`
}

func (r *DescribeReviewTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReviewTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoundPlaysRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 过滤条件：轮播播单标识，数组长度限制：100。
	RoundPlayIds []*string `json:"RoundPlayIds,omitnil,omitempty" name:"RoundPlayIds"`

	// 过滤条件，轮播播单状态，可选值： <li>Enabled：启动状态；</li> <li>Disabled：停止状态。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 过滤条件：轮播播单创建时间。
	CreateTime *TimeRange `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过滤条件：轮播播单更新时间。
	UpdateTime *TimeRange `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 分页偏移量，默认值：0。已经废弃，请根据 ScrollToken 参数进行分批次查询。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRoundPlaysRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 过滤条件：轮播播单标识，数组长度限制：100。
	RoundPlayIds []*string `json:"RoundPlayIds,omitnil,omitempty" name:"RoundPlayIds"`

	// 过滤条件，轮播播单状态，可选值： <li>Enabled：启动状态；</li> <li>Disabled：停止状态。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 过滤条件：轮播播单创建时间。
	CreateTime *TimeRange `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过滤条件：轮播播单更新时间。
	UpdateTime *TimeRange `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 分页偏移量，默认值：0。已经废弃，请根据 ScrollToken 参数进行分批次查询。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRoundPlaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoundPlaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "RoundPlayIds")
	delete(f, "Status")
	delete(f, "CreateTime")
	delete(f, "UpdateTime")
	delete(f, "ScrollToken")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoundPlaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoundPlaysResponseParams struct {
	// 符合过滤条件的轮播播单总数。已经废弃，分批次查询请请使用 ScrollToken 参数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 轮播播单详情列表。
	RoundPlaySet []*RoundPlayInfo `json:"RoundPlaySet,omitnil,omitempty" name:"RoundPlaySet"`

	// 翻页标识，当请求未返回所有数据，该字段表示下一条记录的 ID。当该字段为空，说明已无更多数据。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRoundPlaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoundPlaysResponseParams `json:"Response"`
}

func (r *DescribeRoundPlaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoundPlaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSampleSnapshotTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 采样截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeSampleSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 采样截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	delete(f, "SubAppId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 采样截图模板详情列表。
	SampleSnapshotTemplateSet []*SampleSnapshotTemplate `json:"SampleSnapshotTemplateSet,omitnil,omitempty" name:"SampleSnapshotTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSnapshotByTimeOffsetTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定时间点截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeSnapshotByTimeOffsetTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定时间点截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	delete(f, "SubAppId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 指定时间点截图模板详情列表。
	SnapshotByTimeOffsetTemplateSet []*SnapshotByTimeOffsetTemplate `json:"SnapshotByTimeOffsetTemplateSet,omitnil,omitempty" name:"SnapshotByTimeOffsetTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeStorageDataRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeStorageDataRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeStorageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageDataResponseParams struct {
	// 当前媒体总量。
	MediaCount *uint64 `json:"MediaCount,omitnil,omitempty" name:"MediaCount"`

	// 当前总存储量，单位是字节。
	TotalStorage *uint64 `json:"TotalStorage,omitnil,omitempty" name:"TotalStorage"`

	// 当前标准存储量，单位是字节。
	StandardStorage *uint64 `json:"StandardStorage,omitnil,omitempty" name:"StandardStorage"`

	// 当前低频存储量，单位是字节。
	InfrequentStorage *uint64 `json:"InfrequentStorage,omitnil,omitempty" name:"InfrequentStorage"`

	// 当前归档存储量，单位是字节。
	ArchiveStorage *uint64 `json:"ArchiveStorage,omitnil,omitempty" name:"ArchiveStorage"`

	// 当前深度归档存储量，单位是字节。
	DeepArchiveStorage *uint64 `json:"DeepArchiveStorage,omitnil,omitempty" name:"DeepArchiveStorage"`

	// 各计费区域的存储用量。
	StorageStat []*StorageStatData `json:"StorageStat,omitnil,omitempty" name:"StorageStat"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStorageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageDataResponseParams `json:"Response"`
}

func (r *DescribeStorageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageDetailsRequestParams struct {
	// 起始时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，需大于开始日期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 统计时间粒度，有效值：
	// <li>Minute：以5分钟为粒度。</li>
	// <li>Day：以天为粒度。</li>
	// 默认按时间跨度决定，小于等于1天以5分钟为粒度，大于1天则以天为粒度。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 查询的存储类型，有效值：
	// <li>TotalStorage：存储总量，标准、低频、归档和深度归档存储量之和，不含提前删除量。</li>
	// <li>StandardStorage：标准存储。</li>
	// <li>InfrequentStorage：低频存储。</li>
	// <li>ArchiveStorage：归档存储。</li>
	// <li>DeepArchiveStorage：深度归档存储。</li>
	// <li>DeletedInfrequentStorage：低频存储提前删除量。</li>
	// <li>DeletedArchiveStorage：归档提前删除量。</li>
	// <li>DeletedDeepArchiveStorage：深度归档提前删除量。</li>
	// <li>ArchiveStandardRetrieval：归档标准取回量。</li>
	// <li>ArchiveExpeditedRetrieval：归档快速取回量。</li>
	// <li>ArchiveBulkRetrieval：归档批量取回量。</li>
	// <li>DeepArchiveStandardRetrieval：深度归档标准取回量。</li>
	// <li>DeepArchiveBulkRetrieval：深度归档批量取回量。</li>
	// <li>InfrequentRetrieval：低频存储取回量。</li>
	// 默认值为 TotalStorage。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 查询的存储区域，有效值：
	// <li>Chinese Mainland：中国境内（不包含港澳台）。</li>
	// <li>Outside Chinese Mainland：中国境外。</li>
	// 默认值为 Chinese Mainland。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeStorageDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，需大于开始日期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 统计时间粒度，有效值：
	// <li>Minute：以5分钟为粒度。</li>
	// <li>Day：以天为粒度。</li>
	// 默认按时间跨度决定，小于等于1天以5分钟为粒度，大于1天则以天为粒度。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 查询的存储类型，有效值：
	// <li>TotalStorage：存储总量，标准、低频、归档和深度归档存储量之和，不含提前删除量。</li>
	// <li>StandardStorage：标准存储。</li>
	// <li>InfrequentStorage：低频存储。</li>
	// <li>ArchiveStorage：归档存储。</li>
	// <li>DeepArchiveStorage：深度归档存储。</li>
	// <li>DeletedInfrequentStorage：低频存储提前删除量。</li>
	// <li>DeletedArchiveStorage：归档提前删除量。</li>
	// <li>DeletedDeepArchiveStorage：深度归档提前删除量。</li>
	// <li>ArchiveStandardRetrieval：归档标准取回量。</li>
	// <li>ArchiveExpeditedRetrieval：归档快速取回量。</li>
	// <li>ArchiveBulkRetrieval：归档批量取回量。</li>
	// <li>DeepArchiveStandardRetrieval：深度归档标准取回量。</li>
	// <li>DeepArchiveBulkRetrieval：深度归档批量取回量。</li>
	// <li>InfrequentRetrieval：低频存储取回量。</li>
	// 默认值为 TotalStorage。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 查询的存储区域，有效值：
	// <li>Chinese Mainland：中国境内（不包含港澳台）。</li>
	// <li>Outside Chinese Mainland：中国境外。</li>
	// 默认值为 Chinese Mainland。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

func (r *DescribeStorageDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "Interval")
	delete(f, "StorageType")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageDetailsResponseParams struct {
	// 存储统计数据，每5分钟或每天一条数据。
	Data []*StatDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStorageDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageDetailsResponseParams `json:"Response"`
}

func (r *DescribeStorageDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageRegionsRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeStorageRegionsRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeStorageRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageRegionsResponseParams struct {
	// 存储地域信息列表。
	StorageRegionInfos []*StorageRegionInfo `json:"StorageRegionInfos,omitnil,omitempty" name:"StorageRegionInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStorageRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageRegionsResponseParams `json:"Response"`
}

func (r *DescribeStorageRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAppIdsRequestParams struct {
	// 应用名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签信息，查询指定标签的应用列表。
	Tags []*ResourceTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分页拉取的起始偏移量。默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页拉取的最大返回结果数。默认值：200；最大值：200。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSubAppIdsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签信息，查询指定标签的应用列表。
	Tags []*ResourceTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分页拉取的起始偏移量。默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页拉取的最大返回结果数。默认值：200；最大值：200。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSubAppIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAppIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Tags")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubAppIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAppIdsResponseParams struct {
	// 应用信息集合。
	SubAppIdInfoSet []*SubAppIdInfo `json:"SubAppIdInfoSet,omitnil,omitempty" name:"SubAppIdInfoSet"`

	// 应用总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubAppIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubAppIdsResponseParams `json:"Response"`
}

func (r *DescribeSubAppIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAppIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSuperPlayerConfigsRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 播放器配置名字过滤条件，数组长度限制：100。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 播放器配置类型过滤条件，可选值：
	// <li>Preset：系统预置配置；</li>
	// <li>Custom：用户自定义配置。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeSuperPlayerConfigsRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 播放器配置名字过滤条件，数组长度限制：100。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 播放器配置类型过滤条件，可选值：
	// <li>Preset：系统预置配置；</li>
	// <li>Custom：用户自定义配置。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeSuperPlayerConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSuperPlayerConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Names")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSuperPlayerConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSuperPlayerConfigsResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 播放器配置数组。
	PlayerConfigSet []*PlayerConfig `json:"PlayerConfigSet,omitnil,omitempty" name:"PlayerConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSuperPlayerConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSuperPlayerConfigsResponseParams `json:"Response"`
}

func (r *DescribeSuperPlayerConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSuperPlayerConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// 视频处理任务的任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 视频处理任务的任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
	// 任务类型，取值：<li>Procedure：视频处理任务；</li><li>EditMedia：视频编辑任务；</li><li>SplitMedia：视频拆条任务；</li><li>ComposeMedia：制作媒体文件任务；</li><li>WechatPublish：微信发布任务；</li><li>WechatMiniProgramPublish：微信小程序视频发布任务；</li><li>PullUpload：拉取上传媒体文件任务；</li><li>FastClipMedia：快速剪辑任务；</li><li>RemoveWatermarkTask：智能去除水印任务；</li><li>DescribeFileAttributesTask：获取文件属性任务；</li><li>RebuildMedia：音画质重生任务（不推荐使用）；</li><li>ReviewAudioVideo：音视频审核任务；</li><li>ExtractTraceWatermark：提取溯源水印任务；</li><li>ExtractCopyRightWatermark：提取版权水印任务；</li><li>QualityInspect：音画质检测任务；</li><li>QualityEnhance：音画质重生任务；</li><li>ComplexAdaptiveDynamicStreaming：复杂自适应码流任务。</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务状态，取值：
	// <li>WAITING：等待中；</li>
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成；</li>
	// <li>ABORTED：已终止。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务的创建时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 视频处理任务信息，仅当 TaskType 为 Procedure，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcedureTask *ProcedureTask `json:"ProcedureTask,omitnil,omitempty" name:"ProcedureTask"`

	// 视频编辑任务信息，仅当 TaskType 为 EditMedia，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditMediaTask *EditMediaTask `json:"EditMediaTask,omitnil,omitempty" name:"EditMediaTask"`

	// 微信发布任务信息，仅当 TaskType 为 WechatPublish，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatPublishTask *WechatPublishTask `json:"WechatPublishTask,omitnil,omitempty" name:"WechatPublishTask"`

	// 制作媒体文件任务信息，仅当 TaskType 为 ComposeMedia，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComposeMediaTask *ComposeMediaTask `json:"ComposeMediaTask,omitnil,omitempty" name:"ComposeMediaTask"`

	// 视频拆条任务信息，仅当 TaskType 为 SplitMedia，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SplitMediaTask *SplitMediaTask `json:"SplitMediaTask,omitnil,omitempty" name:"SplitMediaTask"`

	// 微信小程序发布任务信息，仅当 TaskType 为 WechatMiniProgramPublish，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatMiniProgramPublishTask *WechatMiniProgramPublishTask `json:"WechatMiniProgramPublishTask,omitnil,omitempty" name:"WechatMiniProgramPublishTask"`

	// 拉取上传媒体文件任务信息，仅当 TaskType 为 PullUpload，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullUploadTask *PullUploadTask `json:"PullUploadTask,omitnil,omitempty" name:"PullUploadTask"`

	// 视频转码任务信息，仅当 TaskType 为 Transcode，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTask *TranscodeTask2017 `json:"TranscodeTask,omitnil,omitempty" name:"TranscodeTask"`

	// 视频拼接任务信息，仅当 TaskType 为 Concat，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConcatTask *ConcatTask2017 `json:"ConcatTask,omitnil,omitempty" name:"ConcatTask"`

	// 视频剪辑任务信息，仅当 TaskType 为 Clip，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClipTask *ClipTask2017 `json:"ClipTask,omitnil,omitempty" name:"ClipTask"`

	// 截取雪碧图任务信息，仅当 TaskType 为 ImageSprite，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateImageSpriteTask *CreateImageSpriteTask2017 `json:"CreateImageSpriteTask,omitnil,omitempty" name:"CreateImageSpriteTask"`

	// 视频指定时间点截图任务信息，仅当 TaskType 为 SnapshotByTimeOffset，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTask *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetTask,omitnil,omitempty" name:"SnapshotByTimeOffsetTask"`

	// 智能去除水印任务信息，仅当 TaskType 为 RemoveWatermark，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoveWatermarkTask *RemoveWatermarkTask `json:"RemoveWatermarkTask,omitnil,omitempty" name:"RemoveWatermarkTask"`

	// 音画质重生任务信息，仅当 TaskType 为 RebuildMedia，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RebuildMediaTask *RebuildMediaTask `json:"RebuildMediaTask,omitnil,omitempty" name:"RebuildMediaTask"`

	// 提取溯源水印任务信息，仅当 TaskType 为 ExtractTraceWatermark，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtractTraceWatermarkTask *ExtractTraceWatermarkTask `json:"ExtractTraceWatermarkTask,omitnil,omitempty" name:"ExtractTraceWatermarkTask"`

	// 提取版权水印任务信息，仅当 TaskType 为 ExtractCopyRightWatermark，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtractCopyRightWatermarkTask *ExtractCopyRightWatermarkTask `json:"ExtractCopyRightWatermarkTask,omitnil,omitempty" name:"ExtractCopyRightWatermarkTask"`

	// 音视频审核任务信息，仅当 TaskType 为 ReviewAudioVideo，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewAudioVideoTask *ReviewAudioVideoTask `json:"ReviewAudioVideoTask,omitnil,omitempty" name:"ReviewAudioVideoTask"`

	// 该字段已无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReduceMediaBitrateTask *ReduceMediaBitrateTask `json:"ReduceMediaBitrateTask,omitnil,omitempty" name:"ReduceMediaBitrateTask"`

	// 获取文件属性任务信息，仅当 TaskType 为 DescribeFileAttributes，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescribeFileAttributesTask *DescribeFileAttributesTask `json:"DescribeFileAttributesTask,omitnil,omitempty" name:"DescribeFileAttributesTask"`

	// 音画质检测任务信息，仅当 TaskType 为 QualityInspect 时该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityInspectTask *QualityInspectTask `json:"QualityInspectTask,omitnil,omitempty" name:"QualityInspectTask"`

	// 音画质重生任务信息，仅当 TaskType 为 QualityEnhance，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityEnhanceTask *QualityEnhanceTask `json:"QualityEnhanceTask,omitnil,omitempty" name:"QualityEnhanceTask"`

	// 复杂自适应码流任务信息，仅当 TaskType 为 ComplexAdaptiveDynamicStreaming，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplexAdaptiveDynamicStreamingTask *ComplexAdaptiveDynamicStreamingTask `json:"ComplexAdaptiveDynamicStreamingTask,omitnil,omitempty" name:"ComplexAdaptiveDynamicStreamingTask"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 过滤条件：任务状态，可选值：WAITING（等待中）、PROCESSING（处理中）、FINISH（已完成）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 过滤条件：文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 过滤条件：任务创建时间。
	CreateTime *TimeRange `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过滤条件：任务结束时间。
	FinishTime *TimeRange `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 排序方式。Sort.Field 可选：
	// <li> CreateTime 任务创建时间。</li>
	// <li>FinishTime 任务结束时间。</li>
	Sort *SortBy `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 过滤条件：任务状态，可选值：WAITING（等待中）、PROCESSING（处理中）、FINISH（已完成）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 过滤条件：文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 过滤条件：任务创建时间。
	CreateTime *TimeRange `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过滤条件：任务结束时间。
	FinishTime *TimeRange `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 排序方式。Sort.Field 可选：
	// <li> CreateTime 任务创建时间。</li>
	// <li>FinishTime 任务结束时间。</li>
	Sort *SortBy `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`
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
	delete(f, "SubAppId")
	delete(f, "Status")
	delete(f, "FileId")
	delete(f, "CreateTime")
	delete(f, "FinishTime")
	delete(f, "Sort")
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
	TaskSet []*TaskSimpleInfo `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 翻页标识，当请求未返回所有数据，该字段表示下一条记录的 ID。当该字段为空，说明已无更多数据。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转码模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式板；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式。</li>
	ContainerType *string `json:"ContainerType,omitnil,omitempty" name:"ContainerType"`

	// 极速高清过滤条件，用于过滤普通转码或极速高清转码模板，可选值：
	// <li>Common：普通转码模板；</li>
	// <li>TEHD：极速高清模板。</li>
	TEHDType *string `json:"TEHDType,omitnil,omitempty" name:"TEHDType"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转码模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式板；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式。</li>
	ContainerType *string `json:"ContainerType,omitnil,omitempty" name:"ContainerType"`

	// 极速高清过滤条件，用于过滤普通转码或极速高清转码模板，可选值：
	// <li>Common：普通转码模板；</li>
	// <li>TEHD：极速高清模板。</li>
	TEHDType *string `json:"TEHDType,omitnil,omitempty" name:"TEHDType"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "ContainerType")
	delete(f, "TEHDType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 转码模板详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTemplateSet []*TranscodeTemplate `json:"TranscodeTemplateSet,omitnil,omitempty" name:"TranscodeTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeVodDomainsRequestParams struct {
	// 域名列表。当该字段不填时，则默认列出所有域名信息。本字段限制如下：
	// <li>域名个数度最大为 20。</li>
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 分页拉取的最大返回结果数。默认值：20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取的起始偏移量。默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type DescribeVodDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 域名列表。当该字段不填时，则默认列出所有域名信息。本字段限制如下：
	// <li>域名个数度最大为 20。</li>
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 分页拉取的最大返回结果数。默认值：20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取的起始偏移量。默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *DescribeVodDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVodDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVodDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVodDomainsResponseParams struct {
	// 域名总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 域名信息列表。
	DomainSet []*DomainDetailInfo `json:"DomainSet,omitnil,omitempty" name:"DomainSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVodDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVodDomainsResponseParams `json:"Response"`
}

func (r *DescribeVodDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVodDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWatermarkTemplatesRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 水印类型过滤条件，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印；</li>
	// <li>svg：SVG 水印。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 水印模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 返回记录条数
	// <li>默认值：10；</li>
	// <li>最大值：100。</li>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeWatermarkTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 水印类型过滤条件，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印；</li>
	// <li>svg：SVG 水印。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 水印模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// 返回记录条数
	// <li>默认值：10；</li>
	// <li>最大值：100。</li>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "SubAppId")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Definitions")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWatermarkTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWatermarkTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 水印模板详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkTemplateSet []*WatermarkTemplate `json:"WatermarkTemplateSet,omitnil,omitempty" name:"WatermarkTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// <b>关键词应用场景过滤条件，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜的内容识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜的内容识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜的内容识别，等价于 3+4；
	// 可多选，元素间关系为 or，即关键词的应用场景包含该字段集合中任意元素的记录，均符合该条件。
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 关键词过滤条件，数组长度限制：100 个词。
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 标签过滤条件，数组长度限制：20 个词。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeWordSamplesRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// <b>关键词应用场景过滤条件，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜的内容识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜的内容识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜的内容识别，等价于 3+4；
	// 可多选，元素间关系为 or，即关键词的应用场景包含该字段集合中任意元素的记录，均符合该条件。
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 关键词过滤条件，数组长度限制：100 个词。
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 标签过滤条件，数组长度限制：20 个词。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "SubAppId")
	delete(f, "Usages")
	delete(f, "Keywords")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 关键词信息。
	WordSet []*AiSampleWord `json:"WordSet,omitnil,omitempty" name:"WordSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DomainCLSTargetInfo struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 中国大陆地区的日志推送目标。
	ChineseMainlandCLSTargetInfo *AreaCLSTargetInfo `json:"ChineseMainlandCLSTargetInfo,omitnil,omitempty" name:"ChineseMainlandCLSTargetInfo"`

	// 中国大陆以外地区的日志推送目标。
	OutsideChineseMainlandCLSTargetInfo *AreaCLSTargetInfo `json:"OutsideChineseMainlandCLSTargetInfo,omitnil,omitempty" name:"OutsideChineseMainlandCLSTargetInfo"`
}

type DomainDetailInfo struct {
	// 域名名称。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 加速地区信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateAreaInfos []*AccelerateAreaInfo `json:"AccelerateAreaInfos,omitnil,omitempty" name:"AccelerateAreaInfos"`

	// 部署状态，取值有：
	// <li>Online：上线；</li>
	// <li>Deploying：部署中；</li>
	// <li>Locked: 锁定中，出现该状态时，无法对该域名进行部署变更。</li>
	DeployStatus *string `json:"DeployStatus,omitnil,omitempty" name:"DeployStatus"`

	// HTTPS 配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPSConfig *DomainHTTPSConfig `json:"HTTPSConfig,omitnil,omitempty" name:"HTTPSConfig"`

	// [Key 防盗链](https://cloud.tencent.com/document/product/266/14047)配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlSignatureAuthPolicy *UrlSignatureAuthPolicy `json:"UrlSignatureAuthPolicy,omitnil,omitempty" name:"UrlSignatureAuthPolicy"`

	// [Referer 防盗链](https://cloud.tencent.com/document/product/266/14046)配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefererAuthPolicy *RefererAuthPolicy `json:"RefererAuthPolicy,omitnil,omitempty" name:"RefererAuthPolicy"`

	// 域名添加到腾讯云点播系统中的时间。
	// <li>格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。</li>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 域名 QUIC 配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QUICConfig *DomainQUICConfig `json:"QUICConfig,omitnil,omitempty" name:"QUICConfig"`
}

type DomainHTTPSConfig struct {
	// 证书过期时间。
	// <li>格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。</li>
	CertExpireTime *string `json:"CertExpireTime,omitnil,omitempty" name:"CertExpireTime"`

	// 腾讯云 SSL 产品中的证书 ID。
	CloudCertId *string `json:"CloudCertId,omitnil,omitempty" name:"CloudCertId"`
}

type DomainQUICConfig struct {
	// QUIC 配置状态，可选值：
	// <li>Enabled: 启用；</li>
	// <li>Disabled: 禁用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DrmStreamingsInfo struct {
	// 保护类型为 SimpleAES 的转自适应码流模板 ID。
	SimpleAesDefinition *uint64 `json:"SimpleAesDefinition,omitnil,omitempty" name:"SimpleAesDefinition"`

	// 保护类型为 Widevine 的转自适应码流模板 ID。
	WidevineDefinition *uint64 `json:"WidevineDefinition,omitnil,omitempty" name:"WidevineDefinition"`

	// 保护类型为 FairPlay 的转自适应码流模板 ID。
	FairPlayDefinition *uint64 `json:"FairPlayDefinition,omitnil,omitempty" name:"FairPlayDefinition"`
}

type DrmStreamingsInfoForUpdate struct {
	// 保护类型为 SimpleAES 的转自适应码流模板 ID。
	SimpleAesDefinition *uint64 `json:"SimpleAesDefinition,omitnil,omitempty" name:"SimpleAesDefinition"`

	// 保护类型为 Widevine 的转自适应码流模板 ID。
	WidevineDefinition *uint64 `json:"WidevineDefinition,omitnil,omitempty" name:"WidevineDefinition"`

	// 保护类型为 FairPlay 的转自适应码流模板 ID。
	FairPlayDefinition *uint64 `json:"FairPlayDefinition,omitnil,omitempty" name:"FairPlayDefinition"`
}

type DynamicRangeInfo struct {
	// 画面动态范围信息。可取值：
	// <li>SDR：Standard Dynamic Range 标准动态范围；</li>
	// <li>HDR：High Dynamic Range 高动态范围。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 高动态范围类型，当 Type 为 HDR 时有效。目前支持的可取值：
	// <li>hdr10：表示 hdr10 标准；</li>
	// <li>hlg：表示 hlg 标准。</li>
	HDRType *string `json:"HDRType,omitnil,omitempty" name:"HDRType"`
}

type EditMediaFileInfo struct {
	// 视频的 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 视频剪辑起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 视频剪辑结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type EditMediaOutputConfig struct {
	// 输出文件名，最长 64 个字符。缺省由系统指定生成文件名。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 输出文件格式，可选值：mp4、hls。默认是 mp4。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 输出的视频信息。
	VideoStream *EditMediaVideoStream `json:"VideoStream,omitnil,omitempty" name:"VideoStream"`

	// 极速高清转码参数。
	TEHDConfig *EditMediaTEHDConfig `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`
}

// Predefined struct for user
type EditMediaRequestParams struct {
	// 输入视频的类型，可以取的值为  File，Stream 两种。
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 输入的视频文件信息，当 InputType 为 File 时必填。
	FileInfos []*EditMediaFileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 输入的流信息，当 InputType 为 Stream 时必填。
	StreamInfos []*EditMediaStreamInfo `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// 编辑模板 ID，取值有 10，20，不填代表使用 10 模板。
	// <li>10：拼接时，以分辨率最高的输入为基准；</li>
	// <li>20：拼接时，以码率最高的输入为基准。</li>
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// [任务流模板](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF)名字，如果要对生成的新视频执行任务流时填写。
	ProcedureName *string `json:"ProcedureName,omitnil,omitempty" name:"ProcedureName"`

	// 编辑后生成的文件配置。
	OutputConfig *EditMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 标识来源上下文，用于透传用户请求信息，在EditMediaComplete回调和任务流状态变更回调将返回该字段值，最长 1000个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type EditMediaRequest struct {
	*tchttp.BaseRequest
	
	// 输入视频的类型，可以取的值为  File，Stream 两种。
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 输入的视频文件信息，当 InputType 为 File 时必填。
	FileInfos []*EditMediaFileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 输入的流信息，当 InputType 为 Stream 时必填。
	StreamInfos []*EditMediaStreamInfo `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// 编辑模板 ID，取值有 10，20，不填代表使用 10 模板。
	// <li>10：拼接时，以分辨率最高的输入为基准；</li>
	// <li>20：拼接时，以码率最高的输入为基准。</li>
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// [任务流模板](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF)名字，如果要对生成的新视频执行任务流时填写。
	ProcedureName *string `json:"ProcedureName,omitnil,omitempty" name:"ProcedureName"`

	// 编辑后生成的文件配置。
	OutputConfig *EditMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 标识来源上下文，用于透传用户请求信息，在EditMediaComplete回调和任务流状态变更回调将返回该字段值，最长 1000个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
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
	delete(f, "InputType")
	delete(f, "SubAppId")
	delete(f, "FileInfos")
	delete(f, "StreamInfos")
	delete(f, "Definition")
	delete(f, "ProcedureName")
	delete(f, "OutputConfig")
	delete(f, "SessionContext")
	delete(f, "TasksPriority")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditMediaResponseParams struct {
	// 编辑视频的任务 ID，可以通过该 ID 查询编辑任务（任务类型为 EditMedia）的状态。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type EditMediaStreamInfo struct {
	// 录制的流 ID
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// 流剪辑的起始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 流剪辑的结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type EditMediaTEHDConfig struct {
	// 极速高清类型，可选值：<li>TEHD-100 表示极速高清-100;</li> <li>OFF 表示关闭极速高清。</li>不填表示 OFF。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 视频码率上限，当 Type 指定了极速高清类型时有效。
	// 不填或填0表示不设视频码率上限。
	MaxVideoBitrate *uint64 `json:"MaxVideoBitrate,omitnil,omitempty" name:"MaxVideoBitrate"`
}

type EditMediaTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，0 表示成功，其他值表示失败：
	// <li>40000：输入参数不合法，请检查输入参数；</li>
	// <li>60000：源文件错误（如视频数据损坏），请确认源文件是否正常；</li>
	// <li>70000：内部服务错误，建议重试。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 编辑视频任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 视频编辑任务的输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *EditMediaTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 视频编辑任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *EditMediaTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 输出视频的元信息。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 任务类型为 Procedure 的任务 ID。若发起[编辑视频](https://cloud.tencent.com/document/api/266/34783)任务时指定了任务流模板(ProcedureName)，当该任务流模板指定了 MediaProcessTask、AiAnalysisTask、AiRecognitionTask 中的一个或多个时发起该任务。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitnil,omitempty" name:"ProcedureTaskId"`

	// 任务类型为 ReviewAudioVideo 的任务 ID。若发起[编辑视频](https://cloud.tencent.com/document/api/266/34783)任务时指定了任务流模板(ProcedureName)，当该任务流模板指定了 ReviewAudioVideoTask 时，发起该任务。
	ReviewAudioVideoTaskId *string `json:"ReviewAudioVideoTaskId,omitnil,omitempty" name:"ReviewAudioVideoTaskId"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type EditMediaTaskInput struct {
	// 输入视频的来源类型，可以取的值为 File，Stream 两种。
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 输入的视频文件信息，当 InputType 为 File 时，该字段有值。
	FileInfoSet []*EditMediaFileInfo `json:"FileInfoSet,omitnil,omitempty" name:"FileInfoSet"`

	// 输入的流信息，当 InputType 为 Stream 时，该字段有值。
	StreamInfoSet []*EditMediaStreamInfo `json:"StreamInfoSet,omitnil,omitempty" name:"StreamInfoSet"`
}

type EditMediaTaskOutput struct {
	// 文件类型，例如 mp4、flv 等。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 媒体文件播放地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 输出文件名，最长 64 个字符。缺省由系统指定生成文件名。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type EditMediaVideoStream struct {
	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码；</li>
	// <li>libx265：H.265 编码；</li>
	// <li>av1：AOMedia Video 1 编码；</li>
	// <li>H.266：H.266 编码。</li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0 或不填时，表示自动选择最佳视频码率。
	Bitrate *uint64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率取基准分辨率；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按基准分辨率比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按基准分辨率比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率取基准分辨率；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按基准分辨率比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按基准分辨率比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 视频帧率，取值范围：[0, 100]，单位：Hz。
	// 当取值为0，将自动为视频设置帧率。
	// 默认值为 0。
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`
}

type EmptyTrackItem struct {
	// 持续时间，单位为秒。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

// Predefined struct for user
type EnhanceMediaByTemplateRequestParams struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 音画质重生模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 起始偏移时间，单位：秒，不填表示从视频开始截取。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 结束偏移时间，单位：秒，不填表示截取到视频末尾。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 音画质重生后的文件配置。
	OutputConfig *RebuildMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type EnhanceMediaByTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 音画质重生模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 起始偏移时间，单位：秒，不填表示从视频开始截取。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 结束偏移时间，单位：秒，不填表示截取到视频末尾。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 音画质重生后的文件配置。
	OutputConfig *RebuildMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *EnhanceMediaByTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnhanceMediaByTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "Definition")
	delete(f, "SubAppId")
	delete(f, "StartTimeOffset")
	delete(f, "EndTimeOffset")
	delete(f, "OutputConfig")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "TasksPriority")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnhanceMediaByTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnhanceMediaByTemplateResponseParams struct {
	// 音画质重生的任务 ID，可以通过该 ID 查询音画质重生任务的状态。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnhanceMediaByTemplateResponse struct {
	*tchttp.BaseResponse
	Response *EnhanceMediaByTemplateResponseParams `json:"Response"`
}

func (r *EnhanceMediaByTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnhanceMediaByTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhanceMediaQualityOutputConfig struct {
	// 输出文件名，最长 64 个字符。缺省由系统指定生成文件名。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type EnhanceMediaQualityRequestParams struct {
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符，在上传成功后由云点播后台分配。可以在 [视频上传完成事件通知](/document/product/266/7830) 或 [云点播控制台](https://console.cloud.tencent.com/vod/media) 获取该字段。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 音画质重生模板 ID。
	// 针对典型的使用场景，云点播提供了多个[预置模板](https://cloud.tencent.com/document/product/266/102586#50604b3f-0286-4a10-a3f7-18218116aff7)。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生后的媒体文件配置。
	OutputConfig *EnhanceMediaQualityOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，音画质重生完成回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`
}

type EnhanceMediaQualityRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符，在上传成功后由云点播后台分配。可以在 [视频上传完成事件通知](/document/product/266/7830) 或 [云点播控制台](https://console.cloud.tencent.com/vod/media) 获取该字段。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 音画质重生模板 ID。
	// 针对典型的使用场景，云点播提供了多个[预置模板](https://cloud.tencent.com/document/product/266/102586#50604b3f-0286-4a10-a3f7-18218116aff7)。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生后的媒体文件配置。
	OutputConfig *EnhanceMediaQualityOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，音画质重生完成回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`
}

func (r *EnhanceMediaQualityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnhanceMediaQualityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "Definition")
	delete(f, "SubAppId")
	delete(f, "OutputConfig")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "TasksPriority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnhanceMediaQualityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnhanceMediaQualityResponseParams struct {
	// 音画质重生任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnhanceMediaQualityResponse struct {
	*tchttp.BaseResponse
	Response *EnhanceMediaQualityResponseParams `json:"Response"`
}

func (r *EnhanceMediaQualityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnhanceMediaQualityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventContent struct {
	// 事件句柄，调用方必须调用 ConfirmEvents 来确认消息已经收到，确认有效时间 30 秒。失效后，事件可重新被获取。
	EventHandle *string `json:"EventHandle,omitnil,omitempty" name:"EventHandle"`

	// <b>支持事件类型：</b>
	// <li>NewFileUpload：视频上传完成；</li>
	// <li>ProcedureStateChanged：任务流状态变更；</li>
	// <li>FileDeleted：视频删除完成；</li>
	// <li>RestoreMediaComplete：视频取回完成；</li>
	// <li>PullComplete：视频转拉完成；</li>
	// <li>EditMediaComplete：视频编辑完成；</li>
	// <li>SplitMediaComplete：视频拆分完成；</li>
	// <li>ComposeMediaComplete：制作媒体文件完成；</li>
	// <li>WechatMiniProgramPublishComplete：微信小程序发布完成。</li>
	// <li>RemoveWatermark：智能去除水印完成。</li>
	// <li>RebuildMediaComplete：音画质重生完成事件（不推荐使用）。</li>
	// <li>ReviewAudioVideoComplete：音视频审核完成；</li>
	// <li>ExtractTraceWatermarkComplete：提取溯源水印完成；</li>
	// <li>ExtractCopyRightWatermarkComplete：提取版权水印完成；</li>
	// <li>DescribeFileAttributesComplete：获取文件属性完成；</li>
	// <li>QualityInspectComplete：音画质检测完成；</li>
	// <li>QualityEnhanceComplete：音画质重生任务完成；</li>
	// <li>PersistenceComplete：剪辑固化完成；</li>
	// <li>ComplexAdaptiveDynamicStreamingComplete：复杂自适应码流任务完成。</li>
	// <b>兼容 2017 版的事件类型：</b>
	// <li>TranscodeComplete：视频转码完成；</li>
	// <li>ConcatComplete：视频拼接完成；</li>
	// <li>ClipComplete：视频剪辑完成；</li>
	// <li>CreateImageSpriteComplete：视频截取雪碧图完成；</li>
	// <li>CreateSnapshotByTimeOffsetComplete：视频按时间点截图完成。</li>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 视频上传完成事件，当事件类型为 NewFileUpload 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUploadEvent *FileUploadTask `json:"FileUploadEvent,omitnil,omitempty" name:"FileUploadEvent"`

	// 任务流状态变更事件，当事件类型为 ProcedureStateChanged 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcedureStateChangeEvent *ProcedureTask `json:"ProcedureStateChangeEvent,omitnil,omitempty" name:"ProcedureStateChangeEvent"`

	// 文件删除事件，当事件类型为 FileDeleted 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileDeleteEvent *FileDeleteTask `json:"FileDeleteEvent,omitnil,omitempty" name:"FileDeleteEvent"`

	// 视频转拉完成事件，当事件类型为 PullComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullCompleteEvent *PullUploadTask `json:"PullCompleteEvent,omitnil,omitempty" name:"PullCompleteEvent"`

	// 视频编辑完成事件，当事件类型为 EditMediaComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditMediaCompleteEvent *EditMediaTask `json:"EditMediaCompleteEvent,omitnil,omitempty" name:"EditMediaCompleteEvent"`

	// 视频拆分完成事件，当事件类型为 SplitMediaComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SplitMediaCompleteEvent *SplitMediaTask `json:"SplitMediaCompleteEvent,omitnil,omitempty" name:"SplitMediaCompleteEvent"`

	// 制作媒体文件任务完成事件，当事件类型为 ComposeMediaComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComposeMediaCompleteEvent *ComposeMediaTask `json:"ComposeMediaCompleteEvent,omitnil,omitempty" name:"ComposeMediaCompleteEvent"`

	// 视频剪辑完成事件，当事件类型为 ClipComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClipCompleteEvent *ClipTask2017 `json:"ClipCompleteEvent,omitnil,omitempty" name:"ClipCompleteEvent"`

	// 视频转码完成事件，当事件类型为 TranscodeComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeCompleteEvent *TranscodeTask2017 `json:"TranscodeCompleteEvent,omitnil,omitempty" name:"TranscodeCompleteEvent"`

	// 视频截取雪碧图完成事件，当事件类型为 CreateImageSpriteComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateImageSpriteCompleteEvent *CreateImageSpriteTask2017 `json:"CreateImageSpriteCompleteEvent,omitnil,omitempty" name:"CreateImageSpriteCompleteEvent"`

	// 视频拼接完成事件，当事件类型为 ConcatComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConcatCompleteEvent *ConcatTask2017 `json:"ConcatCompleteEvent,omitnil,omitempty" name:"ConcatCompleteEvent"`

	// 视频按时间点截图完成事件，当事件类型为 CreateSnapshotByTimeOffsetComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetCompleteEvent *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetCompleteEvent,omitnil,omitempty" name:"SnapshotByTimeOffsetCompleteEvent"`

	// 微信发布完成事件，当事件类型为 WechatPublishComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatPublishCompleteEvent *WechatPublishTask `json:"WechatPublishCompleteEvent,omitnil,omitempty" name:"WechatPublishCompleteEvent"`

	// 微信小程序发布任务完成事件，当事件类型为 WechatMiniProgramPublishComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatMiniProgramPublishCompleteEvent *WechatMiniProgramPublishTask `json:"WechatMiniProgramPublishCompleteEvent,omitnil,omitempty" name:"WechatMiniProgramPublishCompleteEvent"`

	// 智能去除水印完成事件，当事件类型为 RemoveWatermark 有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoveWatermarkCompleteEvent *RemoveWatermarkTask `json:"RemoveWatermarkCompleteEvent,omitnil,omitempty" name:"RemoveWatermarkCompleteEvent"`

	// 视频取回完成事件，当事件类型为 RestoreMediaComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestoreMediaCompleteEvent *RestoreMediaTask `json:"RestoreMediaCompleteEvent,omitnil,omitempty" name:"RestoreMediaCompleteEvent"`

	// 音画质重生完成事件，当事件类型为 RebuildMediaComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RebuildMediaCompleteEvent *RebuildMediaTask `json:"RebuildMediaCompleteEvent,omitnil,omitempty" name:"RebuildMediaCompleteEvent"`

	// 溯源水印提取完成事件，当事件类型为 ExtractTraceWatermarkComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtractTraceWatermarkCompleteEvent *ExtractTraceWatermarkTask `json:"ExtractTraceWatermarkCompleteEvent,omitnil,omitempty" name:"ExtractTraceWatermarkCompleteEvent"`

	// 版权水印提取完成事件，当事件类型为 ExtractCopyRightWatermarkComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtractCopyRightWatermarkCompleteEvent *ExtractCopyRightWatermarkTask `json:"ExtractCopyRightWatermarkCompleteEvent,omitnil,omitempty" name:"ExtractCopyRightWatermarkCompleteEvent"`

	// 音视频审核完成事件，当事件类型为 ReviewAudioVideoComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewAudioVideoCompleteEvent *ReviewAudioVideoTask `json:"ReviewAudioVideoCompleteEvent,omitnil,omitempty" name:"ReviewAudioVideoCompleteEvent"`

	// 该字段已无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReduceMediaBitrateCompleteEvent *ReduceMediaBitrateTask `json:"ReduceMediaBitrateCompleteEvent,omitnil,omitempty" name:"ReduceMediaBitrateCompleteEvent"`

	// 获取文件属性完成事件，当事件类型为 DescribeFileAttributesComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescribeFileAttributesCompleteEvent *DescribeFileAttributesTask `json:"DescribeFileAttributesCompleteEvent,omitnil,omitempty" name:"DescribeFileAttributesCompleteEvent"`

	// 音画质检测完成事件，当事件类型为 QualityInspectComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityInspectCompleteEvent *QualityInspectTask `json:"QualityInspectCompleteEvent,omitnil,omitempty" name:"QualityInspectCompleteEvent"`

	// 音画质重生完成事件，当事件类型为 QualityEnhanceComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityEnhanceCompleteEvent *QualityEnhanceTask `json:"QualityEnhanceCompleteEvent,omitnil,omitempty" name:"QualityEnhanceCompleteEvent"`

	// 媒体转推状态变化事件，当事件类型为 MediaCastStatusChanged 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaCastStatusChangedEvent *MediaCastEvent `json:"MediaCastStatusChangedEvent,omitnil,omitempty" name:"MediaCastStatusChangedEvent"`

	// 剪辑固化完成事件，当事件类型为 PersistenceComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersistenceCompleteEvent *PersistenceCompleteTask `json:"PersistenceCompleteEvent,omitnil,omitempty" name:"PersistenceCompleteEvent"`

	// 自适应码流任务信息，仅当 EventType 为ComplexAdaptiveDynamicStreamingComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplexAdaptiveDynamicStreamingCompleteEvent *ComplexAdaptiveDynamicStreamingTask `json:"ComplexAdaptiveDynamicStreamingCompleteEvent,omitnil,omitempty" name:"ComplexAdaptiveDynamicStreamingCompleteEvent"`
}

// Predefined struct for user
type ExecuteFunctionRequestParams struct {
	// 调用后端接口名称。
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// 接口参数，具体参数格式调用时与后端协调。
	FunctionArg *string `json:"FunctionArg,omitnil,omitempty" name:"FunctionArg"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type ExecuteFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 调用后端接口名称。
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// 接口参数，具体参数格式调用时与后端协调。
	FunctionArg *string `json:"FunctionArg,omitnil,omitempty" name:"FunctionArg"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
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
	delete(f, "SubAppId")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteFunctionResponseParams struct {
	// 处理结果打包后的字符串，具体与后台一同协调。
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ExtractCopyRightWatermarkRequestParams struct {
	// 需要提取水印的媒体 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 标识来源上下文，用于透传用户请求信息，在 ExtractCopyRightWatermarkComplete 回调和任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type ExtractCopyRightWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 需要提取水印的媒体 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 标识来源上下文，用于透传用户请求信息，在 ExtractCopyRightWatermarkComplete 回调和任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *ExtractCopyRightWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExtractCopyRightWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "SubAppId")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "TasksPriority")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExtractCopyRightWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExtractCopyRightWatermarkResponseParams struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExtractCopyRightWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *ExtractCopyRightWatermarkResponseParams `json:"Response"`
}

func (r *ExtractCopyRightWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExtractCopyRightWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtractCopyRightWatermarkTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，0 表示成功，其他值表示失败：
	// <li>40000：输入参数不合法，请检查输入参数；</li>
	// <li>60000：源文件错误（如视频数据损坏），请确认源文件是否正常；</li>
	// <li>70000：内部服务错误，建议重试。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 提取版权水印任务输入信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *ExtractCopyRightWatermarkTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 提取版权水印任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *ExtractCopyRightWatermarkTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type ExtractCopyRightWatermarkTaskInput struct {
	// 需要提取水印的媒体 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ExtractCopyRightWatermarkTaskOutput struct {
	// 版权信息。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

// Predefined struct for user
type ExtractTraceWatermarkRequestParams struct {
	// 需要提取水印的媒体 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 媒体文件 ID。Url 对应的原始媒体文件 ID。
	// <li><font color=red>注意</font>：此字段必填。</li>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 标识来源上下文，用于透传用户请求信息，在ExtractTraceWatermarkComplete回调和任务流状态变更回调将返回该字段值，最长 1000个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type ExtractTraceWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 需要提取水印的媒体 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 媒体文件 ID。Url 对应的原始媒体文件 ID。
	// <li><font color=red>注意</font>：此字段必填。</li>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 标识来源上下文，用于透传用户请求信息，在ExtractTraceWatermarkComplete回调和任务流状态变更回调将返回该字段值，最长 1000个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *ExtractTraceWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExtractTraceWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "TasksPriority")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExtractTraceWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExtractTraceWatermarkResponseParams struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExtractTraceWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *ExtractTraceWatermarkResponseParams `json:"Response"`
}

func (r *ExtractTraceWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExtractTraceWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtractTraceWatermarkTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，0 表示成功，其他值表示失败：
	// <li>40000：输入参数不合法，请检查输入参数；</li>
	// <li>60000：源文件错误（如视频数据损坏），请确认源文件是否正常；</li>
	// <li>70000：内部服务错误，建议重试。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 提取溯源水印任务输入信息。
	Input *ExtractTraceWatermarkTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 提取溯源水印任务输出信息。
	Output *ExtractTraceWatermarkTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type ExtractTraceWatermarkTaskInput struct {
	// 需要提取水印的媒体 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 媒体文件 ID。Url 对应的原始媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

type ExtractTraceWatermarkTaskOutput struct {
	// 播放者的 ID，以十六进制表示，共6位。
	Uv *string `json:"Uv,omitnil,omitempty" name:"Uv"`

	// 该字段已废弃。
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`
}

type FaceConfigureInfo struct {
	// 人脸识别任务开关，可选值：
	// <li>ON：开启智能人脸识别任务；</li>
	// <li>OFF：关闭智能人脸识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 人脸识别过滤分数，当识别结果达到该分数以上，返回识别结果。默认 95 分。取值范围：0 - 100。
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 默认人物过滤标签，指定需要返回的默认人物的标签。如果未填或者为空，则全部默认人物结果都返回。标签可选值：
	// <li>entertainment：娱乐明星；</li>
	// <li>sport：体育明星；</li>
	// <li>politician：相关人物。</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitnil,omitempty" name:"DefaultLibraryLabelSet"`

	// 用户自定义人物过滤标签，指定需要返回的用户自定义人物的标签。如果未填或者为空，则全部自定义人物结果都返回。
	// 标签个数最多 100 个，每个标签长度最多 16 个字符。
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitnil,omitempty" name:"UserDefineLibraryLabelSet"`

	// 人物库选择，可选值：
	// <li>Default：使用默认人物库；</li>
	// <li>UserDefine：使用用户自定义人物库。</li>
	// <li>All：同时使用默认人物库和用户自定义人物库。</li>
	// 默认值：All，使用系统默认人物库及用户自定义人物库。
	FaceLibrary *string `json:"FaceLibrary,omitnil,omitempty" name:"FaceLibrary"`
}

type FaceConfigureInfoForUpdate struct {
	// 人脸识别任务开关，可选值：
	// <li>ON：开启智能人脸识别任务；</li>
	// <li>OFF：关闭智能人脸识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 人脸识别过滤分数，当识别结果达到该分数以上，返回识别结果。取值范围：0-100。
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 默认人物过滤标签，指定需要返回的默认人物的标签。如果未填或者为空，则全部默认人物结果都返回。标签可选值：
	// <li>entertainment：娱乐明星；</li>
	// <li>sport：体育明星；</li>
	// <li>politician：相关人物。</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitnil,omitempty" name:"DefaultLibraryLabelSet"`

	// 用户自定义人物过滤标签，指定需要返回的用户自定义人物的标签。如果未填或者为空，则全部自定义人物结果都返回。
	// 标签个数最多 100 个，每个标签长度最多 16 个字符。
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitnil,omitempty" name:"UserDefineLibraryLabelSet"`

	// 人物库选择，可选值：
	// <li>Default：使用默认人物库；</li>
	// <li>UserDefine：使用用户自定义人物库。</li>
	// <li>All：同时使用默认人物库和用户自定义人物库。</li>
	FaceLibrary *string `json:"FaceLibrary,omitnil,omitempty" name:"FaceLibrary"`
}

type FaceEnhanceInfo struct {
	// 人脸增强控制开关，可选值：
	// <li>ON：开启人脸增强；</li>
	// <li>OFF：关闭人脸增强。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 人脸增强强度，仅当人脸增强控制开关为 ON 时有效，取值范围：0.0~1.0。
	// 默认：0.0。
	Intensity *float64 `json:"Intensity,omitnil,omitempty" name:"Intensity"`
}

type FastEditMediaFileInfo struct {
	// 媒体的 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 操作的音视频类型，可选值有：
	// <li>Transcode：转码输出；</li>
	// <li>Original：原始音视频。</li>
	// 注意：操作的音视频，必须为 HLS 格式。
	AudioVideoType *string `json:"AudioVideoType,omitnil,omitempty" name:"AudioVideoType"`

	// 当 AudioVideoType 为 Transcode 时有效，表示操作媒体的的转码模板 ID。
	TranscodeDefinition *int64 `json:"TranscodeDefinition,omitnil,omitempty" name:"TranscodeDefinition"`

	// 媒体剪辑起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 媒体剪辑结束的时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

// Predefined struct for user
type FastEditMediaRequestParams struct {
	// 输入的媒体文件信息。最多支持传入100个媒体。
	FileInfos []*FastEditMediaFileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// ClipMode 用来表示剪辑时间点落在一个 TS 分片中间时，是否包含这个分片。共有两种取值： <li>StartInclusiveEndInclusive：当剪辑起始时间点和结束时间点落在一个分片的中间时，都会包含这个分片；</li> <li>StartInclusiveEndExclusive：当起始时间点落在一个分片的中间时，会包含这个分片；而当结束时间点落在一个分片的中间时，不会包含这个分片。</li> 不填时，默认为 StartInclusiveEndInclusive。
	ClipMode *string `json:"ClipMode,omitnil,omitempty" name:"ClipMode"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type FastEditMediaRequest struct {
	*tchttp.BaseRequest
	
	// 输入的媒体文件信息。最多支持传入100个媒体。
	FileInfos []*FastEditMediaFileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// ClipMode 用来表示剪辑时间点落在一个 TS 分片中间时，是否包含这个分片。共有两种取值： <li>StartInclusiveEndInclusive：当剪辑起始时间点和结束时间点落在一个分片的中间时，都会包含这个分片；</li> <li>StartInclusiveEndExclusive：当起始时间点落在一个分片的中间时，会包含这个分片；而当结束时间点落在一个分片的中间时，不会包含这个分片。</li> 不填时，默认为 StartInclusiveEndInclusive。
	ClipMode *string `json:"ClipMode,omitnil,omitempty" name:"ClipMode"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *FastEditMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FastEditMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileInfos")
	delete(f, "ClipMode")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FastEditMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FastEditMediaResponseParams struct {
	// 快速编辑后的视频的媒体文件的唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 快速编辑后的媒体播放地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FastEditMediaResponse struct {
	*tchttp.BaseResponse
	Response *FastEditMediaResponseParams `json:"Response"`
}

func (r *FastEditMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FastEditMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileDeleteResultItem struct {
	// 删除的文件 ID 。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 本次删除的文件部分。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteParts []*MediaDeleteItem `json:"DeleteParts,omitnil,omitempty" name:"DeleteParts"`
}

type FileDeleteTask struct {
	// 删除文件 ID 列表。
	FileIdSet []*string `json:"FileIdSet,omitnil,omitempty" name:"FileIdSet"`

	// 删除文件结果信息列表。
	FileDeleteResultInfo []*FileDeleteResultItem `json:"FileDeleteResultInfo,omitnil,omitempty" name:"FileDeleteResultInfo"`
}

type FileReviewInfo struct {
	// 媒体审核信息\*。
	// 
	// \* 只展示通过 [音视频审核(ReviewAudioVideo)](https://cloud.tencent.com/document/api/266/80283) 或 [图片审核(ReviewImage)](https://cloud.tencent.com/document/api/266/73217) 发起的审核结果信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaReviewInfo *ReviewInfo `json:"MediaReviewInfo,omitnil,omitempty" name:"MediaReviewInfo"`

	// 媒体封面审核信息\*。
	// 
	// \* 只展示通过 [音视频审核(ReviewAudioVideo)](https://cloud.tencent.com/document/api/266/80283) 或 [图片审核(ReviewImage)](https://cloud.tencent.com/document/api/266/73217) 发起的审核结果信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverReviewInfo *ReviewInfo `json:"CoverReviewInfo,omitnil,omitempty" name:"CoverReviewInfo"`
}

type FileUploadTask struct {
	// 文件唯一 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 上传完成后生成的媒体文件基础信息。
	MediaBasicInfo *MediaBasicInfo `json:"MediaBasicInfo,omitnil,omitempty" name:"MediaBasicInfo"`

	// 任务类型为 Procedure 的任务 ID。若视频[上传时指定要执行的任务(procedure)](https://cloud.tencent.com/document/product/266/33475#.E4.BB.BB.E5.8A.A1.E5.8F.91.E8.B5.B7)，当该任务流模板指定了 MediaProcessTask、AiAnalysisTask、AiRecognitionTask 中的一个或多个时发起该任务。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitnil,omitempty" name:"ProcedureTaskId"`

	// 任务类型为 ReviewAudioVideo 的任务 ID。若视频[上传时指定要执行的任务(procedure)](https://cloud.tencent.com/document/product/266/33475#.E4.BB.BB.E5.8A.A1.E5.8F.91.E8.B5.B7)，当该任务流模板指定了 ReviewAudioVideoTask 时，发起该任务。
	ReviewAudioVideoTaskId *string `json:"ReviewAudioVideoTaskId,omitnil,omitempty" name:"ReviewAudioVideoTaskId"`

	// 元信息。包括大小、时长、视频流信息、音频流信息等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`
}

type FileVerifyInfo struct {
	// 文件验证 URL 指引。
	FileVerifyUrl *string `json:"FileVerifyUrl,omitnil,omitempty" name:"FileVerifyUrl"`

	// 文件校验域名列表。
	FileVerifyDomains []*string `json:"FileVerifyDomains,omitnil,omitempty" name:"FileVerifyDomains"`

	// 文件校验文件名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileVerifyName *string `json:"FileVerifyName,omitnil,omitempty" name:"FileVerifyName"`
}

// Predefined struct for user
type ForbidMediaDistributionRequestParams struct {
	// 媒体文件列表，每次最多可提交 20 条。
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// forbid：禁播，recover：解禁。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type ForbidMediaDistributionRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件列表，每次最多可提交 20 条。
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// forbid：禁播，recover：解禁。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *ForbidMediaDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForbidMediaDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileIds")
	delete(f, "Operation")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForbidMediaDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ForbidMediaDistributionResponseParams struct {
	// 不存在的文件 ID 列表。
	NotExistFileIdSet []*string `json:"NotExistFileIdSet,omitnil,omitempty" name:"NotExistFileIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ForbidMediaDistributionResponse struct {
	*tchttp.BaseResponse
	Response *ForbidMediaDistributionResponseParams `json:"Response"`
}

func (r *ForbidMediaDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForbidMediaDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FrameTagConfigureInfo struct {
	// 智能按帧标签任务开关，可选值：
	// <li>ON：开启智能按帧标签任务；</li>
	// <li>OFF：关闭智能按帧标签任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 截帧间隔，单位为秒，当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`
}

type FrameTagConfigureInfoForUpdate struct {
	// 智能按帧标签任务开关，可选值：
	// <li>ON：开启智能按帧标签任务；</li>
	// <li>OFF：关闭智能按帧标签任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 截帧间隔，单位为秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`
}

type HDRInfo struct {
	// 高动态范围类型控制开关，可选值：
	// <li>ON：开启高动态范围类型转换；</li>
	// <li>OFF：关闭高动态范围类型转换。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 高动态范围类型，可选值：
	// <li>hdr10：表示 hdr10 标准；</li>
	// <li>hlg：表示 hlg 标准。</li>
	// 
	// 注意：
	// <li> 仅当高动态范围类型控制开关为 ON 时有效；</li>
	// <li>当画质重生目标参数中指定视频输出参数的视频流编码格式 Codec 为 libx265 时有效。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type HandleCurrentPlaylistRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 轮播播单唯一标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// 操作类型，取值有：<li>Insert：向当前播放列表插入播放节目。</li> <li>InsertTemporary：向当前播放列表临时插入播放节目。只能插入到当前正在播放的节目后面，临时插入的节目只在本次轮播过程生效。</li><li>Delete：删除播放列表中的播放节目。不能删除正在播放的节目。</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 播单节目 ID。当 Operation 为 Insert 时必填，表示插入的节目列表位于该播放节目之后。插入的位置必须在当前正在播放的节目之后。
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// 节目列表。当 Operation 为 Insert、InsertTemporary、Delete 时必填，表示要操作的节目列表。列表长度最大为10。
	RoundPlaylist []*RoundPlayListItemInfo `json:"RoundPlaylist,omitnil,omitempty" name:"RoundPlaylist"`
}

type HandleCurrentPlaylistRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 轮播播单唯一标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// 操作类型，取值有：<li>Insert：向当前播放列表插入播放节目。</li> <li>InsertTemporary：向当前播放列表临时插入播放节目。只能插入到当前正在播放的节目后面，临时插入的节目只在本次轮播过程生效。</li><li>Delete：删除播放列表中的播放节目。不能删除正在播放的节目。</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 播单节目 ID。当 Operation 为 Insert 时必填，表示插入的节目列表位于该播放节目之后。插入的位置必须在当前正在播放的节目之后。
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// 节目列表。当 Operation 为 Insert、InsertTemporary、Delete 时必填，表示要操作的节目列表。列表长度最大为10。
	RoundPlaylist []*RoundPlayListItemInfo `json:"RoundPlaylist,omitnil,omitempty" name:"RoundPlaylist"`
}

func (r *HandleCurrentPlaylistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HandleCurrentPlaylistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "RoundPlayId")
	delete(f, "Operation")
	delete(f, "ItemId")
	delete(f, "RoundPlaylist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HandleCurrentPlaylistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HandleCurrentPlaylistResponseParams struct {
	// 操作成功的节目列表。
	RoundPlaylist []*RoundPlayListItemInfo `json:"RoundPlaylist,omitnil,omitempty" name:"RoundPlaylist"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type HandleCurrentPlaylistResponse struct {
	*tchttp.BaseResponse
	Response *HandleCurrentPlaylistResponseParams `json:"Response"`
}

func (r *HandleCurrentPlaylistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HandleCurrentPlaylistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HeadTailConfigureInfo struct {
	// 视频片头片尾识别任务开关，可选值：
	// <li>ON：开启智能视频片头片尾识别任务；</li>
	// <li>OFF：关闭智能视频片头片尾识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type HeadTailConfigureInfoForUpdate struct {
	// 视频片头片尾识别任务开关，可选值：
	// <li>ON：开启智能视频片头片尾识别任务；</li>
	// <li>OFF：关闭智能视频片头片尾识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type HeadTailTaskInput struct {
	// 片头片尾模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type HeadTailTemplate struct {
	// 片头片尾模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板名，最大支持 64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述，最大支持 256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 片头候选列表。使用时会选择跟正片分辨率最贴近的一个使用，当存在相同的候选时，选择第一个使用，最大支持 5 个。
	HeadCandidateSet []*string `json:"HeadCandidateSet,omitnil,omitempty" name:"HeadCandidateSet"`

	// 片尾候选列表。使用时会选择跟正片分辨率最贴近的一个使用，当存在相同的候选时，选择第一个使用，最大支持 5 个。
	TailCandidateSet []*string `json:"TailCandidateSet,omitnil,omitempty" name:"TailCandidateSet"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li> gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊；</li>
	// <li> white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充；</li>
	// <li> black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：stretch 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type HighlightSegmentItem struct {
	// 置信度。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 片段起始时间偏移。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 片段结束时间偏移。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type HighlightsConfigureInfo struct {
	// 智能精彩片段任务开关，可选值：
	// <li>ON：开启智能精彩片段任务；</li>
	// <li>OFF：关闭智能精彩片段任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type HighlightsConfigureInfoForUpdate struct {
	// 智能精彩片段任务开关，可选值：
	// <li>ON：开启智能精彩片段任务；</li>
	// <li>OFF：关闭智能精彩片段任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ImageBlur struct {
	// 图片模糊的操作类型。可选模式有：
	// <li>Gaussian : 高斯模糊。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 模糊半径，取值范围为1 - 50。当 Type 取值为 Gaussian 时此字段有效。
	Radius *int64 `json:"Radius,omitnil,omitempty" name:"Radius"`

	// 正态分布的标准差，必须大于0。当 Type 取值为 Gaussian 时此字段有效。
	Sigma *int64 `json:"Sigma,omitnil,omitempty" name:"Sigma"`
}

type ImageCenterCut struct {
	// 图片的裁剪模式，可选 Circle 和 Rectangle。
	// <li>Circle ： 内切圆裁剪，输出图片半径为 Radius。</li>
	// <li>Rectangle ： 矩形裁剪，输出图片宽为 Width ， 高为 Height。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 输出图片的宽度，单位为像素，当 Type 取值为 Rectangle 时有效。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 输出图片的高度，单位为像素，当 Type 取值为 Rectangle 时有效。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 输出图片的半径，单位为像素，当 Type 取值为 Circle 时有效。
	Radius *int64 `json:"Radius,omitnil,omitempty" name:"Radius"`
}

type ImageContentReviewInput struct {
	// 图片智能内容审核模板 ID。当前只支持：
	// <li>10：所有审核类型均打开。</li>
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type ImageOperation struct {
	// 图片处理类型。可选类型有：
	// <li>Scale : 图片缩略处理；</li>
	// <li>CenterCut : 图片裁剪处理；</li>
	// <li>Blur : 图片模糊处理。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 图片缩略处理，仅当 Type 为 Scale 时有效。
	Scale *ImageScale `json:"Scale,omitnil,omitempty" name:"Scale"`

	// 图片裁剪处理，仅当 Type 为 CenterCut 时有效。
	CenterCut *ImageCenterCut `json:"CenterCut,omitnil,omitempty" name:"CenterCut"`

	// 图片模糊处理，仅当 Type 为 Blur 时有效。
	Blur *ImageBlur `json:"Blur,omitnil,omitempty" name:"Blur"`
}

type ImageProcessingTemplate struct {
	// 图片处理模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 图片处理模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 图片处理操作数组，操作将以数组顺序执行。
	// <li>长度限制：3。</li>
	Operations []*ImageOperation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ImageReviewUsageDataItem struct {
	// 数据所在时间区间的开始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。如：当时间粒度为天，2018-12-01T00:00:00+08:00，表示2018年12月1日（含）到2018年12月2日（不含）区间。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 次数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type ImageScale struct {
	// 图片缩放的操作类型。可选模式有：
	// <li>WidthFirst : 指定图片的宽为 Width ，高度等比缩放。</li>
	// <li>HeightFirst : 指定图片的高为 Height ，宽度等比缩放。</li>
	// <li>LongEdgeFirst : 指定图片的长边为 LongEdge ，短边等比缩放。</li>
	// <li>ShortEdgeFirst : 指定图片的短边为 ShortEdge ，长边等比缩放。</li>
	// <li>Force : 忽略原图宽高比例，指定图片宽度为 Width，高度为 Height ，强行缩放图片，可能导致目标图片变形。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 输出图片的高度，单位为像素。当 Type 取值为 HeightFirst 或 Force 时此字段有效。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 输出图片的宽度，单位为像素。当 Type 取值为 WidthFirst 或 Force 时此字段有效。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 输出图片的长边长度，单位为像素。当 Type 取值为 LongEdgeFirst 时此字段有效。
	LongEdge *uint64 `json:"LongEdge,omitnil,omitempty" name:"LongEdge"`

	// 输出图片的短边长度，单位为像素。当 Type 取值为 ShortEdgeFirst 时此字段有效。
	ShortEdge *uint64 `json:"ShortEdge,omitnil,omitempty" name:"ShortEdge"`
}

type ImageSpriteTaskInput struct {
	// 雪碧图模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type ImageSpriteTemplate struct {
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 雪碧图模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 采样类型。
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 图片格式。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type ImageTransform struct {
	// 类型，取值有：
	// <li> Rotate：图像旋转。</li>
	// <li> Flip：图像翻转。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 图像以中心点为原点进行旋转的角度，取值范围0~360。当 Type = Rotate 时有效。
	RotateAngle *float64 `json:"RotateAngle,omitnil,omitempty" name:"RotateAngle"`

	// 图像翻转动作，取值有：
	// <li>Horizental：水平翻转，即左右镜像。</li>
	// <li>Vertical：垂直翻转，即上下镜像。</li>
	// 当 Type = Flip 时有效。
	Flip *string `json:"Flip,omitnil,omitempty" name:"Flip"`
}

type ImageWatermarkInput struct {
	// 水印图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串。支持 jpeg、png、gif 图片格式。
	ImageContent *string `json:"ImageContent,omitnil,omitempty" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。取值范围为[8, 4096]。</li>
	// 默认值：10%。
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素。取值范围为0或[8, 4096]。</li>
	// 默认值：0px，表示 Height 按照原始水印图片的宽高比缩放。
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印重复类型。使用场景：水印为动态图像。取值范围：
	// <li>once：动态水印播放完后，不再出现；</li>
	// <li>repeat_last_frame：水印播放完后，停留在最后一帧；</li>
	// <li>repeat：水印循环播放，直到视频结束（默认值）。</li>
	RepeatType *string `json:"RepeatType,omitnil,omitempty" name:"RepeatType"`

	// 图片透明度，取值范围：[0, 100]
	// <li>0：完全不透明</li>
	// <li>100：完全透明</li>
	// 默认值：0。
	Transparency *int64 `json:"Transparency,omitnil,omitempty" name:"Transparency"`
}

type ImageWatermarkInputForUpdate struct {
	// 水印图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串。支持 jpeg、png 图片格式。
	ImageContent *string `json:"ImageContent,omitnil,omitempty" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。取值范围为[8, 4096]。</li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素。取值范围为0或[8, 4096]。</li>
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印重复类型。使用场景：水印为动态图像。取值范围：
	// <li>once：动态水印播放完后，不再出现；</li>
	// <li>repeat_last_frame：水印播放完后，停留在最后一帧；</li>
	// <li>repeat：水印循环播放，直到视频结束。</li>
	RepeatType *string `json:"RepeatType,omitnil,omitempty" name:"RepeatType"`

	// 图片透明度，取值范围：[0, 100]
	// <li>0：完全不透明</li>
	// <li>100：完全透明。</li>
	Transparency *int64 `json:"Transparency,omitnil,omitempty" name:"Transparency"`
}

type ImageWatermarkTemplate struct {
	// 水印图片地址。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素；</li>
	// 0px：表示 Height 按照 Width 对视频宽度的比例缩放。
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印重复类型。使用场景：水印为动态图像。取值范围：
	// <li>once：动态水印播放完后，不再出现；</li>
	// <li>repeat_last_frame：水印播放完后，停留在最后一帧；</li>
	// <li>repeat：水印循环播放，直到视频结束。</li>
	RepeatType *string `json:"RepeatType,omitnil,omitempty" name:"RepeatType"`

	// 图片透明度，取值范围：[0, 100]
	// <li>0：完全不透明</li>
	// <li>100：完全透明。</li>
	Transparency *int64 `json:"Transparency,omitnil,omitempty" name:"Transparency"`
}

// Predefined struct for user
type InspectMediaQualityRequestParams struct {
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符，在上传成功后由云点播后台分配。可以在 [视频上传完成事件通知](/document/product/266/7830) 或 [云点播控制台](https://console.cloud.tencent.com/vod/media) 获取该字段。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 音画质检测模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 来源上下文，用于透传用户请求信息，音画质检测完成回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type InspectMediaQualityRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符，在上传成功后由云点播后台分配。可以在 [视频上传完成事件通知](/document/product/266/7830) 或 [云点播控制台](https://console.cloud.tencent.com/vod/media) 获取该字段。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 音画质检测模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 来源上下文，用于透传用户请求信息，音画质检测完成回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *InspectMediaQualityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InspectMediaQualityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "Definition")
	delete(f, "SubAppId")
	delete(f, "TasksPriority")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InspectMediaQualityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InspectMediaQualityResponseParams struct {
	// 音画质检测任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InspectMediaQualityResponse struct {
	*tchttp.BaseResponse
	Response *InspectMediaQualityResponseParams `json:"Response"`
}

func (r *InspectMediaQualityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InspectMediaQualityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JitterConfigureInfo struct {
	// 视频画面抖动重影检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type JitterConfigureInfoForUpdate struct {
	// 视频画面抖动重影检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type JustInTimeTranscodeTemplate struct {
	// 模板类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 模板名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频参数配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoConfigure *VideoConfigureInfo `json:"VideoConfigure,omitnil,omitempty" name:"VideoConfigure"`

	// 水印参数配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkConfigure *WatermarkConfigureData `json:"WatermarkConfigure,omitnil,omitempty" name:"WatermarkConfigure"`
}

type LicenseUsageDataItem struct {
	// 数据所在时间区间的开始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。如：当时间粒度为天，2018-12-01T00:00:00+08:00，表示2018年12月1日（含）到2018年12月2日（不含）区间。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// License 请求次数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type LiveRealTimeClipMediaSegmentInfo struct {
	// 片段的起始时间。格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 片段的结束时间。格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type LiveRealTimeClipRequestParams struct {
	// 推流直播码。
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// 流剪辑的开始时间，格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 流剪辑的结束时间，格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 是否固化。0 不固化，1 固化。默认不固化。
	IsPersistence *int64 `json:"IsPersistence,omitnil,omitempty" name:"IsPersistence"`

	// 剪辑固化后的视频存储过期时间。格式参照 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。仅 IsPersistence 为 1 时有效，默认剪辑固化的视频永不过期。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 剪辑固化后的视频点播任务流处理，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。仅 IsPersistence 为 1 时有效。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	// 仅 IsPersistence 为 1 时有效。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，[上传完成回调](/document/product/266/7830) 将返回该字段值，最长 250 个字符。仅 IsPersistence 为 1 时有效。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`

	// 会话上下文，用于透传用户请求信息，当指定 Procedure 参数后，[任务流状态变更回调](/document/product/266/9636) 将返回该字段值，最长 1000 个字符。仅 IsPersistence 为 1 时有效。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 是否需要返回剪辑后的视频元信息。0 不需要，1 需要。默认不需要。
	MetaDataRequired *uint64 `json:"MetaDataRequired,omitnil,omitempty" name:"MetaDataRequired"`

	// 云点播中添加的用于时移播放的域名，必须在云直播已经[关联录制模板和开通时移服务](https://cloud.tencent.com/document/product/266/52220#.E6.AD.A5.E9.AA.A43.EF.BC.9A.E5.85.B3.E8.81.94.E5.BD.95.E5.88.B6.E6.A8.A1.E6.9D.BF.3Ca-id.3D.22step3.22.3E.3C.2Fa.3E)。**如果本接口的首次调用时间在 2021-01-01T00:00:00Z 之后，则此字段为必选字段。**
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 剪辑的直播流信息：
	// <li>默认剪辑直播原始流。</li>
	// <li>当StreamInfo中指定的Type为Transcoding，则剪辑TemplateId对应的直播转码流。</li>
	StreamInfo *LiveRealTimeClipStreamInfo `json:"StreamInfo,omitnil,omitempty" name:"StreamInfo"`

	// 系统保留字段，请勿填写。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type LiveRealTimeClipRequest struct {
	*tchttp.BaseRequest
	
	// 推流直播码。
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// 流剪辑的开始时间，格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 流剪辑的结束时间，格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 是否固化。0 不固化，1 固化。默认不固化。
	IsPersistence *int64 `json:"IsPersistence,omitnil,omitempty" name:"IsPersistence"`

	// 剪辑固化后的视频存储过期时间。格式参照 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。仅 IsPersistence 为 1 时有效，默认剪辑固化的视频永不过期。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 剪辑固化后的视频点播任务流处理，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。仅 IsPersistence 为 1 时有效。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	// 仅 IsPersistence 为 1 时有效。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，[上传完成回调](/document/product/266/7830) 将返回该字段值，最长 250 个字符。仅 IsPersistence 为 1 时有效。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`

	// 会话上下文，用于透传用户请求信息，当指定 Procedure 参数后，[任务流状态变更回调](/document/product/266/9636) 将返回该字段值，最长 1000 个字符。仅 IsPersistence 为 1 时有效。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 是否需要返回剪辑后的视频元信息。0 不需要，1 需要。默认不需要。
	MetaDataRequired *uint64 `json:"MetaDataRequired,omitnil,omitempty" name:"MetaDataRequired"`

	// 云点播中添加的用于时移播放的域名，必须在云直播已经[关联录制模板和开通时移服务](https://cloud.tencent.com/document/product/266/52220#.E6.AD.A5.E9.AA.A43.EF.BC.9A.E5.85.B3.E8.81.94.E5.BD.95.E5.88.B6.E6.A8.A1.E6.9D.BF.3Ca-id.3D.22step3.22.3E.3C.2Fa.3E)。**如果本接口的首次调用时间在 2021-01-01T00:00:00Z 之后，则此字段为必选字段。**
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 剪辑的直播流信息：
	// <li>默认剪辑直播原始流。</li>
	// <li>当StreamInfo中指定的Type为Transcoding，则剪辑TemplateId对应的直播转码流。</li>
	StreamInfo *LiveRealTimeClipStreamInfo `json:"StreamInfo,omitnil,omitempty" name:"StreamInfo"`

	// 系统保留字段，请勿填写。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *LiveRealTimeClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LiveRealTimeClipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "IsPersistence")
	delete(f, "ExpireTime")
	delete(f, "Procedure")
	delete(f, "ClassId")
	delete(f, "SourceContext")
	delete(f, "SessionContext")
	delete(f, "MetaDataRequired")
	delete(f, "Host")
	delete(f, "StreamInfo")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LiveRealTimeClipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LiveRealTimeClipResponseParams struct {
	// 剪辑后的视频播放 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 剪辑固化后的视频的媒体文件的唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 剪辑固化后的视频任务流 ID。
	VodTaskId *string `json:"VodTaskId,omitnil,omitempty" name:"VodTaskId"`

	// 剪辑后的视频元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// <span id="p_segmentset">剪辑后的视频片段信息。</span>
	SegmentSet []*LiveRealTimeClipMediaSegmentInfo `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LiveRealTimeClipResponse struct {
	*tchttp.BaseResponse
	Response *LiveRealTimeClipResponseParams `json:"Response"`
}

func (r *LiveRealTimeClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LiveRealTimeClipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LiveRealTimeClipStreamInfo struct {
	// 直播流类型，可选值：
	// <li>Original（原始流，<b>默认值</b>）。</li>
	// <li>Transcoding（转码流）。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 直播转码模板ID。
	// <b>当Type值为"Transcoding"时，必须填写。</b>
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type LiveRecordInfo struct {
	// 直播录制流 ID。
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// 录制起始时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	RecordStartTime *string `json:"RecordStartTime,omitnil,omitempty" name:"RecordStartTime"`

	// 录制结束时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	RecordEndTime *string `json:"RecordEndTime,omitnil,omitempty" name:"RecordEndTime"`
}

type LowLightEnhanceInfo struct {
	// 低光照增强控制开关，可选值：
	// <li>ON：开启低光照增强；</li>
	// <li>OFF：关闭低光照增强。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 低光照增强类型，仅当低光照增强控制开关为 ON 时有效，可选值：
	// <li>normal：正常低光照增强；</li>
	// 默认值：normal。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type ManageTaskRequestParams struct {
	// 视频处理的任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作类型，取值范围：
	// <li>Abort：终止任务。只能终止已发起且状态为等待中（WAITING）的任务。</li>
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type ManageTaskRequest struct {
	*tchttp.BaseRequest
	
	// 视频处理的任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作类型，取值范围：
	// <li>Abort：终止任务。只能终止已发起且状态为等待中（WAITING）的任务。</li>
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
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
	delete(f, "TaskId")
	delete(f, "OperationType")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type MediaAdaptiveDynamicStreamingInfo struct {
	// 转自适应码流信息数组。
	AdaptiveDynamicStreamingSet []*AdaptiveDynamicStreamingInfoItem `json:"AdaptiveDynamicStreamingSet,omitnil,omitempty" name:"AdaptiveDynamicStreamingSet"`
}

type MediaAiAnalysisClassificationItem struct {
	// 智能分类的类别名称。
	Classification *string `json:"Classification,omitnil,omitempty" name:"Classification"`

	// 智能分类的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type MediaAiAnalysisCoverItem struct {
	// 智能封面地址。
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// 智能封面的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type MediaAiAnalysisFrameTagItem struct {
	// 按帧标签名称。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 按帧标签名称的分类列表，CategorySet.N 表示第 N+1级分类。
	// 比如 Tag 为“塔楼”时，CategorySet 包含两个元素：CategorySet.0 为“场景”，CategorySet.1为 “建筑”，表示按帧标签为“塔楼”，且第1级分类是“场景”，第2级分类是“建筑”。
	CategorySet []*string `json:"CategorySet,omitnil,omitempty" name:"CategorySet"`

	// 按帧标签的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type MediaAiAnalysisFrameTagSegmentItem struct {
	// 按帧标签起始的偏移时间。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 按帧标签结束的偏移时间。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 时间片段内的标签列表。
	TagSet []*MediaAiAnalysisFrameTagItem `json:"TagSet,omitnil,omitempty" name:"TagSet"`
}

type MediaAiAnalysisHighlightItem struct {
	// 智能精彩集锦地址。
	HighlightUrl *string `json:"HighlightUrl,omitnil,omitempty" name:"HighlightUrl"`

	// 智能精彩集锦封面地址。
	CovImgUrl *string `json:"CovImgUrl,omitnil,omitempty" name:"CovImgUrl"`

	// 智能精彩集锦的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 智能精彩集锦持续时间。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 智能精彩集锦子片段列表，精彩集锦片段由这些子片段拼接生成。
	SegmentSet []*HighlightSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type MediaAiAnalysisTagItem struct {
	// 标签名称。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 标签的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type MediaAnimatedGraphicsInfo struct {
	// 视频转动图结果信息
	AnimatedGraphicsSet []*MediaAnimatedGraphicsItem `json:"AnimatedGraphicsSet,omitnil,omitempty" name:"AnimatedGraphicsSet"`
}

type MediaAnimatedGraphicsItem struct {
	// 转动图的文件地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 转动图模板 ID，参见[转动图参数模板](https://cloud.tencent.com/document/product/266/33481#.3Cspan-id-.3D-.22zdt.22.3E.3C.2Fspan.3E.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 动图格式，如 gif。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 动图的高度，单位：px。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 动图的宽度，单位：px。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 动图码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 动图大小，单位：字节。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 动图的md5值。
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 动图在视频中的起始时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type MediaAudioStreamItem struct {
	// 音频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 音频流的采样率，单位：hz。
	SamplingRate *int64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// 音频流的编码格式，例如 aac。
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`
}

type MediaBasicInfo struct {
	// 媒体文件名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 媒体文件描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 媒体文件的创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 媒体文件的最近更新时间（如修改视频属性、发起视频处理等会触发更新媒体文件信息的操作），使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 媒体文件的过期时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。“9999-12-31T23:59:59Z”表示永不过期。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 媒体文件的分类 ID。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 媒体文件的分类名称。
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// 媒体文件的分类路径，分类间以“-”分隔，如“新的一级分类 - 新的二级分类”。
	ClassPath *string `json:"ClassPath,omitnil,omitempty" name:"ClassPath"`

	// 媒体文件的封面图片地址。
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// 媒体文件的封装格式，例如 mp4、flv 等。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 原始媒体文件的 URL 地址。
	MediaUrl *string `json:"MediaUrl,omitnil,omitempty" name:"MediaUrl"`

	// 该媒体文件的来源信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceInfo *MediaSourceData `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`

	// 媒体文件存储地区，如 ap-chongqing，参见[地域列表](https://cloud.tencent.com/document/product/266/9760#.E5.B7.B2.E6.94.AF.E6.8C.81.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 媒体文件的标签信息。
	TagSet []*string `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 直播录制文件的唯一标识。
	Vid *string `json:"Vid,omitnil,omitempty" name:"Vid"`

	// 文件类型：
	// <li>Video: 视频文件</li>
	// <li>Audio: 音频文件</li>
	// <li>Image: 图片文件</li>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 文件状态：Normal：正常，Forbidden：封禁。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 媒体文件的存储类别：
	// <li> STANDARD：标准存储。</li>
	// <li> STANDARD_IA：低频存储。</li>
	// <li> ARCHIVE：归档存储。</li>
	// <li> DEEP_ARCHIVE：深度归档存储。</li>
	StorageClass *string `json:"StorageClass,omitnil,omitempty" name:"StorageClass"`
}

type MediaCastEvent struct {
	// 媒体转推 ID。
	CastId *string `json:"CastId,omitnil,omitempty" name:"CastId"`

	// 转推状态，取值有：
	// <li>Working ：运行中；</li>
	// <li>Scheduled ：等待定时时间到达后启动；</li>
	// <li>Stopped ：已经停止转推；</li>
	// <li>Idle ：空闲。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type MediaClassInfo struct {
	// 分类 ID。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 父类 ID，一级分类的父类 ID 为 -1。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 分类名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类级别，一级分类为 0，最大值为 3，即最多允许 4 级分类层次。
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 当前分类的第一级子类 ID 集合。
	SubClassIdSet []*int64 `json:"SubClassIdSet,omitnil,omitempty" name:"SubClassIdSet"`

	// 分类名称（该字段已不推荐使用，建议使用新的分类名称字段 Name）。
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`
}

type MediaContentReviewAsrTextSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 嫌疑片段音视频审核的结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 嫌疑关键词列表。
	KeywordSet []*string `json:"KeywordSet,omitnil,omitempty" name:"KeywordSet"`
}

type MediaContentReviewOcrTextSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 嫌疑片段音视频审核的结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 嫌疑关键词列表。
	KeywordSet []*string `json:"KeywordSet,omitnil,omitempty" name:"KeywordSet"`

	// 嫌疑文字出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	// PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type MediaContentReviewPoliticalSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段分数。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 嫌疑片段涉及令人不适宜的信息的结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 涉及令人不适宜的信息、违规图标名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 嫌疑片段涉及令人不适宜的信息的结果标签。音视频审核模板[画面涉及令人不适宜的信息的任务控制参数](https://cloud.tencent.com/document/api/266/31773#PoliticalImgReviewTemplateInfo)里 LabelSet 参数与此参数取值范围的对应关系：
	// violation_photo：
	// <li>violation_photo：违规图标。</li>
	// politician：
	// <li>nation_politician：国家领导人；</li>
	// <li>province_politician: 省部级领导人；</li>
	// <li>bureau_politician：厅局级领导人；</li>
	// <li>county_politician：县处级领导人；</li>
	// <li>rural_politician：乡科级领导人；</li>
	// <li>sensitive_politician：违规相关人物；</li>
	// <li>foreign_politician：国外领导人。</li>
	// entertainment：
	// <li>sensitive_entertainment：违规娱乐人物。</li>
	// sport：
	// <li>sensitive_sport：违规体育人物。</li>
	// entrepreneur：
	// <li>sensitive_entrepreneur：违规商业人物。</li>
	// scholar：
	// <li>sensitive_scholar：违规教育学者。</li>
	// celebrity：
	// <li>sensitive_celebrity：违规知名人物；</li>
	// <li>historical_celebrity：历史知名人物。</li>
	// military：
	// <li>sensitive_military：违规相关人物。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	//  PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 涉及令人不适宜的信息、违规图标出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// 该字段已废弃，请使用 PicUrlExpireTime。
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitnil,omitempty" name:"PicUrlExpireTimeStamp"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type MediaContentReviewSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段涉及令人反感的信息的分数。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 嫌疑片段涉及令人反感的信息的结果标签。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 嫌疑片段鉴别涉及令人反感的信息的结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	//  PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 该字段已废弃，请使用 PicUrlExpireTime。
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitnil,omitempty" name:"PicUrlExpireTimeStamp"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type MediaDeleteItem struct {
	// 所指定的删除部分。如果未填写该字段则参数无效。可选值有：
	// <li>OriginalFiles（删除原文件，删除后无法发起转码、微信发布等任何视频处理操作）；</li>
	// <li>TranscodeFiles（删除转码文件）；</li>
	// <li>AdaptiveDynamicStreamingFiles（删除转自适应码流文件）；</li>
	// <li>WechatPublishFiles（删除微信发布文件）；</li>
	// <li>WechatMiniProgramPublishFiles（删除微信小程序发布文件）。</li>
	// <font color=red>注意：</font> <li>取值为OriginalFiles时，文件上传时携带的封面文件会被删除；</li>
	// <li>取值为TranscodeFiles时，媒体处理产生的封面文件会被删除。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 删除由Type参数指定的种类下的视频模板号，模板定义参见[转码模板](https://cloud.tencent.com/document/product/266/33478#.3Cspan-id-.3D-.22zm.22-.3E.3C.2Fspan.3E.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	// 默认值为0，表示删除参数Type指定种类下所有的视频。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type MediaImageSpriteInfo struct {
	// 特定规格的雪碧图信息集合，每个元素代表一套相同规格的雪碧图。
	ImageSpriteSet []*MediaImageSpriteItem `json:"ImageSpriteSet,omitnil,omitempty" name:"ImageSpriteSet"`
}

type MediaImageSpriteItem struct {
	// 雪碧图规格，参见[雪碧图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 雪碧图小图的高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 雪碧图小图的宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 每一张雪碧图大图里小图的数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 每一张雪碧图大图的地址。
	ImageUrlSet []*string `json:"ImageUrlSet,omitnil,omitempty" name:"ImageUrlSet"`

	// 雪碧图子图位置与时间关系的 WebVtt 文件地址。WebVtt 文件表明了各个雪碧图小图对应的时间点，以及在雪碧大图里的坐标位置，一般被播放器用于实现预览。
	WebVttUrl *string `json:"WebVttUrl,omitnil,omitempty" name:"WebVttUrl"`
}

type MediaInfo struct {
	// 基础信息。包括视频名称、分类、播放地址、封面图片等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicInfo *MediaBasicInfo `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 元信息。包括大小、时长、视频流信息、音频流信息等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 转码结果信息。包括该视频转码生成的各种码率的视频的地址、规格、码率、分辨率等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeInfo *MediaTranscodeInfo `json:"TranscodeInfo,omitnil,omitempty" name:"TranscodeInfo"`

	// 转动图结果信息。对视频转动图（如 gif）后，动图相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicsInfo *MediaAnimatedGraphicsInfo `json:"AnimatedGraphicsInfo,omitnil,omitempty" name:"AnimatedGraphicsInfo"`

	// 采样截图信息。对视频采样截图后，相关截图信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotInfo *MediaSampleSnapshotInfo `json:"SampleSnapshotInfo,omitnil,omitempty" name:"SampleSnapshotInfo"`

	// 雪碧图信息。对视频截取雪碧图之后，雪碧的相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteInfo *MediaImageSpriteInfo `json:"ImageSpriteInfo,omitnil,omitempty" name:"ImageSpriteInfo"`

	// 指定时间点截图信息。对视频依照指定时间点截图后，各个截图的信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetInfo *MediaSnapshotByTimeOffsetInfo `json:"SnapshotByTimeOffsetInfo,omitnil,omitempty" name:"SnapshotByTimeOffsetInfo"`

	// 视频打点信息。对视频设置的各个打点信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyFrameDescInfo *MediaKeyFrameDescInfo `json:"KeyFrameDescInfo,omitnil,omitempty" name:"KeyFrameDescInfo"`

	// 转自适应码流信息。包括规格、加密类型、打包格式等相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingInfo *MediaAdaptiveDynamicStreamingInfo `json:"AdaptiveDynamicStreamingInfo,omitnil,omitempty" name:"AdaptiveDynamicStreamingInfo"`

	// 小程序审核信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniProgramReviewInfo *MediaMiniProgramReviewInfo `json:"MiniProgramReviewInfo,omitnil,omitempty" name:"MiniProgramReviewInfo"`

	// 字幕信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubtitleInfo *MediaSubtitleInfo `json:"SubtitleInfo,omitnil,omitempty" name:"SubtitleInfo"`

	// 媒体文件唯一标识 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 审核信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewInfo *FileReviewInfo `json:"ReviewInfo,omitnil,omitempty" name:"ReviewInfo"`
}

type MediaInputInfo struct {
	// 视频 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 视频名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 视频自定义 ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type MediaKeyFrameDescInfo struct {
	// 视频打点信息数组。
	KeyFrameDescSet []*MediaKeyFrameDescItem `json:"KeyFrameDescSet,omitnil,omitempty" name:"KeyFrameDescSet"`
}

type MediaKeyFrameDescItem struct {
	// 打点的视频偏移时间，单位：秒。
	TimeOffset *float64 `json:"TimeOffset,omitnil,omitempty" name:"TimeOffset"`

	// 打点的内容字符串，限制 1-128 个字符。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type MediaMetaData struct {
	// 上传的媒体文件大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 视频流码率平均值与音频流码率平均值之和，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 视频拍摄时的选择角度，单位：度。
	Rotate *int64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitnil,omitempty" name:"VideoStreamSet"`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitnil,omitempty" name:"AudioStreamSet"`

	// 视频时长，单位：秒。
	VideoDuration *float64 `json:"VideoDuration,omitnil,omitempty" name:"VideoDuration"`

	// 音频时长，单位：秒。
	AudioDuration *float64 `json:"AudioDuration,omitnil,omitempty" name:"AudioDuration"`

	// 媒体文件的 Md5 值。
	// <li><font color=red>注意</font>：如需要获取媒体文件的 Md5，调用 DescribeFileAttributes 接口，待任务执行完成后获取。</li>
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
}

type MediaMiniProgramReviewElem struct {
	// 音视频审核类型。 
	// <li>Porn：画面涉及令人反感的信息，</li>
	// <li>Porn.Ocr：文字涉及令人反感的信息，</li>
	// <li>Porn.Asr：声音涉及令人反感的信息，</li>
	// <li>Terrorism：画面涉及令人不安全的信息，</li>
	// <li>Political：画面涉及令人不适宜的信息，</li>
	// <li>Political.Ocr：文字涉及令人不适宜的信息，</li>
	// <li>Political.Asr：声音涉及令人不适宜的信息。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 音视频审核意见。
	// <li>pass：确认正常，</li>
	// <li>block：确认违规，</li>
	// <li>review：疑似违规。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 音视频审核结果置信度。取值 0~100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type MediaMiniProgramReviewInfo struct {
	// 音视频审核信息列表。
	MiniProgramReviewList []*MediaMiniProgramReviewInfoItem `json:"MiniProgramReviewList,omitnil,omitempty" name:"MiniProgramReviewList"`
}

type MediaMiniProgramReviewInfoItem struct {
	// 模板id。小程序视频发布的视频所对应的转码模板ID，为0代表原始视频。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 视频元信息。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 小程序音视频审核视频播放地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 小程序视频发布状态：
	// <li>Pass：成功。</li>
	// <li>Rejected：未通过。</li>
	ReviewResult *string `json:"ReviewResult,omitnil,omitempty" name:"ReviewResult"`

	// 小程序音视频审核元素。
	ReviewSummary []*MediaMiniProgramReviewElem `json:"ReviewSummary,omitnil,omitempty" name:"ReviewSummary"`
}

type MediaOutputInfo struct {
	// 输出文件 Bucket 所属地域，如 ap-guangzhou  。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 输出文件 Bucket 。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 输出文件目录，目录名必须以 "/" 结尾。
	Dir *string `json:"Dir,omitnil,omitempty" name:"Dir"`
}

type MediaProcessTaskAdaptiveDynamicStreamingResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 转自适应码流任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 对视频转自适应码流任务的输入。
	Input *AdaptiveDynamicStreamingTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频转自适应码流任务的输出。
	Output *AdaptiveDynamicStreamingInfoItem `json:"Output,omitnil,omitempty" name:"Output"`

	// 转自适应码流任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 转自适应码流任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type MediaProcessTaskAnimatedGraphicResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 转动图任务的输入。
	Input *AnimatedGraphicTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 转动图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaAnimatedGraphicsItem `json:"Output,omitnil,omitempty" name:"Output"`

	// 转动图任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 转动图任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 转动图任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type MediaProcessTaskCoverBySnapshotResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 对视频截图做封面任务的输入。
	Input *CoverBySnapshotTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频截图做封面任务的输出。
	Output *CoverBySnapshotTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 对视频截图做封面任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 截图做封面任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 截图做封面任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type MediaProcessTaskImageSpriteResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 对视频截雪碧图任务的输入。
	Input *ImageSpriteTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频截雪碧图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaImageSpriteItem `json:"Output,omitnil,omitempty" name:"Output"`

	// 对视频截雪碧图任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 截雪碧图任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 截雪碧图任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type MediaProcessTaskInput struct {
	// 视频转码任务列表。
	TranscodeTaskSet []*TranscodeTaskInput `json:"TranscodeTaskSet,omitnil,omitempty" name:"TranscodeTaskSet"`

	// 视频转动图任务列表。
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInput `json:"AnimatedGraphicTaskSet,omitnil,omitempty" name:"AnimatedGraphicTaskSet"`

	// 对视频按时间点截图任务列表。
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTaskSet,omitnil,omitempty" name:"SnapshotByTimeOffsetTaskSet"`

	// 对视频采样截图任务列表。
	SampleSnapshotTaskSet []*SampleSnapshotTaskInput `json:"SampleSnapshotTaskSet,omitnil,omitempty" name:"SampleSnapshotTaskSet"`

	// 对视频截雪碧图任务列表。
	ImageSpriteTaskSet []*ImageSpriteTaskInput `json:"ImageSpriteTaskSet,omitnil,omitempty" name:"ImageSpriteTaskSet"`

	// 对视频截图做封面任务列表。
	CoverBySnapshotTaskSet []*CoverBySnapshotTaskInput `json:"CoverBySnapshotTaskSet,omitnil,omitempty" name:"CoverBySnapshotTaskSet"`

	// 对视频转自适应码流任务列表。
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTaskSet,omitnil,omitempty" name:"AdaptiveDynamicStreamingTaskSet"`
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
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 视频转码任务的查询结果，当任务类型为 Transcode 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitnil,omitempty" name:"TranscodeTask"`

	// 视频转动图任务的查询结果，当任务类型为 AnimatedGraphics 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitnil,omitempty" name:"AnimatedGraphicTask"`

	// 对视频按时间点截图任务的查询结果，当任务类型为 SnapshotByTimeOffset 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitnil,omitempty" name:"SnapshotByTimeOffsetTask"`

	// 对视频采样截图任务的查询结果，当任务类型为 SampleSnapshot 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitnil,omitempty" name:"SampleSnapshotTask"`

	// 对视频截雪碧图任务的查询结果，当任务类型为 ImageSprite 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitnil,omitempty" name:"ImageSpriteTask"`

	// 对视频截图做封面任务的查询结果，当任务类型为 CoverBySnapshot 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverBySnapshotTask *MediaProcessTaskCoverBySnapshotResult `json:"CoverBySnapshotTask,omitnil,omitempty" name:"CoverBySnapshotTask"`

	// 对视频转自适应码流任务的查询结果，当任务类型为 AdaptiveDynamicStreaming 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingTask *MediaProcessTaskAdaptiveDynamicStreamingResult `json:"AdaptiveDynamicStreamingTask,omitnil,omitempty" name:"AdaptiveDynamicStreamingTask"`
}

type MediaProcessTaskSampleSnapshotResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 对视频做采样截图任务输入。
	Input *SampleSnapshotTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频做采样截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSampleSnapshotItem `json:"Output,omitnil,omitempty" name:"Output"`

	// 对视频做采样截图任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 采样截图任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 采样截图任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type MediaProcessTaskSnapshotByTimeOffsetResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 对视频按指定时间点截图任务输入。
	Input *SnapshotByTimeOffsetTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频按指定时间点截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSnapshotByTimeOffsetItem `json:"Output,omitnil,omitempty" name:"Output"`

	// 对视频按指定时间点截图任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 时间点截图任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 时间点截图任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type MediaProcessTaskTranscodeResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 转码任务的输入。
	Input *TranscodeTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 转码任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaTranscodeItem `json:"Output,omitnil,omitempty" name:"Output"`

	// 转码进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 转码任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 转码任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type MediaSampleSnapshotInfo struct {
	// 特定规格的采样截图信息集合，每个元素代表一套相同规格的采样截图。
	SampleSnapshotSet []*MediaSampleSnapshotItem `json:"SampleSnapshotSet,omitnil,omitempty" name:"SampleSnapshotSet"`
}

type MediaSampleSnapshotItem struct {
	// 采样截图规格 ID，参见[采样截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 采样方式，取值范围：
	// <li>Percent：根据百分比间隔采样。</li>
	// <li>Time：根据时间间隔采样。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔
	// <li>当 SampleType 为 Percent 时，该值表示多少百分比一张图。</li>
	// <li>当 SampleType 为 Time 时，该值表示多少时间间隔一张图，单位秒， 第一张图均为视频首帧。</li>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 生成的截图 url 列表。
	ImageUrlSet []*string `json:"ImageUrlSet,omitnil,omitempty" name:"ImageUrlSet"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitnil,omitempty" name:"WaterMarkDefinition"`
}

type MediaSnapshotByTimeOffsetInfo struct {
	// 特定规格的指定时间点截图信息集合。目前每种规格只能有一套截图。
	SnapshotByTimeOffsetSet []*MediaSnapshotByTimeOffsetItem `json:"SnapshotByTimeOffsetSet,omitnil,omitempty" name:"SnapshotByTimeOffsetSet"`
}

type MediaSnapshotByTimeOffsetItem struct {
	// 指定时间点截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 同一规格的截图信息集合，每个元素代表一张截图。
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet,omitnil,omitempty" name:"PicInfoSet"`
}

type MediaSnapshotByTimePicInfoItem struct {
	// 该张截图对应视频文件中的时间偏移，单位为<font color=red>毫秒</font>。
	TimeOffset *float64 `json:"TimeOffset,omitnil,omitempty" name:"TimeOffset"`

	// 该张截图的 URL 地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitnil,omitempty" name:"WaterMarkDefinition"`
}

type MediaSourceData struct {
	// 媒体文件的来源类别：
	// <li>Record：来自录制。如直播录制、直播时移录制等。</li>
	// <li>Upload：来自上传。如拉取上传、服务端上传、客户端 UGC 上传等。</li>
	// <li>VideoProcessing：来自视频处理。如视频拼接、视频剪辑等。</li>
	// <li>TrtcRecord：来自TRTC 伴生录制。</li>
	// <li>WebPageRecord：来自全景录制。</li>
	// <li>Unknown：未知来源。</li>
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 用户创建文件时透传的字段。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`

	// 直播录制信息，当文件来源为 Record 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveRecordInfo *LiveRecordInfo `json:"LiveRecordInfo,omitnil,omitempty" name:"LiveRecordInfo"`

	// TRTC 伴生录制信息，当文件来源为 TrtcRecord 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrtcRecordInfo *TrtcRecordInfo `json:"TrtcRecordInfo,omitnil,omitempty" name:"TrtcRecordInfo"`

	// 全景录制信息，当文件来源为 WebPageRecord 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebPageRecordInfo *WebPageRecordInfo `json:"WebPageRecordInfo,omitnil,omitempty" name:"WebPageRecordInfo"`
}

type MediaSubStreamInfoItem struct {
	// 子流类型，取值范围：
	// <li>audio：纯音频；</li>
	// <li>video：视频（可能包含音频流）。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当子流为视频流时，视频画面宽度，单位：px。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 当子流为视频流时，视频画面高度，单位：px。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 子流媒体文件大小，单位：Byte。
	// <font color=red>注意：</font>在 2023-02-09T16:00:00Z 前处理生成的自适应码流文件此字段为0。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type MediaSubtitleInfo struct {
	// 字幕信息列表。
	SubtitleSet []*MediaSubtitleItem `json:"SubtitleSet,omitnil,omitempty" name:"SubtitleSet"`
}

type MediaSubtitleInput struct {
	// 字幕名字，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字幕语言。常见的取值如下：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>ja：日文；</li>
	// <li>ko：韩文；</li>
	// <li>vi：越南语；</li>
	// <li>ms：马来语；</li>
	// <li>th：泰语；</li>
	// <li>pt：葡萄牙语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ar：阿拉伯语；</li>
	// <li>es：西班牙语；</li>
	// <li>hi：印地语；</li>
	// <li>fr：法语。</li>
	// 其他取值参考 [RFC5646](https://tools.ietf.org/html/rfc5646)
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 字幕格式。取值范围如下：
	// <li>vtt</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 字幕内容，进行 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 字幕的唯一标识。长度不能超过16个字符，可以使用大小写字母、数字、下划线（_）或横杠（-）。不能与媒体文件中现有字幕的唯一标识重复。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type MediaSubtitleItem struct {
	// 字幕的唯一标识。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 字幕名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字幕语言。常见的取值如下：
	// <li>zh：中文；</li>
	// <li>en：英文；</li>
	// <li>ja：日文；</li>
	// <li>ko：韩文；</li>
	// <li>vi：越南语；</li>
	// <li>ms：马来语；</li>
	// <li>th：泰语；</li>
	// <li>pt：葡萄牙语；</li>
	// <li>tr：土耳其语；</li>
	// <li>ar：阿拉伯语；</li>
	// <li>es：西班牙语；</li>
	// <li>hi：印地语；</li>
	// <li>fr：法语。</li>
	// 其他取值参考 [RFC5646](https://tools.ietf.org/html/rfc5646)
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 字幕格式。取值范围如下：
	// <li>vtt</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 字幕 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 字幕来源，取值范围：
	// <li>UserUploaded：用户上传；</li>
	// <li>AIRecognition：AI 识别，通过语音识别或语音翻译生成。</li>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type MediaTrack struct {
	// 轨道类型，取值有：
	// <ul>
	// <li>Video ：视频轨道。视频轨道由以下 Item 组成：<ul><li>VideoTrackItem</li><li>MediaTransitionItem</li> <li>EmptyTrackItem</li></ul> </li>
	// <li>Audio ：音频轨道。音频轨道由以下 Item 组成：<ul><li>AudioTrackItem</li><li>MediaTransitionItem</li><li>EmptyTrackItem</li></ul></li>
	// <li>Sticker ：贴图轨道。贴图轨道以下 Item 组成：<ul><li> StickerTrackItem</li><li>EmptyTrackItem</li></ul></li>	
	// </ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 轨道上的媒体片段列表。
	TrackItems []*MediaTrackItem `json:"TrackItems,omitnil,omitempty" name:"TrackItems"`
}

type MediaTrackItem struct {
	// 片段类型。取值有：
	// <li>Video：视频片段。</li>
	// <li>Audio：音频片段。</li>
	// <li>Sticker：贴图片段。</li>
	// <li>Transition：转场。</li>
	// <li>Empty：空白片段。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 视频片段，当 Type = Video 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoItem *VideoTrackItem `json:"VideoItem,omitnil,omitempty" name:"VideoItem"`

	// 音频片段，当 Type = Audio 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioItem *AudioTrackItem `json:"AudioItem,omitnil,omitempty" name:"AudioItem"`

	// 贴图片段，当 Type = Sticker 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StickerItem *StickerTrackItem `json:"StickerItem,omitnil,omitempty" name:"StickerItem"`

	// 转场，当 Type = Transition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransitionItem *MediaTransitionItem `json:"TransitionItem,omitnil,omitempty" name:"TransitionItem"`

	// 空白片段，当 Type = Empty 时有效。空片段用于时间轴的占位。<li>如需要两个音频片段之间有一段时间的静音，可以用 EmptyTrackItem 来进行占位。</li>
	// <li>使用 EmptyTrackItem 进行占位，来定位某个Item。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmptyItem *EmptyTrackItem `json:"EmptyItem,omitnil,omitempty" name:"EmptyItem"`
}

type MediaTranscodeInfo struct {
	// 各规格的转码信息集合，每个元素代表一个规格的转码结果。
	TranscodeSet []*MediaTranscodeItem `json:"TranscodeSet,omitnil,omitempty" name:"TranscodeSet"`
}

type MediaTranscodeItem struct {
	// 转码后的视频文件地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/266/33476)。
	// <font color=red>注意：取值 0 表示原始文件。</font> 
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 媒体文件总大小，单位：字节。
	// <li>当媒体文件为 HLS 时，大小是 m3u8 和 ts 文件大小的总和。</li>
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 视频的 md5 值。
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitnil,omitempty" name:"VideoStreamSet"`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitnil,omitempty" name:"AudioStreamSet"`

	// 数字水印类型。可选值：
	// <li>Trace 表示经过溯源水印处理；</li>
	// <li>CopyRight 表示经过版权水印处理；</li>
	// <li>None 表示没有经过数字水印处理。</li>
	DigitalWatermarkType *string `json:"DigitalWatermarkType,omitnil,omitempty" name:"DigitalWatermarkType"`

	// 版权信息。
	CopyRightWatermarkText *string `json:"CopyRightWatermarkText,omitnil,omitempty" name:"CopyRightWatermarkText"`
}

type MediaTransitionItem struct {
	// 转场持续时间，单位为秒。进行转场处理的两个媒体片段，第二个片段在轨道上的起始时间会自动进行调整，设置为前面一个片段的结束时间减去转场的持续时间。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 转场操作列表。图像转场操作和音频转场操作各自最多支持一个。
	//
	// Deprecated: Transitions is deprecated.
	Transitions []*TransitionOpertion `json:"Transitions,omitnil,omitempty" name:"Transitions"`

	// 转场操作列表。图像转场操作和音频转场操作各自最多支持一个。
	MediaTransitions []*TransitionOperation `json:"MediaTransitions,omitnil,omitempty" name:"MediaTransitions"`
}

type MediaVideoStreamItem struct {
	// 视频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 视频流的高度，单位：px。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 视频流的宽度，单位：px。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频流的编码格式，例如 h264。
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 帧率，单位：hz。
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 编码标签，仅当 Codec 为 hevc 时有效。
	CodecTag *string `json:"CodecTag,omitnil,omitempty" name:"CodecTag"`

	// 画面动态范围信息。
	// <li><font color=red>注意</font>：在 2023-01-10T00:00:00Z 后处理的转码文件，此字段有效。</li>
	DynamicRangeInfo *DynamicRangeInfo `json:"DynamicRangeInfo,omitnil,omitempty" name:"DynamicRangeInfo"`
}

// Predefined struct for user
type ModifyAIAnalysisTemplateRequestParams struct {
	// 音视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfoForUpdate `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfoForUpdate `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfoForUpdate `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfoForUpdate `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`

	// 智能精彩集锦任务控制参数。
	HighlightConfigure *HighlightsConfigureInfoForUpdate `json:"HighlightConfigure,omitnil,omitempty" name:"HighlightConfigure"`
}

type ModifyAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 音视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfoForUpdate `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfoForUpdate `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfoForUpdate `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfoForUpdate `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`

	// 智能精彩集锦任务控制参数。
	HighlightConfigure *HighlightsConfigureInfoForUpdate `json:"HighlightConfigure,omitnil,omitempty" name:"HighlightConfigure"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ClassificationConfigure")
	delete(f, "TagConfigure")
	delete(f, "CoverConfigure")
	delete(f, "FrameTagConfigure")
	delete(f, "HighlightConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAIAnalysisTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAIAnalysisTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 音视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频片头片尾识别控制参数。
	HeadTailConfigure *HeadTailConfigureInfoForUpdate `json:"HeadTailConfigure,omitnil,omitempty" name:"HeadTailConfigure"`

	// 视频拆条识别控制参数。
	SegmentConfigure *SegmentConfigureInfoForUpdate `json:"SegmentConfigure,omitnil,omitempty" name:"SegmentConfigure"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	AsrTranslateConfigure *AsrTranslateConfigureInfoForUpdate `json:"AsrTranslateConfigure,omitnil,omitempty" name:"AsrTranslateConfigure"`

	// 物体识别控制参数。
	ObjectConfigure *ObjectConfigureInfoForUpdate `json:"ObjectConfigure,omitnil,omitempty" name:"ObjectConfigure"`

	// 截帧间隔，单位为秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`
}

type ModifyAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 音视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频片头片尾识别控制参数。
	HeadTailConfigure *HeadTailConfigureInfoForUpdate `json:"HeadTailConfigure,omitnil,omitempty" name:"HeadTailConfigure"`

	// 视频拆条识别控制参数。
	SegmentConfigure *SegmentConfigureInfoForUpdate `json:"SegmentConfigure,omitnil,omitempty" name:"SegmentConfigure"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`

	// 语音翻译控制参数。
	AsrTranslateConfigure *AsrTranslateConfigureInfoForUpdate `json:"AsrTranslateConfigure,omitnil,omitempty" name:"AsrTranslateConfigure"`

	// 物体识别控制参数。
	ObjectConfigure *ObjectConfigureInfoForUpdate `json:"ObjectConfigure,omitnil,omitempty" name:"ObjectConfigure"`

	// 截帧间隔，单位为秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "HeadTailConfigure")
	delete(f, "SegmentConfigure")
	delete(f, "FaceConfigure")
	delete(f, "OcrFullTextConfigure")
	delete(f, "OcrWordsConfigure")
	delete(f, "AsrFullTextConfigure")
	delete(f, "AsrWordsConfigure")
	delete(f, "AsrTranslateConfigure")
	delete(f, "ObjectConfigure")
	delete(f, "ScreenshotInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAIRecognitionTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 自适应转码模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自适应转码格式，取值范围：
	// <li>HLS；</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// 自适应转码输入流参数信息，最多输入10路流。
	// 注意：各个流的帧率必须保持一致；如果不一致，采用第一个流的帧率作为输出帧率。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 切片类型，当 Format 为 HLS 时有效，可选值：
	// <li>ts：ts 切片；</li>
	// <li>fmp4：fmp4 切片。</li>
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
}

type ModifyAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 自适应转码模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自适应转码格式，取值范围：
	// <li>HLS；</li>
	// <li>MPEG-DASH。</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 是否禁止视频低码率转高码率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// 是否禁止视频分辨率转高分辨率，取值范围：
	// <li>0：否，</li>
	// <li>1：是。</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// 自适应转码输入流参数信息，最多输入10路流。
	// 注意：各个流的帧率必须保持一致；如果不一致，采用第一个流的帧率作为输出帧率。
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 切片类型，当 Format 为 HLS 时有效，可选值：
	// <li>ts：ts 切片；</li>
	// <li>fmp4：fmp4 切片。</li>
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Format")
	delete(f, "DisableHigherVideoBitrate")
	delete(f, "DisableHigherVideoResolution")
	delete(f, "StreamInfos")
	delete(f, "Comment")
	delete(f, "SegmentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAdaptiveDynamicStreamingTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 动图格式，取值为 gif 和 webp。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 帧率，取值范围：[1, 30]，单位：Hz。
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ModifyAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 动图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 动图格式，取值为 gif 和 webp。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 帧率，取值范围：[1, 30]，单位：Hz。
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
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
	delete(f, "SubAppId")
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyClassRequestParams struct {
	// 分类 ID
	ClassId *uint64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 分类名称。长度限制：1-64 个字符。
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type ModifyClassRequest struct {
	*tchttp.BaseRequest
	
	// 分类 ID
	ClassId *uint64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 分类名称。长度限制：1-64 个字符。
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *ModifyClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClassId")
	delete(f, "ClassName")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClassResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClassResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClassResponseParams `json:"Response"`
}

func (r *ModifyClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContentReviewTemplateRequestParams struct {
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 令人不安全的信息的控制参数。
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// 令人反感的信息的控制参数。
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// 令人不适宜的信息控制参数。
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	ProhibitedConfigure *ProhibitedConfigureInfoForUpdate `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`

	// 截帧间隔，单位为秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`

	// 审核结果是否进入审核墙（对审核结果进行人工识别）的开关。
	// <li>ON：是；</li>
	// <li>OFF：否。</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitnil,omitempty" name:"ReviewWallSwitch"`
}

type ModifyContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 令人不安全的信息的控制参数。
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// 令人反感的信息的控制参数。
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// 令人不适宜的信息控制参数。
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// 违禁控制参数。违禁内容包括：
	// <li>谩骂；</li>
	// <li>涉毒违法。</li>
	ProhibitedConfigure *ProhibitedConfigureInfoForUpdate `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`

	// 截帧间隔，单位为秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`

	// 审核结果是否进入审核墙（对审核结果进行人工识别）的开关。
	// <li>ON：是；</li>
	// <li>OFF：否。</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitnil,omitempty" name:"ReviewWallSwitch"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "TerrorismConfigure")
	delete(f, "PornConfigure")
	delete(f, "PoliticalConfigure")
	delete(f, "ProhibitedConfigure")
	delete(f, "UserDefineConfigure")
	delete(f, "ScreenshotInterval")
	delete(f, "ReviewWallSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContentReviewTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyDefaultDistributionConfigRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 分发域名，取值为点播域名列表里的域名。不填或者填空，表示不修改域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分发协议，取值为 HTTP 或者 HTTPS。
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// 播放密钥，由大小写字母（a - Z）或者数字（0 - 9）组成，长度在8 - 20个字符之间。
	PlayKey *string `json:"PlayKey,omitnil,omitempty" name:"PlayKey"`
}

type ModifyDefaultDistributionConfigRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 分发域名，取值为点播域名列表里的域名。不填或者填空，表示不修改域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分发协议，取值为 HTTP 或者 HTTPS。
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// 播放密钥，由大小写字母（a - Z）或者数字（0 - 9）组成，长度在8 - 20个字符之间。
	PlayKey *string `json:"PlayKey,omitnil,omitempty" name:"PlayKey"`
}

func (r *ModifyDefaultDistributionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultDistributionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Domain")
	delete(f, "Scheme")
	delete(f, "PlayKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDefaultDistributionConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultDistributionConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDefaultDistributionConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDefaultDistributionConfigResponseParams `json:"Response"`
}

func (r *ModifyDefaultDistributionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultDistributionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultStorageRegionRequestParams struct {
	// 默认的存储地域，必须是已经开通的地域（通过 DescribeStorageRegions 接口查询）。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type ModifyDefaultStorageRegionRequest struct {
	*tchttp.BaseRequest
	
	// 默认的存储地域，必须是已经开通的地域（通过 DescribeStorageRegions 接口查询）。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *ModifyDefaultStorageRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultStorageRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StorageRegion")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDefaultStorageRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultStorageRegionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDefaultStorageRegionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDefaultStorageRegionResponseParams `json:"Response"`
}

func (r *ModifyDefaultStorageRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultStorageRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnhanceMediaTemplateRequestParams struct {
	// 音画质重生模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 音画质重生视频控制控制信息。
	RebuildVideoInfo *RebuildVideoInfo `json:"RebuildVideoInfo,omitnil,omitempty" name:"RebuildVideoInfo"`

	// 音画质重生音频控制控制信息。
	RebuildAudioInfo *RebuildAudioInfo `json:"RebuildAudioInfo,omitnil,omitempty" name:"RebuildAudioInfo"`

	// 输出目标视频控制信息。
	TargetVideoInfo *RebuildMediaTargetVideoStream `json:"TargetVideoInfo,omitnil,omitempty" name:"TargetVideoInfo"`

	// 输出目标音频控制信息。
	TargetAudioInfo *RebuildMediaTargetAudioStream `json:"TargetAudioInfo,omitnil,omitempty" name:"TargetAudioInfo"`

	// 输出文件封装格式，可选值：mp4、flv、hls。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

type ModifyEnhanceMediaTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 音画质重生模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 音画质重生模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 音画质重生视频控制控制信息。
	RebuildVideoInfo *RebuildVideoInfo `json:"RebuildVideoInfo,omitnil,omitempty" name:"RebuildVideoInfo"`

	// 音画质重生音频控制控制信息。
	RebuildAudioInfo *RebuildAudioInfo `json:"RebuildAudioInfo,omitnil,omitempty" name:"RebuildAudioInfo"`

	// 输出目标视频控制信息。
	TargetVideoInfo *RebuildMediaTargetVideoStream `json:"TargetVideoInfo,omitnil,omitempty" name:"TargetVideoInfo"`

	// 输出目标音频控制信息。
	TargetAudioInfo *RebuildMediaTargetAudioStream `json:"TargetAudioInfo,omitnil,omitempty" name:"TargetAudioInfo"`

	// 输出文件封装格式，可选值：mp4、flv、hls。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

func (r *ModifyEnhanceMediaTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnhanceMediaTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RebuildVideoInfo")
	delete(f, "RebuildAudioInfo")
	delete(f, "TargetVideoInfo")
	delete(f, "TargetAudioInfo")
	delete(f, "Container")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnhanceMediaTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnhanceMediaTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEnhanceMediaTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnhanceMediaTemplateResponseParams `json:"Response"`
}

func (r *ModifyEnhanceMediaTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnhanceMediaTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEventConfigRequestParams struct {
	// 接收事件通知的方式。
	// <li>PUSH：[HTTP 回调通知](https://cloud.tencent.com/document/product/266/33779)；</li>
	// <li>PULL：[基于消息队列的可靠通知](https://cloud.tencent.com/document/product/266/33779)。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 采用 [HTTP 回调通知](https://cloud.tencent.com/document/product/266/33779) 接收方式时，用于接收 3.0 格式回调的地址。
	// 注意：如果带 NotificationUrl  参数且值为空字符串时将会清空 3.0 格式回调地址。
	NotificationUrl *string `json:"NotificationUrl,omitnil,omitempty" name:"NotificationUrl"`

	// 是否接收 [视频上传完成](https://cloud.tencent.com/document/product/266/7830) 事件通知， 默认 "OFF" 为忽略该事件通知，"ON" 为接收事件通知。
	UploadMediaCompleteEventSwitch *string `json:"UploadMediaCompleteEventSwitch,omitnil,omitempty" name:"UploadMediaCompleteEventSwitch"`

	// 是否接收 [视频删除完成](https://cloud.tencent.com/document/product/266/13434) 事件通知，  默认 "OFF" 为忽略该事件通知，"ON" 为接收事件通知。
	DeleteMediaCompleteEventSwitch *string `json:"DeleteMediaCompleteEventSwitch,omitnil,omitempty" name:"DeleteMediaCompleteEventSwitch"`

	// 是否接收剪辑固化完成事件通知，  默认 "OFF" 为忽略该事件通知，"ON" 为接收事件通知。
	PersistenceCompleteEventSwitch *string `json:"PersistenceCompleteEventSwitch,omitnil,omitempty" name:"PersistenceCompleteEventSwitch"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type ModifyEventConfigRequest struct {
	*tchttp.BaseRequest
	
	// 接收事件通知的方式。
	// <li>PUSH：[HTTP 回调通知](https://cloud.tencent.com/document/product/266/33779)；</li>
	// <li>PULL：[基于消息队列的可靠通知](https://cloud.tencent.com/document/product/266/33779)。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 采用 [HTTP 回调通知](https://cloud.tencent.com/document/product/266/33779) 接收方式时，用于接收 3.0 格式回调的地址。
	// 注意：如果带 NotificationUrl  参数且值为空字符串时将会清空 3.0 格式回调地址。
	NotificationUrl *string `json:"NotificationUrl,omitnil,omitempty" name:"NotificationUrl"`

	// 是否接收 [视频上传完成](https://cloud.tencent.com/document/product/266/7830) 事件通知， 默认 "OFF" 为忽略该事件通知，"ON" 为接收事件通知。
	UploadMediaCompleteEventSwitch *string `json:"UploadMediaCompleteEventSwitch,omitnil,omitempty" name:"UploadMediaCompleteEventSwitch"`

	// 是否接收 [视频删除完成](https://cloud.tencent.com/document/product/266/13434) 事件通知，  默认 "OFF" 为忽略该事件通知，"ON" 为接收事件通知。
	DeleteMediaCompleteEventSwitch *string `json:"DeleteMediaCompleteEventSwitch,omitnil,omitempty" name:"DeleteMediaCompleteEventSwitch"`

	// 是否接收剪辑固化完成事件通知，  默认 "OFF" 为忽略该事件通知，"ON" 为接收事件通知。
	PersistenceCompleteEventSwitch *string `json:"PersistenceCompleteEventSwitch,omitnil,omitempty" name:"PersistenceCompleteEventSwitch"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *ModifyEventConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEventConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "NotificationUrl")
	delete(f, "UploadMediaCompleteEventSwitch")
	delete(f, "DeleteMediaCompleteEventSwitch")
	delete(f, "PersistenceCompleteEventSwitch")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEventConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEventConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEventConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEventConfigResponseParams `json:"Response"`
}

func (r *ModifyEventConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEventConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHeadTailTemplateRequestParams struct {
	// 片头片尾模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板名，长度限制 64 个字符。不传代表不修改。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述，长度限制 256 个字符。不传代表不修改，传空代表清空。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 片头候选列表，填写视频的 FileId。转码时将自动选择与正片宽高比最接近的一个片头（相同宽高比时，靠前的候选项优先）。最多支持 5 个候选片头。不传代表不修改，传空数组代表清空。
	HeadCandidateSet []*string `json:"HeadCandidateSet,omitnil,omitempty" name:"HeadCandidateSet"`

	// 片尾候选列表，填写视频的 FileId。转码时将自动选择与正片宽高比最接近的一个片尾（相同宽高比时，靠前的候选项优先）。最多支持 5 个候选片头。不传代表不修改，传空数组代表清空。
	TailCandidateSet []*string `json:"TailCandidateSet,omitnil,omitempty" name:"TailCandidateSet"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li> gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊；</li>
	// <li> white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充；</li>
	// <li> black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值为不修改。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type ModifyHeadTailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 片头片尾模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板名，长度限制 64 个字符。不传代表不修改。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述，长度限制 256 个字符。不传代表不修改，传空代表清空。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 片头候选列表，填写视频的 FileId。转码时将自动选择与正片宽高比最接近的一个片头（相同宽高比时，靠前的候选项优先）。最多支持 5 个候选片头。不传代表不修改，传空数组代表清空。
	HeadCandidateSet []*string `json:"HeadCandidateSet,omitnil,omitempty" name:"HeadCandidateSet"`

	// 片尾候选列表，填写视频的 FileId。转码时将自动选择与正片宽高比最接近的一个片尾（相同宽高比时，靠前的候选项优先）。最多支持 5 个候选片头。不传代表不修改，传空数组代表清空。
	TailCandidateSet []*string `json:"TailCandidateSet,omitnil,omitempty" name:"TailCandidateSet"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li> gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊；</li>
	// <li> white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充；</li>
	// <li> black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值为不修改。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

func (r *ModifyHeadTailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHeadTailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "HeadCandidateSet")
	delete(f, "TailCandidateSet")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHeadTailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHeadTailTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyHeadTailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHeadTailTemplateResponseParams `json:"Response"`
}

func (r *ModifyHeadTailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHeadTailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSpriteTemplateRequestParams struct {
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 雪碧图中小图的宽度，取值范围： [128, 4096]，单位：px。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 雪碧图中小图的高度，取值范围： [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 图片格式，取值：
	// <li> jpg：jpg 格式；</li>
	// <li> png：png 格式；</li>
	// <li> webp：webp 格式。</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type ModifyImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 雪碧图中小图的宽度，取值范围： [128, 4096]，单位：px。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 雪碧图中小图的高度，取值范围： [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 图片格式，取值：
	// <li> jpg：jpg 格式；</li>
	// <li> png：png 格式；</li>
	// <li> webp：webp 格式。</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
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
	delete(f, "SubAppId")
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyJustInTimeTranscodeTemplateRequestParams struct {
	// 模板名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 视频参数配置。
	VideoConfigure *VideoConfigureInfoForUpdate `json:"VideoConfigure,omitnil,omitempty" name:"VideoConfigure"`

	// 水印参数配置。
	WatermarkConfigure *WatermarkConfigureInfoForUpdate `json:"WatermarkConfigure,omitnil,omitempty" name:"WatermarkConfigure"`

	// 模板描述，长度限制256个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ModifyJustInTimeTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 视频参数配置。
	VideoConfigure *VideoConfigureInfoForUpdate `json:"VideoConfigure,omitnil,omitempty" name:"VideoConfigure"`

	// 水印参数配置。
	WatermarkConfigure *WatermarkConfigureInfoForUpdate `json:"WatermarkConfigure,omitnil,omitempty" name:"WatermarkConfigure"`

	// 模板描述，长度限制256个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *ModifyJustInTimeTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyJustInTimeTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	delete(f, "VideoConfigure")
	delete(f, "WatermarkConfigure")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyJustInTimeTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyJustInTimeTranscodeTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyJustInTimeTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyJustInTimeTranscodeTemplateResponseParams `json:"Response"`
}

func (r *ModifyJustInTimeTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyJustInTimeTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMediaInfoRequestParams struct {
	// 媒体文件唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 媒体文件名称，最长 64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 媒体文件描述，最长 128 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 媒体文件分类 ID。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 媒体文件过期时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 视频封面图片文件（如 jpeg, png 等）进行 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式。
	CoverData *string `json:"CoverData,omitnil,omitempty" name:"CoverData"`

	// 新增的一组视频打点信息，如果某个偏移时间已存在打点，则会进行覆盖操作，单个媒体文件最多 100 个打点信息。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	AddKeyFrameDescs []*MediaKeyFrameDescItem `json:"AddKeyFrameDescs,omitnil,omitempty" name:"AddKeyFrameDescs"`

	// 要删除的一组视频打点信息的时间偏移，单位：秒。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	DeleteKeyFrameDescs []*float64 `json:"DeleteKeyFrameDescs,omitnil,omitempty" name:"DeleteKeyFrameDescs"`

	// 取值 1 表示清空视频打点信息，其他值无意义。
	// 同一个请求里，ClearKeyFrameDescs 与 AddKeyFrameDescs 不能同时出现。
	ClearKeyFrameDescs *int64 `json:"ClearKeyFrameDescs,omitnil,omitempty" name:"ClearKeyFrameDescs"`

	// 新增的一组标签，单个媒体文件最多 16 个标签，单个标签最多 32 个字符。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	AddTags []*string `json:"AddTags,omitnil,omitempty" name:"AddTags"`

	// 要删除的一组标签。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	DeleteTags []*string `json:"DeleteTags,omitnil,omitempty" name:"DeleteTags"`

	// 取值 1 表示清空媒体文件所有标签，其他值无意义。
	// 同一个请求里，ClearTags 与 AddTags 不能同时出现。
	ClearTags *int64 `json:"ClearTags,omitnil,omitempty" name:"ClearTags"`

	// 新增一组字幕。单个媒体文件最多 16 个字幕。同一个请求中，AddSubtitles 中指定的字幕 Id 必须与 DeleteSubtitleIds 都不相同。
	AddSubtitles []*MediaSubtitleInput `json:"AddSubtitles,omitnil,omitempty" name:"AddSubtitles"`

	// 待删除字幕的唯一标识。同一个请求中，AddSubtitles 中指定的字幕 Id 必须与 DeleteSubtitleIds 都不相同。
	DeleteSubtitleIds []*string `json:"DeleteSubtitleIds,omitnil,omitempty" name:"DeleteSubtitleIds"`

	// 取值 1 表示清空媒体文件所有的字幕信息，其他值无意义。
	// 同一个请求里，ClearSubtitles 与 AddSubtitles不能同时出现。
	ClearSubtitles *int64 `json:"ClearSubtitles,omitnil,omitempty" name:"ClearSubtitles"`
}

type ModifyMediaInfoRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 媒体文件名称，最长 64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 媒体文件描述，最长 128 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 媒体文件分类 ID。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 媒体文件过期时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 视频封面图片文件（如 jpeg, png 等）进行 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式。
	CoverData *string `json:"CoverData,omitnil,omitempty" name:"CoverData"`

	// 新增的一组视频打点信息，如果某个偏移时间已存在打点，则会进行覆盖操作，单个媒体文件最多 100 个打点信息。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	AddKeyFrameDescs []*MediaKeyFrameDescItem `json:"AddKeyFrameDescs,omitnil,omitempty" name:"AddKeyFrameDescs"`

	// 要删除的一组视频打点信息的时间偏移，单位：秒。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	DeleteKeyFrameDescs []*float64 `json:"DeleteKeyFrameDescs,omitnil,omitempty" name:"DeleteKeyFrameDescs"`

	// 取值 1 表示清空视频打点信息，其他值无意义。
	// 同一个请求里，ClearKeyFrameDescs 与 AddKeyFrameDescs 不能同时出现。
	ClearKeyFrameDescs *int64 `json:"ClearKeyFrameDescs,omitnil,omitempty" name:"ClearKeyFrameDescs"`

	// 新增的一组标签，单个媒体文件最多 16 个标签，单个标签最多 32 个字符。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	AddTags []*string `json:"AddTags,omitnil,omitempty" name:"AddTags"`

	// 要删除的一组标签。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	DeleteTags []*string `json:"DeleteTags,omitnil,omitempty" name:"DeleteTags"`

	// 取值 1 表示清空媒体文件所有标签，其他值无意义。
	// 同一个请求里，ClearTags 与 AddTags 不能同时出现。
	ClearTags *int64 `json:"ClearTags,omitnil,omitempty" name:"ClearTags"`

	// 新增一组字幕。单个媒体文件最多 16 个字幕。同一个请求中，AddSubtitles 中指定的字幕 Id 必须与 DeleteSubtitleIds 都不相同。
	AddSubtitles []*MediaSubtitleInput `json:"AddSubtitles,omitnil,omitempty" name:"AddSubtitles"`

	// 待删除字幕的唯一标识。同一个请求中，AddSubtitles 中指定的字幕 Id 必须与 DeleteSubtitleIds 都不相同。
	DeleteSubtitleIds []*string `json:"DeleteSubtitleIds,omitnil,omitempty" name:"DeleteSubtitleIds"`

	// 取值 1 表示清空媒体文件所有的字幕信息，其他值无意义。
	// 同一个请求里，ClearSubtitles 与 AddSubtitles不能同时出现。
	ClearSubtitles *int64 `json:"ClearSubtitles,omitnil,omitempty" name:"ClearSubtitles"`
}

func (r *ModifyMediaInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ClassId")
	delete(f, "ExpireTime")
	delete(f, "CoverData")
	delete(f, "AddKeyFrameDescs")
	delete(f, "DeleteKeyFrameDescs")
	delete(f, "ClearKeyFrameDescs")
	delete(f, "AddTags")
	delete(f, "DeleteTags")
	delete(f, "ClearTags")
	delete(f, "AddSubtitles")
	delete(f, "DeleteSubtitleIds")
	delete(f, "ClearSubtitles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMediaInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMediaInfoResponseParams struct {
	// 新的视频封面 URL。
	// * 注意：仅当请求携带 CoverData 时此返回值有效。 *
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// 新增的字幕信息。
	AddedSubtitleSet []*MediaSubtitleItem `json:"AddedSubtitleSet,omitnil,omitempty" name:"AddedSubtitleSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMediaInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMediaInfoResponseParams `json:"Response"`
}

func (r *ModifyMediaInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMediaStorageClassRequestParams struct {
	// 媒体文件唯一标识列表，最大长度：100。
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 目标存储类型。可选值有：
	// <li> STANDARD：标准存储。</li>
	// <li> STANDARD_IA：低频存储。</li>
	// <li> ARCHIVE：归档存储。</li>
	// <li> DEEP_ARCHIVE：深度归档存储。</li>
	StorageClass *string `json:"StorageClass,omitnil,omitempty" name:"StorageClass"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 取回模式。当文件的存储类型从归档或深度归档转换为标准存储时，需要指定取回（也称为解冻）操作的模式，具体说明请参考[数据取回及取回模式](https://cloud.tencent.com/document/product/266/56196#retake)。
	// 当媒体文件目前的存储类型为归档存储时，有以下取值：
	// <li>Expedited：极速模式。</li>
	// <li>Standard：标准模式。</li>
	// <li>Bulk：批量模式。</li>
	// 当媒体文件目前的存储类型为深度归档存储时，有以下取值：
	// <li>Standard：标准模式。</li>
	// <li>Bulk：批量模式。</li>
	RestoreTier *string `json:"RestoreTier,omitnil,omitempty" name:"RestoreTier"`
}

type ModifyMediaStorageClassRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件唯一标识列表，最大长度：100。
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 目标存储类型。可选值有：
	// <li> STANDARD：标准存储。</li>
	// <li> STANDARD_IA：低频存储。</li>
	// <li> ARCHIVE：归档存储。</li>
	// <li> DEEP_ARCHIVE：深度归档存储。</li>
	StorageClass *string `json:"StorageClass,omitnil,omitempty" name:"StorageClass"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 取回模式。当文件的存储类型从归档或深度归档转换为标准存储时，需要指定取回（也称为解冻）操作的模式，具体说明请参考[数据取回及取回模式](https://cloud.tencent.com/document/product/266/56196#retake)。
	// 当媒体文件目前的存储类型为归档存储时，有以下取值：
	// <li>Expedited：极速模式。</li>
	// <li>Standard：标准模式。</li>
	// <li>Bulk：批量模式。</li>
	// 当媒体文件目前的存储类型为深度归档存储时，有以下取值：
	// <li>Standard：标准模式。</li>
	// <li>Bulk：批量模式。</li>
	RestoreTier *string `json:"RestoreTier,omitnil,omitempty" name:"RestoreTier"`
}

func (r *ModifyMediaStorageClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaStorageClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileIds")
	delete(f, "StorageClass")
	delete(f, "SubAppId")
	delete(f, "RestoreTier")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMediaStorageClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMediaStorageClassResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMediaStorageClassResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMediaStorageClassResponseParams `json:"Response"`
}

func (r *ModifyMediaStorageClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaStorageClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonSampleRequestParams struct {
	// 素材 ID。
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 名称，长度限制：128 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 素材应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于不适宜的内容识别，等价于 Review.Face。
	// 3. All：用于内容识别、不适宜的内容识别，等价于 1+2。
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 五官操作信息。
	FaceOperationInfo *AiSampleFaceOperation `json:"FaceOperationInfo,omitnil,omitempty" name:"FaceOperationInfo"`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil,omitempty" name:"TagOperationInfo"`
}

type ModifyPersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// 素材 ID。
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 名称，长度限制：128 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 素材应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于不适宜的内容识别，等价于 Review.Face。
	// 3. All：用于内容识别、不适宜的内容识别，等价于 1+2。
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 五官操作信息。
	FaceOperationInfo *AiSampleFaceOperation `json:"FaceOperationInfo,omitnil,omitempty" name:"FaceOperationInfo"`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil,omitempty" name:"TagOperationInfo"`
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
	delete(f, "SubAppId")
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
	Person *AiSamplePerson `json:"Person,omitnil,omitempty" name:"Person"`

	// 处理失败的五官信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitnil,omitempty" name:"FailFaceInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyQualityInspectTemplateRequestParams struct {
	// 模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 截帧间隔，单位为秒，最小值为 1。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`

	// 视频画面抖动重影检测的控制参数。
	JitterConfigure *JitterConfigureInfoForUpdate `json:"JitterConfigure,omitnil,omitempty" name:"JitterConfigure"`

	// 视频画面模糊检测的控制参数。
	BlurConfigure *BlurConfigureInfoForUpdate `json:"BlurConfigure,omitnil,omitempty" name:"BlurConfigure"`

	// 视频画面低光、过曝检测的控制参数。
	AbnormalLightingConfigure *AbnormalLightingConfigureInfoForUpdate `json:"AbnormalLightingConfigure,omitnil,omitempty" name:"AbnormalLightingConfigure"`

	// 视频画面花屏检测的控制参数。
	CrashScreenConfigure *CrashScreenConfigureInfoForUpdate `json:"CrashScreenConfigure,omitnil,omitempty" name:"CrashScreenConfigure"`

	// 视频画面黑边、白边、黑屏、白屏检测的控制参数。
	BlackWhiteEdgeConfigure *BlackWhiteEdgeConfigureInfoForUpdate `json:"BlackWhiteEdgeConfigure,omitnil,omitempty" name:"BlackWhiteEdgeConfigure"`

	// 视频画面噪点检测的控制参数。
	NoiseConfigure *NoiseConfigureInfoForUpdate `json:"NoiseConfigure,omitnil,omitempty" name:"NoiseConfigure"`

	// 视频画面马赛克检测的控制参数。
	MosaicConfigure *MosaicConfigureInfoForUpdate `json:"MosaicConfigure,omitnil,omitempty" name:"MosaicConfigure"`

	// 视频画面二维码检测的控制参数。
	QRCodeConfigure *QRCodeConfigureInfoForUpdate `json:"QRCodeConfigure,omitnil,omitempty" name:"QRCodeConfigure"`

	// 音频（静音、低音、爆音）检测的控制参数。
	VoiceConfigure *VoiceConfigureInfoForUpdate `json:"VoiceConfigure,omitnil,omitempty" name:"VoiceConfigure"`

	// 视频画面质量评价的控制参数。
	QualityEvaluationConfigure *QualityEvaluationConfigureInfoForUpdate `json:"QualityEvaluationConfigure,omitnil,omitempty" name:"QualityEvaluationConfigure"`
}

type ModifyQualityInspectTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 截帧间隔，单位为秒，最小值为 1。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`

	// 视频画面抖动重影检测的控制参数。
	JitterConfigure *JitterConfigureInfoForUpdate `json:"JitterConfigure,omitnil,omitempty" name:"JitterConfigure"`

	// 视频画面模糊检测的控制参数。
	BlurConfigure *BlurConfigureInfoForUpdate `json:"BlurConfigure,omitnil,omitempty" name:"BlurConfigure"`

	// 视频画面低光、过曝检测的控制参数。
	AbnormalLightingConfigure *AbnormalLightingConfigureInfoForUpdate `json:"AbnormalLightingConfigure,omitnil,omitempty" name:"AbnormalLightingConfigure"`

	// 视频画面花屏检测的控制参数。
	CrashScreenConfigure *CrashScreenConfigureInfoForUpdate `json:"CrashScreenConfigure,omitnil,omitempty" name:"CrashScreenConfigure"`

	// 视频画面黑边、白边、黑屏、白屏检测的控制参数。
	BlackWhiteEdgeConfigure *BlackWhiteEdgeConfigureInfoForUpdate `json:"BlackWhiteEdgeConfigure,omitnil,omitempty" name:"BlackWhiteEdgeConfigure"`

	// 视频画面噪点检测的控制参数。
	NoiseConfigure *NoiseConfigureInfoForUpdate `json:"NoiseConfigure,omitnil,omitempty" name:"NoiseConfigure"`

	// 视频画面马赛克检测的控制参数。
	MosaicConfigure *MosaicConfigureInfoForUpdate `json:"MosaicConfigure,omitnil,omitempty" name:"MosaicConfigure"`

	// 视频画面二维码检测的控制参数。
	QRCodeConfigure *QRCodeConfigureInfoForUpdate `json:"QRCodeConfigure,omitnil,omitempty" name:"QRCodeConfigure"`

	// 音频（静音、低音、爆音）检测的控制参数。
	VoiceConfigure *VoiceConfigureInfoForUpdate `json:"VoiceConfigure,omitnil,omitempty" name:"VoiceConfigure"`

	// 视频画面质量评价的控制参数。
	QualityEvaluationConfigure *QualityEvaluationConfigureInfoForUpdate `json:"QualityEvaluationConfigure,omitnil,omitempty" name:"QualityEvaluationConfigure"`
}

func (r *ModifyQualityInspectTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQualityInspectTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ScreenshotInterval")
	delete(f, "JitterConfigure")
	delete(f, "BlurConfigure")
	delete(f, "AbnormalLightingConfigure")
	delete(f, "CrashScreenConfigure")
	delete(f, "BlackWhiteEdgeConfigure")
	delete(f, "NoiseConfigure")
	delete(f, "MosaicConfigure")
	delete(f, "QRCodeConfigure")
	delete(f, "VoiceConfigure")
	delete(f, "QualityEvaluationConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQualityInspectTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQualityInspectTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQualityInspectTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQualityInspectTemplateResponseParams `json:"Response"`
}

func (r *ModifyQualityInspectTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQualityInspectTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRebuildMediaTemplateRequestParams struct {
	// 视频重生模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *string `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 视频重生模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 视频重生模板描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频重生视频控制信息。
	RebuildVideoInfo *RebuildVideoInfo `json:"RebuildVideoInfo,omitnil,omitempty" name:"RebuildVideoInfo"`

	// 视频重生音频控制信息。
	RebuildAudioInfo *RebuildAudioInfo `json:"RebuildAudioInfo,omitnil,omitempty" name:"RebuildAudioInfo"`

	// 输出目标视频控制信息。
	TargetVideoInfo *RebuildMediaTargetVideoStream `json:"TargetVideoInfo,omitnil,omitempty" name:"TargetVideoInfo"`

	// 输出目标音频控制信息。
	TargetAudioInfo *RebuildMediaTargetAudioStream `json:"TargetAudioInfo,omitnil,omitempty" name:"TargetAudioInfo"`

	// 输出文件封装格式，可选值：mp4、flv、hls。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

type ModifyRebuildMediaTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 视频重生模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *string `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 视频重生模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 视频重生模板描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频重生视频控制信息。
	RebuildVideoInfo *RebuildVideoInfo `json:"RebuildVideoInfo,omitnil,omitempty" name:"RebuildVideoInfo"`

	// 视频重生音频控制信息。
	RebuildAudioInfo *RebuildAudioInfo `json:"RebuildAudioInfo,omitnil,omitempty" name:"RebuildAudioInfo"`

	// 输出目标视频控制信息。
	TargetVideoInfo *RebuildMediaTargetVideoStream `json:"TargetVideoInfo,omitnil,omitempty" name:"TargetVideoInfo"`

	// 输出目标音频控制信息。
	TargetAudioInfo *RebuildMediaTargetAudioStream `json:"TargetAudioInfo,omitnil,omitempty" name:"TargetAudioInfo"`

	// 输出文件封装格式，可选值：mp4、flv、hls。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

func (r *ModifyRebuildMediaTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRebuildMediaTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RebuildVideoInfo")
	delete(f, "RebuildAudioInfo")
	delete(f, "TargetVideoInfo")
	delete(f, "TargetAudioInfo")
	delete(f, "Container")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRebuildMediaTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRebuildMediaTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRebuildMediaTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRebuildMediaTemplateResponseParams `json:"Response"`
}

func (r *ModifyRebuildMediaTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRebuildMediaTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReviewTemplateRequestParams struct {
	// 审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 需要返回的违规标签列表，可选值为：
	// <li>Porn：色情；</li>
	// <li>Terror：暴力；</li>
	// <li>Polity：不适宜的信息；</li>
	// <li>Illegal：违法；</li>
	// <li>Abuse：谩骂；</li>
	// <li>Ad：广告；</li>
	// <li>Moan：娇喘。</li>
	// 
	// 注意：不填表示不更新。
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ModifyReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 审核模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 需要返回的违规标签列表，可选值为：
	// <li>Porn：色情；</li>
	// <li>Terror：暴力；</li>
	// <li>Polity：不适宜的信息；</li>
	// <li>Illegal：违法；</li>
	// <li>Abuse：谩骂；</li>
	// <li>Ad：广告；</li>
	// <li>Moan：娇喘。</li>
	// 
	// 注意：不填表示不更新。
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`
}

func (r *ModifyReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReviewTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "Labels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReviewTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyReviewTemplateResponseParams `json:"Response"`
}

func (r *ModifyReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReviewTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoundPlayRequestParams struct {
	// 轮播播单唯一标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 启播时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 轮播列表。
	// <li>数组长度限制：100。</li>
	RoundPlaylist []*RoundPlayListItemInfo `json:"RoundPlaylist,omitnil,omitempty" name:"RoundPlaylist"`

	// 轮播播单名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 轮播播单描述信息，长度限制：256 个字符。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 播放状态，可选值：<li>Disabled：停止播放。</li><li>Enabled：启播时长到达后启动播放。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 播放模式，可选值：
	// <li>Loop：循环播放播单；</li>
	// <li>Linear：单次播放，播单播放完停止播放。</li>
	PlayBackMode *string `json:"PlayBackMode,omitnil,omitempty" name:"PlayBackMode"`
}

type ModifyRoundPlayRequest struct {
	*tchttp.BaseRequest
	
	// 轮播播单唯一标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 启播时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 轮播列表。
	// <li>数组长度限制：100。</li>
	RoundPlaylist []*RoundPlayListItemInfo `json:"RoundPlaylist,omitnil,omitempty" name:"RoundPlaylist"`

	// 轮播播单名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 轮播播单描述信息，长度限制：256 个字符。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 播放状态，可选值：<li>Disabled：停止播放。</li><li>Enabled：启播时长到达后启动播放。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 播放模式，可选值：
	// <li>Loop：循环播放播单；</li>
	// <li>Linear：单次播放，播单播放完停止播放。</li>
	PlayBackMode *string `json:"PlayBackMode,omitnil,omitempty" name:"PlayBackMode"`
}

func (r *ModifyRoundPlayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoundPlayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoundPlayId")
	delete(f, "SubAppId")
	delete(f, "StartTime")
	delete(f, "RoundPlaylist")
	delete(f, "Name")
	delete(f, "Desc")
	delete(f, "Status")
	delete(f, "PlayBackMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoundPlayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoundPlayResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRoundPlayResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoundPlayResponseParams `json:"Response"`
}

func (r *ModifyRoundPlayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoundPlayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySampleSnapshotTemplateRequestParams struct {
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// 图片格式，取值为 jpg 和 png。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type ModifySampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// 图片格式，取值为 jpg 和 png。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
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
	delete(f, "SubAppId")
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifySnapshotByTimeOffsetTemplateRequestParams struct {
	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式，取值可以为 jpg 和 png。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type ModifySnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式，取值可以为 jpg 和 png。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
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
	delete(f, "SubAppId")
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifySubAppIdInfoRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 应用名称，长度限制：40个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用简介，长度限制： 300个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifySubAppIdInfoRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 应用名称，长度限制：40个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用简介，长度限制： 300个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifySubAppIdInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubAppIdInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubAppIdInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubAppIdInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySubAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubAppIdInfoResponseParams `json:"Response"`
}

func (r *ModifySubAppIdInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubAppIdInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubAppIdStatusRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 应用状态，取值范围：
	// <li>On：启用。</li>
	// <li>Off：停用。</li>
	// <li>Destroyed：销毁。</li>
	// 当前状态如果是 Destoying ，不能进行启用操作，需要等待销毁完成后才能重新启用。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifySubAppIdStatusRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 应用状态，取值范围：
	// <li>On：启用。</li>
	// <li>Off：停用。</li>
	// <li>Destroyed：销毁。</li>
	// 当前状态如果是 Destoying ，不能进行启用操作，需要等待销毁完成后才能重新启用。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifySubAppIdStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubAppIdStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubAppIdStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubAppIdStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySubAppIdStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubAppIdStatusResponseParams `json:"Response"`
}

func (r *ModifySubAppIdStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubAppIdStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySuperPlayerConfigRequestParams struct {
	// 播放器配置名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 播放的音视频类型，可选值：
	// <li>AdaptiveDynamicStream：自适应码流输出；</li>
	// <li>Transcode：转码输出；</li>
	// <li>Original：原始音视频。</li>
	AudioVideoType *string `json:"AudioVideoType,omitnil,omitempty" name:"AudioVideoType"`

	// 播放 DRM 保护的自适应码流开关：
	// <li>ON：开启，表示仅播放 DRM  保护的自适应码流输出；</li>
	// <li>OFF：关闭，表示播放未加密的自适应码流输出。</li>
	DrmSwitch *string `json:"DrmSwitch,omitnil,omitempty" name:"DrmSwitch"`

	// 允许输出的未加密的自适应码流模板 ID。
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitnil,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// 允许输出的 DRM 自适应码流模板内容。
	DrmStreamingsInfo *DrmStreamingsInfoForUpdate `json:"DrmStreamingsInfo,omitnil,omitempty" name:"DrmStreamingsInfo"`

	// 允许输出的转码模板 ID。
	TranscodeDefinition *uint64 `json:"TranscodeDefinition,omitnil,omitempty" name:"TranscodeDefinition"`

	// 允许输出的雪碧图模板 ID。
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitnil,omitempty" name:"ImageSpriteDefinition"`

	// 播放器对于不同分辨率的子流展示名字。
	ResolutionNames []*ResolutionNameInfo `json:"ResolutionNames,omitnil,omitempty" name:"ResolutionNames"`

	// 播放时使用的域名。填 Default 表示使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 播放时使用的 Scheme。取值范围：
	// <li>Default：使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的 Scheme；</li>
	// <li>HTTP；</li>
	// <li>HTTPS。</li>
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ModifySuperPlayerConfigRequest struct {
	*tchttp.BaseRequest
	
	// 播放器配置名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 播放的音视频类型，可选值：
	// <li>AdaptiveDynamicStream：自适应码流输出；</li>
	// <li>Transcode：转码输出；</li>
	// <li>Original：原始音视频。</li>
	AudioVideoType *string `json:"AudioVideoType,omitnil,omitempty" name:"AudioVideoType"`

	// 播放 DRM 保护的自适应码流开关：
	// <li>ON：开启，表示仅播放 DRM  保护的自适应码流输出；</li>
	// <li>OFF：关闭，表示播放未加密的自适应码流输出。</li>
	DrmSwitch *string `json:"DrmSwitch,omitnil,omitempty" name:"DrmSwitch"`

	// 允许输出的未加密的自适应码流模板 ID。
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitnil,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// 允许输出的 DRM 自适应码流模板内容。
	DrmStreamingsInfo *DrmStreamingsInfoForUpdate `json:"DrmStreamingsInfo,omitnil,omitempty" name:"DrmStreamingsInfo"`

	// 允许输出的转码模板 ID。
	TranscodeDefinition *uint64 `json:"TranscodeDefinition,omitnil,omitempty" name:"TranscodeDefinition"`

	// 允许输出的雪碧图模板 ID。
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitnil,omitempty" name:"ImageSpriteDefinition"`

	// 播放器对于不同分辨率的子流展示名字。
	ResolutionNames []*ResolutionNameInfo `json:"ResolutionNames,omitnil,omitempty" name:"ResolutionNames"`

	// 播放时使用的域名。填 Default 表示使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 播放时使用的 Scheme。取值范围：
	// <li>Default：使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的 Scheme；</li>
	// <li>HTTP；</li>
	// <li>HTTPS。</li>
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *ModifySuperPlayerConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySuperPlayerConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	delete(f, "AudioVideoType")
	delete(f, "DrmSwitch")
	delete(f, "AdaptiveDynamicStreamingDefinition")
	delete(f, "DrmStreamingsInfo")
	delete(f, "TranscodeDefinition")
	delete(f, "ImageSpriteDefinition")
	delete(f, "ResolutionNames")
	delete(f, "Domain")
	delete(f, "Scheme")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySuperPlayerConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySuperPlayerConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySuperPlayerConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifySuperPlayerConfigResponseParams `json:"Response"`
}

func (r *ModifySuperPlayerConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySuperPlayerConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTranscodeTemplateRequestParams struct {
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a、wav。其中，mp3、flac、ogg、m4a、wav 为纯音频文件。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// 视频流配置参数。
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// 音频流配置参数。
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// 极速高清转码参数。
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// 切片类型，当 Container 为 hls 时有效，可选值：
	// <li>ts：ts 切片；</li>
	// <li>fmp4：fmp4 切片。</li>
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
}

type ModifyTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a、wav。其中，mp3、flac、ogg、m4a、wav 为纯音频文件。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// 视频流配置参数。
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// 音频流配置参数。
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// 极速高清转码参数。
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// 切片类型，当 Container 为 hls 时有效，可选值：
	// <li>ts：ts 切片；</li>
	// <li>fmp4：fmp4 切片。</li>
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
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
	delete(f, "SubAppId")
	delete(f, "Container")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	delete(f, "VideoTemplate")
	delete(f, "AudioTemplate")
	delete(f, "TEHDConfig")
	delete(f, "SegmentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTranscodeTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyVodDomainAccelerateConfigRequestParams struct {
	// 需要设置加速配置的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 区域，可选值：
	// <li>Chinese Mainland：中国境内（不包含港澳台）。</li>
	// <li>Outside Chinese Mainland: 中国境外。</li>
	// <li>Global: 全球范围。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 开启或者关闭所选区域的域名加速，可选值：
	// <li>Enabled: 开启。</li>
	// <li>Disabled：关闭。</li>
	// 开启中国境内加速的域名，需要先[备案域名](/document/product/243/18905)。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type ModifyVodDomainAccelerateConfigRequest struct {
	*tchttp.BaseRequest
	
	// 需要设置加速配置的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 区域，可选值：
	// <li>Chinese Mainland：中国境内（不包含港澳台）。</li>
	// <li>Outside Chinese Mainland: 中国境外。</li>
	// <li>Global: 全球范围。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 开启或者关闭所选区域的域名加速，可选值：
	// <li>Enabled: 开启。</li>
	// <li>Disabled：关闭。</li>
	// 开启中国境内加速的域名，需要先[备案域名](/document/product/243/18905)。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *ModifyVodDomainAccelerateConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVodDomainAccelerateConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Area")
	delete(f, "Status")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVodDomainAccelerateConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVodDomainAccelerateConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVodDomainAccelerateConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVodDomainAccelerateConfigResponseParams `json:"Response"`
}

func (r *ModifyVodDomainAccelerateConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVodDomainAccelerateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVodDomainConfigRequestParams struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// [Referer 防盗链](/document/product/266/14046)规则。
	RefererAuthPolicy *RefererAuthPolicy `json:"RefererAuthPolicy,omitnil,omitempty" name:"RefererAuthPolicy"`

	// [Key 防盗链](/document/product/266/14047)规则。
	UrlSignatureAuthPolicy *UrlSignatureAuthPolicy `json:"UrlSignatureAuthPolicy,omitnil,omitempty" name:"UrlSignatureAuthPolicy"`

	// QUIC 配置。
	QUICConfig *DomainQUICConfig `json:"QUICConfig,omitnil,omitempty" name:"QUICConfig"`
}

type ModifyVodDomainConfigRequest struct {
	*tchttp.BaseRequest
	
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// [Referer 防盗链](/document/product/266/14046)规则。
	RefererAuthPolicy *RefererAuthPolicy `json:"RefererAuthPolicy,omitnil,omitempty" name:"RefererAuthPolicy"`

	// [Key 防盗链](/document/product/266/14047)规则。
	UrlSignatureAuthPolicy *UrlSignatureAuthPolicy `json:"UrlSignatureAuthPolicy,omitnil,omitempty" name:"UrlSignatureAuthPolicy"`

	// QUIC 配置。
	QUICConfig *DomainQUICConfig `json:"QUICConfig,omitnil,omitempty" name:"QUICConfig"`
}

func (r *ModifyVodDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVodDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SubAppId")
	delete(f, "RefererAuthPolicy")
	delete(f, "UrlSignatureAuthPolicy")
	delete(f, "QUICConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVodDomainConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVodDomainConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVodDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVodDomainConfigResponseParams `json:"Response"`
}

func (r *ModifyVodDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVodDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWatermarkTemplateRequestParams struct {
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// 图片水印模板，该字段仅对图片水印模板有效。
	ImageTemplate *ImageWatermarkInputForUpdate `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// 文字水印模板，该字段仅对文字水印模板有效。
	TextTemplate *TextWatermarkTemplateInputForUpdate `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG 水印模板，该字段仅对 SVG 水印模板有效。
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`
}

type ModifyWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// 图片水印模板，该字段仅对图片水印模板有效。
	ImageTemplate *ImageWatermarkInputForUpdate `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// 文字水印模板，该字段仅对文字水印模板有效。
	TextTemplate *TextWatermarkTemplateInputForUpdate `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG 水印模板，该字段仅对 SVG 水印模板有效。
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`
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
	delete(f, "SubAppId")
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
	// 图片水印地址，仅当 ImageTemplate.ImageContent 非空，该字段有值。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// <b>关键词应用场景，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜的内容识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜的内容识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜的内容识别，等价于 3+4；
	// 7. All：包含以上全部，等价于 1+2+3+4。
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil,omitempty" name:"TagOperationInfo"`
}

type ModifyWordSampleRequest struct {
	*tchttp.BaseRequest
	
	// 关键词，长度限制：128 个字符。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// <b>关键词应用场景，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过音频识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行不适宜的内容识别；
	// 4. Review.Asr：通过音频识别技术，进行不适宜的内容识别；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、音频识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、音频识别技术，进行不适宜的内容识别，等价于 3+4；
	// 7. All：包含以上全部，等价于 1+2+3+4。
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil,omitempty" name:"TagOperationInfo"`
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
	delete(f, "SubAppId")
	delete(f, "Usages")
	delete(f, "TagOperationInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWordSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWordSampleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type MosaicConfigureInfo struct {
	// 视频画面马赛克检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type MosaicConfigureInfoForUpdate struct {
	// 视频画面马赛克检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type MosaicInput struct {
	// 原点位置，目前仅支持：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，马赛克原点为图片或文字的左上角。</li>
	// 默认值：TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// 马赛克原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示马赛克 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示马赛克 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 马赛克原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示马赛克 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示马赛克 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// 马赛克的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示马赛克 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示马赛克 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 默认值：10%。
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 马赛克的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示马赛克 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示马赛克 Height 单位为像素，如 100px 表示 Height 为 100 像素。</li>
	// 默认值：10%。
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 马赛克的起始时间偏移，单位：秒。不填或填0，表示马赛克从画面出现时开始显现。
	// <li>不填或填0，表示马赛克从画面开始就出现；</li>
	// <li>当数值大于0时（假设为 n），表示马赛克从画面开始的第 n 秒出现；</li>
	// <li>当数值小于0时（假设为 -n），表示马赛克从离画面结束 n 秒前开始出现。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 马赛克的结束时间偏移，单位：秒。
	// <li>不填或填0，表示马赛克持续到画面结束；</li>
	// <li>当数值大于0时（假设为 n），表示马赛克持续到第 n 秒时消失；</li>
	// <li>当数值小于0时（假设为 -n），表示马赛克持续到离画面结束 n 秒前消失。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type NoiseConfigureInfo struct {
	// 视频画面噪点检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type NoiseConfigureInfoForUpdate struct {
	// 视频画面噪点检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ObjectConfigureInfo struct {
	// 物体识别任务开关，可选值：
	// <li>ON：开启智能物体识别任务；</li>
	// <li>OFF：关闭智能物体识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 物体库选择，可选值：
	// <li>Default：使用默认物体库；</li>
	// <li>UserDefine：使用用户自定义物体库。</li>
	// <li>All：同时使用默认物体库和用户自定义物体库。</li>
	// 默认值： All，同时使用默认物体库和用户自定义物体库。
	ObjectLibrary *string `json:"ObjectLibrary,omitnil,omitempty" name:"ObjectLibrary"`
}

type ObjectConfigureInfoForUpdate struct {
	// 物体识别任务开关，可选值：
	// <li>ON：开启智能物体识别任务；</li>
	// <li>OFF：关闭智能物体识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 物体库选择，可选值：
	// <li>Default：使用默认物体库；</li>
	// <li>UserDefine：使用用户自定义物体库。</li>
	// <li>All：同时使用默认物体库和用户自定义物体库。</li>
	ObjectLibrary *string `json:"ObjectLibrary,omitnil,omitempty" name:"ObjectLibrary"`
}

type OcrFullTextConfigureInfo struct {
	// 文本全文识别任务开关，可选值：
	// <li>ON：开启智能文本全文识别任务；</li>
	// <li>OFF：关闭智能文本全文识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type OcrFullTextConfigureInfoForUpdate struct {
	// 文本全文识别任务开关，可选值：
	// <li>ON：开启智能文本全文识别任务；</li>
	// <li>OFF：关闭智能文本全文识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type OcrWordsConfigureInfo struct {
	// 文本关键词识别任务开关，可选值：
	// <li>ON：开启文本关键词识别任务；</li>
	// <li>OFF：关闭文本关键词识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`
}

type OcrWordsConfigureInfoForUpdate struct {
	// 文本关键词识别任务开关，可选值：
	// <li>ON：开启文本关键词识别任务；</li>
	// <li>OFF：关闭文本关键词识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`
}

type OutputAudioStream struct {
	// 音频流的编码格式，可选值：
	// <li>libfdk_aac：适合 mp4 文件。</li>
	// 默认值：libfdk_aac。
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 音频流的采样率，可选值：
	// <li>16000</li>
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	// 默认值：16000。
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频声道数，可选值：
	// <li>1：单声道 。</li>
	// <li>2：双声道</li>
	// 默认值：2。
	AudioChannel *int64 `json:"AudioChannel,omitnil,omitempty" name:"AudioChannel"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。
	// 当取值为 0，将自动设置音频码率。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`
}

type OutputVideoStream struct {
	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码 </li>
	// 默认值：libx264。
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 视频帧率，取值范围：[0, 60]，单位：Hz。
	// 默认值：0，表示和第一个视频轨的第一个视频片段的视频帧率一致。
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`
}

// Predefined struct for user
type ParseStreamingManifestRequestParams struct {
	// 待解析的索引文件内容。
	MediaManifestContent *string `json:"MediaManifestContent,omitnil,omitempty" name:"MediaManifestContent"`

	// 视频索引文件格式。默认 m3u8 格式。
	// <li>m3u8</li>
	// <li>mpd</li>
	ManifestType *string `json:"ManifestType,omitnil,omitempty" name:"ManifestType"`
}

type ParseStreamingManifestRequest struct {
	*tchttp.BaseRequest
	
	// 待解析的索引文件内容。
	MediaManifestContent *string `json:"MediaManifestContent,omitnil,omitempty" name:"MediaManifestContent"`

	// 视频索引文件格式。默认 m3u8 格式。
	// <li>m3u8</li>
	// <li>mpd</li>
	ManifestType *string `json:"ManifestType,omitnil,omitempty" name:"ManifestType"`
}

func (r *ParseStreamingManifestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseStreamingManifestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaManifestContent")
	delete(f, "ManifestType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseStreamingManifestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseStreamingManifestResponseParams struct {
	// 分片文件列表。
	MediaSegmentSet []*string `json:"MediaSegmentSet,omitnil,omitempty" name:"MediaSegmentSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ParseStreamingManifestResponse struct {
	*tchttp.BaseResponse
	Response *ParseStreamingManifestResponseParams `json:"Response"`
}

func (r *ParseStreamingManifestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseStreamingManifestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PersistenceCompleteTask struct {
	// 固化生成的媒体 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 剪辑固化的来源，有以下三种。
	// <li>SimpleHlsClip：来自简单 HLS 剪辑；</li>
	// <li>FastEditMedia：来自快速媒体编辑；</li>
	// <li>LiveRealTimeClip:来自直播即时剪辑。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersistenceSource *string `json:"PersistenceSource,omitnil,omitempty" name:"PersistenceSource"`
}

type PlayStatFileInfo struct {
	// 播放统计数据所属日期，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 播放统计文件的 URL 地址。播放统计文件内容为：
	// <li> date：播放日期。</li>
	// <li> file_id：视频文件 ID。</li>
	// <li> ip_count：去重后的客户端 IP 数。</li>
	// <li> flux：播放流量，单位：字节。</li>
	// <li> play_times：总的播放次数。</li>
	// <li> pc_play_times：PC 端播放次数。</li>
	// <li> mobile_play_times：移动端播放次数。</li>
	// <li> iphone_play_times：iPhone 端播放次数。</li>
	// <li> android_play_times：Android 端播放次数。</li>
	// <li> host_name	域名。</li>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type PlayStatInfo struct {
	// 数据所在时间区间的开始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。如：当时间粒度为天，2018-12-01T00:00:00+08:00，表示2018年12月1日（含）到2018年12月2日（不含）区间。
	// <li>表示小时级别数据时，2019-08-22T00:00:00+08:00表示2019-08-22日0点到1点的统计数据。</li>
	// <li>表示天级别数据时，2019-08-22T00:00:00+08:00表示2019-08-22日的统计数据。</li>
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 媒体文件ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 播放次数。
	PlayTimes *uint64 `json:"PlayTimes,omitnil,omitempty" name:"PlayTimes"`

	// 播放流量，单位：字节。
	Traffic *uint64 `json:"Traffic,omitnil,omitempty" name:"Traffic"`
}

type PlayerConfig struct {
	// 播放器配置名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 播放器配置类型，取值范围：
	// <li>Preset：系统预置配置；</li>
	// <li>Custom：用户自定义配置。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 播放的音视频类型，可选值有：
	// <li>AdaptiveDynamicStream：自适应码流输出；</li>
	// <li>Transcode：转码输出；</li>
	// <li>Original：原始音视频。</li>
	AudioVideoType *string `json:"AudioVideoType,omitnil,omitempty" name:"AudioVideoType"`

	// 播放 DRM 保护的自适应码流开关：
	// <li>ON：开启，表示仅播放 DRM  保护的自适应码流输出；</li>
	// <li>OFF：关闭，表示播放未加密的自适应码流输出。</li>
	DrmSwitch *string `json:"DrmSwitch,omitnil,omitempty" name:"DrmSwitch"`

	// 允许输出的未加密的自适应码流模板 ID。
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitnil,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// 允许输出的 DRM 自适应码流模板内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrmStreamingsInfo *DrmStreamingsInfo `json:"DrmStreamingsInfo,omitnil,omitempty" name:"DrmStreamingsInfo"`

	// 允许输出的转码模板 ID。
	TranscodeDefinition *uint64 `json:"TranscodeDefinition,omitnil,omitempty" name:"TranscodeDefinition"`

	// 允许输出的雪碧图模板 ID。
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitnil,omitempty" name:"ImageSpriteDefinition"`

	// 播放器对不于不同分辨率的子流展示名字。
	ResolutionNameSet []*ResolutionNameInfo `json:"ResolutionNameSet,omitnil,omitempty" name:"ResolutionNameSet"`

	// 播放器配置创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 播放器配置最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 播放时使用的域名。值为 Default，表示使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 播放时使用的 Scheme。取值范围：
	// <li>Default：使用[默认分发配置](https://cloud.tencent.com/document/product/266/33373)中的 Scheme；</li>
	// <li>HTTP；</li>
	// <li>HTTPS。</li>
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type PoliticalAsrReviewTemplateInfo struct {
	// 语音鉴别涉及令人不适宜的信息的任务开关，可选值：
	// <li>ON：开启语音鉴别涉及令人不适宜的信息的任务；</li>
	// <li>OFF：关闭语音鉴别的涉及令人不适宜的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定需人工复核是否违规的分数阈值，当音视频审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`

	// 判定涉嫌违规的分数阈值，当音视频审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`
}

type PoliticalAsrReviewTemplateInfoForUpdate struct {
	// 语音鉴别涉及令人不适宜的信息的任务开关，可选值：
	// <li>ON：开启语音鉴别涉及令人不适宜的信息的任务；</li>
	// <li>OFF：关闭语音鉴别涉及令人不适宜的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当音视频审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当音视频审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PoliticalConfigureInfo struct {
	// 画面鉴别涉及令人不适宜的信息的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImgReviewInfo *PoliticalImgReviewTemplateInfo `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// 语音鉴别涉及令人不适宜的信息的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrReviewInfo *PoliticalAsrReviewTemplateInfo `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// 文本鉴别涉及令人不适宜的信息的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrReviewInfo *PoliticalOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type PoliticalConfigureInfoForUpdate struct {
	// 画面鉴别涉及令人不适宜的信息的控制参数。
	ImgReviewInfo *PoliticalImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// 语音鉴别涉及令人不适宜的信息的控制参数。
	AsrReviewInfo *PoliticalAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// 文本鉴别涉及令人不适宜的信息的控制参数。
	OcrReviewInfo *PoliticalOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type PoliticalImageResult struct {
	// 鉴别涉及令人不适宜信息的评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 鉴别涉及令人不适宜信息的结果建议，取值范围：
	// <li>pass；</li>
	// <li>review；</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 涉及令人不适宜的信息、违规图标名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 涉及令人不适宜的信息、违规图标出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type PoliticalImgReviewTemplateInfo struct {
	// 画面鉴别涉及令人不适宜的信息的任务开关，可选值：
	// <li>ON：开启画面鉴别涉及令人不适宜的信息的任务；</li>
	// <li>OFF：关闭画面鉴别涉及令人不适宜的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 画面鉴别涉及令人不适宜的信息的过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>violation_photo：违规图标；</li>
	// <li>politician：相关人物；</li>
	// <li>entertainment：娱乐人物；</li>
	// <li>sport：体育人物；</li>
	// <li>entrepreneur：商业人物；</li>
	// <li>scholar：教育学者；</li>
	// <li>celebrity：知名人物；</li>
	// <li>military：相关人物。</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 97 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 95 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PoliticalImgReviewTemplateInfoForUpdate struct {
	// 画面鉴别涉及令人不适宜的信息的任务开关，可选值：
	// <li>ON：开启画面鉴别涉及令人不适宜的信息的任务；</li>
	// <li>OFF：关闭画面鉴别涉及令人不适宜的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 画面鉴别涉及令人不适宜的信息的过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>violation_photo：违规图标；</li>
	// <li>politician：相关人物；</li>
	// <li>entertainment：娱乐人物；</li>
	// <li>sport：体育人物；</li>
	// <li>entrepreneur：商业人物；</li>
	// <li>scholar：教育学者；</li>
	// <li>celebrity：知名人物；</li>
	// <li>military：相关人物。</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfo struct {
	// 文本鉴别涉及令人不适宜的信息的任务开关，可选值：
	// <li>ON：开启文本鉴别涉及令人不适宜的信息的任务；</li>
	// <li>OFF：关闭文本鉴别涉及令人不适宜的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfoForUpdate struct {
	// 文本鉴别涉及令人不适宜的信息的任务开关，可选值：
	// <li>ON：开启文本鉴别涉及令人不适宜的信息的任务；</li>
	// <li>OFF：关闭文本鉴别涉及令人不适宜的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfo struct {
	// 语音鉴别涉及令人反感的信息的任务开关，可选值：
	// <li>ON：开启语音鉴别涉及令人反感的信息的任务；</li>
	// <li>OFF：关闭语音鉴别涉及令人反感的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfoForUpdate struct {
	// 语音鉴别涉及令人反感的信息的任务开关，可选值：
	// <li>ON：开启语音鉴别涉及令人反感的信息的任务；</li>
	// <li>OFF：关闭语音鉴别涉及令人反感的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornConfigureInfo struct {
	// 画面鉴别涉及令人反感的信息的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImgReviewInfo *PornImgReviewTemplateInfo `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// 语音鉴别涉及令人反感的信息的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrReviewInfo *PornAsrReviewTemplateInfo `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// 文本鉴别涉及令人反感的信息的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrReviewInfo *PornOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type PornConfigureInfoForUpdate struct {
	// 画面鉴别涉及令人反感的信息的控制参数。
	ImgReviewInfo *PornImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// 语音鉴别涉及令人反感的信息的控制参数。
	AsrReviewInfo *PornAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// 文本鉴别涉及令人反感的信息的控制参数。
	OcrReviewInfo *PornOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type PornImageResult struct {
	// 鉴别涉及令人反感的信息的评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 鉴别涉及令人反感的信息的结果建议，取值范围：
	// <li>pass；</li>
	// <li>review；</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 鉴别涉及令人反感的信息的结果标签，取值范围：
	// <li>porn：色情；</li>
	// <li>sexy：性感；</li>
	// <li>vulgar：低俗；</li>
	// <li>intimacy：亲密行为。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`
}

type PornImgReviewTemplateInfo struct {
	// 画面鉴别涉及令人反感的信息的任务开关，可选值：
	// <li>ON：开启画面鉴别涉及令人反感的信息的任务；</li>
	// <li>OFF：关闭画面鉴别涉及令人反感的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 画面鉴别涉及令人反感的信息的过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>porn：色情；</li>
	// <li>vulgar：低俗；</li>
	// <li>intimacy：亲密行为；</li>
	// <li>sexy：性感。</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 90 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 0 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornImgReviewTemplateInfoForUpdate struct {
	// 画面鉴别涉及令人反感的信息的任务开关，可选值：
	// <li>ON：开启画面鉴别涉及令人反感的信息的任务；</li>
	// <li>OFF：关闭画面鉴别涉及令人反感的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 画面鉴别涉及令人反感的信息的过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>porn：色情；</li>
	// <li>vulgar：低俗；</li>
	// <li>intimacy：亲密行为；</li>
	// <li>sexy：性感。</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfo struct {
	// 文本鉴别涉及令人反感的信息的任务开关，可选值：
	// <li>ON：开启文本鉴别涉及令人反感的信息的任务；</li>
	// <li>OFF：关闭文本鉴别涉及令人反感的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfoForUpdate struct {
	// 文本鉴别涉及令人反感的信息的任务开关，可选值：
	// <li>ON：开启文本鉴别涉及令人反感的信息的任务；</li>
	// <li>OFF：关闭文本鉴别涉及令人反感的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type ProcedureReviewAudioVideoTaskInput struct {
	// 审核模板。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 审核的内容，可选值：
	// <li>Media：原始音视频；</li>
	// <li>Cover：封面。</li>
	// 不填或填空数组时，默认为审核 Media。
	ReviewContents []*string `json:"ReviewContents,omitnil,omitempty" name:"ReviewContents"`
}

type ProcedureTask struct {
	// 音视频处理任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 已弃用，请使用各个具体任务的 ErrCode。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 已弃用，请使用各个具体任务的 Message。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 媒体文件 ID
	// <li>若流程由 [ProcessMedia](https://cloud.tencent.com/document/product/266/33427) 发起，该字段表示 [MediaInfo](https://cloud.tencent.com/document/product/266/31773#MediaInfo) 的 FileId；</li>
	// <li>若流程由 [ProcessMediaByUrl](https://cloud.tencent.com/document/product/266/33426) 发起，该字段表示 [MediaInputInfo](https://cloud.tencent.com/document/product/266/31773#MediaInputInfo) 的 Id。</li>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 媒体文件名称
	// <li>若流程由 [ProcessMedia](https://cloud.tencent.com/document/product/266/33427) 发起，该字段表示 [MediaInfo](https://cloud.tencent.com/document/product/266/31773#MediaInfo) 的 BasicInfo.Name；</li>
	// <li>若流程由 [ProcessMediaByUrl](https://cloud.tencent.com/document/product/266/33426) 发起，该字段表示 [MediaInputInfo](https://cloud.tencent.com/document/product/266/31773#MediaInputInfo) 的 Name。</li>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 媒体文件地址
	// <li>若流程由 [ProcessMedia](https://cloud.tencent.com/document/product/266/33427) 发起，该字段表示 [MediaInfo](https://cloud.tencent.com/document/product/266/31773#MediaInfo) 的 BasicInfo.MediaUrl；</li>
	// <li>若流程由 [ProcessMediaByUrl](https://cloud.tencent.com/document/product/266/33426) 发起，该字段表示 [MediaInputInfo](https://cloud.tencent.com/document/product/266/31773#MediaInputInfo) 的 Url。</li>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 原始音视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 音视频处理任务的执行状态与结果。
	MediaProcessResultSet []*MediaProcessTaskResult `json:"MediaProcessResultSet,omitnil,omitempty" name:"MediaProcessResultSet"`

	// 音视频审核任务的执行状态与结果。
	AiContentReviewResultSet []*AiContentReviewResult `json:"AiContentReviewResultSet,omitnil,omitempty" name:"AiContentReviewResultSet"`

	// 音视频内容分析任务的执行状态与结果。
	AiAnalysisResultSet []*AiAnalysisResult `json:"AiAnalysisResultSet,omitnil,omitempty" name:"AiAnalysisResultSet"`

	// 音视频内容识别任务的执行状态与结果。
	AiRecognitionResultSet []*AiRecognitionResult `json:"AiRecognitionResultSet,omitnil,omitempty" name:"AiRecognitionResultSet"`

	// 任务流的优先级，取值范围为 [-10, 10]。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式。
	// <li>Finish：只有当任务流全部执行完毕时，才发起一次事件通知；</li>
	// <li>Change：只要任务流中每个子任务的状态发生变化，都进行事件通知；</li>
	// <li>None：不接受该任务流回调。</li>
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 操作者。取值范围：
	// <li>System: 表示系统触发。</li>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 操作类型。取值范围：
	// <li>TSC: 表示使用极速高清进行智能降码。</li>
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type ProcedureTemplate struct {
	// 任务流名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务流模板类型，取值范围：
	// <li>Preset：系统预置任务流模板；</li>
	// <li>Custom：用户自定义任务流模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频处理类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// AI 智能审核类型任务参数 \*。
	// <font color=red>\*：该参数用于发起旧版审核，不建议使用。推荐使用 ReviewAudioVideoTask 参数发起审核。</font> 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// AI 智能内容分析类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// AI 内容识别类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// 微信小程序发布任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniProgramPublishTask *WechatMiniProgramPublishTaskInput `json:"MiniProgramPublishTask,omitnil,omitempty" name:"MiniProgramPublishTask"`

	// 音视频审核类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewAudioVideoTask *ProcedureReviewAudioVideoTaskInput `json:"ReviewAudioVideoTask,omitnil,omitempty" name:"ReviewAudioVideoTask"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ProcessImageRequestParams struct {
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符。本接口要求媒体文件必须是图片格式。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 操作类型。现在仅支持填 ContentReview，表示内容智能识别。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 图片内容智能识别参数，当 Operation 为 ContentReview 时该字段有效。
	ContentReviewInput *ImageContentReviewInput `json:"ContentReviewInput,omitnil,omitempty" name:"ContentReviewInput"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type ProcessImageRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符。本接口要求媒体文件必须是图片格式。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 操作类型。现在仅支持填 ContentReview，表示内容智能识别。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 图片内容智能识别参数，当 Operation 为 ContentReview 时该字段有效。
	ContentReviewInput *ImageContentReviewInput `json:"ContentReviewInput,omitnil,omitempty" name:"ContentReviewInput"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *ProcessImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "Operation")
	delete(f, "ContentReviewInput")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessImageResponseParams struct {
	// 图片内容智能识别任务结果。
	ContentReviewResultSet []*ContentReviewResult `json:"ContentReviewResultSet,omitnil,omitempty" name:"ContentReviewResultSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProcessImageResponse struct {
	*tchttp.BaseResponse
	Response *ProcessImageResponseParams `json:"Response"`
}

func (r *ProcessImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaByProcedureRequestParams struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// [任务流模板](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF)名字。
	ProcedureName *string `json:"ProcedureName,omitnil,omitempty" name:"ProcedureName"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 任务流的优先级，数值越大优先级越高，取值范围是-10到10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果 7 天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type ProcessMediaByProcedureRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// [任务流模板](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF)名字。
	ProcedureName *string `json:"ProcedureName,omitnil,omitempty" name:"ProcedureName"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 任务流的优先级，数值越大优先级越高，取值范围是-10到10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果 7 天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *ProcessMediaByProcedureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaByProcedureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "ProcedureName")
	delete(f, "SubAppId")
	delete(f, "TasksPriority")
	delete(f, "TasksNotifyMode")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaByProcedureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaByProcedureResponseParams struct {
	// 任务类型为 Procedure 的任务 ID，当入参 ProcedureName 对应的任务流模板指定了 MediaProcessTask、AiAnalysisTask、AiRecognitionTask 中的一个或多个时发起该任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型为 ReviewAudioVideo 的任务 ID，当入参 ProcedureName 对应的任务流模板指定了 ReviewAudioVideoTask 时，发起该任务。
	ReviewAudioVideoTaskId *string `json:"ReviewAudioVideoTaskId,omitnil,omitempty" name:"ReviewAudioVideoTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProcessMediaByProcedureResponse struct {
	*tchttp.BaseResponse
	Response *ProcessMediaByProcedureResponseParams `json:"Response"`
}

func (r *ProcessMediaByProcedureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaByProcedureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaByUrlRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// API 已经<font color='red'>不再维护</font>。推荐使用的替代 API 请参考接口描述。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 输出文件 COS 路径信息。
	OutputInfo *MediaOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type ProcessMediaByUrlRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// API 已经<font color='red'>不再维护</font>。推荐使用的替代 API 请参考接口描述。
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 输出文件 COS 路径信息。
	OutputInfo *MediaOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *ProcessMediaByUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaByUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "InputInfo")
	delete(f, "OutputInfo")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "TasksPriority")
	delete(f, "TasksNotifyMode")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaByUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaByUrlResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProcessMediaByUrlResponse struct {
	*tchttp.BaseResponse
	Response *ProcessMediaByUrlResponseParams `json:"Response"`
}

func (r *ProcessMediaByUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaByUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaRequestParams struct {
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符，在上传成功后由云点播后台分配。可以在 [视频上传完成事件通知](/document/product/266/7830) 或 [云点播控制台](https://console.cloud.tencent.com/vod/media) 获取该字段。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// 音视频内容审核类型任务参数 \*。
	// <font color=red>\* 不建议使用</font>，推荐使用 [音视频审核(ReviewAudioVideo)](https://cloud.tencent.com/document/api/266/80283) 或 [图片审核(ReviewImage)](https://cloud.tencent.com/document/api/266/73217)。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// 音视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// 音视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type ProcessMediaRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符，在上传成功后由云点播后台分配。可以在 [视频上传完成事件通知](/document/product/266/7830) 或 [云点播控制台](https://console.cloud.tencent.com/vod/media) 获取该字段。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// 音视频内容审核类型任务参数 \*。
	// <font color=red>\* 不建议使用</font>，推荐使用 [音视频审核(ReviewAudioVideo)](https://cloud.tencent.com/document/api/266/80283) 或 [图片审核(ReviewImage)](https://cloud.tencent.com/document/api/266/73217)。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// 音视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// 音视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
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
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "TasksPriority")
	delete(f, "TasksNotifyMode")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ProductInstance struct {
	// 预付费商品实例类型，取值有：
	// <li>StarterPackage：点播新手包。</li>
	// <li>MiniProgramPlugin：点播小程序插件。</li>
	// <li>ResourcePackage：点播资源包。</li>
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 资源包实例起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 资源包实例过期日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 资源包实例ID。对应每个资源包，系统会分配相应的资源。续费或者升级资源包时，需要带上这个资源ID。
	ProductInstanceId *string `json:"ProductInstanceId,omitnil,omitempty" name:"ProductInstanceId"`

	// 系统最近一次扣除资源包的日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	LastConsumeDate *string `json:"LastConsumeDate,omitnil,omitempty" name:"LastConsumeDate"`

	// 资源包绑定 License 状态，取值有：
	// <li>0：未绑定。</li>
	// <li>1：已绑定。</li>
	BindStatus *int64 `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`

	// 预付费资源包实例中包含的资源包列表。
	//
	// Deprecated: ProductInstanceResourceSet is deprecated.
	ProductInstanceResourceSet []*ProductInstanceRecource `json:"ProductInstanceResourceSet,omitnil,omitempty" name:"ProductInstanceResourceSet"`

	// 预付费资源包实例中包含的资源包列表。
	ResourceSet []*ProductInstanceResource `json:"ResourceSet,omitnil,omitempty" name:"ResourceSet"`

	// 资源包实例的状态，取值有：
	// <li>Effective：生效，可用于计费抵扣。</li>
	// <li>Isolated：隔离，不可用于计费抵扣。</li>
	ProductInstanceStatus *string `json:"ProductInstanceStatus,omitnil,omitempty" name:"ProductInstanceStatus"`

	// 资源包实例的可退还状态，取值有：
	// <li>FullRefund：可全额退款。</li>
	// <li>Denied：不可退款。</li>
	RefundStatus *string `json:"RefundStatus,omitnil,omitempty" name:"RefundStatus"`

	// 自动续费状态，取值有：
	// <li>Never：不自动续费。</li>
	// <li>Expire：到期自动续费。</li>
	// <li>ExpireOrUseOut：到期或用完自动续费。</li>
	// <li>NotSupport：不支持。</li>
	RenewStatus *string `json:"RenewStatus,omitnil,omitempty" name:"RenewStatus"`
}

type ProductInstanceRecource struct {
	// 资源类型。
	// <li>Storage：存储资源包。</li>
	// <li>Traffic：流量资源包。</li>
	// <li>Transcode：普通转码资源包。</li>
	// <li>TESHD：极速高清转码资源包。</li>
	// <li>Review：音视频审核转码资源包。</li>
	// <li>MediaProcess：媒体处理时长资源包。</li>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源包额度。
	// <li>音视频存储资源包，单位为字节。</li>
	// <li>音视频转码资源包，单位为秒。</li>
	// <li>音视频审核资源包，单位为秒。</li>
	// <li>音视频极速高清资源包，单位为秒。</li>
	// <li>音视频加速资源包，单位为字节。</li>
	// <li>媒体处理时长资源包，单位为秒。</li>
	Amount *int64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 资源包余量。
	// <li>音视频存储资源包，单位为字节。</li>
	// <li>音视频转码资源包，单位为秒。</li>
	// <li>音视频审核资源包，单位为秒。</li>
	// <li>音视频极速高清资源包，单位为秒。</li>
	// <li>音视频加速资源包，单位为字节。</li>
	// <li>媒体处理时长资源包，单位为秒。</li>
	Left *int64 `json:"Left,omitnil,omitempty" name:"Left"`
}

type ProductInstanceResource struct {
	// 资源类型。
	// <li>Storage：存储资源包。</li>
	// <li>Traffic：流量资源包。</li>
	// <li>Transcode：普通转码资源包。</li>
	// <li>TESHD：极速高清转码资源包。</li>
	// <li>Review：音视频审核转码资源包。</li>
	// <li>MediaProcess：媒体处理时长资源包。</li>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源包额度。
	// <li>音视频存储资源包，单位为字节。</li>
	// <li>音视频转码资源包，单位为秒。</li>
	// <li>音视频审核资源包，单位为秒。</li>
	// <li>音视频极速高清资源包，单位为秒。</li>
	// <li>音视频加速资源包，单位为字节。</li>
	// <li>媒体处理时长资源包，单位为秒。</li>
	Amount *int64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 资源包余量。
	// <li>音视频存储资源包，单位为字节。</li>
	// <li>音视频转码资源包，单位为秒。</li>
	// <li>音视频审核资源包，单位为秒。</li>
	// <li>音视频极速高清资源包，单位为秒。</li>
	// <li>音视频加速资源包，单位为字节。</li>
	// <li>媒体处理时长资源包，单位为秒。</li>
	Left *int64 `json:"Left,omitnil,omitempty" name:"Left"`
}

type ProhibitedAsrReviewTemplateInfo struct {
	// 语音违禁任务开关，可选值：
	// <li>ON：开启语音违禁任务；</li>
	// <li>OFF：关闭语音违禁任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type ProhibitedAsrReviewTemplateInfoForUpdate struct {
	// 语音违禁任务开关，可选值：
	// <li>ON：开启语音违禁任务；</li>
	// <li>OFF：关闭语音违禁任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type ProhibitedConfigureInfo struct {
	// 语音违禁控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrReviewInfo *ProhibitedAsrReviewTemplateInfo `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// 文本违禁控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrReviewInfo *ProhibitedOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type ProhibitedConfigureInfoForUpdate struct {
	// 语音违禁控制参数。
	AsrReviewInfo *ProhibitedAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// 文本违禁控制参数。
	OcrReviewInfo *ProhibitedOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type ProhibitedOcrReviewTemplateInfo struct {
	// 文本违禁任务开关，可选值：
	// <li>ON：开启文本违禁任务；</li>
	// <li>OFF：关闭文本违禁任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type ProhibitedOcrReviewTemplateInfoForUpdate struct {
	// 文本违禁任务开关，可选值：
	// <li>ON：开启文本违禁任务；</li>
	// <li>OFF：关闭文本违禁任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

// Predefined struct for user
type PullEventsRequestParams struct {
	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type PullEventsRequest struct {
	*tchttp.BaseRequest
	
	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *PullEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExtInfo")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullEventsResponseParams struct {
	// 事件列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventSet []*EventContent `json:"EventSet,omitnil,omitempty" name:"EventSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PullEventsResponse struct {
	*tchttp.BaseResponse
	Response *PullEventsResponseParams `json:"Response"`
}

func (r *PullEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullUploadRequestParams struct {
	// 要拉取的媒体 URL，暂不支持拉取 Dash 格式（可以支持 HLS）。
	// 支持的扩展名详见[媒体类型](https://cloud.tencent.com/document/product/266/9760#.E5.AA.92.E4.BD.93.E7.B1.BB.E5.9E.8B)。请确保媒体 URL 可以访问。
	MediaUrl *string `json:"MediaUrl,omitnil,omitempty" name:"MediaUrl"`

	// 媒体文件类型（扩展名），支持的类型详见[媒体类型](https://cloud.tencent.com/document/product/266/9760#.E5.AA.92.E4.BD.93.E7.B1.BB.E5.9E.8B)。
	// 如果 MediaType 不填或取值为空字符串，将根据 MediaUrl 自动获取文件类型。
	MediaType *string `json:"MediaType,omitnil,omitempty" name:"MediaType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 媒体名称。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 要拉取的视频封面 URL。支持的文件格式：gif、jpeg（jpg）、png。
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// 媒体后续任务操作，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 媒体文件过期时间，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 指定上传园区，仅适用于对上传地域有特殊需求的用户：
	// <li>不填默认上传至您的[默认地域](https://cloud.tencent.com/document/product/266/14059?from=11329#.E5.AD.98.E5.82.A8.E5.9C.B0.E5.9F.9F.E6.AD.A5.E9.AA.A4)。</li>
	// <li>若指定上传园区，请先确认[上传存储设置](https://cloud.tencent.com/document/product/266/14059?from=11329#.E5.AD.98.E5.82.A8.E5.9C.B0.E5.9F.9F.E6.AD.A5.E9.AA.A4)已经开启相应的存储地域。</li>
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 分类ID，用于对媒体进行分类管理，可通过[创建分类](https://cloud.tencent.com/document/product/266/7812)接口，创建分类，获得分类 ID。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，当指定 Procedure 任务后，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// 来源上下文，用于透传用户请求信息，[上传完成回调](/document/product/266/7830) 将返回该字段值，最长 250 个字符。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`
}

type PullUploadRequest struct {
	*tchttp.BaseRequest
	
	// 要拉取的媒体 URL，暂不支持拉取 Dash 格式（可以支持 HLS）。
	// 支持的扩展名详见[媒体类型](https://cloud.tencent.com/document/product/266/9760#.E5.AA.92.E4.BD.93.E7.B1.BB.E5.9E.8B)。请确保媒体 URL 可以访问。
	MediaUrl *string `json:"MediaUrl,omitnil,omitempty" name:"MediaUrl"`

	// 媒体文件类型（扩展名），支持的类型详见[媒体类型](https://cloud.tencent.com/document/product/266/9760#.E5.AA.92.E4.BD.93.E7.B1.BB.E5.9E.8B)。
	// 如果 MediaType 不填或取值为空字符串，将根据 MediaUrl 自动获取文件类型。
	MediaType *string `json:"MediaType,omitnil,omitempty" name:"MediaType"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 媒体名称。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 要拉取的视频封面 URL。支持的文件格式：gif、jpeg（jpg）、png。
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// 媒体后续任务操作，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 媒体文件过期时间，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 指定上传园区，仅适用于对上传地域有特殊需求的用户：
	// <li>不填默认上传至您的[默认地域](https://cloud.tencent.com/document/product/266/14059?from=11329#.E5.AD.98.E5.82.A8.E5.9C.B0.E5.9F.9F.E6.AD.A5.E9.AA.A4)。</li>
	// <li>若指定上传园区，请先确认[上传存储设置](https://cloud.tencent.com/document/product/266/14059?from=11329#.E5.AD.98.E5.82.A8.E5.9C.B0.E5.9F.9F.E6.AD.A5.E9.AA.A4)已经开启相应的存储地域。</li>
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 分类ID，用于对媒体进行分类管理，可通过[创建分类](https://cloud.tencent.com/document/product/266/7812)接口，创建分类，获得分类 ID。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，当指定 Procedure 任务后，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// 来源上下文，用于透传用户请求信息，[上传完成回调](/document/product/266/7830) 将返回该字段值，最长 250 个字符。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`
}

func (r *PullUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaUrl")
	delete(f, "MediaType")
	delete(f, "SubAppId")
	delete(f, "MediaName")
	delete(f, "CoverUrl")
	delete(f, "Procedure")
	delete(f, "ExpireTime")
	delete(f, "StorageRegion")
	delete(f, "ClassId")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	delete(f, "SourceContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullUploadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullUploadResponseParams struct {
	// 拉取上传视频的任务 ID，可以通过该 ID 查询拉取上传任务的状态。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PullUploadResponse struct {
	*tchttp.BaseResponse
	Response *PullUploadResponseParams `json:"Response"`
}

func (r *PullUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PullUploadTask struct {
	// 拉取上传任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，0 表示成功，其他值表示失败：
	// <li>40000：输入参数不合法，请检查输入参数；</li>
	// <li>60000：源文件错误（如视频数据损坏），请确认源文件是否正常；</li>
	// <li>70000：内部服务错误，建议重试。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 拉取上传完成后生成的视频 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 拉取上传完成后生成的媒体文件基础信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaBasicInfo *MediaBasicInfo `json:"MediaBasicInfo,omitnil,omitempty" name:"MediaBasicInfo"`

	// 输出视频的元信息。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 拉取上传完成后生成的播放地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 任务类型为 Procedure 的任务 ID。若[拉取上传](https://cloud.tencent.com/document/api/266/35575)时指定了媒体后续任务操作(Procedure)，当该任务流模板指定了 MediaProcessTask、AiAnalysisTask、AiRecognitionTask 中的一个或多个时发起该任务。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitnil,omitempty" name:"ProcedureTaskId"`

	// 任务类型为 ReviewAudioVideo 的任务 ID。若[拉取上传](https://cloud.tencent.com/document/api/266/35575)时指定了媒体后续任务操作(Procedure)，当该任务流模板指定了 ReviewAudioVideoTask 时，发起该任务。
	ReviewAudioVideoTaskId *string `json:"ReviewAudioVideoTaskId,omitnil,omitempty" name:"ReviewAudioVideoTaskId"`

	// 来源上下文，用于透传用户请求信息，[URL 拉取视频上传完成](https://cloud.tencent.com/document/product/266/7831)将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 拉取上传进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

// Predefined struct for user
type PushUrlCacheRequestParams struct {
	// 预热的 URL 列表，单次最多指定20个 URL。
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type PushUrlCacheRequest struct {
	*tchttp.BaseRequest
	
	// 预热的 URL 列表，单次最多指定20个 URL。
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *PushUrlCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PushUrlCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PushUrlCacheResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PushUrlCacheResponse struct {
	*tchttp.BaseResponse
	Response *PushUrlCacheResponseParams `json:"Response"`
}

func (r *PushUrlCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QRCodeConfigureInfo struct {
	// 视频画面二维码检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type QRCodeConfigureInfoForUpdate struct {
	// 视频画面二维码检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type QualityEnhanceTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，0 表示成功，其他值表示失败：
	// <li>40000：输入参数不合法，请检查输入参数；</li>
	// <li>60000：源文件错误（如视频数据损坏），请确认源文件是否正常；</li>
	// <li>70000：内部服务错误，建议重试。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 音画质重生任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 音画质重生任务的输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *QualityEnhanceTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音画质重生任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *QualityEnhanceTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音画质重生输出视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type QualityEnhanceTaskInput struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 音画质重生模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type QualityEnhanceTaskOutput struct {
	// 文件类型，例如 mp4、flv 等。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 媒体文件播放地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 输出文件名，最长 64 个字符。缺省由系统指定生成文件名。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type QualityEvaluationConfigureInfo struct {
	// 视频画面质量评价检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 视频画面质量评价过滤阈值，结果只返回低于该值的时间段，默认值为 60。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`
}

type QualityEvaluationConfigureInfoForUpdate struct {
	// 视频画面质量评价开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 视频画面质量评价过滤阈值，结果只返回低于该值的时间段。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`
}

type QualityInspectItem struct {
	// 异常片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 异常片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 检测出异常的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	// <font color=red>注意：</font> 仅当 Type 取值为下列之一时，本字段取值有效：
	// <li>BlackWhiteEdge：黑白边；</li>
	// <li>Mosaic：马赛克；</li>
	// <li>QRCode：二维码；</li>
	// <li>AppletCode：小程序码；</li>
	// <li>BarCode：条形码。</li>
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// 置信度，取值范围：[0, 100]。
	// <font color=red>注意：</font> 仅当 Type 取值为下列之一时，本字段取值有效：
	// <li>Mosaic：马赛克；</li>
	// <li>QRCode：二维码；</li>
	// <li>AppletCode：小程序码；</li>
	// <li>BarCode：条形码。</li>
	Confidence *uint64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type QualityInspectResultItem struct {
	// 异常类型，取值范围：
	// <li>Jitter：抖动；</li>
	// <li>Blur：模糊；</li>
	// <li>LowLighting：低光照；</li>
	// <li>HighLighting：过曝；</li>
	// <li>CrashScreen：花屏；</li>
	// <li>BlackWhiteEdge：黑白边；</li>
	// <li>SolidColorScreen：纯色屏；</li>
	// <li>Noise：噪点；</li>
	// <li>Mosaic：马赛克；</li>
	// <li>QRCode：二维码；</li>
	// <li>AppletCode：小程序码；</li>
	// <li>BarCode：条形码；</li>
	// <li>LowVoice：低音；</li>
	// <li>HighVoice：爆音；</li>
	// <li>NoVoice：静音；</li>
	// <li>LowEvaluation：无参考打分低于阈值。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 异常片段列表。
	// <font color=red>注意：</font> 该列表最多仅展示前 100 个元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*QualityInspectItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 异常片段列表文件 URL。文件 内容 为  JSON，数据结构与 SegmentSet 字段一致。（文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 异常片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type QualityInspectTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音画质检测输入音视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 音画质检测任务输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *QualityInspectTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音画质检测任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *QualityInspectTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，音画质检测完成回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type QualityInspectTaskInput struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 音画质检测模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type QualityInspectTaskOutput struct {
	// 媒体文件是否无音频轨，取值范围：
	// <li>0：否，即有音频轨；</li>
	// <li>1：是，即无音频轨。</li>
	NoAudio *int64 `json:"NoAudio,omitnil,omitempty" name:"NoAudio"`

	// 媒体文件是否无视频轨，取值范围：
	// <li>0：否，即有视频轨；</li>
	// <li>1：是，即无视频轨。</li>
	NoVideo *int64 `json:"NoVideo,omitnil,omitempty" name:"NoVideo"`

	// 视频画面质量评分，取值范围：[0, 100]。
	QualityEvaluationScore *uint64 `json:"QualityEvaluationScore,omitnil,omitempty" name:"QualityEvaluationScore"`

	// 音画质检测出的异常项列表。
	QualityInspectResultSet []*QualityInspectResultItem `json:"QualityInspectResultSet,omitnil,omitempty" name:"QualityInspectResultSet"`
}

type QualityInspectTemplateItem struct {
	// 模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板类型，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 截帧间隔，单位为秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitnil,omitempty" name:"ScreenshotInterval"`

	// 视频画面抖动重影检测的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JitterConfigure *JitterConfigureInfo `json:"JitterConfigure,omitnil,omitempty" name:"JitterConfigure"`

	// 视频画面模糊检测的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlurConfigure *BlurConfigureInfo `json:"BlurConfigure,omitnil,omitempty" name:"BlurConfigure"`

	// 视频画面低光、过曝检测的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AbnormalLightingConfigure *AbnormalLightingConfigureInfo `json:"AbnormalLightingConfigure,omitnil,omitempty" name:"AbnormalLightingConfigure"`

	// 视频画面花屏检测的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrashScreenConfigure *CrashScreenConfigureInfo `json:"CrashScreenConfigure,omitnil,omitempty" name:"CrashScreenConfigure"`

	// 视频画面黑边、白边、黑屏、白屏检测的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlackWhiteEdgeConfigure *BlackWhiteEdgeConfigureInfo `json:"BlackWhiteEdgeConfigure,omitnil,omitempty" name:"BlackWhiteEdgeConfigure"`

	// 视频画面噪点检测的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoiseConfigure *NoiseConfigureInfo `json:"NoiseConfigure,omitnil,omitempty" name:"NoiseConfigure"`

	// 视频画面马赛克检测的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MosaicConfigure *MosaicConfigureInfo `json:"MosaicConfigure,omitnil,omitempty" name:"MosaicConfigure"`

	// 视频画面二维码检测的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QRCodeConfigure *QRCodeConfigureInfo `json:"QRCodeConfigure,omitnil,omitempty" name:"QRCodeConfigure"`

	// 视频画面质量评价的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityEvaluationConfigure *QualityEvaluationConfigureInfo `json:"QualityEvaluationConfigure,omitnil,omitempty" name:"QualityEvaluationConfigure"`

	// 音频（静音、低音、爆音）检测的控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceConfigure *VoiceConfigureInfo `json:"VoiceConfigure,omitnil,omitempty" name:"VoiceConfigure"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type RebuildAudioInfo struct {
	// 音频降噪控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioDenoiseInfo *AudioDenoiseInfo `json:"AudioDenoiseInfo,omitnil,omitempty" name:"AudioDenoiseInfo"`
}

// Predefined struct for user
type RebuildMediaByTemplateRequestParams struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 视频重生模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *string `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 起始偏移时间，单位：秒，不填表示从视频开始截取。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 结束偏移时间，单位：秒，不填表示截取到视频末尾。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 视频重生后的文件配置。
	OutputConfig *RebuildMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type RebuildMediaByTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 视频重生模板 ID。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *string `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 起始偏移时间，单位：秒，不填表示从视频开始截取。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 结束偏移时间，单位：秒，不填表示截取到视频末尾。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 视频重生后的文件配置。
	OutputConfig *RebuildMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *RebuildMediaByTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebuildMediaByTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "Definition")
	delete(f, "SubAppId")
	delete(f, "StartTimeOffset")
	delete(f, "EndTimeOffset")
	delete(f, "OutputConfig")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "TasksPriority")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebuildMediaByTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebuildMediaByTemplateResponseParams struct {
	// 视频重生的任务 ID，可以通过该 ID 查询视频重生任务的状态。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RebuildMediaByTemplateResponse struct {
	*tchttp.BaseResponse
	Response *RebuildMediaByTemplateResponseParams `json:"Response"`
}

func (r *RebuildMediaByTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebuildMediaByTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebuildMediaOutputConfig struct {
	// 输出文件名，最长 64 个字符。缺省由系统指定生成文件名。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type RebuildMediaRequestParams struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 起始偏移时间，单位：秒，不填表示从视频开始截取。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 结束偏移时间，单位：秒，不填表示截取到视频末尾。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 画质修复控制参数。
	RepairInfo *RepairInfo `json:"RepairInfo,omitnil,omitempty" name:"RepairInfo"`

	// 智能插帧控制参数。
	VideoFrameInterpolationInfo *VideoFrameInterpolationInfo `json:"VideoFrameInterpolationInfo,omitnil,omitempty" name:"VideoFrameInterpolationInfo"`

	// 画面超分控制参数。
	SuperResolutionInfo *SuperResolutionInfo `json:"SuperResolutionInfo,omitnil,omitempty" name:"SuperResolutionInfo"`

	// 高动态范围类型控制参数。
	HDRInfo *HDRInfo `json:"HDRInfo,omitnil,omitempty" name:"HDRInfo"`

	// 视频降噪控制参数。
	VideoDenoiseInfo *VideoDenoiseInfo `json:"VideoDenoiseInfo,omitnil,omitempty" name:"VideoDenoiseInfo"`

	// 音频降噪控制参数。
	AudioDenoiseInfo *AudioDenoiseInfo `json:"AudioDenoiseInfo,omitnil,omitempty" name:"AudioDenoiseInfo"`

	// 色彩增强控制参数。
	ColorInfo *ColorEnhanceInfo `json:"ColorInfo,omitnil,omitempty" name:"ColorInfo"`

	// 细节增强控制参数。
	SharpInfo *SharpEnhanceInfo `json:"SharpInfo,omitnil,omitempty" name:"SharpInfo"`

	// 人脸增强控制参数。
	FaceInfo *FaceEnhanceInfo `json:"FaceInfo,omitnil,omitempty" name:"FaceInfo"`

	// 低光照控制参数。
	LowLightInfo *LowLightEnhanceInfo `json:"LowLightInfo,omitnil,omitempty" name:"LowLightInfo"`

	// 去划痕控制参数。
	ScratchRepairInfo *ScratchRepairInfo `json:"ScratchRepairInfo,omitnil,omitempty" name:"ScratchRepairInfo"`

	// 去伪影（毛刺）控制参数。
	ArtifactRepairInfo *ArtifactRepairInfo `json:"ArtifactRepairInfo,omitnil,omitempty" name:"ArtifactRepairInfo"`

	// 音画质重生输出目标参数。
	TargetInfo *RebuildMediaTargetInfo `json:"TargetInfo,omitnil,omitempty" name:"TargetInfo"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type RebuildMediaRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 起始偏移时间，单位：秒，不填表示从视频开始截取。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 结束偏移时间，单位：秒，不填表示截取到视频末尾。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 画质修复控制参数。
	RepairInfo *RepairInfo `json:"RepairInfo,omitnil,omitempty" name:"RepairInfo"`

	// 智能插帧控制参数。
	VideoFrameInterpolationInfo *VideoFrameInterpolationInfo `json:"VideoFrameInterpolationInfo,omitnil,omitempty" name:"VideoFrameInterpolationInfo"`

	// 画面超分控制参数。
	SuperResolutionInfo *SuperResolutionInfo `json:"SuperResolutionInfo,omitnil,omitempty" name:"SuperResolutionInfo"`

	// 高动态范围类型控制参数。
	HDRInfo *HDRInfo `json:"HDRInfo,omitnil,omitempty" name:"HDRInfo"`

	// 视频降噪控制参数。
	VideoDenoiseInfo *VideoDenoiseInfo `json:"VideoDenoiseInfo,omitnil,omitempty" name:"VideoDenoiseInfo"`

	// 音频降噪控制参数。
	AudioDenoiseInfo *AudioDenoiseInfo `json:"AudioDenoiseInfo,omitnil,omitempty" name:"AudioDenoiseInfo"`

	// 色彩增强控制参数。
	ColorInfo *ColorEnhanceInfo `json:"ColorInfo,omitnil,omitempty" name:"ColorInfo"`

	// 细节增强控制参数。
	SharpInfo *SharpEnhanceInfo `json:"SharpInfo,omitnil,omitempty" name:"SharpInfo"`

	// 人脸增强控制参数。
	FaceInfo *FaceEnhanceInfo `json:"FaceInfo,omitnil,omitempty" name:"FaceInfo"`

	// 低光照控制参数。
	LowLightInfo *LowLightEnhanceInfo `json:"LowLightInfo,omitnil,omitempty" name:"LowLightInfo"`

	// 去划痕控制参数。
	ScratchRepairInfo *ScratchRepairInfo `json:"ScratchRepairInfo,omitnil,omitempty" name:"ScratchRepairInfo"`

	// 去伪影（毛刺）控制参数。
	ArtifactRepairInfo *ArtifactRepairInfo `json:"ArtifactRepairInfo,omitnil,omitempty" name:"ArtifactRepairInfo"`

	// 音画质重生输出目标参数。
	TargetInfo *RebuildMediaTargetInfo `json:"TargetInfo,omitnil,omitempty" name:"TargetInfo"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *RebuildMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebuildMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "StartTimeOffset")
	delete(f, "EndTimeOffset")
	delete(f, "RepairInfo")
	delete(f, "VideoFrameInterpolationInfo")
	delete(f, "SuperResolutionInfo")
	delete(f, "HDRInfo")
	delete(f, "VideoDenoiseInfo")
	delete(f, "AudioDenoiseInfo")
	delete(f, "ColorInfo")
	delete(f, "SharpInfo")
	delete(f, "FaceInfo")
	delete(f, "LowLightInfo")
	delete(f, "ScratchRepairInfo")
	delete(f, "ArtifactRepairInfo")
	delete(f, "TargetInfo")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "TasksPriority")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebuildMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebuildMediaResponseParams struct {
	// 音画质重生的任务 ID，可以通过该 ID 查询音画质重生任务的状态。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RebuildMediaResponse struct {
	*tchttp.BaseResponse
	Response *RebuildMediaResponseParams `json:"Response"`
}

func (r *RebuildMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebuildMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebuildMediaTargetAudioStream struct {
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
	// <li>libfdk_aac。</li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。
	// 当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 音频流的采样率，可选值：
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 
	// 单位：Hz。
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频通道方式，可选值：
	// <li>1：单通道</li>
	// <li>2：双通道</li>
	// <li>6：立体声</li>
	// 
	// 当媒体的封装格式是音频格式时（flac，ogg，mp3，m4a）时，声道数不允许设为立体声。
	// 默认值：2。
	AudioChannel *int64 `json:"AudioChannel,omitnil,omitempty" name:"AudioChannel"`
}

type RebuildMediaTargetInfo struct {
	// 输出文件名，最长 64 个字符。缺省由系统指定生成文件名。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 描述信息，最长 128 个字符。缺省描述信息为空。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 输出文件封装格式，可选值：mp4、flv、hls。默认mp4。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 输出的视频信息。
	VideoStream *RebuildMediaTargetVideoStream `json:"VideoStream,omitnil,omitempty" name:"VideoStream"`

	// 输出的音频信息。
	AudioStream *RebuildMediaTargetAudioStream `json:"AudioStream,omitnil,omitempty" name:"AudioStream"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`
}

type RebuildMediaTargetVideoStream struct {
	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码；</li>
	// <li>libx265：H.265 编码；</li>
	// <li>av1：AOMedia Video 1 编码。</li>
	// 默认视频流的编码格式为 H.264 编码。
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 视频帧率，取值范围：[0, 100]，单位：Hz。 当取值为 0，表示帧率和原始视频保持一致。
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 
	// 默认值：0。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 
	// 默认值：0。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li>stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 
	// 默认值：stretch 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 关键帧 I 帧之间的间隔，取值范围：0 和 [1, 100000]，单位：帧数。
	// 当填 0 或不填时，系统将自动设置 gop 长度。
	Gop *int64 `json:"Gop,omitnil,omitempty" name:"Gop"`
}

type RebuildMediaTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，0 表示成功，其他值表示失败：
	// <li>40000：输入参数不合法，请检查输入参数；</li>
	// <li>60000：源文件错误（如视频数据损坏），请确认源文件是否正常；</li>
	// <li>70000：内部服务错误，建议重试。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 音画质重生任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 音画质重生任务的输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *RebuildMediaTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音画质重生任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *RebuildMediaTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 音画质重生输出视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type RebuildMediaTaskInput struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 起始偏移时间，单位：秒，不填表示从视频开始截取。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 结束偏移时间，单位：秒，不填表示截取到视频末尾。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 音画质重生模版号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 画质修复控制参数。
	RepairInfo *RepairInfo `json:"RepairInfo,omitnil,omitempty" name:"RepairInfo"`

	// 智能插帧控制参数。
	VideoFrameInterpolationInfo *VideoFrameInterpolationInfo `json:"VideoFrameInterpolationInfo,omitnil,omitempty" name:"VideoFrameInterpolationInfo"`

	// 画面超分控制参数。
	SuperResolutionInfo *SuperResolutionInfo `json:"SuperResolutionInfo,omitnil,omitempty" name:"SuperResolutionInfo"`

	// 高动态范围类型控制参数。
	HDRInfo *HDRInfo `json:"HDRInfo,omitnil,omitempty" name:"HDRInfo"`

	// 视频降噪控制参数。
	VideoDenoiseInfo *VideoDenoiseInfo `json:"VideoDenoiseInfo,omitnil,omitempty" name:"VideoDenoiseInfo"`

	// 音频降噪控制参数。
	AudioDenoiseInfo *AudioDenoiseInfo `json:"AudioDenoiseInfo,omitnil,omitempty" name:"AudioDenoiseInfo"`

	// 色彩增强控制参数。
	ColorInfo *ColorEnhanceInfo `json:"ColorInfo,omitnil,omitempty" name:"ColorInfo"`

	// 细节增强控制参数。
	SharpInfo *SharpEnhanceInfo `json:"SharpInfo,omitnil,omitempty" name:"SharpInfo"`

	// 人脸增强控制参数。
	FaceInfo *FaceEnhanceInfo `json:"FaceInfo,omitnil,omitempty" name:"FaceInfo"`

	// 低光照控制参数。
	LowLightInfo *LowLightEnhanceInfo `json:"LowLightInfo,omitnil,omitempty" name:"LowLightInfo"`

	// 去划痕控制参数。
	ScratchRepairInfo *ScratchRepairInfo `json:"ScratchRepairInfo,omitnil,omitempty" name:"ScratchRepairInfo"`

	// 去伪影（毛刺）控制参数。
	ArtifactRepairInfo *ArtifactRepairInfo `json:"ArtifactRepairInfo,omitnil,omitempty" name:"ArtifactRepairInfo"`

	// 音画质重生输出目标参数。
	TargetInfo *RebuildMediaTargetInfo `json:"TargetInfo,omitnil,omitempty" name:"TargetInfo"`
}

type RebuildMediaTaskOutput struct {
	// 文件类型，例如 mp4、flv 等。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 媒体文件播放地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 输出文件名，最长 64 个字符。缺省由系统指定生成文件名。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type RebuildMediaTemplate struct {
	// 音画质重生模板号。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板类型，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 音画质重生模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 音画质重生模板描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 音画质重生视频控制信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RebuildVideoInfo *RebuildVideoInfo `json:"RebuildVideoInfo,omitnil,omitempty" name:"RebuildVideoInfo"`

	// 音画质重生音频控制信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RebuildAudioInfo *RebuildAudioInfo `json:"RebuildAudioInfo,omitnil,omitempty" name:"RebuildAudioInfo"`

	// 输出视频控制信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetVideoInfo *RebuildMediaTargetVideoStream `json:"TargetVideoInfo,omitnil,omitempty" name:"TargetVideoInfo"`

	// 输出音频控制信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetAudioInfo *RebuildMediaTargetAudioStream `json:"TargetAudioInfo,omitnil,omitempty" name:"TargetAudioInfo"`

	// 封装格式。可选值：mp4、hls。默认是 mp4。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值 0。
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type RebuildVideoInfo struct {
	// 画质修复控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepairInfo *RepairInfo `json:"RepairInfo,omitnil,omitempty" name:"RepairInfo"`

	// 智能插帧控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoFrameInterpolationInfo *VideoFrameInterpolationInfo `json:"VideoFrameInterpolationInfo,omitnil,omitempty" name:"VideoFrameInterpolationInfo"`

	// 画面超分控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuperResolutionInfo *SuperResolutionInfo `json:"SuperResolutionInfo,omitnil,omitempty" name:"SuperResolutionInfo"`

	// 高动态范围类型控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HDRInfo *HDRInfo `json:"HDRInfo,omitnil,omitempty" name:"HDRInfo"`

	// 视频降噪控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoDenoiseInfo *VideoDenoiseInfo `json:"VideoDenoiseInfo,omitnil,omitempty" name:"VideoDenoiseInfo"`

	// 色彩增强控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColorInfo *ColorEnhanceInfo `json:"ColorInfo,omitnil,omitempty" name:"ColorInfo"`

	// 细节增强控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SharpInfo *SharpEnhanceInfo `json:"SharpInfo,omitnil,omitempty" name:"SharpInfo"`

	// 人脸增强控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceInfo *FaceEnhanceInfo `json:"FaceInfo,omitnil,omitempty" name:"FaceInfo"`

	// 低光照控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LowLightInfo *LowLightEnhanceInfo `json:"LowLightInfo,omitnil,omitempty" name:"LowLightInfo"`

	// 去划痕控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScratchRepairInfo *ScratchRepairInfo `json:"ScratchRepairInfo,omitnil,omitempty" name:"ScratchRepairInfo"`

	// 去伪影控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArtifactRepairInfo *ArtifactRepairInfo `json:"ArtifactRepairInfo,omitnil,omitempty" name:"ArtifactRepairInfo"`
}

type ReduceMediaBitrateAdaptiveDynamicStreamingResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 对视频转自适应码流任务的输入。
	Input *AdaptiveDynamicStreamingTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频转自适应码流任务的输出。
	Output *AdaptiveDynamicStreamingInfoItem `json:"Output,omitnil,omitempty" name:"Output"`
}

type ReduceMediaBitrateMediaProcessTaskResult struct {
	// 任务的类型，可以取的值有：
	// <li>Transcode：转码</li>
	// <li>AdaptiveDynamicStreaming：自适应码流</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 降码率任务中视频转码任务的查询结果，当任务类型为 Transcode 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTask *ReduceMediaBitrateTranscodeResult `json:"TranscodeTask,omitnil,omitempty" name:"TranscodeTask"`

	// 降码率任务中对视频转自适应码流任务的查询结果，当任务类型为 AdaptiveDynamicStreaming 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingTask *ReduceMediaBitrateAdaptiveDynamicStreamingResult `json:"AdaptiveDynamicStreamingTask,omitnil,omitempty" name:"AdaptiveDynamicStreamingTask"`
}

type ReduceMediaBitrateTask struct {
	// 视频处理任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 媒体文件名称。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 媒体文件地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 原始视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 降码率任务执行状态与结果。
	MediaProcessResultSet []*ReduceMediaBitrateMediaProcessTaskResult `json:"MediaProcessResultSet,omitnil,omitempty" name:"MediaProcessResultSet"`

	// 任务流的优先级，取值范围为 [-10, 10]。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式。
	// <li>Finish：只有当任务流全部执行完毕时，才发起一次事件通知；</li>
	// <li>None：不接受该任务流回调。</li>
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type ReduceMediaBitrateTranscodeResult struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 转码任务的输入。
	Input *TranscodeTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 转码任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaTranscodeItem `json:"Output,omitnil,omitempty" name:"Output"`

	// 转码进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 转码任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 转码任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type RefererAuthPolicy struct {
	// [Referer 防盗链](https://cloud.tencent.com/document/product/266/14046)设置状态，可选值：
	// <li>Enabled: 启用；</li>
	// <li>Disabled: 禁用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Referer 校验类型，可选值：
	// <li>Black：黑名单方式校验。HTTP 请求携带了 Referers 列表中的某个 Referer 将被拒绝访问。</li>
	// <li>White：白名单方式校验。HTTP 请求携带了 Referers 列表中的 Referer 时才允许访问。</li>
	// 当 Status 取值为 Enabled 时，AuthType 必须赋值。
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 用于校验的 Referer 列表，最大支持20个 Referer。当 Status 取值为 Enabled 时， Referers 不能为空数组。Referer 的格式参考域名的格式。
	Referers []*string `json:"Referers,omitnil,omitempty" name:"Referers"`

	// 是否允许空 Referer 访问本域名，可选值：
	// <li>Yes： 是。</li>
	// <li>No： 否。</li>
	// 当 Status 取值为 Enabled 时，BlankRefererAllowed 必须赋值。
	BlankRefererAllowed *string `json:"BlankRefererAllowed,omitnil,omitempty" name:"BlankRefererAllowed"`
}

// Predefined struct for user
type RefreshUrlCacheRequestParams struct {
	// 刷新的 URL 列表，单次最多指定20个 URL。
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type RefreshUrlCacheRequest struct {
	*tchttp.BaseRequest
	
	// 刷新的 URL 列表，单次最多指定20个 URL。
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *RefreshUrlCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshUrlCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshUrlCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshUrlCacheResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RefreshUrlCacheResponse struct {
	*tchttp.BaseResponse
	Response *RefreshUrlCacheResponseParams `json:"Response"`
}

func (r *RefreshUrlCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshUrlCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveWaterMarkTaskInput struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

type RemoveWaterMarkTaskOutput struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 文件类型，例如 mp4、mp3 等。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 媒体文件播放地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件名称，最长 64 个字符。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 元信息。包括大小、时长、视频流信息、音频流信息等。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`
}

// Predefined struct for user
type RemoveWatermarkRequestParams struct {
	// 媒体文件 ID 。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 该字段已无效。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`
}

type RemoveWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID 。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 该字段已无效。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitnil,omitempty" name:"TasksNotifyMode"`
}

func (r *RemoveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "TasksPriority")
	delete(f, "TasksNotifyMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveWatermarkResponseParams struct {
	// 任务 ID 。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *RemoveWatermarkResponseParams `json:"Response"`
}

func (r *RemoveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveWatermarkTask struct {
	// 任务 ID 。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败：
	// <li>40000：输入参数不合法，请检查输入参数；</li>
	// <li>60000：源文件错误（如视频数据损坏），请确认源文件是否正常；</li>
	// <li>70000：内部服务错误，建议重试。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 智能去除水印任务的输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *RemoveWaterMarkTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 智能去除水印任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *RemoveWaterMarkTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type RepairInfo struct {
	// 画质修复控制开关，可选值：
	// <li>ON：开启画质修复；</li>
	// <li>OFF：关闭画质修复。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 画质修复类型，仅当画质修复控制开关为 ON 时有效，可选值：
	// <li>weak：轻画质修复；</li>
	// <li>normal：正常画质修复；</li>
	// <li>strong：强画质修复。</li>
	// 默认值：weak。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type ResetProcedureTemplateRequestParams struct {
	// 任务流名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// AI 智能内容审核类型任务参数 \*。
	// <font color=red>\*：该参数用于发起旧版审核，不建议使用。推荐使用 ReviewAudioVideoTask 参数发起审核。</font> 
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// AI 智能内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// AI 内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// 音视频审核类型任务参数。
	ReviewAudioVideoTask *ProcedureReviewAudioVideoTaskInput `json:"ReviewAudioVideoTask,omitnil,omitempty" name:"ReviewAudioVideoTask"`
}

type ResetProcedureTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 任务流名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// AI 智能内容审核类型任务参数 \*。
	// <font color=red>\*：该参数用于发起旧版审核，不建议使用。推荐使用 ReviewAudioVideoTask 参数发起审核。</font> 
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// AI 智能内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// AI 内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// 音视频审核类型任务参数。
	ReviewAudioVideoTask *ProcedureReviewAudioVideoTaskInput `json:"ReviewAudioVideoTask,omitnil,omitempty" name:"ReviewAudioVideoTask"`
}

func (r *ResetProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetProcedureTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	delete(f, "Comment")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "ReviewAudioVideoTask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetProcedureTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetProcedureTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ResetProcedureTemplateResponseParams `json:"Response"`
}

func (r *ResetProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetProcedureTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResolutionNameInfo struct {
	// 视频短边长度，单位：像素。
	MinEdgeLength *uint64 `json:"MinEdgeLength,omitnil,omitempty" name:"MinEdgeLength"`

	// 展示名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ResourceTag struct {
	// 标签键。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type RestoreMediaRequestParams struct {
	// 媒体文件唯一标识列表，最大长度：100。
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 解冻出的临时媒体文件的可访问持续时长，必须大于0，单位为“天”。
	RestoreDay *uint64 `json:"RestoreDay,omitnil,omitempty" name:"RestoreDay"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 解冻模式。当媒体文件当前的存储类型为归档存储时，有以下取值：
	// <li>极速模式：Expedited，解冻任务在5分钟后完成。</li>
	// <li>标准模式：Standard，解冻任务在5小时后完成 。</li>
	// <li>批量模式：Bulk，，解冻任务在12小时后完成。</li>
	// 当媒体文件的存储类型为深度归档存储时，有以下取值：
	// <li>标准模式：Standard，解冻任务在24小时后完成。</li>
	// <li>批量模式：Bulk，解冻任务在48小时后完成。</li>
	RestoreTier *string `json:"RestoreTier,omitnil,omitempty" name:"RestoreTier"`
}

type RestoreMediaRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件唯一标识列表，最大长度：100。
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 解冻出的临时媒体文件的可访问持续时长，必须大于0，单位为“天”。
	RestoreDay *uint64 `json:"RestoreDay,omitnil,omitempty" name:"RestoreDay"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 解冻模式。当媒体文件当前的存储类型为归档存储时，有以下取值：
	// <li>极速模式：Expedited，解冻任务在5分钟后完成。</li>
	// <li>标准模式：Standard，解冻任务在5小时后完成 。</li>
	// <li>批量模式：Bulk，，解冻任务在12小时后完成。</li>
	// 当媒体文件的存储类型为深度归档存储时，有以下取值：
	// <li>标准模式：Standard，解冻任务在24小时后完成。</li>
	// <li>批量模式：Bulk，解冻任务在48小时后完成。</li>
	RestoreTier *string `json:"RestoreTier,omitnil,omitempty" name:"RestoreTier"`
}

func (r *RestoreMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileIds")
	delete(f, "RestoreDay")
	delete(f, "SubAppId")
	delete(f, "RestoreTier")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreMediaResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreMediaResponse struct {
	*tchttp.BaseResponse
	Response *RestoreMediaResponseParams `json:"Response"`
}

func (r *RestoreMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreMediaTask struct {
	// 文件ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 文件原始存储类型。
	OriginalStorageClass *string `json:"OriginalStorageClass,omitnil,omitempty" name:"OriginalStorageClass"`

	// 文件目标存储类型。对于临时取回，目标存储类型与原始存储类型相同。
	TargetStorageClass *string `json:"TargetStorageClass,omitnil,omitempty" name:"TargetStorageClass"`

	// 取回模式，取值：
	// <li>Expedited：极速模式</li>
	// <li>Standard：标准模式</li>
	// <li>Bulk：批量模式</li>
	RestoreTier *string `json:"RestoreTier,omitnil,omitempty" name:"RestoreTier"`

	// 临时取回副本有效期，单位：天。对于永久取回，取值为0。
	RestoreDay *int64 `json:"RestoreDay,omitnil,omitempty" name:"RestoreDay"`

	// 该字段已废弃。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 该字段已废弃。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type ReviewAudioVideoRequestParams struct {
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符，在上传成功后由云点播后台分配。可以在 [视频上传完成事件通知](/document/product/266/7830) 或 [云点播控制台](https://console.cloud.tencent.com/vod/media) 获取该字段。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 审核的内容，可选值有：
	// <li>Media：原始音视频；</li>
	// <li>Cover：封面。</li>
	// 不填或填空数组时，默认为审核 Media。
	ReviewContents []*string `json:"ReviewContents,omitnil,omitempty" name:"ReviewContents"`

	// 审核模板 ID，默认值为 10。取值范围：
	// <li>10：[预置模板](https://cloud.tencent.com/document/product/266/33476#.E9.A2.84.E7.BD.AE.E9.9F.B3.E8.A7.86.E9.A2.91.E5.AE.A1.E6.A0.B8.E6.A8.A1.E6.9D.BF.5B.5D(id.3Averify))，支持检测的违规标签包括色情（Porn）、暴力（Terror）、不适宜的信息（Polity）和娇喘（Moan）。</li>
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 来源上下文，用于透传用户请求信息，音视频审核完成回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type ReviewAudioVideoRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符，在上传成功后由云点播后台分配。可以在 [视频上传完成事件通知](/document/product/266/7830) 或 [云点播控制台](https://console.cloud.tencent.com/vod/media) 获取该字段。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 审核的内容，可选值有：
	// <li>Media：原始音视频；</li>
	// <li>Cover：封面。</li>
	// 不填或填空数组时，默认为审核 Media。
	ReviewContents []*string `json:"ReviewContents,omitnil,omitempty" name:"ReviewContents"`

	// 审核模板 ID，默认值为 10。取值范围：
	// <li>10：[预置模板](https://cloud.tencent.com/document/product/266/33476#.E9.A2.84.E7.BD.AE.E9.9F.B3.E8.A7.86.E9.A2.91.E5.AE.A1.E6.A0.B8.E6.A8.A1.E6.9D.BF.5B.5D(id.3Averify))，支持检测的违规标签包括色情（Porn）、暴力（Terror）、不适宜的信息（Polity）和娇喘（Moan）。</li>
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 来源上下文，用于透传用户请求信息，音视频审核完成回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *ReviewAudioVideoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReviewAudioVideoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "ReviewContents")
	delete(f, "Definition")
	delete(f, "TasksPriority")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReviewAudioVideoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReviewAudioVideoResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReviewAudioVideoResponse struct {
	*tchttp.BaseResponse
	Response *ReviewAudioVideoResponseParams `json:"Response"`
}

func (r *ReviewAudioVideoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReviewAudioVideoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviewAudioVideoSegmentItem struct {
	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段涉及令人反感的信息的分数。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 嫌疑片段鉴别涉及违规信息的结果建议，取值范围：
	// <li>review：疑似违规，建议复审；</li>
	// <li>block：确认违规，建议封禁。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 嫌疑片段最可能的违规的标签，取值范围：
	// <li>Porn：色情；</li>
	// <li>Terror：暴力；</li>
	// <li>Polity：不适宜的信息；</li>
	// <li>Ad：广告；</li>
	// <li>Illegal：违法；</li>
	// <li>Abuse：谩骂；</li>
	// <li>Moan：娇喘。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 违规子标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 嫌疑片段违禁的形式，取值范围：
	// <li>Image：画面上的人物或图标；</li>
	// <li>OCR：画面上的文字；</li>
	// <li>ASR：语音中的文字；</li>
	// <li>Voice：声音。</li>
	Form *string `json:"Form,omitnil,omitempty" name:"Form"`

	// 当 Form 为 Image 或 OCR 时有效，表示嫌疑人物、图标或文字出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// 当 Form 为 OCR 或 ASR 时有效，表示识别出来的 OCR 或 ASR 文本内容。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 当 Form 为 OCR 或 ASR 时有效，表示嫌疑片段命中的违规关键词列表。
	KeywordSet []*string `json:"KeywordSet,omitnil,omitempty" name:"KeywordSet"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	//  PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type ReviewAudioVideoTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 音视频审核任务的输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *ReviewAudioVideoTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 音视频审核任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *ReviewAudioVideoTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，音视频审核完成回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type ReviewAudioVideoTaskInput struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 音视频审核模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 审核的内容，可选值：
	// <li>Media：原始音视频；</li>
	// <li>Cover：封面。</li>
	ReviewContents []*string `json:"ReviewContents,omitnil,omitempty" name:"ReviewContents"`
}

type ReviewAudioVideoTaskOutput struct {
	// 音视频内容审核的结果建议，取值范围：
	// <li>pass：建议通过；</li>
	// <li>review：建议复审；</li>
	// <li>block：建议封禁。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 当 Suggestion 为 review 或 block 时有效，表示音视频最可能的违规的标签，取值范围：
	// <li>Porn：色情；</li>
	// <li>Terror：暴力；</li>
	// <li>Polity：不适宜的信息；</li>
	// <li>Ad：广告；</li>
	// <li>Illegal：违法；</li>
	// <li>Abuse：谩骂；</li>
	// <li>Moan：娇喘。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 当 Suggestion 为 review 或 block 时有效，表示音视频最可能的违禁的形式，取值范围：
	// <li>Image：画面上的人物或图标；</li>
	// <li>OCR：画面上的文字；</li>
	// <li>ASR：语音中的文字；</li>
	// <li>Voice：声音。</li>
	Form *string `json:"Form,omitnil,omitempty" name:"Form"`

	// 有违规信息的嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 10个 元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*ReviewAudioVideoSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 涉及违规信息的嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 涉及违规信息的嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`

	// 封面审核结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverReviewResult *ReviewImageResult `json:"CoverReviewResult,omitnil,omitempty" name:"CoverReviewResult"`
}

// Predefined struct for user
type ReviewImageRequestParams struct {
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符。本接口要求媒体文件必须是图片格式。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 图片审核模板 ID，取值范围：
	// <li>10：预置模板，支持检测的违规标签包括色情（Porn）、暴力（Terror）和不适宜的信息（Polity）。</li>
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type ReviewImageRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID，即该文件在云点播上的全局唯一标识符。本接口要求媒体文件必须是图片格式。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 图片审核模板 ID，取值范围：
	// <li>10：预置模板，支持检测的违规标签包括色情（Porn）、暴力（Terror）和不适宜的信息（Polity）。</li>
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *ReviewImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReviewImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "Definition")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReviewImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReviewImageResponseParams struct {
	// 图片审核任务结果。
	// <font color=red>注意：该字段已废弃，建议使用 MediaReviewResult。</font> 
	ReviewResultSet []*ContentReviewResult `json:"ReviewResultSet,omitnil,omitempty" name:"ReviewResultSet"`

	// 图片审核任务结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaReviewResult *ReviewImageResult `json:"MediaReviewResult,omitnil,omitempty" name:"MediaReviewResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReviewImageResponse struct {
	*tchttp.BaseResponse
	Response *ReviewImageResponseParams `json:"Response"`
}

func (r *ReviewImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReviewImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviewImageResult struct {
	// 图片审核的结果建议，取值范围：
	// <li>pass：建议通过；</li>
	// <li>review：建议复审；</li>
	// <li>block：建议封禁。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 当 Suggestion 为 review 或 block 时有效，表示最可能的违规的标签，取值范围：
	// <li>Porn：色情；</li>
	// <li>Terror：暴力；</li>
	// <li>Polity：不适宜的信息；</li>
	// <li>Ad：广告；</li>
	// <li>Illegal：违法；</li>
	// <li>Abuse：谩骂。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 当 Suggestion 为 review 或 block 时有效，表示最可能的违禁的形式，取值范围：
	// <li>Image：画面上的人物或图标；</li>
	// <li>OCR：画面上的文字。</li>
	Form *string `json:"Form,omitnil,omitempty" name:"Form"`

	// 有违规信息的嫌疑的视频片段列表。
	// <font color=red>注意</font> ：该列表最多仅展示前 10个 元素。如希望获得完整结果，请从 SegmentSetFileUrl 对应的文件中获取。
	SegmentSet []*ReviewImageSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// 涉及违规信息的嫌疑的视频片段列表文件 URL。文件的内容为 JSON，数据结构与 SegmentSet 字段一致。 （文件不会永久存储，到达SegmentSetFileUrlExpireTime 时间点后文件将被删除）。
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitnil,omitempty" name:"SegmentSetFileUrl"`

	// 涉及违规信息的嫌疑的视频片段列表文件 URL 失效时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitnil,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type ReviewImageSegmentItem struct {
	// 嫌疑片段涉及令人反感的信息的分数。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 嫌疑片段鉴别涉及违规信息的结果建议，取值范围：
	// <li>review：疑似违规，建议复审；</li>
	// <li>block：确认违规，建议封禁。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 嫌疑片段最可能的违规的标签，取值范围：
	// <li>Porn：色情；</li>
	// <li>Terror：暴力；</li>
	// <li>Polity：不适宜的信息；</li>
	// <li>Ad：广告；</li>
	// <li>Illegal：违法；</li>
	// <li>Abuse：谩骂。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 违规子标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 嫌疑片段违禁的形式，取值范围：
	// <li>Image：画面上的人物或图标；</li>
	// <li>OCR：画面上的文字。</li>
	Form *string `json:"Form,omitnil,omitempty" name:"Form"`

	// 嫌疑人物、图标或文字出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// 当 Form 为 OCR 时有效，表示识别出来的 OCR 文本内容。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 当 Form 为 OCR 时有效，表示嫌疑片段命中的违规关键词列表。
	KeywordSet []*string `json:"KeywordSet,omitnil,omitempty" name:"KeywordSet"`
}

type ReviewInfo struct {
	// 审核模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 审核的结果建议，取值范围：
	// <li>pass：建议通过；</li>
	// <li>review：建议复审；</li>
	// <li>block：建议封禁。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 审核类型，当 Suggestion 为 review 或 block 时有效，格式为：Form.Label。
	// Form 表示违禁的形式，取值范围：
	// <li>Image：画面上的人物或图标；</li>
	// <li>OCR：画面上的文字；</li>
	// <li>ASR：语音中的文字；</li>
	// <li>Voice：声音。</li>
	// Label 表示违禁的标签，取值范围：
	// <li>Porn：色情；</li>
	// <li>Terror：暴力；</li>
	// <li>Polity：不适宜的信息；</li>
	// <li>Ad：广告；</li>
	// <li>Illegal：违法；</li>
	// <li>Abuse：谩骂；</li>
	// <li>Moan：娇喘。</li>
	TypeSet []*string `json:"TypeSet,omitnil,omitempty" name:"TypeSet"`

	// 审核时间，使用  [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	ReviewTime *string `json:"ReviewTime,omitnil,omitempty" name:"ReviewTime"`
}

type ReviewTemplate struct {
	// 审核模版唯一标签。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 模板类型，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 需要返回的违规标签列表。
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type RoundPlayFilePlayInfo struct {
	// 播单节目的 ID，由系统分配。
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// 媒体文件标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 启播时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	StartPlayTime *string `json:"StartPlayTime,omitnil,omitempty" name:"StartPlayTime"`

	// 播放时长，单位为秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 播放进度，单位为秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type RoundPlayInfo struct {
	// 轮播播单标识。
	RoundPlayId *string `json:"RoundPlayId,omitnil,omitempty" name:"RoundPlayId"`

	// 启播时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 轮播列表。
	RoundPlaylist []*RoundPlayListItemInfo `json:"RoundPlaylist,omitnil,omitempty" name:"RoundPlaylist"`

	// 轮播播单名称，长度限制：64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 轮播播单描述信息，长度限制：256 个字符。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 播放状态，可选值：
	// <li>Enabled：启动状态；</li>
	// <li>Disabled：停止状态。</li>
	// 默认值：Enabled。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 播放模式，可选值：
	// <li>Loop：循环播放播单；</li>
	// <li>Linear：单次播放，播单播放完停止播放。</li>
	// 默认值：Loop。
	PlayBackMode *string `json:"PlayBackMode,omitnil,omitempty" name:"PlayBackMode"`

	// 轮播播放地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 创建时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#52)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type RoundPlayListItemInfo struct {
	// 媒体文件标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 播放的音视频类型，可选值：
	// <li>Transcode：转码输出；转码输出会有多个模版，必须指定 Definition 字段</li>
	// <li>Original：原始音视频。</li>
	// Type 对应的格式必须为 HLS 格式。
	AudioVideoType *string `json:"AudioVideoType,omitnil,omitempty" name:"AudioVideoType"`

	// 播放节目的 ID，由系统分配。
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// 指定播放的转码模版，当 AudioVideoType 为 Transcode 时必须指定。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type SDMCDrmKeyProviderInfo struct {
	// 华曦达分配的用户 ID。最大长度为128个字符。
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 华曦达分配的用户密钥 ID。最大长度为128个字符。
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 华曦达分配的用户密钥内容。最大长度为128个字符。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 华曦达分配的 FairPlay 证书地址。该地址需使用 HTTPS 协议，最大长度为1024个字符。
	FairPlayCertificateUrl *string `json:"FairPlayCertificateUrl,omitnil,omitempty" name:"FairPlayCertificateUrl"`
}

type SampleSnapshotTaskInput struct {
	// 采样截图模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`
}

type SampleSnapshotTemplate struct {
	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 采样截图模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 采样截图类型。
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔。
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 填充方式，当截图配置宽高参数与原始视频的宽高比不一致时，对截图的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type ScratchRepairInfo struct {
	// 去划痕控制开关，可选值：
	// <li>ON：开启去划痕；</li>
	// <li>OFF：关闭去划痕。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 去划痕强度，仅当去划痕控制开关为 ON 时有效，取值范围：0.0~1.0。
	// 默认：0.0。
	Intensity *float64 `json:"Intensity,omitnil,omitempty" name:"Intensity"`

	// 去划痕类型，仅当去划痕控制开关为 ON 时有效，可选值：
	// <li>normal：正常去划痕；</li>
	// 默认值：normal。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type SearchMediaRequestParams struct {
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 文件 ID 集合，匹配集合中的任意元素。
	// <li>数组长度限制：10。</li>
	// <li>单个 ID 长度限制：40个字符。</li>
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 文件名集合，模糊匹配媒体文件的文件名，匹配度越高，排序越优先。
	// <li>单个文件名长度限制：100个字符。</li>
	// <li>数组长度限制：10。</li>
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 文件名前缀，前缀匹配媒体文件的文件名。
	// <li>单个文件名前缀长度限制：100个字符。</li>
	// <li>数组长度限制：10。</li>
	NamePrefixes []*string `json:"NamePrefixes,omitnil,omitempty" name:"NamePrefixes"`

	// 文件描述集合，模糊匹配媒体文件的描述，匹配度越高，排序越优先。
	// <li>单个描述长度限制：100个字符。</li>
	// <li>数组长度限制：10。</li>
	Descriptions []*string `json:"Descriptions,omitnil,omitempty" name:"Descriptions"`

	// 分类 ID 集合，匹配集合指定 ID 的分类及其所有子类。
	// <li>数组长度限制：10。</li>
	ClassIds []*int64 `json:"ClassIds,omitnil,omitempty" name:"ClassIds"`

	// 标签集合，匹配集合中任意元素。
	// <li>单个标签长度限制：32个字符。</li>
	// <li>数组长度限制：16。</li>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 文件类型。匹配集合中的任意元素：
	// <li>Video: 视频文件</li>
	// <li>Audio: 音频文件</li>
	// <li>Image: 图片文件</li>
	Categories []*string `json:"Categories,omitnil,omitempty" name:"Categories"`

	// 媒体文件来源集合，来源取值参见 [SourceType](https://cloud.tencent.com/document/product/266/31773#MediaSourceData)。
	// <li>数组长度限制：10。</li>
	SourceTypes []*string `json:"SourceTypes,omitnil,omitempty" name:"SourceTypes"`

	// 推流直播码集合。匹配集合中的任意元素。
	// <li>数组长度限制：10。</li>
	StreamIds []*string `json:"StreamIds,omitnil,omitempty" name:"StreamIds"`

	// 匹配创建时间在此时间段内的文件。
	// <li>包含所指定的头尾时间点。</li>
	CreateTime *TimeRange `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 匹配过期时间在此时间段内的文件，无法检索到已过期文件。
	// <li>包含所指定的头尾时间点。</li>
	ExpireTime *TimeRange `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 媒体文件存储地区，如 ap-chongqing，参见[地域列表](https://cloud.tencent.com/document/product/266/9760#.E5.B7.B2.E6.94.AF.E6.8C.81.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)。
	// <li>单个存储地区长度限制：20个字符。</li>
	// <li>数组长度限制：20。</li>
	StorageRegions []*string `json:"StorageRegions,omitnil,omitempty" name:"StorageRegions"`

	// 存储类型数组。可选值有：
	// <li> STANDARD：标准存储。</li>
	// <li> STANDARD_IA：低频存储。</li>
	// <li> ARCHIVE：归档存储。</li>
	// <li> DEEP_ARCHIVE：深度归档存储。</li>
	StorageClasses []*string `json:"StorageClasses,omitnil,omitempty" name:"StorageClasses"`

	// 媒体文件封装格式集合，匹配集合中任意元素。
	// <li>数组长度限制：10。</li>
	MediaTypes []*string `json:"MediaTypes,omitnil,omitempty" name:"MediaTypes"`

	// 媒体文件状态，匹配集合中任意元素。
	// <li> Normal：正常；</li>
	// <li> SystemForbidden：平台封禁；</li>
	// <li> Forbidden：主动封禁。</li>
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 媒体文件审核结果，匹配集合中任意元素。
	// <li> pass：审核通过；</li>
	// <li> review：疑似违规，建议复审；</li>
	// <li> block：确认违规，建议封禁；</li>
	// <li> notModerated：未审核。</li>
	ReviewResults []*string `json:"ReviewResults,omitnil,omitempty" name:"ReviewResults"`

	// TRTC 应用 ID 集合。匹配集合中的任意元素。
	// <li>数组长度限制：10。</li>
	TrtcSdkAppIds []*uint64 `json:"TrtcSdkAppIds,omitnil,omitempty" name:"TrtcSdkAppIds"`

	// TRTC 房间 ID 集合。匹配集合中的任意元素。
	// <li>单个房间 ID 长度限制：64个字符；</li>
	// <li>数组长度限制：10。</li>
	TrtcRoomIds []*string `json:"TrtcRoomIds,omitnil,omitempty" name:"TrtcRoomIds"`

	// 指定所有媒体文件需要返回的信息，可同时指定多个信息，N 从 0 开始递增。如果未填写该字段，默认返回所有信息。选项有：
	// <li>basicInfo（视频基础信息）。</li>
	// <li>metaData（视频元信息）。</li>
	// <li>transcodeInfo（视频转码结果信息）。</li>
	// <li>animatedGraphicsInfo（视频转动图结果信息）。</li>
	// <li>imageSpriteInfo（视频雪碧图信息）。</li>
	// <li>snapshotByTimeOffsetInfo（视频指定时间点截图信息）。</li>
	// <li>sampleSnapshotInfo（采样截图信息）。</li>
	// <li>keyFrameDescInfo（打点信息）。</li>
	// <li>adaptiveDynamicStreamingInfo（转自适应码流信息）。</li>
	// <li>miniProgramReviewInfo（小程序审核信息）。</li>
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方式。
	// <li>Sort.Field 可选 CreateTime 。</li>
	// <li>当 Text、 Names 或 Descriptions 不为空时，Sort.Field 字段无效， 搜索结果将以匹配度排序。</li>
	Sort *SortBy `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <div id="p_offset">分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	// <li>取值范围：Offset + Limit 不超过5000。（参见：<a href="#maxResultsDesc">接口返回结果数限制</a>）</li></div>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <div id="p_limit">分页返回的记录条数，默认值：10。将返回第 Offset 到第 Offset+Limit-1 条。
	// <li>取值范围：Offset + Limit 不超过5000。（参见：<a href="#maxResultsDesc">接口返回结果数限制</a>）</li></div>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// （不推荐：应使用 Names、NamePrefixes 或 Descriptions 替代）
	// 搜索文本，模糊匹配媒体文件名称或描述信息，匹配项越多，匹配度越高，排序越优先。长度限制：64个字符。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// （不推荐：应使用 SourceTypes 替代）
	// 媒体文件来源，来源取值参见 [SourceType](https://cloud.tencent.com/document/product/266/31773#MediaSourceData)。
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// （不推荐：应使用 StreamIds 替代）
	// 推流直播码。
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// （不推荐：应使用 CreateTime 替代）
	// 创建时间的开始时间。
	// <li>大于等于开始时间。</li>
	// <li>当 CreateTime.After 也存在时，将优先使用 CreateTime.After。</li>
	// <li>格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。</li>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// （不推荐：应使用 CreateTime 替代）
	// 创建时间的结束时间。
	// <li>小于结束时间。</li>
	// <li>当 CreateTime.Before 也存在时，将优先使用 CreateTime.Before。</li>
	// <li>格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。</li>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 该字段已无效。
	Vids []*string `json:"Vids,omitnil,omitempty" name:"Vids"`

	// 该字段已无效。
	Vid *string `json:"Vid,omitnil,omitempty" name:"Vid"`
}

type SearchMediaRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 文件 ID 集合，匹配集合中的任意元素。
	// <li>数组长度限制：10。</li>
	// <li>单个 ID 长度限制：40个字符。</li>
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 文件名集合，模糊匹配媒体文件的文件名，匹配度越高，排序越优先。
	// <li>单个文件名长度限制：100个字符。</li>
	// <li>数组长度限制：10。</li>
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 文件名前缀，前缀匹配媒体文件的文件名。
	// <li>单个文件名前缀长度限制：100个字符。</li>
	// <li>数组长度限制：10。</li>
	NamePrefixes []*string `json:"NamePrefixes,omitnil,omitempty" name:"NamePrefixes"`

	// 文件描述集合，模糊匹配媒体文件的描述，匹配度越高，排序越优先。
	// <li>单个描述长度限制：100个字符。</li>
	// <li>数组长度限制：10。</li>
	Descriptions []*string `json:"Descriptions,omitnil,omitempty" name:"Descriptions"`

	// 分类 ID 集合，匹配集合指定 ID 的分类及其所有子类。
	// <li>数组长度限制：10。</li>
	ClassIds []*int64 `json:"ClassIds,omitnil,omitempty" name:"ClassIds"`

	// 标签集合，匹配集合中任意元素。
	// <li>单个标签长度限制：32个字符。</li>
	// <li>数组长度限制：16。</li>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 文件类型。匹配集合中的任意元素：
	// <li>Video: 视频文件</li>
	// <li>Audio: 音频文件</li>
	// <li>Image: 图片文件</li>
	Categories []*string `json:"Categories,omitnil,omitempty" name:"Categories"`

	// 媒体文件来源集合，来源取值参见 [SourceType](https://cloud.tencent.com/document/product/266/31773#MediaSourceData)。
	// <li>数组长度限制：10。</li>
	SourceTypes []*string `json:"SourceTypes,omitnil,omitempty" name:"SourceTypes"`

	// 推流直播码集合。匹配集合中的任意元素。
	// <li>数组长度限制：10。</li>
	StreamIds []*string `json:"StreamIds,omitnil,omitempty" name:"StreamIds"`

	// 匹配创建时间在此时间段内的文件。
	// <li>包含所指定的头尾时间点。</li>
	CreateTime *TimeRange `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 匹配过期时间在此时间段内的文件，无法检索到已过期文件。
	// <li>包含所指定的头尾时间点。</li>
	ExpireTime *TimeRange `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 媒体文件存储地区，如 ap-chongqing，参见[地域列表](https://cloud.tencent.com/document/product/266/9760#.E5.B7.B2.E6.94.AF.E6.8C.81.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)。
	// <li>单个存储地区长度限制：20个字符。</li>
	// <li>数组长度限制：20。</li>
	StorageRegions []*string `json:"StorageRegions,omitnil,omitempty" name:"StorageRegions"`

	// 存储类型数组。可选值有：
	// <li> STANDARD：标准存储。</li>
	// <li> STANDARD_IA：低频存储。</li>
	// <li> ARCHIVE：归档存储。</li>
	// <li> DEEP_ARCHIVE：深度归档存储。</li>
	StorageClasses []*string `json:"StorageClasses,omitnil,omitempty" name:"StorageClasses"`

	// 媒体文件封装格式集合，匹配集合中任意元素。
	// <li>数组长度限制：10。</li>
	MediaTypes []*string `json:"MediaTypes,omitnil,omitempty" name:"MediaTypes"`

	// 媒体文件状态，匹配集合中任意元素。
	// <li> Normal：正常；</li>
	// <li> SystemForbidden：平台封禁；</li>
	// <li> Forbidden：主动封禁。</li>
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 媒体文件审核结果，匹配集合中任意元素。
	// <li> pass：审核通过；</li>
	// <li> review：疑似违规，建议复审；</li>
	// <li> block：确认违规，建议封禁；</li>
	// <li> notModerated：未审核。</li>
	ReviewResults []*string `json:"ReviewResults,omitnil,omitempty" name:"ReviewResults"`

	// TRTC 应用 ID 集合。匹配集合中的任意元素。
	// <li>数组长度限制：10。</li>
	TrtcSdkAppIds []*uint64 `json:"TrtcSdkAppIds,omitnil,omitempty" name:"TrtcSdkAppIds"`

	// TRTC 房间 ID 集合。匹配集合中的任意元素。
	// <li>单个房间 ID 长度限制：64个字符；</li>
	// <li>数组长度限制：10。</li>
	TrtcRoomIds []*string `json:"TrtcRoomIds,omitnil,omitempty" name:"TrtcRoomIds"`

	// 指定所有媒体文件需要返回的信息，可同时指定多个信息，N 从 0 开始递增。如果未填写该字段，默认返回所有信息。选项有：
	// <li>basicInfo（视频基础信息）。</li>
	// <li>metaData（视频元信息）。</li>
	// <li>transcodeInfo（视频转码结果信息）。</li>
	// <li>animatedGraphicsInfo（视频转动图结果信息）。</li>
	// <li>imageSpriteInfo（视频雪碧图信息）。</li>
	// <li>snapshotByTimeOffsetInfo（视频指定时间点截图信息）。</li>
	// <li>sampleSnapshotInfo（采样截图信息）。</li>
	// <li>keyFrameDescInfo（打点信息）。</li>
	// <li>adaptiveDynamicStreamingInfo（转自适应码流信息）。</li>
	// <li>miniProgramReviewInfo（小程序审核信息）。</li>
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方式。
	// <li>Sort.Field 可选 CreateTime 。</li>
	// <li>当 Text、 Names 或 Descriptions 不为空时，Sort.Field 字段无效， 搜索结果将以匹配度排序。</li>
	Sort *SortBy `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <div id="p_offset">分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	// <li>取值范围：Offset + Limit 不超过5000。（参见：<a href="#maxResultsDesc">接口返回结果数限制</a>）</li></div>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <div id="p_limit">分页返回的记录条数，默认值：10。将返回第 Offset 到第 Offset+Limit-1 条。
	// <li>取值范围：Offset + Limit 不超过5000。（参见：<a href="#maxResultsDesc">接口返回结果数限制</a>）</li></div>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// （不推荐：应使用 Names、NamePrefixes 或 Descriptions 替代）
	// 搜索文本，模糊匹配媒体文件名称或描述信息，匹配项越多，匹配度越高，排序越优先。长度限制：64个字符。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// （不推荐：应使用 SourceTypes 替代）
	// 媒体文件来源，来源取值参见 [SourceType](https://cloud.tencent.com/document/product/266/31773#MediaSourceData)。
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// （不推荐：应使用 StreamIds 替代）
	// 推流直播码。
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// （不推荐：应使用 CreateTime 替代）
	// 创建时间的开始时间。
	// <li>大于等于开始时间。</li>
	// <li>当 CreateTime.After 也存在时，将优先使用 CreateTime.After。</li>
	// <li>格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。</li>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// （不推荐：应使用 CreateTime 替代）
	// 创建时间的结束时间。
	// <li>小于结束时间。</li>
	// <li>当 CreateTime.Before 也存在时，将优先使用 CreateTime.Before。</li>
	// <li>格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。</li>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 该字段已无效。
	Vids []*string `json:"Vids,omitnil,omitempty" name:"Vids"`

	// 该字段已无效。
	Vid *string `json:"Vid,omitnil,omitempty" name:"Vid"`
}

func (r *SearchMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "FileIds")
	delete(f, "Names")
	delete(f, "NamePrefixes")
	delete(f, "Descriptions")
	delete(f, "ClassIds")
	delete(f, "Tags")
	delete(f, "Categories")
	delete(f, "SourceTypes")
	delete(f, "StreamIds")
	delete(f, "CreateTime")
	delete(f, "ExpireTime")
	delete(f, "StorageRegions")
	delete(f, "StorageClasses")
	delete(f, "MediaTypes")
	delete(f, "Status")
	delete(f, "ReviewResults")
	delete(f, "TrtcSdkAppIds")
	delete(f, "TrtcRoomIds")
	delete(f, "Filters")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Text")
	delete(f, "SourceType")
	delete(f, "StreamId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Vids")
	delete(f, "Vid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchMediaResponseParams struct {
	// 符合搜索条件的记录总数。
	// <li>最大值：5000。当命中记录数超过5000时，该字段将返回 5000，而非实际命中总数。</li>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 媒体文件信息列表。
	MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitnil,omitempty" name:"MediaInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchMediaResponse struct {
	*tchttp.BaseResponse
	Response *SearchMediaResponseParams `json:"Response"`
}

func (r *SearchMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SegmentConfigureInfo struct {
	// 视频拆条识别任务开关，可选值：
	// <li>ON：开启智能视频拆条识别任务；</li>
	// <li>OFF：关闭智能视频拆条识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type SegmentConfigureInfoForUpdate struct {
	// 视频拆条识别任务开关，可选值：
	// <li>ON：开启智能视频拆条识别任务；</li>
	// <li>OFF：关闭智能视频拆条识别任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type SetCLSPushTargetRequestParams struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 要设置的中国大陆地区的日志推送目标。
	ChineseMainlandCLSTargetInfo *AreaCLSTargetInfo `json:"ChineseMainlandCLSTargetInfo,omitnil,omitempty" name:"ChineseMainlandCLSTargetInfo"`

	// 要设置的中国大陆以外地区的日志推送目标。
	OutsideChineseMainlandCLSTargetInfo *AreaCLSTargetInfo `json:"OutsideChineseMainlandCLSTargetInfo,omitnil,omitempty" name:"OutsideChineseMainlandCLSTargetInfo"`
}

type SetCLSPushTargetRequest struct {
	*tchttp.BaseRequest
	
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *int64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 要设置的中国大陆地区的日志推送目标。
	ChineseMainlandCLSTargetInfo *AreaCLSTargetInfo `json:"ChineseMainlandCLSTargetInfo,omitnil,omitempty" name:"ChineseMainlandCLSTargetInfo"`

	// 要设置的中国大陆以外地区的日志推送目标。
	OutsideChineseMainlandCLSTargetInfo *AreaCLSTargetInfo `json:"OutsideChineseMainlandCLSTargetInfo,omitnil,omitempty" name:"OutsideChineseMainlandCLSTargetInfo"`
}

func (r *SetCLSPushTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCLSPushTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SubAppId")
	delete(f, "ChineseMainlandCLSTargetInfo")
	delete(f, "OutsideChineseMainlandCLSTargetInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCLSPushTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCLSPushTargetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetCLSPushTargetResponse struct {
	*tchttp.BaseResponse
	Response *SetCLSPushTargetResponseParams `json:"Response"`
}

func (r *SetCLSPushTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCLSPushTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDrmKeyProviderInfoRequestParams struct {
	// 华曦达（SDMC）相关的 DRM 密钥提供商信息。
	SDMCInfo *SDMCDrmKeyProviderInfo `json:"SDMCInfo,omitnil,omitempty" name:"SDMCInfo"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

type SetDrmKeyProviderInfoRequest struct {
	*tchttp.BaseRequest
	
	// 华曦达（SDMC）相关的 DRM 密钥提供商信息。
	SDMCInfo *SDMCDrmKeyProviderInfo `json:"SDMCInfo,omitnil,omitempty" name:"SDMCInfo"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`
}

func (r *SetDrmKeyProviderInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDrmKeyProviderInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SDMCInfo")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetDrmKeyProviderInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDrmKeyProviderInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetDrmKeyProviderInfoResponse struct {
	*tchttp.BaseResponse
	Response *SetDrmKeyProviderInfoResponseParams `json:"Response"`
}

func (r *SetDrmKeyProviderInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDrmKeyProviderInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVodDomainCertificateRequestParams struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 本次接口调用的操作类型，可选值：
	// <li>Set：设置域名的 HTTPS 所需证书；</li>
	// <li>Clear: 清空域名 HTTPS 配置，清空后该域名无法支持 HTTPS 访问。</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// [腾讯云SSL](/document/product/400/7572) 上传的证书 ID。清空域名 HTTPS 配置时无需填写该字段。
	CertID *string `json:"CertID,omitnil,omitempty" name:"CertID"`
}

type SetVodDomainCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 本次接口调用的操作类型，可选值：
	// <li>Set：设置域名的 HTTPS 所需证书；</li>
	// <li>Clear: 清空域名 HTTPS 配置，清空后该域名无法支持 HTTPS 访问。</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// [腾讯云SSL](/document/product/400/7572) 上传的证书 ID。清空域名 HTTPS 配置时无需填写该字段。
	CertID *string `json:"CertID,omitnil,omitempty" name:"CertID"`
}

func (r *SetVodDomainCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVodDomainCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Operation")
	delete(f, "SubAppId")
	delete(f, "CertID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetVodDomainCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVodDomainCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetVodDomainCertificateResponse struct {
	*tchttp.BaseResponse
	Response *SetVodDomainCertificateResponseParams `json:"Response"`
}

func (r *SetVodDomainCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVodDomainCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SharpEnhanceInfo struct {
	// 细节增强控制开关，可选值：
	// <li>ON：开启细节增强；</li>
	// <li>OFF：关闭细节增强。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 细节增强强度，仅当细节增强控制开关为 ON 时有效，取值范围：0.0~1.0。
	// 默认：0.0。
	Intensity *float64 `json:"Intensity,omitnil,omitempty" name:"Intensity"`
}

type SimpleAesEdkPair struct {
	// 加密后的数据密钥。
	Edk *string `json:"Edk,omitnil,omitempty" name:"Edk"`

	// 数据密钥。返回的数据密钥 DK 为 Base64 编码字符串。
	Dk *string `json:"Dk,omitnil,omitempty" name:"Dk"`
}

// Predefined struct for user
type SimpleHlsClipRequestParams struct {
	// 需要裁剪的腾讯云点播 HLS 视频 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 裁剪的开始偏移时间，单位秒。默认 0，即从视频开头开始裁剪。负数表示距离视频结束多少秒开始裁剪。例如 -10 表示从倒数第 10 秒开始裁剪。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 裁剪的结束偏移时间，单位秒。默认 0，即裁剪到视频尾部。负数表示距离视频结束多少秒结束裁剪。例如 -10 表示到倒数第 10 秒结束裁剪。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 是否固化。0 不固化，1 固化。默认不固化。
	IsPersistence *int64 `json:"IsPersistence,omitnil,omitempty" name:"IsPersistence"`

	// 剪辑固化后的视频存储过期时间。格式参照 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。仅 IsPersistence 为 1 时有效，默认剪辑固化的视频永不过期。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 剪辑固化后的视频点播任务流处理，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。仅 IsPersistence 为 1 且 Precision 为 Rough 时有效。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	// 仅 IsPersistence 为 1 时有效。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，[上传完成回调](/document/product/266/7830) 将返回该字段值，最长 250 个字符。仅 IsPersistence 为 1 时有效。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`

	// 会话上下文，用于透传用户请求信息，当指定 Procedure 参数后，[任务流状态变更回调](/document/product/266/9636) 将返回该字段值，最长 1000 个字符。仅 IsPersistence 为 1 时有效。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 裁剪精度，取值有：<li>Rough: 粗略裁剪，最小剪辑精度是单个 ts 分片；</li><li>Precise: 精确裁剪，做到按照剪辑时间点的毫秒级精确剪辑。</li> 默认取值 Rough。
	Precision *string `json:"Precision,omitnil,omitempty" name:"Precision"`

	// 输出视频类型，取值有：<li>hls: 输出 hls 文件；</li><li>mp4：输出 mp4 文件，MP4 文件的大小不超过5G，时长小于2小时。仅当 Precision 选择 Precise 且 IsPersistence  选择0时有效，即只有非固化的精确剪辑时支持输出 MP4。</li>默认取值 hls。
	OutputMediaType *string `json:"OutputMediaType,omitnil,omitempty" name:"OutputMediaType"`

	// 保留字段，特殊用途时使用。 示例值：""
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type SimpleHlsClipRequest struct {
	*tchttp.BaseRequest
	
	// 需要裁剪的腾讯云点播 HLS 视频 URL。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 裁剪的开始偏移时间，单位秒。默认 0，即从视频开头开始裁剪。负数表示距离视频结束多少秒开始裁剪。例如 -10 表示从倒数第 10 秒开始裁剪。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 裁剪的结束偏移时间，单位秒。默认 0，即裁剪到视频尾部。负数表示距离视频结束多少秒结束裁剪。例如 -10 表示到倒数第 10 秒结束裁剪。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 是否固化。0 不固化，1 固化。默认不固化。
	IsPersistence *int64 `json:"IsPersistence,omitnil,omitempty" name:"IsPersistence"`

	// 剪辑固化后的视频存储过期时间。格式参照 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。仅 IsPersistence 为 1 时有效，默认剪辑固化的视频永不过期。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 剪辑固化后的视频点播任务流处理，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。仅 IsPersistence 为 1 且 Precision 为 Rough 时有效。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	// 仅 IsPersistence 为 1 时有效。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，[上传完成回调](/document/product/266/7830) 将返回该字段值，最长 250 个字符。仅 IsPersistence 为 1 时有效。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`

	// 会话上下文，用于透传用户请求信息，当指定 Procedure 参数后，[任务流状态变更回调](/document/product/266/9636) 将返回该字段值，最长 1000 个字符。仅 IsPersistence 为 1 时有效。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 裁剪精度，取值有：<li>Rough: 粗略裁剪，最小剪辑精度是单个 ts 分片；</li><li>Precise: 精确裁剪，做到按照剪辑时间点的毫秒级精确剪辑。</li> 默认取值 Rough。
	Precision *string `json:"Precision,omitnil,omitempty" name:"Precision"`

	// 输出视频类型，取值有：<li>hls: 输出 hls 文件；</li><li>mp4：输出 mp4 文件，MP4 文件的大小不超过5G，时长小于2小时。仅当 Precision 选择 Precise 且 IsPersistence  选择0时有效，即只有非固化的精确剪辑时支持输出 MP4。</li>默认取值 hls。
	OutputMediaType *string `json:"OutputMediaType,omitnil,omitempty" name:"OutputMediaType"`

	// 保留字段，特殊用途时使用。 示例值：""
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

func (r *SimpleHlsClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SimpleHlsClipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "SubAppId")
	delete(f, "StartTimeOffset")
	delete(f, "EndTimeOffset")
	delete(f, "IsPersistence")
	delete(f, "ExpireTime")
	delete(f, "Procedure")
	delete(f, "ClassId")
	delete(f, "SourceContext")
	delete(f, "SessionContext")
	delete(f, "Precision")
	delete(f, "OutputMediaType")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SimpleHlsClipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SimpleHlsClipResponseParams struct {
	// 裁剪后的视频地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 裁剪后的视频元信息。目前`Size`，`Rotate`，`VideoDuration`，`AudioDuration` 几个字段暂时缺省，没有真实数据。
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 剪辑固化后的视频的媒体文件的唯一标识。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 剪辑固化后的视频任务流 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SimpleHlsClipResponse struct {
	*tchttp.BaseResponse
	Response *SimpleHlsClipResponseParams `json:"Response"`
}

func (r *SimpleHlsClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SimpleHlsClipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotByTimeOffset2017 struct {
	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 截图的具体时间点，单位：毫秒。
	TimeOffset *uint64 `json:"TimeOffset,omitnil,omitempty" name:"TimeOffset"`

	// 截图输出文件地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type SnapshotByTimeOffsetTask2017 struct {
	// 截图任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 截图文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 截图结果信息。
	SnapshotInfoSet []*SnapshotByTimeOffset2017 `json:"SnapshotInfoSet,omitnil,omitempty" name:"SnapshotInfoSet"`
}

type SnapshotByTimeOffsetTaskInput struct {
	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 截图时间点列表，时间点支持 s、% 两种格式：
	// <li>当字符串以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// <li>当字符串以 % 结尾，表示时间点为视频时长的百分比大小，如10%表示时间点为视频前第10%的时间。</li>
	ExtTimeOffsetSet []*string `json:"ExtTimeOffsetSet,omitnil,omitempty" name:"ExtTimeOffsetSet"`

	// 截图时间点列表，单位为<font color=red>毫秒</font>。此参数已不再建议使用，建议您使用 ExtTimeOffsetSet 参数。
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitnil,omitempty" name:"TimeOffsetSet"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`
}

type SnapshotByTimeOffsetTemplate struct {
	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 指定时间点截图模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 截图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 图片格式。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 填充方式，当截图配置宽高参数与原始视频的宽高比不一致时，对截图的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type SortBy struct {
	// 排序字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 排序方式，可选值：Asc（升序）、Desc（降序）
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type SpecificationDataItem struct {
	// 任务规格。
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 统计数据。
	Data []*TaskStatDataItem `json:"Data,omitnil,omitempty" name:"Data"`
}

type SplitMediaOutputConfig struct {
	// 输出文件名，最长 64 个字符。缺省由系统指定生成文件名。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 输出文件格式，可选值：mp4、hls。默认是 mp4。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type SplitMediaRequestParams struct {
	// 视频的 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 视频拆条任务信息列表，最多同时支持100个拆条信息。
	Segments []*SplitMediaTaskConfig `json:"Segments,omitnil,omitempty" name:"Segments"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 标识来源上下文，用于透传用户请求信息，在 SplitMediaComplete 回调和任务流状态变更回调将返回该字段值，最长 1000个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`
}

type SplitMediaRequest struct {
	*tchttp.BaseRequest
	
	// 视频的 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 视频拆条任务信息列表，最多同时支持100个拆条信息。
	Segments []*SplitMediaTaskConfig `json:"Segments,omitnil,omitempty" name:"Segments"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 标识来源上下文，用于透传用户请求信息，在 SplitMediaComplete 回调和任务流状态变更回调将返回该字段值，最长 1000个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于任务去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`
}

func (r *SplitMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "Segments")
	delete(f, "SubAppId")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "TasksPriority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SplitMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SplitMediaResponseParams struct {
	// 视频拆条的任务 ID，可以通过该 ID 查询拆条任务（任务类型为 SplitMedia）的状态。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SplitMediaResponse struct {
	*tchttp.BaseResponse
	Response *SplitMediaResponseParams `json:"Response"`
}

func (r *SplitMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitMediaTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/266/50368#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 视频拆条任务详细信息列表。
	FileInfoSet []*SplitMediaTaskSegmentInfo `json:"FileInfoSet,omitnil,omitempty" name:"FileInfoSet"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 视频拆条任务进度，取值范围 [0-100] 。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type SplitMediaTaskConfig struct {
	// 视频拆条起始的偏移时间，单位：秒。
	// <li>不填或填0，表示转码后的视频从原始视频的起始位置开始；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频从原始视频的第 n 秒位置开始；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频从原始视频结束 n 秒前的位置开始。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 视频拆条结束的偏移时间，单位：秒。
	// <li>不填或填0，表示转码后的视频持续到原始视频的末尾终止；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频持续到原始视频第 n 秒时终止；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频持续到原始视频结束 n 秒前终止。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// [任务流模板](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF)名字，如果要对生成的新视频执行任务流时填写。
	ProcedureName *string `json:"ProcedureName,omitnil,omitempty" name:"ProcedureName"`

	// 视频拆条输出信息。
	OutputConfig *SplitMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`
}

type SplitMediaTaskInput struct {
	// 视频的 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 视频拆条起始的偏移时间，单位：秒。
	// <li>不填或填0，表示转码后的视频从原始视频的起始位置开始；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频从原始视频的第 n 秒位置开始；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频从原始视频结束 n 秒前的位置开始。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 视频拆条结束的偏移时间，单位：秒。
	// <li>不填或填0，表示转码后的视频持续到原始视频的末尾终止；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频持续到原始视频第 n 秒时终止；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频持续到原始视频结束 n 秒前终止。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// [任务流模板](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF)名字，如果要对生成的新视频执行任务流时填写。
	ProcedureName *string `json:"ProcedureName,omitnil,omitempty" name:"ProcedureName"`

	// 视频拆条输出信息。
	OutputConfig *SplitMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`
}

type SplitMediaTaskSegmentInfo struct {
	// 视频拆条任务输入信息。
	Input *SplitMediaTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 视频拆条任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *TaskOutputMediaInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// 任务类型为 Procedure 的任务 ID。若发起[视频拆条](https://cloud.tencent.com/document/api/266/51098)任务时，视频拆条任务信息列表指定了任务流模板(ProcedureName)，当该任务流模板指定了 MediaProcessTask、AiAnalysisTask、AiRecognitionTask 中的一个或多个时发起该任务。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitnil,omitempty" name:"ProcedureTaskId"`

	// 任务类型为 ReviewAudioVideo 的任务 ID。若发起[视频拆条](https://cloud.tencent.com/document/api/266/51098)任务时，视频拆条任务信息列表指定了任务流模板(ProcedureName)，当该任务流模板指定了 ReviewAudioVideoTask 时，发起该任务。
	ReviewAudioVideoTaskId *string `json:"ReviewAudioVideoTaskId,omitnil,omitempty" name:"ReviewAudioVideoTaskId"`
}

type StatDataItem struct {
	// 数据所在时间区间的开始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。如：当时间粒度为天，2018-12-01T00:00:00+08:00，表示2018年12月1日（含）到2018年12月2日（不含）区间。
	// <li>表示小时级别数据时，2019-08-22T00:00:00+08:00表示2019-08-22日0点到1点的统计数据。</li>
	// <li>表示天级别数据时，2019-08-22T00:00:00+08:00表示2019-08-22日的统计数据。</li>
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 数据大小。
	// <li>存储空间的数据，单位是字节。</li>
	// <li>转码时长的数据，单位是秒。</li>
	// <li>流量数据，单位是字节。</li>
	// <li>带宽数据，单位是比特每秒。</li>
	// <li>直播剪辑数据，单位是秒。</li>
	// <li>轮播数据，单位是秒。</li>
	// <li>子应用数量数据，单位是个。</li>
	// <li>日志服务用量数据，单位是条。</li>
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type StickerTrackItem struct {
	// 贴图片段的媒体素材来源，可以是：
	// <li>点播的媒体文件 ID；</li>
	// <li>其他媒体文件的下载 URL。</li>
	// 注意：当使用其他媒体文件的下载 URL 作为素材来源，且开启了访问控制（如防盗链）时，需要在 URL 携带访问控制参数（如防盗链签名）。
	SourceMedia *string `json:"SourceMedia,omitnil,omitempty" name:"SourceMedia"`

	// 贴图的持续时间，单位为秒。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 贴图在轨道上的起始时间，单位为秒。
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 原点位置，取值有：
	// <li>Center：坐标原点为中心位置，如画布中心。</li>
	// 默认值：Center。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// 贴图原点距离画布原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示贴图 XPos 为画布宽度指定百分比的位置，如 10% 表示 XPos 为画布宽度的 10%。</li><li>当字符串以 px 结尾，表示贴图 XPos 单位为像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 贴图原点距离画布原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示贴图 YPos 为画布高度指定百分比的位置，如 10% 表示 YPos 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示贴图 YPos 单位为像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// 贴图的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示贴图 Width 为画布宽度的百分比大小，如 10% 表示 Width 为画布宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示贴图 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取贴图素材本身的 Width、Height。</li>
	// <li>当 Width 为空0，Height 非空，则 Width 按比例缩放</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 贴图的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示贴图 Height 为画布高度的百分比大小，如 10% 表示 Height 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示贴图 Height 单位为像素，如 100px 表示 Hieght 为 100 像素。</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取贴图素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按比例缩放</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 对贴图进行的操作，如图像旋转等。
	ImageOperations []*ImageTransform `json:"ImageOperations,omitnil,omitempty" name:"ImageOperations"`
}

type StorageRegionInfo struct {
	// 存储地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 存储地域描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 状态，是否开通，取值有：
	// <li>opened：已经开通。</li>
	// <li>unopened：未开通。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否默认的存储地域，true：是；false：否。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 存储区域，取值有：
	// <li>Chinese Mainland：中国境内（不包含港澳台）。</li>
	// <li>Outside Chinese Mainland：中国境外。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type StorageStatData struct {
	// 点播存储的计费区域，可能值：
	// <li>Chinese Mainland：中国境内（不包含港澳台）。</li>
	// <li>Outside Chinese Mainland：中国境外。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 当前总存储量，单位是字节。
	TotalStorage *uint64 `json:"TotalStorage,omitnil,omitempty" name:"TotalStorage"`

	// 当前低频存储量，单位是字节。
	InfrequentStorage *uint64 `json:"InfrequentStorage,omitnil,omitempty" name:"InfrequentStorage"`

	// 当前标准存储量，单位是字节。
	StandardStorage *uint64 `json:"StandardStorage,omitnil,omitempty" name:"StandardStorage"`

	// 当前归档存储量，单位是字节。
	ArchiveStorage *uint64 `json:"ArchiveStorage,omitnil,omitempty" name:"ArchiveStorage"`

	// 当前深度归档存储量，单位是字节。
	DeepArchiveStorage *uint64 `json:"DeepArchiveStorage,omitnil,omitempty" name:"DeepArchiveStorage"`
}

type SubAppIdInfo struct {
	// 子应用 ID。
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 子应用名称。
	SubAppIdName *string `json:"SubAppIdName,omitnil,omitempty" name:"SubAppIdName"`

	// 子应用简介。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 子应用创建时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 子应用状态，有效值：
	// <li>On：启用；</li>
	// <li>Off：停用。</li>
	// <li>Destroying：销毁中。</li>
	// <li>Destroyed：销毁完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 子应用名称（该字段已不推荐使用，建议使用新的子应用名称字段 SubAppIdName）。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type SubtitleFormatsOperation struct {
	// 操作类型，取值范围：
	// <li>add：添加 Formats 指定的格式列表；</li>
	// <li>delete：删除 Formats 指定的格式列表；<l/i>
	// <li>reset：将已配置的格式列表重置为  Formats 指定的格式列表。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 字幕格式列表，取值范围：
	// <li>vtt：生成 WebVTT 字幕文件；</li>
	// <li>srt：生成 SRT 字幕文件。</li>
	Formats []*string `json:"Formats,omitnil,omitempty" name:"Formats"`
}

type SuperResolutionInfo struct {
	// 画面超分控制开关，可选值：
	// <li>ON：开启画面超分；</li>
	// <li>OFF：关闭画面超分。</li>
	// 当开启画面超分时，默认2倍超分。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 画面超分类型，仅当画面超分控制开关为 ON 时有效，可选值：
	// <li>lq：针对低清晰度有较多噪声视频的超分；</li>
	// <li>hq：针对高清晰度视频超分。</li>
	// 默认值：lq。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 超分倍数，可选值：2。
	// 默认值：2。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
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
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印的高度，支持 px，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素；当填 0px 且
	//  Width 不为 0px 时，表示水印的高度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的高度取原始 SVG 图像的高度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Height 为视频宽度的百分比大小，如 10W% 表示 Height 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Height 为视频高度的百分比大小，如 10H% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Height 为视频短边的百分比大小，如 10S% 表示 Height 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Height 为视频长边的百分比大小，如 10L% 表示 Height 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 H%。</li>
	// 默认值为 0px。
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`
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
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印的高度，支持 px，%，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素；当填 0px 且
	//  Width 不为 0px 时，表示水印的高度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的高度取原始 SVG 图像的高度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Height 为视频宽度的百分比大小，如 10W% 表示 Height 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Height 为视频高度的百分比大小，如 10H% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Height 为视频短边的百分比大小，如 10S% 表示 Height 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Height 为视频长边的百分比大小，如 10L% 表示 Height 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 H%。
	// 默认值为 0px。
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印周期配置，用于配置水印周期性地显示与隐藏。
	// 主要使用场景是：为了视频防遮标，在视频多个地方设置水印，这些水印按固定顺序周期性地显示与隐藏。
	// 例如，设置 A、B、C、D 4 个水印分别位于视频的左上角、右上角、右下角、左下角处，视频开始时，{ A 显示 5 秒 -> B 显示 5 秒 -> C 显示 5 秒 -> D 显示 5 秒 } -> A 显示 5 秒 -> B 显示 5 秒 -> ...，任何时刻只显示一处水印。
	// 花括号 {} 表示由 A、B、C、D 4 个水印组成的大周期，可以看出每个大周期持续 20 秒。
	// 可以看出，A、B、C、D 都是周期性地显示 5 秒、隐藏 15 秒，且四者有固定的显示顺序。
	// 此配置项即用来描述单个水印的周期配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleConfig *WatermarkCycleConfigForUpdate `json:"CycleConfig,omitnil,omitempty" name:"CycleConfig"`
}

type TEHDConfig struct {
	// 极速高清类型，可选值：<li>TEHD-100 表示极速高清-100;</li> <li>OFF 表示关闭极速高清。</li>不填表示 OFF。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 视频码率上限，当 Type 指定了极速高清类型时有效。
	// 不填或填0表示不设视频码率上限。
	MaxVideoBitrate *uint64 `json:"MaxVideoBitrate,omitnil,omitempty" name:"MaxVideoBitrate"`
}

type TEHDConfigForUpdate struct {
	// 极速高清类型，可选值：<li>TEHD-100 表示极速高清-100;</li> <li>OFF 表示关闭极速高清。</li>不填表示不修改。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 视频码率上限，不填代表不修改。
	MaxVideoBitrate *uint64 `json:"MaxVideoBitrate,omitnil,omitempty" name:"MaxVideoBitrate"`
}

type TagConfigureInfo struct {
	// 智能标签任务开关，可选值：
	// <li>ON：开启智能标签任务；</li>
	// <li>OFF：关闭智能标签任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type TagConfigureInfoForUpdate struct {
	// 智能标签任务开关，可选值：
	// <li>ON：开启智能标签任务；</li>
	// <li>OFF：关闭智能标签任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type TaskOutputMediaInfo struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 转拉完成后生成的媒体文件基础信息。
	MediaBasicInfo *MediaBasicInfo `json:"MediaBasicInfo,omitnil,omitempty" name:"MediaBasicInfo"`
}

type TaskSimpleInfo struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态。取值：WAITING（等待中）、PROCESSING（处理中）、FINISH（已完成）、ABORTED（已终止）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 视频 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 任务类型，取值：
	// <li>Procedure：视频处理任务；</li>
	// <li>EditMedia：视频编辑任务；</li>
	// <li>ReduceMediaBitrate：降码率任务；</li>
	// <li>WechatDistribute：微信发布任务；</li>
	// <li>ReviewAudioVideo：音视频审核任务。</li>
	// 兼容 2017 版的任务类型：
	// <li>Transcode：视频转码任务；</li>
	// <li>SnapshotByTimeOffset：视频截图任务；</li>
	// <li>Concat：视频拼接任务；</li>
	// <li>Clip：视频剪辑任务；</li>
	// <li>ImageSprites：截取雪碧图任务。</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务开始执行时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。若任务尚未开始，该字段为空。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 任务结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。若任务尚未完成，该字段为空。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type TaskStatData struct {
	// 任务类型。
	// <li> Transcoding: 普通转码</li>
	// <li> Transcoding-TESHD: 极速高清转码</li>
	// <li> Editing: 视频编辑</li>
	// <li> Editing-TESHD: 极速高清视频编辑</li>
	// <li> AdaptiveBitrateStreaming: 自适应码流</li>
	// <li> ContentAudit: 内容审核</li>
	// <li> ContentRecognition: 内容识别</li>
	// <li> RemoveWatermark: 去水印</li>
	// <li> ExtractTraceWatermark: 提取水印</li>
	// <li> AddTraceWatermark: 添加水印</li>
	// <li> RebuildMedia: 音画质重生</li>
	// <li> QualityInspect: 音画质检测</li>
	// <li>Transcode: 转码，包含普通转码、极速高清和视频编辑（不推荐使用）</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务数统计数据概览，用量单位为秒。
	Summary []*TaskStatDataItem `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 不同规格任务统计数据详情。
	// 转码规格：
	// <li>Remuxing: 转封装</li>
	// <li>Audio: 音频转码</li>
	// <li>Standard.H264.SD: H.264编码方式标清转码</li>
	// <li>Standard.H264.HD: H.264编码方式高清转码</li>
	// <li>Standard.H264.FHD: H.264编码方式全高清转码</li>
	// <li>Standard.H264.2K: H.264编码方式2K转码</li>
	// <li>Standard.H264.4K: H.264编码方式4K转码</li>
	// <li>Standard.H265.SD: H.265编码方式标清转码</li>
	// <li>Standard.H265.HD: H.265编码方式高清转码</li>
	// <li>Standard.H265.FHD: H.265编码方式全高清转码</li>
	// <li>Standard.H265.2K: H.265编码方式2K转码</li>
	// <li>Standard.H265.4K: H.265编码方式4K转码</li>
	// <li>TESHD-10.H264.SD: H.264编码方式标清极速高清转码</li>
	// <li>TESHD-10.H264.HD: H.264编码方式高清极速高清转码</li>
	// <li>TESHD-10.H264.FHD: H.264编码方式全高清极速高清转码</li>
	// <li>TESHD-10.H264.2K: H.264编码方式2K极速高清转码</li>
	// <li>TESHD-10.H264.4K: H.264编码方式4K极速高清转码</li>
	// <li>TESHD-10.H265.SD: H.265编码方式标清极速高清转码</li>
	// <li>TESHD-10.H265.HD: H.265编码方式高清极速高清转码</li>
	// <li>TESHD-10.H265.FHD: H.265编码方式全高清极速高清转码</li>
	// <li>TESHD-10.H265.2K: H.265编码方式2K极速高清转码</li>
	// <li>TESHD-10.H265.4K: H.265编码方式4K极速高清转码</li>
	// <li>Edit.Audio: 音频编辑</li>
	// <li>Edit.H264.SD: H.264编码方式标清视频编辑</li>
	// <li>Edit.H264.HD: H.264编码方式高清视频编辑</li>
	// <li>Edit.H264.FHD: H.264编码方式全高清视频编辑</li>
	// <li>Edit.H264.2K: H.264编码方式2K视频编辑</li>
	// <li>Edit.H264.4K: H.264编码方式4K视频编辑</li>
	// <li>Edit.H265.SD: H.265编码方式标清视频编辑</li>
	// <li>Edit.H265.HD: H.265编码方式高清视频编辑</li>
	// <li>Edit.H265.FHD: H.265编码方式全高清视频编辑</li>
	// <li>Edit.H265.2K: H.265编码方式2K视频编辑</li>
	// <li>Edit.H265.4K: H.265编码方式4K视频编辑</li>
	// <li>Edit.TESHD-10.H264.SD: H.264编码方式标清极速高清视频编辑</li>
	// <li>Edit.TESHD-10.H264.HD: H.264编码方式高清极速高清视频编辑</li>
	// <li>Edit.TESHD-10.H264.FHD: H.264编码方式全高清极速高清视频编辑</li>
	// <li>Edit.TESHD-10.H264.2K: H.264编码方式2K极速高清视频编辑</li>
	// <li>Edit.TESHD-10.H264.4K: H.264编码方式4K极速高清视频编辑</li>
	// <li>Edit.TESHD-10.H265.SD: H.265编码方式标清极速高清视频编辑</li>
	// <li>Edit.TESHD-10.H265.HD: H.265编码方式高清极速高清视频编辑</li>
	// <li>Edit.TESHD-10.H265.FHD: H.265编码方式全高清极速高清视频编辑</li>
	// <li>Edit.TESHD-10.H265.2K: H.265编码方式2K极速高清视频编辑</li>
	// <li>Edit.TESHD-10.H265.4K: H.265编码方式4K极速高清视频编辑</li>
	// 去水印、音画质重生规格：
	// <li>480P: 短边 ≤ 480px</li>
	// <li>720P: 短边 ≤ 720px</li>
	// <li>1080P: 短边 ≤ 1080px</li>
	// <li>2K: 短边 ≤ 1440px</li>
	// <li>4K: 短边 ≤ 2160px</li>
	// <li>8K: 短边 ≤ 4320px</li>
	// <li>Audio: 音频</li>
	Details []*SpecificationDataItem `json:"Details,omitnil,omitempty" name:"Details"`
}

type TaskStatDataItem struct {
	// 数据所在时间区间的开始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#52)。如：当时间粒度为天，2018-12-01T00:00:00+08:00，表示2018年12月1日（含）到2018年12月2日（不含）区间。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 任务数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 任务用量。
	Usage *int64 `json:"Usage,omitnil,omitempty" name:"Usage"`
}

type TempCertificate struct {
	// 临时安全证书 Id。
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 临时安全证书 Key。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// Token 值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 证书无效的时间，返回 Unix 时间戳，精确到秒。
	ExpiredTime *uint64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`
}

type TerrorismConfigureInfo struct {
	// 画面鉴别涉及令人不安全的信息的任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImgReviewInfo *TerrorismImgReviewTemplateInfo `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// 文字鉴别涉及令人不安全的信息的任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrReviewInfo *TerrorismOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type TerrorismConfigureInfoForUpdate struct {
	// 画面鉴别涉及令人不安全的信息的任务控制参数。
	ImgReviewInfo *TerrorismImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// 文本鉴别涉及令人不安全的信息的任务控制参数。
	OcrReviewInfo *TerrorismOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type TerrorismImageResult struct {
	// 鉴别涉及令人不安全的信息的评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 鉴别涉及令人不安全的信息的结果建议，取值范围：
	// <li>pass；</li>
	// <li>review；</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 鉴别涉及令人不安全的信息的结果标签，取值范围：
	// <li>guns：武器枪支；</li>
	// <li>crowd：人群聚集；</li>
	// <li>police：警察部队；</li>
	// <li>bloody：血腥画面；</li>
	// <li>banners：暴恐旗帜；</li>
	// <li>explosion：爆炸火灾；</li>
	// <li>scenario：暴恐画面。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`
}

type TerrorismImgReviewTemplateInfo struct {
	// 画面鉴别涉及令人不安全的信息的任务开关，可选值：
	// <li>ON：开启画面鉴别涉及令人不安全的信息的任务；</li>
	// <li>OFF：关闭画面鉴别涉及令人不安全的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 画面鉴别涉及令人不安全的信息的过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>guns：武器枪支；</li>
	// <li>crowd：人群聚集；</li>
	// <li>bloody：血腥画面；</li>
	// <li>police：警察部队；</li>
	// <li>banners：暴恐旗帜；</li>
	// <li>militant：武装分子；</li>
	// <li>explosion：爆炸火灾；</li>
	// <li>terrorists：暴恐人物；</li>
	// <li>scenario：暴恐画面。</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 90 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 80 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type TerrorismImgReviewTemplateInfoForUpdate struct {
	// 画面鉴别涉及令人不安全的信息的任务开关，可选值：
	// <li>ON：开启画面鉴别涉及令人不安全的信息的任务；</li>
	// <li>OFF：关闭画面鉴别涉及令人不安全的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 画面鉴别涉及令人不安全的信息的过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>guns：武器枪支；</li>
	// <li>crowd：人群聚集；</li>
	// <li>bloody：血腥画面；</li>
	// <li>police：警察部队；</li>
	// <li>banners：暴恐旗帜；</li>
	// <li>militant：武装分子；</li>
	// <li>explosion：爆炸火灾；</li>
	// <li>terrorists：暴恐人物；</li>
	// <li>scenario：暴恐画面。</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type TerrorismOcrReviewTemplateInfo struct {
	// 文本鉴别涉及令人不安全的信息的任务开关，可选值：
	// <li>ON：开启文本鉴别涉及令人不安全的信息的任务；</li>
	// <li>OFF：关闭文本鉴别涉及令人不安全的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type TerrorismOcrReviewTemplateInfoForUpdate struct {
	// 文本鉴别涉及令人不安全的信息的任务开关，可选值：
	// <li>ON：开启文本鉴别涉及令人不安全的信息的任务；</li>
	// <li>OFF：关闭文本鉴别涉及令人不安全的信息的任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type TextWatermarkTemplateInput struct {
	// 字体类型，目前可以支持两种：
	// <li>simkai.ttf：可以支持中文和英文；</li>
	// <li>arial.ttf：仅支持英文。</li>
	FontType *string `json:"FontType,omitnil,omitempty" name:"FontType"`

	// 字体大小，格式：Npx，N 为数值。
	FontSize *string `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// 字体颜色，格式：0xRRGGBB，默认值：0xFFFFFF（白色）。
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// 文字透明度，取值范围：(0, 1]
	// <li>0：完全透明</li>
	// <li>1：完全不透明</li>
	// 默认值：1。
	FontAlpha *float64 `json:"FontAlpha,omitnil,omitempty" name:"FontAlpha"`
}

type TextWatermarkTemplateInputForUpdate struct {
	// 字体类型，目前可以支持两种：
	// <li>simkai.ttf：可以支持中文和英文；</li>
	// <li>arial.ttf：仅支持英文。</li>
	FontType *string `json:"FontType,omitnil,omitempty" name:"FontType"`

	// 字体大小，格式：Npx，N 为数值。
	FontSize *string `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// 字体颜色，格式：0xRRGGBB，默认值：0xFFFFFF（白色）。
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// 文字透明度，取值范围：(0, 1]
	// <li>0：完全透明</li>
	// <li>1：完全不透明</li>
	FontAlpha *float64 `json:"FontAlpha,omitnil,omitempty" name:"FontAlpha"`
}

type TimeRange struct {
	// <li>大于等于此时间（起始时间）。</li>
	// <li>格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。</li>
	After *string `json:"After,omitnil,omitempty" name:"After"`

	// <li>小于此时间（结束时间）。</li>
	// <li>格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。</li>
	Before *string `json:"Before,omitnil,omitempty" name:"Before"`
}

type TraceWatermarkInput struct {
	// 溯源水印任务开关，此字段必填，可选值：
	// <li>ON：开启溯源水印；</li>
	// <li>OFF：关闭溯源水印。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 该字段已废弃，请勿使用。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type TranscodePlayInfo2017 struct {
	// 播放地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/266/33476)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`
}

type TranscodeTask2017 struct {
	// 转码任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 被转码文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 被转码文件名称。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 视频时长，单位：秒。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 封面地址。
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// 视频转码后生成的播放信息。
	PlayInfoSet []*TranscodePlayInfo2017 `json:"PlayInfoSet,omitnil,omitempty" name:"PlayInfoSet"`
}

type TranscodeTaskInput struct {
	// 视频转码模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// 溯源水印，该功能已不再支持使用。
	TraceWatermark *TraceWatermarkInput `json:"TraceWatermark,omitnil,omitempty" name:"TraceWatermark"`

	// 版权水印，该功能已不再支持使用。
	CopyRightWatermark *CopyRightWatermarkInput `json:"CopyRightWatermark,omitnil,omitempty" name:"CopyRightWatermark"`

	// 马赛克列表，最大可支持 10 张。
	MosaicSet []*MosaicInput `json:"MosaicSet,omitnil,omitempty" name:"MosaicSet"`

	// 片头片尾列表，支持多片头片尾，最大可支持 10 个。
	HeadTailSet []*HeadTailTaskInput `json:"HeadTailSet,omitnil,omitempty" name:"HeadTailSet"`

	// 转码后的视频的起始时间偏移，单位：秒。
	// <li>不填或填0，表示转码后的视频从原始视频的起始位置开始；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频从原始视频的第 n 秒位置开始；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频从原始视频结束 n 秒前的位置开始。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 转码后视频的终止时间偏移，单位：秒。
	// <li>不填或填0，表示转码后的视频持续到原始视频的末尾终止；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频持续到原始视频第 n 秒时终止；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频持续到原始视频结束 n 秒前终止。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type TranscodeTemplate struct {
	// 转码模板唯一标识。
	Definition *string `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 封装格式，取值：mp4、flv、hls、mp3、flac、ogg。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 转码模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 模板类型，取值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否去除视频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// 视频流配置参数，仅当 RemoveVideo 为 0，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，仅当 RemoveAudio 为 0，该字段有效 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// 极速高清转码参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式板。</li>
	ContainerType *string `json:"ContainerType,omitnil,omitempty" name:"ContainerType"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 切片类型，仅当 Container 为 hls 时有效。
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
}

type TransitionOperation struct {
	// 转场类型，取值有：
	// <ul>
	// <li>图像的转场操作，用于两个视频片段图像间的转场处理：
	// <ul>
	// <li>ImageFadeInFadeOut：图像淡入淡出。 </li>
	// <li>BowTieHorizontal：水平蝴蝶结。 </li>
	// <li>BowTieVertical：垂直蝴蝶结。 </li>
	// <li>ButterflyWaveScrawler：晃动。 </li>
	// <li>Cannabisleaf：枫叶。 </li>
	// <li>Circle：弧形收放。 </li>
	// <li>CircleCrop：圆环聚拢。 </li>
	// <li>Circleopen：椭圆聚拢。 </li>
	// <li>Crosswarp：横向翘曲。 </li>
	// <li>Cube：立方体。 </li>
	// <li>DoomScreenTransition：幕布。 </li>
	// <li>Doorway：门廊。 </li>
	// <li>Dreamy：波浪。 </li>
	// <li>DreamyZoom：水平聚拢。 </li>
	// <li>FilmBurn：火烧云。 </li>
	// <li>GlitchMemories：抖动。 </li>
	// <li>Heart：心形。 </li>
	// <li>InvertedPageCurl：翻页。 </li>
	// <li>Luma：腐蚀。 </li>
	// <li>Mosaic：九宫格。 </li>
	// <li>Pinwheel：风车。 </li>
	// <li>PolarFunction：椭圆扩散。 </li>
	// <li>PolkaDotsCurtain：弧形扩散。 </li>
	// <li>Radial：雷达扫描 </li>
	// <li>RotateScaleFade：上下收放。 </li>
	// <li>Squeeze：上下聚拢。 </li>
	// <li>Swap：放大切换。 </li>
	// <li>Swirl：螺旋。 </li>
	// <li>UndulatingBurnOutSwirl：水流蔓延。 </li>
	// <li>Windowblinds：百叶窗。 </li>
	// <li>WipeDown：向下收起。 </li>
	// <li>WipeLeft：向左收起。 </li>
	// <li>WipeRight：向右收起。 </li>
	// <li>WipeUp：向上收起。 </li>
	// <li>ZoomInCircles：水波纹。 </li>
	// </ul>
	// </li>
	// <li>音频的转场操作，用于两个音频片段间的转场处理：
	// <ul>
	// <li>AudioFadeInFadeOut：声音淡入淡出。 </li>
	// </ul>
	// </li>
	// </ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type TransitionOpertion struct {
	// 转场类型，取值有：
	// <ul>
	// <li>图像的转场操作，用于两个视频片段图像间的转场处理：
	// <ul>
	// <li>ImageFadeInFadeOut：图像淡入淡出。 </li>
	// <li>BowTieHorizontal：水平蝴蝶结。 </li>
	// <li>BowTieVertical：垂直蝴蝶结。 </li>
	// <li>ButterflyWaveScrawler：晃动。 </li>
	// <li>Cannabisleaf：枫叶。 </li>
	// <li>Circle：弧形收放。 </li>
	// <li>CircleCrop：圆环聚拢。 </li>
	// <li>Circleopen：椭圆聚拢。 </li>
	// <li>Crosswarp：横向翘曲。 </li>
	// <li>Cube：立方体。 </li>
	// <li>DoomScreenTransition：幕布。 </li>
	// <li>Doorway：门廊。 </li>
	// <li>Dreamy：波浪。 </li>
	// <li>DreamyZoom：水平聚拢。 </li>
	// <li>FilmBurn：火烧云。 </li>
	// <li>GlitchMemories：抖动。 </li>
	// <li>Heart：心形。 </li>
	// <li>InvertedPageCurl：翻页。 </li>
	// <li>Luma：腐蚀。 </li>
	// <li>Mosaic：九宫格。 </li>
	// <li>Pinwheel：风车。 </li>
	// <li>PolarFunction：椭圆扩散。 </li>
	// <li>PolkaDotsCurtain：弧形扩散。 </li>
	// <li>Radial：雷达扫描 </li>
	// <li>RotateScaleFade：上下收放。 </li>
	// <li>Squeeze：上下聚拢。 </li>
	// <li>Swap：放大切换。 </li>
	// <li>Swirl：螺旋。 </li>
	// <li>UndulatingBurnOutSwirl：水流蔓延。 </li>
	// <li>Windowblinds：百叶窗。 </li>
	// <li>WipeDown：向下收起。 </li>
	// <li>WipeLeft：向左收起。 </li>
	// <li>WipeRight：向右收起。 </li>
	// <li>WipeUp：向上收起。 </li>
	// <li>ZoomInCircles：水波纹。 </li>
	// </ul>
	// </li>
	// <li>音频的转场操作，用于两个音频片段间的转场处理：
	// <ul>
	// <li>AudioFadeInFadeOut：声音淡入淡出。 </li>
	// </ul>
	// </li>
	// </ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type TrtcRecordInfo struct {
	// TRTC 应用 ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC 房间 ID。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 录制任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 参与录制的用户 ID 列表。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

type UrlSignatureAuthPolicy struct {
	// [Key 防盗链](https://cloud.tencent.com/document/product/266/14047)设置状态，可选值：
	// <li>Enabled: 启用。</li>
	// <li>Disabled: 禁用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// [Key 防盗链](https://cloud.tencent.com/document/product/266/14047)中用于生成签名的密钥。
	// EncryptedKey 字符串的长度为8~40个字节，不能包含不可见字符。
	EncryptedKey *string `json:"EncryptedKey,omitnil,omitempty" name:"EncryptedKey"`
}

type UserDefineAsrTextReviewTemplateInfo struct {
	// 用户自定语音审核任务开关，可选值：
	// <li>ON：开启自定义语音审核任务；</li>
	// <li>OFF：关闭自定义语音审核任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 用户自定义语音过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义语音关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineAsrTextReviewTemplateInfoForUpdate struct {
	// 用户自定语音审核任务开关，可选值：
	// <li>ON：开启自定义语音审核任务；</li>
	// <li>OFF：关闭自定义语音审核任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 用户自定义语音过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义语音关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineConfigureInfo struct {
	// 用户自定义人物音视频审核控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceReviewInfo *UserDefineFaceReviewTemplateInfo `json:"FaceReviewInfo,omitnil,omitempty" name:"FaceReviewInfo"`

	// 用户自定义语音音视频审核控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfo `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// 用户自定义文本音视频审核控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type UserDefineConfigureInfoForUpdate struct {
	// 用户自定义人物音视频审核控制参数。
	FaceReviewInfo *UserDefineFaceReviewTemplateInfoForUpdate `json:"FaceReviewInfo,omitnil,omitempty" name:"FaceReviewInfo"`

	// 用户自定义语音音视频审核控制参数。
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// 用户自定义文本音视频审核控制参数。
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type UserDefineFaceReviewTemplateInfo struct {
	// 用户自定义人物音视频审核任务开关，可选值：
	// <li>ON：开启自定义人物音视频审核任务；</li>
	// <li>OFF：关闭自定义人物音视频审核任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 用户自定义人物过滤标签，音视频审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义人物库的时，需要添加对应人物标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 97 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 95 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineFaceReviewTemplateInfoForUpdate struct {
	// 用户自定义人物音视频审核任务开关，可选值：
	// <li>ON：开启自定义人物音视频审核任务；</li>
	// <li>OFF：关闭自定义人物音视频审核任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 用户自定义人物过滤标签，音视频审核结果包含选择的标签则返回结果，如果过滤标签为空，则音视频审核结果全部返回。如果要使用标签过滤功能，添加自定义人物库的时，需要添加对应人物标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当音视频审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当音视频审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfo struct {
	// 用户自定文本音视频审核任务开关，可选值：
	// <li>ON：开启自定义文本音视频审核任务；</li>
	// <li>OFF：关闭自定义文本音视频审核任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 用户自定义文本过滤标签，音视频审核结果包含选择的标签则返回结果，如果过滤标签为空，则音视频审核结果全部返回。如果要使用标签过滤功能，添加自定义文本关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfoForUpdate struct {
	// 用户自定文本音视频审核任务开关，可选值：
	// <li>ON：开启自定义文本音视频审核任务；</li>
	// <li>OFF：关闭自定义文本音视频审核任务。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 用户自定义文本过滤标签，音视频审核结果包含选择的标签则返回结果，如果过滤标签为空，则音视频审核结果全部返回。如果要使用标签过滤功能，添加自定义文本关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 判定涉嫌违规的分数阈值，当审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

// Predefined struct for user
type VerifyDomainRecordRequestParams struct {
	// 需要接入点播的加速域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 验证方式：
	// <li>dns：DNS 解析验证；</li>
	// <li>fIle：文件验证。</li>
	// 
	// 默认值：dns。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

type VerifyDomainRecordRequest struct {
	*tchttp.BaseRequest
	
	// 需要接入点播的加速域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 验证方式：
	// <li>dns：DNS 解析验证；</li>
	// <li>fIle：文件验证。</li>
	// 
	// 默认值：dns。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

func (r *VerifyDomainRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDomainRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "VerifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyDomainRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyDomainRecordResponseParams struct {
	// 是否验证成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyDomainRecordResponse struct {
	*tchttp.BaseResponse
	Response *VerifyDomainRecordResponseParams `json:"Response"`
}

func (r *VerifyDomainRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDomainRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VideoConfigureInfo struct {
	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 1920]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 
	// 默认值：0。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 1920]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 
	// 默认值：0。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 视频流的码率，取值范围：0 和 [128, 10000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`
}

type VideoConfigureInfoForUpdate struct {
	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 1920]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 
	// 默认值：0。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 1920]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 
	// 默认值：0。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 视频流的码率，取值范围：0 和 [128, 10000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`
}

type VideoDenoiseInfo struct {
	// 视频降噪控制开关，可选值：
	// <li>ON：开启视频降噪；</li>
	// <li>OFF：关闭视频降噪。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 视频降噪类型，仅当视频降噪控制开关为 ON 时有效，可选值：
	// <li>weak：轻视频降噪；</li>
	// <li>strong：强视频降噪。</li>
	// 默认值：weak。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type VideoFrameInterpolationInfo struct {
	// 智能插帧控制开关，可选值：
	// <li>ON：开启智能插帧；</li>
	// <li>OFF：关闭智能插帧。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 智能插帧帧率，帧率范围为 (0, 100]，仅当智能插帧控制开关为 ON 时有效。默认跟源文件帧率一致。
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`
}

type VideoTemplateInfo struct {
	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码；</li>
	// <li>libx265：H.265 编码；</li>
	// <li>av1：AOMedia Video 1 编码；</li>
	// <li>H.266：H.266 编码。</li>
	// <font color=red>注意：</font>
	// <li> av1，H.266 编码容器目前只支持 mp4 ；</li>
	// <li> H.266 目前只支持恒定 CRF 码率控制方式。 </li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 视频帧率，取值范围：[0, 100]，单位：Hz。
	// 当取值为 0，表示帧率和原始视频保持一致。
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 8192]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 8192]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 视频恒定码率控制因子，取值范围为[1, 51]。
	// 
	// <font color=red>注意：</font>
	// <li>如果指定该参数，将使用 CRF 的码率控制方式做转码（视频码率将不再生效）；</li>
	// <li>当指定视频流编码格式为 H.266 时，该字段必填，推荐值为 28；</li>
	// <li>如果没有特殊需求，不建议指定该参数。</li>
	Vcrf *uint64 `json:"Vcrf,omitnil,omitempty" name:"Vcrf"`

	// 关键帧 I 帧之间的间隔，取值范围：0 和 [1, 100000]，单位：帧数。
	// 当填 0 或不填时，系统将自动设置 gop 长度。
	Gop *uint64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// 当原始视频为 HDR（High Dynamic Range）时，转码输出是否依然保持 HDR。取值范围：
	// <li>ON: 如果原始文件是 HDR，则转码输出保持 HDR；否则转码输出为 SDR （Standard Dynamic Range）。</li>
	// <li>OFF: 无论原始文件是 HDR 还是 SDR，转码输出均为 SDR。</li>
	// 默认值：OFF。
	PreserveHDRSwitch *string `json:"PreserveHDRSwitch,omitnil,omitempty" name:"PreserveHDRSwitch"`

	// 编码标签，仅当视频流的编码格式为 H.265 编码时有效，可选值：
	// <li>hvc1 表示 hvc1 标签；</li>
	// <li>hev1 表示 hev1 标签。 </li>
	// 默认值：hvc1。
	CodecTag *string `json:"CodecTag,omitnil,omitempty" name:"CodecTag"`
}

type VideoTemplateInfoForUpdate struct {
	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码；</li>
	// <li>libx265：H.265 编码；</li>
	// <li>av1：AOMedia Video 1 编码；</li>
	// <li>H.266：H.266 编码。</li>
	// <font color=red>注意：</font>
	// <li> av1，H.266 编码容器目前只支持 mp4 ；</li>
	// <li> H.266 目前只支持恒定 CRF 码率控制方式。 </li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 视频帧率，取值范围：[0, 100]，单位：Hz。
	// 当取值为 0，表示帧率和原始视频保持一致。
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 8192]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 8192]，单位：px。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// <li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li>
	// <li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊填充。</li>
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 视频恒定码率控制因子，取值范围为[1, 51]，填 0 表示禁用该参数。
	// 
	// <font color=red>注意：</font>
	// <li>如果指定该参数，将使用 CRF 的码率控制方式做转码（视频码率将不再生效）；</li>
	// <li>当指定视频流编码格式为 H.266 时，该字段必填，推荐值为 28；</li>
	// <li>如果没有特殊需求，不建议指定该参数。</li>
	Vcrf *uint64 `json:"Vcrf,omitnil,omitempty" name:"Vcrf"`

	// 关键帧 I 帧之间的间隔，取值范围：0 和 [1, 100000]，单位：帧数。
	// 当填 0 或不填时，系统将自动设置 gop 长度。
	Gop *uint64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// 当原始视频为 HDR（High Dynamic Range）时，转码输出是否依然保持 HDR。取值范围：
	// <li>ON: 如果原始文件是 HDR，则转码输出保持 HDR；否则转码输出为 SDR （Standard Dynamic Range）。</li>
	// <li>OFF: 无论原始文件是 HDR 还是 SDR，转码输出均为 SDR。</li>
	PreserveHDRSwitch *string `json:"PreserveHDRSwitch,omitnil,omitempty" name:"PreserveHDRSwitch"`

	// 编码标签，仅当视频流的编码格式为 H.265 编码时有效，可选值：
	// <li>hvc1 表示 hvc1 标签；</li>
	// <li>hev1 表示 hev1 标签。 </li>
	// 默认值：hvc1。
	CodecTag *string `json:"CodecTag,omitnil,omitempty" name:"CodecTag"`
}

type VideoTrackItem struct {
	// 视频片段的媒体素材来源，可以是：
	// <li>点播的媒体文件 ID；</li>
	// <li>其他媒体文件的下载 URL。</li>
	// 注意：当使用其他媒体文件的下载 URL 作为素材来源，且开启了访问控制（如防盗链）时，需要在 URL 携带访问控制参数（如防盗链签名）。
	SourceMedia *string `json:"SourceMedia,omitnil,omitempty" name:"SourceMedia"`

	// 视频片段取自素材文件的起始时间，单位为秒。默认为0。
	SourceMediaStartTime *float64 `json:"SourceMediaStartTime,omitnil,omitempty" name:"SourceMediaStartTime"`

	// 视频片段时长，单位为秒。默认取视频素材本身长度，表示截取全部素材。如果源文件是图片，Duration需要大于0。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 视频片段目标时长，单位为秒。
	// <li>当 TargetDuration 不填或填0时，表示目标时长和 Duration 一致；</li>
	// <li>当 TargetDuration 取大于0的值时，将对视频片段做快进或慢放等处理，使得输出片段的时长等于 TargetDuration。</li>
	TargetDuration *float64 `json:"TargetDuration,omitnil,omitempty" name:"TargetDuration"`

	// 视频原点位置，取值有：
	// <li>Center：坐标原点为中心位置，如画布中心。</li>
	// 默认值 ：Center。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// 视频片段原点距离画布原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 XPos 为画布宽度指定百分比的位置，如 10% 表示 XPos 为画布口宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示视频片段 XPos 单位为像素，如 100px 表示 XPos 为100像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 视频片段原点距离画布原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 YPos 为画布高度指定百分比的位置，如 10% 表示 YPos 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示视频片段 YPos 单位为像素，如 100px 表示 YPos 为100像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// 视频片段的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 Width 为画布宽度的百分比大小，如 10% 表示 Width 为画布宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示视频片段 Width 单位为像素，如 100px 表示 Width 为100像素。</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取视频素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按比例缩放</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频片段的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 Height 为画布高度的百分比大小，如 10% 表示 Height 为画布高度的 10%；
	// </li><li>当字符串以 px 结尾，表示视频片段 Height 单位为像素，如 100px 表示 Height 为100像素。</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取视频素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按比例缩放</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 对音频进行操作，如静音等。
	AudioOperations []*AudioTransform `json:"AudioOperations,omitnil,omitempty" name:"AudioOperations"`

	// 对图像进行的操作，如图像旋转等。
	ImageOperations []*ImageTransform `json:"ImageOperations,omitnil,omitempty" name:"ImageOperations"`
}

type VoiceConfigureInfo struct {
	// 音频（静音、低音、爆音）检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type VoiceConfigureInfoForUpdate struct {
	// 音频（静音、低音、爆音）检测开关，可选值：
	// <li>ON：开启；</li>
	// <li>OFF：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type WatermarkConfigureData struct {
	// 是否启用水印。可取值：
	// <li>ON：表示启用水印；</li>
	// <li>OFF：表示关闭水印。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 水印 Url。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 水印的宽度。
	// <li>字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印的高度。
	// <li>字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印原点距离视频图像坐标原点的水平位置。字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%。
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`
}

type WatermarkConfigureInfo struct {
	// 是否启用水印。可取值：
	// <li>ON：表示启用水印；</li>
	// <li>OFF：表示关闭水印。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 水印图片 Base64 编码后的字符串。支持 jpeg、png 图片格式。
	ImageContent *string `json:"ImageContent,omitnil,omitempty" name:"ImageContent"`

	// 水印的宽度。
	// <li>字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印的高度。
	// <li>字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印原点距离视频图像坐标原点的水平位置。字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%。
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`
}

type WatermarkConfigureInfoForUpdate struct {
	// 是否启用水印。可取值：
	// <li>ON：表示启用水印；</li>
	// <li>OFF：表示关闭水印。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 水印图片 Base64 编码后的字符串。支持 jpeg、png 图片格式。
	ImageContent *string `json:"ImageContent,omitnil,omitempty" name:"ImageContent"`

	// 水印的宽度。
	// <li>字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印的高度。
	// <li>字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印原点距离视频图像坐标原点的水平位置。字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%。
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`
}

type WatermarkCycleConfigForUpdate struct {
	// 水印在视频里第一次出现的播放时间点，单位：秒。
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 在一个水印周期内，水印显示的持续时间，单位：秒。
	DisplayDuration *float64 `json:"DisplayDuration,omitnil,omitempty" name:"DisplayDuration"`

	// 一个水印周期的持续时间，单位：秒。
	// 填 0 表示水印只持续一个水印周期（即在整个视频里只显示 DisplayDuration 秒）。
	CycleDuration *float64 `json:"CycleDuration,omitnil,omitempty" name:"CycleDuration"`
}

type WatermarkInput struct {
	// 水印模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 文字内容，长度不超过100个字符。仅当水印类型为文字水印时填写。
	// 文字水印不支持截图打水印。
	TextContent *string `json:"TextContent,omitnil,omitempty" name:"TextContent"`

	// SVG 内容。长度不超过 2000000 个字符。仅当水印类型为 SVG 水印时填写。
	// SVG 水印不支持截图打水印。
	SvgContent *string `json:"SvgContent,omitnil,omitempty" name:"SvgContent"`

	// 水印的起始时间偏移，单位：秒。不填或填0，表示水印从画面出现时开始显现。
	// <li>不填或填0，表示水印从画面开始就出现；</li>
	// <li>当数值大于0时（假设为 n），表示水印从画面开始的第 n 秒出现；</li>
	// <li>当数值小于0时（假设为 -n），表示水印从离画面结束 n 秒前开始出现。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 水印的结束时间偏移，单位：秒。
	// <li>不填或填0，表示水印持续到画面结束；</li>
	// <li>当数值大于0时（假设为 n），表示水印持续到第 n 秒时消失；</li>
	// <li>当数值小于0时（假设为 -n），表示水印持续到离画面结束 n 秒前消失。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type WatermarkTemplate struct {
	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 水印类型，取值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 水印模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 水印图片原点距离视频图像原点的水平位置。
	// <li>当字符串以 % 结尾，表示水印 Left 为视频宽度指定百分比的位置，如 10% 表示 Left 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Left 为视频宽度指定像素的位置，如 100px 表示 Left 为 100 像素。</li>
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 水印图片原点距离视频图像原点的垂直位置。
	// <li>当字符串以 % 结尾，表示水印 Top 为视频高度指定百分比的位置，如 10% 表示 Top 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Top 为视频高度指定像素的位置，如 100px 表示 Top 为 100 像素。</li>
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// 图片水印模板，仅当 Type 为 image，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTemplate *ImageWatermarkTemplate `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// 文字水印模板，仅当 Type 为 text，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG 水印模板，当 Type 为 svg，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 原点位置，可选值：
	// <li>topLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>topRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>bottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>bottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下。；</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`
}

// Predefined struct for user
type WeChatMiniProgramPublishRequestParams struct {
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 发布视频所对应的转码模板 ID，为0代表原始视频。
	SourceDefinition *int64 `json:"SourceDefinition,omitnil,omitempty" name:"SourceDefinition"`
}

type WeChatMiniProgramPublishRequest struct {
	*tchttp.BaseRequest
	
	// 媒体文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <b>点播[应用](/document/product/266/14574) ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 发布视频所对应的转码模板 ID，为0代表原始视频。
	SourceDefinition *int64 `json:"SourceDefinition,omitnil,omitempty" name:"SourceDefinition"`
}

func (r *WeChatMiniProgramPublishRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WeChatMiniProgramPublishRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "SourceDefinition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WeChatMiniProgramPublishRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WeChatMiniProgramPublishResponseParams struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type WeChatMiniProgramPublishResponse struct {
	*tchttp.BaseResponse
	Response *WeChatMiniProgramPublishResponseParams `json:"Response"`
}

func (r *WeChatMiniProgramPublishResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WeChatMiniProgramPublishResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebPageRecordInfo struct {
	// 全景录制地址。
	RecordUrl *string `json:"RecordUrl,omitnil,omitempty" name:"RecordUrl"`

	// 全景录制任务 ID。
	RecordTaskId *string `json:"RecordTaskId,omitnil,omitempty" name:"RecordTaskId"`
}

type WechatMiniProgramPublishTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// WAITING：等待中；
	// PROCESSING：处理中；
	// FINISH：已完成。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 发布视频文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 发布视频所对应的转码模板 ID，为 0 代表原始视频。
	SourceDefinition *uint64 `json:"SourceDefinition,omitnil,omitempty" name:"SourceDefinition"`

	// 微信小程序视频发布状态，取值：
	// <li>Pass：发布成功；</li>
	// <li>Failed：发布失败；</li>
	// <li>Rejected：音视频审核未通过。</li>
	PublishResult *string `json:"PublishResult,omitnil,omitempty" name:"PublishResult"`
}

type WechatMiniProgramPublishTaskInput struct {
	// 发布视频所对应的转码模板 ID，为 0 代表原始视频。
	SourceDefinition *uint64 `json:"SourceDefinition,omitnil,omitempty" name:"SourceDefinition"`
}

type WechatPublishTask struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// WAITING：等待中；
	// PROCESSING：处理中；
	// FINISH：已完成。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 发布视频文件 ID。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 微信发布模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 发布视频所对应的转码模板 ID，为 0 代表原始视频。
	SourceDefinition *uint64 `json:"SourceDefinition,omitnil,omitempty" name:"SourceDefinition"`

	// 微信发布状态，取值：
	// <li>FAIL：失败；</li>
	// <li>SUCCESS：成功；</li>
	// <li>AUDITNOTPASS：审核未通过；</li>
	// <li>NOTTRIGGERED：尚未发起微信发布。</li>
	WechatStatus *string `json:"WechatStatus,omitnil,omitempty" name:"WechatStatus"`

	// 微信 Vid。
	WechatVid *string `json:"WechatVid,omitnil,omitempty" name:"WechatVid"`

	// 微信地址。
	WechatUrl *string `json:"WechatUrl,omitnil,omitempty" name:"WechatUrl"`
}