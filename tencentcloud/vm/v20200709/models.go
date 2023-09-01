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

package v20200709

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AudioResult struct {
	// 是否命中
	// 0 未命中
	// 1 命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitnil" name:"HitFlag"`

	// 命中的标签
	// Porn 色情
	// Polity 政治
	// Illegal 违法
	// Abuse 谩骂
	// Terror 暴恐
	// Ad 广告
	// Moan 呻吟
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 得分，0-100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil" name:"Score"`

	// 音频ASR文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 音频片段存储URL，有效期为1天
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 音频时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitnil" name:"Duration"`

	// 拓展字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 文本审核结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextResults []*AudioResultDetailTextResult `json:"TextResults,omitnil" name:"TextResults"`

	// 音频呻吟审核结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	MoanResults []*AudioResultDetailMoanResult `json:"MoanResults,omitnil" name:"MoanResults"`

	// 音频语种检测结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanguageResults []*AudioResultDetailLanguageResult `json:"LanguageResults,omitnil" name:"LanguageResults"`
}

type AudioResultDetailLanguageResult struct {
	// 语种
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil" name:"Score"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *float64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *float64 `json:"EndTime,omitnil" name:"EndTime"`

	// 子标签码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabelCode *string `json:"SubLabelCode,omitnil" name:"SubLabelCode"`
}

type AudioResultDetailMoanResult struct {
	// 固定为Moan
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 分数
	Score *int64 `json:"Score,omitnil" name:"Score"`

	// 开始时间
	StartTime *float64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *float64 `json:"EndTime,omitnil" name:"EndTime"`

	// 子标签码
	SubLabelCode *string `json:"SubLabelCode,omitnil" name:"SubLabelCode"`
}

type AudioResultDetailTextResult struct {
	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 命中的关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitnil" name:"Keywords"`

	// 命中的LibId
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibId *string `json:"LibId,omitnil" name:"LibId"`

	// 命中的LibName
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitnil" name:"LibName"`

	// 得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil" name:"Score"`

	// 词库类型 1 黑白库 2 自定义库
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibType *int64 `json:"LibType,omitnil" name:"LibType"`

	// 审核建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`
}

type AudioSegments struct {
	// 截帧时间。
	// 点播文件：该值为相对于视频偏移时间，单位为秒，例如：0，5，10
	// 直播流：该值为时间戳，例如：1594650717
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetTime *string `json:"OffsetTime,omitnil" name:"OffsetTime"`

	// 结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AudioResult `json:"Result,omitnil" name:"Result"`
}

type BucketInfo struct {
	// 腾讯云对象存储，存储桶名称
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`

	// 地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 对象Key
	Object *string `json:"Object,omitnil" name:"Object"`
}

// Predefined struct for user
type CancelTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *CancelTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

// Predefined struct for user
type CancelTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelTaskResponseParams `json:"Response"`
}

func (r *CancelTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBizConfigRequestParams struct {
	// 业务ID，仅限英文字母、数字和下划线（_）组成，长度不超过8位
	BizType *string `json:"BizType,omitnil" name:"BizType"`

	// 审核分类信息
	MediaModeration *MediaModerationConfig `json:"MediaModeration,omitnil" name:"MediaModeration"`

	// 业务名称，用于标识业务场景，长度不超过32位
	BizName *string `json:"BizName,omitnil" name:"BizName"`

	// 审核内容，可选：Polity (政治); Porn (色情); Illegal(违法);Abuse (谩骂); Terror (暴恐); Ad (广告); Custom (自定义);
	ModerationCategories []*string `json:"ModerationCategories,omitnil" name:"ModerationCategories"`
}

type CreateBizConfigRequest struct {
	*tchttp.BaseRequest
	
	// 业务ID，仅限英文字母、数字和下划线（_）组成，长度不超过8位
	BizType *string `json:"BizType,omitnil" name:"BizType"`

	// 审核分类信息
	MediaModeration *MediaModerationConfig `json:"MediaModeration,omitnil" name:"MediaModeration"`

	// 业务名称，用于标识业务场景，长度不超过32位
	BizName *string `json:"BizName,omitnil" name:"BizName"`

	// 审核内容，可选：Polity (政治); Porn (色情); Illegal(违法);Abuse (谩骂); Terror (暴恐); Ad (广告); Custom (自定义);
	ModerationCategories []*string `json:"ModerationCategories,omitnil" name:"ModerationCategories"`
}

func (r *CreateBizConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBizConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizType")
	delete(f, "MediaModeration")
	delete(f, "BizName")
	delete(f, "ModerationCategories")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBizConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBizConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBizConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateBizConfigResponseParams `json:"Response"`
}

func (r *CreateBizConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBizConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoModerationTaskRequestParams struct {
	// 业务类型, 定义 模版策略，输出存储配置。如果没有BizType，可以先参考 【创建业务配置】接口进行创建
	BizType *string `json:"BizType,omitnil" name:"BizType"`

	// 任务类型：可选VIDEO（点播视频），LIVE_VIDEO（直播视频）
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入的任务信息，最多可以同时创建10个任务
	Tasks []*TaskInput `json:"Tasks,omitnil" name:"Tasks"`

	// 回调签名key，具体可以查看签名文档。
	Seed *string `json:"Seed,omitnil" name:"Seed"`

	// 接收审核信息回调地址，如果设置，则审核过程中产生的违规音频片段和画面截帧发送此接口
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 审核排队优先级。当您有多个视频审核任务排队时，可以根据这个参数控制排队优先级。用于处理插队等逻辑。默认该参数为0
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`
}

type CreateVideoModerationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 业务类型, 定义 模版策略，输出存储配置。如果没有BizType，可以先参考 【创建业务配置】接口进行创建
	BizType *string `json:"BizType,omitnil" name:"BizType"`

	// 任务类型：可选VIDEO（点播视频），LIVE_VIDEO（直播视频）
	Type *string `json:"Type,omitnil" name:"Type"`

	// 输入的任务信息，最多可以同时创建10个任务
	Tasks []*TaskInput `json:"Tasks,omitnil" name:"Tasks"`

	// 回调签名key，具体可以查看签名文档。
	Seed *string `json:"Seed,omitnil" name:"Seed"`

	// 接收审核信息回调地址，如果设置，则审核过程中产生的违规音频片段和画面截帧发送此接口
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 审核排队优先级。当您有多个视频审核任务排队时，可以根据这个参数控制排队优先级。用于处理插队等逻辑。默认该参数为0
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`
}

func (r *CreateVideoModerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoModerationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizType")
	delete(f, "Type")
	delete(f, "Tasks")
	delete(f, "Seed")
	delete(f, "CallbackUrl")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoModerationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoModerationTaskResponseParams struct {
	// 任务创建结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*TaskResult `json:"Results,omitnil" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVideoModerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateVideoModerationTaskResponseParams `json:"Response"`
}

