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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AdaptiveDynamicStreamingTaskInput struct {

	// 转自适应码流模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`

	// 转自适应码流后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 转自适应码流后，manifest 文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 转自适应码流（HLS）后，二级 index 文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}_{trackType}_{trackDefinition}.{format}`。
	SubStreamManifestObjectName *string `json:"SubStreamManifestObjectName,omitempty" name:"SubStreamManifestObjectName"`

	// 转自适应码流后，分片文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}_{trackType}_{trackDefinition}_{number}.{format}`。
	SegmentObjectName *string `json:"SegmentObjectName,omitempty" name:"SegmentObjectName"`

	// 转自适应码流后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type AiAnalysisTaskInput struct {

	// 视频内容分析模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiContentReviewTaskInput struct {

	// 视频内容审核模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskInput struct {

	// 视频智能识别模板 ID 。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AnimatedGraphicTaskInput struct {

	// 视频转动图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 动图在视频中的开始时间，单位为秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间，单位为秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 转动图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 转动图后文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_animatedGraphic_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`
}

type AnimatedGraphicsTemplate struct {

	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 转动图模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 转动图模板描述。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 动图宽度（或长边）的最大值。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 动图高度（或短边）的最大值。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 动图格式。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 帧率。
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 图片质量。
	Quality *float64 `json:"Quality,omitempty" name:"Quality"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
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
	// <li>libfdk_aac；</li>
	// <li>libmp3lame。</li>
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。
	// 当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频流的采样率，可选值：
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 音频通道方式，可选值：
	// <li>1：单通道</li>
	// <li>2：双通道</li>
	// <li>6：立体声</li>
	// 默认值：2。
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
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
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。 当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频流的采样率，可选值：
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 音频通道方式，可选值：
	// <li>1：单通道</li>
	// <li>2：双通道</li>
	// <li>6：立体声</li>
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type CosFileUploadTrigger struct {

	// 工作流绑定的 COS Bucket 名，如 TopRankVideo-125xxx88。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 工作流绑定的 COS Bucket 所属园区，如 ap-chongiqng。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 工作流绑定的输入路径目录，如`/movie/201907/`，不填代表根目录`/`。
	Dir *string `json:"Dir,omitempty" name:"Dir"`

	// 工作流允许触发的文件格式列表，如 ["mp4", "flv", "mov"]。不填代表所有格式的文件都可以触发工作流。
	Formats []*string `json:"Formats,omitempty" name:"Formats" list`
}

type CosInputInfo struct {

	// 视频处理对象文件所在的 COS Bucket 名，如 TopRankVideo-125xxx88。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 视频处理对象文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 视频处理对象文件的输入路径，如`/movie/201907/WildAnimal.mov`。
	Object *string `json:"Object,omitempty" name:"Object"`
}

type CosOutputStorage struct {

	// 视频处理生成的文件输出的目标 Bucket 名，如 TopRankVideo-125xxx88。如果不填，表示继承上层。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 视频处理生成的文件输出的目标 Bucket 的园区，如 ap-chongqing。如果不填，表示继承上层。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type CreateAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 帧率，取值范围：[1, 30]，单位：Hz。
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 动图格式，取值为 gif 和 webp。默认为 gif。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitempty" name:"Quality"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转动图模板唯一标识。
		Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAnimatedGraphicsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest

	// 雪碧图中小图的宽度，取值范围： [128, 4096]，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 雪碧图中小图的高度，取值范围： [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitempty" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitempty" name:"ColumnCount"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageSpriteTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 雪碧图模板唯一标识。
		Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageSpriteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageSpriteTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 图片宽度，取值范围： [128, 4096]，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 图片高度，取值范围： [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 图片格式，取值为 jpg 和 png。默认为 jpg。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateSampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSampleSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 采样截图模板唯一标识。
		Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSampleSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSampleSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest

	// 图片宽度，取值范围： [128, 4096]，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 图片高度，取值范围： [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 图片格式，取值可以为 jpg 和 png。默认为 jpg。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateSnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 时间点截图模板唯一标识。
		Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSnapshotByTimeOffsetTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 视频流配置参数，当 RemoveVideo 为 0，该字段必填。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，当 RemoveAudio 为 0，该字段必填。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitempty" name:"AudioTemplate"`
}

