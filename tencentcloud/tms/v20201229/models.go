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
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DetailResults struct {

	// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义词库。
	// 以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 建议您拿到判断结果后的执行操作。
	// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 该标签下命中的关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 该标签模型命中的分值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 仅当Label为Custom自定义关键词时有效，表示自定义关键词库类型，1:黑白库，2：自定义库
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibType *int64 `json:"LibType,omitempty" name:"LibType"`

	// 仅当Label为Custom自定义关键词时有效，表示自定义库id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// 仅当Labe为Custom自定义关键词时有效，表示自定义库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitempty" name:"LibName"`
}

type Device struct {

	// 用户IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// Mac地址
	Mac *string `json:"Mac,omitempty" name:"Mac"`

	// 设备指纹Token
	TokenId *string `json:"TokenId,omitempty" name:"TokenId"`

	// 设备指纹ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备序列号
	IMEI *string `json:"IMEI,omitempty" name:"IMEI"`

	// IOS设备，Identifier For Advertising（广告标识符）
	IDFA *string `json:"IDFA,omitempty" name:"IDFA"`

	// IOS设备，IDFV - Identifier For Vendor（应用开发商标识符）
	IDFV *string `json:"IDFV,omitempty" name:"IDFV"`
}

type RiskDetails struct {

	// 风险类别，RiskAccount，RiskIP, RiskIMEI
	Label *string `json:"Label,omitempty" name:"Label"`

	// 风险等级，1:疑似，2：恶意
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

type TextModerationRequest struct {
	*tchttp.BaseRequest

	// 文本内容Base64编码。原文长度需小于15000字节，即5000个汉字以内。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 该字段用于标识业务场景。您可以在内容安全控制台创建对应的ID，配置不同的内容审核策略，通过接口调用，默认不填为0，后端使用默认策略。 -- 该字段暂未开放。
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 数据ID，英文字母、下划线、-组成，不超过64个字符
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 账号相关信息字段，填入后可识别违规风险账号。
	User *User `json:"User,omitempty" name:"User"`

	// 设备相关信息字段，填入后可识别违规风险设备。
	Device *Device `json:"Device,omitempty" name:"Device"`
}

func (r *TextModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return errors.New("TextModerationRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TextModerationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 您在入参时所填入的Biztype参数。 -- 该字段暂未开放。
		BizType *string `json:"BizType,omitempty" name:"BizType"`

		// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义词库。
	// 以及其他令人反感、不安全或不适宜的内容类型。
		Label *string `json:"Label,omitempty" name:"Label"`

		// 建议您拿到判断结果后的执行操作。
	// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
		Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

		// 文本命中的关键词信息，用于提示您文本违规的具体原因，可能会返回多个命中的关键词。（如：加我微信）
	// 如返回值为空，Score不为空，即识别结果（Label）是来自于语义模型判断的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

		// 机器判断当前分类的置信度，取值范围：0.00~100.00。分数越高，表示越有可能属于当前分类。
	// （如：色情 99.99，则该样本属于色情的置信度非常高。）
		Score *int64 `json:"Score,omitempty" name:"Score"`

		// 接口识别样本后返回的详细结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DetailResults []*DetailResults `json:"DetailResults,omitempty" name:"DetailResults"`

		// 接口识别样本中存在违规账号风险的检测结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		RiskDetails []*RiskDetails `json:"RiskDetails,omitempty" name:"RiskDetails"`

		// 扩展字段，用于特定信息返回，不同客户/Biztype下返回信息不同。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Extra *string `json:"Extra,omitempty" name:"Extra"`

		// 请求参数中的DataId
	// 注意：此字段可能返回 null，表示取不到有效值。
		DataId *string `json:"DataId,omitempty" name:"DataId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {

	// 用户账号ID，如填写，会根据账号历史恶意情况，判定消息有害结果，特别是有利于可疑恶意情况下的辅助判断。账号可以填写微信uin、QQ号、微信openid、QQopenid、字符串等。该字段和账号类别确定唯一账号。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 账号类别，"1-微信uin 2-QQ号 3-微信群uin 4-qq群号 5-微信openid 6-QQopenid 7-其它string"
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 性别 默认0 未知 1 男性 2 女性
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 年龄 默认0 未知
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 用户等级，默认0 未知 1 低 2 中 3 高
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`
}
