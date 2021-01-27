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

package v20201103

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountInfo struct {

	// 用户账号类型（默认开通 QQ 开放账号、手机号，手机 MD5 账号类型查询。如需使用微信开放账号，则需要 提交工单 由腾讯云进行资格审核，审核通过后方可正常使用微信开放账号）：
	// 1：QQ开放账号。
	// 2：微信开放账号。
	// 4：手机号（暂仅支持国内手机号）。
	// 8：设备号（imei/imeiMD5/idfa/idfaMd5）。
	// 0：其他。
	// 10004：手机号MD5（标准中国大陆手机号11位，MD5后取32位小写值）
	AccountType *uint64 `json:"AccountType,omitempty" name:"AccountType"`

	// QQ账号信息，AccountType是1时，该字段必填。
	QQAccount *QQAccountInfo `json:"QQAccount,omitempty" name:"QQAccount"`

	// 微信账号信息，AccountType是2时，该字段必填。
	WeChatAccount *WeChatAccountInfo `json:"WeChatAccount,omitempty" name:"WeChatAccount"`

	// 其它账号信息，AccountType是0、4、8或10004时，该字段必填。
	OtherAccount *OtherAccountInfo `json:"OtherAccount,omitempty" name:"OtherAccount"`
}

type InputDetails struct {

	// 字段名称
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 字段值
	FieldValue *string `json:"FieldValue,omitempty" name:"FieldValue"`
}

