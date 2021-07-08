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

	// 该字段用于返回审核内容是否命中审核模型；取值：0（**未命中**）、1（**命中**）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示具体的后续建议操作。<br>
	// 返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 该字段用于返回当前标签下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 该字段用于返回音频文件经ASR识别后的文本信息。最长可识别**5小时**的音频文件，若超出时长限制，接口将会报错。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 该字段用于返回音频片段存储的链接地址，该地址有效期为1天。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 该字段用于返回音频文件的时长，单位为秒。
	Duration *string `json:"Duration,omitempty" name:"Duration"`

	// 该字段用于返回额外附加信息，不同客户或Biztype下返回信息不同。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 该字段用于返回音频文件经ASR识别后产生的文本的详细审核结果。具体结果内容请参见AudioResultDetailLanguageResult数据结构的细节描述。
	TextResults []*AudioResultDetailTextResult `json:"TextResults,omitempty" name:"TextResults"`

	// 该字段用于返回音频文件呻吟检测的详细审核结果。具体结果内容请参见AudioResultDetailMoanResult数据结构的细节描述。
	MoanResults []*AudioResultDetailMoanResult `json:"MoanResults,omitempty" name:"MoanResults"`

	// 该字段用于返回音频小语种检测的详细审核结果。具体结果内容请参见AudioResultDetailLanguageResult数据结构的细节描述。
	LanguageResults []*AudioResultDetailLanguageResult `json:"LanguageResults,omitempty" name:"LanguageResults"`
}

type AudioResultDetailLanguageResult struct {

	// 该字段用于返回语种检测结果所对应的语种标签，目前支持：**Arabic**（阿拉伯语）、**English**（英语）、**Mandarin**（普通话）、**Tibetan**（藏语）、**Uyghur**（维语）、**Other**（其他上面5类之外）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 该参数用于返回当前标签下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高**），越高代表音频越有可能属于当前返回的语种标签；如：*Uyghur 99*，则表明该音频非常有可能属于维语内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 该参数用于返回对应语种标签的片段在音频文件内的开始时间，单位为毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *float64 `json:"StartTime,omitempty" name:"StartTime"`

	// 该参数用于返回对应语种标签的片段在音频文件内的结束时间，单位为毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *float64 `json:"EndTime,omitempty" name:"EndTime"`

	// *内测中，敬请期待*
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabelCode *string `json:"SubLabelCode,omitempty" name:"SubLabelCode"`
}

type AudioResultDetailMoanResult struct {

	// 该字段用于返回检测结果需要检测的内容类型，此处固定为**Moan**（呻吟）以调用呻吟检测功能。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 该字段用于返回呻吟检测的置信度，取值范围：0（**置信度最低**）-100（**置信度最高**），越高代表音频越有可能属于呻吟内容。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 该字段用于返回对应呻吟标签的片段在音频文件内的开始时间，单位为毫秒。
	StartTime *float64 `json:"StartTime,omitempty" name:"StartTime"`

	// 该字段用于返回对应呻吟标签的片段在音频文件内的结束时间，单位为毫秒。
	EndTime *float64 `json:"EndTime,omitempty" name:"EndTime"`

	// *内测中，敬请期待*
	SubLabelCode *string `json:"SubLabelCode,omitempty" name:"SubLabelCode"`
}

type AudioResultDetailTextResult struct {

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 该字段用于返回ASR识别出的文本内容命中的关键词信息，用于标注内容违规的具体原因（如：加我微信）。该参数可能会有多个返回值，代表命中的多个关键词；若返回值为空，Score不为空，则代表识别结果所对应的恶意标签（Label）来自于语义模型判断的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 该字段**仅当Label为Custom：自定义关键词时该参数有效**,用于返回自定义库的ID，以方便自定义库管理和配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// 该字段**仅当Label为Custom：自定义关键词时该参数有效**,用于返回自定义库的名称,以方便自定义库管理和配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// 该字段用于返回当前标签下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高**），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示具体的后续建议操作。<br>
	// 返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 该字段用于返回自定义关键词对应的词库类型，取值为**1**（黑白库）和**2**（自定义关键词库），若未配置自定义关键词库,则默认值为1（黑白库匹配）。
	LibType *int64 `json:"LibType,omitempty" name:"LibType"`
}

