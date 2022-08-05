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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountInfo struct {
	// 账号类型
	AccountType *uint64 `json:"AccountType,omitempty" name:"AccountType"`

	// QQ账号信息，AccountType是1时，该字段必填。
	QQAccount *QQAccountInfo `json:"QQAccount,omitempty" name:"QQAccount"`

	// 微信账号信息，AccountType是2时，该字段必填。
	WeChatAccount *WeChatAccountInfo `json:"WeChatAccount,omitempty" name:"WeChatAccount"`

	// 其它账号信息，AccountType是0、4、8或10004时，该字段必填。
	OtherAccount *OtherAccountInfo `json:"OtherAccount,omitempty" name:"OtherAccount"`
}

// Predefined struct for user
type DescribeRiskAssessmentRequestParams struct {

}

type DescribeRiskAssessmentRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRiskAssessmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskAssessmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskAssessmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskAssessmentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRiskAssessmentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskAssessmentResponseParams `json:"Response"`
}

func (r *DescribeRiskAssessmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskAssessmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskTrendsRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputFrontRisk `json:"BusinessSecurityData,omitempty" name:"BusinessSecurityData"`
}

type DescribeRiskTrendsRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputFrontRisk `json:"BusinessSecurityData,omitempty" name:"BusinessSecurityData"`
}

func (r *DescribeRiskTrendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskTrendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskTrendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskTrendsResponseParams struct {
	// 业务出参
	Data *OutputFrontRiskData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRiskTrendsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskTrendsResponseParams `json:"Response"`
}

func (r *DescribeRiskTrendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InputCryptoManageMarketingRisk struct {
	// 是否授权
	IsAuthorized *string `json:"IsAuthorized,omitempty" name:"IsAuthorized"`

	// 加密类型
	CryptoType *string `json:"CryptoType,omitempty" name:"CryptoType"`

	// 加密内容
	CryptoContent *string `json:"CryptoContent,omitempty" name:"CryptoContent"`
}

type InputDetails struct {
	// 字段名称
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 字段值
	FieldValue *string `json:"FieldValue,omitempty" name:"FieldValue"`
}

type InputFrontRisk struct {
	// 事件ID
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 趋势类型
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 当前开始时间
	CurrentStartTime *string `json:"CurrentStartTime,omitempty" name:"CurrentStartTime"`

	// 当前结束时间
	CurrentEndTime *string `json:"CurrentEndTime,omitempty" name:"CurrentEndTime"`
}

type InputManageMarketingRisk struct {
	// 账号信息。
	Account *AccountInfo `json:"Account,omitempty" name:"Account"`

	// 场景类型：场景SceneCode, 控制台上新建对应的场景并获取对应的值；
	// 例如：e_register_protection_1521184361
	// 控制台链接：https://console.cloud.tencent.com/rce/risk/sceneroot；
	SceneCode *string `json:"SceneCode,omitempty" name:"SceneCode"`

	// 登录来源的外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 时间戳
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
	Details []*InputDetails `json:"Details,omitempty" name:"Details"`

	// 可选填写。详情请跳转至SponsorInfo查看。
	Sponsor *SponsorInfo `json:"Sponsor,omitempty" name:"Sponsor"`

	// 可选填写。详情请跳转至OnlineScamInfo查看。
	OnlineScam *OnlineScamInfo `json:"OnlineScam,omitempty" name:"OnlineScam"`

	// 平台: 1android
	Platform *string `json:"Platform,omitempty" name:"Platform"`
}

// Predefined struct for user
type ManageMarketingRiskRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputManageMarketingRisk `json:"BusinessSecurityData,omitempty" name:"BusinessSecurityData"`

	// 业务入参
	BusinessCryptoData *InputCryptoManageMarketingRisk `json:"BusinessCryptoData,omitempty" name:"BusinessCryptoData"`
}

type ManageMarketingRiskRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputManageMarketingRisk `json:"BusinessSecurityData,omitempty" name:"BusinessSecurityData"`

	// 业务入参
	BusinessCryptoData *InputCryptoManageMarketingRisk `json:"BusinessCryptoData,omitempty" name:"BusinessCryptoData"`
}

