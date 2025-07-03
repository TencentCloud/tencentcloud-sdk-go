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

package v20210922

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AudioResult struct {
	// 该字段用于返回审核内容是否命中审核模型；取值：0（**未命中**）、1（**命中**）。
	HitFlag *int64 `json:"HitFlag,omitnil,omitempty" name:"HitFlag"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示具体的后续建议操作。<br>
	// 返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回当前标签下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回音频文件经ASR识别后的文本信息。最长可识别**5小时**的音频文件，若超出时长限制，接口将会报错。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 该字段用于返回审核结果的访问链接（URL）。<br>备注：链接默认有效期为12小时。如果您需要更长时效的链接，请使用[COS预签名](https://cloud.tencent.com/document/product/1265/104001)功能更新签名时效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 该字段用于返回音频文件的时长，单位为毫秒。
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 该字段用于返回输入参数中的额外附加信息（Extra），如未配置则默认返回值为空。<br>备注：不同客户或Biztype下返回信息不同，如需配置该字段请提交工单咨询或联系售后专员处理。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 该字段用于返回音频文件经ASR识别后产生的文本的详细审核结果。具体结果内容请参见AudioResultDetailLanguageResult数据结构的细节描述。
	TextResults []*AudioResultDetailTextResult `json:"TextResults,omitnil,omitempty" name:"TextResults"`

	// 该字段用于返回音频文件呻吟检测的详细审核结果。具体结果内容请参见AudioResultDetailMoanResult数据结构的细节描述。
	MoanResults []*AudioResultDetailMoanResult `json:"MoanResults,omitnil,omitempty" name:"MoanResults"`

	// 该字段用于返回音频小语种检测的详细审核结果。具体结果内容请参见AudioResultDetailLanguageResult数据结构的细节描述。
	LanguageResults []*AudioResultDetailLanguageResult `json:"LanguageResults,omitnil,omitempty" name:"LanguageResults"`

	// 该字段用于返回当前标签（Lable）下的二级标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 识别类标签结果信息列表
	RecognitionResults []*RecognitionResult `json:"RecognitionResults,omitnil,omitempty" name:"RecognitionResults"`

	// 该字段用于返回音频文件说话人检测的详细审核结果
	SpeakerResults []*SpeakerResult `json:"SpeakerResults,omitnil,omitempty" name:"SpeakerResults"`

	// 该字段用于返回音频文件出行检测的详细审核结果
	TravelResults []*TravelResult `json:"TravelResults,omitnil,omitempty" name:"TravelResults"`

	// 该字段用于返回音频文件的三级标签
	SubTag *string `json:"SubTag,omitnil,omitempty" name:"SubTag"`

	// 该字段用于返回音频文件的三级标签码
	SubTagCode *string `json:"SubTagCode,omitnil,omitempty" name:"SubTagCode"`

	// 该字段用于返回音频文件歌曲识别的详细审核结果
	LabelResults []*LabelResult `json:"LabelResults,omitnil,omitempty" name:"LabelResults"`

	// 审核命中类型
	HitType *string `json:"HitType,omitnil,omitempty" name:"HitType"`
}

type AudioResultDetailLanguageResult struct {
	// 语种
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 得分
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 开始时间
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 子标签码
	SubLabelCode *string `json:"SubLabelCode,omitnil,omitempty" name:"SubLabelCode"`
}

type AudioResultDetailMoanResult struct {
	// 该字段用于返回检测结果需要检测的内容类型，此处固定为**Moan**（呻吟）以调用呻吟检测功能。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回呻吟检测的置信度，取值范围：0（**置信度最低**）-100（**置信度最高**），越高代表音频越有可能属于呻吟内容。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回对应呻吟标签的片段在音频文件内的开始时间，单位为毫秒。
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 该字段用于返回对应呻吟标签的片段在音频文件内的结束时间，单位为毫秒。
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// *内测中，敬请期待*
	SubLabelCode *string `json:"SubLabelCode,omitnil,omitempty" name:"SubLabelCode"`

	// 该字段用于返回当前标签（Lable）下的二级标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回基于恶意标签的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`
}

