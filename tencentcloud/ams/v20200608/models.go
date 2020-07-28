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

package v20200608

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AudioResult struct {

	// 是否命中
	// 0 未命中
	// 1 命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 命中的标签
	// Porn 色情
	// Polity 政治
	// Illegal 违法
	// Abuse 谩骂
	// Terror 暴恐
	// Ad 广告
	// Moan 呻吟
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 得分，0-100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 音频ASR文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 音频片段存储URL，有效期为1天
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 音频时长
	Duration *string `json:"Duration,omitempty" name:"Duration"`

	// 拓展字段
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 文本审核结果
	TextResults []*AudioResultDetailTextResult `json:"TextResults,omitempty" name:"TextResults" list`

	// 音频呻吟审核结果
	MoanResults []*AudioResultDetailMoanResult `json:"MoanResults,omitempty" name:"MoanResults" list`

	// 音频语种检测结果
	LanguageResults []*AudioResultDetailLanguageResult `json:"LanguageResults,omitempty" name:"LanguageResults" list`
}

type AudioResultDetailLanguageResult struct {

	// 语种
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *float64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *float64 `json:"EndTime,omitempty" name:"EndTime"`

	// 子标签码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabelCode *string `json:"SubLabelCode,omitempty" name:"SubLabelCode"`
}

type AudioResultDetailMoanResult struct {

	// 固定为Moan
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 分数
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 开始时间
	StartTime *float64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *float64 `json:"EndTime,omitempty" name:"EndTime"`

	// 子标签码
	SubLabelCode *string `json:"SubLabelCode,omitempty" name:"SubLabelCode"`
}

