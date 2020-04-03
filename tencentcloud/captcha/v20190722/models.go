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

package v20190722

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CaptchaOperDataInterceptUnit struct {

	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// 停止验证数量
	AllStopCnt *float64 `json:"AllStopCnt,omitempty" name:"AllStopCnt"`

	// 图片停止加载数量
	PicStopCnt *float64 `json:"PicStopCnt,omitempty" name:"PicStopCnt"`

	// 策略拦截数量
	StrategyStopCnt *float64 `json:"StrategyStopCnt,omitempty" name:"StrategyStopCnt"`
}

type CaptchaOperDataLoadTimeUnit struct {

	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// Market加载时间
	MarketLoadTime *float64 `json:"MarketLoadTime,omitempty" name:"MarketLoadTime"`

	// AppId加载时间
	AppIdLoadTime *float64 `json:"AppIdLoadTime,omitempty" name:"AppIdLoadTime"`
}

type CaptchaOperDataRes struct {

	// 验证码加载耗时数据返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperDataLoadTimeUnitArray []*CaptchaOperDataLoadTimeUnit `json:"OperDataLoadTimeUnitArray,omitempty" name:"OperDataLoadTimeUnitArray" list`

	// 验证码拦截情况数据返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperDataInterceptUnitArray []*CaptchaOperDataInterceptUnit `json:"OperDataInterceptUnitArray,omitempty" name:"OperDataInterceptUnitArray" list`

	// 验证码尝试次数数据返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperDataTryTimesUnitArray []*CaptchaOperDataTryTimesUnit `json:"OperDataTryTimesUnitArray,omitempty" name:"OperDataTryTimesUnitArray" list`

	// 验证码尝试次数分布数据返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperDataTryTimesDistributeUnitArray []*CaptchaOperDataTryTimesDistributeUnit `json:"OperDataTryTimesDistributeUnitArray,omitempty" name:"OperDataTryTimesDistributeUnitArray" list`
}

type CaptchaOperDataTryTimesDistributeUnit struct {

	// 尝试次数
	TryCount *int64 `json:"TryCount,omitempty" name:"TryCount"`

	// 用户请求数量
	UserCount *int64 `json:"UserCount,omitempty" name:"UserCount"`
}

type CaptchaOperDataTryTimesUnit struct {

	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// 平均尝试次数
	CntPerPass []*float64 `json:"CntPerPass,omitempty" name:"CntPerPass" list`

	// market平均尝试次数
	MarketCntPerPass *float64 `json:"MarketCntPerPass,omitempty" name:"MarketCntPerPass"`
}

type CaptchaQueryData struct {

	// 数量
	Cnt *int64 `json:"Cnt,omitempty" name:"Cnt"`

	// 时间
	Date *string `json:"Date,omitempty" name:"Date"`
}

type CaptchaUserAllAppId struct {

	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 注册应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 腾讯云APPID
	TcAppId *int64 `json:"TcAppId,omitempty" name:"TcAppId"`
}

type DescribeCaptchaAppIdInfoRequest struct {
	*tchttp.BaseRequest

	// 验证码应用注册APPID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`
}

func (r *DescribeCaptchaAppIdInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaAppIdInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 界面风格
		SchemeColor *string `json:"SchemeColor,omitempty" name:"SchemeColor"`

		// 语言
		Language *int64 `json:"Language,omitempty" name:"Language"`

		// 场景
		SceneType *int64 `json:"SceneType,omitempty" name:"SceneType"`

		// 防控风险等级
		EvilInterceptGrade *int64 `json:"EvilInterceptGrade,omitempty" name:"EvilInterceptGrade"`

		// 智能验证
		SmartVerify *int64 `json:"SmartVerify,omitempty" name:"SmartVerify"`

		// 智能引擎
		SmartEngine *int64 `json:"SmartEngine,omitempty" name:"SmartEngine"`

		// 验证码类型
		CapType *int64 `json:"CapType,omitempty" name:"CapType"`

		// 应用名称
		AppName *string `json:"AppName,omitempty" name:"AppName"`

		// 域名限制
		DomainLimit *string `json:"DomainLimit,omitempty" name:"DomainLimit"`

		// 邮件告警
	// 注意：此字段可能返回 null，表示取不到有效值。
		MailAlarm []*string `json:"MailAlarm,omitempty" name:"MailAlarm" list`

		// 流量控制
		TrafficThreshold *int64 `json:"TrafficThreshold,omitempty" name:"TrafficThreshold"`

		// 加密key
		EncryptKey *string `json:"EncryptKey,omitempty" name:"EncryptKey"`

		// 是否全屏
		TopFullScreen *int64 `json:"TopFullScreen,omitempty" name:"TopFullScreen"`

		// 成功返回0 其它失败
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 返回操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCaptchaAppIdInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaAppIdInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaDataRequest struct {
	*tchttp.BaseRequest

	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间
	End *int64 `json:"End,omitempty" name:"End"`

	// 查询类型
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeCaptchaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码 0 成功 其它失败
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 数据数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*CaptchaQueryData `json:"Data,omitempty" name:"Data" list`

		// 返回信息描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCaptchaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaDataSumRequest struct {
	*tchttp.BaseRequest

	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间
	End *int64 `json:"End,omitempty" name:"End"`
}

func (r *DescribeCaptchaDataSumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaDataSumRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaDataSumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求总量
		GetSum *int64 `json:"GetSum,omitempty" name:"GetSum"`

		// 请求验证成功量
		VfySuccSum *int64 `json:"VfySuccSum,omitempty" name:"VfySuccSum"`

		// 请求验证量
		VfySum *int64 `json:"VfySum,omitempty" name:"VfySum"`

		// 拦截攻击量
		AttackSum *int64 `json:"AttackSum,omitempty" name:"AttackSum"`

		// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// 成功返回0  其它失败
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 票据校验量
		CheckTicketSum *int64 `json:"CheckTicketSum,omitempty" name:"CheckTicketSum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCaptchaDataSumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaDataSumResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaOperDataRequest struct {
	*tchttp.BaseRequest

	// 验证码应用ID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *uint64 `json:"Start,omitempty" name:"Start"`

	// 查询类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeCaptchaOperDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaOperDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaOperDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功返回 0 其它失败
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// 用户操作数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *CaptchaOperDataRes `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCaptchaOperDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaOperDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaResultRequest struct {
	*tchttp.BaseRequest

	// 固定填值：9。可在控制台配置不同验证码类型。
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// 验证码返回给用户的票据
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// 用户操作来源的外网 IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 验证票据需要的随机字符串
	Randstr *string `json:"Randstr,omitempty" name:"Randstr"`

	// 验证码应用ID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 用于服务器端校验验证码票据的验证密钥，请妥善保密，请勿泄露给第三方
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// 业务 ID，网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 场景 ID，网站或应用的业务下有多个场景使用此服务，通过此 ID 区分统计数据
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// mac 地址或设备唯一标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机设备号
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 是否返回前端获取验证码时间，取值1：需要返回
	NeedGetCaptchaTime *int64 `json:"NeedGetCaptchaTime,omitempty" name:"NeedGetCaptchaTime"`
}

func (r *DescribeCaptchaResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 1	OK	验证通过
	// 6	user code len error	验证码长度不匹配
	// 7	captcha no match	验证码答案不匹配/Randstr参数不匹配
	// 8	verify timeout	验证码签名超时
	// 9	Sequnce repeat	验证码签名重放
	// 10	Sequnce invalid	验证码签名序列
	// 11	Cookie invalid	验证码cooking信息不合法
	// 12	sig len error	签名长度错误
	// 13	verify ip no match	ip不匹配
	// 15	decrypt fail	验证码签名解密失败
	// 16	appid no match	验证码强校验appid错误
	// 17	cmd no much	验证码系统命令不匹配
	// 18	uin no match	号码不匹配
	// 19	seq redirect	重定向验证
	// 20	opt no vcode	操作使用pt免验证码校验错误
	// 21	diff	差别，验证错误
	// 22	captcha type not match	验证码类型与拉取时不一致
	// 23	verify type error	验证类型错误
	// 24	invalid pkg	非法请求包
	// 25	bad visitor	策略拦截
	// 26	system busy	系统内部错误
	// 100	param err	appsecretkey 参数校验错误
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 状态描述及验证错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// [0,100]，恶意等级
	// 注意：此字段可能返回 null，表示取不到有效值。
		EvilLevel *int64 `json:"EvilLevel,omitempty" name:"EvilLevel"`

		// 前端获取验证码时间，时间戳格式
	// 注意：此字段可能返回 null，表示取不到有效值。
		GetCaptchaTime *int64 `json:"GetCaptchaTime,omitempty" name:"GetCaptchaTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCaptchaResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaUserAllAppIdRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCaptchaUserAllAppIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaUserAllAppIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaUserAllAppIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户注册的所有Appid和应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*CaptchaUserAllAppId `json:"Data,omitempty" name:"Data" list`

		// 成功返回 0  其它失败
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 返回操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCaptchaUserAllAppIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaUserAllAppIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCaptchaAppIdInfoRequest struct {
	*tchttp.BaseRequest

	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 域名限制
	DomainLimit *string `json:"DomainLimit,omitempty" name:"DomainLimit"`

	// 场景类型
	SceneType *int64 `json:"SceneType,omitempty" name:"SceneType"`

	// 验证码类型
	CapType *int64 `json:"CapType,omitempty" name:"CapType"`

	// 风险级别
	EvilInterceptGrade *int64 `json:"EvilInterceptGrade,omitempty" name:"EvilInterceptGrade"`

	// 智能检测
	SmartVerify *int64 `json:"SmartVerify,omitempty" name:"SmartVerify"`

	// 开启智能引擎
	SmartEngine *int64 `json:"SmartEngine,omitempty" name:"SmartEngine"`

	// web风格
	SchemeColor *string `json:"SchemeColor,omitempty" name:"SchemeColor"`

	// 语言
	CaptchaLanguage *int64 `json:"CaptchaLanguage,omitempty" name:"CaptchaLanguage"`

	// 告警邮箱
	MailAlarm *string `json:"MailAlarm,omitempty" name:"MailAlarm"`

	// 是否全屏
	TopFullScreen *int64 `json:"TopFullScreen,omitempty" name:"TopFullScreen"`

	// 流量限制
	TrafficThreshold *int64 `json:"TrafficThreshold,omitempty" name:"TrafficThreshold"`
}

func (r *UpdateCaptchaAppIdInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCaptchaAppIdInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCaptchaAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码 0 成功，其它失败
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 返回操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCaptchaAppIdInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCaptchaAppIdInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