type AudioSegments struct {

	// 该字段用于返回音频片段的开始时间，单位为秒。对于点播文件，该参数代表对应音频相对于完整音轨的偏移时间，如0（代表不偏移），5（音轨开始后5秒），10（音轨开始后10秒）；对于直播文件，该参数则返回对应音频片段开始时的Unix时间戳，如：1594650717。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetTime *string `json:"OffsetTime,omitempty" name:"OffsetTime"`

	// 该字段用于返回音频片段的具体审核结果，详细内容敬请参考AudioResult数据结构的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AudioResult `json:"Result,omitempty" name:"Result"`
}

type BucketInfo struct {

	// 该字段用于标识腾讯云对象存储的存储桶名称,关于文件桶的详细信息敬请参考 [腾讯云存储相关说明](https://cloud.tencent.com/document/product/436/44352)。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 该字段用于标识腾讯云对象存储的托管机房的分布地区，对象存储 COS 的数据存放在这些地域的存储桶中。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 该字段用于标识腾讯云对象存储的对象Key,对象z作为基本单元被存放在存储桶中；用户可以通过腾讯云控制台、API、SDK 等多种方式管理对象。有关对象的详细描述敬请参阅相应 [产品文档](https://cloud.tencent.com/document/product/436/13324)。
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAudioModerationSyncTaskRequest struct {
	*tchttp.BaseRequest

	// Biztype为策略的具体的编号，用于接口调度，在内容安全控制台中可配置。不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。Biztype仅为数字、字母与下划线的组合，长度为3-32个字符；调用时不传入Biztype代表采用默认的识别策略。
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 数据标识，可以由英文字母、数字、下划线、-、@#组成，不超过64个字符
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 音频文件资源格式，当前为mp3，wav，请按照实际文件格式填入
	FileFormat *string `json:"FileFormat,omitempty" name:"FileFormat"`

	// 文件名称，可以由英文字母、数字、下划线、-、@#组成，不超过64个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据Base64编码，短音频同步接口仅传入可音频内容；
	// 支持范围：文件大小不能超过5M，时长不可超过60s，码率范围为8-16Kbps；
	// 支持格式：wav、mp3
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 音频资源访问链接，与FileContent参数必须二选一输入；
	// 支持范围：同FileContent；
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

func (r *CreateAudioModerationSyncTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAudioModerationSyncTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizType")
	delete(f, "DataId")
	delete(f, "FileFormat")
	delete(f, "Name")
	delete(f, "FileContent")
	delete(f, "FileUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAudioModerationSyncTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAudioModerationSyncTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求接口时传入的数据标识
		DataId *string `json:"DataId,omitempty" name:"DataId"`

		// 文件名称，可以由英文字母、数字、下划线、-、@#组成，不超过64个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
		Name *string `json:"Name,omitempty" name:"Name"`

		// Biztype为策略的具体的编号，用于接口调度，在内容安全控制台中可配置。不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。Biztype仅为数字、字母与下划线的组合，长度为3-32个字符；调用时不传入Biztype代表采用默认的识别策略。
		BizType *string `json:"BizType,omitempty" name:"BizType"`

		// 智能审核服务对于内容违规类型的等级，可选值：
	// Pass 建议通过；
	// Reveiw 建议复审；
	// Block 建议屏蔽；
	// 注意：此字段可能返回 null，表示取不到有效值。
		Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

		// 智能审核服务对于内容违规类型的判断，详见返回值列表
	// 如：Label：Porn（色情）；
	// 注意：此字段可能返回 null，表示取不到有效值。
		Label *string `json:"Label,omitempty" name:"Label"`

		// 音频文本，备注：这里的文本最大只返回前1000个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
		AsrText *string `json:"AsrText,omitempty" name:"AsrText"`

		// 音频中对话内容审核结果；
	// 注意：此字段可能返回 null，表示取不到有效值。
		TextResults []*TextResult `json:"TextResults,omitempty" name:"TextResults"`

		// 音频中低俗内容审核结果；
	// 注意：此字段可能返回 null，表示取不到有效值。
		MoanResults []*MoanResult `json:"MoanResults,omitempty" name:"MoanResults"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAudioModerationSyncTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAudioModerationSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAudioModerationTaskRequest struct {
	*tchttp.BaseRequest

	// 输入的任务信息，最多可以同时创建10个任务
	Tasks []*TaskInput `json:"Tasks,omitempty" name:"Tasks"`

	// 默认为：default，客户可以在音频审核控制台配置自己的BizType
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 审核类型，这里可选：AUDIO (点播音频)和 LIVE_AUDIO（直播音频），默认为 AUDIIO
	Type *string `json:"Type,omitempty" name:"Type"`

	// （可选）回调签名key，具体可以查看 回调签名示例
	Seed *string `json:"Seed,omitempty" name:"Seed"`

	// 接收审核信息回调地址，如果设置，则审核过程中产生的违规音频片段和画面截帧发送此接口
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

func (r *CreateAudioModerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAudioModerationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tasks")
	delete(f, "BizType")
	delete(f, "Type")
	delete(f, "Seed")
	delete(f, "CallbackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAudioModerationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAudioModerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务创建结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Results []*TaskResult `json:"Results,omitempty" name:"Results"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAudioModerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAudioModerationTaskResponse) FromJsonString(s string) error {
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

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 审核时传入的数据Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		DataId *string `json:"DataId,omitempty" name:"DataId"`

		// 业务类型，用户可以在控制台查看自己配置的BizType
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
		Labels []*TaskLabel `json:"Labels,omitempty" name:"Labels"`

		// 输入的媒体信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		InputInfo *InputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

		// 音频文本，备注：这里的文本最大只返回前1000个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
		AudioText *string `json:"AudioText,omitempty" name:"AudioText"`

		// 音频片段审核信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		AudioSegments []*AudioSegments `json:"AudioSegments,omitempty" name:"AudioSegments"`

		// 错误类型，如果任务状态为Error，则该字段不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorType *string `json:"ErrorType,omitempty" name:"ErrorType"`

		// 错误描述，如果任务状态为Error，则该字段不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorDescription *string `json:"ErrorDescription,omitempty" name:"ErrorDescription"`

		// 任务创建时间，格式为 ISO 8601
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 任务最后更新时间，格式为 ISO 8601
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InputInfo struct {

	// 该字段表示文件访问类型，取值为**URL**（资源链接）和**COS** (腾讯云对象存储)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 该字段表示文件访问的链接地址，格式为标准URL格式。<br> 备注：当Type为URL时此字段不为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 该字段表示文件访问的腾讯云存储桶信息。<br> 备注：当Type为COS时此字段不为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketInfo *BucketInfo `json:"BucketInfo,omitempty" name:"BucketInfo"`
}

type MediaInfo struct {

	// 该字段用于返回传入的媒体文件的编码格式，如wav、mp3、aac、flac、amr、3gp、 m4a、wma、ogg、ape等。
	Codecs *string `json:"Codecs,omitempty" name:"Codecs"`

	// 该字段用于返回对传入的流媒体文件进行分片的片段时长，单位为秒。**默认值为15秒**，支持用户自定义配置。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// *内测中，敬请期待*
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// *内测中，敬请期待*
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// *内测中，敬请期待*
	Thumbnail *string `json:"Thumbnail,omitempty" name:"Thumbnail"`
}

type MoanResult struct {

	// 固定取值为Moan（呻吟/娇喘），如音频中无复杂类型「MoanResult」的返回则代表改音频中无呻吟/娇喘相关违规内容；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 机器判断当前分类的置信度，取值范围：0~100。分数越高，表示越有可能属于当前分类。
	// （如：Moan 99，则该样本属于呻吟/娇喘的置信度非常高。）
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 建议您拿到判断结果后的执行操作。
	// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 违规事件开始时间，单位为毫秒（ms）；
	StartTime *float64 `json:"StartTime,omitempty" name:"StartTime"`

	// 违规事件结束时间，单位为毫秒（ms）；
	EndTime *float64 `json:"EndTime,omitempty" name:"EndTime"`
}

type StorageInfo struct {

	// 该字段表示文件访问类型，取值为**URL**（资源链接）和**COS** (腾讯云对象存储)；该字段应当与传入的访问类型相对应，可用于强校验并方便系统快速识别访问地址；若不传入此参数，则默认值为URL，此时系统将自动判定访问地址类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 该字段表示文件访问的链接地址，格式为标准URL格式。<br> 备注：当Type为URL时此字段不为空，该参数与BucketInfo参数须传入其中之一
	Url *string `json:"Url,omitempty" name:"Url"`

	// 该字段表示文件访问的腾讯云存储桶信息。<br> 备注：当Type为COS时此字段不为空，该参数与Url参数须传入其中之一。
	BucketInfo *BucketInfo `json:"BucketInfo,omitempty" name:"BucketInfo"`
}

type TaskData struct {

	// 该字段用于返回音频审核任务数据所对应的数据ID，方便后续查询和管理审核任务。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 该字段用于返回音频审核任务所生成的任务ID，用于标识具体审核任务，方便后续查询和管理。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 该字段用于返回所查询内容的任务状态。
	// <br>取值：**FINISH**（任务已完成）、**PENDING** （任务等待中）、**RUNNING** （任务进行中）、**ERROR** （任务出错）、**CANCELLED** （任务已取消）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 该字段用于返回音频审核任务所对应的任务名称，方便后续查询和管理审核任务。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 该字段用于返回调用音频审核接口时传入的BizType参数，方便数据的辨别和管理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 该字段用于返回调用音频审核接口时输入的音频审核类型，取值为：**AUDIO**（点播音频）和**LIVE_AUDIO**（直播音频），默认值为AUDIO。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 该字段用于返回基于恶意标签的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 输入信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaInfo *MediaInfo `json:"MediaInfo,omitempty" name:"MediaInfo"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*TaskLabel `json:"Labels,omitempty" name:"Labels"`

	// 该字段用于返回被查询任务创建的时间，格式采用 ISO 8601标准。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 该字段用于返回被查询任务最后更新时间，格式采用 ISO 8601标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type TaskFilter struct {

	// 该字段用于传入任务对应的业务类型供筛选器进行筛选。Biztype为策略的具体的编号，用于接口调度，在内容安全控制台中可配置。不同Biztype关联不同的业务场景与审核策略，调用前请确认正确的Biztype。Biztype仅为**数字、字母与下划线的组合**，长度为3-32个字符。<br>备注：在不传入该参数时筛选器默认不筛选业务类型。
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 该字段用于传入音频审核对应的任务类型供筛选器进行筛选，取值为：**VIDEO**（点播视频审核），**AUDIO**（点播音频审核）， **LIVE_VIDEO**（直播视频审核）, **LIVE_AUDIO**（直播音频审核）。<br>备注：在不传入该参数时筛选器默认不筛选任务类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 该字段用于传入音频审核对应的建议操作供筛选器进行筛选，取值为：**Block**：建议屏蔽，**Review**：建议人工复审，**Pass**：建议通过。<br>备注：在不传入该参数时筛选器默认不筛选建议操作。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 该字段用于传入审核任务的任务状态供筛选器进行筛选，取值为：**FINISH**（任务已完成）、**PENDING** （任务等待中）、**RUNNING** （任务进行中）、**ERROR** （任务出错）、**CANCELLED** （任务已取消）。<br>备注：在不传入该参数时筛选器默认不筛选任务状态。
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type TaskInput struct {

	// 选填参数，该字段表示您为待检测对象分配的数据ID，传入后可方便您对文件进行标识和管理。<br>取值：由英文字母（大小写均可）、数字及四个特殊符号（_，-，@，#）组成，**长度不超过64个字符**。
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 选填参数，该字段表示音频审核任务所对应的任务名称，方便后续查询和管理审核任务。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 必填参数，该字段表示审核文件的访问参数，用于获取审核媒体文件，该参数内包括访问类型和访问地址。
	Input *StorageInfo `json:"Input,omitempty" name:"Input"`
}

type TaskLabel struct {

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 该字段用于返回当前标签对应的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 该字段用于返回当前标签（Label）下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容；*色情 0*，则表明该文本不属于色情内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type TaskResult struct {

	// 该字段用于返回创建音频审核任务时在TaskInput结构内传入的DataId，用于标识具体审核任务。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 该字段用于返回音频审核任务所生成的任务ID，用于标识具体审核任务，方便后续查询和管理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 该字段用于返回任务创建的状态，如返回OK则代表任务创建成功，其他返回值可参考公共错误码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// **仅在Code的返回值为错误码时生效**，用于返回错误的详情内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type TextResult struct {

	// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义词库。
	// 以及其他令人反感、不安全或不适宜的内容类型。
	// 
	// 如音频中无复杂类型「TextResults」的返回则代表改音频中无相关违规内容；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 命中的关键词，为空则代表该违规内容出自于模型的判断；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 命中关键词库的库标识；
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// 命中关键词库的名字；
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// 机器判断当前分类的置信度，取值范围：0~100。分数越高，表示越有可能属于当前分类。
	// （如：Porn 99，则该样本属于色情的置信度非常高。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 建议您拿到判断结果后的执行操作。
	// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 自定义词库的类型，自定义词库相关的信息可登录控制台中查看；
	// 
	// 1：自定义黑白库；
	// 
	// 2：自定义库；
	LibType *int64 `json:"LibType,omitempty" name:"LibType"`
}