type InputManageMarketingRisk struct {

	// 账号信息。
	Account *AccountInfo `json:"Account,omitempty" name:"Account"`

	// 场景code
	SceneCode *string `json:"SceneCode,omitempty" name:"SceneCode"`

	// 登录来源的外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 用户操作时间戳，单位秒（格林威治时间精确到秒，如1501590972）。
	PostTime *uint64 `json:"PostTime,omitempty" name:"PostTime"`

	// 用户唯一标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 设备指纹token。
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`

	// 设备指纹BusinessId
	DeviceBusinessId *int64 `json:"DeviceBusinessId,omitempty" name:"DeviceBusinessId"`

	// 业务ID。网站或应用在多个业务中使用此服务，通过此ID区分统计数据。
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 昵称，UTF-8 编码。
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 用户邮箱地址（非系统自动生成）。
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 是否识别设备异常：
	// 0：不识别。
	// 1：识别。
	CheckDevice *int64 `json:"CheckDevice,omitempty" name:"CheckDevice"`

	// 用户HTTP请求中的Cookie进行2次hash的值，只要保证相同Cookie的hash值一致即可。
	CookieHash *string `json:"CookieHash,omitempty" name:"CookieHash"`

	// 用户HTTP请求的Referer值。
	Referer *string `json:"Referer,omitempty" name:"Referer"`

	// 用户HTTP请求的User-Agent值。
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// 用户HTTP请求的X-Forwarded-For值。
	XForwardedFor *string `json:"XForwardedFor,omitempty" name:"XForwardedFor"`

	// MAC地址或设备唯一标识。
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机制造商ID，如果手机注册，请带上此信息。
	VendorId *string `json:"VendorId,omitempty" name:"VendorId"`

	// 设备类型：
	// 1：Android
	// 2：IOS
	DeviceType *int64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 详细信息
	Details []*InputDetails `json:"Details,omitempty" name:"Details" list`

	// 可选填写。详情请跳转至SponsorInfo查看。
	Sponsor *SponsorInfo `json:"Sponsor,omitempty" name:"Sponsor"`

	// 可选填写。详情请跳转至OnlineScamInfo查看。
	OnlineScam *OnlineScamInfo `json:"OnlineScam,omitempty" name:"OnlineScam"`
}

type ManageMarketingRiskRequest struct {
	*tchttp.BaseRequest

	// 业务入参
	BusinessSecurityData *InputManageMarketingRisk `json:"BusinessSecurityData,omitempty" name:"BusinessSecurityData"`
}

func (r *ManageMarketingRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageMarketingRiskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManageMarketingRiskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务出参
		Data *OutputManageMarketingRisk `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageMarketingRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageMarketingRiskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OnlineScamInfo struct {

	// 内容标签。
	ContentLabel *string `json:"ContentLabel,omitempty" name:"ContentLabel"`

	// 内容风险等级：
	// 0：正常。
	// 1：可疑。
	ContentRiskLevel *int64 `json:"ContentRiskLevel,omitempty" name:"ContentRiskLevel"`

	// 内容产生形式：
	// 0：对话。
	// 1：广播。
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`

	// 诈骗账号类型：
	// 1：11位手机号。
	// 2：QQ账号。
	FraudType *int64 `json:"FraudType,omitempty" name:"FraudType"`

	// 诈骗账号，手机号或QQ账号。
	FraudAccount *string `json:"FraudAccount,omitempty" name:"FraudAccount"`
}

type OtherAccountInfo struct {

	// 其它账号信息：
	// AccountType是4时，填入真实的手机号（如13123456789）。
	// AccountType是8时，支持 imei、idfa、imeiMD5、idfaMD5 入参。
	// AccountType是0时，填入账号信息。
	// AccountType是10004时，填入手机号的MD5值。
	// 注：imeiMd5 加密方式为：imei 明文小写后，进行 MD5 加密，加密后取小写值。IdfaMd5 加密方式为：idfa 明文大写后，进行 MD5 加密，加密后取小写值。
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 手机号，若 AccountType 是4（手机号）、或10004（手机号 MD5），则无需重复填写，否则填入对应的手机号（如13123456789）。
	MobilePhone *string `json:"MobilePhone,omitempty" name:"MobilePhone"`

	// 用户设备号。若 AccountType 是8（设备号），则无需重复填写，否则填入对应的设备号。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type OutputManageMarketingRisk struct {

	// 返回码。0表示成功，非0标识失败错误码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// UTF-8编码，出错消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 业务详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OutputManageMarketingRiskValue `json:"Value,omitempty" name:"Value"`

	// 控制台显示的req_id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UUid *string `json:"UUid,omitempty" name:"UUid"`
}

type OutputManageMarketingRiskValue struct {

	// 账号ID。对应输入参数：
	// AccountType是1时，对应QQ的OpenID。
	// AccountType是2时，对应微信的OpenID/UnionID。
	// AccountType是4时，对应手机号。
	// AccountType是8时，对应imei、idfa、imeiMD5或者idfaMD5。
	// AccountType是0时，对应账号信息。
	// AccountType是10004时，对应手机号的MD5。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 操作时间戳，单位秒（对应输入参数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostTime *uint64 `json:"PostTime,omitempty" name:"PostTime"`

	// 对应输入参数，AccountType 是 QQ 或微信开放账号时，用于标识 QQ 或微信用户登录后关联业务自身的账号ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateAccount *string `json:"AssociateAccount,omitempty" name:"AssociateAccount"`

	// 操作来源的外网IP（对应输入参数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 风险值
	// pass : 无恶意
	// review：需要人工审核
	// reject：拒绝，高风险恶意
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 风险类型，请参考官网风险类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType []*int64 `json:"RiskType,omitempty" name:"RiskType" list`
}

type QQAccountInfo struct {

	// QQ的OpenID。
	QQOpenId *string `json:"QQOpenId,omitempty" name:"QQOpenId"`

	// QQ分配给网站或应用的AppId，用来唯一标识网站或应用。
	AppIdUser *string `json:"AppIdUser,omitempty" name:"AppIdUser"`

	// 用于标识QQ用户登录后所关联业务自身的账号ID。
	AssociateAccount *string `json:"AssociateAccount,omitempty" name:"AssociateAccount"`

	// 账号绑定的手机号。
	MobilePhone *string `json:"MobilePhone,omitempty" name:"MobilePhone"`

	// 用户设备号。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type SponsorInfo struct {

	// 助力场景建议填写：活动发起人微信OpenID。
	SponsorOpenId *string `json:"SponsorOpenId,omitempty" name:"SponsorOpenId"`

	// 助力场景建议填写：发起人设备号。
	SponsorDeviceNumber *string `json:"SponsorDeviceNumber,omitempty" name:"SponsorDeviceNumber"`

	// 助力场景建议填写：发起人手机号。
	SponsorPhone *string `json:"SponsorPhone,omitempty" name:"SponsorPhone"`

	// 助力场景建议填写：发起人IP。
	SponsorIp *string `json:"SponsorIp,omitempty" name:"SponsorIp"`

	// 助力场景建议填写：活动链接。
	CampaignUrl *string `json:"CampaignUrl,omitempty" name:"CampaignUrl"`
}

type WeChatAccountInfo struct {

	// 微信的OpenID/UnionID 。
	WeChatOpenId *string `json:"WeChatOpenId,omitempty" name:"WeChatOpenId"`

	// 微信开放账号类型：
	// 1：微信公众号/微信第三方登录。
	// 2：微信小程序。
	WeChatSubType *uint64 `json:"WeChatSubType,omitempty" name:"WeChatSubType"`

	// 随机串。如果WeChatSubType是2，该字段必填。Token签名随机数，建议16个字符。
	RandStr *string `json:"RandStr,omitempty" name:"RandStr"`

	// 如果WeChatSubType是1，填入授权的access_token（注意：不是普通access_token，详情请参阅官方说明文档。获取网页版本的access_token时，scope字段必需填写snsapi_userinfo。
	// 如果WeChatSubType是2，填入以session_key为密钥签名随机数RandStr（hmac_sha256签名算法）得到的字符串。
	WeChatAccessToken *string `json:"WeChatAccessToken,omitempty" name:"WeChatAccessToken"`

	// 用于标识微信用户登录后所关联业务自身的账号ID。
	AssociateAccount *string `json:"AssociateAccount,omitempty" name:"AssociateAccount"`

	// 账号绑定的手机号。
	MobilePhone *string `json:"MobilePhone,omitempty" name:"MobilePhone"`

	// 用户设备号。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}