type AudioResultDetailTextResult struct {
	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回ASR识别出的文本内容命中的关键词信息，用于标注内容违规的具体原因（如：加我微信）。该参数可能会有多个返回值，代表命中的多个关键词；若返回值为空，Score不为空，则代表识别结果所对应的恶意标签（Label）来自于语义模型判断的返回值。
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 该字段**仅当Label为Custom：自定义关键词时该参数有效**,用于返回自定义库的ID，以方便自定义库管理和配置。
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// 该字段**仅当Label为Custom：自定义关键词时该参数有效**,用于返回自定义库的名称,以方便自定义库管理和配置。
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 该字段用于返回当前标签下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高**），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回自定义关键词对应的词库类型，取值为**1**（黑白库）和**2**（自定义关键词库），若未配置自定义关键词库,则默认值为1（黑白库匹配）。
	LibType *int64 `json:"LibType,omitnil,omitempty" name:"LibType"`

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示具体的后续建议操作。<br>
	// 返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回当前标签（Lable）下的二级标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回命中的关键词信息
	HitInfos []*OcrHitInfo `json:"HitInfos,omitnil,omitempty" name:"HitInfos"`
}

type AudioSegments struct {
	// 截帧时间。
	// 点播文件：该值为相对于视频偏移时间，单位为秒，例如：0，5，10
	// 直播流：该值为时间戳，例如：1594650717
	OffsetTime *string `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 结果集
	Result *AudioResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

type BucketInfo struct {
	// 腾讯云对象存储，存储桶名称
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 对象Key
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`
}

