// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AudioLabelResult struct {
	// 场景
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 建议
	Suggestion *int64 `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 歌曲名等
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分数
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 开始时间
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

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

	// 该字段用于返回音频文件说话人检测的详细审核结果。
	SpeakerResults []*SpeakerResult `json:"SpeakerResults,omitnil,omitempty" name:"SpeakerResults"`

	// 该字段用于返回音频文件歌曲检测的详细审核结果。
	LabelResults []*AudioLabelResult `json:"LabelResults,omitnil,omitempty" name:"LabelResults"`

	// 该字段用于返回音频文件出行检测的详细审核结果。
	TravelResults []*TravelResult `json:"TravelResults,omitnil,omitempty" name:"TravelResults"`

	// 三级标签
	SubTag *string `json:"SubTag,omitnil,omitempty" name:"SubTag"`

	// 三级标签码
	SubTagCode *string `json:"SubTagCode,omitnil,omitempty" name:"SubTagCode"`
}

type AudioResultDetailLanguageResult struct {
	// 该字段用于返回对应的语言种类信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该参数用于返回当前标签下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高**），越高代表音频越有可能属于当前返回的语种标签；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该参数用于返回对应语种标签的片段在音频文件内的开始时间，单位为毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 该参数用于返回对应语种标签的片段在音频文件内的结束时间，单位为毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// *内测中，敬请期待*
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabelCode *string `json:"SubLabelCode,omitnil,omitempty" name:"SubLabelCode"`
}

type AudioResultDetailMoanResult struct {
	// 该字段用于返回检测结果需要检测的内容类型，此处固定为**Moan**（呻吟）以调用呻吟检测功能。
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
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

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示具体的后续建议操作。<br>
	// 返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回自定义关键词对应的词库类型，取值为**1**（黑白库）和**2**（自定义关键词库），若未配置自定义关键词库,则默认值为1（黑白库匹配）。
	LibType *int64 `json:"LibType,omitnil,omitempty" name:"LibType"`

	// 该字段用于返回当前标签（Lable）下的二级标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回命中的关键词信息
	HitInfos []*HitInfo `json:"HitInfos,omitnil,omitempty" name:"HitInfos"`
}

type AudioSegments struct {
	// 该字段用于返回音频片段的开始时间，单位为秒。对于点播文件，该参数代表对应音频相对于完整音轨的偏移时间，如0（代表不偏移），5（音轨开始后5秒），10（音轨开始后10秒）；对于直播文件，该参数则返回对应音频片段开始时的Unix时间戳，如：1594650717。
	OffsetTime *string `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 该字段用于返回音频片段的具体审核结果，详细内容敬请参考AudioResult数据结构的描述。
	Result *AudioResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

type BucketInfo struct {
	// 该字段用于标识腾讯云对象存储的存储桶名称,关于文件桶的详细信息敬请参考 [腾讯云存储相关说明](https://cloud.tencent.com/document/product/436/44352)。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 该字段用于标识腾讯云对象存储的托管机房的分布地区，对象存储 COS 的数据存放在这些地域的存储桶中。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 该字段用于标识腾讯云对象存储的对象Key,对象z作为基本单元被存放在存储桶中；用户可以通过腾讯云控制台、API、SDK 等多种方式管理对象。有关对象的详细描述敬请参阅相应 [产品文档](https://cloud.tencent.com/document/product/436/13324)。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`
}

