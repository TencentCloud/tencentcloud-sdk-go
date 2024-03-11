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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccountInfo struct {
	// 用户账号类型；默认开通QQOpenId、手机号MD5权限；如果需要使用微信OpenId入参，则需要"提交工单"或联系对接人进行资格审核，审核通过后方可正常使用微信开放账号。
	// 1：QQ开放账号
	// 2：微信开放账号
	// 8：设备号，仅支持IMEI、IMEIMD5、IDFA、IDFAMD5
	// 10004：手机号MD5，中国大陆11位手机号进行MD5加密，取32位小写值。
	AccountType *uint64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// QQ账号信息，AccountType是"1"时，该字段必填。
	QQAccount *QQAccountInfo `json:"QQAccount,omitnil,omitempty" name:"QQAccount"`

	// 微信账号信息，AccountType是"2"时，该字段必填。
	WeChatAccount *WeChatAccountInfo `json:"WeChatAccount,omitnil,omitempty" name:"WeChatAccount"`

	// 其它账号信息，AccountType是8或10004时，该字段必填。
	OtherAccount *OtherAccountInfo `json:"OtherAccount,omitnil,omitempty" name:"OtherAccount"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BusinessSecurityData *InputFrontRisk `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type DescribeRiskTrendsRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputFrontRisk `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
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
	Data *OutputFrontRiskData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 是否授权：1已授权，否则未授权。
	//  调用全栈式风控引擎接口服务时，客户需先明确授权
	// 
	IsAuthorized *string `json:"IsAuthorized,omitnil,omitempty" name:"IsAuthorized"`

	// 加密类型：1AES加密
	CryptoType *string `json:"CryptoType,omitnil,omitempty" name:"CryptoType"`

	// 加密内容，非空时接口采用加密模式。
	CryptoContent *string `json:"CryptoContent,omitnil,omitempty" name:"CryptoContent"`
}

type InputDetails struct {
	// 字段名称
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 字段值
	FieldValue *string `json:"FieldValue,omitnil,omitempty" name:"FieldValue"`
}

type InputFrontRisk struct {
	// 事件ID
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 趋势类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 当前开始时间
	CurrentStartTime *string `json:"CurrentStartTime,omitnil,omitempty" name:"CurrentStartTime"`

	// 当前结束时间
	CurrentEndTime *string `json:"CurrentEndTime,omitnil,omitempty" name:"CurrentEndTime"`
}

type InputManageMarketingRisk struct {
	// 用户账号类型；默认开通QQOpenId、手机号MD5权限；如果需要使用微信OpenId入参，则需要"提交工单"或联系对接人进行资格审核，审核通过后方可正常使用微信开放账号。
	// 1：QQ开放账号
	// 2：微信开放账号
	// 8：设备号，仅支持IMEI、IMEIMD5、IDFA、IDFAMD5
	// 10004：手机号MD5，中国大陆11位手机号进行MD5加密，取32位小写值。
	Account *AccountInfo `json:"Account,omitnil,omitempty" name:"Account"`

	// 场景码，用于识别和区分不同的业务场景，可在控制台上新建和管理
	// 控制台链接：https://console.cloud.tencent.com/rce/risk/strategy/scene-root
	// 活动防刷默认场景码：e_activity_antirush 
	// 登录保护默认场景码：e_login_protection
	// 注册保护默认场景码：e_register_protection
	SceneCode *string `json:"SceneCode,omitnil,omitempty" name:"SceneCode"`

	// 用户外网ip（传入用户非外网ip会影响判断结果）。
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 用户操作时间戳，精确到秒。
	PostTime *uint64 `json:"PostTime,omitnil,omitempty" name:"PostTime"`

	// 业务平台用户唯一标识，支持自定义。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 设备指纹DeviceToken值，集成设备指纹后获取；如果集成了相应的设备指纹，该字段必填。
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// 设备指纹 BusinessId。
	DeviceBusinessId *int64 `json:"DeviceBusinessId,omitnil,omitempty" name:"DeviceBusinessId"`

	// 业务ID。网站或应用在多个业务中使用此服务，通过此ID区分统计数据。
	BusinessId *uint64 `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 昵称，UTF-8 编码。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 用户邮箱地址。
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// 是否识别设备异常：
	// 0：不识别。
	// 1：识别。
	CheckDevice *int64 `json:"CheckDevice,omitnil,omitempty" name:"CheckDevice"`

	// 用户HTTP请求中的Cookie进行2次hash的值，只要保证相同Cookie的hash值一致即可。
	CookieHash *string `json:"CookieHash,omitnil,omitempty" name:"CookieHash"`

	// 用户HTTP请求的Referer值。
	Referer *string `json:"Referer,omitnil,omitempty" name:"Referer"`

	// 用户HTTP请求的User-Agent值。
	UserAgent *string `json:"UserAgent,omitnil,omitempty" name:"UserAgent"`

	// 用户HTTP请求的X-Forwarded-For值。
	XForwardedFor *string `json:"XForwardedFor,omitnil,omitempty" name:"XForwardedFor"`

	// MAC地址或设备唯一标识。
	MacAddress *string `json:"MacAddress,omitnil,omitempty" name:"MacAddress"`

	// 手机制造商ID，如果手机注册，请带上此信息。
	VendorId *string `json:"VendorId,omitnil,omitempty" name:"VendorId"`

	// 设备类型，账号类型（AccountType）为8时填写。
	// 1:IMEI；国际移动设备识别号（15-17位数字）；
	// 2:IMEIMD5；国际移动设备识别号，通过MD5加密后取32位小写值；
	// 3:IDFA；
	// 4:IDFAMD5；国际移动设备识别号，通过MD5加密后取32位小写值。
	DeviceType *int64 `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 扩展字段。
	Details []*InputDetails `json:"Details,omitnil,omitempty" name:"Details"`

	// 邀请助力场景相关信息。
	Sponsor *SponsorInfo `json:"Sponsor,omitnil,omitempty" name:"Sponsor"`

	// 详情请跳转至OnlineScamInfo查看。
	OnlineScam *OnlineScamInfo `json:"OnlineScam,omitnil,omitempty" name:"OnlineScam"`

	// 1：Android
	// 2：iOS
	// 3：H5
	// 4：小程序
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

// Predefined struct for user
type ManageMarketingRiskRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputManageMarketingRisk `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`

	// 业务入参
	BusinessCryptoData *InputCryptoManageMarketingRisk `json:"BusinessCryptoData,omitnil,omitempty" name:"BusinessCryptoData"`
}

type ManageMarketingRiskRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputManageMarketingRisk `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`

	// 业务入参
	BusinessCryptoData *InputCryptoManageMarketingRisk `json:"BusinessCryptoData,omitnil,omitempty" name:"BusinessCryptoData"`
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
	Data *OutputManageMarketingRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ContentLabel *string `json:"ContentLabel,omitnil,omitempty" name:"ContentLabel"`

	// 内容风险等级：
	// 0：正常。
	// 1：可疑。
	ContentRiskLevel *int64 `json:"ContentRiskLevel,omitnil,omitempty" name:"ContentRiskLevel"`

	// 内容产生形式：
	// 0：对话。
	// 1：广播。
	ContentType *int64 `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 类型
	FraudType *int64 `json:"FraudType,omitnil,omitempty" name:"FraudType"`

	// 账号
	FraudAccount *string `json:"FraudAccount,omitnil,omitempty" name:"FraudAccount"`
}

