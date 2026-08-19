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

package v20260130

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AddPromotionEvent struct {
	// <p>营销活动ID</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>营销活动名称</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>营销活动描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>邀请人ID</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>营销活动关联的优惠券</p>
	Coupon *Coupon `json:"Coupon,omitnil,omitempty" name:"Coupon"`

	// <p>营销活动关联的积分活动</p>
	Point *CreditPoint `json:"Point,omitnil,omitempty" name:"Point"`

	// <p>参加营销活动结果</p>
	Result *Result `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Address struct {
	// <p>国家</p><p>参数格式：符合ISO 3166标准</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>省份</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>城市</p>
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// <p>地区</p>
	District *string `json:"District,omitnil,omitempty" name:"District"`

	// <p>详细地址</p>
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// <p>邮政编码</p>
	ZipCode *string `json:"ZipCode,omitnil,omitempty" name:"ZipCode"`
}

type Amount struct {
	// <p>原始货币类型</p><p>参数格式：符合ISO 4217标准</p>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// <p>原始金额</p>
	OriginalAmount *float64 `json:"OriginalAmount,omitnil,omitempty" name:"OriginalAmount"`

	// <p>当前币种对美金的汇率</p>
	ExchangeRateUSD *float64 `json:"ExchangeRateUSD,omitnil,omitempty" name:"ExchangeRateUSD"`

	// <p>当前币种对人民币的汇率</p>
	ExchangeRateCNY *float64 `json:"ExchangeRateCNY,omitnil,omitempty" name:"ExchangeRateCNY"`
}

type App struct {
	// <p>应用程序运行的移动设备的操作系统类型</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>应用程序运行的移动设备的操作系统版本</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`

	// <p>应用程序运行的移动设备的生产厂商</p>
	DeviceManufacturer *string `json:"DeviceManufacturer,omitnil,omitempty" name:"DeviceManufacturer"`

	// <p>应用程序运行的移动设备的型号</p>
	DeviceModel *string `json:"DeviceModel,omitnil,omitempty" name:"DeviceModel"`

	// <p>应用程序运行的移动设备的唯一ID，对于iOS为IFV标识，对于Android为Android ID</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>应用程序名称</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>应用程序版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>应用程序提供的语言</p>
	ClientLanguage *string `json:"ClientLanguage,omitnil,omitempty" name:"ClientLanguage"`
}