func (r *CreateVideoModerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoModerationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// 任务ID，创建任务后返回的TaskId字段
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 是否展示所有分片，默认只展示命中规则的分片
	ShowAllSegments *bool `json:"ShowAllSegments,omitnil" name:"ShowAllSegments"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，创建任务后返回的TaskId字段
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 是否展示所有分片，默认只展示命中规则的分片
	ShowAllSegments *bool `json:"ShowAllSegments,omitnil" name:"ShowAllSegments"`
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
	delete(f, "ShowAllSegments")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 审核时传入的数据Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataId *string `json:"DataId,omitnil" name:"DataId"`

	// 业务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizType *string `json:"BizType,omitnil" name:"BizType"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 状态，可选值：
	// FINISH 已完成
	// PENDING 等待中
	// RUNNING 进行中
	// ERROR 出错
	// CANCELLED 已取消
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 审核建议
	// 可选：
	// Pass 通过
	// Reveiw 建议复审
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 审核结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*TaskLabel `json:"Labels,omitnil" name:"Labels"`

	// 媒体解码信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaInfo *MediaInfo `json:"MediaInfo,omitnil" name:"MediaInfo"`

	// 任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputInfo *InputInfo `json:"InputInfo,omitnil" name:"InputInfo"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// 在秒后重试
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryInSeconds *int64 `json:"TryInSeconds,omitnil" name:"TryInSeconds"`

	// 图片结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSegments []*ImageSegments `json:"ImageSegments,omitnil" name:"ImageSegments"`

	// 音频结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioSegments []*AudioSegments `json:"AudioSegments,omitnil" name:"AudioSegments"`

	// 如果返回的状态为ERROR，该字段会标记错误类型。
	// 可选值：：
	// DECODE_ERROR: 解码失败。（输入资源中可能包含无法解码的视频）
	// URL_ERROR：下载地址验证失败。
	// TIMEOUT_ERROR：处理超时。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorType *string `json:"ErrorType,omitnil" name:"ErrorType"`

	// 审核任务错误日志。当Error不为空时，会展示该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDescription *string `json:"ErrorDescription,omitnil" name:"ErrorDescription"`

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
type DescribeVideoStatRequestParams struct {
	// 审核类型 1: 机器审核; 2: 人工审核
	AuditType *int64 `json:"AuditType,omitnil" name:"AuditType"`

	// 查询条件
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type DescribeVideoStatRequest struct {
	*tchttp.BaseRequest
	
	// 审核类型 1: 机器审核; 2: 人工审核
	AuditType *int64 `json:"AuditType,omitnil" name:"AuditType"`

	// 查询条件
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeVideoStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoStatResponseParams struct {
	// 识别结果统计
	Overview *Overview `json:"Overview,omitnil" name:"Overview"`

	// 识别量统计
	TrendCount []*TrendCount `json:"TrendCount,omitnil" name:"TrendCount"`

	// 违规数据分布
	EvilCount []*EvilCount `json:"EvilCount,omitnil" name:"EvilCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVideoStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoStatResponseParams `json:"Response"`
}

func (r *DescribeVideoStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EvilCount struct {
	// 违规类型：
	// Terror	24001
	// Porn	20002
	// Polity	20001
	// Ad	20105
	// Abuse	20007	
	// Illegal	20006	
	// Spam	25001	
	// Moan	26001
	EvilType *string `json:"EvilType,omitnil" name:"EvilType"`

	// 分布类型总量
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type FileOutput struct {
	// 存储的Bucket
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`

	// Cos Region
	Region *string `json:"Region,omitnil" name:"Region"`

	// 对象前缀
	ObjectPrefix *string `json:"ObjectPrefix,omitnil" name:"ObjectPrefix"`
}

type Filters struct {
	// 查询字段：
	// 策略BizType
	// 子账号SubUin
	// 日期区间DateRange
	Name *string `json:"Name,omitnil" name:"Name"`

	// 查询值
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type ImageResult struct {
	// 违规标志
	// 0 未命中
	// 1 命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitnil" name:"HitFlag"`

	// 命中的标签
	// Porn 色情
	// Sexy 性感
	// Polity 政治
	// Illegal 违法
	// Abuse 谩骂
	// Terror 暴恐
	// Ad 广告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil" name:"Score"`

	// 画面截帧图片结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*ImageResultResult `json:"Results,omitnil" name:"Results"`

	// 图片URL地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 附加字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`
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
	Scene *string `json:"Scene,omitnil" name:"Scene"`

	// 是否命中
	// 0 未命中
	// 1 命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitnil" name:"HitFlag"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 子标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitnil" name:"SubLabel"`

	// 分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil" name:"Score"`

	// 如果命中场景为涉政，则该数据为人物姓名列表，否则null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Names []*string `json:"Names,omitnil" name:"Names"`

	// 图片OCR文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 其他详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*ImageResultsResultDetail `json:"Details,omitnil" name:"Details"`
}

type ImageResultsResultDetail struct {
	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// OCR识别文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 位置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *ImageResultsResultDetailLocation `json:"Location,omitnil" name:"Location"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibId *string `json:"LibId,omitnil" name:"LibId"`

	// 库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitnil" name:"LibName"`

	// 命中的关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitnil" name:"Keywords"`

	// 建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil" name:"Score"`

	// 子标签码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabelCode *string `json:"SubLabelCode,omitnil" name:"SubLabelCode"`
}

type ImageResultsResultDetailLocation struct {
	// x坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *float64 `json:"X,omitnil" name:"X"`

	// y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *float64 `json:"Y,omitnil" name:"Y"`

	// 宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitnil" name:"Height"`

	// 旋转角度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rotate *float64 `json:"Rotate,omitnil" name:"Rotate"`
}

type ImageSegments struct {
	// 截帧时间。
	// 点播文件：该值为相对于视频偏移时间，单位为秒，例如：0，5，10
	// 直播流：该值为时间戳，例如：1594650717
	OffsetTime *string `json:"OffsetTime,omitnil" name:"OffsetTime"`

	// 画面截帧结果详情
	Result *ImageResult `json:"Result,omitnil" name:"Result"`
}

type InputInfo struct {
	// 传入的类型可选：URL，COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Url地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 桶信息。当输入当时COS时，该字段不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketInfo *string `json:"BucketInfo,omitnil" name:"BucketInfo"`
}

type MediaInfo struct {
	// 编码格式
	Codecs *string `json:"Codecs,omitnil" name:"Codecs"`

	// 流检测时分片时长
	// 注意：此字段可能返回 0，表示取不到有效值。
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// 宽，单位为像素
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 高，单位为像素
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type MediaModerationConfig struct {
	// 是否使用OCR，默认为true
	UseOCR *bool `json:"UseOCR,omitnil" name:"UseOCR"`

	// 是否使用音频，默认为true。视频场景下，默认为 false
	UseAudio *bool `json:"UseAudio,omitnil" name:"UseAudio"`

	// 图片取帧频率, 单位（秒/帧），默认 5， 可选 1 ～ 300
	ImageFrequency *int64 `json:"ImageFrequency,omitnil" name:"ImageFrequency"`

	// 音频片段长度。单位为：秒
	AudioFrequency *int64 `json:"AudioFrequency,omitnil" name:"AudioFrequency"`

	// 临时文件存储位置
	SegmentOutput *FileOutput `json:"SegmentOutput,omitnil" name:"SegmentOutput"`

	// 回调地址
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`
}

type Overview struct {
	// 总调用量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 总调用时长
	TotalHour *int64 `json:"TotalHour,omitnil" name:"TotalHour"`

	// 通过量
	PassCount *int64 `json:"PassCount,omitnil" name:"PassCount"`

	// 通过时长
	PassHour *int64 `json:"PassHour,omitnil" name:"PassHour"`

	// 违规量
	EvilCount *int64 `json:"EvilCount,omitnil" name:"EvilCount"`

	// 违规时长
	EvilHour *int64 `json:"EvilHour,omitnil" name:"EvilHour"`

	// 疑似违规量
	SuspectCount *int64 `json:"SuspectCount,omitnil" name:"SuspectCount"`

	// 疑似违规时长
	SuspectHour *int64 `json:"SuspectHour,omitnil" name:"SuspectHour"`
}

type StorageInfo struct {
	// 类型 可选：
	// URL 资源链接类型
	// COS 腾讯云对象存储类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 资源链接
	Url *string `json:"Url,omitnil" name:"Url"`

	// 腾讯云存储桶信息
	BucketInfo *BucketInfo `json:"BucketInfo,omitnil" name:"BucketInfo"`
}

type TaskInput struct {
	// 数据ID
	DataId *string `json:"DataId,omitnil" name:"DataId"`

	// 任务名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 任务输入
	Input *StorageInfo `json:"Input,omitnil" name:"Input"`
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
	Label *string `json:"Label,omitnil" name:"Label"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 得分，分数是 0 ～ 100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil" name:"Score"`
}

type TaskResult struct {
	// 请求时传入的DataId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataId *string `json:"DataId,omitnil" name:"DataId"`

	// TaskId，任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 错误码。如果code为OK，则表示创建成功，其他则参考公共错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 如果错误，该字段表示错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`
}

type TrendCount struct {
	// 总调用量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 总调用时长
	TotalHour *int64 `json:"TotalHour,omitnil" name:"TotalHour"`

	// 通过量
	PassCount *int64 `json:"PassCount,omitnil" name:"PassCount"`

	// 通过时长
	PassHour *int64 `json:"PassHour,omitnil" name:"PassHour"`

	// 违规量
	EvilCount *int64 `json:"EvilCount,omitnil" name:"EvilCount"`

	// 违规时长
	EvilHour *int64 `json:"EvilHour,omitnil" name:"EvilHour"`

	// 疑似违规量
	SuspectCount *int64 `json:"SuspectCount,omitnil" name:"SuspectCount"`

	// 疑似违规时长
	SuspectHour *int64 `json:"SuspectHour,omitnil" name:"SuspectHour"`

	// 日期
	Date *string `json:"Date,omitnil" name:"Date"`
}