type AudioResultDetailTextResult struct {

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 命中的关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`

	// 命中的LibId
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// 命中的LibName
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// 得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 审核建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 词库类型 1 黑白库 2 自定义库
	LibType *int64 `json:"LibType,omitempty" name:"LibType"`
}

type AudioSegments struct {

	// 截帧时间。
	// 点播文件：该值为相对于视频偏移时间，单位为秒，例如：0，5，10
	// 直播流：该值为时间戳，例如：1594650717
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetTime *string `json:"OffsetTime,omitempty" name:"OffsetTime"`

	// 结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AudioResult `json:"Result,omitempty" name:"Result"`
}

type BucketInfo struct {

	// 腾讯云对象存储，存储桶名称
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 对象Key
	Object *string `json:"Object,omitempty" name:"Object"`
}

type CreateAudioModerationTaskRequest struct {
	*tchttp.BaseRequest

	// 业务类型, 定义 模版策略，输出存储配置。如果没有BizType，可以先参考 【创建业务配置】接口进行创建
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 异步检测结果回调通知接收URL。支持HTTP和HTTPS
	Type *string `json:"Type,omitempty" name:"Type"`

	// 回调签名key，具体可以查看签名文档。
	Seed *string `json:"Seed,omitempty" name:"Seed"`

	// 接收审核信息回调地址，如果设置，则审核过程中产生的违规音频片段和画面截帧发送此接口
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 输入的任务信息，最多可以同时创建10个任务
	Tasks []*TaskInput `json:"Tasks,omitempty" name:"Tasks" list`
}

func (r *CreateAudioModerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAudioModerationTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAudioModerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务创建结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Results []*TaskResult `json:"Results,omitempty" name:"Results" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAudioModerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAudioModerationTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBizConfigRequest struct {
	*tchttp.BaseRequest

	// 业务ID
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 审核分类信息
	MediaModeration *MediaModerationConfig `json:"MediaModeration,omitempty" name:"MediaModeration"`

	// 页面名称
	BizName *string `json:"BizName,omitempty" name:"BizName"`

	// 审核内容，可选：Polity (政治); Porn (色情); Illegal(违法);Abuse (谩骂); Terror (暴恐); Ad (广告);
	ModerationCategories []*string `json:"ModerationCategories,omitempty" name:"ModerationCategories" list`
}

func (r *CreateBizConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBizConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBizConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBizConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBizConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBizConfigRequest struct {
	*tchttp.BaseRequest

	// 审核业务类类型
	BizType *string `json:"BizType,omitempty" name:"BizType"`
}

func (r *DescribeBizConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBizConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBizConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务类型
		BizType *string `json:"BizType,omitempty" name:"BizType"`

		// 业务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		BizName *string `json:"BizName,omitempty" name:"BizName"`

		// 审核范围
		ModerationCategories []*string `json:"ModerationCategories,omitempty" name:"ModerationCategories" list`

		// 多媒体审核配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		MediaModeration *MediaModerationConfig `json:"MediaModeration,omitempty" name:"MediaModeration"`

		// 创建时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 更新时间
		UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBizConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBizConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 任务ID，创建任务后返回的TaskId字段
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 是否展示所有分片，默认只展示命中规则的分片
	ShowAllSegments *bool `json:"ShowAllSegments,omitempty" name:"ShowAllSegments"`
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

		// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 审核时传入的数据Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		DataId *string `json:"DataId,omitempty" name:"DataId"`

		// 业务类型，用于调用识别策略模板；
	// （暂未发布功能，敬请期待）
	// 注意：此字段可能返回 null，表示取不到有效值。
		BizType *string `json:"BizType,omitempty" name:"BizType"`

		// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		Name *string `json:"Name,omitempty" name:"Name"`

		// 查询内容审核任务的状态，可选值：
	// FINISH 已完成
	// PENDING 等待中
	// RUNNING 进行中
	// ERROR 出错
	// CANCELLED 已取消
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务类型：可选AUDIO（点播音频），LIVE_AUDIO（直播音频）
	// 注意：此字段可能返回 null，表示取不到有效值。
		Type *string `json:"Type,omitempty" name:"Type"`

		// 智能审核服务对于内容违规类型的等级，可选值：
	// Pass 建议通过；
	// Reveiw 建议复审；
	// Block 建议屏蔽；
	// 注意：此字段可能返回 null，表示取不到有效值。
		Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

		// 智能审核服务对于内容违规类型的判断，详见返回值列表
	// 如：Label：Porn（色情）；
	// 注意：此字段可能返回 null，表示取不到有效值。
		Labels []*TaskLabel `json:"Labels,omitempty" name:"Labels" list`

		// 传入媒体的解码信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		MediaInfo *MediaInfo `json:"MediaInfo,omitempty" name:"MediaInfo"`

		// 审核任务的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		InputInfo *InputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

		// 审核任务的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 审核任务的更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

		// 在N秒后重试
	// 注意：此字段可能返回 null，表示取不到有效值。
		TryInSeconds *int64 `json:"TryInSeconds,omitempty" name:"TryInSeconds"`

		// 视频/音频审核中的音频结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		AudioSegments []*AudioSegments `json:"AudioSegments,omitempty" name:"AudioSegments" list`

		// 视频审核中的图片结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImageSegments []*ImageSegments `json:"ImageSegments,omitempty" name:"ImageSegments" list`

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

type FileOutput struct {

	// 存储的Bucket
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Cos Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// 对象前缀
	ObjectPrefix *string `json:"ObjectPrefix,omitempty" name:"ObjectPrefix"`
}

type ImageResult struct {

	// 违规标志
	// 0 未命中
	// 1 命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 命中的标签
	// Porn 色情
	// Sexy 性感
	// Polity 政治
	// Illegal 违法
	// Abuse 谩骂
	// Terror 暴恐
	// Ad 广告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 得分
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 画面截帧图片结果集
	Results []*ImageResultResult `json:"Results,omitempty" name:"Results" list`

	// 图片URL地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 附加字段
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

type ImageResultResult struct {

	// 场景
	// Porn 色情
	// Sexy 性感
	// Polity 政治
	// Illegal 违法
	// Abuse 谩骂
	// Terror 暴恐
	// Ad 广告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 是否命中
	// 0 未命中
	// 1 命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 子标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitempty" name:"SubLabel"`

	// 分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 如果命中场景为涉政，则该数据为人物姓名列表，否则null
	Names []*string `json:"Names,omitempty" name:"Names" list`

	// 图片OCR文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 其他详情
	Details []*ImageResultsResultDetail `json:"Details,omitempty" name:"Details" list`
}

type ImageResultsResultDetail struct {

	// 位置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location []*ImageResultsResultDetailLocation `json:"Location,omitempty" name:"Location" list`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// OCR识别文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 标签
	Label *string `json:"Label,omitempty" name:"Label"`

	// 库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// 库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// 命中的关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`

	// 建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 子标签码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabelCode *string `json:"SubLabelCode,omitempty" name:"SubLabelCode"`
}

type ImageResultsResultDetailLocation struct {

	// x坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *float64 `json:"X,omitempty" name:"X"`

	// y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *float64 `json:"Y,omitempty" name:"Y"`

	// 宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 旋转角度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rotate *float64 `json:"Rotate,omitempty" name:"Rotate"`
}

type ImageSegments struct {

	// 画面截帧结果详情
	Result *ImageResult `json:"Result,omitempty" name:"Result"`

	// 截帧时间。
	// 点播文件：该值为相对于视频偏移时间，单位为秒，例如：0，5，10
	// 直播流：该值为时间戳，例如：1594650717
	OffsetTime *string `json:"OffsetTime,omitempty" name:"OffsetTime"`
}

type InputInfo struct {

	// 传入的类型可选：URL，COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// Url地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 桶信息。当输入当时COS时，该字段不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketInfo *BucketInfo `json:"BucketInfo,omitempty" name:"BucketInfo"`
}

type MediaInfo struct {

	// 编码格式
	Codecs *string `json:"Codecs,omitempty" name:"Codecs"`

	// 流检测时分片时长
	// 注意：此字段可能返回 0，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 宽，单位为像素
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高，单位为像素
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type MediaModerationConfig struct {

	// 音频截帧频率。默认一分钟
	AudioFrequency *int64 `json:"AudioFrequency,omitempty" name:"AudioFrequency"`

	// 图片取帧频率, 单位（秒/帧），默认 5， 可选 1 ～ 300
	ImageFrequency *int64 `json:"ImageFrequency,omitempty" name:"ImageFrequency"`

	// 异步回调地址。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 临时文件存储位置
	SegmentOutput *FileOutput `json:"SegmentOutput,omitempty" name:"SegmentOutput"`

	// 是否使用OCR，默认为true
	UseOCR *bool `json:"UseOCR,omitempty" name:"UseOCR"`

	// 是否使用音频。（音频场景下，该值永远为true）
	UseAudio *bool `json:"UseAudio,omitempty" name:"UseAudio"`
}

type StorageInfo struct {

	// 类型 可选：
	// URL 资源链接类型
	// COS 腾讯云对象存储类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 资源链接
	Url *string `json:"Url,omitempty" name:"Url"`

	// 腾讯云存储桶信息
	BucketInfo *BucketInfo `json:"BucketInfo,omitempty" name:"BucketInfo"`
}

type TaskInput struct {

	// 数据ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务输入
	Input *StorageInfo `json:"Input,omitempty" name:"Input"`
}

type TaskLabel struct {

	// 命中的标签
	// Porn 色情
	// Sexy 性感
	// Polity 政治
	// Illegal 违法
	// Abuse 谩骂
	// Terror 暴恐
	// Ad 广告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 得分，分数是 0 ～ 100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type TaskResult struct {

	// 请求时传入的DataId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// TaskId，任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 错误码。如果code为OK，则表示创建成功，其他则参考公共错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 如果错误，该字段表示错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}