// Predefined struct for user
type AssessDeviceRiskPremiumProRequestParams struct {
	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type AssessDeviceRiskPremiumProRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *AssessDeviceRiskPremiumProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskPremiumProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessDeviceRiskPremiumProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessDeviceRiskPremiumProResponseParams struct {
	// <p>设备风险评估高级版返回结果</p>
	Data *AssessDeviceRiskPremiumRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessDeviceRiskPremiumProResponse struct {
	*tchttp.BaseResponse
	Response *AssessDeviceRiskPremiumProResponseParams `json:"Response"`
}

func (r *AssessDeviceRiskPremiumProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskPremiumProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessDeviceRiskPremiumRsp struct {
	// <p>决策信息</p>
	Decision *Decision `json:"Decision,omitnil,omitempty" name:"Decision"`

	// <p>设备风险分信息</p>
	Score *DataScore `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>设备基础信息</p>
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>IP环境基础信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environment *Environment `json:"Environment,omitnil,omitempty" name:"Environment"`
}

// Predefined struct for user
type AssessDeviceRiskProRequestParams struct {
	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type AssessDeviceRiskProRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *AssessDeviceRiskProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessDeviceRiskProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessDeviceRiskProResponseParams struct {
	// <p>设备风险评估基础版返回结果</p>
	Data *AssessDeviceRiskRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessDeviceRiskProResponse struct {
	*tchttp.BaseResponse
	Response *AssessDeviceRiskProResponseParams `json:"Response"`
}

func (r *AssessDeviceRiskProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessDeviceRiskRsp struct {
	// <p>设备风险分信息</p>
	Score *DataScore `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>设备基础信息</p>
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>IP环境基础信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environment *Environment `json:"Environment,omitnil,omitempty" name:"Environment"`
}

// Predefined struct for user
type AssessEnvironmentRiskRequestParams struct {
	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type AssessEnvironmentRiskRequest struct {
	*tchttp.BaseRequest
	
	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *AssessEnvironmentRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessEnvironmentRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessEnvironmentRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessEnvironmentRiskResponseParams struct {
	// <p>环境风险评估返回结果</p>
	Data *AssessEnvironmentRiskRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessEnvironmentRiskResponse struct {
	*tchttp.BaseResponse
	Response *AssessEnvironmentRiskResponseParams `json:"Response"`
}

func (r *AssessEnvironmentRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessEnvironmentRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessEnvironmentRiskRsp struct {
	// <p>IP环境风险分信息</p>
	Score *DataScore `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>IP环境基础信息</p>
	Environment *Environment `json:"Environment,omitnil,omitempty" name:"Environment"`
}

// Predefined struct for user
type AssessRiskRequestParams struct {
	// <p>事件码，标准事件包含：</p><p>枚举值：</p><ul><li>login： 登录</li><li>register： 注册</li><li>create_order： 创建订单</li><li>transaction： 交易支付</li><li>charge_back： 拒付</li><li>sms： 短信</li><li>logout： 登出</li><li>modify_account： 修改账号</li><li>modify_password： 修改密码</li><li>security_verification： 安全验证</li><li>add_promotion： 参加营销活动</li><li>redeem： 兑奖</li><li>withdraw： 提现</li><li>cust_event： 自定义事件，cust_xxx</li><li>scan_code： 扫码</li><li>lucky_draw： 抽奖</li><li>task： 做任务</li><li>invitation： 邀请</li><li>claim_red_packet： 领红包</li><li>browse： 浏览</li></ul><p>自定义事件可与RCE约定后进行风险评估</p>
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// <p>事件的发生时间</p><p>参数格式：符合ISO 8601标准的带UTC时区的毫秒级时间</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>用户当前会话 ID， 用于关联用户登录前后的动作，如果没有传UserId，则SessionId必传，如缺失则可填充空字符串</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>事件详情，根据您输入的事件码传入对应的事件信息</p>
	EventDetail *EventDetail `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`

	// <p>用户在您系统中的唯一ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>用户邮箱</p>
	UserEmail *string `json:"UserEmail,omitnil,omitempty" name:"UserEmail"`

	// <p>用户提供的联系方式</p><p>参数格式：符合E.164标准的带“+”、地区编码和号码的格式</p>
	UserPhone *string `json:"UserPhone,omitnil,omitempty" name:"UserPhone"`

	// <p>web浏览器相关信息，若您已集成我们的设备指纹SDK，则无需传入此字段</p>
	Browser *Browser `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>应用程序、操作系统和移动设备详细信息，若您已集成我们的设备指纹SDK，则无需传入此字段</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`

	// <p>数据授权信息，国内地域必填</p>
	DataAuthorization *DataAuthorization `json:"DataAuthorization,omitnil,omitempty" name:"DataAuthorization"`

	// <p>手机号码加密方式，国内地域必填</p><p>枚举值：</p><ul><li>md5： md5加密</li><li>plain： 明文</li></ul>
	UserPhoneEncrypt *string `json:"UserPhoneEncrypt,omitnil,omitempty" name:"UserPhoneEncrypt"`

	// <p>微信开放账号</p>
	WeChatOpenId *string `json:"WeChatOpenId,omitnil,omitempty" name:"WeChatOpenId"`

	// <p>QQ开放账号</p>
	QQOpenId *string `json:"QQOpenId,omitnil,omitempty" name:"QQOpenId"`

	// <p>QQ应用ID，当传入QQ开放账号时，该字段必填，QQ分配给网站或应用的AppId，用来唯一标识网站或应用</p>
	QQAppId *string `json:"QQAppId,omitnil,omitempty" name:"QQAppId"`
}

type AssessRiskRequest struct {
	*tchttp.BaseRequest
	
	// <p>事件码，标准事件包含：</p><p>枚举值：</p><ul><li>login： 登录</li><li>register： 注册</li><li>create_order： 创建订单</li><li>transaction： 交易支付</li><li>charge_back： 拒付</li><li>sms： 短信</li><li>logout： 登出</li><li>modify_account： 修改账号</li><li>modify_password： 修改密码</li><li>security_verification： 安全验证</li><li>add_promotion： 参加营销活动</li><li>redeem： 兑奖</li><li>withdraw： 提现</li><li>cust_event： 自定义事件，cust_xxx</li><li>scan_code： 扫码</li><li>lucky_draw： 抽奖</li><li>task： 做任务</li><li>invitation： 邀请</li><li>claim_red_packet： 领红包</li><li>browse： 浏览</li></ul><p>自定义事件可与RCE约定后进行风险评估</p>
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// <p>事件的发生时间</p><p>参数格式：符合ISO 8601标准的带UTC时区的毫秒级时间</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>用户当前会话 ID， 用于关联用户登录前后的动作，如果没有传UserId，则SessionId必传，如缺失则可填充空字符串</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>事件详情，根据您输入的事件码传入对应的事件信息</p>
	EventDetail *EventDetail `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`

	// <p>用户在您系统中的唯一ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>用户邮箱</p>
	UserEmail *string `json:"UserEmail,omitnil,omitempty" name:"UserEmail"`

	// <p>用户提供的联系方式</p><p>参数格式：符合E.164标准的带“+”、地区编码和号码的格式</p>
	UserPhone *string `json:"UserPhone,omitnil,omitempty" name:"UserPhone"`

	// <p>web浏览器相关信息，若您已集成我们的设备指纹SDK，则无需传入此字段</p>
	Browser *Browser `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>应用程序、操作系统和移动设备详细信息，若您已集成我们的设备指纹SDK，则无需传入此字段</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`

	// <p>数据授权信息，国内地域必填</p>
	DataAuthorization *DataAuthorization `json:"DataAuthorization,omitnil,omitempty" name:"DataAuthorization"`

	// <p>手机号码加密方式，国内地域必填</p><p>枚举值：</p><ul><li>md5： md5加密</li><li>plain： 明文</li></ul>
	UserPhoneEncrypt *string `json:"UserPhoneEncrypt,omitnil,omitempty" name:"UserPhoneEncrypt"`

	// <p>微信开放账号</p>
	WeChatOpenId *string `json:"WeChatOpenId,omitnil,omitempty" name:"WeChatOpenId"`

	// <p>QQ开放账号</p>
	QQOpenId *string `json:"QQOpenId,omitnil,omitempty" name:"QQOpenId"`

	// <p>QQ应用ID，当传入QQ开放账号时，该字段必填，QQ分配给网站或应用的AppId，用来唯一标识网站或应用</p>
	QQAppId *string `json:"QQAppId,omitnil,omitempty" name:"QQAppId"`
}

func (r *AssessRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventCode")
	delete(f, "EventTime")
	delete(f, "SessionId")
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	delete(f, "EventDetail")
	delete(f, "UserId")
	delete(f, "UserEmail")
	delete(f, "UserPhone")
	delete(f, "Browser")
	delete(f, "App")
	delete(f, "DataAuthorization")
	delete(f, "UserPhoneEncrypt")
	delete(f, "WeChatOpenId")
	delete(f, "QQOpenId")
	delete(f, "QQAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessRiskResponseParams struct {
	// <p>事件风险评估结果</p>
	Data *AssessRiskRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessRiskResponse struct {
	*tchttp.BaseResponse
	Response *AssessRiskResponseParams `json:"Response"`
}

func (r *AssessRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessRiskRsp struct {
	// <p>决策信息</p>
	Decision *Decision `json:"Decision,omitnil,omitempty" name:"Decision"`

	// <p>风险分，根据您开启的产品服务计算的评分结果</p>
	Score *Score `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>扩展信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo []*Cust `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`
}

type Billing struct {
	// <p>账单地址</p>
	Address *Address `json:"Address,omitnil,omitempty" name:"Address"`

	// <p>账单联系电话</p><p>参数格式：符合E.164标准的带“+”、地区编码和号码的格式</p>
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// <p>账单邮箱</p>
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// <p>账单接收人姓名</p>
	Recipient *string `json:"Recipient,omitnil,omitempty" name:"Recipient"`
}

type BrowseEvent struct {
	// <p>当前浏览网页的类型，例如主页、搜索页等</p>
	PageType *string `json:"PageType,omitnil,omitempty" name:"PageType"`

	// <p>当前浏览的网页URL</p>
	PageUrl *string `json:"PageUrl,omitnil,omitempty" name:"PageUrl"`

	// <p>浏览耗时</p><p>单位：毫秒</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>网页内容类型，例如广告、视频、文章等</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>网页内容ID</p>
	ContentId *string `json:"ContentId,omitnil,omitempty" name:"ContentId"`

	// <p>上一个网页的类型，例如主页、搜索页等</p>
	ReferPageType *string `json:"ReferPageType,omitnil,omitempty" name:"ReferPageType"`

	// <p>上一个网页URL</p>
	ReferPageUrl *string `json:"ReferPageUrl,omitnil,omitempty" name:"ReferPageUrl"`

	// <p>游客账号ID</p>
	GuestId *string `json:"GuestId,omitnil,omitempty" name:"GuestId"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Browser struct {
	// <p>与网站交互的浏览器的用户代理</p>
	UserAgent *string `json:"UserAgent,omitnil,omitempty" name:"UserAgent"`

	// <p>浏览器支持的用户请求语言</p><p>参数格式：符合ISO 3166标准</p>
	AcceptLanguage *string `json:"AcceptLanguage,omitnil,omitempty" name:"AcceptLanguage"`

	// <p>浏览器当前网站内容的语言</p><p>参数格式：符合ISO 3166标准</p>
	ContentLanguage *string `json:"ContentLanguage,omitnil,omitempty" name:"ContentLanguage"`
}

type Card struct {
	// <p>发卡行识别码卡号前6位</p><p>参数格式：符合ISO 13616-1标准</p>
	CardBin *string `json:"CardBin,omitnil,omitempty" name:"CardBin"`

	// <p>发卡行识别码卡号后4位</p><p>参数格式：符合ISO 13616-1标准</p>
	LastFourDigits *string `json:"LastFourDigits,omitnil,omitempty" name:"LastFourDigits"`

	// <p>发行国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>发行银行</p>
	Bank *string `json:"Bank,omitnil,omitempty" name:"Bank"`

	// <p>支付卡类型</p><p>枚举值：</p><ul><li>credit： 信用卡</li><li>debit： 借记卡</li><li>charge： 签账卡</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>支付卡品牌</p>
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// <p>支付卡等级</p>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>持有者姓名</p>
	HolderName *string `json:"HolderName,omitnil,omitempty" name:"HolderName"`

	// <p>过期日期</p><p>参数格式：YYYY-MM-DD</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type ChargeBackEvent struct {
	// <p>交易ID</p>
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`

	// <p>订单 ID，当一笔交易关联多个订单（合并支付）时请输入所有订单ID</p>
	OrderId []*string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// <p>拒付理由码，参考各卡组织定义的拒付码，例如：10.1、13.1、 4870、4871等</p>
	ChargeBackCode *string `json:"ChargeBackCode,omitnil,omitempty" name:"ChargeBackCode"`

	// <p>拒付原因，参考各卡组织定义的拒付原因，例如：未收到商品、欺诈等</p>
	ChargeBackReason *string `json:"ChargeBackReason,omitnil,omitempty" name:"ChargeBackReason"`

	// <p>拒付申诉阶段</p><p>枚举值：</p><ul><li>need_response： 需要商家回应</li><li>information_supplied： 商家已提供信息</li><li>chargeback_reversed： 拒付已撤销</li><li>chargeback_sustained： 拒付已成立</li></ul>
	ChargeBackProcess *string `json:"ChargeBackProcess,omitnil,omitempty" name:"ChargeBackProcess"`

	// <p>拒付金额</p>
	ChargeBackAmount *Amount `json:"ChargeBackAmount,omitnil,omitempty" name:"ChargeBackAmount"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type ClaimRedPacketEvent struct {
	// <p>营销活动ID</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>营销活动名称</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>营销活动描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>邀请人ID</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>红包ID</p>
	RedPacketId *string `json:"RedPacketId,omitnil,omitempty" name:"RedPacketId"`

	// <p>红包类型，如手气红包、口令红包、均分红包等</p>
	RedPacketType *string `json:"RedPacketType,omitnil,omitempty" name:"RedPacketType"`

	// <p>红包金额</p>
	RedPacketAmount *Amount `json:"RedPacketAmount,omitnil,omitempty" name:"RedPacketAmount"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Coupon struct {
	// <p>优惠券ID</p>
	CouponId *string `json:"CouponId,omitnil,omitempty" name:"CouponId"`

	// <p>优惠券名称</p>
	CouponName *string `json:"CouponName,omitnil,omitempty" name:"CouponName"`

	// <p>优惠券开始时间</p><p>参数格式：符合ISO 8601标准的带UTC时区的毫秒级时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>优惠券过期时间</p><p>参数格式：符合ISO 8601标准的带UTC时区的毫秒级时间</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>折扣百分比，如果折扣为 10%，则发送“0.1”</p>
	PercentageRate *float64 `json:"PercentageRate,omitnil,omitempty" name:"PercentageRate"`

	// <p>折扣金额</p>
	DiscountAmount *Amount `json:"DiscountAmount,omitnil,omitempty" name:"DiscountAmount"`

	// <p>优惠券门槛</p>
	Threshold *float64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`
}

type CreateOrderEvent struct {
	// <p>订单ID</p>
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// <p>订单金额</p>
	Amount *Amount `json:"Amount,omitnil,omitempty" name:"Amount"`

	// <p>商家信息</p>
	Merchant *Merchant `json:"Merchant,omitnil,omitempty" name:"Merchant"`

	// <p>账单信息</p>
	Billing *Billing `json:"Billing,omitnil,omitempty" name:"Billing"`

	// <p>商品信息</p>
	Items []*Item `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>物流信息</p>
	Delivery *Delivery `json:"Delivery,omitnil,omitempty" name:"Delivery"`

	// <p>营销活动信息</p>
	Promotions []*Promotion `json:"Promotions,omitnil,omitempty" name:"Promotions"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type CreditPoint struct {
	// <p>积分分值</p>
	Point *float64 `json:"Point,omitnil,omitempty" name:"Point"`

	// <p>积分类型</p>
	PointType *string `json:"PointType,omitnil,omitempty" name:"PointType"`
}

type Cust struct {
	// <p>标识符</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type CustEvent struct {
	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type DataAuthorization struct {
	// <p>数据委托方，客户主体名称</p>
	DataProviderName *string `json:"DataProviderName,omitnil,omitempty" name:"DataProviderName"`

	// <p>数据受托方，腾讯云主体名称，固定填：腾讯云计算（北京）有限责任公司</p>
	DataRecipientName *string `json:"DataRecipientName,omitnil,omitempty" name:"DataRecipientName"`

	// <p>客户请求RCE所提供的用户数据类型，支持多选</p><p>枚举值：</p><ul><li>1： 手机号</li><li>2： 微信开放账号</li><li>3： QQ开放账号</li><li>4： IP地址</li><li>5： URL网址</li><li>999： 其他</li></ul>
	UserDataType []*int64 `json:"UserDataType,omitnil,omitempty" name:"UserDataType"`

	// <p>客户是否已按合规指南要求获取用户授权，同意客户委托腾讯云处理入参信息</p><p>枚举值：</p><ul><li>true： 已授权</li><li>false： 未授权</li></ul>
	IsAuthorized *bool `json:"IsAuthorized,omitnil,omitempty" name:"IsAuthorized"`

	// <p>客户是否已按合规指南要求获取用户授权，同意腾讯云结合客户提供的信息，对已合法收集的用户数据进行必要处理得出服务结果，并返回给客户</p><p>枚举值：</p><ul><li>true： 已授权</li><li>false： 未授权</li></ul>
	IsOrderHanding *bool `json:"IsOrderHanding,omitnil,omitempty" name:"IsOrderHanding"`

	// <p>客户获得的用户授权期限Unix时间戳（单位秒），不填默认无固定期限</p>
	AuthorizationDeadline *int64 `json:"AuthorizationDeadline,omitnil,omitempty" name:"AuthorizationDeadline"`

	// <p>客户获得用户授权所依赖的协议地址</p>
	PrivacyPolicyLink *string `json:"PrivacyPolicyLink,omitnil,omitempty" name:"PrivacyPolicyLink"`
}

type DataScore struct {
	// <p>风险等级</p>
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>风险标签</p>
	RiskLabels []*RiskLabel `json:"RiskLabels,omitnil,omitempty" name:"RiskLabels"`

	// <p>综合风险分数。</p><p>取值范围：[1, 1000]</p><p>数值越大，风险越大。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskScore *int64 `json:"RiskScore,omitnil,omitempty" name:"RiskScore"`
}

type Decision struct {
	// <p>决策结果</p><ul><li>pass：通过</li><li>review：复审</li><li>reject：拒绝</li></ul>
	DecisionResult *string `json:"DecisionResult,omitnil,omitempty" name:"DecisionResult"`

	// <p>命中策略后的决策动作，可在控制台配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Disposition *string `json:"Disposition,omitnil,omitempty" name:"Disposition"`
}

type Delivery struct {
	// <p>物流方式</p><ul><li>physical：物理投送</li><li>electonic：电子投送</li></ul>
	DeliveryMethod *string `json:"DeliveryMethod,omitnil,omitempty" name:"DeliveryMethod"`

	// <p>物流费用</p>
	DeliveryAmount *Amount `json:"DeliveryAmount,omitnil,omitempty" name:"DeliveryAmount"`

	// <p>收货地址</p>
	DeliveryAddress *Address `json:"DeliveryAddress,omitnil,omitempty" name:"DeliveryAddress"`

	// <p>收货人电话</p><p>参数格式：符合E.164标准的带“+”、地区编码和号码的格式</p>
	ConsigneePhone *string `json:"ConsigneePhone,omitnil,omitempty" name:"ConsigneePhone"`

	// <p>收货人邮箱</p>
	ConsigneeEmail *string `json:"ConsigneeEmail,omitnil,omitempty" name:"ConsigneeEmail"`

	// <p>收货人姓名</p>
	ConsigneeName *string `json:"ConsigneeName,omitnil,omitempty" name:"ConsigneeName"`

	// <p>是否加急</p>
	Expedited *bool `json:"Expedited,omitnil,omitempty" name:"Expedited"`

	// <p>物流厂商，一般是物流的公司</p>
	DeliveryCarrier *string `json:"DeliveryCarrier,omitnil,omitempty" name:"DeliveryCarrier"`

	// <p>物流追踪单号</p>
	DeliveryTracking *string `json:"DeliveryTracking,omitnil,omitempty" name:"DeliveryTracking"`
}

type Device struct {
	// <p>设备ID</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>App版本信息</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>品牌</p>
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// <p>客户端IP</p>
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// <p>机型</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>网络类型</p>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>应用包名</p>
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// <p>平台</p><p>枚举值：</p><ul><li>2： Android</li><li>3： IOS</li><li>4： H5</li><li>5： 微信小程序</li></ul>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统版本</p>
	SystemVersion *string `json:"SystemVersion,omitnil,omitempty" name:"SystemVersion"`

	// <p>SDK版本</p>
	SdkBuildVersion *string `json:"SdkBuildVersion,omitnil,omitempty" name:"SdkBuildVersion"`

	// <p>验签token，验签功能启用请联系我们。</p>
	SignToken *string `json:"SignToken,omitnil,omitempty" name:"SignToken"`

	// <p>token生成时间戳，毫秒级。</p>
	TokenTime *string `json:"TokenTime,omitnil,omitempty" name:"TokenTime"`

	// <p>隐私浏览器类型，当检测到隐私浏览器时返回，仅H5。</p>
	PrivacyBrowser *string `json:"PrivacyBrowser,omitnil,omitempty" name:"PrivacyBrowser"`
}

type DigitalOrder struct {
	// <p>数字资产</p>
	DigitalAsset *string `json:"DigitalAsset,omitnil,omitempty" name:"DigitalAsset"`

	// <p>数字资产类型</p><p>枚举值：</p><ul><li>coin： 代币</li><li>commodity： 大宗商品</li><li>crypto： 加密货币</li><li>fiat： 法币</li><li>token： 通证</li><li>stock： 股票</li><li>bond： 债券</li></ul>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>订单类型</p><p>枚举值：</p><ul><li>limit： 限价单</li><li>market： 市价单</li><li>stop_limit： 止损限价单</li><li>stop_loss： 止损单</li><li>take_profit： 止盈单</li><li>take_profit_limit： 止盈限价单</li></ul>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>数字资产的数量</p>
	Volume *float64 `json:"Volume,omitnil,omitempty" name:"Volume"`
}

type Environment struct {
	// <p>IP地理位置信息</p>
	Location *IPLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// <p>IP基础网络信息</p>
	Network *IPNetwork `json:"Network,omitnil,omitempty" name:"Network"`
}

type EventDetail struct {
	// <p>登录</p>
	Login *LoginEvent `json:"Login,omitnil,omitempty" name:"Login"`

	// <p>注册（变更用户信息）</p>
	Register *RegisterEvent `json:"Register,omitnil,omitempty" name:"Register"`

	// <p>创建订单</p>
	CreateOrder *CreateOrderEvent `json:"CreateOrder,omitnil,omitempty" name:"CreateOrder"`

	// <p>交易支付</p>
	Transaction *TransactionEvent `json:"Transaction,omitnil,omitempty" name:"Transaction"`

	// <p>短信</p>
	Sms *SMSEvent `json:"Sms,omitnil,omitempty" name:"Sms"`

	// <p>拒付</p>
	ChargeBack *ChargeBackEvent `json:"ChargeBack,omitnil,omitempty" name:"ChargeBack"`

	// <p>登出</p>
	Logout *LogoutEvent `json:"Logout,omitnil,omitempty" name:"Logout"`

	// <p>修改账号</p>
	ModifyAccount *ModifyAccountEvent `json:"ModifyAccount,omitnil,omitempty" name:"ModifyAccount"`

	// <p>修改密码</p>
	ModifyPassword *ModifyPasswordEvent `json:"ModifyPassword,omitnil,omitempty" name:"ModifyPassword"`

	// <p>安全验证</p>
	SecurityVerification *SecurityVerificationEvent `json:"SecurityVerification,omitnil,omitempty" name:"SecurityVerification"`

	// <p>参加营销活动</p>
	AddPromotion *AddPromotionEvent `json:"AddPromotion,omitnil,omitempty" name:"AddPromotion"`

	// <p>兑奖</p>
	Redeem *RedeemEvent `json:"Redeem,omitnil,omitempty" name:"Redeem"`

	// <p>提现</p>
	Withdraw *WithdrawEvent `json:"Withdraw,omitnil,omitempty" name:"Withdraw"`

	// <p>自定义事件</p>
	CustEvent *CustEvent `json:"CustEvent,omitnil,omitempty" name:"CustEvent"`

	// <p>扫码</p>
	ScanCode *ScanCodeEvent `json:"ScanCode,omitnil,omitempty" name:"ScanCode"`

	// <p>抽奖</p>
	LuckyDraw *LuckyDrawEvent `json:"LuckyDraw,omitnil,omitempty" name:"LuckyDraw"`

	// <p>做任务</p>
	Task *TaskEvent `json:"Task,omitnil,omitempty" name:"Task"`

	// <p>邀请</p>
	Invitation *InvitationEvent `json:"Invitation,omitnil,omitempty" name:"Invitation"`

	// <p>领红包</p>
	ClaimRedPacket *ClaimRedPacketEvent `json:"ClaimRedPacket,omitnil,omitempty" name:"ClaimRedPacket"`

	// <p>浏览</p>
	Browse *BrowseEvent `json:"Browse,omitnil,omitempty" name:"Browse"`
}

type IPLocation struct {
	// <p>IP地址所属国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>IP地址所属省份</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>IP地址所属城市</p>
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// <p>IP地址所属地区</p>
	District *string `json:"District,omitnil,omitempty" name:"District"`

	// <p>IP地址的经度</p>
	Longitude *string `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// <p>IP地址的纬度</p>
	Latitude *string `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// <p>IP地址所属时区</p>
	Timezone *string `json:"Timezone,omitnil,omitempty" name:"Timezone"`

	// <p>IP地址的邮政编码</p>
	ZipCode *string `json:"ZipCode,omitnil,omitempty" name:"ZipCode"`
}

type IPNetwork struct {
	// <p>互联网服务提供商</p>
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// <p>自治系统号</p>
	ASN *string `json:"ASN,omitnil,omitempty" name:"ASN"`

	// <p>IP注册组织名称</p>
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// <p>是否保留IP</p>
	IsReserved *bool `json:"IsReserved,omitnil,omitempty" name:"IsReserved"`

	// <p>是否网关IP</p>
	IsGateway *bool `json:"IsGateway,omitnil,omitempty" name:"IsGateway"`

	// <p>是否任播网络</p>
	IsAnycast *bool `json:"IsAnycast,omitnil,omitempty" name:"IsAnycast"`

	// <p>是否移动网络</p>
	IsMobile *bool `json:"IsMobile,omitnil,omitempty" name:"IsMobile"`

	// <p>是否动态IP</p>
	IsDynamic *bool `json:"IsDynamic,omitnil,omitempty" name:"IsDynamic"`

	// <p>是否网络出口</p>
	IsEgress *bool `json:"IsEgress,omitnil,omitempty" name:"IsEgress"`

	// <p>是否域名解析</p>
	IsDNS *bool `json:"IsDNS,omitnil,omitempty" name:"IsDNS"`

	// <p>是否教育机构</p>
	IsEducation *bool `json:"IsEducation,omitnil,omitempty" name:"IsEducation"`

	// <p>是否组织机构</p>
	IsInstitution *bool `json:"IsInstitution,omitnil,omitempty" name:"IsInstitution"`

	// <p>是否企业专线</p>
	IsCompany *bool `json:"IsCompany,omitnil,omitempty" name:"IsCompany"`

	// <p>是否家用宽带</p>
	IsResidence *bool `json:"IsResidence,omitnil,omitempty" name:"IsResidence"`

	// <p>是否云服务</p>
	IsCloudService *bool `json:"IsCloudService,omitnil,omitempty" name:"IsCloudService"`

	// <p>是否基础设施</p>
	IsInfrastructure *bool `json:"IsInfrastructure,omitnil,omitempty" name:"IsInfrastructure"`

	// <p>是否邮箱服务</p>
	IsMXServer *bool `json:"IsMXServer,omitnil,omitempty" name:"IsMXServer"`
}

type InvitationEvent struct {
	// <p>受邀请人ID</p>
	InviteeUserId *string `json:"InviteeUserId,omitnil,omitempty" name:"InviteeUserId"`

	// <p>营销活动ID</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>营销活动名称</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>营销活动描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>受邀请人电话号码</p><p>参数格式：符合E.164标准的带“+”、地区编码和号码的格式</p>
	InviteePhone *string `json:"InviteePhone,omitnil,omitempty" name:"InviteePhone"`

	// <p>邀请码</p>
	InvitationCode *string `json:"InvitationCode,omitnil,omitempty" name:"InvitationCode"`

	// <p>邀请链接</p>
	InvitationUrl *string `json:"InvitationUrl,omitnil,omitempty" name:"InvitationUrl"`

	// <p>邀请渠道，如微信、抖音、小红书等</p>
	InvitationChannel *string `json:"InvitationChannel,omitnil,omitempty" name:"InvitationChannel"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Inviter struct {
	// <p>邀请人ID</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>邀请人电话号码</p><p>参数格式：符合E.164标准的带“+”、地区编码和号码的格式</p>
	InviterPhone *string `json:"InviterPhone,omitnil,omitempty" name:"InviterPhone"`

	// <p>邀请码</p>
	InviteCode *string `json:"InviteCode,omitnil,omitempty" name:"InviteCode"`

	// <p>邀请渠道</p>
	InviteChannel *string `json:"InviteChannel,omitnil,omitempty" name:"InviteChannel"`
}

type Item struct {
	// <p>商品ID</p>
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// <p>商品名称</p>
	ItemName *string `json:"ItemName,omitnil,omitempty" name:"ItemName"`

	// <p>商品类别</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>商品单价</p>
	Price *Amount `json:"Price,omitnil,omitempty" name:"Price"`

	// <p>如果商品有UPC码（Universal Product Code），请提供</p>
	UPC *string `json:"UPC,omitnil,omitempty" name:"UPC"`

	// <p>如果商品有EAN码（European Article Number），请提供</p>
	EAN *string `json:"EAN,omitnil,omitempty" name:"EAN"`

	// <p>如果商品有SKU码（Stock Keeping Unit），请提供</p>
	SKU *string `json:"SKU,omitnil,omitempty" name:"SKU"`

	// <p>如果商品有ISBN码（International Standard Book Number ），请提供</p>
	ISBN *string `json:"ISBN,omitnil,omitempty" name:"ISBN"`

	// <p>商品品牌</p>
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// <p>商品数量</p>
	Quantity *int64 `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// <p>生产厂商</p>
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// <p>商品标签</p>
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type LoginEvent struct {
	// <p>用户基础信息</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>用户登录时输入的用户名</p>
	UserLoginName *string `json:"UserLoginName,omitnil,omitempty" name:"UserLoginName"`

	// <p>登录结果</p>
	LoginResult *Result `json:"LoginResult,omitnil,omitempty" name:"LoginResult"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type LogoutEvent struct {
	// <p>用户基础信息</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>用户登录时输入的用户名</p>
	UserLoginName *string `json:"UserLoginName,omitnil,omitempty" name:"UserLoginName"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type LuckyDrawEvent struct {
	// <p>营销活动ID</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>营销活动名称</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>营销活动描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>邀请人ID</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>抽奖次数</p><p>单位：次数</p>
	LuckyDrawCount *int64 `json:"LuckyDrawCount,omitnil,omitempty" name:"LuckyDrawCount"`

	// <p>抽奖类型</p>
	LuckyDrawType *string `json:"LuckyDrawType,omitnil,omitempty" name:"LuckyDrawType"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Merchant struct {
	// <p>商家ID</p>
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// <p>商家名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>商家的注册时间</p><p>参数格式：符合ISO 8601标准的带UTC时区的毫秒级时间</p>
	RegisterTime *string `json:"RegisterTime,omitnil,omitempty" name:"RegisterTime"`

	// <p>商家类别代码</p><p>参数格式：符合ISO 18245标准的4位编号</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>商家电话</p><p>参数格式：符合E.164标准的带“+”、地区编码和号码的格式</p>
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// <p>商家邮件</p>
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// <p>商家店铺网址</p>
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// <p>商家地址</p>
	Address *Address `json:"Address,omitnil,omitempty" name:"Address"`

	// <p>商家等级</p>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>经营类型</p><p>枚举值：</p><ul><li>person： 个人</li><li>company： 企业</li></ul>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// <p>商家在售商品数量</p>
	GoodsQuantity *int64 `json:"GoodsQuantity,omitnil,omitempty" name:"GoodsQuantity"`

	// <p>商家历史销售数量</p>
	HistoricSalesQuantity *int64 `json:"HistoricSalesQuantity,omitnil,omitempty" name:"HistoricSalesQuantity"`

	// <p>商家历史销售总额</p>
	HistoricSalesAmount *Amount `json:"HistoricSalesAmount,omitnil,omitempty" name:"HistoricSalesAmount"`
}

type ModifyAccountEvent struct {
	// <p>用户基础信息</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>用户填写的个人信息</p>
	Person *Person `json:"Person,omitnil,omitempty" name:"Person"`

	// <p>用户填写的账单地址</p>
	BillingAddress *Address `json:"BillingAddress,omitnil,omitempty" name:"BillingAddress"`

	// <p>用户填写的收货地址</p>
	DeliveryAddress *Address `json:"DeliveryAddress,omitnil,omitempty" name:"DeliveryAddress"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type ModifyPasswordEvent struct {
	// <p>修改原因</p><p>枚举值：</p><ul><li>user_modify： 用户主动修改</li><li>forgot_password： 忘记密码</li><li>forced_reset： 系统强制重置</li></ul>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Order struct {
	// <p>订单ID</p>
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// <p>订单金额</p>
	Amount *Amount `json:"Amount,omitnil,omitempty" name:"Amount"`

	// <p>商品信息</p>
	Items []*Item `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>物流信息</p>
	Delivery *Delivery `json:"Delivery,omitnil,omitempty" name:"Delivery"`
}

type PaymentMethod struct {
	// <p>支付方式</p><p>枚举值：</p><ul><li>cash： 现金</li><li>check： 支票</li><li>credit_card： 信用卡</li><li>debit_card： 借记卡</li><li>crypto_currency： 加密货币</li><li>digital_wallet： 数字钱包</li><li>gift_card： 礼品卡</li><li>points： 积分</li><li>in_app_purchase： APP内购买</li><li>electronic_fund_transfer： 电子资金转账</li><li>financing： 融资</li><li>invoice： 发票</li><li>prepaid_card： 预付卡</li><li>sepa_credit： SEPA信用转账</li></ul>
	PaymentType *string `json:"PaymentType,omitnil,omitempty" name:"PaymentType"`

	// <p>支付渠道</p>
	PaymentChannel *string `json:"PaymentChannel,omitnil,omitempty" name:"PaymentChannel"`

	// <p>银行卡信息，当用支付方式是credit_card、debit_card时必填</p>
	Card *Card `json:"Card,omitnil,omitempty" name:"Card"`

	// <p>SEPA直接借记授权</p><p>枚举值：</p><ul><li>true： 是</li><li>false： 否</li></ul>
	SEPADirectDebitMandate *bool `json:"SEPADirectDebitMandate,omitnil,omitempty" name:"SEPADirectDebitMandate"`

	// <p>数字钱包</p>
	DigitalWallet *Wallet `json:"DigitalWallet,omitnil,omitempty" name:"DigitalWallet"`
}

type PaymentResult struct {
	// <p>支付状态</p><p>枚举值：</p><ul><li>success： 成功</li><li>failure： 失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>支付失败原因</p>
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// <p>是否使用3DS，枚举值：</p><ul><li>是：true</li><li>否：false</li></ul>
	ThreeDomainSecure *bool `json:"ThreeDomainSecure,omitnil,omitempty" name:"ThreeDomainSecure"`

	// <p>ECI返回码</p>
	ECICode *string `json:"ECICode,omitnil,omitempty" name:"ECICode"`

	// <p>AVS响应结果（地址验证）</p>
	AVSCode *string `json:"AVSCode,omitnil,omitempty" name:"AVSCode"`

	// <p>CVC验证结果（交易真实性验证）</p>
	CVCCode *string `json:"CVCCode,omitnil,omitempty" name:"CVCCode"`
}

type Person struct {
	// <p>姓名全称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>性别</p>
	Gender *string `json:"Gender,omitnil,omitempty" name:"Gender"`

	// <p>出生日期</p><p>参数格式：YYYY-MM-DD</p>
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// <p>学历</p>
	Degree *string `json:"Degree,omitnil,omitempty" name:"Degree"`

	// <p>职业</p>
	Occupation *string `json:"Occupation,omitnil,omitempty" name:"Occupation"`
}

type Promotion struct {
	// <p>营销活动ID</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>营销活动名称</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>营销活动描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>邀请人ID</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>优惠券</p>
	Coupon *Coupon `json:"Coupon,omitnil,omitempty" name:"Coupon"`

	// <p>积分</p>
	CreditPoint *CreditPoint `json:"CreditPoint,omitnil,omitempty" name:"CreditPoint"`
}

type PromotionCode struct {
	// <p>活动码ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>活动码类型，例如：qrcode-二维码、barcode-条形码、miniprogram_code-小程序码</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>活动码图片URL或链接</p>
	ImageLink *string `json:"ImageLink,omitnil,omitempty" name:"ImageLink"`

	// <p>营销活动码使用地址</p>
	Address *Address `json:"Address,omitnil,omitempty" name:"Address"`

	// <p>营销活动码关联的商品</p>
	Items []*Item `json:"Items,omitnil,omitempty" name:"Items"`
}

type RedeemEvent struct {
	// <p>营销活动ID</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>营销活动名称</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>营销活动描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>邀请人ID</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>兑奖关联的订单信息</p>
	Order *Order `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>兑奖结果</p>
	Result *Result `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type RegisterEvent struct {
	// <p>注册结果</p>
	RegisterResult *Result `json:"RegisterResult,omitnil,omitempty" name:"RegisterResult"`

	// <p>用户基础信息</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>用户注册时填写的个人信息</p>
	Person *Person `json:"Person,omitnil,omitempty" name:"Person"`

	// <p>用户注册时填写的账单地址</p>
	BillingAddress *Address `json:"BillingAddress,omitnil,omitempty" name:"BillingAddress"`

	// <p>用户注册时填写的收货地址</p>
	DeliveryAddress *Address `json:"DeliveryAddress,omitnil,omitempty" name:"DeliveryAddress"`

	// <p>邀请人信息</p>
	Inviter *Inviter `json:"Inviter,omitnil,omitempty" name:"Inviter"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

// Predefined struct for user
type ReportEventRequestParams struct {
	// <p>事件码，标准事件包含：</p><p>枚举值：</p><ul><li>login： 登录</li><li>register： 注册</li><li>create_order： 创建订单</li><li>transaction： 交易支付</li><li>charge_back： 拒付</li><li>sms： 短信</li><li>logout： 登出</li><li>modify_account： 修改账号</li><li>modify_password： 修改密码</li><li>security_verification： 安全验证</li><li>add_promotion： 参加营销活动</li><li>redeem： 兑奖</li><li>withdraw： 提现</li><li>cust_event： 自定义事件，cust_xxx</li><li>scan_code： 扫码</li><li>lucky_draw： 抽奖</li><li>task： 做任务</li><li>invitation： 邀请</li><li>claim_red_packet： 领红包</li><li>browse： 浏览</li></ul><p>自定义事件可与RCE约定后进行风险评估</p>
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// <p>事件的发生时间</p><p>参数格式：符合ISO 8601标准的带UTC时区的毫秒级时间</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>用户当前会话 ID， 用于关联用户登录前后的动作，如果没有传UserId，则SessionId必传，如缺失则可填充空字符串</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>事件详情，根据您输入的事件码传入对应的事件信息</p>
	EventDetail *EventDetail `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`

	// <p>用户在您系统中的唯一ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>用户邮箱</p>
	UserEmail *string `json:"UserEmail,omitnil,omitempty" name:"UserEmail"`

	// <p>用户提供的联系方式</p><p>参数格式：符合E.164标准的带“+”、地区编码和号码的格式</p>
	UserPhone *string `json:"UserPhone,omitnil,omitempty" name:"UserPhone"`

	// <p>web浏览器相关信息，若您已集成我们的设备指纹SDK，则无需传入此字段</p>
	Browser *Browser `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>应用程序、操作系统和移动设备详细信息，若您已集成我们的设备指纹SDK，则无需传入此字段</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`

	// <p>数据授权信息，国内地域必填</p>
	DataAuthorization *DataAuthorization `json:"DataAuthorization,omitnil,omitempty" name:"DataAuthorization"`

	// <p>手机号码加密方式，国内地域必填</p><p>枚举值：</p><ul><li>md5： md5加密</li><li>plain： 明文</li></ul>
	UserPhoneEncrypt *string `json:"UserPhoneEncrypt,omitnil,omitempty" name:"UserPhoneEncrypt"`

	// <p>微信开放账号</p>
	WeChatOpenId *string `json:"WeChatOpenId,omitnil,omitempty" name:"WeChatOpenId"`

	// <p>QQ开放账号</p>
	QQOpenId *string `json:"QQOpenId,omitnil,omitempty" name:"QQOpenId"`

	// <p>QQ应用ID，当传入QQ开放账号时，该字段必填，QQ分配给网站或应用的AppId，用来唯一标识网站或应用</p>
	QQAppId *string `json:"QQAppId,omitnil,omitempty" name:"QQAppId"`
}

type ReportEventRequest struct {
	*tchttp.BaseRequest
	
	// <p>事件码，标准事件包含：</p><p>枚举值：</p><ul><li>login： 登录</li><li>register： 注册</li><li>create_order： 创建订单</li><li>transaction： 交易支付</li><li>charge_back： 拒付</li><li>sms： 短信</li><li>logout： 登出</li><li>modify_account： 修改账号</li><li>modify_password： 修改密码</li><li>security_verification： 安全验证</li><li>add_promotion： 参加营销活动</li><li>redeem： 兑奖</li><li>withdraw： 提现</li><li>cust_event： 自定义事件，cust_xxx</li><li>scan_code： 扫码</li><li>lucky_draw： 抽奖</li><li>task： 做任务</li><li>invitation： 邀请</li><li>claim_red_packet： 领红包</li><li>browse： 浏览</li></ul><p>自定义事件可与RCE约定后进行风险评估</p>
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// <p>事件的发生时间</p><p>参数格式：符合ISO 8601标准的带UTC时区的毫秒级时间</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>用户当前会话 ID， 用于关联用户登录前后的动作，如果没有传UserId，则SessionId必传，如缺失则可填充空字符串</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址（IPv4或IPv6）</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>事件详情，根据您输入的事件码传入对应的事件信息</p>
	EventDetail *EventDetail `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`

	// <p>用户在您系统中的唯一ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>用户邮箱</p>
	UserEmail *string `json:"UserEmail,omitnil,omitempty" name:"UserEmail"`

	// <p>用户提供的联系方式</p><p>参数格式：符合E.164标准的带“+”、地区编码和号码的格式</p>
	UserPhone *string `json:"UserPhone,omitnil,omitempty" name:"UserPhone"`

	// <p>web浏览器相关信息，若您已集成我们的设备指纹SDK，则无需传入此字段</p>
	Browser *Browser `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>应用程序、操作系统和移动设备详细信息，若您已集成我们的设备指纹SDK，则无需传入此字段</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`

	// <p>数据授权信息，国内地域必填</p>
	DataAuthorization *DataAuthorization `json:"DataAuthorization,omitnil,omitempty" name:"DataAuthorization"`

	// <p>手机号码加密方式，国内地域必填</p><p>枚举值：</p><ul><li>md5： md5加密</li><li>plain： 明文</li></ul>
	UserPhoneEncrypt *string `json:"UserPhoneEncrypt,omitnil,omitempty" name:"UserPhoneEncrypt"`

	// <p>微信开放账号</p>
	WeChatOpenId *string `json:"WeChatOpenId,omitnil,omitempty" name:"WeChatOpenId"`

	// <p>QQ开放账号</p>
	QQOpenId *string `json:"QQOpenId,omitnil,omitempty" name:"QQOpenId"`

	// <p>QQ应用ID，当传入QQ开放账号时，该字段必填，QQ分配给网站或应用的AppId，用来唯一标识网站或应用</p>
	QQAppId *string `json:"QQAppId,omitnil,omitempty" name:"QQAppId"`
}

func (r *ReportEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventCode")
	delete(f, "EventTime")
	delete(f, "SessionId")
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	delete(f, "EventDetail")
	delete(f, "UserId")
	delete(f, "UserEmail")
	delete(f, "UserPhone")
	delete(f, "Browser")
	delete(f, "App")
	delete(f, "DataAuthorization")
	delete(f, "UserPhoneEncrypt")
	delete(f, "WeChatOpenId")
	delete(f, "QQOpenId")
	delete(f, "QQAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportEventResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReportEventResponse struct {
	*tchttp.BaseResponse
	Response *ReportEventResponseParams `json:"Response"`
}

func (r *ReportEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Result struct {
	// <p>实际是否完成状态</p><p>枚举值：</p><ul><li>success： 成功</li><li>failure： 失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>失败原因</p>
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`
}

type RiskLabel struct {
	// <p>风险ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>风险描述</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type SMSEvent struct {
	// <p>用户基础信息</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>本次短信发送标识 ID</p>
	SMSId *string `json:"SMSId,omitnil,omitempty" name:"SMSId"`

	// <p>用户实际完成验证码时间</p><p>参数格式：符合ISO 8601标准的带UTC时区的毫秒级时间</p>
	ReceivedTime *string `json:"ReceivedTime,omitnil,omitempty" name:"ReceivedTime"`

	// <p>记录用户收到短信的动作</p><ul><li>no_action：用户无动作</li><li>safe：用户确认本人操作</li><li>compromised：用户反馈为第三方操作</li></ul>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// <p>短信回执结果</p>
	SMSResult *Result `json:"SMSResult,omitnil,omitempty" name:"SMSResult"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type ScanCodeEvent struct {
	// <p>营销活动码</p>
	PromotionCode *PromotionCode `json:"PromotionCode,omitnil,omitempty" name:"PromotionCode"`

	// <p>营销活动ID</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>营销活动名称</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>营销活动描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>邀请人ID</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组， 示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Score struct {
	// <p>风险分值，范围[1, 1000]，分值越大，风险越高</p>
	RiskScore *int64 `json:"RiskScore,omitnil,omitempty" name:"RiskScore"`

	// <p>风险标签</p>
	RiskLabels []*RiskLabel `json:"RiskLabels,omitnil,omitempty" name:"RiskLabels"`
}

type SecurityVerificationEvent struct {
	// <p>安全验证所处的事件类型</p><p>枚举值：</p><ul><li>register： 注册</li><li>login： 登录</li><li>modify_account： 修改账号</li><li>modify_password： 修改密码</li><li>create_order： 创建订单</li><li>transaction： 交易支付</li><li>modify_order： 修改订单</li><li>withdraw： 提现</li><li>add_promotion： 参加营销活动</li><li>redeem： 兑奖</li></ul>
	VerificationEvent *string `json:"VerificationEvent,omitnil,omitempty" name:"VerificationEvent"`

	// <p>安全验证类型，sms-短信、phone_call-电话、email-邮件、captcha-验证码、shared_knowledge-共享知识、face-人脸、fingerprint-指纹等</p>
	VerificationType *string `json:"VerificationType,omitnil,omitempty" name:"VerificationType"`

	// <p>安全验证的内容，例如：用于验证的电话号码、邮件、验证码或者问题，当安全验证类型是sms、phone_call、email、captcha、shared_knowledge时输入</p>
	VerificationContent *string `json:"VerificationContent,omitnil,omitempty" name:"VerificationContent"`

	// <p>安全验证结果</p>
	VerificationResult *Result `json:"VerificationResult,omitnil,omitempty" name:"VerificationResult"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type TaskEvent struct {
	// <p>营销活动ID</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>营销活动名称</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>营销活动描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>邀请人ID</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>任务ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务名称</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>任务类型，如签到打卡、观看广告、累计步数等</p>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// <p>任务完成耗时</p><p>单位：毫秒</p>
	TaskCostTime *int64 `json:"TaskCostTime,omitnil,omitempty" name:"TaskCostTime"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type TransactionEvent struct {
	// <p>交易唯一标识</p>
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`

	// <p>您系统中的订单 ID，当一笔交易关联多个订单（合并支付）时请输入所有订单ID</p>
	OrderId []*string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// <p>交易金额</p>
	PaymentAmount *Amount `json:"PaymentAmount,omitnil,omitempty" name:"PaymentAmount"`

	// <p>支付方式，支持多种支付方式</p>
	PaymentMethod *PaymentMethod `json:"PaymentMethod,omitnil,omitempty" name:"PaymentMethod"`

	// <p>交易类型</p><p>枚举值：</p><ul><li>sale： 一次性完成授权与扣款（最常见）</li><li>authorize： 仅授权（冻结金额）</li><li>capture： 执行扣款（在授权后）</li><li>void： 取消待处理的授权或扣款</li><li>refund： 退款（部分或全部）</li><li>deposit： 向账户存款</li><li>withdrawal： 从账户提现</li><li>transfer： 账户间转账</li><li>buy： 购买资产（如加密货币）</li><li>sell： 出售资产</li><li>send： 发送资金/资产（如跨钱包转账）</li><li>receive： 接收资金/资产</li></ul><p>默认值：sale</p>
	TransactionType *string `json:"TransactionType,omitnil,omitempty" name:"TransactionType"`

	// <p>账单信息</p>
	Billing *Billing `json:"Billing,omitnil,omitempty" name:"Billing"`

	// <p>物流信息</p>
	Delivery *Delivery `json:"Delivery,omitnil,omitempty" name:"Delivery"`

	// <p>商家信息</p>
	Merchant *Merchant `json:"Merchant,omitnil,omitempty" name:"Merchant"`

	// <p>支付结果</p>
	PaymentResult *PaymentResult `json:"PaymentResult,omitnil,omitempty" name:"PaymentResult"`

	// <p>接收方的用户ID，适用于 transfer 交易类型</p>
	TransferRecipientUserId *string `json:"TransferRecipientUserId,omitnil,omitempty" name:"TransferRecipientUserId"`

	// <p>发送方的物理地址，适用于 transfer 交易类型</p>
	TransferSentAddress *Address `json:"TransferSentAddress,omitnil,omitempty" name:"TransferSentAddress"`

	// <p>接收方的物理地址，适用于 transfer 交易类型</p>
	TransferReceivedAddress *Address `json:"TransferReceivedAddress,omitnil,omitempty" name:"TransferReceivedAddress"`

	// <p>数字订单列表</p>
	DigitalOrders []*DigitalOrder `json:"DigitalOrders,omitnil,omitempty" name:"DigitalOrders"`

	// <p>接收加密货币的钱包</p>
	ReceiverWallet *Wallet `json:"ReceiverWallet,omitnil,omitempty" name:"ReceiverWallet"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type User struct {
	// <p>用户等级</p>
	UserLevel *string `json:"UserLevel,omitnil,omitempty" name:"UserLevel"`

	// <p>用户积分</p>
	UserPoint *CreditPoint `json:"UserPoint,omitnil,omitempty" name:"UserPoint"`

	// <p>用户类型</p>
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type Wallet struct {
	// <p>钱包类型</p><p>枚举值：</p><ul><li>crypto： 加密货币</li><li>digital： 数字货币</li><li>fiat： 法币</li></ul>
	WalletType *string `json:"WalletType,omitnil,omitempty" name:"WalletType"`

	// <p>钱包地址，通常为钱包的唯一标识</p>
	WalletAddress *string `json:"WalletAddress,omitnil,omitempty" name:"WalletAddress"`

	// <p>钱包归属人姓名</p>
	WalletHolderName *string `json:"WalletHolderName,omitnil,omitempty" name:"WalletHolderName"`

	// <p>钱包供应商，wechat、alipay、paypal等</p>
	WalletProvider *string `json:"WalletProvider,omitnil,omitempty" name:"WalletProvider"`
}

type WithdrawEvent struct {
	// <p>提现金额</p>
	Amount *Amount `json:"Amount,omitnil,omitempty" name:"Amount"`

	// <p>提现方式</p><p>枚举值：</p><ul><li>card： 银行卡</li><li>wallet： 电子钱包</li></ul>
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>提现银行卡，当提现方式是card时必填</p>
	Card *Card `json:"Card,omitnil,omitempty" name:"Card"`

	// <p>提现数字钱包，当提现方式是wallet时必填</p>
	Wallet *Wallet `json:"Wallet,omitnil,omitempty" name:"Wallet"`

	// <p>提现结果</p>
	Result *Result `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>与RCE约定的定制化信息，为K:V 格式的对象数组，示例：[{&quot;Key&quot;: &quot;ApproverName&quot;, &quot;Value&quot;: &quot;bob&quot;},{&quot;Key&quot;:&quot;ApproverPhone&quot;,&quot;Value&quot;: &quot;+86131****5678&quot;}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}