// Predefined struct for user
type CancelTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateVideoModerationTaskRequestParams struct {
	// 该字段表示特定审核策略的编号，用于接口调度。需要提前在[内容安全控制台](https://console.cloud.tencent.com/cms/clouds/manage)中创建策略后获取该Biztype字段，传入该字段，会根据业务场景在审核时调用相应的审核策略。 备注：Biztype仅为数字、字母与下划线的组合，长度为3-32个字符；不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 任务类型：可选VIDEO（点播视频），LIVE_VIDEO（直播视频）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 输入的任务信息，最多可以同时创建10个任务
	Tasks []*TaskInput `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 可选参数，该字段表示回调签名的key信息，用于保证数据的安全性。 签名方法为在返回的HTTP头部添加 X-Signature 的字段，值为： seed + body 的 SHA256 编码和Hex字符串，在收到回调数据后，可以根据返回的body，用 **sha256(seed + body)**, 计算出 `X-Signature` 进行验证。<br>具体使用实例可参考 [回调签名示例](https://cloud.tencent.com/document/product/1265/104001#42dd87d2-580f-46cf-a953-639a787d1eda)。
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 接收审核信息回调地址。如果设置了该字段，在审核过程中发现违规音频片段和画面截帧结果将发送至该接口。更多详情请参阅[回调配置说明](https://cloud.tencent.com/document/product/1265/104001)。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 审核排队优先级。当您有多个视频审核任务排队时，可以根据这个参数控制排队优先级。用于处理插队等逻辑。默认该参数为0
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 该字段表示待检测对象对应的用户相关信息，若填入则可甄别相应违规风险用户
	User *User `json:"User,omitnil,omitempty" name:"User"`
}

type CreateVideoModerationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 该字段表示特定审核策略的编号，用于接口调度。需要提前在[内容安全控制台](https://console.cloud.tencent.com/cms/clouds/manage)中创建策略后获取该Biztype字段，传入该字段，会根据业务场景在审核时调用相应的审核策略。 备注：Biztype仅为数字、字母与下划线的组合，长度为3-32个字符；不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 任务类型：可选VIDEO（点播视频），LIVE_VIDEO（直播视频）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 输入的任务信息，最多可以同时创建10个任务
	Tasks []*TaskInput `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 可选参数，该字段表示回调签名的key信息，用于保证数据的安全性。 签名方法为在返回的HTTP头部添加 X-Signature 的字段，值为： seed + body 的 SHA256 编码和Hex字符串，在收到回调数据后，可以根据返回的body，用 **sha256(seed + body)**, 计算出 `X-Signature` 进行验证。<br>具体使用实例可参考 [回调签名示例](https://cloud.tencent.com/document/product/1265/104001#42dd87d2-580f-46cf-a953-639a787d1eda)。
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 接收审核信息回调地址。如果设置了该字段，在审核过程中发现违规音频片段和画面截帧结果将发送至该接口。更多详情请参阅[回调配置说明](https://cloud.tencent.com/document/product/1265/104001)。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 审核排队优先级。当您有多个视频审核任务排队时，可以根据这个参数控制排队优先级。用于处理插队等逻辑。默认该参数为0
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 该字段表示待检测对象对应的用户相关信息，若填入则可甄别相应违规风险用户
	User *User `json:"User,omitnil,omitempty" name:"User"`
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
	delete(f, "User")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoModerationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoModerationTaskResponseParams struct {
	// 任务创建结果
	Results []*TaskResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 是否展示所有分片，默认只展示命中规则的分片
	ShowAllSegments *bool `json:"ShowAllSegments,omitnil,omitempty" name:"ShowAllSegments"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，创建任务后返回的TaskId字段
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 是否展示所有分片，默认只展示命中规则的分片
	ShowAllSegments *bool `json:"ShowAllSegments,omitnil,omitempty" name:"ShowAllSegments"`
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
	// 该字段用于返回创建视频审核任务后返回的任务ID（在Results参数中），用于标识需要查询任务详情的审核任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 该字段用于返回调用视频审核接口时传入的数据ID参数，方便数据的辨别和管理。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段用于返回调用视频审核接口时传入的BizType参数，方便数据的辨别和管理。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段用于返回调用视频审核接口时传入的TaskInput参数中的任务名称，方便任务的识别与管理。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 该字段用于返回所查询内容的任务状态。
	// <br>取值：**FINISH**（任务已完成）、**PENDING** （任务等待中）、**RUNNING** （任务进行中）、**ERROR** （任务出错）、**CANCELLED** （任务已取消）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 该字段用于返回调用视频审核接口时输入的视频审核类型，取值为：**VIDEO**（点播视频）和**LIVE_VIDEO**（直播视频），默认值为VIDEO。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 该字段用于返回基于恶意标签的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	Labels []*TaskLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 该字段用于返回输入媒体文件的详细信息，包括编解码格式、分片时长等信息。详细内容敬请参考MediaInfo数据结构的描述。
	MediaInfo *MediaInfo `json:"MediaInfo,omitnil,omitempty" name:"MediaInfo"`

	// 该字段用于返回审核服务的媒体内容信息，主要包括传入文件类型和访问地址。
	InputInfo *InputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 该字段用于返回被查询任务创建的时间，格式采用 ISO 8601标准。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 该字段用于返回被查询任务最后更新时间，格式采用 ISO 8601标准。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 在秒后重试
	TryInSeconds *int64 `json:"TryInSeconds,omitnil,omitempty" name:"TryInSeconds"`

	// 该字段用于返回视频中截帧审核的结果，详细返回内容敬请参考ImageSegments数据结构的描述。
	ImageSegments []*ImageSegments `json:"ImageSegments,omitnil,omitempty" name:"ImageSegments"`

	// 该字段用于返回视频中音频审核的结果，详细返回内容敬请参考AudioSegments数据结构的描述。
	AudioSegments []*AudioSegments `json:"AudioSegments,omitnil,omitempty" name:"AudioSegments"`

	// 当任务状态为Error时，返回对应错误的类型，取值：
	// **DECODE_ERROR**: 解码失败。（输入资源中可能包含无法解码的视频）
	// **URL_ERROR**：下载地址验证失败。
	// **TIMEOUT_ERROR**：处理超时。
	// **CALLBACK_ERRORR**：回调错误。
	// **MODERATION_ERROR**：审核失败。
	// **URL_NOT_SUPPORTED**：源文件太大或没有图片音频帧
	// 任务状态非Error时默认返回为空。
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// 当任务状态为Error时，该字段用于返回对应错误的详细描述，任务状态非Error时默认返回为空。
	ErrorDescription *string `json:"ErrorDescription,omitnil,omitempty" name:"ErrorDescription"`

	// 该字段用于返回检测结果所对应的标签。如果未命中恶意，返回Normal，如果命中恶意，则返回Labels中优先级最高的标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回音频文件识别出的对应文本内容，最大支持**前1000个字符**。
	AudioText *string `json:"AudioText,omitnil,omitempty" name:"AudioText"`

	// 该字段用于返回音频文件识别出的对应文本内容。
	Asrs []*RcbAsr `json:"Asrs,omitnil,omitempty" name:"Asrs"`

	// 该字段用于返回检测结果明细数据相关的cos url	
	SegmentCosUrlList *SegmentCosUrlList `json:"SegmentCosUrlList,omitnil,omitempty" name:"SegmentCosUrlList"`

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
	// 该参数表示任务列表每页展示的任务条数，**默认值为10，最大值为100**（每页展示10条任务）。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 该参数表示任务筛选器的输入参数，可根据业务类型、审核文件类型、处理建议及任务状态筛选想要查看的审核任务，具体参数内容请参见TaskFilter数据结构的详细描述。
	Filter *TaskFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 该参数表示翻页时使用的Token信息，由系统自动生成，并在翻页时向下一个生成的页面传递此参数，以方便快速翻页功能的实现。当到最后一页时，该字段为空。
	PageToken *string `json:"PageToken,omitnil,omitempty" name:"PageToken"`

	// 该参数表示任务列表的开始时间，格式为ISO8601标准的时间戳。**默认值为最近3天**，若传入该参数，则在这一时间到EndTime之间的任务将会被筛选出来。<br>备注：该参数与Filter共同起到任务筛选作用，二者作用无先后顺序。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 该参数表示任务列表的结束时间，格式为ISO8601标准的时间戳。**默认值为空**，若传入该参数，则在这StartTime到这一时间之间的任务将会被筛选出来。<br>备注：该参数与Filter共同起到任务筛选作用，二者作用无先后顺序。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 该参数表示任务列表每页展示的任务条数，**默认值为10，最大值为100**（每页展示10条任务）。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 该参数表示任务筛选器的输入参数，可根据业务类型、审核文件类型、处理建议及任务状态筛选想要查看的审核任务，具体参数内容请参见TaskFilter数据结构的详细描述。
	Filter *TaskFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 该参数表示翻页时使用的Token信息，由系统自动生成，并在翻页时向下一个生成的页面传递此参数，以方便快速翻页功能的实现。当到最后一页时，该字段为空。
	PageToken *string `json:"PageToken,omitnil,omitempty" name:"PageToken"`

	// 该参数表示任务列表的开始时间，格式为ISO8601标准的时间戳。**默认值为最近3天**，若传入该参数，则在这一时间到EndTime之间的任务将会被筛选出来。<br>备注：该参数与Filter共同起到任务筛选作用，二者作用无先后顺序。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 该参数表示任务列表的结束时间，格式为ISO8601标准的时间戳。**默认值为空**，若传入该参数，则在这StartTime到这一时间之间的任务将会被筛选出来。<br>备注：该参数与Filter共同起到任务筛选作用，二者作用无先后顺序。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 该字段用于返回当前查询的任务总量，格式为int字符串。
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 该字段用于返回当前页的任务详细数据，具体输出内容请参见TaskData数据结构的详细描述。
	Data []*TaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// 该字段用于返回翻页时使用的Token信息，由系统自动生成，并在翻页时向下一个生成的页面传递此参数，以方便快速翻页功能的实现。当到最后一页时，该字段为空。
	PageToken *string `json:"PageToken,omitnil,omitempty" name:"PageToken"`

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

type ImageResult struct {
	// 违规标志
	// 0 未命中
	// 1 命中
	HitFlag *int64 `json:"HitFlag,omitnil,omitempty" name:"HitFlag"`

	// 命中的标签
	// Porn 色情
	// Sexy 性感
	// Polity 政治
	// Illegal 违法
	// Abuse 谩骂
	// Terror 暴恐
	// Ad 广告
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 得分
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 画面截帧图片结果集
	Results []*ImageResultResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 该字段用于返回审核结果的访问链接（URL）。<br>备注：链接默认有效期为12小时。如果您需要更长时效的链接，请使用[COS预签名](https://cloud.tencent.com/document/product/1265/104001)功能更新签名时效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 附加字段
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 二级标签
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 场景结果
	RecognitionResults []*RecognitionResult `json:"RecognitionResults,omitnil,omitempty" name:"RecognitionResults"`

	// 审核命中类型
	HitType *string `json:"HitType,omitnil,omitempty" name:"HitType"`
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
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 是否命中
	// 0 未命中
	// 1 命中
	HitFlag *int64 `json:"HitFlag,omitnil,omitempty" name:"HitFlag"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 子标签
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 分数
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 人物名称列表，如未识别，则为null
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 图片OCR文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 其他详情
	Details []*ImageResultsResultDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 审核命中类型
	HitType *string `json:"HitType,omitnil,omitempty" name:"HitType"`
}

type ImageResultsResultDetail struct {
	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// OCR识别文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 位置信息
	Location *ImageResultsResultDetailLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 库ID
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// 库名称
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 命中的关键词
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 得分
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 子标签码
	SubLabelCode *string `json:"SubLabelCode,omitnil,omitempty" name:"SubLabelCode"`

	// 子标签
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// Ocr文本命中信息
	OcrHitInfos []*OcrHitInfo `json:"OcrHitInfos,omitnil,omitempty" name:"OcrHitInfos"`
}

type ImageResultsResultDetailLocation struct {
	// x坐标
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// y坐标
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 旋转角度
	Rotate *float64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`
}

type ImageSegments struct {
	// 截帧时间。
	// 点播文件：该值为相对于视频偏移时间，单位为秒，例如：0，5，10
	// 直播流：该值为时间戳，例如：1594650717
	OffsetTime *string `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 画面截帧结果详情
	Result *ImageResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 截帧毫秒时间
	OffsetusTime *string `json:"OffsetusTime,omitnil,omitempty" name:"OffsetusTime"`
}

type InputInfo struct {
	// 传入的类型可选：URL，COS
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Url地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 桶信息。当输入当时COS时，该字段不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketInfo *string `json:"BucketInfo,omitnil,omitempty" name:"BucketInfo"`
}

type LabelResult struct {
	// 场景
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 建议
	Suggestion *int64 `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分数
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 开始时间
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type MediaInfo struct {
	// 编码格式
	Codecs *string `json:"Codecs,omitnil,omitempty" name:"Codecs"`

	// 流检测时分片时长
	// 注意：此字段可能返回 0，表示取不到有效值。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 宽，单位为像素
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高，单位为像素
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 封面
	Thumbnail *string `json:"Thumbnail,omitnil,omitempty" name:"Thumbnail"`
}

type OcrHitInfo struct {
	// 关键词
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 关键词内容
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 自定义库名
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 位置信息
	Positions []*TextPosition `json:"Positions,omitnil,omitempty" name:"Positions"`
}

type RcbAsr struct {
	// 该字段用于返回音频文件识别出的对应文本内容，最大支持**前1000个字符**。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 该字段用于返回被查询任务创建的时间，格式采用 ISO 8601标准。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

type RecognitionResult struct {
	// 可能的取值有：Teenager 、Gender
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 识别标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SegmentCosUrlList struct {
	// 全量图片片段的cos url
	ImageAllUrl *string `json:"ImageAllUrl,omitnil,omitempty" name:"ImageAllUrl"`

	// 全量音频片段的cos url
	AudioAllUrl *string `json:"AudioAllUrl,omitnil,omitempty" name:"AudioAllUrl"`

	// 违规图片片段的cos url
	ImageBlockUrl *string `json:"ImageBlockUrl,omitnil,omitempty" name:"ImageBlockUrl"`

	// 违规音频片段的cos url
	AudioBlockUrl *string `json:"AudioBlockUrl,omitnil,omitempty" name:"AudioBlockUrl"`

	// 全量音频识别文本的cos url
	AsrUrl *string `json:"AsrUrl,omitnil,omitempty" name:"AsrUrl"`
}

type SpeakerResult struct {
	// 标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 分数
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 开始时间
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type StorageInfo struct {
	// 类型 可选：
	// URL 资源链接类型
	// COS 腾讯云对象存储类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 资源链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 腾讯云存储桶信息
	BucketInfo *BucketInfo `json:"BucketInfo,omitnil,omitempty" name:"BucketInfo"`
}

type Tag struct {
	// 根据Label字段确定具体名称：
	// 当Label 为Teenager 时 Name可能取值有：Teenager 
	// 当Label 为Gender 时 Name可能取值有：Male 、Female
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 置信分：0～100，数值越大表示置信度越高
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 识别开始偏移时间，单位：毫秒
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 识别结束偏移时间，单位：毫秒
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type TaskData struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 输入的数据ID
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 业务类型
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态，可选：PENDING，RUNNING，ERROR，FINISH，CANCELLED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 处理建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 标签
	Labels []*TaskLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 媒体信息
	MediaInfo *MediaInfo `json:"MediaInfo,omitnil,omitempty" name:"MediaInfo"`

	// 输入信息
	InputInfo *InputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type TaskFilter struct {
	// 任务业务类型
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 任务类型，可选：VIDEO，AUDIO， LIVE_VIDEO, LIVE_AUDIO
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 建议，可选：Pass, Review,Block
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 状态，可选：PENDING，RUNNING，ERROR，FINISH，CANCELLED
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`
}

type TaskInput struct {
	// 数据ID
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 任务名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务输入
	Input *StorageInfo `json:"Input,omitnil,omitempty" name:"Input"`
}

type TaskLabel struct {
	// 该字段用于返回检测结果所对应的恶意标签。
	// 返回值：Porn：色情，Abuse：谩骂，Ad：广告；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 审核建议，可选值：
	// Pass 通过，
	// Review 建议人审，
	// Block 确认违规
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 得分，分数是 0 ～ 100
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 命中的二级标签
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`
}

type TaskResult struct {
	// 请求时传入的DataId
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// TaskId，任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 错误码。如果code为OK，则表示创建成功，其他则参考公共错误码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 如果错误，该字段表示错误详情
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type TextPosition struct {
	// 命中关键词在文本中的起始位置
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 命中关键词在文本中的结束位置
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

type TravelResult struct {
	// 一级标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 二级标签
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 风险等级
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 出行音频角色
	AudioRole *string `json:"AudioRole,omitnil,omitempty" name:"AudioRole"`

	// 出行语音文本
	AudioText *string `json:"AudioText,omitnil,omitempty" name:"AudioText"`

	// 开始时间
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type User struct {
	// 业务用户ID 如填写，会根据账号历史恶意情况，判定消息有害结果，特别是有利于可疑恶意情况下的辅助判断。账号可以填写微信uin、QQ号、微信openid、QQopenid、字符串等。该字段和账号类别确定唯一账号。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 该字段表示业务用户ID对应的账号类型，取值：1-微信uin，2-QQ号，3-微信群uin，4-qq群号，5-微信openid，6-QQopenid，7-其它string。
	// 该字段与账号ID参数（UserId）配合使用可确定唯一账号。
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 性别 默认0 未知 1 男性 2 女性
	Gender *uint64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 年龄 默认0 未知
	Age *uint64 `json:"Age,omitnil,omitempty" name:"Age"`

	// 用户等级，默认0 未知 1 低 2 中 3 高
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 用户简介，长度不超过5000字
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 用户头像图片链接
	HeadUrl *string `json:"HeadUrl,omitnil,omitempty" name:"HeadUrl"`

	// 群聊场景房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 群聊场景群ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 群聊场景群用户数
	GroupSize *int64 `json:"GroupSize,omitnil,omitempty" name:"GroupSize"`

	// 消息接收者ID
	ReceiverId *string `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 消息生成时间，毫秒
	SendTime *string `json:"SendTime,omitnil,omitempty" name:"SendTime"`
}