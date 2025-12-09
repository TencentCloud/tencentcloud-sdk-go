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

// Predefined struct for user
type CreateFinancialLLMTaskRequestParams struct {
	// 接口使用的识别策略 ID，请参考 [快速指引](https://cloud.tencent.com/document/product/1124/124604) 获取该值。  
	// 示例值：TencentCloudFinancialLLMDefault
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 送审内容的格式，有两个可选值：
	// - 1：代表送审内容为**文档**，如DOC文档
	// - 2：代表送审内容为**纯文本**
	// 
	// 示例值：1
	ContentType *int64 `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 若送审内容为文档（ContentType=1），需要传入具体格式，当前支持：DOC、DOCX、PDF。  
	// 说明：若送审内容为纯文本（ContentType=2），则本字段传空（FileType=""）。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 送审内容的传入方式如下：
	// - 若为文档类，需传入文档的URL（原文档文字数不超过10,000字），例如：http://xxxxxxxxxxxx/financial_test.doc
	// - 若为纯文本类，请以UTF-8格式进行Base64编码后传入（编码后字符数不超过10,000字），例如：5piO5aSpNjAz5LiA5a6a5rao
	// 
	// 示例值：5piO5aSpNjAz5LiA5a6a5rao
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type CreateFinancialLLMTaskRequest struct {
	*tchttp.BaseRequest
	
	// 接口使用的识别策略 ID，请参考 [快速指引](https://cloud.tencent.com/document/product/1124/124604) 获取该值。  
	// 示例值：TencentCloudFinancialLLMDefault
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 送审内容的格式，有两个可选值：
	// - 1：代表送审内容为**文档**，如DOC文档
	// - 2：代表送审内容为**纯文本**
	// 
	// 示例值：1
	ContentType *int64 `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 若送审内容为文档（ContentType=1），需要传入具体格式，当前支持：DOC、DOCX、PDF。  
	// 说明：若送审内容为纯文本（ContentType=2），则本字段传空（FileType=""）。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 送审内容的传入方式如下：
	// - 若为文档类，需传入文档的URL（原文档文字数不超过10,000字），例如：http://xxxxxxxxxxxx/financial_test.doc
	// - 若为纯文本类，请以UTF-8格式进行Base64编码后传入（编码后字符数不超过10,000字），例如：5piO5aSpNjAz5LiA5a6a5rao
	// 
	// 示例值：5piO5aSpNjAz5LiA5a6a5rao
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *CreateFinancialLLMTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFinancialLLMTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizType")
	delete(f, "ContentType")
	delete(f, "FileType")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFinancialLLMTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFinancialLLMTaskResponseParams struct {
	// 本次请求返回的任务ID将用于后续查询接口，以获取对应的审校结果。
	// 示例值：3570106e-b156-45d9-8af5-99b46f7eb2f9。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFinancialLLMTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateFinancialLLMTaskResponseParams `json:"Response"`
}

func (r *CreateFinancialLLMTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFinancialLLMTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailResults struct {
	// 该字段用于返回检测结果所对应的全部恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回对应当前标签的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测文本命中的关键词信息，用于标注文本违规的具体原因（如：*加我微信*）。该参数可能会有多个返回值，代表命中的多个关键词；如返回值为空且Score不为空，则代表识别结果所对应的恶意标签（Label）是来自于语义模型判断的返回值。
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 该字段用于返回当前标签（Label）下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容；*色情 0*，则表明该文本不属于色情内容。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回自定义关键词对应的词库类型，取值为**1**（黑白库）和**2**（自定义关键词库），若未配置自定义关键词库,则默认值为1（黑白库匹配）。
	LibType *int64 `json:"LibType,omitnil,omitempty" name:"LibType"`

	// 该字段用于返回自定义库的ID，以方便自定义库管理和配置。
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// 该字段用于返回自定义库的名称,以方便自定义库管理和配置。
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 该字段用于返回当前标签（Label）下的二级标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回当前一级标签（Label）下的关键词、子标签及分数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 该字段用于返回违规文本命中信息
	HitInfos []*HitInfo `json:"HitInfos,omitnil,omitempty" name:"HitInfos"`
}

type Device struct {
	// 该字段表示业务用户对应设备的IP地址。<br>
	// 备注:目前仅支持IPv4地址记录，不支持IPv6地址记录。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 该字段表示业务用户对应的MAC地址，以方便设备识别与管理；其格式与取值与标准MAC地址一致。
	Mac *string `json:"Mac,omitnil,omitempty" name:"Mac"`

	// *内测中，敬请期待。*
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// *内测中，敬请期待。*
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 该字段表示业务用户对应设备的**IMEI码**（国际移动设备识别码），该识别码可用于识别每一部独立的手机等移动通信设备，方便设备识别与管理。<br>备注：格式为**15-17位纯数字**。
	IMEI *string `json:"IMEI,omitnil,omitempty" name:"IMEI"`

	// **iOS设备专用**，该字段表示业务用户对应的**IDFA**(广告标识符),这是由苹果公司提供的用于标识用户的广告标识符，由一串16进制的32位数字和字母组成。<br>
	// 备注：苹果公司自2021年iOS14更新后允许用户手动关闭或者开启IDFA，故此字符串标记有效性可能有所降低。
	IDFA *string `json:"IDFA,omitnil,omitempty" name:"IDFA"`

	// **iOS设备专用**，该字段表示业务用户对应的**IDFV**(应用开发商标识符),这是由苹果公司提供的用于标注应用开发商的标识符，由一串16进制的32位数字和字母组成，可被用于唯一标识设备。
	IDFV *string `json:"IDFV,omitnil,omitempty" name:"IDFV"`
}

type FinancialLLMViolationDetail struct {
	// 违规点
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 处置建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 违规原因列表
	Reasons []*FinancialLLMViolationReason `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

type FinancialLLMViolationReason struct {
	// 违规原文片段
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// 违规原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

// Predefined struct for user
type GetFinancialLLMTaskResultRequestParams struct {
	// 该值对应创建任务接口里返回的TaskId字段值，创建任务接口见[创建金融大模型审校任务](https://cloud.tencent.com/document/product/1124/124463)。
	// 示例值：3570106e-b156-45d9-8af5-99b46f7eb2f9
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetFinancialLLMTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 该值对应创建任务接口里返回的TaskId字段值，创建任务接口见[创建金融大模型审校任务](https://cloud.tencent.com/document/product/1124/124463)。
	// 示例值：3570106e-b156-45d9-8af5-99b46f7eb2f9
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetFinancialLLMTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFinancialLLMTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFinancialLLMTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFinancialLLMTaskResultResponseParams struct {
	// TaskId对应的任务的状态：
	// - Success: 任务已完成
	// - Processing: 任务进行中，建议10秒后再查询
	// - Failed: 任务失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 该字段标识服务检测到的违规点，具体内容参阅数据结构[FinancialLLMViolationDetail](https://cloud.tencent.com/document/api/1124/51861#FinancialLLMViolationDetail)
	Details []*FinancialLLMViolationDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 本次检测的违规点列表
	ReviewedLabels []*string `json:"ReviewedLabels,omitnil,omitempty" name:"ReviewedLabels"`

	// 审校任务的开始时间
	// 示例值：2025-09-25 19:42:35
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 若审校任务失败（Status="Failed"），该字段返回失败的具体原因。
	// 示例值：文档解析失败
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// 该字段为历史结构字段，不再推荐使用。
	ModerationResult *string `json:"ModerationResult,omitnil,omitempty" name:"ModerationResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFinancialLLMTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *GetFinancialLLMTaskResultResponseParams `json:"Response"`
}

func (r *GetFinancialLLMTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFinancialLLMTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HitInfo struct {
	// 标识模型命中还是关键词命中
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 命中关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 自定义词库名称
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 位置信息
	Positions []*Positions `json:"Positions,omitnil,omitempty" name:"Positions"`
}

type Positions struct {
	// 关键词起始位置
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 关键词结束位置
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

type RiskDetails struct {
	// 该字段用于返回账号信息检测对应的风险类别，取值为：**RiskAccount**（账号存在风险）、**RiskIP**（IP地址存在风险）、**RiskIMEI**（移动设备识别码存在风险）。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回账号信息检测对应的风险等级，取值为：**1**（疑似存在风险）和**2**（存在恶意风险）。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`
}

type SentimentAnalysis struct {
	// 情感标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 标签分数，取值范围0到100
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 情感分析明细
	Detail *SentimentDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 响应码，成功为"OK"，失败为"InternalError"
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 异常信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type SentimentDetail struct {
	// 正向分数，取值范围0到100
	Positive *int64 `json:"Positive,omitnil,omitempty" name:"Positive"`

	// 负向分数，取值范围0到100
	Negative *int64 `json:"Negative,omitnil,omitempty" name:"Negative"`
}

type Tag struct {
	// 该字段用于返回命中的关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 该字段用于返回子标签
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回子标签对应的分数
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`
}

// Predefined struct for user
type TextModerationRequestParams struct {
	// 待检测的文本内容，需为UTF-8编码并以Base64格式传入。
	// 示例值：5L2g55qE5Lil6LCo6K6p5L2g5Y+R546w77yM5Lqn5ZOB57uP55CG5Y+r5YmR6Z2S
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 接口使用的识别策略编号，需在[控制台](https://console.cloud.tencent.com/cms/clouds/manage)获取。详细获取方式请参考以下链接：
	// - **内容安全**（详见步骤四：策略配置）：[点击这里](https://cloud.tencent.com/document/product/1124/37119)
	// - **AI生成识别**（详见服务对接->方式二）：[点击这里](https://cloud.tencent.com/document/product/1124/118694)
	// 
	// 示例值：TencentCloudDefault
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段表示您为待检测文本分配的数据ID，作用是方便您对数据进行标识和管理。
	// 取值：可由英文字母、数字、四种特殊符号（_，-，@，#）组成，**长度不超过64个字符**。
	// 示例值：a6127dd-c2a0-43e7-a3da-d27022d39ba7
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段标识用户信息，传入后可增强甄别有违规风险的发布者账号。
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 该字段标识设备信息，传入后可增强甄别有违规风险的发布者设备。
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// Content字段的原始语种，枚举值包括 zh 和 en：
	// - 推荐使用 zh
	// - en 适用于纯英文内容，耗时较高。若需使用 en，请先通过[反馈工单](https://console.cloud.tencent.com/workorder/category?level1_id=141&level2_id=1287&source=14&data_title=%E6%96%87%E6%9C%AC%E5%86%85%E5%AE%B9%E5%AE%89%E5%85%A8&step=1)确认
	// 
	// 示例值：zh
	SourceLanguage *string `json:"SourceLanguage,omitnil,omitempty" name:"SourceLanguage"`

	// 服务类型，枚举值包括 TEXT 和 TEXT_AIGC：
	// TEXT：内容安全
	// TEXT_AIGC：AI生成识别
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 适用于上下文关联审核场景，若多条文本内容需要联合审核，通过该字段关联会话。
	// 示例值：7e8f9a0b1c2d3e4f5a6b7c8d9e0f1a2b
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type TextModerationRequest struct {
	*tchttp.BaseRequest
	
	// 待检测的文本内容，需为UTF-8编码并以Base64格式传入。
	// 示例值：5L2g55qE5Lil6LCo6K6p5L2g5Y+R546w77yM5Lqn5ZOB57uP55CG5Y+r5YmR6Z2S
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 接口使用的识别策略编号，需在[控制台](https://console.cloud.tencent.com/cms/clouds/manage)获取。详细获取方式请参考以下链接：
	// - **内容安全**（详见步骤四：策略配置）：[点击这里](https://cloud.tencent.com/document/product/1124/37119)
	// - **AI生成识别**（详见服务对接->方式二）：[点击这里](https://cloud.tencent.com/document/product/1124/118694)
	// 
	// 示例值：TencentCloudDefault
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段表示您为待检测文本分配的数据ID，作用是方便您对数据进行标识和管理。
	// 取值：可由英文字母、数字、四种特殊符号（_，-，@，#）组成，**长度不超过64个字符**。
	// 示例值：a6127dd-c2a0-43e7-a3da-d27022d39ba7
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段标识用户信息，传入后可增强甄别有违规风险的发布者账号。
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 该字段标识设备信息，传入后可增强甄别有违规风险的发布者设备。
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// Content字段的原始语种，枚举值包括 zh 和 en：
	// - 推荐使用 zh
	// - en 适用于纯英文内容，耗时较高。若需使用 en，请先通过[反馈工单](https://console.cloud.tencent.com/workorder/category?level1_id=141&level2_id=1287&source=14&data_title=%E6%96%87%E6%9C%AC%E5%86%85%E5%AE%B9%E5%AE%89%E5%85%A8&step=1)确认
	// 
	// 示例值：zh
	SourceLanguage *string `json:"SourceLanguage,omitnil,omitempty" name:"SourceLanguage"`

	// 服务类型，枚举值包括 TEXT 和 TEXT_AIGC：
	// TEXT：内容安全
	// TEXT_AIGC：AI生成识别
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 适用于上下文关联审核场景，若多条文本内容需要联合审核，通过该字段关联会话。
	// 示例值：7e8f9a0b1c2d3e4f5a6b7c8d9e0f1a2b
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *TextModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextModerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "BizType")
	delete(f, "DataId")
	delete(f, "User")
	delete(f, "Device")
	delete(f, "SourceLanguage")
	delete(f, "Type")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextModerationResponseParams struct {
	// 该字段用于回显检测对象请求参数中的 BizType，与输入的 BizType 值对应。
	// 示例值：TencentCloudDefault
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 用于标识对本次请求的处置建议，共三种返回值。
	// 返回值：**Block**: 建议直接做违规处置，**Review**: 建议人工二次确认，**Pass**: 未识别到风险。
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测结果（DetailResults）中所对应的**优先级最高的恶意标签**，表示模型推荐的审核结果，建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；以及其他令人反感、不安全或不适宜的内容类型
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 对应 Label 字段下的二级子标签，表示该 Label 下更细分的违规点。
	// 示例值：SexualBehavior（该值为 Porn 下的一个二级标签）
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段标识 SubLabel 的置信度，取值范围为 0 - 100，值越高代表置信度越高。
	// 示例值：85
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段标识被检测文本所命中的关键词，可能返回0个或多个关键词。
	// 示例值：["优惠券", "线下兑换"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 该字段返回的检测的详细信息，返回值信息可参阅对应数据结构 DetailResults 的详细描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailResults []*DetailResults `json:"DetailResults,omitnil,omitempty" name:"DetailResults"`

	// 该字段标识入参 User 的检测结果，具体内容参阅数据结构 RiskDetails。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskDetails []*RiskDetails `json:"RiskDetails,omitnil,omitempty" name:"RiskDetails"`

	// 该字段用于返回根据您的需求配置的附加信息（Extra），如未配置则默认返回值为空。
	// 备注：不同客户或Biztype下返回信息不同，如需配置该字段请提交工单咨询或联系售后专员处理。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 该字段用于回显检测对象请求参数中的 DataId，与输入的 DataId 值对应。
	// 示例值：a6127dd-c2a0-43e7-a3da-d27022d39ba7
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 历史上下文关联的字段，不再推荐使用。上下文关联审核可通过入参的 SessionId 来实现。
	ContextText *string `json:"ContextText,omitnil,omitempty" name:"ContextText"`

	// 该字段为历史结构字段，不再推荐使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SentimentAnalysis *SentimentAnalysis `json:"SentimentAnalysis,omitnil,omitempty" name:"SentimentAnalysis"`

	// 该字段为历史结构字段，不再推荐使用。
	HitType *string `json:"HitType,omitnil,omitempty" name:"HitType"`

	// 该字段用于回显检测对象请求参数中的 SessionId，与输入的 SessionId 值对应。
	// 示例值：7e8f9a0b1c2d3e4f5a6b7c8d9e0f1a2b
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextModerationResponse struct {
	*tchttp.BaseResponse
	Response *TextModerationResponseParams `json:"Response"`
}

func (r *TextModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {
	// 该字段表示业务用户ID,填写后，系统可根据账号过往违规历史优化审核结果判定，有利于存在可疑违规风险时的辅助判断。<br>
	// 备注：该字段可传入微信openid、QQopenid、字符串等账号信息，与账号类别参数（AccountType）配合使用可确定唯一账号。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 该字段表示业务用户对应的账号昵称信息。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 该字段表示业务用户ID对应的账号类型，取值：**1**-微信uin，**2**-QQ号，**3**-微信群uin，**4**-qq群号，**5**-微信openid，**6**-QQopenid，**7**-其它string。<br>
	// 该字段与账号ID参数（UserId）配合使用可确定唯一账号。
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 该字段表示业务用户对应账号的性别信息。<br>
	// 取值：**0**（默认值，代表性别未知）、**1**（男性）、**2**（女性）。
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 该字段表示业务用户对应账号的年龄信息。<br>
	// 取值：**0**（默认值，代表年龄未知）-（**自定义年龄上限**）之间的整数。
	Age *int64 `json:"Age,omitnil,omitempty" name:"Age"`

	// 该字段表示业务用户对应账号的等级信息。<br>
	// 取值：**0**（默认值，代表等级未知）、**1**（等级较低）、**2**（等级中等）、**3**（等级较高），目前**暂不支持自定义等级**。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 该字段表示业务用户对应账号的手机号信息，支持全球各地区手机号的记录。<br>
	// 备注：请保持手机号格式的统一，如区号格式（086/+86）等。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 该字段表示业务用户头像图片的访问链接(URL)，支持PNG、JPG、JPEG、BMP、GIF、WEBP格式。
	// 备注：头像图片大小不超过5MB，建议分辨率不低于256x256；图片下载时间限制为3秒，超过则会返回下载超时。
	HeadUrl *string `json:"HeadUrl,omitnil,omitempty" name:"HeadUrl"`

	// 该字段表示业务用户的简介信息，支持汉字、英文及特殊符号，长度不超过5000个汉字字符。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 该字段表示业务群聊场景时的房间ID。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 该字段表示消息接受者ID
	ReceiverId *string `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 消息生成时间，精确到毫秒
	SendTime *int64 `json:"SendTime,omitnil,omitempty" name:"SendTime"`
}