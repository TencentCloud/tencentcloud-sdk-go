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

package v20200224

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type QueryActivityAntiRushRequest struct {
	*tchttp.BaseRequest

	// 用户账号类型（默认开通 QQ 开放账号、手机号，手机 MD5 账号类型查询。如需使用微信开放账号，则需要 提交工单 由腾讯云进行资格审核，审核通过后方可正常使用微信开放账号）：
	// 1：QQ 开放帐号。
	// 2：微信开放账号。
	// 4：手机号。
	// 0：其他。
	// 10004：手机号 MD5。
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 用户 ID 不同的 accountType 对应不同的用户 ID。如果是 QQ，则填入对应的 openid，微信用户则填入对应的 openid/unionid，手机号则填入对应真实用户手机号（如13123456789）。
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 用户的真实外网 IP。若填入非外网有效ip，会返回level=0的风控结果，risktype中会有205的风险码返回作为标识
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 用户操作时间戳。
	PostTime *string `json:"PostTime,omitempty" name:"PostTime"`

	// accountType 是QQ开放账号时，该参数必填，表示 QQ 开放平台分配给网站或应用的 AppID，用来唯一标识网站或应用。
	AppIdU *string `json:"AppIdU,omitempty" name:"AppIdU"`

	// 昵称，UTF-8 编码。
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 手机号。若 accountType 选4（手机号）、或10004（手机号 MD5），则无需重复填写。否则填入对应的手机号（如15912345687）。accountType为1或2时，该字段支持MD5值；
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 用户邮箱地址。
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 注册时间戳。
	RegisterTime *string `json:"RegisterTime,omitempty" name:"RegisterTime"`

	// 注册来源的外网 IP。
	RegisterIp *string `json:"RegisterIp,omitempty" name:"RegisterIp"`

	// 用户 HTTP 请求中的 cookie 进行2次 hash 的值，只要保证相同 cookie 的 hash 值一致即可。
	CookieHash *string `json:"CookieHash,omitempty" name:"CookieHash"`

	// 地址。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 登录来源：
	// 0：其他。
	// 1：PC 网页。
	// 2：移动页面。
	// 3：App。
	// 4：微信公众号。
	LoginSource *string `json:"LoginSource,omitempty" name:"LoginSource"`

	// 登录方式：
	// 0：其他。
	// 1：手动账号密码输入。
	// 2：动态短信密码登录。
	// 3：二维码扫描登录。
	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`

	// 登录耗时，单位：秒。
	LoginSpend *string `json:"LoginSpend,omitempty" name:"LoginSpend"`

	// 用户操作的目的 ID，如点赞等，该字段就是被点赞的消息 ID，如果是投票，则为被投号码的 ID。
	RootId *string `json:"RootId,omitempty" name:"RootId"`

	// 用户 HTTP 请求的 referer 值。
	Referer *string `json:"Referer,omitempty" name:"Referer"`

	// 登录成功后跳转页面。
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 用户 HTTP 请求的 userAgent。
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// 用户 HTTP 请求中的 x_forward_for。
	XForwardedFor *string `json:"XForwardedFor,omitempty" name:"XForwardedFor"`

	// 用户操作过程中鼠标单击次数。
	MouseClickCount *string `json:"MouseClickCount,omitempty" name:"MouseClickCount"`

	// 用户操作过程中键盘单击次数。
	KeyboardClickCount *string `json:"KeyboardClickCount,omitempty" name:"KeyboardClickCount"`

	// MAC 地址或设备唯一标识。
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机制造商 ID，如果手机注册，请带上此信息。
	VendorId *string `json:"VendorId,omitempty" name:"VendorId"`

	// 手机设备号。支持以下格式：
	// 1.imei明文
	// 2.idfa明文,
	// 3.imei小写后MD5值小写
	// 4.idfa大写后MD5值小写
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// App 客户端版本。
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 业务 ID 网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 1：微信公众号。
	// 2：微信小程序。
	WxSubType *string `json:"WxSubType,omitempty" name:"WxSubType"`

	// Token 签名随机数，WxSubType为微信小程序时必填，建议16个字符。
	RandNum *string `json:"RandNum,omitempty" name:"RandNum"`

	// 如果 accountType为2而且wxSubType有填，该字段必选，否则不需要填写；
	// 如果是微信小程序（WxSubType=2），该字段为以ssesion_key为key去签名随机数radnNum得到的值（ hmac_sha256签名算法）；如果是微信公众号或第三方登录，则为授权的access_token（网页版本的access_Token,而且获取token的scope字段必需填写snsapi_userinfo；）
	WxToken *string `json:"WxToken,omitempty" name:"WxToken"`

	// 是否识别设备异常：
	// 0：不识别。
	// 1：识别。
	CheckDevice *string `json:"CheckDevice,omitempty" name:"CheckDevice"`
}

func (r *QueryActivityAntiRushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryActivityAntiRushRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryActivityAntiRushResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作时间戳，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PostTime *string `json:"PostTime,omitempty" name:"PostTime"`

		// 用户操作的真实外网 IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

		// 0：表示无恶意。
	// 1 - 4：恶意等级由低到高。
		Level *int64 `json:"Level,omitempty" name:"Level"`

		// 风险类型。
	// 
	// 账号风险：
	// 
	// 1，账号信用低，账号近期存在因恶意被处罚历史，网络低活跃，被举报等因素；
	// 2，垃圾账号，疑似批量注册小号，近期存在严重违规或大量举报；
	// 3，无效账号，送检账号参数无法成功解析，请检查微信openid是否有误 ，QQopenid是否与appidU对应，手机号是否有误。
	// 4，黑名单，该账号在业务侧有过拉黑记录
	// 5，白名单，该账号在业务侧有过加白名单记录
	// 
	// 行为风险：
	// 101，批量操作，存在ip/设备/环境等因素的聚集性异常；
	// 102，自动机，疑似自动机批量请求；
	// 104，微信登录态无效，检查wxToken参数，是否已经失效；
	// 
	// 环境风险：
	// 201，环境异常，操作ip/设备/环境存在异常。当前ip为非常用ip或恶意ip段；
	// 205，非公网有效ip，传进来的IP地址为内网ip地址或者ip保留地址；
	// 206，设备异常，该设备存在异常的使用行为
		RiskType []*int64 `json:"RiskType,omitempty" name:"RiskType" list`

		// accountType是QQ或微信开放账号时，用于标识QQ或微信用户登录后关联业务自身的账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		AssociateAccount *string `json:"AssociateAccount,omitempty" name:"AssociateAccount"`

		// 用户ID 
	// accountType不同对应不同的用户ID。如果是QQ或微信用户则填入对应的openId
	// 注意：此字段可能返回 null，表示取不到有效值。
		Uid *string `json:"Uid,omitempty" name:"Uid"`

		// 用户操作的目的ID 
	// 比如：点赞，该字段就是被点 赞的消息 id，如果是投票，就是被投号码的 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		RootId *string `json:"RootId,omitempty" name:"RootId"`

		// 业务侧错误码。成功时返回Success，错误时返回具体业务错误原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryActivityAntiRushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryActivityAntiRushResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
