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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type QueryRegisterProtectionRequestParams struct {
	// 注册来源的外网 IP。
	RegisterIp *string `json:"RegisterIp,omitempty" name:"RegisterIp"`

	// 用户 ID 不同的 accountType 对应不同的用户 ID。如果是 QQ，则填入对应的 openid，微信用户则填入对应的 openid/unionid，手机号则填入对应真实用户手机号（如13123456789）。
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 注册时间戳，单位：秒。
	RegisterTime *string `json:"RegisterTime,omitempty" name:"RegisterTime"`

	// 用户账号类型（QQ 开放帐号、微信开放账号需要 提交工单 由腾讯云进行资格审核）：
	// 1：QQ 开放帐号。
	// 2：微信开放账号。
	// 4：手机号。
	// 0：其他。
	// 10004：手机号 MD5。
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// accountType 是 QQ 或微信开放账号时，该参数必填，表示 QQ 或微信分配给网站或应用的 AppID，用来唯一标识网站或应用。
	AppIdU *string `json:"AppIdU,omitempty" name:"AppIdU"`

	// accountType 是 QQ 或微信开放账号时，用于标识 QQ 或微信用户登录后关联业务自身的账号 ID。
	AssociateAccount *string `json:"AssociateAccount,omitempty" name:"AssociateAccount"`

	// 昵称，UTF-8 编码。
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 手机号：国家代码-手机号， 如0086-15912345687（0086前不需要+号）。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 用户邮箱地址（非系统自动生成）。
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 地址。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 用户 HTTP 请求中的 cookie 进行2次 hash 的值，只要保证相同 cookie 的 hash 值一致即可。
	CookieHash *string `json:"CookieHash,omitempty" name:"CookieHash"`

	// 注册来源：
	// 0：其他。
	// 1：PC 网页。
	// 2：移动页面。
	// 3：App。
	// 4：微信公众号。
	RegisterSource *string `json:"RegisterSource,omitempty" name:"RegisterSource"`

	// 用户 HTTP 请求的 referer 值。
	Referer *string `json:"Referer,omitempty" name:"Referer"`

	// 注册成功后跳转页面。
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 用户 HTTP 请求的 userAgent。
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// 用户 HTTP 请求中的 x_forward_for。
	XForwardedFor *string `json:"XForwardedFor,omitempty" name:"XForwardedFor"`

	// 用户操作过程中鼠标单击次数。
	MouseClickCount *string `json:"MouseClickCount,omitempty" name:"MouseClickCount"`

	// 用户操作过程中键盘单击次数。
	KeyboardClickCount *string `json:"KeyboardClickCount,omitempty" name:"KeyboardClickCount"`

	// 注册结果：
	// 0：失败。
	// 1：成功。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 失败原因：
	// 0：其他。
	// 1：参数错误。
	// 2：帐号冲突。
	// 3：验证错误。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 登录耗时，单位：秒。
	RegisterSpend *string `json:"RegisterSpend,omitempty" name:"RegisterSpend"`

	// MAC 地址或设备唯一标识。
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机制造商 ID，如果手机注册，请带上此信息。
	VendorId *string `json:"VendorId,omitempty" name:"VendorId"`

	// App 客户端版本。
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 手机设备号。
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 业务 ID 网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 1：微信公众号。
	// 2：微信小程序。
	WxSubType *string `json:"WxSubType,omitempty" name:"WxSubType"`

	// Token 签名随机数，微信小程序必填，建议16个字符。
	RandNum *string `json:"RandNum,omitempty" name:"RandNum"`

	// 如果是微信小程序，该字段为以 ssesion_key 为 key 去签名随机数 radnNum 得到的值（hmac_sha256签名算法）。
	// 如果是微信公众号或第三方登录，则为授权的 access_token（注意：不是普通 access_token，具体看 微信官方文档）。
	WxToken *string `json:"WxToken,omitempty" name:"WxToken"`
}

type QueryRegisterProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 注册来源的外网 IP。
	RegisterIp *string `json:"RegisterIp,omitempty" name:"RegisterIp"`

	// 用户 ID 不同的 accountType 对应不同的用户 ID。如果是 QQ，则填入对应的 openid，微信用户则填入对应的 openid/unionid，手机号则填入对应真实用户手机号（如13123456789）。
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 注册时间戳，单位：秒。
	RegisterTime *string `json:"RegisterTime,omitempty" name:"RegisterTime"`

	// 用户账号类型（QQ 开放帐号、微信开放账号需要 提交工单 由腾讯云进行资格审核）：
	// 1：QQ 开放帐号。
	// 2：微信开放账号。
	// 4：手机号。
	// 0：其他。
	// 10004：手机号 MD5。
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// accountType 是 QQ 或微信开放账号时，该参数必填，表示 QQ 或微信分配给网站或应用的 AppID，用来唯一标识网站或应用。
	AppIdU *string `json:"AppIdU,omitempty" name:"AppIdU"`

	// accountType 是 QQ 或微信开放账号时，用于标识 QQ 或微信用户登录后关联业务自身的账号 ID。
	AssociateAccount *string `json:"AssociateAccount,omitempty" name:"AssociateAccount"`

	// 昵称，UTF-8 编码。
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 手机号：国家代码-手机号， 如0086-15912345687（0086前不需要+号）。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 用户邮箱地址（非系统自动生成）。
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 地址。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 用户 HTTP 请求中的 cookie 进行2次 hash 的值，只要保证相同 cookie 的 hash 值一致即可。
	CookieHash *string `json:"CookieHash,omitempty" name:"CookieHash"`

	// 注册来源：
	// 0：其他。
	// 1：PC 网页。
	// 2：移动页面。
	// 3：App。
	// 4：微信公众号。
	RegisterSource *string `json:"RegisterSource,omitempty" name:"RegisterSource"`

	// 用户 HTTP 请求的 referer 值。
	Referer *string `json:"Referer,omitempty" name:"Referer"`

	// 注册成功后跳转页面。
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 用户 HTTP 请求的 userAgent。
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// 用户 HTTP 请求中的 x_forward_for。
	XForwardedFor *string `json:"XForwardedFor,omitempty" name:"XForwardedFor"`

	// 用户操作过程中鼠标单击次数。
	MouseClickCount *string `json:"MouseClickCount,omitempty" name:"MouseClickCount"`

	// 用户操作过程中键盘单击次数。
	KeyboardClickCount *string `json:"KeyboardClickCount,omitempty" name:"KeyboardClickCount"`

	// 注册结果：
	// 0：失败。
	// 1：成功。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 失败原因：
	// 0：其他。
	// 1：参数错误。
	// 2：帐号冲突。
	// 3：验证错误。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 登录耗时，单位：秒。
	RegisterSpend *string `json:"RegisterSpend,omitempty" name:"RegisterSpend"`

	// MAC 地址或设备唯一标识。
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机制造商 ID，如果手机注册，请带上此信息。
	VendorId *string `json:"VendorId,omitempty" name:"VendorId"`

	// App 客户端版本。
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 手机设备号。
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 业务 ID 网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 1：微信公众号。
	// 2：微信小程序。
	WxSubType *string `json:"WxSubType,omitempty" name:"WxSubType"`

	// Token 签名随机数，微信小程序必填，建议16个字符。
	RandNum *string `json:"RandNum,omitempty" name:"RandNum"`

	// 如果是微信小程序，该字段为以 ssesion_key 为 key 去签名随机数 radnNum 得到的值（hmac_sha256签名算法）。
	// 如果是微信公众号或第三方登录，则为授权的 access_token（注意：不是普通 access_token，具体看 微信官方文档）。
	WxToken *string `json:"WxToken,omitempty" name:"WxToken"`
}

func (r *QueryRegisterProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRegisterProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegisterIp")
	delete(f, "Uid")
	delete(f, "RegisterTime")
	delete(f, "AccountType")
	delete(f, "AppIdU")
	delete(f, "AssociateAccount")
	delete(f, "NickName")
	delete(f, "PhoneNumber")
	delete(f, "EmailAddress")
	delete(f, "Address")
	delete(f, "CookieHash")
	delete(f, "RegisterSource")
	delete(f, "Referer")
	delete(f, "JumpUrl")
	delete(f, "UserAgent")
	delete(f, "XForwardedFor")
	delete(f, "MouseClickCount")
	delete(f, "KeyboardClickCount")
	delete(f, "Result")
	delete(f, "Reason")
	delete(f, "RegisterSpend")
	delete(f, "MacAddress")
	delete(f, "VendorId")
	delete(f, "AppVersion")
	delete(f, "Imei")
	delete(f, "BusinessId")
	delete(f, "WxSubType")
	delete(f, "RandNum")
	delete(f, "WxToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryRegisterProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRegisterProtectionResponseParams struct {
	// 业务侧错误码，成功时返回 Success，错误时返回具体业务错误原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`

	// accountType 是 QQ 或微信开放账号时，用于标识 QQ 或微信用户登录后关联业务自身的账号 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateAccount *string `json:"AssociateAccount,omitempty" name:"AssociateAccount"`

	// 注册时间戳，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterTime *string `json:"RegisterTime,omitempty" name:"RegisterTime"`

	// 用户 ID 不同的 accountType 对应不同的用户 ID。如果是 QQ，则填入对应的 openid，微信用户则填入对应的 openid/unionid，手机号则填入对应真实用户手机号（如13123456789）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 注册来源的外网 IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterIp *string `json:"RegisterIp,omitempty" name:"RegisterIp"`

	// 0：表示无恶意。
	// 1 - 4：恶意等级由低到高。
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 风险类型。
	RiskType []*int64 `json:"RiskType,omitempty" name:"RiskType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryRegisterProtectionResponse struct {
	*tchttp.BaseResponse
	Response *QueryRegisterProtectionResponseParams `json:"Response"`
}

func (r *QueryRegisterProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRegisterProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}