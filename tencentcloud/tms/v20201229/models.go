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
	// 该字段表示待检测对象的文本内容，文本需要按utf-8格式编码，长度不能超过10000个字符（按unicode编码计算），并进行 Base64加密
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 该字段表示使用的策略的具体编号，该字段需要先在[内容安全控制台](https://console.cloud.tencent.com/cms/clouds/manage)中配置。
	// 备注：不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段表示您为待检测对象分配的数据ID，传入后可方便您对文件进行标识和管理。<br>取值：由英文字母（大小写均可）、数字及四个特殊符号（_，-，@，#）组成，**长度不超过64个字符**
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段表示待检测对象对应的用户相关信息，传入后可便于甄别相应违规风险用户
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 该字段表示待检测对象对应的设备相关信息，传入后可便于甄别相应违规风险设备
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// 表示Content的原始语种，枚举值包括 "en" 和 "zh"。其中，"en" 表示英文，"zh" 表示中文。非中文场景的处理耗时较高，具体情况取决于送审文本长度，非中文场景需[反馈工单](https://console.cloud.tencent.com/workorder/category?level1_id=141&level2_id=1287&source=14&data_title=%E6%96%87%E6%9C%AC%E5%86%85%E5%AE%B9%E5%AE%89%E5%85%A8&step=1)确认。
	SourceLanguage *string `json:"SourceLanguage,omitnil,omitempty" name:"SourceLanguage"`

	// 审核的业务类型，枚举值包括"TEXT"和"TEXT_AIGC"。其中"TEXT"表示传统文本审核，"TEXT_AIGC"表示AI生成检测（生成检测能力具体能力了解可[参见文档](https://cloud.tencent.com/document/product/1124/118694)）。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type TextModerationRequest struct {
	*tchttp.BaseRequest
	
	// 该字段表示待检测对象的文本内容，文本需要按utf-8格式编码，长度不能超过10000个字符（按unicode编码计算），并进行 Base64加密
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 该字段表示使用的策略的具体编号，该字段需要先在[内容安全控制台](https://console.cloud.tencent.com/cms/clouds/manage)中配置。
	// 备注：不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段表示您为待检测对象分配的数据ID，传入后可方便您对文件进行标识和管理。<br>取值：由英文字母（大小写均可）、数字及四个特殊符号（_，-，@，#）组成，**长度不超过64个字符**
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段表示待检测对象对应的用户相关信息，传入后可便于甄别相应违规风险用户
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 该字段表示待检测对象对应的设备相关信息，传入后可便于甄别相应违规风险设备
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// 表示Content的原始语种，枚举值包括 "en" 和 "zh"。其中，"en" 表示英文，"zh" 表示中文。非中文场景的处理耗时较高，具体情况取决于送审文本长度，非中文场景需[反馈工单](https://console.cloud.tencent.com/workorder/category?level1_id=141&level2_id=1287&source=14&data_title=%E6%96%87%E6%9C%AC%E5%86%85%E5%AE%B9%E5%AE%89%E5%85%A8&step=1)确认。
	SourceLanguage *string `json:"SourceLanguage,omitnil,omitempty" name:"SourceLanguage"`

	// 审核的业务类型，枚举值包括"TEXT"和"TEXT_AIGC"。其中"TEXT"表示传统文本审核，"TEXT_AIGC"表示AI生成检测（生成检测能力具体能力了解可[参见文档](https://cloud.tencent.com/document/product/1124/118694)）。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextModerationResponseParams struct {
	// 该字段用于返回请求参数中的BizType参数
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段用于返回检测结果（DetailResults）中所对应的**优先级最高的恶意标签**，表示模型推荐的审核结果，建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；以及其他令人反感、不安全或不适宜的内容类型
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回当前标签（Label）下被检测文本命中的关键词信息，用于标注文本违规的具体原因（如：*加我微信*）。该参数可能会有多个返回值，代表命中的多个关键词；如返回值为空且Score不为空，则代表识别结果所对应的恶意标签（Label）是来自于语义模型判断的返回值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 该字段用于返回当前标签（Label）下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容；*色情 0*，则表明该文本不属于色情内容
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回基于文本风险库审核的详细结果，返回值信息可参阅对应数据结构（DetailResults）的详细描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailResults []*DetailResults `json:"DetailResults,omitnil,omitempty" name:"DetailResults"`

	// 该字段用于返回文本检测中存在违规风险的账号检测结果，主要包括违规风险类别和风险等级信息，具体内容可参阅对应数据结构（RiskDetails）的详细描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskDetails []*RiskDetails `json:"RiskDetails,omitnil,omitempty" name:"RiskDetails"`

	// 该字段用于返回根据您的需求配置的额外附加信息（Extra），如未配置则默认返回值为空。<br>备注：不同客户或Biztype下返回信息不同，如需配置该字段请提交工单咨询或联系售后专员处理
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 该字段用于返回检测对象对应请求参数中的DataId，与输入的DataId字段中的内容对应
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段用于返回当前标签（Label）下的二级标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回上下文关联文本
	ContextText *string `json:"ContextText,omitnil,omitempty" name:"ContextText"`

	// 情感分析结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SentimentAnalysis *SentimentAnalysis `json:"SentimentAnalysis,omitnil,omitempty" name:"SentimentAnalysis"`

	// 该字段用于标识本次审核决策归因，比如text_nlp_tianji标识是由nlp tianji模型给出的审核决策，text_keyword_public标识命中了业务的关键词库
	HitType *string `json:"HitType,omitnil,omitempty" name:"HitType"`

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