func (r *ManageMarketingRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageMarketingRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	delete(f, "BusinessCryptoData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageMarketingRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageMarketingRiskResponseParams struct {
	// 业务出参
	Data *OutputManageMarketingRisk `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ManageMarketingRiskResponse struct {
	*tchttp.BaseResponse
	Response *ManageMarketingRiskResponseParams `json:"Response"`
}

func (r *ManageMarketingRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 类型
	FraudType *int64 `json:"FraudType,omitempty" name:"FraudType"`

	// 账号
	FraudAccount *string `json:"FraudAccount,omitempty" name:"FraudAccount"`
}

type OtherAccountInfo struct {
	// id
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 手机号
	MobilePhone *string `json:"MobilePhone,omitempty" name:"MobilePhone"`

	// id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type OutputFrontRisk struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*OutputFrontRiskValue `json:"Value,omitempty" name:"Value"`
}

type OutputFrontRiskData struct {
	// 返回码[0：成功；非0：标识失败错误码]。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 出错消息[UTF-8编码]。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*OutputFrontRisk `json:"Value,omitempty" name:"Value"`
}

type OutputFrontRiskValue struct {
	// 请求次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Requests *int64 `json:"Requests,omitempty" name:"Requests"`

	// 日期标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *string `json:"Index,omitempty" name:"Index"`
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
	// 账号风险 
	// 1 账号信用低,账号近期存在因恶意被处罚历史，网络低活跃，被举报等因素
	// 11 疑似 低活跃账号,账号活跃度与正常用户有差异
	// 2 垃圾账号 疑似批量注册小号，近期存在严重违规或大量举报
	// 21 疑似小号 账号有疑似线上养号，小号等行为
	// 22 疑似违规账号 账号曾有违规行为、曾被举报过、曾因违规被处罚过等
	// 3 无效账号 送检账号参数无法成功解析，请检查微信 openid 是否有误/appid与QQopenid无法关联/微信openid权限是否有开通/手机号是否为中国大陆手机号；
	// 4 黑名单 该账号在业务侧有过拉黑记录
	// 5 白名单 业务自行有添加过白名单记录
	// 行为风险 
	// 101 批量操作 存在 ip/设备/环境等因素的聚集性异常
	// 1011 疑似 IP 属性聚集，出现 IP 聚集
	// 1012 疑似 设备属性聚集 出现设备聚集
	// 102 自动机 疑似自动机批量请求
	// 103 恶意行为-网赚 疑似网赚
	// 104 微信登录态无效 检查 WeChatAccessToken 参数，是否已经失效；
	// 201 环境风险 环境异常 操作 ip/设备/环境存在异常。当前 ip 为非常用 ip 或恶意 ip 段
	// 2011 疑似 非常用IP 请求 当前请求 IP 非该账号常用 IP
	// 2012 疑似 IP 异常 使用 idc 机房 ip 或 使用代理 ip 或 使用恶意 ip 等
	// 205 非公网有效ip 传进来的 IP 地址为内网 ip 地址或者 ip 保留地址；
	// 设备风险
	// 206  设备异常 该设备存在异常的使用行为
	// 2061 疑似 非常用设备 当前请求的设备非该账号常用设备
	// 2062 疑似 虚拟设备 请求设备为模拟器、脚本、云设备等虚拟设备
	// 2063 疑似 群控设备 请求设备为猫池、手机墙等群控设备
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType []*int64 `json:"RiskType,omitempty" name:"RiskType"`

	// 唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConstId *string `json:"ConstId,omitempty" name:"ConstId"`
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
	// OpenID
	SponsorOpenId *string `json:"SponsorOpenId,omitempty" name:"SponsorOpenId"`

	// 设备号
	SponsorDeviceNumber *string `json:"SponsorDeviceNumber,omitempty" name:"SponsorDeviceNumber"`

	// 手机号
	SponsorPhone *string `json:"SponsorPhone,omitempty" name:"SponsorPhone"`

	// IP
	SponsorIp *string `json:"SponsorIp,omitempty" name:"SponsorIp"`

	// 链接
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

	// token
	WeChatAccessToken *string `json:"WeChatAccessToken,omitempty" name:"WeChatAccessToken"`

	// 用于标识微信用户登录后所关联业务自身的账号ID。
	AssociateAccount *string `json:"AssociateAccount,omitempty" name:"AssociateAccount"`

	// 账号绑定的手机号。
	MobilePhone *string `json:"MobilePhone,omitempty" name:"MobilePhone"`

	// 用户设备号。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}