// Predefined struct for user
type CancelTaskRequestParams struct {
	// 该字段表示创建视频审核任务后返回的任务ID（在Results参数中），用于标识需要取消的审核任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest
	
	// 该字段表示创建视频审核任务后返回的任务ID（在Results参数中），用于标识需要取消的审核任务。
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
	// 该参数用于传入审核任务的任务类型，取值：**VIDEO**（点播视频），**LIVE_VIDEO**（直播视频）。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 该字段表示输入的视频审核任务信息，具体输入内容请参见TaskInput数据结构的详细描述。<br> 备注：最多同时可创建**10个任务**。
	Tasks []*TaskInput `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 该字段表示策略的具体编号，用于接口调度，在内容安全控制台中可配置。若不传入Biztype参数（留空），则代表采用biztype为default的策略，没有则需要新建；传入则会在审核时根据业务场景取相应的审核策略。<br>备注：Biztype仅为数字、字母与下划线的组合，长度为3-32个字符；不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 可选参数，该字段表示回调签名的key信息，用于保证数据的安全性。 签名方法为在返回的HTTP头部添加 X-Signature 的字段，值为： seed + body 的 SHA256 编码和Hex字符串，在收到回调数据后，可以根据返回的body，用 **sha256(seed + body)**, 计算出 `X-Signature` 进行验证。<br>具体使用实例可参考 [回调签名示例](https://cloud.tencent.com/document/product/1265/51885)。
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 可选参数，该字段表示接受审核信息回调的地址，格式为URL链接默认格式。配置成功后，审核过程中产生的违规音视频片段将通过此接口发送。回调返回内容格式请参考 [回调签名示例](https://cloud.tencent.com/document/product/1265/51879#.E7.A4.BA.E4.BE.8B2-.E5.9B.9E.E8.B0.83.E7.AD.BE.E5.90.8D.E7.A4.BA.E4.BE.8B) <br>备注：音频默认截取时长为**15秒**，视频截帧默认为**5秒**截取一张图片；若用户自行配置截取间隔，则按照用户配置返回相应片段。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 可选参数，该参数用于传入审核任务的优先级。当您有多个视频审核任务排队时，可以根据这个参数控制排队优先级，用于处理插队等逻辑；该参数**默认值为0**。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type CreateVideoModerationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 该参数用于传入审核任务的任务类型，取值：**VIDEO**（点播视频），**LIVE_VIDEO**（直播视频）。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 该字段表示输入的视频审核任务信息，具体输入内容请参见TaskInput数据结构的详细描述。<br> 备注：最多同时可创建**10个任务**。
	Tasks []*TaskInput `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 该字段表示策略的具体编号，用于接口调度，在内容安全控制台中可配置。若不传入Biztype参数（留空），则代表采用biztype为default的策略，没有则需要新建；传入则会在审核时根据业务场景取相应的审核策略。<br>备注：Biztype仅为数字、字母与下划线的组合，长度为3-32个字符；不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 可选参数，该字段表示回调签名的key信息，用于保证数据的安全性。 签名方法为在返回的HTTP头部添加 X-Signature 的字段，值为： seed + body 的 SHA256 编码和Hex字符串，在收到回调数据后，可以根据返回的body，用 **sha256(seed + body)**, 计算出 `X-Signature` 进行验证。<br>具体使用实例可参考 [回调签名示例](https://cloud.tencent.com/document/product/1265/51885)。
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 可选参数，该字段表示接受审核信息回调的地址，格式为URL链接默认格式。配置成功后，审核过程中产生的违规音视频片段将通过此接口发送。回调返回内容格式请参考 [回调签名示例](https://cloud.tencent.com/document/product/1265/51879#.E7.A4.BA.E4.BE.8B2-.E5.9B.9E.E8.B0.83.E7.AD.BE.E5.90.8D.E7.A4.BA.E4.BE.8B) <br>备注：音频默认截取时长为**15秒**，视频截帧默认为**5秒**截取一张图片；若用户自行配置截取间隔，则按照用户配置返回相应片段。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 可选参数，该参数用于传入审核任务的优先级。当您有多个视频审核任务排队时，可以根据这个参数控制排队优先级，用于处理插队等逻辑；该参数**默认值为0**。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
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

// Predefined struct for user
type CreateVideoModerationTaskResponseParams struct {
	// 该字段用于返回任务创建的结果，具体输出内容请参见TaskResult数据结构的详细描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 该字段表示创建视频审核任务后返回的任务ID（在Results参数中），用于标识需要查询任务详情的审核任务。
	// <br>备注：查询接口单次最大查询量为**20条每次**。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 该布尔字段表示是否展示全部的视频片段，取值：True(展示全部的视频分片)、False(只展示命中审核规则的视频分片)；默认值为False。
	ShowAllSegments *bool `json:"ShowAllSegments,omitnil,omitempty" name:"ShowAllSegments"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 该字段表示创建视频审核任务后返回的任务ID（在Results参数中），用于标识需要查询任务详情的审核任务。
	// <br>备注：查询接口单次最大查询量为**20条每次**。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 该布尔字段表示是否展示全部的视频片段，取值：True(展示全部的视频分片)、False(只展示命中审核规则的视频分片)；默认值为False。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 该字段用于返回调用视频审核接口时传入的数据ID参数，方便数据的辨别和管理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段用于返回调用视频审核接口时传入的BizType参数，方便数据的辨别和管理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段用于返回调用视频审核接口时传入的TaskInput参数中的任务名称，方便任务的识别与管理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 该字段用于返回所查询内容的任务状态。
	// <br>取值：**FINISH**（任务已完成）、**PENDING** （任务等待中）、**RUNNING** （任务进行中）、**ERROR** （任务出错）、**CANCELLED** （任务已取消）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 该字段用于返回调用视频审核接口时输入的视频审核类型，取值为：**VIDEO**（点播视频）和**LIVE_VIDEO**（直播视频），默认值为VIDEO。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 该字段用于返回基于恶意标签的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*TaskLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 该字段用于返回输入媒体文件的详细信息，包括编解码格式、分片时长等信息。详细内容敬请参考MediaInfo数据结构的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaInfo *MediaInfo `json:"MediaInfo,omitnil,omitempty" name:"MediaInfo"`

	// 该字段用于返回审核服务的媒体内容信息，主要包括传入文件类型和访问地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputInfo *InputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 该字段用于返回被查询任务创建的时间，格式采用 ISO 8601标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 该字段用于返回被查询任务最后更新时间，格式采用 ISO 8601标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 该字段用于返回视频中截帧审核的结果，详细返回内容敬请参考ImageSegments数据结构的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSegments []*ImageSegments `json:"ImageSegments,omitnil,omitempty" name:"ImageSegments"`

	// 该字段用于返回视频中音频审核的结果，详细返回内容敬请参考AudioSegments数据结构的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioSegments []*AudioSegments `json:"AudioSegments,omitnil,omitempty" name:"AudioSegments"`

	// 当任务状态为Error时，返回对应错误的类型，取值：**DECODE_ERROR**: 解码失败。（输入资源中可能包含无法解码的视频）
	// **URL_ERROR**：下载地址验证失败。
	// **TIMEOUT_ERROR**：处理超时。任务状态非Error时默认返回为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// 当任务状态为Error时，该字段用于返回对应错误的详细描述，任务状态非Error时默认返回为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDescription *string `json:"ErrorDescription,omitnil,omitempty" name:"ErrorDescription"`

	// 该字段用于返回检测结果所对应的标签。如果未命中恶意，返回Normal，如果命中恶意，则返回Labels中优先级最高的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回检测结果明细数据相关的cos url
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentCosUrlList *SegmentCosUrlList `json:"SegmentCosUrlList,omitnil,omitempty" name:"SegmentCosUrlList"`

	// 该字段用于返回音频审核的ASR识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioText *string `json:"AudioText,omitnil,omitempty" name:"AudioText"`

	// 在秒后重试
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryInSeconds *int64 `json:"TryInSeconds,omitnil,omitempty" name:"TryInSeconds"`

	// 该字段用于返回音频文件识别出的对应文本内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Asrs []*RcbAsr `json:"Asrs,omitnil,omitempty" name:"Asrs"`

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
	// 该参数表示任务列表每页展示的任务条数，**默认值为10**（每页展示10条任务）。
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
	
	// 该参数表示任务列表每页展示的任务条数，**默认值为10**（每页展示10条任务）。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 该字段用于返回当前页的任务详细数据，具体输出内容请参见TaskData数据结构的详细描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// 该字段用于返回翻页时使用的Token信息，由系统自动生成，并在翻页时向下一个生成的页面传递此参数，以方便快速翻页功能的实现。当到最后一页时，该字段为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
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

type HitInfo struct {
	// 关键词
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 命中关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 命中的自定义词库名
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 关键词位置信息
	Positions []*TextPosition `json:"Positions,omitnil,omitempty" name:"Positions"`
}

type ImageResult struct {
	// 该参数用于标识审核内容是否命中恶意标签，取值：0（**未命中**）和1（**命中**）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitnil,omitempty" name:"HitFlag"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示具体的后续建议操作。<br>
	// 返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回当前标签下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 -性行为 99*，则表明该文本非常有可能属于色情性行为内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回图像审核结果的子结果，详细内容敬请参考ImageResultResult数据结构的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*ImageResultResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 该字段用于返回审核结果的访问链接（URL）。<br>备注：链接默认有效期为12小时。如果您需要更长时效的链接，请使用[COS预签名](https://cloud.tencent.com/document/product/1265/104001)功能更新签名时效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 该字段用于返回输入参数中的额外附加信息（Extra），如未配置则默认返回值为空。<br>备注：不同客户或Biztype下返回信息不同，如需配置该字段请提交工单咨询或联系售后专员处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 该字段用于返回当前标签（Lable）下的二级标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回仅识别图片元素的模型结果；包括：场景模型命中的标签、置信度和位置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecognitionResults []*RecognitionResult `json:"RecognitionResults,omitnil,omitempty" name:"RecognitionResults"`
}

type ImageResultResult struct {
	// 该字段用于返回检测结果所对应的恶意场景。返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**AppLogo**：广告台标，**Custom**：自定义违规，以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 该参数用于标识审核内容是否命中恶意标签，取值：0（**未命中**）和1（**命中**）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitnil,omitempty" name:"HitFlag"`

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示具体的后续建议操作。<br>
	// 返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回恶意标签下对应的子标签的检测结果，如：*Porn-SexBehavior*等子标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回当前标签下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 -性行为 99*，则表明该文本非常有可能属于色情性行为内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回审核图片在敏感场景下命中的特定对象名称列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 该字段用于返回图片OCR文本识别的检测结果，识别**上限在5000字节内**。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 该字段用于返回图像审核子结果的其他详细信息，如文本位置、自定义库等。详细返回内容敬请参考ImageResultsResultDetail数据结构的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*ImageResultsResultDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type ImageResultsResultDetail struct {
	// 该字段用于返回调用视频审核接口时传入的TaskInput参数中的任务名称，方便任务的识别与管理。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 该字段用于返回图片OCR文本识别的检测结果，识别**上限在5000字节内**。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 该字段用于返回图像审核子结果的详细位置信息，如坐标、大小、旋转角度等。详细返回内容敬请参考ImageResultsResultDetailLocation数据结构的描述。
	Location *ImageResultsResultDetailLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段**仅当Label为Custom：自定义关键词时该参数有效**,用于返回自定义库的ID，以方便自定义库管理和配置。
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// 该字段**仅当Label为Custom：自定义关键词时该参数有效**,用于返回自定义库的名称,以方便自定义库管理和配置。
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 该字段用于返回检测文本命中的关键词信息，用于标注文本违规的具体原因（如：*加我微信*）。该参数可能会有多个返回值，代表命中的多个关键词；如返回值为空且Score不为空，则代表识别结果所对应的恶意标签（Label）是来自于语义模型判断的返回值。
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示具体的后续建议操作。<br>
	// 返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回当前标签下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回恶意标签下对应的子标签的检测结果，如：*Porn-SexBehavior*等子标签。
	SubLabelCode *string `json:"SubLabelCode,omitnil,omitempty" name:"SubLabelCode"`

	// 该字段用于返回恶意标签下对应的子标签的检测结果，如：*Porn-SexBehavior*等子标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回OCR命中的关键词信息。
	OcrHitInfos []*HitInfo `json:"OcrHitInfos,omitnil,omitempty" name:"OcrHitInfos"`
}

type ImageResultsResultDetailLocation struct {
	// 该参数用于标识OCR检测框左上角位置的**横坐标**（x）所在的像素位置，结合剩余参数可唯一确定检测框的大小和位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// 该参数用于标识OCR检测框左上角位置的**纵坐标**（y）所在的像素位置，结合剩余参数可唯一确定检测框的大小和位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 该参数用于标识OCR检测框的宽度（**由左上角出发在x轴向右延伸的长度**）。结合剩余参数可唯一确定检测框的大小和位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 该参数用于标识OCR检测框的高度（**由左上角出发在y轴向下延伸的长度**）。结合剩余参数可唯一确定检测框的大小和位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 该参数用于标识OCR检测框的旋转角度，该参数结合X和Y两个坐标参数可唯一确定检测框的具体位置；取值：0-360（**角度制**），方向为**逆时针旋**转。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rotate *float64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`
}

type ImageSegments struct {
	// 该字段用于返回视频片段的截帧时间，单位为秒。对于点播文件，该参数代表对应截取图片相对于视频的偏移时间，如0（代表不偏移），5（视频开始后5秒），10（视频开始后10秒）；对于直播文件，该参数则返回对应图片的Unix时间戳，如：1594650717。
	OffsetTime *string `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 该字段用于返回视频片段的具体截帧审核结果，详细内容敬请参考ImageResult数据结构的描述。
	Result *ImageResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 该字段用于返回视频片段的具体截帧审核时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 该字段用于返回视频片段的截帧时间，单位为豪秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetusTime *string `json:"OffsetusTime,omitnil,omitempty" name:"OffsetusTime"`
}

type InputInfo struct {
	// 该字段表示文件访问类型，取值为**URL**（资源链接）和**COS** (腾讯云对象存储)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 该字段表示文件访问的链接地址，格式为标准URL格式。<br> 备注：当Type为URL时此字段不为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 该字段表示文件访问的腾讯云存储桶信息。<br> 备注：当Type为COS时此字段不为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketInfo *string `json:"BucketInfo,omitnil,omitempty" name:"BucketInfo"`
}

type MediaInfo struct {
	// 该字段用于返回对传入的视频流进行分片的片段时长，单位为秒。**默认值为5秒**，支持用户自定义配置。<br>备注：仅在审核文件为流媒体时生效；此字段返回0则代表未取到有效值。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type RcbAsr struct {
	// 该字段用于返回音频文件识别出的对应文本内容，最大支持前1000个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 该字段用于返回被查询任务创建的时间，格式采用 ISO 8601标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

type RecognitionResult struct {
	// 可能的取值有：Teenager 、Gender
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 识别标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SegmentCosUrlList struct {
	// 全量图片片段的cos url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageAllUrl *string `json:"ImageAllUrl,omitnil,omitempty" name:"ImageAllUrl"`

	// 全量音频片段的cos url
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioAllUrl *string `json:"AudioAllUrl,omitnil,omitempty" name:"AudioAllUrl"`

	// 违规图片片段的cos url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageBlockUrl *string `json:"ImageBlockUrl,omitnil,omitempty" name:"ImageBlockUrl"`

	// 违规音频片段的cos url
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioBlockUrl *string `json:"AudioBlockUrl,omitnil,omitempty" name:"AudioBlockUrl"`

	// 全量音频识别文本的cos url
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 该字段表示文件访问类型，取值为**URL**（资源链接）和**COS** (腾讯云对象存储)；该字段应当与传入的访问类型相对应，可用于强校验并方便系统快速识别访问地址；若不传入此参数，则默认值为URL，此时系统将自动判定访问地址类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 该字段表示文件访问的链接地址，格式为标准URL格式。<br> 备注：当Type为URL时此字段不为空，该参数与BucketInfo参数须传入其中之一
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 该字段表示文件访问的腾讯云存储桶信息。<br> 备注：当Type为COS时此字段不为空，该参数与Url参数须传入其中之一。
	BucketInfo *BucketInfo `json:"BucketInfo,omitnil,omitempty" name:"BucketInfo"`
}

type Tag struct {
	// 根据Label字段确定具体名称：
	// 当Label 为Teenager 时 Name可能取值有：Teenager 
	// 当Label 为Gender 时 Name可能取值有：Male 、Female
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 置信分：0～100，数值越大表示置信度越高
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 识别开始偏移时间，单位：毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 识别结束偏移时间，单位：毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type TaskData struct {
	// 该字段用于返回视频审核任务数据所对应的数据ID，方便后续查询和管理审核任务。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段用于返回视频审核任务所生成的任务ID，用于标识具体审核任务，方便后续查询和管理。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 该字段用于返回所查询内容的任务状态。
	// <br>取值：**FINISH**（任务已完成）、**PENDING** （任务等待中）、**RUNNING** （任务进行中）、**ERROR** （任务出错）、**CANCELLED** （任务已取消）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 该字段用于返回视频审核任务所对应的任务名称，方便后续查询和管理审核任务。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 该字段用于返回调用视频审核接口时传入的BizType参数，方便数据的辨别和管理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段用于返回调用音频审核接口时输入的音频审核类型，取值为：**VIDEO**（点播视频）和**LIVE_VIDEO**（直播视频），默认值为VIDEO。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 该字段用于返回基于恶意标签的后续操作建议。当您获取到判定结果后，返回值表示具体的后续建议操作。<br>
	// 返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	Labels []*TaskLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 该字段用于返回输入媒体文件的详细信息，包括编码格式、分片时长等信息。详细内容敬请参考MediaInfo数据结构的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaInfo *MediaInfo `json:"MediaInfo,omitnil,omitempty" name:"MediaInfo"`

	// 该字段用于返回被查询任务创建的时间，格式采用 ISO 8601标准。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 该字段用于返回被查询任务最后更新时间，格式采用 ISO 8601标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 该字段用于返回审核服务的媒体内容信息，主要包括传入文件类型和访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputInfo *InputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`
}

type TaskFilter struct {
	// 该字段用于传入任务对应的业务类型供筛选器进行筛选。Biztype为策略的具体的编号，用于接口调度，在内容安全控制台中可配置。不同Biztype关联不同的业务场景与审核策略，调用前请确认正确的Biztype。Biztype仅为**数字、字母与下划线的组合**，长度为3-32个字符。<br>备注：在不传入该参数时筛选器默认不筛选业务类型。
	BizType []*string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段用于传入视频审核对应的任务类型供筛选器进行筛选，取值为：**VIDEO**（点播视频审核），**AUDIO**（点播音频审核）， **LIVE_VIDEO**（直播视频审核）, **LIVE_AUDIO**（直播音频审核）。<br>备注：在不传入该参数时筛选器默认不筛选任务类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 该字段用于传入视频审核对应的建议操作供筛选器进行筛选，取值为：**Block**：建议屏蔽，**Review**：建议人工复审，**Pass**：建议通过。<br>备注：在不传入该参数时筛选器默认不筛选建议操作。
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于传入审核任务的任务状态供筛选器进行筛选，取值为：**FINISH**（任务已完成）、**PENDING** （任务等待中）、**RUNNING** （任务进行中）、**ERROR** （任务出错）、**CANCELLED** （任务已取消）。<br>备注：在不传入该参数时筛选器默认不筛选任务状态。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`
}

type TaskInput struct {
	// 选填参数，该字段表示您为待检测对象分配的数据ID，传入后可方便您对文件进行标识和管理。<br>取值：由英文字母（大小写均可）、数字及四个特殊符号（_，-，@，#）组成，**长度不超过64个字符**。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 选填参数，该字段表示审核任务所对应的任务名称，方便后续查询和管理审核任务。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 必填参数，该字段表示审核文件的访问参数，用于获取审核媒体文件，该参数内包括访问类型和访问地址。
	Input *StorageInfo `json:"Input,omitnil,omitempty" name:"Input"`
}

type TaskLabel struct {
	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告，**Custom**：自定义违规；以及其他令人反感、不安全或不适宜的内容类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回当前标签（Label）下的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回当前标签（Label）下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容；*色情 0*，则表明该文本不属于色情内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回当前标签（Lable）下的二级标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`
}

type TaskResult struct {
	// 该字段用于返回创建视频审核任务时在TaskInput结构内传入的DataId，用于标识具体审核任务。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段用于返回视频审核任务所生成的任务ID，用于标识具体审核任务，方便后续查询和管理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 该字段用于返回任务创建的状态，如返回OK则代表任务创建成功，其他返回值可参考公共错误码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// **仅在Code的返回值为错误码时生效**，用于返回错误的详情内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type TextPosition struct {
	// 关键词在文本中的起始位置
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 关键词在文本中的结束位置
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

type TravelResult struct {
	// 一级标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 二级标签
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 风险等级
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 音频角色
	AudioRole *string `json:"AudioRole,omitnil,omitempty" name:"AudioRole"`

	// 音频语音文本
	AudioText *string `json:"AudioText,omitnil,omitempty" name:"AudioText"`

	// 开始时间
	StartTime *float64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *float64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}