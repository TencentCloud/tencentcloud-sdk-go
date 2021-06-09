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

package v20201229

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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
	// Illegal 违法
	// Abuse 谩骂
	// Ad 广告
	// Moan 呻吟
	// 以及其他令人反感、不安全或不适宜的内容类型。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitempty" name:"Duration"`

	// 拓展字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 文本审核结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextResults []*AudioResultDetailTextResult `json:"TextResults,omitempty" name:"TextResults"`

	// 音频呻吟审核结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	MoanResults []*AudioResultDetailMoanResult `json:"MoanResults,omitempty" name:"MoanResults"`

	// 音频语言类别检测结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanguageResults []*AudioResultDetailLanguageResult `json:"LanguageResults,omitempty" name:"LanguageResults"`
}

type AudioResultDetailLanguageResult struct {

	// 语言分类，如中文、英文等；
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
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 命中的LibId
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// 命中的LibName
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// 得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 词库类型 1 黑白库 2 自定义库
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibType *int64 `json:"LibType,omitempty" name:"LibType"`

	// 审核建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
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

type CancelTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CancelTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVideoModerationTaskRequest struct {
	*tchttp.BaseRequest

	// 任务类型：可选VIDEO（点播视频），LIVE_VIDEO（直播视频）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 输入的任务信息，最多可以同时创建10个任务
	Tasks []*TaskInput `json:"Tasks,omitempty" name:"Tasks"`

	// 业务类型, 定义 模版策略，输出存储配置。默认为default，客户可以在视频审核控制台创建自己的 BizType
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// （可选）回调签名key，具体可以查看回调签名示例
	Seed *string `json:"Seed,omitempty" name:"Seed"`

	// 接收审核信息回调地址，如果设置，则审核过程中产生的违规音频片段和画面截帧发送此接口
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 审核排队优先级。当您有多个视频审核任务排队时，可以根据这个参数控制排队优先级。用于处理插队等逻辑。默认该参数为0
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *CreateVideoModerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoModerationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Tasks")
	delete(f, "BizType")
	delete(f, "Seed")
	delete(f, "CallbackUrl")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoModerationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateVideoModerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务创建结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Results []*TaskResult `json:"Results,omitempty" name:"Results"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVideoModerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoModerationTaskResponse) FromJsonString(s string) error {
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ShowAllSegments")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 审核时传入的数据Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		DataId *string `json:"DataId,omitempty" name:"DataId"`

		// 业务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		BizType *string `json:"BizType,omitempty" name:"BizType"`

		// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		Name *string `json:"Name,omitempty" name:"Name"`

		// 状态，可选值：
	// FINISH 已完成
	// PENDING 等待中
	// RUNNING 进行中
	// ERROR 出错
	// CANCELLED 已取消
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		Type *string `json:"Type,omitempty" name:"Type"`

		// 审核建议
	// 可选：
	// Pass 通过
	// Reveiw 建议复审
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
		Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

		// 审核结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Labels []*TaskLabel `json:"Labels,omitempty" name:"Labels"`

		// 媒体解码信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		MediaInfo *MediaInfo `json:"MediaInfo,omitempty" name:"MediaInfo"`

		// 任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		InputInfo *InputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

		// 创建时间，格式为 ISO 8601
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 更新时间，格式为 ISO 8601
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

		// 视频中截帧审核的结果。
	// 数据有效期为24小时，如需要储存更长时间请在已配置的COS储存桶中设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImageSegments []*ImageSegments `json:"ImageSegments,omitempty" name:"ImageSegments"`

		// 视频中音频审核的结果。
	// 数据有效期为24小时，如需要储存更长时间请在已配置的COS储存桶中设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AudioSegments []*AudioSegments `json:"AudioSegments,omitempty" name:"AudioSegments"`

		// 如果返回的状态为ERROR，该字段会标记错误类型。
	// 可选值：：
	// DECODE_ERROR: 解码失败。（输入资源中可能包含无法解码的视频）
	// URL_ERROR：下载地址验证失败。
	// TIMEOUT_ERROR：处理超时。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorType *string `json:"ErrorType,omitempty" name:"ErrorType"`

		// 审核任务错误日志。当Error不为空时，会展示该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorDescription *string `json:"ErrorDescription,omitempty" name:"ErrorDescription"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 每页展示多少条。（默认展示10条）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤参数
	Filter *TaskFilter `json:"Filter,omitempty" name:"Filter"`

	// 翻页token，在向前或向后翻页时需要
	PageToken *string `json:"PageToken,omitempty" name:"PageToken"`

	// 开始时间。默认是最近3天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。默认为空
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Filter")
	delete(f, "PageToken")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务总量，为 int 字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *string `json:"Total,omitempty" name:"Total"`

		// 当前页数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*TaskData `json:"Data,omitempty" name:"Data"`

		// 翻页Token，当已经到最后一页时，该字段为空
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageToken *string `json:"PageToken,omitempty" name:"PageToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// Illegal 违法
	// Abuse 谩骂
	// Ad 广告
	// 以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 画面截帧图片结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*ImageResultResult `json:"Results,omitempty" name:"Results"`

	// 图片URL地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 附加字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

type ImageResultResult struct {

	// 场景
	// Porn 色情
	// Sexy 性感
	// Illegal 违法
	// Abuse 谩骂
	// Ad 广告
	// 以及其他令人反感、不安全或不适宜的内容类型。
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

	// 识别结果一级标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 识别结果二级标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitempty" name:"SubLabel"`

	// 分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 适应特定场景，则该数据为名称列表，否则为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Names []*string `json:"Names,omitempty" name:"Names"`

	// 图片OCR文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 其他详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*ImageResultsResultDetail `json:"Details,omitempty" name:"Details"`
}

type ImageResultsResultDetail struct {

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// OCR识别文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 位置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *ImageResultsResultDetailLocation `json:"Location,omitempty" name:"Location"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// 库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// 命中的关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

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

	// 截帧时间。单位为秒。
	// 点播文件：该值为相对于视频偏移时间，单位为秒，例如：0，5，10
	// 直播流：该值为时间戳，例如：1594650717
	OffsetTime *string `json:"OffsetTime,omitempty" name:"OffsetTime"`

	// 画面截帧结果详情
	Result *ImageResult `json:"Result,omitempty" name:"Result"`
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
	BucketInfo *string `json:"BucketInfo,omitempty" name:"BucketInfo"`
}

type MediaInfo struct {

	// 流检测时分片时长
	// 注意：此字段可能返回 0，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
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

type TaskData struct {

	// 输入的数据ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 状态，可选：PENDING，RUNNING，ERROR，FINISH，CANCELLED
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 业务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 建议，可选：Pass， Block 和 Review
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 任务违规标签
	Labels []*TaskLabel `json:"Labels,omitempty" name:"Labels"`

	// 媒体信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaInfo *MediaInfo `json:"MediaInfo,omitempty" name:"MediaInfo"`

	// 创建时间（ iso 8601 格式）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（ iso 8601 格式）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type TaskFilter struct {

	// 任务业务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizType []*string `json:"BizType,omitempty" name:"BizType"`

	// 任务类型，可选：VIDEO，AUDIO， LIVE_VIDEO, LIVE_AUDIO
	Type *string `json:"Type,omitempty" name:"Type"`

	// 建议，可选：Pass, Review,Block
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 状态，可选：PENDING，RUNNING，ERROR，FINISH，CANCELLED
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
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
	// Illegal 违法
	// Abuse 谩骂
	// Ad 广告
	// 以及其他令人反感、不安全或不适宜的内容类型。
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