type OtherAccountInfo struct {
	// 其他账号信息；
	// AccountType是8时，填入设备号（IMEI、IMEIMD5、IDFA、IDFAMD5）
	// AccountType是10004时，填入中国大陆标准11位手机号的MD5值
	// 注释：
	// MD5手机号加密方式，中国大陆11位手机号进行MD5加密，加密后取32位小写值
	// 设备号加密方式，对IMEI、IDFA明文进行MD5加密，加密后取32位小写值。
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// MD5手机号,AccountType是10004时，此处无需重复填写。
	MobilePhone *string `json:"MobilePhone,omitnil,omitempty" name:"MobilePhone"`

	// 用户设备号，AccountType是8时，此处无需重复填写。
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type OutputFrontRisk struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*OutputFrontRiskValue `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputFrontRiskData struct {
	// 返回码[0：成功；非0：标识失败错误码]。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 出错消息[UTF-8编码]。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*OutputFrontRisk `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputFrontRiskValue struct {
	// 请求次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Requests *int64 `json:"Requests,omitnil,omitempty" name:"Requests"`

	// 日期标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`
}

type OutputManageMarketingRisk struct {
	// 返回码。0表示成功，非0标识失败错误码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// UTF-8编码，出错消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 业务详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OutputManageMarketingRiskValue `json:"Value,omitnil,omitempty" name:"Value"`

	// 控制台显示的req_id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UUid *string `json:"UUid,omitnil,omitempty" name:"UUid"`
}

type OutputManageMarketingRiskValue struct {
	// 账号ID：对应输入参数。
	// 当AccountType为1时，对应QQ的OpenId。
	// 当AccountType为2时，对应微信的OpenId/UnionId。
	// 当AccountType为8时，对应IMEI、IDFA、IMEIMD5或者IDFAMD5。
	// 当AccountType为10004时，对应手机号的MD5值。
	// 请注意：此字段可能返回null，表示无法获取有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 操作时间戳，单位秒（对应输入参数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostTime *uint64 `json:"PostTime,omitnil,omitempty" name:"PostTime"`

	// 业务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateAccount *string `json:"AssociateAccount,omitnil,omitempty" name:"AssociateAccount"`

	// 操作来源的外网IP（对应输入参数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 风险等级
	// pass：无恶意
	// review：低风险，需要人工审核
	// reject：高风险，建议拦截
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 风险类型，可能同时命中多个风险类型
	// 1: 账号信用低，账号近期存在因恶意被处罚历史，网络低活跃，被举报等因素。
	// 11: 疑似低活跃账号，账号活跃度与正常用户有差异。
	// 2: 垃圾账号，疑似批量注册小号，近期存在严重违规或大量举报。
	// 21: 疑似小号，账号有疑似线上养号，小号等行为。
	// 22: 疑似违规账号，账号曾有违规行为、曾被举报过、曾因违规被处罚过等。
	// 3: 无效账号，送检账号参数无法成功解析，请检查微信 OpenId 是否有误/AppId 与 QQ OpenId 无法关联/微信 OpenId 权限是否开通/手机号是否为中国大陆手机号；
	// 4: 黑名单，该账号在业务侧有过拉黑记录。
	// 5: 白名单，业务自行有添加过白名单记录。
	// 101: 批量操作，存在 IP/设备/环境等因素的聚集性异常。
	// 1011: 疑似 IP 属性聚集，出现 IP 聚集。
	// 1012: 疑似设备属性聚集，出现设备聚集。
	// 102: 自动机，疑似自动机批量请求。
	// 103: 恶意行为-网赚，疑似网赚。
	// 104: 微信登录态无效，检查 WeChatAccessToken 参数，是否已经失效。
	// 201: 环境风险，环境异常操作 IP/设备/环境存在异常。当前 IP 为非常用 IP 或恶意 IP 段。
	// 2011: 疑似非常用IP，请求当前请求 IP 非该账号常用 IP。
	// 2012: 疑似 IP 异常，使用 IDC 机房 IP 或使用代理 IP 或使用恶意 IP 等。
	// 205: 非公网有效 IP，传进来的 IP 地址为内网 IP 地址或者 IP 保留地址。
	// 206: 设备异常，该设备存在异常的使用行为。
	// 2061: 疑似非常用设备，当前请求的设备非该账号常用设备。
	// 2062: 疑似虚拟设备，请求设备为模拟器、脚本、云设备等虚拟设备。
	// 2063: 疑似群控设备，请求设备为猫池、手机墙等群控设备。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType []*int64 `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// 设备指纹ID，如果集成了设备指纹，并传入了正确的DeviceToken和Platform，该字段正常输出；如果DeviceToken异常（校验不通过），则会在RiskType中返回"-1"标签，ConstId字段为空；如果没有集成设备指纹ConstId字段默认为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConstId *string `json:"ConstId,omitnil,omitempty" name:"ConstId"`

	// 风险扩展数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskInformation *string `json:"RiskInformation,omitnil,omitempty" name:"RiskInformation"`
}

type QQAccountInfo struct {
	// QQ的OpenId。
	QQOpenId *string `json:"QQOpenId,omitnil,omitempty" name:"QQOpenId"`

	// QQ分配给网站或应用的AppId，用来唯一标识网站或应用。
	AppIdUser *string `json:"AppIdUser,omitnil,omitempty" name:"AppIdUser"`

	// 用于标识QQ用户登录后所关联业务自身的账号ID。
	AssociateAccount *string `json:"AssociateAccount,omitnil,omitempty" name:"AssociateAccount"`

	// 账号绑定的MD5手机号，
	// 注释：只支中国大陆11位手机号MD5加密后位的32位小写字符串。
	MobilePhone *string `json:"MobilePhone,omitnil,omitempty" name:"MobilePhone"`

	// 用户设备号，支持IMEI、IMEIMD5、IDFA、IDFAMD5
	// 注释：IMEIMD5、IDFAMD5加密方式，对IMEI、IDFA明文进行MD5加密，加密后取32位小写值。
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type SponsorInfo struct {
	// 助力场景建议填写：活动发起人微信OpenId。
	SponsorOpenId *string `json:"SponsorOpenId,omitnil,omitempty" name:"SponsorOpenId"`

	// 助力场景建议填写：发起人设备号
	SponsorDeviceNumber *string `json:"SponsorDeviceNumber,omitnil,omitempty" name:"SponsorDeviceNumber"`

	// 助力场景建议填写：发起人的MD5手机号
	SponsorPhone *string `json:"SponsorPhone,omitnil,omitempty" name:"SponsorPhone"`

	// 助力场景建议填写：发起人IP
	SponsorIp *string `json:"SponsorIp,omitnil,omitempty" name:"SponsorIp"`

	// 助力场景建议填写：活动链接
	CampaignUrl *string `json:"CampaignUrl,omitnil,omitempty" name:"CampaignUrl"`
}

type WeChatAccountInfo struct {
	// 微信的OpenId/UnionId。
	WeChatOpenId *string `json:"WeChatOpenId,omitnil,omitempty" name:"WeChatOpenId"`

	// 微信开放账号类型：
	// 1：微信公众号/微信第三方登录。
	// 2：微信小程序。
	WeChatSubType *uint64 `json:"WeChatSubType,omitnil,omitempty" name:"WeChatSubType"`

	// 随机串。如果WeChatSubType是2，该字段必填。Token签名随机数，建议16个字符。
	RandStr *string `json:"RandStr,omitnil,omitempty" name:"RandStr"`

	// 如果WeChatSubType 是1，填入授权的 access_token（注意：不是普通 access_token，详情请参阅官方说明文档。获取网页版本的 access_token 时，scope 字段必需填写snsapi_userinfo
	// 如果WeChatSubType是2，填入以session_key 为密钥签名随机数RandStr（hmac_sha256签名算法）得到的字符串。
	WeChatAccessToken *string `json:"WeChatAccessToken,omitnil,omitempty" name:"WeChatAccessToken"`

	// 用于标识微信用户登录后所关联业务自身的账号ID。
	AssociateAccount *string `json:"AssociateAccount,omitnil,omitempty" name:"AssociateAccount"`

	// 账号绑定的MD5手机号，
	// 注释：只支持标准中国大陆11位手机号MD5加密后位的32位小写字符串。
	MobilePhone *string `json:"MobilePhone,omitnil,omitempty" name:"MobilePhone"`

	// 用户设备号，支持IMEI、IMEIMD5、IDFA、IDFAMD5
	// 注释：IMEIMD5、IDFAMD5加密方式，对IMEI、IDFA明文进行MD5加密，加密后取32位小写值。
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}