func (r *CreateTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转码模板唯一标识。
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// 水印类型，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	// 默认值：TopLeft。目前，当 Type 为 image，该字段仅支持 TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 图片水印模板，仅当 Type 为 image，该字段必填且有效。
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// 文字水印模板，仅当 Type 为 text，该字段必填且有效。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG 水印模板，仅当 Type 为 svg，该字段必填且有效。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitempty" name:"SvgTemplate"`
}

func (r *CreateWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWatermarkTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 水印模板唯一标识。
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// 水印图片地址，仅当 Type 为 image，该字段有效。
		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWatermarkTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest

	// 工作流名称，最多128字符。同一个用户该名称唯一。
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 工作流绑定的触发规则，当上传视频命中该规则到该对象时即触发工作流。
	Trigger *WorkflowTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// 视频处理的文件输出存储位置。不填则继承 Trigger 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 视频处理生成的文件输出的目标目录，如`/movie/201907/`。如果不填，表示与触发文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// 任务的事件通知配置，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TaskPriority *int64 `json:"TaskPriority,omitempty" name:"TaskPriority"`
}

func (r *CreateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWorkflowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 工作流 ID。
		WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWorkflowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest

	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAnimatedGraphicsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest

	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageSpriteTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageSpriteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageSpriteTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteSampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSampleSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSampleSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSampleSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest

	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteSnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotByTimeOffsetTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWatermarkTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWatermarkTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowRequest struct {
	*tchttp.BaseRequest

	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

func (r *DeleteWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWorkflowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWorkflowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAnimatedGraphicsTemplatesRequest struct {
	*tchttp.BaseRequest

	// 转动图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAnimatedGraphicsTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAnimatedGraphicsTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAnimatedGraphicsTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 转动图模板详情列表。
		AnimatedGraphicsTemplateSet []*AnimatedGraphicsTemplate `json:"AnimatedGraphicsTemplateSet,omitempty" name:"AnimatedGraphicsTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAnimatedGraphicsTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAnimatedGraphicsTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSpriteTemplatesRequest struct {
	*tchttp.BaseRequest

	// 雪碧图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeImageSpriteTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSpriteTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSpriteTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 雪碧图模板详情列表。
		ImageSpriteTemplateSet []*ImageSpriteTemplate `json:"ImageSpriteTemplateSet,omitempty" name:"ImageSpriteTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSpriteTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSpriteTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSampleSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest

	// 采样截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSampleSnapshotTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSampleSnapshotTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSampleSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 采样截图模板详情列表。
		SampleSnapshotTemplateSet []*SampleSnapshotTemplate `json:"SampleSnapshotTemplateSet,omitempty" name:"SampleSnapshotTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSampleSnapshotTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSampleSnapshotTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotByTimeOffsetTemplatesRequest struct {
	*tchttp.BaseRequest

	// 指定时间点截图模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSnapshotByTimeOffsetTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSnapshotByTimeOffsetTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotByTimeOffsetTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 指定时间点截图模板详情列表。
		SnapshotByTimeOffsetTemplateSet []*SnapshotByTimeOffsetTemplate `json:"SnapshotByTimeOffsetTemplateSet,omitempty" name:"SnapshotByTimeOffsetTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotByTimeOffsetTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSnapshotByTimeOffsetTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 视频处理任务的任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务类型，目前仅有取值 WorkflowTask。
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// 任务状态，取值：
	// <li>WAITING：等待中；</li>
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务的创建时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
		BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

		// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

		// 视频处理任务信息，仅当 TaskType 为 WorkflowTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		WorkflowTask *WorkflowTask `json:"WorkflowTask,omitempty" name:"WorkflowTask"`

		// 任务的事件通知信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

		// 任务流的优先级，取值范围为 [-10, 10]。
		TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

		// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长50个字符，不带或者带空字符串表示不做去重。
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长1000个字符。
		SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：任务状态，可选值：WAITING（等待中）、PROCESSING（处理中）、FINISH（已完成）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务概要列表。
		TaskSet []*TaskSimpleInfo `json:"TaskSet,omitempty" name:"TaskSet" list`

		// 翻页标识，当请求未返回所有数据，该字段表示下一条记录的 ID。当该字段为空字符串，说明已无更多数据。
		ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest

	// 转码模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式板；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式。</li>
	ContainerType *string `json:"ContainerType,omitempty" name:"ContainerType"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 转码模板详情列表。
		TranscodeTemplateSet []*TranscodeTemplate `json:"TranscodeTemplateSet,omitempty" name:"TranscodeTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户状态，取值：
	// <li>Normal：正常 ；</li>
	// <li> Closed：下线；</li>
	// <li> Arrearage：欠费停服。</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// 用户付费类型，取值：
	// <li>DailyPayment：日结付费 ；</li>
	// <li>MonthlyPayment：月结付费。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
		PaymentType *string `json:"PaymentType,omitempty" name:"PaymentType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWatermarkTemplatesRequest struct {
	*tchttp.BaseRequest

	// 水印模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 水印类型过滤条件，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数
	// <li>默认值：10；</li>
	// <li>最大值：100。</li>
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeWatermarkTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWatermarkTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWatermarkTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 水印模板详情列表。
		WatermarkTemplateSet []*WatermarkTemplate `json:"WatermarkTemplateSet,omitempty" name:"WatermarkTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWatermarkTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWatermarkTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkflowsRequest struct {
	*tchttp.BaseRequest

	// 工作流 ID 过滤条件，数组长度限制：100。
	WorkflowIds []*int64 `json:"WorkflowIds,omitempty" name:"WorkflowIds" list`

	// 工作流状态，取值范围：
	// <li>Enabled：已启用，</li>
	// <li>Disabled：已禁用。</li>
	// 不填此参数，则不区分工作流状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWorkflowsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 工作流信息数组。
		WorkflowInfoSet []*WorkflowInfo `json:"WorkflowInfoSet,omitempty" name:"WorkflowInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWorkflowsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableWorkflowRequest struct {
	*tchttp.BaseRequest

	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

func (r *DisableWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableWorkflowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableWorkflowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableWorkflowRequest struct {
	*tchttp.BaseRequest

	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

func (r *EnableWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableWorkflowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableWorkflowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageSpriteTaskInput struct {

	// 雪碧图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 截取雪碧图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 截取雪碧图后，雪碧图图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_imageSprite_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 截取雪碧图后，Web VTT 文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_imageSprite_{definition}.{format}`。
	WebVttObjectName *string `json:"WebVttObjectName,omitempty" name:"WebVttObjectName"`

	// 截取雪碧图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type ImageSpriteTemplate struct {

	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 雪碧图模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 雪碧图中小图的宽度。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 雪碧图中小图的高度。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 采样类型。
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 采样间隔。
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitempty" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitempty" name:"ColumnCount"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ImageWatermarkInput struct {

	// 水印图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串。支持 jpeg、png 图片格式。
	ImageContent *string `json:"ImageContent,omitempty" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 默认值：10%。
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 默认值：0px，表示 Height 按照原始水印图片的宽高比缩放。
	Height *string `json:"Height,omitempty" name:"Height"`
}

type ImageWatermarkInputForUpdate struct {

	// 水印图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串。支持 jpeg、png 图片格式。
	ImageContent *string `json:"ImageContent,omitempty" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 0px 表示 Height 按照 Width 对视频宽度的比例缩放。
	Height *string `json:"Height,omitempty" name:"Height"`
}

type ImageWatermarkTemplate struct {

	// 水印图片地址。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素；</li>
	// 0px：表示 Height 按照 Width 对视频宽度的比例缩放。
	Height *string `json:"Height,omitempty" name:"Height"`
}

type MediaAnimatedGraphicsItem struct {

	// 转动图文件的存储位置。
	Storage *TaskOutputStorage `json:"Storage,omitempty" name:"Storage"`

	// 转动图的文件路径。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 转动图模板 ID，参见[转动图参数模板](https://cloud.tencent.com/document/product/266/33481#.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 动图格式，如 gif。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 动图的高度，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 动图的宽度，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 动图码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 动图大小，单位：字节。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 动图的md5值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 动图在视频中的起始时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type MediaAudioStreamItem struct {

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

type MediaImageSpriteItem struct {

	// 雪碧图规格，参见[雪碧图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 雪碧图小图的高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 雪碧图小图的宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 每一张雪碧图大图里小图的数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 每一张雪碧图大图的路径。
	ImagePathSet []*string `json:"ImagePathSet,omitempty" name:"ImagePathSet" list`

	// 雪碧图子图位置与时间关系的 WebVtt 文件路径。WebVtt 文件表明了各个雪碧图小图对应的时间点，以及在在雪碧大图里的坐标位置，一般被播放器用于实现预览。
	WebVttPath *string `json:"WebVttPath,omitempty" name:"WebVttPath"`

	// 雪碧图文件的存储位置。
	Storage *TaskOutputStorage `json:"Storage,omitempty" name:"Storage"`
}

type MediaInputInfo struct {

	// 视频处理对象的类型，现在仅支持 COS。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当 Type 为 COS 时有效，则该项为必填，表示视频处理 COS 对象信息。
	CosInputInfo *CosInputInfo `json:"CosInputInfo,omitempty" name:"CosInputInfo"`
}

type MediaMetaData struct {

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
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet" list`

	// 音频流信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet" list`

	// 视频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoDuration *float64 `json:"VideoDuration,omitempty" name:"VideoDuration"`

	// 音频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioDuration *float64 `json:"AudioDuration,omitempty" name:"AudioDuration"`
}

type MediaProcessTaskAnimatedGraphicResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 转动图任务的输入。
	Input *AnimatedGraphicTaskInput `json:"Input,omitempty" name:"Input"`

	// 转动图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaAnimatedGraphicsItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskImageSpriteResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频截雪碧图任务的输入。
	Input *ImageSpriteTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频截雪碧图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaImageSpriteItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskInput struct {

	// 视频转码任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTaskSet []*TranscodeTaskInput `json:"TranscodeTaskSet,omitempty" name:"TranscodeTaskSet" list`

	// 视频转动图任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInput `json:"AnimatedGraphicTaskSet,omitempty" name:"AnimatedGraphicTaskSet" list`

	// 对视频按时间点截图任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTaskSet,omitempty" name:"SnapshotByTimeOffsetTaskSet" list`

	// 对视频采样截图任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotTaskSet []*SampleSnapshotTaskInput `json:"SampleSnapshotTaskSet,omitempty" name:"SampleSnapshotTaskSet" list`

	// 对视频截雪碧图任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteTaskSet []*ImageSpriteTaskInput `json:"ImageSpriteTaskSet,omitempty" name:"ImageSpriteTaskSet" list`

	// 对视频转自适应码流任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTaskSet,omitempty" name:"AdaptiveDynamicStreamingTaskSet" list`
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
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频转码任务的查询结果，当任务类型为 Transcode 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitempty" name:"TranscodeTask"`

	// 视频转动图任务的查询结果，当任务类型为 AnimatedGraphics 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitempty" name:"AnimatedGraphicTask"`

	// 对视频按时间点截图任务的查询结果，当任务类型为 SnapshotByTimeOffset 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitempty" name:"SnapshotByTimeOffsetTask"`

	// 对视频采样截图任务的查询结果，当任务类型为 SampleSnapshot 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitempty" name:"SampleSnapshotTask"`

	// 对视频截雪碧图任务的查询结果，当任务类型为 ImageSprite 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitempty" name:"ImageSpriteTask"`
}

type MediaProcessTaskSampleSnapshotResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频做采样截图任务输入。
	Input *SampleSnapshotTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频做采样截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSampleSnapshotItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskSnapshotByTimeOffsetResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频按指定时间点截图任务输入。
	Input *SnapshotByTimeOffsetTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频按指定时间点截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSnapshotByTimeOffsetItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskTranscodeResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 转码任务的输入。
	Input *TranscodeTaskInput `json:"Input,omitempty" name:"Input"`

	// 转码任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaTranscodeItem `json:"Output,omitempty" name:"Output"`
}

type MediaSampleSnapshotItem struct {

	// 采样截图规格 ID，参见[采样截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 采样方式，取值范围：
	// <li>Percent：根据百分比间隔采样。</li>
	// <li>Time：根据时间间隔采样。</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 采样间隔
	// <li>当 SampleType 为 Percent 时，该值表示多少百分比一张图。</li>
	// <li>当 SampleType 为 Time 时，该值表示多少时间间隔一张图，单位秒， 第一张图均为视频首帧。</li>
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 截图后文件的存储位置。
	Storage *TaskOutputStorage `json:"Storage,omitempty" name:"Storage"`

	// 生成的截图 path 列表。
	ImagePathSet []*string `json:"ImagePathSet,omitempty" name:"ImagePathSet" list`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition" list`
}

type MediaSnapshotByTimeOffsetItem struct {

	// 指定时间点截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 同一规格的截图信息集合，每个元素代表一张截图。
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet,omitempty" name:"PicInfoSet" list`

	// 指定时间点截图文件的存储位置。
	Storage *TaskOutputStorage `json:"Storage,omitempty" name:"Storage"`
}

type MediaSnapshotByTimePicInfoItem struct {

	// 该张截图对应视频文件中的时间偏移，单位为<font color=red>毫秒</font>。
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// 该张截图的路径。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition" list`
}

type MediaTranscodeItem struct {

	// 转码后文件的目标存储。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 转码后的视频文件路径。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/266/33478#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 媒体文件总大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 视频的 md5 值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 音频流信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet" list`

	// 视频流信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet" list`
}

type MediaVideoStreamItem struct {

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

type ModifyAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest

	// 转动图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 转动图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 动图宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 动图格式，取值为 gif 和 webp。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 帧率，取值范围：[1, 30]，单位：Hz。
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 图片质量，取值范围：[1, 100]，默认值为 75。
	Quality *float64 `json:"Quality,omitempty" name:"Quality"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *ModifyAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAnimatedGraphicsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest

	// 雪碧图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 雪碧图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 雪碧图中小图的宽度，取值范围： [128, 4096]，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 雪碧图中小图的高度，取值范围： [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 采样类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// 雪碧图中小图的行数。
	RowCount *uint64 `json:"RowCount,omitempty" name:"RowCount"`

	// 雪碧图中小图的列数。
	ColumnCount *uint64 `json:"ColumnCount,omitempty" name:"ColumnCount"`
}

func (r *ModifyImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSpriteTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageSpriteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSpriteTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 采样截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 图片宽度，取值范围： [128, 4096]，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 图片高度，取值范围： [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 采样截图类型，取值：
	// <li>Percent：按百分比。</li>
	// <li>Time：按时间间隔。</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 采样间隔。
	// <li>当 SampleType 为 Percent 时，指定采样间隔的百分比。</li>
	// <li>当 SampleType 为 Time 时，指定采样间隔的时间，单位为秒。</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// 图片格式，取值为 jpg 和 png。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *ModifySampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySampleSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySampleSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySampleSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest

	// 指定时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 指定时间点截图模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 图片宽度，取值范围： [128, 4096]，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 图片高度，取值范围： [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 图片格式，取值可以为 jpg 和 png。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *ModifySnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySnapshotByTimeOffsetTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字节。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 视频流配置参数。
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// 音频流配置参数。
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitempty" name:"AudioTemplate"`
}

func (r *ModifyTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	// 目前，当 Type 为 image，该字段仅支持 TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 图片水印模板，该字段仅对图片水印模板有效。
	ImageTemplate *ImageWatermarkInputForUpdate `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// 文字水印模板，该字段仅对文字水印模板有效。
	TextTemplate *TextWatermarkTemplateInputForUpdate `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG水印模板，当 Type 为 svg，该字段必填。当 Type 为 image 或 text，该字段无效。
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitempty" name:"SvgTemplate"`
}

func (r *ModifyWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWatermarkTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图片水印地址，仅当 ImageTemplate.ImageContent 非空，该字段有效。
		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWatermarkTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NumberFormat struct {

	// `{number}`变量的起始值，默认为0。
	InitialValue *uint64 `json:"InitialValue,omitempty" name:"InitialValue"`

	// `{number}`变量的增长步长，默认为1。
	Increment *uint64 `json:"Increment,omitempty" name:"Increment"`

	// `{number}`变量的最小长度，不足时补占位符。默认为1。
	MinLength *uint64 `json:"MinLength,omitempty" name:"MinLength"`

	// `{number}`变量的长度不足时，补充的占位符。默认为"0"。
	PlaceHolder *string `json:"PlaceHolder,omitempty" name:"PlaceHolder"`
}

type ProcessMediaRequest struct {
	*tchttp.BaseRequest

	// 视频处理的文件输入信息。
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// 视频处理输出文件的目标存储。不填则继承 InputInfo 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 视频处理生成的文件输出的目标目录，如`/movie/201907/`。如果不填，表示与 InputInfo 中文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

	// 任务流的优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

func (r *ProcessMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetWorkflowRequest struct {
	*tchttp.BaseRequest

	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 工作流名称，最多128字符。同一个用户该名称唯一。
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 工作流绑定的触发规则，当上传视频命中该规则到该对象时即触发工作流。
	Trigger *WorkflowTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// 视频处理的文件输出配置。不填则继承 Trigger 中的存储位置。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 视频处理生成的文件输出的目标目录，如`/movie/201907/`。如果不填，表示与触发文件所在的目录一致，即`{inputDir}`。
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TaskPriority *int64 `json:"TaskPriority,omitempty" name:"TaskPriority"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`
}

func (r *ResetWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetWorkflowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetWorkflowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SampleSnapshotTaskInput struct {

	// 采样截图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`

	// 采样截图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 采样截图后图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_sampleSnapshot_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 采样截图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type SampleSnapshotTemplate struct {

	// 采样截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 采样截图模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 图片宽度。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 图片高度。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 图片格式。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 采样截图类型。
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 采样间隔。
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SnapshotByTimeOffsetTaskInput struct {

	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 截图时间点列表，单位为<font color=red>毫秒</font>。
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitempty" name:"TimeOffsetSet" list`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`

	// 时间点截图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 时间点截图后图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_snapshotByTimeOffset_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 时间点截图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type SnapshotByTimeOffsetTemplate struct {

	// 时间点截图模板唯一标识。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 模板类型，取值范围：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间点截图模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 图片宽度。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 图片高度。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 图片格式。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
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
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度，支持 px，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素；当填 0px 且
	//  Width 不为 0px 时，表示水印的高度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的高度取原始 SVG 图像的高度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Height 为视频宽度的百分比大小，如 10W% 表示 Height 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Height 为视频高度的百分比大小，如 10H% 表示 Height 为��频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Height 为视频短边的百分比大小，如 10S% 表示 Height 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Height 为视频长边的百分比大小，如 10L% 表示 Height 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 H%。</li>
	// 默认值为 0px。
	Height *string `json:"Height,omitempty" name:"Height"`
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
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度，支持 px，%，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像��；当填 0px 且
	//  Width 不为 0px 时，表示水印的高度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的高度取原始 SVG 图像的高度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Height 为视频宽度的百分比大小，如 10W% 表示 Height 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Height 为视频高度的百分比大小，如 10H% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Height 为视频短边的百分比大小，如 10S% 表示 Height 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Height 为视频长边的百分比大小，如 10L% 表示 Height 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 H%。
	// 默认值为 0px。
	Height *string `json:"Height,omitempty" name:"Height"`
}

type TaskNotifyConfig struct {

	// CMQ 的模型，有 Queue 和 Topic 两种。
	CmqModel *string `json:"CmqModel,omitempty" name:"CmqModel"`

	// CMQ 的园区，如 sh，bj 等。
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// 当模型为 Queue 时有效，表示接收事件通知的 CMQ 的队列名。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 当模型为 Topic 时有效，表示接收事件通知的 CMQ 的主题名。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 工作流通知的模式，可取值有 Finish 和 Change，不填代表 Finish。
	NotifyMode *string `json:"NotifyMode,omitempty" name:"NotifyMode"`
}

type TaskOutputStorage struct {

	// 视频处理输出对象存储位置的类型，现在仅支持 COS。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当 Type 为 COS 时有效，则该项为必填，表示视频处理 COS 输出位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosOutputStorage *CosOutputStorage `json:"CosOutputStorage,omitempty" name:"CosOutputStorage"`
}

type TaskSimpleInfo struct {

	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型，目前仅支持 WorkflowTask。
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务开始执行时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。若任务尚未开始，该字段为：0000-00-00T00:00:00Z。
	BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

	// 任务结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。若任务尚未完成，该字段为：0000-00-00T00:00:00Z。
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
}

type TextWatermarkTemplateInput struct {

	// 字体类型，目前可以支持两种：
	// <li>simkai.ttf：可以支持中文和英文；</li>
	// <li>arial.ttf：仅支持英文。</li>
	FontType *string `json:"FontType,omitempty" name:"FontType"`

	// 字体大小，格式：Npx，N 为数值。
	FontSize *string `json:"FontSize,omitempty" name:"FontSize"`

	// 字体颜色，格式：0xRRGGBB，默认值：0xFFFFFF（白色）。
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// 文字透明度，取值范围：(0, 1]
	// <li>0：完全透明</li>
	// <li>1：完全不透明</li>
	// 默认值：1。
	FontAlpha *float64 `json:"FontAlpha,omitempty" name:"FontAlpha"`
}

type TextWatermarkTemplateInputForUpdate struct {

	// 字体类型，目前可以支持两种：
	// <li>simkai.ttf：可以支持中文和英文；</li>
	// <li>arial.ttf：仅支持英文。</li>
	FontType *string `json:"FontType,omitempty" name:"FontType"`

	// 字体大小，格式：Npx，N 为数值。
	FontSize *string `json:"FontSize,omitempty" name:"FontSize"`

	// 字体颜色，格式：0xRRGGBB，默认值：0xFFFFFF（白色）。
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// 文字透明度，取值范围：(0, 1]
	// <li>0：完全透明</li>
	// <li>1：完全不透明</li>
	FontAlpha *float64 `json:"FontAlpha,omitempty" name:"FontAlpha"`
}

type TranscodeTaskInput struct {

	// 视频转码模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`

	// 转码后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 转码后主文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_transcode_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 转码后分片文件的输出路径（转码 HLS 时 ts 的路径），只能为相对路径。如果不填，则默认为：`{inputName}_transcode_{definition}_{number}.{format}`。
	SegmentObjectName *string `json:"SegmentObjectName,omitempty" name:"SegmentObjectName"`

	// 转码后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type TranscodeTemplate struct {

	// 转码模板唯一标识。
	Definition *string `json:"Definition,omitempty" name:"Definition"`

	// 封装格式，取值：mp4、flv、hls、mp3、flac、ogg。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 转码模板名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 模板类型，取值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否去除视频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 视频流配置参数，仅当 RemoveVideo 为 0，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，仅当 RemoveAudio 为 0，该字段有效 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式板。</li>
	ContainerType *string `json:"ContainerType,omitempty" name:"ContainerType"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type VideoTemplateInfo struct {

	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码</li>
	// <li>libx265：H.265 编码</li>
	// 目前 H.265 编码必须指定分辨率，并且需要在 640*480 以内。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 视频帧率，取值范围：[0, 60]，单位：Hz。
	// 当取值为 0，表示帧率和原始视频保持一致。
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type VideoTemplateInfoForUpdate struct {

	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码</li>
	// <li>libx265：H.265 编码</li>
	// 目前 H.265 编码必须指定分辨率，并且需要在 640*480 以内。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 视频帧率，取值范围：[0, 60]，单位：Hz。
	// 当取值为 0，表示帧率和原始视频保持一致。
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type WatermarkInput struct {

	// 水印模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 文字内容，长度不超过100个字符。仅当水印类型为文字水印时填写。
	TextContent *string `json:"TextContent,omitempty" name:"TextContent"`

	// SVG 内容。长度不超过 2000000 个字符。仅当水印类型为 SVG 水印时填写。
	SvgContent *string `json:"SvgContent,omitempty" name:"SvgContent"`
}

type WatermarkTemplate struct {

	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 水印类型，取值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 水印模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 水印图片原点距离视频图像原点的水平位置。
	// <li>当字符串以 % 结尾，表示水印 Left 为视频宽度指定百分比的位置，如 10% 表示 Left 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Left 为视频宽度指定像素的位置，如 100px 表示 Left 为 100 像素。</li>
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 水印图片原点距离视频图像原点的垂直位置。
	// <li>当字符串以 % 结尾，表示水印 Top 为视频高度指定百分比的位置，如 10% 表示 Top 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Top 为视频高度指定像素的位置，如 100px 表示 Top 为 100 像素。</li>
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 图片水印模板，仅当 Type 为 image，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTemplate *ImageWatermarkTemplate `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// 文字水印模板，仅当 Type 为 text，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG 水印模板，当 Type 为 svg，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitempty" name:"SvgTemplate"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 原点位置，可选值：
	// <li>topLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>topRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>bottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>bottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下。；</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`
}

type WorkflowInfo struct {

	// 工作流 ID。
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 工作流名称。
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 工作流状态，取值范围：
	// <li>Enabled：已启用，</li>
	// <li>Disabled：已禁用。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 工作流绑定的输入规则，当上传视频命中该规则到该对象时即触发工作流。
	Trigger *WorkflowTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// 视频处理的文件输出存储位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 视频处理类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// 任务的事件通知信息，不填代表不获取事件通知。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TaskPriority *int64 `json:"TaskPriority,omitempty" name:"TaskPriority"`

	// 视频处理生成的文件输出的目标目录，如`/movie/201907/`。
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// 工作流创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 工作流最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type WorkflowTask struct {

	// 视频处理任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 视频处理的目标文件信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// 原始视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 视频处理任务的执行状态与结果。
	MediaProcessResultSet []*MediaProcessTaskResult `json:"MediaProcessResultSet,omitempty" name:"MediaProcessResultSet" list`
}

type WorkflowTrigger struct {

	// 触发器的类型，目前仅支持 CosFileUpload。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当 Type 为 CosFileUpload 时必填且有效，为 COS 触发规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosFileUploadTrigger *CosFileUploadTrigger `json:"CosFileUploadTrigger,omitempty" name:"CosFileUploadTrigger"`
}
