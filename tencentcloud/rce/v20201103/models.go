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
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountInfo struct {

	// 用户账号类型（默认开通 QQ 开放账号、手机号，手机 MD5 账号类型查询。如需使用微信开放账号，则需要 提交工单 由腾讯云进行资格审核，审核通过后方可正常使用微信开放账号）：
	// 1：QQ开放账号。
	// 2：微信开放账号。
	// 4：手机号（暂仅支持国内手机号）。
	// 8：设备号（imei/imeiMD5/idfa/idfaMD5）。
	// 0：其他。
	// 10004：支持手机号MD5加密和sha256加密；
	// 标准中国大陆手机号11位，MD5加密后取长度32位小写字符串；例如：手机号13112345678的Md5加密结果为手"dafc728802534d51fbf85c70313a2bd2"
	// 标准中国大陆手机号11位，sha256加密后取长度为64位的小写字符串；例如：手机号13112345678的sha256加密的结果为“9f46715cff1a9ac969ec01924111f7f3697a97ad98a4fd53e15a78d79d1f3551”
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

	// 场景类型：场景SceneCode, 控制台上新建对应的场景并获取对应的值；
	// 例如：e_register_protection_1521184361
	// 控制台链接：https://console.cloud.tencent.com/rce/risk/sceneroot；
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
	// FieldName 清单
	// Android serial_number String 否 设备序列号
	// Android carrier String 否 运营商；-1: 获取失败，0: 其他，1: 移动，2: 联通，3: 电信，4: 铁通
	// Android mcc_mnc String 否 netOperator MCC+MNC
	// Android model String 否 手机型号
	// Android os_system String 否 操作系统
	// Android vendor_id String 否 设备品牌 “华为”“oppo”“小米”
	// Android device_version String 否 设备版本
	// Android android_api_level String 否 安卓API等级
	// Android phone_chip_info String 否 手机芯片信息
	// Android resolution_w String 否 屏幕分辨率宽，保留整数
	// Android resolution_h String 否 屏幕分辨率高，保留整数
	// Android brightness String 否 屏幕亮度
	// Android bluetooth_address String 否 蓝牙地址
	// Android baseband_version String 否 基带版本
	// Android kernel_version String 否 kernel 版本
	// Android cpu_core String 否 CPU 核数
	// Android cpu_model String 否 CPU 型号
	// Android memory String 否 内存容量，单位转换为 GB
	// Android storage String 否 存储容量，单位转换为 GB
	// Android volume String 否 手机音量
	// Android battery_power String 否 电池电量
	// Android language String 否 语言
	// Android package_name String 否 软件包名
	// Android App_version String 否 App 版本号
	// Android App_name String 否 App 显示名称
	// Android is_debug String 否 是否 debug；0 为正常模式，1 为 debug 模式；其他值无效
	// Android is_root String 否 是否越狱；0 为正常，1 为越狱；其他值无效
	// Android is_proxy String 否 是否启动代理；0 为未开启，1 为开启；其他值无效
	// Android is_emulator String 否 是否模拟器；0 为未开启，1 为开启；其他值无效
	// Android charge_status String 否 充电状态；1-不在充电，2-USB充电，3-电源充电
	// Android network_type String 否 网络类型：2G/3G/4G/5G/WiFi/WWAN/other
	// Android wifi_mac String 否 WiFi MAC地址
	// IOS model String 否 机器型号 iPhone11
	// IOS memory String 否 内存容量，单位转换为 GB
	// IOS os_system String 否 操作系统
	// IOS device_version String 否 设备版本
	// IOS phone_chip_info String 否 手机芯片信息
	// IOS device_name String 否 设备名称 "xxx 的 iPhone"，"xxx's IPhone" 等等
	// IOS uptime String 否 开机时间
	// IOS language String 否 系统语言
	// IOS carrier String 否 运营商
	// IOS cpu_model String 否 CPU 型号
	// IOS cpu_core String 否 CPU 个数
	// IOS volume String 否 手机音量
	// IOS battery_power String 否 电池电量
	// IOS resolution_w String 否 屏幕分辨率宽，保留整数
	// IOS resolution_h String 否 屏幕分辨率高，保留整数
	// IOS package_name String 否 App 包名
	// IOS App_version String 否 App 版本号
	// IOS App_name String 否 App 显示名称
	// IOS is_debug String 否 是否 debug；0 为正常模式，1 为 debug 模式；其他值无效
	// IOS is_root String 否 是否越狱；0 为正常，1 为越狱；其他值无效
	// IOS is_proxy String 否 是否启动代理；0 为未开启，1 为开启；其他值无效
	// IOS is_emulator String 否 是否模拟器；0 为未开启，1 为开启；其他值无效
	// IOS charge_status String 否 充电状态；1-不在充电，2-USB充电，3-电源充电
	// IOS network_type String 否 网络类型：2G/3G/4G/5G/WiFi/WWAN/other
	// IOS wifi_mac String 否 WiFi MAC地址
	// 其他 os_system String 否 操作系统
	// 其他 browser String 否 浏览器信息
	// 其他 from_url String 否 来源链接
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageMarketingRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return errors.New("ManageMarketingRiskRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
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
